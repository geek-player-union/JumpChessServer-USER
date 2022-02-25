package session

import (
	"encoding/json"
	"github.com/geek-player-union/JumpChessServer-USER/database"
	"strings"
)

type Instruction struct {
	Command string
	Message string
	Data    interface{}
}

func handle(command string, s *session) Instruction {
	recv := &Instruction{}
	send := Instruction{}
	err := json.Unmarshal([]byte(command), recv)
	if err != nil {
		return Instruction{
			Command: "SYNTAX_ERROR",
			Message: "invalid json syntax",
		}
	}

	recv.Command = strings.ToUpper(recv.Command)
	send.Command = recv.Command + "_"
	var message string
	var data interface{}

	switch recv.Command {
	case "LOGIN":
		message, data = handleLogin(recv.Data)
		if data != nil {
			s.uid = data.(*database.User).ID
		}
	case "BUY":
		message = handleBuy(recv.Data, s.uid)
	default:
		return Instruction{
			Command: "UNKNOWN_ERROR",
		}
	}

	if len(message) == 0 {
		send.Command += "DONE"
	} else {
		send.Command += "ERROR"
	}

	send.Message = message
	send.Data = data
	return send
}
