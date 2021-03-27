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
	mux := http.NewServeMux()
	port := utils.GetPort()

	app := &application{
		logError: logger.Error,
		logInfo:  logger.Info,
	}

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:     port,
		ErrorLog: logger.Error,
		Handler:  mux,
	}

	logger.Info.Printf("Server listening on port %s", port)
	err := server.ListenAndServe()

	if err != nil {
		logger.Error.Fatal(err)
	}
}
