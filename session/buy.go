package session

import "github.com/geek-player-union/JumpChessServer-USER/database"

func handleBuy(d interface{}, uid int64) Instruction {
	data, ok := d.(struct {
		ItemId int `json:"itemid"`
	})

	if !ok {
		return Instruction{
			Command: "SYNTAX_ERROR",
		}
	}

	if data.ItemId >= database.ItemNumber || data.ItemId < 0 {
		return Instruction{
			Command: "INVALID_ITEM",
		}
	}

	user := database.GetUserById(uid)
	if user == nil {
		return Instruction{
			Command: "NEED_LOGIN",
		}
	}

	return Instruction{
		Command: user.BuyItem(data.ItemId),
	}
}
