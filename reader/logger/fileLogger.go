package logger

import (
	"fmt"
	"os"
)

type FileLogger struct {
	File string
}

func (fl FileLogger) Log(msg string) {
	// os.Open() // has mode O_RDONLY
	fd, err := os.OpenFile(fl.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		panic(err.Error())
	}

	defer fd.Close()

	_, err = fd.Write([]byte(formatLog(msg)))
	if err != nil {
		fmt.Println(err)
	}
}

func (cl FileLogger) String() string {
	return "FileLogger: File - " + cl.File
}
