package portrait

import (
	"net/http"
	"controller"
	"html/template"
	"os"
	"io"
	"path/filepath"
	"time"
	"model"
	"dao"
	"strconv"
)

type PortraitHandler struct {
}

func (p *PortraitHandler) View(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles(handler.GetView("portrait.html"))
	portraits := dao.GetPortraits(handler.DEFAULT_START, handler.DEFAULT_END)
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
	}
	t.Execute(w, portraits)

}

func (p *PortraitHandler) GetPortraits(w http.ResponseWriter, req *http.Request) {

	var portraits []model.Portrait
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		portraits = dao.GetPortraits(handler.DEFAULT_START, handler.DEFAULT_END)
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		portraits = dao.GetPortraits(startSeq, endSeq)
	}
	w.Write(model.MarshalResponse(0, portraits))

}

func (p *PortraitHandler) Add(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	userName := req.Form.Get("name")
	password := req.Form.Get("password")
	user := model.User{Name: userName, Password: password, LoginTime: time.Now()}

	if !dao.UploadPermission(user) {
		w.Write(model.MarshalResponse(1, "用户身份不对"))
		return
	}

	file, head, err := req.FormFile("file")
	if err != nil {
		w.Write(model.MarshalResponse(1, err.Error()))
		return
	}
	defer file.Close()

	photoName := req.Form.Get("portraitName")
	upTime := time.Now()
	photoPath := "/static/portraits/" + head.Filename
	upUser := userName

	//创建文件
	fW, err := os.Create(filepath.Join(handler.PortraitPath, head.Filename))
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

	portrait := model.Portrait{Name: photoName, Path: photoPath, UpTime: upTime, UpUser: upUser}
	dao.AddPortrait(portrait)
	dao.UpdateUserLogin(user)

	w.Write(model.MarshalResponse(0, "success"))
}
