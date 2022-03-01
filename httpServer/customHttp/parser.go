package customHttp

import (
	"errors"
	"strings"
)

type Headers map[string]string

type HttpRequest struct {
	Method   string
	Path     string
	Protocol string
	Headers  Headers
	Body     string
}

type HttpResponse struct {
	Code     int
	Protocol string
	Date     string
	Headers  Headers
	Body     string
}

func parseHttpRequest(req string) (HttpRequest, error) {
	if req == "" {
		return HttpRequest{}, errors.New("the request body cannot be empty")
	}

	reqLines := strings.Split(req, "\n")

	method, path, protocol := parseReqLine(reqLines[0])

	return HttpRequest{
		Method:   method,
		Protocol: protocol,
		Path:     path,
	}, nil
}

func parseReqLine(line string) (string, string, string) {
	lineParts := strings.Split(line, " ")
	return lineParts[0], lineParts[1], lineParts[2]
}
