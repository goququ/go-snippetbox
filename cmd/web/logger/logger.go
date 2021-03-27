package logger

import (
	"log"
	"os"
)

var Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var Info = log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)
