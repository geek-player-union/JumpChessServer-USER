package session

import (
	"github.com/geek-player-union/JumpChessServer-USER/database"
)

func handleLogin(commands []string) Instruction {
	if len(commands) != 2 {
		return Instruction{
			Command: "SYNTAX_ERROR",
			Message: "invalid parameters for command 'login'",
		}
	}

	uid := commands[0]
	password := commands[1]

	user, reason := database.CheckUserLogin(uid, password)
	ret := Instruction{
		Command: reason,
		Data:    user,
	}

	return ret
}
