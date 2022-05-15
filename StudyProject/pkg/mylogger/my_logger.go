package mylogger

import (
	"log"
	"os"
	"sync"
)

type myLogger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

var myLog *myLogger
var once sync.Once

func NewMyLogger(whereToWrite string) *myLogger {
	if myLog == nil {
		if whereToWrite == "file" {
			once.Do(initLoggingToFile())
		}
		return myLog
	}

	return myLog
}

func initLoggingToFile() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error initialise myLogger, error: %s\n", err.Error())
	}
	myLog := &myLogger{
		infoLogger:    log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(file, "WARNING", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(file, "ERROR", log.Ldate|log.Ltime|log.Lshortfile),
	}
	myLog.infoLogger.Printf("myLogger initialised")
}

func initLoggingToConsole() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error initialise myLogger, error: %s\n", err.Error())
	}
	myLog := &myLogger{
		infoLogger:    log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(file, "WARNING", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(file, "ERROR", log.Ldate|log.Ltime|log.Lshortfile),
	}
	myLog.infoLogger.Printf("myLogger initialised")
}
