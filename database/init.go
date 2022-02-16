package database

import (
	"fmt"
	"github.com/geek-player-union/JumpChessServer-USER/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Init() {
	dbConnectStr := fmt.Sprintf("%s:%s@/%s?charset=utf8", config.MysqlUsername, config.MysqlPassword, config.MysqlDBName)

	var err error
	engine, err = xorm.NewEngine("mysql", dbConnectStr)
	if err != nil {
		panic(err)
	}

	err = engine.Sync2(User{})
	if err != nil {
		panic(err)
	}
}
