package setting

import (
	"log"
	"path"
	"runtime"

	"github.com/go-ini/ini"
)

type Setting struct {
	ServiceHTTPAddr          string
	ServicePort              int
	ServiceGrpcAddr          string
	ServiceGrpcPort          int
	ServiceEnv               string
	ServiceMaxRequestsPerSec int

	DatabaseDSN    string
	DatabaseDriver string

	RedisAddr   string
	RedisPrefix string

	JWTSecret      string
	JWTTokenExpSec int

	JobNamespace string

	ZipkinReporter     string
	ZipkinEndPointHost string
	ZipkinEndpointName string

	TimeZone string

	KafkaConsumerTopic   string
	KafkaConsumerGroup   string
	KafkaConsumerBrokers []string

	MetricsHTTPAddr string
	MetricsHTTPPort int

	DefaultRPCKeepAliveSec int
	DefaultRPCTimeout      int

	// grpc venus service
	VenusServiceHostAddr     string
	VenusServiceKeepAliveSec int
	VenusServiceTimeout      int
}

var Config = &Setting{ServicePort: 80}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

func Init(config ...string) error {

	filePath := getCurrentPath() + "/conf.test.ini"

	if len(config) == 1 {
		filePath = config[0]
	}

	cfg, err := ini.Load(filePath)
	if err != nil {
		log.Fatalf("setting setup, fail to parse '%s': %v", filePath, err)
		return err
	}

	err = cfg.Section("Setting").MapTo(Config)
	if err != nil {
		log.Fatalf("server setup, fail to parse '%s': %v", filePath, err)
		return err
	}

	return nil
}
