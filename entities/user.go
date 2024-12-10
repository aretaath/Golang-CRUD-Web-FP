package entities

import "time"

type User struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
