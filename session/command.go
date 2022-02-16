package session

import (
	"encoding/json"
	"strings"
)

type Instruction struct {
	Command string
	Message string
	Data    interface{}
}

func handle(command string, s *session) Instruction {
	recv := &Instruction{}
	err := json.Unmarshal([]byte(command), recv)
	if err != nil {
		return Instruction{
			Command: "SYNTAX_ERROR",
			Message: "invalid json syntax",
		}
	}

	recv.Command = strings.ToUpper(recv.Command)
	switch recv.Command {
	case "LOGIN":
		ret, uid := handleLogin(recv.Data)
		s.uid = uid
		return ret
	case "BUY":
		return handleBuy(recv.Data, s.uid)
	default:
		return Instruction{
			Command: "UNKNOWN_COMMAND",
			Message: "unknown command error",
		}
	}
}
