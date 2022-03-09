package log

import (
	"fmt"
	"time"
)

type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(msg string) {
	fmt.Printf("[%v] %v", time.Now().Format("DD:MM:YYYY")))
}