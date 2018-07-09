package model

import (
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	Permission   int       `json:"permission"` //0 upload picture
	RegisterTime time.Time `json:"register_time"`
	LoginTime    time.Time `json:"login_time"`
}
