package env

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"venus/env/global"
	jobintf "venus/job/intf"
	"venus/model"
)

// connection database
func configDatabase() error {
	if err := model.Init(global.Config.DatabaseDriver, global.Config.DatabaseDSN); err != nil {
		return err
	}
	return nil
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,

		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(global.Config.RedisAddr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// connection redis
func configRedis() error {

	pool := newPool()

	conn := pool.Get()

	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return err
	}

	connF := func() redis.Conn {
		return pool.Get()
	}

	global.Redis.Pool = pool
	global.Redis.Conn = connF
	global.Tasker = jobintf.NewTasker(pool, global.Config.JobNamespace)

	return nil
}

func timeLocation() error {
	loc, err := time.LoadLocation(global.Config.TimeZone)
	if err != nil {
		return err
	}

	global.TimeLocation = loc
	return nil
}
