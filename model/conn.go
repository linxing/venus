package model

import (
	"time"

	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

//If your MySQL server uses idle timeouts,
//or is actively pruning connections,
//you don’t need to call (*DB).SetConnMaxLifetime in your production services.
//It’s no longer needed as the driver can now gracefully detect and retry stale connections.
//Setting a max lifetime for connections simply causes unnecessary churn by killing and re-opening healthy connections.

//A good pattern to manage high-throughput access to MySQL is
//configuring your sql.(*DB) pool ((*DB).SetMaxIdleConns and (*DB).SetMaxOpenConns) with values that support your peak-hour traffic for the service,
//and making sure that your MySQL server is actively pruning idle connections during off-hours.
//These pruned connections are detected by the MySQL driver and re-created when necessary.
//Ref: https://github.blog/2020-05-20-three-bugs-in-the-go-mysql-driver/

func NewEngine(driver, dsn string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		return nil, err
	}

	engine.SetConnMaxLifetime(5 * time.Minute)
	engine.SetMaxOpenConns(64)
	engine.SetMaxIdleConns(32)
	//engine.TZLocation = time.Local
	//engine.DatabaseTZ = time.Local

	return engine, nil
}
