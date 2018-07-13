package main

import (
	"server"
)

func main() {

	web := server.Server{}
	web.InitServer()
	web.Start()

}
