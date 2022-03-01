package main

import (
	"github.com/svetlimladenov/go-samples/httpServer/customHttp"
)

const (
	PORT = 3000
	HOST = "0.0.0.0"
)

var router = map[string]customHttp.Controller{
	"/hello-world": HelloWorldController{},
}

func main() {
	customHttp.Listen(PORT, router)
}
