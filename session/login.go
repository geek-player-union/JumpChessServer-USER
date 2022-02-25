package session

import (
	"fmt"
	"github.com/geek-player-union/JumpChessServer-USER/database"
)

func handleLogin(d interface{}) (string, *database.User) {
	data, ok := d.(map[string]interface{})

	if !ok {
		return "SYNTAX_ERROR", nil
	}
	fmt.Println(data)
	return database.CheckUserLogin(data["account"].(string), data["code"].(string))
}
