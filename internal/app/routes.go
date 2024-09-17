package app

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", app.healthcheck)

	mux.HandleFunc("GET /r5/CodeSystem/{id}", app.codesystem)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
