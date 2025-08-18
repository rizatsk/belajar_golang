package entity

import "time"

type User struct {
	Id         string
	Name       string
	Email      string
	Created_dt time.Time
}
