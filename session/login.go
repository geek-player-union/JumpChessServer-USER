package session

import (
	"github.com/geek-player-union/JumpChessServer-USER/database"
)

func handleLogin(d interface{}) (Instruction, int64) {
	data, ok := d.(struct {
		Account string
		Code    string
	})

	if !ok {
		return Instruction{
			Command: "SYNTAX_ERROR",
			Message: "invalid parameters for command 'login'",
		}, -1
	}

	user, reason := database.CheckUserLogin(data.Account, data.Code)
	return Instruction{
		Command: reason,
		Data:    user,
	}, user.ID
}
