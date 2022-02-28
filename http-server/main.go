package main

import "fmt"

const (
	PORT = 3000
	HOST = "0.0.0.0"
)

func main() {
	fmt.Printf("Server listening on %v:%v\n", HOST, PORT)
}
