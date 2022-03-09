package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/svetlimladenov/go-samples/reader/logger"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, err
}

func rot13(r byte) byte {
	if r >= 'a' && r <= 'z' {
		// Rotate lowercase letters 13 places.
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		// Rotate uppercase letters 13 places.
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	// Do nothing.
	return r
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (mr MyReader) Read(b []byte) (int, error) {
	for i := 0; i < cap(b); i++ {
		ascii := int('A')
		b[i] = byte(ascii)
	}

	return cap(b), nil
}

func emptyInterfaceSamples() {
	var i interface{} = "string here"
	describeEmptyInterface(i) // (<nil>, <nil>)

	str := i.(string)

	fmt.Println(str)

	i = logger.FileLogger{File: "text.txt"}

	describeEmptyInterface(i)

	appLogger, ok := i.(logger.FileLogger)
	fmt.Println(ok)

	appLogger.Log("TEST")
	fmt.Print(appLogger) // will try to call the String()
}

type CustomLogger struct{}

func (cl *CustomLogger) Log(msg string) {
	if cl == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println("I am looger, but not from the `logger` package")
}

func loggersDemo() {
	var appLogger logger.Logger

	describe(appLogger) // <nil>, <nil> - the value of the interface is nil, if we call .Log - error runtime

	// appLogger.Log("A") -> panic: runtime error: invalid memory address or nil pointer dereference

	// ConsoleLogger

	appLogger = logger.ConsoleLogger{}

	appLogger.Log("Testing the console logger")

	describe(appLogger)

	// FileLogger

	appLogger = logger.FileLogger{File: "text.txt"}

	appLogger.Log("Testing the file logger")

	describe(appLogger)

	// HttpLogger

	appLogger = &logger.HttpLogger{Address: "http://logserver.dev"} // we must use the & to make it work ? , because of how we implement the interface

	appLogger.Log("Http")

	describe(appLogger)

	// Custom Logger

	var customLogger *CustomLogger

	appLogger = customLogger

	appLogger.Log("test")

	describe(appLogger)

}

func describe(logger logger.Logger) {
	fmt.Printf("(%v, %T)\n", logger, logger)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
