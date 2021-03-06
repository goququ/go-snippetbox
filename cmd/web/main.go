package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fPort := strings.Join([]string{":", port}, "")

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server listening on port %s", fPort)
	err := http.ListenAndServe(fPort, mux)

	if err != nil {
		log.Fatal(err)
	}
}
