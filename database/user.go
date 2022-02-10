package database

import "fmt"

type User struct {
	ID       int64  `json:"id" xorm:"pk autoincr"`
	Username string `json:"username" xorm:"varchar(50) notnull unique"`
	Password string `json:"password" xorm:"varchar(50) notnull"`
	Avatar   string `json:"avatar" xorm:"varchar(50) notnull"`
	Exp      int64  `json:"exp" xorm:"bigint notnull"`
	Coin     int64  `json:"coin" xorm:"bigint notnull"`
	Level    int64  `json:"level" xorm:"bigint notnull"`
	Banned   bool   `json:"banned" xorm:"bool notnull"`
}

func (u *User) ToProtocolString() string {
	ret := fmt.Sprintf("USER %s %s %d %d %d ", u.Username, u.Avatar, u.Level, u.Exp, u.Coin)

	items := u.GetItems()

	var hasItem [ItemNumber]bool
	for _, item := range items {
		hasItem[item.Id] = true
	}

	for _, has := range hasItem {
		if has {
			ret += "1"
		} else {
			ret += "0"
		}
	}

	return ret
}

func (u *User) GetItems() []Item {
	// TODO: finish
	return nil
}

func CheckUserLogin(uid int, password string) (*User, string) {
	// TODO: finish CODE_ERROR & ACCOUNT_ERROR & OK
	return nil, "ACCOUNT_ERROR"
}
