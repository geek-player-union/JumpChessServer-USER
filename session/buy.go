package session

import "github.com/geek-player-union/JumpChessServer-USER/database"

func handleBuy(d interface{}, uid int64) string {
	data, ok := d.(map[string]interface{})

	if !ok {
		return "SYNTAX_ERROR"
	}
	itemID := data["itemid"].(float64)
	if itemID >= database.ItemNumber || itemID < 0 {
		return "INVALID_ITEM"
	}

	user := database.GetUserById(uid)
	if user == nil {
		return "NEED_LOGIN"
	}

	return user.BuyItem(int(itemID))
}
