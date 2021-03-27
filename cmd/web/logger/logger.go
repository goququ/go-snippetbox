package logger

import (
	"log"
	"os"
)

type AppLogger struct {
	LogError *log.Logger
	LogInfo  *log.Logger
}

var instance = AppLogger{
	log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile),
	log.New(os.Stderr, "Info\t", log.Ldate|log.Ltime),
}
