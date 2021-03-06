package main

import (
	"log"
	"net/http"
)

const port = ""

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
