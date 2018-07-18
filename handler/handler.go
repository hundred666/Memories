package handler

import (
	"util"
	"log"
	"path/filepath"
	"time"
)

var ViewPath string
var AdminPath string
var PortraitPath string
var MovePath string
var ExpireTime time.Duration

const DEFAULT_START = 0
const DEFAULT_END = -1

func InitHD() {
	log.Println("init path")

	viewPath, err := util.GetViewPath()
	if err != nil {
		log.Fatal(err)
	}
	ViewPath = viewPath
	AdminPath = filepath.Join(viewPath, "admin")
	staticPath, err := util.GetStaticPath()
	if err != nil {
		log.Fatal(err)
	}
	PortraitPath = filepath.Join(staticPath, "portraits")
	MovePath = filepath.Join(staticPath, "moves")
	ExpireTime = util.ExpireTime()

}

func GetView(viewName string) string {
	return filepath.Join(ViewPath, viewName)
}

func GetAdmin(viewName string) string {
	return filepath.Join(AdminPath, viewName)
}
