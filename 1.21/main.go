package main

import "fmt"

// Целевой интерфейс, который ожидает клиент
type Logger interface {
	Println(msg string)
}

// Существующий несовместимый тип ExternalLogLibrary
type ExternalLogLibrary struct{}

func (e *ExternalLogLibrary) Log(message string) {
	fmt.Println("External Log:", message)
}

// Адаптер
type LogAdapter struct {
	extLogger *ExternalLogLibrary
}

// Адаптер реализует интерфейс Logger
func (a *LogAdapter) Println(msg string) {
	a.extLogger.Log(msg) //делегирует вызов методу Log
}

func doSomething(logger Logger) {
	logger.Println("Hello from adapter pattern!")
}

func main() {
	external := &ExternalLogLibrary{}

	adapter := &LogAdapter{extLogger: external}

	doSomething(adapter)
}
