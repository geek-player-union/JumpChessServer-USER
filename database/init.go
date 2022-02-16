package database

import (
	"fmt"
	"github.com/geek-player-union/JumpChessServer-USER/config"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Init() {
	dbConnectStr := fmt.Sprintf("%s:%s", config.MysqlUsername, config.MysqlPassword)

	var err error
	engine, err = xorm.NewEngine(config.MysqlDBName, dbConnectStr)
	if err != nil {
		panic(err)
	}

	err = engine.Sync2(User{})
	if err != nil {
		panic(err)
	}
}
