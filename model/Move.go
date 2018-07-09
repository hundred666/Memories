package model

import "time"

type Move struct {
	Id      int       `json:"id"`
	Content string    `json:"content"`
	ImgPath string    `json:"img_path"`
	User    string    `json:"user"`
	Time    time.Time `json:"time"`
}
