package util

import (
	"time"
)

func TrimContent(str string) (string) {
	if len([]rune(str)) < 10 {
		return str
	}
	return string([]rune(str)[:10]) + "..."
}

func ParseTime(t time.Time) (string) {
	return t.Format("01-02 15:04")
}
