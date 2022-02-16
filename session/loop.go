package session

import (
	"encoding/json"
	"strings"
)

func (s *session) loop() {
	defer s.connection.Close()
	buf := make([]byte, 512)
	var err error
	var n int
	var message string
	var command string
	for !s.terminated {
		n, err = s.connection.Read(buf)
		if err != nil || n < 1 {
			break
		}
		message += string(buf)

		for strings.Index(message, "\n") >= 0 {
			command = strings.Split(message, "\n")[0]
			message = message[len(command):]
			if message[0] == '\n' {
				message = message[1:]
			}
			retBytes, _ := json.Marshal(handle(command, s))
			ret := string(retBytes) + "\n"
			if len(ret) == 0 {
				s.terminated = true
			} else {
				_, err = s.connection.Write([]byte(ret))
				if err != nil {
					s.terminated = true
				}
			}
		}
	}
}
