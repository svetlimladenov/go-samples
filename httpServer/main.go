package httpServer

import (
	"fmt"

	"github.com/svetlimladenov/go-samples/http-server/customHttp"
)

const (
	PORT = 3000
	HOST = "0.0.0.0"
)

func main() {
	fmt.Printf("Server listening on %v:%v\n", HOST, PORT)
	customHttp.Listen()
}
