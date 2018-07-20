package service

import (
	"model"
	"dao"
	"fmt"
	"util"
)

func AddComment(u model.User, c model.Comment) (error) {
	_, err := dao.AddComment(c)
	commentContent := util.TrimContent(c.Content)
	commentTime := util.ParseTime(c.CommentTime)
	content := fmt.Sprintf("%v在%v在懒熊添加了一个评论,内容是：%v", u.Name, commentTime, commentContent)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: c.CommentTime}
	dao.AddAnnounce(a)
	return err
}

func DelComment(u model.User, c model.Comment) (error) {
	_, err := dao.DelComment(c)
	/*commentTime := util.ParseTime(c.CommentTime)
	content := fmt.Sprintf("%v在%v在懒熊删除了一个评论", u.Name, commentTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: c.CommentTime}
	dao.AddAnnounce(a)*/
	return err
}

func UpdateComment(u model.User, c model.Comment) (error) {
	_, err := dao.UpdateComment(c)
	/*commentTime := util.ParseTime(c.CommentTime)
	content := fmt.Sprintf("%v在%v在懒熊更新了一个评论", u.Name, commentTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: c.CommentTime}
	dao.AddAnnounce(a)*/
	return err
}

func GetComment(id int, args ...string) (model.Comment) {
	return dao.GetComment(id, args...)
}

func GetComments(start int, end int, args ...string) ([]model.Comment) {
	return dao.GetComments(start, end, args...)

}

func GetCommentCount() (int) {
	return dao.GetCommentCount()
}
