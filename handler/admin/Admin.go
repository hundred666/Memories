package admin

import (
	"net/http"
	"handler"
	"log"
	"html/template"
	"dao"
	"model"
	"strconv"
)

type AdminHandler struct {
}

const ADMIN = "ADMIN"

func (a *AdminHandler) LoginCheck(req *http.Request) bool {
	if _, err := req.Cookie("uname"); err == nil {
		return true
	}
	return false
}

func (a *AdminHandler) Index(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	if !a.LoginCheck(req) {
		data["Login"] = 1
	} else {
		data["Login"] = 0
		userCookie, _ := req.Cookie("uname")
		userName := userCookie.Value
		data["UserName"] = userName
		userCount := dao.GetUserCount()
		moveCount := dao.GetMoveCount()
		commentCount := dao.GetCommentCount()
		portraitCount := dao.GetPortraitCount()

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
	if !a.LoginCheck(req) {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	commentId, err := strconv.Atoi(req.Form.Get("commentId"))
	commentUser := req.Form.Get("commentUser")
	content := req.Form.Get("content")
	comment := model.Comment{
		Id:      commentId,
		UName:   commentUser,
		Content: content}
	_, err = dao.UpdateComment(comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment update success"))
}

func (a *AdminHandler) DelComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
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
	_, err = dao.DelComment(comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment update success"))
}

func (a *AdminHandler) UpdateMove(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	moveId, err := strconv.Atoi(req.Form.Get("moveId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "moveId error"))
		return
	}
	user:=req.Form.Get("moveUser")
	content := req.Form.Get("content")
	move := model.Move{Id: moveId, Content: content,User:user}
	_, err = dao.UpdateMove(move)
	if err != nil {
		w.Write(model.MarshalResponse(1, "move update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "move update success"))
}

func (a *AdminHandler) DelMove(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	moveId, err := strconv.Atoi(req.Form.Get("moveId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "moveId error"))
		return
	}

	move := model.Move{Id: moveId}
	_, err = dao.DelMove(move)
	if err != nil {
		log.Println()
		w.Write(model.MarshalResponse(1, "move delete failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "move delete success"))
}

func (a *AdminHandler) UpdateMoveComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
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
		UName:   commentUser,
		Content: content}
	_, err = dao.UpdateMoveComment(move, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment update success"))
}

func (a *AdminHandler) DelMoveComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
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
	_, err = dao.DelMoveComment(move, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, "comment delete failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "comment delete success"))

}

func (a *AdminHandler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
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
	permission, err := strconv.Atoi(req.Form.Get("permission"))
	if err != nil {
		permission = 0
	}

	user := model.User{Id: userId, Name: userName, Password: password, Permission: permission}
	err = dao.UpdateUser(user)
	if err != nil {
		w.Write(model.MarshalResponse(1, "user update failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "user update success"))
}

func (a *AdminHandler) DelUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if !a.LoginCheck(req) {
		w.Write(model.MarshalResponse(1, "not login"))
		return
	}
	userId, err := strconv.Atoi(req.Form.Get("userId"))
	if err != nil {
		w.Write(model.MarshalResponse(1, "userId error"))
		return
	}

	user := model.User{Id: userId}
	_, err = dao.DelUser(user)
	if err != nil {
		w.Write(model.MarshalResponse(1, "user delete failed"))
		return
	}
	w.Write(model.MarshalResponse(0, "user delete success"))
}
