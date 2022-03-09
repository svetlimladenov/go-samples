package customHttp

import (
	"fmt"
	"net"
)

type Server struct {
	Port   int
	Router map[string]Controller
}

// In general all methods on a given type, should have either value or pointer receivers, but not both !!!
func (s *Server) Listen() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.Port))

	if err != nil {
		return err
	}

	fmt.Printf("Listening on %v", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Corrupted connection")
			continue
		}

		handleConnection(conn, s.Router)
	}
}

func handleConnection(conn net.Conn, router map[string]Controller) {
	reqString := read(conn)

	req, err := parseHttpRequest(reqString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	controller, exist := router[req.Path]

	if !exist {
		fmt.Println("Error 404 Not Found")
		return
	}

	if err != nil {
		fmt.Println("Error 500 Internal Server Error")
		return
	}

	res, err := controller.Handle(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write([]byte(formatResponse(res)))

	conn.Close()
}

func formatResponse(res HttpResponse) string {
	return `
HTTP/1.1 200 OK
Date: Mon, 27 Jul 2009 12:28:53 GMT
Server: Apache/2.2.14 (Win32)
Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
Content-Length: 88
Content-Type: text/html
Connection: Closed

<html>
<body>
<h1>Hello, World!</h1>
</body>
</html>
	`
}

func read(conn net.Conn) string {
	var reqString string
	for {
		buffer := make([]byte, 1000)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println(err)
			break
		}

		reqString += string(buffer[:n])

		if err != net.ErrClosed {
			break
		}
	}

	return reqString
}
