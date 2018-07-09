package handler

import (
	"util"
	"log"
	"path/filepath"
)

var ViewPath string
var AdminPath string
var PortraitPath string
var MovePath string

const DEFAULT_START = 0
const DEFAULT_END = -1

func init() {
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

}

func GetView(viewName string) string {
	return filepath.Join(ViewPath, viewName)
}

func GetAdmin(viewName string) string {
	return filepath.Join(AdminPath, viewName)
}
