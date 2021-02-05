package global

import (
	"time"

	"github.com/gomodule/redigo/redis"

	jobintf "venus/job/intf"
	"venus/pkg/venus"
	"venus/setting"
)

const (
	DateLayout = "2006-01-02 15:04:05"
)

var (
	Config setting.Setting

	Redis struct {
		Pool *redis.Pool
		Conn func() redis.Conn
	}

	Tasker jobintf.TaskerInterface

	VenusService venus.VenusServiceClient

	TimeLocation = time.UTC
)
