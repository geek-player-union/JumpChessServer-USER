package database

type UserItem struct {
	Id     int64 `json:"id" xorm:"pk autoincr"`
	UserID int64 `json:"user_id" xorm:"not null references user(id)"`
	ItemID int64 `json:"item_id" xorm:"not null"`
}
