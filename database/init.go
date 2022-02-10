package database

import (
	"fmt"
	"github.com/geek-player-union/JumpChessServer-USER/config"
)

var engine *xorm.Engine

func Init() {
	dbConnectStr := fmt.Sprintf("%s:%s", config.MysqlUsername, config.MysqlPassword)

	engine = xorm.NewEngine("mysql", dbConnectStr)
}
