package main

import (
	"log"
	"net/http"

	"github.com/goququ/snippetbox/cmd/web/logger"
	"github.com/goququ/snippetbox/cmd/web/utils"
)

type application struct {
	logError *log.Logger
	logInfo  *log.Logger
}

func main() {
	port := utils.GetPort()

	app := &application{
		logError: logger.Error,
		logInfo:  logger.Info,
	}

	server := &http.Server{
		Addr:     port,
		ErrorLog: app.logError,
		Handler:  app.routes(),
	}

	app.logInfo.Printf("Server listening on port %s", port)
	err := server.ListenAndServe()

	if err != nil {
		app.logError.Fatal(err)
	}
}
