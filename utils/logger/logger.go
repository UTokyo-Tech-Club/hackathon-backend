package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FATAL = iota // 0
	ERROR
	WARNING
	INFO
	DEBUG
)

type Logger struct {
	logger *log.Logger
	level  int
}

var logger *Logger

func init() {
	logger = newLogger()
}

func newLogger() *Logger {
	l := log.New(os.Stdout, "", 2|log.Lshortfile) // include timestamp and line number
	return &Logger{
		logger: l,
		level:  setLogLevel(),
	}
}

func setLogLevel() int {
	logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch logLevel {
	case "FATAL":
		return FATAL
	case "ERROR":
		return ERROR
	case "WARNING":
		return WARNING
	case "INFO":
		return INFO
	case "DEBUG":
		return DEBUG
	default:
		return INFO
	}
}

func Fatal(v ...interface{}) {
	if logger.level >= ERROR {
		logger.logger.Output(2, "[FATAL] "+fmt.Sprintln(v...))
	}
}

func Error(v ...interface{}) {
	if logger.level >= ERROR {
		logger.logger.Output(2, "[ERROR] "+fmt.Sprintln(v...))
	}
}

func Warning(v ...interface{}) {
	if logger.level >= WARNING {
		logger.logger.Output(2, "[WARNING] "+fmt.Sprintln(v...))
	}
}

func Info(v ...interface{}) {
	if logger.level >= INFO {
		logger.logger.Output(2, "[INFO] "+fmt.Sprintln(v...))
	}
}

func Debug(v ...interface{}) {
	if logger.level >= DEBUG {
		logger.logger.Output(2, "[DEBUG] "+fmt.Sprintln(v...))
	}
}
