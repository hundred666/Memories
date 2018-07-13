package server

import (
	"net/http"
	"util"
	"router"
	"log"
	"dao"
	"handler"
)

type Server struct {
	BasePath   string
	Config     map[string]string
	HttpServer *http.Server
}

func (s *Server) InitServer() {
	log.Println("init server")

	dao.InitDB()
	handler.InitHD()
	router.InitRT()

	s.BasePath = util.GetWebPath()
	s.Config = util.WebConf()
	s.HttpServer = &http.Server{
		Addr:         util.WebAddr(),
		Handler:      router.Mux,
		ReadTimeout:  util.ReadTimeout(),
		WriteTimeout: util.WriteTimeout(),
	}
	log.Println("init server ok")
}

func (s *Server) Start() {
	log.Println("server start")
	log.Println(s.HttpServer.ListenAndServe())
}

func (s *Server) GetConf(key string) (string) {
	if p, ok := s.Config[key]; ok {
		return p
	}
	return ""
}
