package model

import "time"

type User struct {
	Id        uint
	Name      string
	CreatedAt time.Time
}
