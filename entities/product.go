package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	User        User
	Quantity    int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
