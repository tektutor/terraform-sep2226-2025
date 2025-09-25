package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Logger interface
type Logger interface {
	Log(message string)
}

// ConsoleLogger prints to stdout
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), message)
}

// FileLogger writes to a file
type FileLogger struct {
	Filename string
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "[%s] %s\n", time.Now().Format(time.RFC3339), message)
}

// JSONLogger encodes messages as JSON
type JSONLogger struct{}

func (j JSONLogger) Log(message string) {
	entry := map[string]string{
		"time":    time.Now().Format(time.RFC3339),
		"message": message,
	}
	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
}

// App can use any logger
type App struct {
	logger Logger
}

func (a App) Run() {
	a.logger.Log("Application started")
	// Simulate work
	time.Sleep(500 * time.Millisecond)
	a.logger.Log("Application finished")
}

func main() {
	// Swap loggers here without touching App
	app1 := App{logger: ConsoleLogger{}}
	app2 := App{logger: FileLogger{Filename: "app.log"}}
	app3 := App{logger: JSONLogger{}}

	fmt.Println("== Console Logger ==")
	app1.Run()

	fmt.Println("\n== File Logger ==")
	app2.Run()
	fmt.Println("Check app.log file")

	fmt.Println("\n== JSON Logger ==")
	app3.Run()
}
