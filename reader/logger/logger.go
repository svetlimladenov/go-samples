package logger

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(msg string)
}

func formatLog(msg string) string {
	return fmt.Sprintf("[%v]:%v\n", time.Now().String(), msg)
}
