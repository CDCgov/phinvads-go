package app

import (
	"github.com/CDCgov/phinvads-go/internal/ui"
	"github.com/gorilla/mux"
	"github.com/vearutop/statigz"
	"github.com/vearutop/statigz/brotli"
)

func (app *Application) setupClientRoutes(s *mux.Router) {
	fs := statigz.FileServer(ui.Files, brotli.AddEncoding, statigz.EncodeOnInit)
	s.PathPrefix("/assets/").Handler(fs)
	s.HandleFunc("/search", app.directSearch).Methods("GET")
	s.HandleFunc("/", app.home).Methods("GET")
	s.HandleFunc("/toggle-banner/{action}", app.handleBannerToggle).Methods("GET")
	s.HandleFunc("/load-hot-topics", app.getAllHotTopics).Methods("GET")
}

func (app *Application) setupFhirRoutes(s *mux.Router) {
	s.HandleFunc("/CodeSystem/{id}", app.getFHIRCodeSystemByID).Methods("GET")
}

// The routes() method returns a servemux containing our application routes.
func (app *Application) setupApiRoutes(s *mux.Router) {
	s.StrictSlash(true) // match on /api or /api/
	s.HandleFunc("/", app.healthcheck).Methods("GET")

	s.HandleFunc("/code-systems", app.getAllCodeSystems).Methods("GET")
	s.HandleFunc("/code-systems/{id}", app.getCodeSystemByID).Methods("GET")

	s.HandleFunc("/code-system-concepts", app.getAllCodeSystemConcepts).Methods("GET")
	s.HandleFunc("/code-system-concepts/{id}", app.getCodeSystemConceptByID).Methods("GET")

	s.HandleFunc("/value-sets", app.getAllValueSets).Methods("GET")
	s.HandleFunc("/value-sets/{id}", app.getValueSetByID).Methods("GET")
	s.HandleFunc("/value-sets/{oid}/versions", app.getValueSetVersionsByValueSetOID).Methods("GET")

	s.HandleFunc("/value-set-versions/{id}", app.getValueSetVersionByID).Methods("GET")

	s.HandleFunc("/views", app.getAllViews).Methods("GET")
	s.HandleFunc("/views/{id}", app.getViewByID).Methods("GET")

	s.HandleFunc("/view-versions/{id}", app.getViewVersionByID).Methods("GET")
	s.HandleFunc("/view-versions-by-view/{viewId}", app.getViewVersionsByViewID).Methods("GET")

	s.HandleFunc("/value-set-concepts/{id}", app.getValueSetConceptByID).Methods("GET")
	s.HandleFunc("/value-set-concepts/value-set-version/{valueSetVersionId}", app.getValueSetConceptsByVersionID).Methods("GET")
	s.HandleFunc("/value-set-concepts/code-system/{codeSystemOid}", app.getValueSetConceptsByCodeSystemOID).Methods("GET")

	s.HandleFunc("/search", app.formSearch).Methods("POST")
}
