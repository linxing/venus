package model

import (
	xormtest "github.com/linxing/xormtest"
)

var (
	initBeans = []interface{}{
		User{},
	}
)

func MockNewDB(name string, beans ...interface{}) (*xormtest.DB, error) {
	if beans == nil {
		beans = initBeans
	}
	return xormtest.NewDB("mysql", "root:root@tcp(localhost:3306)/", name, beans...)
}

func NewDB(tables ...interface{}) (*xormtest.DB, error) {
	db, err := MockNewDB("venus_test", tables...)
	if err != nil {
		return nil, err
	}

	StubEngine(db.Engine())
	return db, nil
}
