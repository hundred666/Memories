package admin

import (
	"net/http"
	"handler"
	"log"
	"html/template"
	"service"
	"model"
	"strconv"
	"time"
)

type AdminHandler struct {
}

const ADMIN = "ADMIN"

func (a *AdminHandler) LoginCheck(req *http.Request) (bool, model.User) {
	var user model.User
	id, err := req.Cookie("userId")
	if err != nil {
		return false, user
	}
	userId, _ := strconv.Atoi(id.Value)
	user = service.GetUserById(userId)
	return true, user

}

func (a *AdminHandler) Index(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["Login"] = 0
	if login, user := a.LoginCheck(req); !login {
		data["Login"] = 1
	} else {

		data["UserName"] = user.Name
		userCount := service.GetUserCount()
		moveCount := service.GetMoveCount()
		commentCount := service.GetCommentCount()
		portraitCount := service.GetPortraitCount()

		data["Users"] = userCount
		data["Moves"] = moveCount
		data["Comments"] = commentCount
		data["Portraits"] = portraitCount
	}
	t, err := template.ParseFiles(handler.GetAdmin("index.html"))
	if err != nil {
		log.Fatal(err.Error())
	}

	t.Execute(w, data)
}

func (a *AdminHandler) UpdateComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, user := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	commentId, err := strconv.Atoi(req.Form.Get("commentId"))
	commentUser := req.Form.Get("commentUser")
	content := req.Form.Get("commentContent")
	comment := model.Comment{
		Id:      commentId,
		User:    commentUser,
		Content: content}
	err = service.UpdateComment(user, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment update success"))
}

func (a *AdminHandler) DelComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, user := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	commentId, err := strconv.Atoi(req.Form.Get("commentId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment type error"))
		return
	}
	comment := model.Comment{
		Id: commentId}
	err = service.DelComment(user, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment update success"))
}

func (a *AdminHandler) UpdateMove(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, loginUser := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	moveId, err := strconv.Atoi(req.Form.Get("moveId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "moveId error"))
		return
	}
	user := req.Form.Get("moveUser")
	content := req.Form.Get("moveContent")
	move := model.Move{Id: moveId, Content: content, User: user}
	err = service.UpdateMove(loginUser, move)
	if err != nil {
		w.Write(model.MarshalResponse(1, "move update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "move update success"))
}

func (a *AdminHandler) DelMove(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, user := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	moveId, err := strconv.Atoi(req.Form.Get("moveId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "moveId error"))
		return
	}

	move := model.Move{Id: moveId}
	err = service.DelMove(user, move)
	if err != nil {
		log.Println()
		w.Write(model.MarshalResponse(1, "move delete failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "move delete success"))
}

func (a *AdminHandler) UpdateMoveComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, user := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}

	moveId, err := strconv.Atoi(req.Form.Get("moveId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "moveId error"))
		return
	}
	move := model.Move{Id: moveId}

	commentId, err := strconv.Atoi(req.Form.Get("commentId"))
	commentUser := req.Form.Get("commentUser")
	content := req.Form.Get("content")
	comment := model.Comment{
		Id:      commentId,
		User:    commentUser,
		Content: content}

	err = service.UpdateMoveComment(user, move, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment update success"))
}

func (a *AdminHandler) DelMoveComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, user := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	moveId, err := strconv.Atoi(req.Form.Get("moveId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "moveId error"))
		return
	}
	move := model.Move{Id: moveId}
	commentId, err := strconv.Atoi(req.Form.Get("commentId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment type error"))
		return
	}
	comment := model.Comment{Id: commentId}
	err = service.DelMoveComment(user, move, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment delete failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment delete success"))

}

func (a *AdminHandler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, _ := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	userId, err := strconv.Atoi(req.Form.Get("userId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "userId error"))
		return
	}
	userName := req.Form.Get("username")
	password := req.Form.Get("password")
	permission, err := strconv.Atoi(req.Form.Get("userPermission"))
	if err != nil {
		permission = 0
	}

	user := model.User{Id: userId, Name: userName, Password: password, Permission: permission}
	err = service.UpdateUser(user)
	if err != nil {
		w.Write(model.MarshalResponse(1, "user update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "user update success"))
}

func (a *AdminHandler) DelUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, _ := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	userId, err := strconv.Atoi(req.Form.Get("userId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "userId error"))
		return
	}

	user := model.User{Id: userId}
	err = service.DelUser(user)
	if err != nil {
		w.Write(model.MarshalResponse(1, "user delete failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "user delete success"))
}

func (a *AdminHandler) UpdateAnnounce(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, _ := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	announceId, err := strconv.Atoi(req.Form.Get("announceId"))
	announceUser := req.Form.Get("announceUser")
	content := req.Form.Get("announceContent")
	display := false
	d := req.Form.Get("announceDisplay")
	if d == "on" {
		display = true
	}

	announce := model.Announce{
		Id:      announceId,
		User:    announceUser,
		Display: display,
		Content: content}

	err = service.UpdateAnnounce(announce)
	if err != nil {
		w.Write(model.MarshalResponse(1, "announce update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "announce update success"))
}

func (a *AdminHandler) DelAnnounce(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, _ := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	announceId, err := strconv.Atoi(req.Form.Get("announceId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "announce type error"))
		return
	}
	announce := model.Announce{
		Id: announceId}
	_, err = service.DelAnnounce(announce)
	if err != nil {
		w.Write(model.MarshalResponse(1, "announce del failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "announce del success"))
}

func (a *AdminHandler) AddAnnounce(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	login, user := a.LoginCheck(req)
	if !login {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}

	content := req.Form.Get("newAnnounceContent")
	announceTime := time.Now()

	announce := model.Announce{
		Content: content,
		Display: true,
		Prior:   0,
		User:    user.Name,
		Time:    announceTime,
	}
	err := service.AddAnnounce(user, announce)
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	w.Write(model.MarshalResponse(0, "success"))
}
