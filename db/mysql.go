package db

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"log"
	"time"
)

func Init(username, password, host, port, database string, syncDB, debug bool) {
	err := orm.RegisterDriver("default", orm.DRMySQL)
	if err != nil {
		log.Panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", username, password, host, port, database)
	err = orm.RegisterDataBase("default", "mysql", dsn,
		orm.MaxIdleConnections(10),
		orm.MaxOpenConnections(30),
		orm.ConnMaxLifetime(3600*time.Second))
	if err != nil {
		log.Panic(err)
	}
	err = orm.SetDataBaseTZ("default", time.Local)
	if err != nil {
		log.Panic(err)
	}
	if syncDB {
		err = orm.RunSyncdb("default", false, true)
		if err != nil {
			log.Panic(err)
		}
	}
	if debug {
		orm.Debug = debug
	}
}
