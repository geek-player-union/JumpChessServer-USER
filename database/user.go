package database

import (
	"encoding/json"
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

func (u *User) ToJsonString() string {
	ret, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	return string(ret)
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

	return "BUY_DONE"
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

func CheckUserLogin(uid string, password string) (*User, string) {
	id, err := strconv.Atoi(uid)
	if err != nil {
		return nil, "INVALID_ACCOUNT"
	}

	user := GetUserById(int64(id))
	if user == nil {
		return nil, "NO_SUCH_USER"
	}

	if user.Password != password {
		return nil, "CODE_ERROR"
	}

	if user.Online {
		return nil, "ONLINE_ERROR"
	}

	user.Online = true
	if !user.Update() {
		return nil, "SERVER_ERROR"
	}

	return user, "USER"
}
