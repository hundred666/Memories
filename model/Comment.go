package model

import (
	"time"
)

type Comment struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	CommentTime time.Time `json:"comment_time"`
	CommentType int       `json:"comment_type"` //0->lanxiong 1->move comment
	User        string    `json:"user"`
	IP          string    `json:"ip"`
	UA          string    `json:"ua"`
}
