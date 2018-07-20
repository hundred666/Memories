package dao

import (
	"testing"
	"model"
	"time"
	"fmt"
	"strings"
)

func TestInserComment(t *testing.T) {
	c := model.Comment{User: "abc", Content: "dasffdwef", CommentTime: time.Now(), CommentType: 1, IP: "ABC", UA: "GOO"}
	_, err := AddComment(c)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestGetComments(t *testing.T) {
	start, end := 0, -1
	comments := GetComments(start, end)
	fmt.Println(comments)
}

func TestGetPortraits(t *testing.T) {
	start, end := 0, -1
	portraits := GetPortraits(start, end)
	fmt.Println(portraits)
}

func TestSplit(t *testing.T) {
	a := "/static/portraits/wx_camera_1529469909420.jpg"
	fmt.Println(strings.SplitAfterN(a, "/", 3))
}

func TestDelMove(t *testing.T) {
	move := model.Move{Id: 1}
	fmt.Println(DelMove(move))
}

func TestDelPortrait(t *testing.T) {
	p := model.Portrait{Id: 9, Path: "/static/portraits/k.png"}
	fmt.Println(DelPortrait(p))
}

func TestGetUserCount(t *testing.T) {
	fmt.Println(GetUserCount())
}

