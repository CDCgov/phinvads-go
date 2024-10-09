package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) clientRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("internal/ui/assets"))))
	r.HandleFunc("/search", app.directSearch).Methods("GET")
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/toggle-banner/{action}", app.handleBannerToggle).Methods("GET")
	r.HandleFunc("/load-hot-topics", app.getAllHotTopics).Methods("GET")
	return r
}

func (app *Application) fhirRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/r5/CodeSystem/{id}", app.getFHIRCodeSystemByID).Methods("GET")
	return r
}

// The routes() method returns a servemux containing our application routes.
func (app *Application) apiRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api", app.healthcheck).Methods("GET")

	r.HandleFunc("/api/code-systems", app.getAllCodeSystems).Methods("GET")
	r.HandleFunc("/api/code-systems/{id}", app.getCodeSystemByID).Methods("GET")

	r.HandleFunc("/api/code-system-concepts", app.getAllCodeSystemConcepts).Methods("GET")
	r.HandleFunc("/api/code-system-concepts/{id}", app.getCodeSystemConceptByID).Methods("GET")

	r.HandleFunc("/api/value-sets", app.getAllValueSets).Methods("GET")
	r.HandleFunc("/api/value-sets/{id}", app.getValueSetByID).Methods("GET")
	r.HandleFunc("/api/value-sets/{oid}/versions", app.getValueSetVersionsByValueSetOID).Methods("GET")

	r.HandleFunc("/api/value-set-versions/{id}", app.getValueSetVersionByID).Methods("GET")

	r.HandleFunc("/api/views", app.getAllViews).Methods("GET")
	r.HandleFunc("/api/views/{id}", app.getViewByID).Methods("GET")

	r.HandleFunc("/api/view-versions/{id}", app.getViewVersionByID).Methods("GET")
	r.HandleFunc("/api/view-versions-by-view/{viewId}", app.getViewVersionsByViewID).Methods("GET")

	r.HandleFunc("/api/value-set-concepts/{id}", app.getValueSetConceptByID).Methods("GET")
	r.HandleFunc("/api/value-set-concepts/value-set-version/{valueSetVersionId}", app.getValueSetConceptsByVersionID).Methods("GET")
	r.HandleFunc("/api/value-set-concepts/code-system/{codeSystemOid}", app.getValueSetConceptsByCodeSystemOID).Methods("GET")

	r.HandleFunc("/api/search", app.formSearch).Methods("POST")

	return r
}
