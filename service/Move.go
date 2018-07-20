package service

import (
	"model"
	"dao"
	"util"
	"fmt"
)

func AddMove(u model.User, m model.Move) (error) {
	err := dao.AddMove(m)
	moveContent := util.TrimContent(m.Content)
	moveTime := util.ParseTime(m.Time)
	content := fmt.Sprintf("%v %v 添加了一个记忆瞬间,内容是：%v", u.Name, moveTime, moveContent)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: m.Time}
	dao.AddAnnounce(a)
	return err
}

func AddMoveComment(u model.User, m model.Move, c model.Comment) (error) {
	err := dao.AddMoveComment(m, c)
	commentContent := util.TrimContent(c.Content)
	commentTime := util.ParseTime(c.CommentTime)
	content := fmt.Sprintf("%v在%v在记忆添加了一个评论,内容是：%v", u.Name, commentTime, commentContent)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: c.CommentTime}
	dao.AddAnnounce(a)
	return err
}

func GetMoves(start int, end int, args ...string) ([]model.Move) {
	return dao.GetMoves(start, end,args...)
}

func GetMove(id int) model.Move {
	return dao.GetMove(id)
}

func GetMoveComments(m model.Move) []model.Comment {
	return dao.GetMoveComments(m)
}
func UpdateMove(u model.User, m model.Move) (error) {
	_, err := dao.UpdateMove(m)
	/*moveContent := util.TrimContent(m.Content)
	moveTime := util.ParseTime(m.Time)
	content := fmt.Sprintf("%v在%v更新了一个记忆瞬间,内容是：%v", u.Name, moveTime, moveContent)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: m.Time}
	dao.AddAnnounce(a)*/
	return err
}

func UpdateMoveComment(u model.User, m model.Move, c model.Comment) (error) {
	_, err := dao.UpdateMoveComment(m, c)
	/*commentContent := util.TrimContent(c.Content)
	moveContent := util.TrimContent(m.Content)
	commentTime := util.ParseTime(c.CommentTime)
	content := fmt.Sprintf("%v在%v在 %v 更新了一个评论,内容是：%v", u.Name, commentTime, moveContent, commentContent)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: c.CommentTime}
	dao.AddAnnounce(a)*/
	return err
}

func DelMove(u model.User, m model.Move) (error) {
	_, err := dao.DelMove(m)
	/*moveTime := util.ParseTime(m.Time)
	content := fmt.Sprintf("%v在%v删除了一个记忆瞬间", u.Name, moveTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: m.Time}
	dao.AddAnnounce(a)*/
	return err
}

func DelMoveComment(u model.User, m model.Move, c model.Comment) (error) {
	_, err := dao.DelMoveComment(m, c)
	/*moveTime := util.ParseTime(m.Time)
	content := fmt.Sprintf("%v在%v删除了一个记忆评论", u.Name, moveTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: m.Time}
	dao.AddAnnounce(a)*/
	return err
}

func GetMoveCount() (int) {
	return dao.GetMoveCount()
}
