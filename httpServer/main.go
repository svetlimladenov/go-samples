package main

import (
	"log"

	"github.com/svetlimladenov/go-samples/httpServer/customHttp"
)

const (
	PORT = 3000
)

var router = map[string]customHttp.Controller{
	"/hello-world": HelloWorldController{},
}

func main() {
	server := customHttp.Server{
		Port:   PORT,
		Router: router,
	}

	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
