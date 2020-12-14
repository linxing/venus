package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	cachePrefix = "captcha_"
)

type Captcha struct {
	Value     string
	ExpiredAt int64
	Store     func() redis.Conn
}

func NewCaptcha(conn func() redis.Conn, expiredAt int64) *Captcha {
	return &Captcha{
		Store:     conn,
		ExpiredAt: expiredAt,
	}
}

func genCaptchaKey(captcha string) string {
	return cachePrefix + captcha
}

func (c *Captcha) GetCaptcha(urlPrefix string) (string, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	captcha := r.Intn(1000) + 1000

	captchaStr := fmt.Sprintf("%d", captcha)

	key := genCaptchaKey(captchaStr)

	conn := c.Store()
	defer conn.Close()

	_, err := conn.Do("SET", key, urlPrefix, "EX", c.ExpiredAt)
	if err != nil {
		return "", err
	}
	return captchaStr, nil
}

func (c *Captcha) VerifyCaptcha(urlPrefix, captcha string) (bool, error) {

	conn := c.Store()
	defer conn.Close()

	key := genCaptchaKey(captcha)

	v, err := redis.String(conn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return false, err
	}

	if v == urlPrefix {
		return true, nil
	}

	return false, nil
}
