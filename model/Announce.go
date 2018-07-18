package model

import "time"

type Announce struct {
	Id      int       `json:"id"`
	Content string    `json:"content"`
	Display bool      `json:"display"`
	Prior   int       `json:"prior"`
	User    string    `json:"user"`
	Time    time.Time `json:"time"`
}
