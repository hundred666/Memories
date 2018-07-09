package util

import (
	"os"
	"bufio"
	"io"
	"log"
	"strings"
)

func ParseWebConf() (map[string]string, error) {
	conf := make(map[string]string)
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
        line, err := rd.ReadString('\n')
        kv := strings.Split(line, "=")
		if len(kv) != 2 {
			continue
		}
		conf[kv[0]] = kv[1]
        if err != nil {
            if err == io.EOF {
                log.Println("file read")
                break
            } else {
                log.Println("file read error")
                break
            }
        }
    }

	return conf, nil
}
