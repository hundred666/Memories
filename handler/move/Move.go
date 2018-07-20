package move

import (
	"net/http"
	"time"
	"service"
	"model"
	"os"
	"handler"
	"io"
	"path/filepath"
	"strconv"
	"html/template"
)

type MoveHandler struct {
}

const MOVE = "MOVE"

func (m *MoveHandler) ViewMoves(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles(handler.GetView("moves.html"))
	moves := service.GetMoves(handler.DEFAULT_START, handler.DEFAULT_END, "ORDER BY Id DESC")
	t.Execute(w, moves)
}

func (m *MoveHandler) ViewMoveDetail(w http.ResponseWriter, req *http.Request) {

	vars := req.URL.Query()
	mId := vars.Get("mid")
	if mId == "" {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	id, _ := strconv.Atoi(mId)
	move := service.GetMove(id)
	t, _ := template.ParseFiles(handler.GetView("moveDetail.html"))
	t.Execute(w, move)
}

func (m *MoveHandler) AddMove(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	content := req.Form.Get("content")
	userName := req.Form.Get("username")
	password := req.Form.Get("password")
	user := model.User{Name: userName, Password: password, LoginTime: time.Now()}

	if !service.UserLogin(user) {
		w.Write(model.MarshalResponse(1, "用户登录失败"))
		return
	}
	file, head, err := req.FormFile("file")
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	defer file.Close()

	upTime := time.Now()

	imgPath := "/static/moves/" + head.Filename

	fW, err := os.Create(filepath.Join(handler.MovePath, head.Filename))
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	move := model.Move{Content: content, ImgPath: imgPath, User: userName, Time: upTime}
	service.AddMove(user, move)
	service.UpdateUserLogin(user)
	w.Write(model.MarshalResponse(0, "success"))

}

func (m *MoveHandler) GetMove(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	mId := req.Form.Get("mid")
	if mId == "" {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	id, _ := strconv.Atoi(mId)
	move := service.GetMove(id)
	w.Write(model.MarshalResponse(0, move))

}

func (m *MoveHandler) GetMoves(w http.ResponseWriter, req *http.Request) {
	var moves []model.Move
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		moves = service.GetMoves(handler.DEFAULT_START, handler.DEFAULT_END)
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		moves = service.GetMoves(startSeq, endSeq)
	}
	w.Write(model.MarshalResponse(0, moves))
}

func (m *MoveHandler) AddComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	userName := req.Form.Get("username")
	password := req.Form.Get("password")
	user := model.User{Name: userName, Password: password, LoginTime: time.Now()}

	if !service.UserLogin(user) {
		w.Write(model.MarshalResponse(1, "用户登录失败"))
		return
	}
	service.UpdateUserLogin(user)

	mId := req.Form.Get("moveId")

	content := req.Form.Get("commentContent")
	commentTime := time.Now()
	commentType := 1
	ip := req.RemoteAddr
	ua := req.UserAgent()

	comment := model.Comment{
		User:        user.Name,
		Content:     content,
		CommentTime: commentTime,
		CommentType: commentType,
		IP:          ip,
		UA:          ua}

	if mId == "" {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	id, _ := strconv.Atoi(mId)
	if service.GetMove(id).Id == 0 {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	move := model.Move{Id: id}
	err := service.AddMoveComment(user, move, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	w.Write(model.MarshalResponse(0, "success"))

}

func (m *MoveHandler) GetComments(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	mId := vars.Get("moveId")
	if mId == "" {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	id, _ := strconv.Atoi(mId)
	move := model.Move{Id: id}
	comments := service.GetMoveComments(move)
	w.Write(model.MarshalResponse(0, comments))
}
