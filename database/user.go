package database

import (
	"encoding/json"
)

type User struct {
	ID       int64  `json:"id" xorm:"pk autoincr"`
	Username string `json:"username" xorm:"varchar(50) notnull unique"`
	Password string `json:"password" xorm:"varchar(50) notnull"`
	Avatar   string `json:"avatar" xorm:"varchar(50) notnull"`
	Exp      int64  `json:"exp" xorm:"bigint notnull"`
	Coin     int64  `json:"coin" xorm:"bigint notnull"`
	Level    int64  `json:"level" xorm:"bigint notnull"`
	Banned   bool   `json:"banned" xorm:"bool notnull"`
	Online   bool   `json:"online" xorm:"bool notnull"`
	Rank     int64  `json:"rank" xorm:"bigint notnull"`
	Assert   string `json:"assert" xorm:"varchar(50) notnull"`
}

func (u *User) ToJsonString() string {
	ret, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	return string(ret)
}

func BuyItems(id int64, ItemId int) string {
	user := &User{ID: id}
	has, _ := engine.Get(user)

	if !has {
		return "DATABASE_ERROR"
	}

	if user.Assert[ItemId] != '0' {
		return "RECURRING_ERROR"
	}

	if user.Coin < Items[ItemId].Price {
		return "COIN_NOT_ENOUGH"
	}

	assertBytes := []byte(user.Assert)
	assertBytes[ItemId] = '1'
	user.Assert = string(assertBytes)
	user.Coin -= Items[ItemId].Price
	_, err := engine.Update(user)
	if err != nil {
		return "DATABASE_ERROR"
	}
	
	return "BUY_DONE"
}

func CheckUserLogin(uid string, password string) (*User, string) {
	user := &User{}
	has, _ := engine.Where("ID = ?", uid).Get(user)

	if has {
		if user.Password == password && user.Online == false {
			user.Online = true
			_, err := engine.Update(user)
			if err != nil {
				return nil, "SERVER_ERROR"
			}
			return user, "USER" //OK
		} else if user.Password != password {
			return nil, "CODE_ERROR"
		} else {
			return nil, "ONLINE_ERROR"
		}
	} else {
		return nil, "ACCOUNT_ERROR"
	}
}
