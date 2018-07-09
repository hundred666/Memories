package move

import (
	"net/http"
	"time"
	"dao"
	"model"
	"os"
	"controller"
	"io"
	"path/filepath"
	"strconv"
	"html/template"
	"log"
)

type MoveHandler struct {
}

const MOVE = "MOVE"

func (m *MoveHandler) ViewMoves(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles(handler.GetView("moves.html"))
	moves := dao.GetMoves(handler.DEFAULT_START, handler.DEFAULT_END,"ORDER BY Id DESC")
	t.Execute(w, moves)
}

func (m *MoveHandler)ViewMoveDetail(w http.ResponseWriter, req *http.Request){

	vars := req.URL.Query()
	mId:=vars.Get("mid")
	if mId == "" {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	id, _ := strconv.Atoi(mId)
	move := dao.GetMove(id)
	t, _ := template.ParseFiles(handler.GetView("moveDetail.html"))
	t.Execute(w, move)
}


func (m *MoveHandler) AddMove(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	content := req.Form.Get("content")
	userName := req.Form.Get("username")
	password := req.Form.Get("password")
	user := model.User{Name: userName, Password: password, LoginTime: time.Now()}

	if !dao.UserLogin(user) {
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
	dao.AddMove(move)
	dao.UpdateUserLogin(user)
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
	move := dao.GetMove(id)
	w.Write(model.MarshalResponse(0, move))

}

func (m *MoveHandler) GetMoves(w http.ResponseWriter, req *http.Request) {
	var moves []model.Move
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		moves = dao.GetMoves(handler.DEFAULT_START, handler.DEFAULT_END)
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		moves = dao.GetMoves(startSeq, endSeq)
	}
	w.Write(model.MarshalResponse(0, moves))
}

func (m *MoveHandler) AddComment(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	userName := req.Form.Get("username")
	log.Println(userName)
	password := req.Form.Get("password")
	log.Println(password)
	user := model.User{Name: userName, Password: password, LoginTime: time.Now()}

	if !dao.UserLogin(user) {
		w.Write(model.MarshalResponse(1, "用户登录失败"))
		return
	}
	dao.UpdateUserLogin(user)

	mId := req.Form.Get("mid")

	content := req.Form.Get("content")
	commentTime := time.Now()
	commentType := 1
	ip := req.RemoteAddr
	ua := req.UserAgent()

	comment := model.Comment{
		UName:       user.Name,
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
	if dao.GetMove(id).Id == 0 {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	move := model.Move{Id: id}
	err := dao.AddMoveComment(move, comment)
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	w.Write(model.MarshalResponse(0, "success"))

}

func (m *MoveHandler) GetComments(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	mId:=vars.Get("mid")
	if mId == "" {
		w.Write(model.MarshalResponse(1, "no resource"))
		return
	}
	id, _ := strconv.Atoi(mId)
	move := model.Move{Id: id}
	comments := dao.GetMoveComments(move)
	w.Write(model.MarshalResponse(0, comments))
}
