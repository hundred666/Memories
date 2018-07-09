package model

import "time"

type MCR struct {
	Id   int       `json:"id"`
	Mid  int       `json:"mid"`
	Cid  int       `json:"cid"`
	Time time.Time `json:"time"`
}
