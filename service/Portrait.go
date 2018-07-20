package service

import (
	"model"
	"dao"
	"util"
	"fmt"
)

func AddPortrait(u model.User, p model.Portrait) (error) {
	err := dao.AddPortrait(p)
	portraitTime := util.ParseTime(p.Time)
	content := fmt.Sprintf("%v在%v上传了一张帅照", u.Name, portraitTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: p.Time}
	dao.AddAnnounce(a)
	return err
}

func DelPortrait(u model.User, p model.Portrait) (error) {
	_, err := dao.DelPortrait(p)
	/*portraitTime := util.ParseTime(p.Time)
	content := fmt.Sprintf("%v在%v删除了一张帅照", u.Name, portraitTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: p.Time}
	dao.AddAnnounce(a)*/
	return err
}

func UpdatePortrait(u model.User, p model.Portrait) (error) {
	_, err := dao.UpdatePortrait(p)
	/*portraitTime := util.ParseTime(p.Time)
	content := fmt.Sprintf("%v在%v更新了一张帅照", u.Name, portraitTime)
	a := model.Announce{Content: content, Display: true, Prior: 2, User: u.Name, Time: p.Time}
	dao.AddAnnounce(a)*/
	return err
}

func GetPortraits(start int, end int, args ...string) ([]model.Portrait) {
	return dao.GetPortraits(start, end,args...)
}

func GetPortraitCount() (int) {
	return dao.GetPortraitCount()
}
