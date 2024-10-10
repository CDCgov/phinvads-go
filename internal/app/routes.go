package app

import (
	"net/http"

	"github.com/CDCgov/phinvads-go/internal/ui"
	"github.com/justinas/alice"
	"github.com/vearutop/statigz"
	"github.com/vearutop/statigz/brotli"
)

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/", statigz.FileServer(ui.Files, brotli.AddEncoding, statigz.EncodeOnInit))

	mux.HandleFunc("GET /", app.home)

	mux.HandleFunc("GET /api", app.healthcheck)

	mux.HandleFunc("GET /api/code-systems", app.getAllCodeSystems)
	mux.HandleFunc("GET /api/code-systems/{id}", app.getCodeSystemByID)

	mux.HandleFunc("GET /api/code-system-concepts", app.getAllCodeSystemConcepts)
	mux.HandleFunc("GET /api/code-system-concepts/{id}", app.getCodeSystemConceptByID)

	mux.HandleFunc("GET /api/value-sets", app.getAllValueSets)
	mux.HandleFunc("GET /api/value-sets/{id}", app.getValueSetByID)
	mux.HandleFunc("GET /api/value-sets/{oid}/versions", app.getValueSetVersionsByValueSetOID)

	mux.HandleFunc("GET /api/value-set-versions/{id}", app.getValueSetVersionByID)

	mux.HandleFunc("GET /api/views", app.getAllViews)
	mux.HandleFunc("GET /api/views/{id}", app.getViewByID)

	mux.HandleFunc("GET /api/view-versions/{id}", app.getViewVersionByID)
	mux.HandleFunc("GET /api/view-versions-by-view/{viewId}", app.getViewVersionsByViewID)

	mux.HandleFunc("GET /api/value-set-concepts/{id}", app.getValueSetConceptByID)
	mux.HandleFunc("GET /api/value-set-concepts/value-set-version/{valueSetVersionId}", app.getValueSetConceptsByVersionID)
	mux.HandleFunc("GET /api/value-set-concepts/code-system/{codeSystemOid}", app.getValueSetConceptsByCodeSystemOID)

	mux.HandleFunc("GET /r5/CodeSystem/{id}", app.getFHIRCodeSystemByID)

	mux.HandleFunc("GET /toggle-banner/{action}", app.handleBannerToggle)
	mux.HandleFunc("GET /load-hot-topics", app.getAllHotTopics)

	mux.HandleFunc("POST /api/search", app.formSearch)
	mux.HandleFunc("GET /search", app.directSearch)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
