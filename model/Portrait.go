package model

import "time"

type Portrait struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Path string    `json:"path"`
	Time time.Time `json:"time"`
	User string    `json:"user"`
}
