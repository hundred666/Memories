package user

import (
	"net/http"
	"model"
	"dao"
	"handler"
	"strconv"
	"time"
)

type UserHandler struct {
}

const USER = "USER"

func (u *UserHandler) LoginCheck(req *http.Request) bool {
	if _, err := req.Cookie("userId"); err == nil {
		return true
	}
	return false
}

func (u *UserHandler) Login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var user model.User
	var cookie http.Cookie
	userName := req.Form.Get("username")
	password := req.Form.Get("password")
	if userName != "" && password != "" {
		user.Name = userName
		user.Password = password
		if dao.UserExist(userName) {
			if dao.AdminPermission(user) {
				user = dao.GetUserByName(user.Name)
				cookie = http.Cookie{Name: "userId", Value: strconv.Itoa(user.Id), Path: "/", MaxAge: 86400}
				http.SetCookie(w, &cookie)
				w.Write(model.MarshalResponse(0, "login succeed"))
				return
			}
		} else {
			user.RegisterTime = time.Now()
			user.LoginTime = time.Now()
			user.Permission = 0
			dao.AddUser(user)
			w.Write(model.MarshalResponse(0, "register succeed"))
			return
		}
	}
	w.Write(model.MarshalResponse(1, "failed"))
}

func (u *UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	if u.LoginCheck(req) {
		cookie := http.Cookie{Name: "userId", Path: "/", MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

func (u *UserHandler) Register(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var user model.User
	userName := req.Form.Get("name")
	password := req.Form.Get("password")
	if userName != "" && password != "" {
		user.Name = userName
		user.Password = password
		if dao.UserExist(userName) {
			w.Write(model.MarshalResponse(1, "user exist"))
			return
		} else {
			user.RegisterTime = time.Now()
			user.LoginTime = time.Now()
			user.Permission = 0
			dao.AddUser(user)
			w.Write(model.MarshalResponse(0, "register succeed"))
			return
		}
	} else {
		w.Write(model.MarshalResponse(1, "please input correct msg"))
		return
	}
}

func (u *UserHandler) GetUsers(w http.ResponseWriter, req *http.Request) {
	var users []model.User
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		users = dao.GetUsers(handler.DEFAULT_START, handler.DEFAULT_END)
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		users = dao.GetUsers(startSeq, endSeq)
	}
	w.Write(model.MarshalResponse(0, users))
}

func (u *UserHandler) GetUser(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	userId := vars.Get("userId")
	if userId == "" {
		w.Write(model.MarshalResponse(1, "user id wrong"))
		return
	}
	id, _ := strconv.Atoi(userId)
	user := dao.GetUserById(id)
	w.Write(model.MarshalResponse(0, user))
}
