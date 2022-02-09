package session

import "fmt"

// 返回 true 代表可以结束 Socket
func handle(commands []string) string {
	// TODO: Delete
	fmt.Println("Command received: " + fmt.Sprint(commands))

	if len(commands) == 0 {
		return "empty"
	}

	switch commands[0] {
	case "login":
		return handleLogin(commands[1:])
	default:
		return "unknown"
	}
}
