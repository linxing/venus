package consumer

import (
	"sync"
	"sync/atomic"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ConsumerManager struct {
	client          sarama.Client
	consumer        sarama.Consumer
	offsetManager   sarama.OffsetManager
	poms            []sarama.PartitionOffsetManager
	messageConsumer func(*sarama.ConsumerMessage) error
	consumerGroup   string
	topic           string
	stopped         atomic.Value
}

func NewConsumerManager(brokers []string, topic string, consumerGroup string, messageConsumer func(*sarama.ConsumerMessage) error) (*ConsumerManager, error) {

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	kafkaClient, err := sarama.NewClient(brokers, kafkaConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	kafkaConsumer, err := sarama.NewConsumerFromClient(kafkaClient)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	kafkaOffsetManager, err := sarama.NewOffsetManagerFromClient(consumerGroup, kafkaClient)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &ConsumerManager{
		client:          kafkaClient,
		consumer:        kafkaConsumer,
		offsetManager:   kafkaOffsetManager,
		messageConsumer: messageConsumer,
		consumerGroup:   consumerGroup,
		topic:           topic,
	}, nil
}

func (cm *ConsumerManager) Run() error {

	waitGroup := sync.WaitGroup{}

	partitions, err := cm.consumer.Partitions(cm.topic)
	if err != nil {
		return errors.WithStack(err)
	}

	poms := make([]sarama.PartitionOffsetManager, 0, len(partitions))

	for _, partition := range partitions {
		waitGroup.Add(1)

		pom, err := cm.offsetManager.ManagePartition(cm.topic, partition)
		if err != nil {
			return errors.WithStack(err)
		}

		poms = append(poms, pom)
		go cm.consumePartition(pom, partition, &waitGroup)
	}

	cm.poms = poms

	waitGroup.Wait()

	return nil
}

func (cm *ConsumerManager) consumePartition(pom sarama.PartitionOffsetManager, partition int32, wg *sync.WaitGroup) {

	defer wg.Done()

	offset, metadata := pom.NextOffset()

	pc, err := cm.consumer.ConsumePartition(cm.topic, partition, offset)
	if err != nil {
		if err != sarama.ErrOffsetOutOfRange {
			logrus.Fatal(err)
		} else {
			offset = sarama.OffsetNewest
			pc, err = cm.consumer.ConsumePartition(cm.topic, partition, offset)
			if err != nil {
				logrus.Fatal(err)
			}
		}
	}

	defer pc.Close()

	for cm.stopped.Load() == nil {
		select {
		case err := <-pc.Errors():
			logrus.Fatal(err)
		case msg := <-pc.Messages():
			err := cm.messageConsumer(msg)
			if err != nil {
				logrus.Fatal(err)
			} else {
				pom.MarkOffset(msg.Offset+1, metadata)
			}
		}
	}
}

func (cm *ConsumerManager) Close() {

	cm.stopped.Store(true)

	for _, poms := range cm.poms {
		if err := poms.Close(); err != nil {
			logrus.Fatal(err)
		}
	}

	if cm.offsetManager != nil {
		if err := cm.offsetManager.Close(); err != nil {
			logrus.Fatal(err)
		}
	}

	if cm.consumer != nil {
		if err := cm.consumer.Close(); err != nil {
			logrus.Fatal(err)
		}
	}

	if cm.client != nil {
		if err := cm.client.Close(); err != nil {
			logrus.Fatal(err)
		}
	}
}
