package database

type Item struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

const ItemNumber = 5

var Items = []Item{{
	Id:          0,
	Name:        "Elf Boy",
	Description: "This is an elf boy.",
	Price:       1,
}, {
	Id:          1,
	Name:        "Elf Girl",
	Description: "This is a elf girl.",
	Price:       1,
}, {
	Id:          2,
	Name:        "Knight",
	Description: "This is a knight.",
	Price:       1,
}, {
	Id:          3,
	Name:        "Witch Boy",
	Description: "This is a witch boy.",
	Price:       1,
}, {
	Id:          4,
	Name:        "Witch Girl",
	Description: "This is a witch girl.",
	Price:       1,
}}
