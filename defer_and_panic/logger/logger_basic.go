package logger

import (
	"fmt"
	"os"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// String возвращает строковое представление уровня
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

var minLevel Level = DEBUG

func SetMinLevel(level Level) {
	minLevel = level
}
func Log(level Level, msg string, args ...any) {
	if level < minLevel {
		return
	}

	timestamp := time.Now().UTC().Format(time.RFC3339)
	// Формируем сообщение
	output := fmt.Sprintf("%s [%s] %s", timestamp, level.String(), msg)
	if len(args) > 0 {
		output += fmt.Sprintf("%v", args)
	}
	fmt.Fprintln(os.Stdout, output)
}
func Debug(msg string, args ...any) {
	Log(DEBUG, msg, args...)
}
func Info(msg string, args ...any) {
	Log(INFO, msg, args...)
}
func Warning(msg string, args ...any) {
	Log(WARNING, msg, args...)
}
func Error(msg string, args ...any) {
	Log(ERROR, msg, args...)
}
func Fatal(msg string, args ...any) {
	Log(FATAL, msg, args...)
	panic(fmt.Sprintf(msg))
}
