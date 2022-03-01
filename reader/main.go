package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileReader()
}

func fileReader() {
	f, err := os.Open("text")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 20)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n])) // :n so we read only the first n - read
		}
	}
}
