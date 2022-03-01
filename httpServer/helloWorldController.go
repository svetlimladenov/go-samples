package main

import (
	"github.com/svetlimladenov/go-samples/httpServer/customHttp"
)

type HelloWorldController struct {
}

func (c HelloWorldController) Handle(req customHttp.HttpRequest) (customHttp.HttpResponse, error) {
	return customHttp.HttpResponse{
		Code: 200,
		Body: "<h1>Hello World</h1>",
	}, nil
}

func (c HelloWorldController) Read(buf []byte) (n int, err error) {
	buf[1] = 1
	return 1, nil
}
