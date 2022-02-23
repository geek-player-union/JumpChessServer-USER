package session

import "github.com/geek-player-union/JumpChessServer-USER/database"

func handleBuy(d interface{}, uid int64) string {
	data, ok := d.(struct {
		ItemId int `json:"itemid"`
	})

	if !ok {
		return "SYNTAX_ERROR"
	}

	if data.ItemId >= database.ItemNumber || data.ItemId < 0 {
		return "INVALID_ITEM"
	}

	user := database.GetUserById(uid)
	if user == nil {
		return "NEED_LOGIN"
	}

	return user.BuyItem(data.ItemId)
}
