package session

import "net"

type session struct {
	uid        int
	connection net.Conn
	terminated bool
}

func StartSession(connection net.Conn) {
	s := &session{
		uid:        -1,
		connection: connection,
		terminated: false,
	}

	go s.loop()
}
