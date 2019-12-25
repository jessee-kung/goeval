package logger

import (
    "io"
    "log"
	"os"
	"bytes"
	"fmt"
)

var (
	// buf is the buffer space for global logger
	buf bytes.Buffer

	// Logger is the handle for a global logger
    Logger *log.Logger
)

// init initialize a global logger for whole project
func init() {
	Logger = log.New(&buf, "[GoEval] ", log.Ldate|log.Ltime)
	// Logger.SetOutput(io.MultiWriter(logFile, os.Stdout))
	Logger.SetOutput(io.MultiWriter(os.Stdout))
}

// Trace output the logs in trace level
func Trace(format string, args ...interface{}) {
	Logger.Println(fmt.Sprintf(format, args...))
}

// Info output the logs in info level
func Info(format string, args ...interface{}) {
	Logger.Println(fmt.Sprintf(format, args...))
}

// Warning output the logs in warning level
func Warning(format string, args ...interface{}) {
	Logger.Println(fmt.Sprintf(format, args...))
}

// Error output the logs in error level
func Error(format string, args ...interface{}) {
	Logger.Println(fmt.Sprintf(format, args...))
}

// Fatal output the logs in fatal level and call os.exit()
func Fatal(format string, args ...interface{}) {
	Logger.Fatalln(fmt.Sprintf(format, args...))
}