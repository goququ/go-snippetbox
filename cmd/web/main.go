package main

import (
	"log"
	"net/http"

	"github.com/goququ/snippetbox/cmd/web/db"
	"github.com/goququ/snippetbox/cmd/web/logger"
	"github.com/goququ/snippetbox/cmd/web/utils"
	"github.com/goququ/snippetbox/pkg/models/psql"
)

type application struct {
	logError    *log.Logger
	logInfo     *log.Logger
	snippets    *psql.SnippetModel
	projectRoot string
}

func main() {
	port := utils.GetPort()

	myDB, err := db.Open()
	if err != nil {
		logger.Error.Fatal(err)
	}
	defer myDB.Close()

	projectRoot, err := utils.GetProjectRoot()
	if err != nil {
		logger.Error.Fatal(err)
	}

	app := &application{
		logError:    logger.Error,
		logInfo:     logger.Info,
		snippets:    &psql.SnippetModel{DB: myDB},
		projectRoot: projectRoot,
	}

	server := &http.Server{
		Addr:     port,
		ErrorLog: app.logError,
		Handler:  app.routes(),
	}

	app.logInfo.Printf("Server listening on port %s", port)

	if err := server.ListenAndServe(); err != nil {
		app.logError.Fatal(err)
	}
}
