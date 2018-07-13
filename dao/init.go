package dao

import (
	"util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dataSource string

func InitDB() {
	log.Println("init db")
	database, name, password := util.Database()
	dataSource = fmt.Sprintf("%s:%s@/%s?charset=utf8&loc=Asia%%2FShanghai&parseTime=true", name, password, database)
}
