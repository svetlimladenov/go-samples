package logger

import "os"

type FileLogger struct {
	File string
}

func (fl FileLogger) Log(msg string) {
	fd, err := os.Open(fl.File)

	if err != nil {
		panic(err.Error())
	}

	defer fd.Close()

	fd.Write([]byte(msg))
}
