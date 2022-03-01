package logger

import "fmt"

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(msg string) {
	fmt.Println(formatLog(msg))
}
