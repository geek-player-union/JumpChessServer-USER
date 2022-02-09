package server

import (
	"fmt"
	"github.com/geek-player-union/JumpChessServer-USER/session"
	"net"
)

func Start(port int) {
	serverSocket, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server started at port %d\n", port)

	for {
		connection, err := serverSocket.Accept()
		if err != nil {
			panic(err)
		}

		session.StartSession(connection)
	}
}
