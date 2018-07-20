package announce

import (
	"net/http"
	"service"
	"model"
	"time"
	"handler"
	"strconv"
)

type AnnounceHandler struct {
}

const ANNOUNCE = "ANNOUNCE"

func (a *AnnounceHandler) GetLatestAnnounce(w http.ResponseWriter, req *http.Request) {
	announce := service.GetLatestAnnounce()
	period := time.Now().Sub(announce.Time)
	if announce.Id == 0 || announce.Display == false || period > handler.ExpireTime {
		w.Write(model.MarshalResponse(2, "not show"))
		return
	}
	w.Write(model.MarshalResponse(0, announce))
}

func (a *AnnounceHandler) GetAnnounce(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	announceId := vars.Get("announceId")
	if announceId == "" {
		w.Write(model.MarshalResponse(1, "announce id wrong"))
		return
	}
	id, _ := strconv.Atoi(announceId)
	comment := service.GetAnnounce(id)
	w.Write(model.MarshalResponse(0, comment))
}

func (a *AnnounceHandler) GetAnnounces(w http.ResponseWriter, req *http.Request) {
	var announces []model.Announce
	req.ParseForm()
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	if start == "" || end == "" {
		announces = service.GetAnnounces(handler.DEFAULT_START, handler.DEFAULT_END)
	} else {
		startSeq, _ := strconv.Atoi(start)
		endSeq, _ := strconv.Atoi(end)
		announces = service.GetAnnounces(startSeq, endSeq)
	}
	w.Write(model.MarshalResponse(0, announces))
}
