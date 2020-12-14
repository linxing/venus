package envtest

import (
	"io"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"

	"venus/env/global"
	jobintf "venus/job/intf"
	"venus/model"
	"venus/setting"
)

type mockenv struct {
	testing.TB
	closers []io.Closer
}

func (me *mockenv) Close() {
	for _, closer := range me.closers {
		if err := closer.Close(); err != nil {
			me.Error(err)
		}
	}
}

func NewMockEnv(t testing.TB, opts ...MockOption) *mockenv {
	me := &mockenv{t, make([]io.Closer, 0)}
	opts = append([]MockOption{mockConfigOption()}, opts...)
	for _, opt := range opts {
		opt(me)
	}

	return me
}

type MockOption func(*mockenv)

func mockConfigOption() MockOption {
	return func(me *mockenv) {
		if err := setting.Init(); err != nil {
			me.Fatal()
		}
		global.Config = *setting.Config
	}
}

func MockDBOption(beans ...interface{}) MockOption {
	return func(me *mockenv) {
		var (
			dbs = make([]io.Closer, 0)
			err error
		)

		defer func() {
			if err == nil {
				me.closers = append(me.closers, dbs...)
				return
			}
			for _, db := range dbs {
				err := db.Close()
				if err != nil {
					return
				}
			}
			me.Fatal(err)
		}()

		venusDB, err := model.MockNewDB("venus", beans...)
		if err != nil {
			return
		}

		model.StubEngine(venusDB.Engine())

		dbs = append(dbs, venusDB)
	}
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,

		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(addr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func MockRedisOption() MockOption {
	return func(me *mockenv) {

		addr := global.Config.RedisAddr
		if addr == "" {
			addr = "redis://127.0.0.1:6379/2"
		}

		pool := newPool(addr)
		conn := func() redis.Conn {
			return pool.Get()
		}

		func() {
			connf := conn()
			defer connf.Close()

			ticker := time.NewTicker(10 * time.Millisecond)
			defer ticker.Stop()

			for {
				for range ticker.C {
					reply, _ := redis.String(connf.Do("SET", "redisLock", 1, "EX", 3600, "NX"))
					if reply == "OK" {
						return
					}
				}
			}
		}()

		global.Redis.Pool = pool
		global.Redis.Conn = conn
		global.Tasker = jobintf.NewTasker(pool, global.Config.JobNamespace)

		me.closers = append(me.closers, newRedisCloser(pool, func() error {
			conn := pool.Get()
			defer conn.Close()

			keys, err := redis.Values(conn.Do("KEYS", "*"))
			if err != nil {
				return err
			}
			if len(keys) > 0 {
				_, err = conn.Do("DEL", keys...)
			}

			return err
		}))
	}
}

type redisPoolCloser struct {
	pool   *redis.Pool
	closer func() error
}

func newRedisCloser(p *redis.Pool, closer func() error) io.Closer {
	return &redisPoolCloser{p, closer}
}

func (rp *redisPoolCloser) Close() error {
	defer rp.pool.Close()

	return rp.closer()
}
