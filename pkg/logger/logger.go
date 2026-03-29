package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var logLevel = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
)

var logLevelColors = map[LogLevel]string{
	DEBUG: colorCyan,
	INFO:  colorGreen,
	WARN:  colorYellow,
	ERROR: colorRed,
	FATAL: colorPurple,
}

type Logger struct {
	mu          sync.Mutex
	logger      *log.Logger
	serviceName string
}

func NewLogger(loggerName string) (*Logger, error) {
	return &Logger{
		logger:      log.New(os.Stdout, "", 0),
		serviceName: loggerName,
	}, nil
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := fmt.Sprintf(format, args...)
	l.logger.Printf("%s[%v][%v] %v\n", logLevelColors[level], logLevel[level], l.serviceName, message)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
	os.Exit(1)
}
