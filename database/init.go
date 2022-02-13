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
	engine, err = xorm.NewEngine("mysql", dbConnectStr)

	if err != nil {
		return
	}
}
