package util

import (
	"path/filepath"
	"os"
)

func GetWebPath() (string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	webPath := filepath.Join(dir, "src")
	_, err = os.Stat(webPath)
	if err == nil {
		return webPath
	}
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	webPath = filepath.Join(dir, "src")
	_, err = os.Stat(webPath)
	if err == nil {
		return webPath
	}
	return ""
}

func GetConfPath() (string, error) {

	webPath := GetWebPath()
	confPath := filepath.Join(webPath, "conf", "web.conf")
	return confPath, nil
}

func GetStaticPath() (string, error) {

	webPath := GetWebPath()

	staticPath := filepath.Join(webPath, "static")
	return staticPath, nil
}

func GetViewPath() (string, error) {
	webPath := GetWebPath()
	viewPath := filepath.Join(webPath, "view")
	return viewPath, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
