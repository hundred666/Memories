package util

import (
	"os"
	"bufio"
	"strings"
	"time"
	"io"
)

var conf map[string]string

func parseWebConf() (map[string]string, error) {
	if _, ok := conf["new"]; ok {
		return conf, nil
	}
	conf = make(map[string]string)

	confFile, err := GetConfPath()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(confFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		kv := strings.Split(string(line), "=")
		if len(kv) != 2 {
			continue
		}
		conf[kv[0]] = kv[1]
	}
	conf["new"] = "ok"

	return conf, nil
}

func WebAddr() (string) {
	if addr, ok := conf["Addr"]; ok {
		return addr
	}
	return ":701"
}

func ReadTimeout() (time.Duration) {
	if t, ok := conf["ReadTimeout"]; ok {
		if readTimeout, err := time.ParseDuration(t); err == nil {
			return readTimeout
		}
	}
	return 20 * time.Second
}

func WriteTimeout() (time.Duration) {
	if t, ok := conf["WriteTimeout"]; ok {
		if writeTimeout, err := time.ParseDuration(t); err == nil {
			return writeTimeout
		}
	}
	return 20 * time.Second
}

func WebConf() (map[string]string) {
	return conf
}

func Database() (database string, name string, password string) {
	if v, ok := conf["Database"]; ok {
		database = v
	}
	if v, ok := conf["Name"]; ok {
		name = v
	}
	if v, ok := conf["Password"]; ok {
		password = v
	}
	return
}
