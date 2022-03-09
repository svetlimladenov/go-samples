package logger

import "fmt"

type HttpLogger struct {
	Address string
}

func (hl *HttpLogger) Log(msg string) {
	// send the log to some server with http
	fmt.Printf("Sending http req to %v\n", hl.Address)
}
