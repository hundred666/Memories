package service

import (
	"model"
	"dao"
	"util"
	"fmt"
)

func AddAnnounce(u model.User, a model.Announce) (error) {
	announceContent := util.TrimContent(a.Content)
	announceTime := util.ParseTime(a.Time)
	content := fmt.Sprintf("%v在%v添加了一个通知,内容是：%v", u.Name, announceTime, announceContent)
	a = model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: a.Time}
	return dao.AddAnnounce(a)
}

func GetLatestAnnounce() (model.Announce) {
	return dao.GetLatestAnnounce()
}

func DelAnnounce(a model.Announce) (int, error) {
	return dao.DelAnnounce(a)
}

func GetAnnounce(id int) (model.Announce) {
	return dao.GetAnnounce(id)
}

func GetAnnounces(start int, end int, args ...string) ([]model.Announce) {
	return dao.GetAnnounces(start, end,args...)
}

func UpdateAnnounce(a model.Announce) (error) {
	_, err := dao.UpdateAnnounce(a)
	return err
}
