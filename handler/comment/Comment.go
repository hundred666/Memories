package comment

import (
	"net/http"
	"model"
	"time"
	"service"
	"strconv"
	"handler"
)

type CommentHandler struct {
}

const COMMENT = "COMMENT"

func (c *CommentHandler) GetComments(w http.ResponseWriter, req *http.Request) {
	var comments []model.Comment
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		comments = service.GetComments(handler.DEFAULT_START, handler.DEFAULT_END, "WHERE CommentType=0")
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		comments = service.GetComments(startSeq, endSeq, "WHERE CommentType=0")
	}
	w.Write(model.MarshalResponse(0, comments))

}

func (c *CommentHandler) GetComment(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	commentId := vars.Get("commentId")
	if commentId == "" {
		w.Write(model.MarshalResponse(1, "comment id wrong"))
		return
	}
	id, _ := strconv.Atoi(commentId)
	comment := service.GetComment(id)
	w.Write(model.MarshalResponse(0, comment))

}

func (c *CommentHandler) GetAllComments(w http.ResponseWriter, req *http.Request) {
	var comments []model.Comment
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		comments = service.GetComments(handler.DEFAULT_START, handler.DEFAULT_END)
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		comments = service.GetComments(startSeq, endSeq)
	}
	w.Write(model.MarshalResponse(0, comments))

}

func (c *CommentHandler) AddComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var user model.User
	userName := req.Form.Get("username")
	password := req.Form.Get("password")
	user.Name = userName
	user.Password = password
	if !service.UserLogin(user) {
		w.Write(model.MarshalResponse(1, "用户登录失败"))
		return
	}
	service.UpdateUserLogin(user)

	content := req.Form.Get("commentContent")
	commentTime := time.Now()
	commentType := 0
	ip := req.RemoteAddr
	ua := req.UserAgent()

	comment := model.Comment{
		User:        user.Name,
		Content:     content,
		CommentTime: commentTime,
		CommentType: commentType,
		IP:          ip,
		UA:          ua}
	err := service.AddComment(user, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	w.Write(model.MarshalResponse(0, "success"))
}
