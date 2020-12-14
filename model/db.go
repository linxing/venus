package model

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const (
	DefaultPageSize int = 20
	MaxPageSize     int = 1000
)

var (
	engine *xorm.Engine
)

func Init(driver string, dsn string) error {

	var err error

	engine, err = NewEngine(driver, dsn)
	if err != nil {
		return err
	}

	// for debug
	// engine.ShowSQL(true)

	return engine.Ping()
}

func GetSession(ctx context.Context) *xorm.Session {
	return engine.NewSession().Context(ctx)
}

func GetEngine() *xorm.Engine {
	return engine
}

func StubEngine(mockedEngine *xorm.Engine) {
	engine = mockedEngine
}
