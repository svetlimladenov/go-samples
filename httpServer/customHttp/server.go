package customHttp

import (
	"fmt"
	"net"
)

func Listen(port int, routing map[string]Controller) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		fmt.Println("Cannot open connection to :8000")
	}

	fmt.Printf("Listening on %v", listener.Addr().String())

	for {
		conn, err := listener.Accept() // blocks and waits
		fmt.Println("Accepting new connection")
		if err != nil {
			fmt.Println("Corrupted connection")
			continue
		}

		buffer := make([]byte, 10000)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Cannot read the request")
			continue
		}

		req, err := parseHttpRequest(string(buffer[:]))
		if err != nil {
			fmt.Println("Something went wrong, can't parse the request")
			continue
		}

		fmt.Println(req)
		controller, exist := routing[req.Path]

		if !exist {
			// respond with 404 not found
			fmt.Println("Error 404 Not Found")
			continue
		}

		res, err := controller.Handle(req)

		if err != nil {
			// response with 500
			fmt.Println("Error 500 Internal Server Error")
			continue
		}

		fmt.Printf("res: %v\n", res)
	}

	return nil
}
