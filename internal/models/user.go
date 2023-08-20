package models

import "time"

type User struct {
	Id       int64
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
}
