package main

import "fmt"

type Logger interface {
	Log(message string)
}

func Process(l Logger) {
	l.Log("app started")
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("[CONSOLE]:", message)
}

type FileLogger struct {
	filename string
}

func (f FileLogger) Log(message string) {
	fmt.Println("[FILE]:", message)
}

func main() {

	// Console logger
	consoleLogger := ConsoleLogger{}
	Process(consoleLogger)

	// File logger
	fileLogger := FileLogger{filename: "app.log"}
	Process(fileLogger)

	// Можно напрямую вызывать
	consoleLogger.Log("console direct call")
	fileLogger.Log("file direct call")
}
