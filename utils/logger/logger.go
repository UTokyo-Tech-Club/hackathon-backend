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

func NewLogger() *Logger {
	l := log.New(os.Stdout, "", 2|log.Lshortfile) // include timestamp and line number
	return &Logger{
		logger: l,
		level:  setLogLevel(),
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.level >= ERROR {
		l.logger.Output(2, "[FATAL] "+fmt.Sprintln(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level >= ERROR {
		l.logger.Output(2, "[ERROR] "+fmt.Sprintln(v...))
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if l.level >= WARNING {
		l.logger.Output(2, "[WARNING] "+fmt.Sprintln(v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level >= INFO {
		l.logger.Output(2, "[INFO] "+fmt.Sprintln(v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level >= DEBUG {
		l.logger.Output(2, "[DEBUG] "+fmt.Sprintln(v...))
	}
}
