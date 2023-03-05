package models

import (
	"time"
)

type Transaction struct {
	// "id":1,"full_name":"Liana Jaquest","quantity":9,"price":"$94.79"
	Id        int
	Fullname  string
	Quantity  int
	Price     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
