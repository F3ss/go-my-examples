package mylogger

import (
	"log"
	"os"
	"sync"
)

// MyLogger Singleton
type MyLogger struct {
	Info     *log.Logger
	Warrning *log.Logger
	Error    *log.Logger
}

var myLog *MyLogger
var once sync.Once

// New func for get *MyLogger
func New() *MyLogger {
	if myLog == nil {
		once.Do(func() {
			file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatalf("Error initialize myLogger, error: %s\n", err.Error())
			}
			myLog := &MyLogger{
				Info:     log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
				Warrning: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
				Error:    log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
			}
			myLog.Info.Printf("myLogger initialized")
		})
		return myLog
	}

	return myLog
}
