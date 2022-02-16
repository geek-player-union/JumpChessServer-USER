package session

import (
	"encoding/json"
	"strings"
)

func (s *session) loop() {
	defer s.connection.Close()
	buf := make([]byte, 512)
	var message string
	var command string
	for !s.terminated {
		if n, err := s.connection.Read(buf); err != nil || n < 1 {
			break
		}
		message += string(buf)
		for strings.Index(message, "\n") >= 0 {
			command = strings.Split(message, "\n")[0]
			message = message[len(command)+1:]
			retBytes, _ := json.Marshal(handle(command, s))
			ret := string(retBytes) + "\n"
			if _, err := s.connection.Write([]byte(ret)); err != nil {
				s.terminated = true
			}
		}
	}
}
