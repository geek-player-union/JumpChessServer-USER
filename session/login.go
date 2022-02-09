package session

import "fmt"

func handleLogin(commands []string) string {
	if len(commands) != 2 {
		return "error"
	}
	uid := commands[0]
	password := commands[1]

	// TODO: login logic
	fmt.Println(uid, password)

	return "OK"
}
