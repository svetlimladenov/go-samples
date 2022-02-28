package customHttp

import "errors"

type HttpMethod int64

const (
	GET HttpMethod = iota
	POST
	PUT
	PATCH
	DELETE
)

type HttpRequst struct {
	Method HttpMethod
	Body   string
}

func parseHttpRequest(req string) (HttpRequst, error) {
	if req == "" {
		return HttpRequst{}, errors.New("the request body cannot be empty")
	}

	return HttpRequst{
		Method: GET,
		Body:   req,
	}, nil
}
