package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"venus/consumer"
	"venus/env"
	"venus/env/global"
	"venus/setting"
)

var (
	BuildTime = "20200101"
	GitTag    = "v0.0.1"
)

func waitForTeardown(cm *consumer.ConsumerManager) {

	sigCh := make(chan os.Signal, 1)

	signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)

	<-sigCh

	cm.Close()
}

func startPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", global.Config.MetricsHTTPAddr, global.Config.MetricsHTTPPort), nil)
	if err != nil {
		logrus.Fatalf("fail to serve prometheus: %s", err)
	}
}

func init() {

	if os.Args[1] == "-v" {
		fmt.Println("build at: " + BuildTime)
		fmt.Println("git tag: " + GitTag)
		os.Exit(0)
	}

	conf := flag.String("conf.ini", "setting/conf.test.ini", "config file")
	flag.Parse()

	if err := setting.Init(*conf); err != nil {
		panic(err)
	}

	global.Config = *setting.Config

	if err := env.Configure(); err != nil {
		log.Printf("[info] init env fail %+v", err)
		return
	}
}

func venusConsumerEvent(msg *sarama.ConsumerMessage) error {
	// TODO complete the envent handler
	return nil
}

func main() {

	go startPrometheus()

	newConsumer, err := consumer.NewConsumerManager(global.Config.KafkaConsumerBrokers, global.Config.KafkaConsumerTopic, global.Config.KafkaConsumerGroup, venusConsumerEvent)
	if err != nil {
		panic(err)
	}

	go waitForTeardown(newConsumer)

	err = newConsumer.Run()
	if err != nil {
		panic(err)
	}
}
