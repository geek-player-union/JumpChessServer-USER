package session

import (
	"github.com/geek-player-union/JumpChessServer-USER/database"
	"strconv"
)

func handleLogin(commands []string) Instruction {
	if len(commands) != 2 {
		return Instruction{
			Command: "SYNTAX_ERROR",
			Message: "invalid parameters for command 'login'",
		}
	}

	uidStr := commands[0]
	password := commands[1]
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return Instruction{
			Command: "SYNTAX_ERROR",
			Message: "invalid uid format",
		}
	}

	user, reason := database.CheckUserLogin(uid, password)
	ret := Instruction{
		Command: reason,
		Data:    user,
	}

	return ret
}
