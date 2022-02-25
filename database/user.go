package database

import (
	"strconv"
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

func (u *User) BuyItem(ItemId int) string {
	if u.Assert[ItemId] != '0' {
		return "RECURRING_ERROR"
	}

	if u.Coin < Items[ItemId].Price {
		return "COIN_NOT_ENOUGH"
	}

	assertBytes := []byte(u.Assert)
	assertBytes[ItemId] = '1'
	u.Assert = string(assertBytes)
	u.Coin -= Items[ItemId].Price
	if !u.Update() {
		return "SERVER_ERROR"
	}

	return ""
}

func (u *User) Update() bool {
	_, err := engine.Update(u)
	if err != nil {
		return false
	}
	return true
}

func GetUserById(uid int64) *User {
	user := &User{ID: uid}
	has, err := engine.Get(user)
	if !has || err != nil {
		return nil
	}
	return user
}

func CheckUserLogin(uid string, password string) (string, *User) {
	id, err := strconv.Atoi(uid)
	if err != nil {
		return "ACCOUNT_ERROR", nil
	}

	user := GetUserById(int64(id))
	if user == nil {
		return "ACCOUNT_ERROR", nil
	}

	if user.Password != password {
		return "CODE_ERROR", nil
	}

	//if user.Online {
	//	return "ONLINE_ERROR", -1
	//}

	user.Online = true
	if !user.Update() {
		return "SERVER_ERROR", nil
	}

	return "", user
}
