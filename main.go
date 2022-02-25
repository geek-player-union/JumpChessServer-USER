package main

import (
	"github.com/geek-player-union/JumpChessServer-USER/config"
	"github.com/geek-player-union/JumpChessServer-USER/database"
	"github.com/geek-player-union/JumpChessServer-USER/server"
)

func main() {
	database.Init()
	server.Start(config.Port)
}
