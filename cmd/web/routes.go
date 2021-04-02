package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	standartMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))
	mux.Get("/snippet/create/new", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standartMiddleware.Then(mux)
}
