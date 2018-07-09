package model

import "time"

type Portrait struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Path   string    `json:"path"`
	UpTime time.Time `json:"up_time"`
	UpUser string    `json:"up_user"`
}
