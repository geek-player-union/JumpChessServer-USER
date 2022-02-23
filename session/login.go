package session

import (
	"github.com/geek-player-union/JumpChessServer-USER/database"
)

func handleLogin(d interface{}) (string, int64) {
	data, ok := d.(struct {
		Account string
		Code    string
	})

	if !ok {
		return "SYNTAX_ERROR", -1
	}

	return database.CheckUserLogin(data.Account, data.Code)
}
