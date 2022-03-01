package customHttp

type Controller interface {
	Handle(req HttpRequest) (HttpResponse, error)
}
