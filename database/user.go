package database

import (
    "fmt"
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

func BetItems(uid string, password string, ItemId int) (bool, string) {
    user := &User{}
    has, _ := engine.Where("ID = ?", uid).Get(user)

    if has {
        if user.Password == password && user.Assert[ItemId] == '0' && user.Coin >= Items[ItemId].Price {
            AssertBytes := []byte(user.Assert)
            AssertBytes[ItemId] = '1'
            user.Assert =string(AssertBytes)
            user.Coin-=Items[ItemId].Price
            _, err := engine.Update(user)
            if err != nil {
                return false, "SERVER_ERROR"
            }
            return true, "BUY_DONE"
        } else if user.Password != password {
            return false, "CONFIRM_ERROR"
        }else if user.Assert[ItemId] == '1' {
            return false, "RECURRING_ERROR"
        } else {
            return false, "COIN_ERROR"
        }
    } else {
        return false,"DATABASE_ERROR"
    }
}

func CheckUserLogin(uid string, password string) (*User, string) {
    user := &User{}
    has, _ := engine.Where("ID = ?", uid).Get(user)

    if has {
        if user.Password == password && user.Online== false {
            user.Online=true
            _, err := engine.Update(user)
            if err != nil {
                return nil, "SERVER_ERROR"
            }
            return user, "USER"     //OK
        } else if user.Password != password {
            return nil, "CODE_ERROR"
        } else {
            return nil, "ONLINE_ERROR"
        }
    } else {
        return nil, "ACCOUNT_ERROR"
    }
}
