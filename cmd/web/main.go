package main

import (
	"net/http"
	"os"

	"github.com/goququ/snippetbox/cmd/web/logger"
)

func main() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	port = ":" + port

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

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
