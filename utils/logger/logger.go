package logger

import (
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

type Logger struct {
	logger *log.Logger
	level  int
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		level:  setLogLevel(),
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.level >= ERROR {
		l.logger.SetPrefix("[FATAL] ")
		l.logger.Fatalln(v...)
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level >= ERROR {
		l.logger.SetPrefix("[ERROR] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if l.level >= WARNING {
		l.logger.SetPrefix("[WARNING] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level >= INFO {
		l.logger.SetPrefix("[INFO] ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level >= DEBUG {
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Println(v...)
	}
}
