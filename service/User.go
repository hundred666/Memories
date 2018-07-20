package service

import (
	"model"
	"dao"
)

func AddUser(user model.User) (error) {
	return dao.AddUser(user)
}

func UpdateUserLogin(user model.User) (error) {
	return dao.UpdateUserLogin(user)
}

func UpdateUser(user model.User) (error) {
	return dao.UpdateUser(user)
}

func DelUser(user model.User) (error) {
	_, err := dao.DelUser(user)
	return err
}

func GetUserByName(Name string) (model.User) {
	return dao.GetUserByName(Name)
}

func GetUserById(id int) (model.User) {
	return dao.GetUserById(id)
}

func UserLogin(user model.User) (bool) {
	return dao.UserLogin(user)
}

func AdminPermission(user model.User) (bool) {
	return dao.AdminPermission(user)
}

func UploadPermission(user model.User) (bool) {
	return dao.UploadPermission(user)
}

func UserExist(Name string) (bool) {
	return dao.UserExist(Name)
}

func GetUsers(start int, end int, args ...string) []model.User {
	return dao.GetUsers(start, end,args...)
}

func GetUserCount() (int) {
	return dao.GetUserCount()
}
