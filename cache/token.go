package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	"venus/env/global"
)

type (
	tokenStatic struct{}
)

var (
	TokenStatic = new(tokenStatic)
)

func GetTokenKey(k string) string {
	return fmt.Sprintf("%s:%s", global.Config.RedisPrefix, k)
}

func (*tokenStatic) SetData(k string, v string, expSec int64) (bool, error) {

	conn := global.Redis.Conn()
	defer conn.Close()

	resp, err := conn.Do("SET", k, v, "EX", expSec)
	return resp != nil, errors.WithStack(err)
}

func (*tokenStatic) GetData(k string) (string, error) {

	conn := global.Redis.Conn()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", k))
	if err != nil && err != redis.ErrNil {
		return "", errors.WithStack(err)
	}

	return s, nil
}
