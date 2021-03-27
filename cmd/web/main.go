package main

import (
	"net/http"
	"os"
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
		Addr: port,
		// ErrorLog: logger.ErrorLogger,
		Handler: mux,
	}

	// logger.InfoLogger.Printf("Server listening on port %s", port)
	err := server.ListenAndServe()

	if err != nil {
		// logger.ErrorLogger.Fatal(err)
	}
}
