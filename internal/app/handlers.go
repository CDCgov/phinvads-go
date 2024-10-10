package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/CDCgov/phinvads-go/internal/app/fhir/r5"
	"github.com/CDCgov/phinvads-go/internal/database/models"
	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
	customErrors "github.com/CDCgov/phinvads-go/internal/errors"
	"github.com/CDCgov/phinvads-go/internal/ui/components"
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
)

func (app *Application) healthcheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("status: OK"))
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getAllCodeSystems(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	codeSystems, err := rp.GetAllCodeSystems(r.Context())
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(codeSystems)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getCodeSystemByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")
	id_type, err := determineIdType(id)
	if err != nil {
		customErrors.BadRequest(w, r, err, app.logger)
		return
	}

	var codeSystem *xo.CodeSystem
	switch id_type {
	case Oid:
		codeSystem, err = rp.GetCodeSystemByOID(r.Context(), id)
	case Id:
		codeSystem, err = rp.GetCodeSystemByID(r.Context(), id)
	case Unknown:
		panic("unreachable!")
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getCodeSystemById",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(codeSystem)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getFHIRCodeSystemByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")
	id_type, err := determineIdType(id)
	if err != nil {
		customErrors.BadRequest(w, r, err, app.logger)
		return
	}

	var codeSystem *xo.CodeSystem
	switch id_type {
	case Oid:
		codeSystem, err = rp.GetCodeSystemByOID(r.Context(), id)
	case Id:
		codeSystem, err = rp.GetCodeSystemByID(r.Context(), id)
	case Unknown:
		panic("unreachable!")
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getCodeSystemById",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	conceptCount, err := models.GetCodeSystemConceptCount(r.Context(), app.db, codeSystem.Oid)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	concepts, err := rp.GetCodeSystemConceptsByCodeSystemOID(r.Context(), app.db, codeSystem)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fhirCodeSystem, err := r5.SerializeCodeSystemToFhir(codeSystem, conceptCount, concepts)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	marshaller, err := jsonformat.NewMarshaller(false, "", "", fhirversion.R4)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	fhirJson, err := marshaller.MarshalResource(fhirCodeSystem)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	_, err = w.Write(fhirJson)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getAllViews(w http.ResponseWriter, r *http.Request) {
	views, err := models.GetAllViews(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(views)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getViewByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	view, err := rp.GetViewByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(view)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getViewVersionByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	viewVersion, err := rp.GetViewVersionByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewVersionByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(viewVersion)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getViewVersionsByViewID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	viewId := r.PathValue("viewId")

	viewVersions, err := rp.GetViewVersionByViewId(r.Context(), viewId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", viewId)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewVersionsByViewID",
				Id:     viewId,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(viewVersions)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getAllCodeSystemConcepts(w http.ResponseWriter, r *http.Request) {
	codeSystemConcepts, err := models.GetAllCodeSystemConcepts(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(codeSystemConcepts)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getCodeSystemConceptByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")

	codeSystemConcept, err := rp.GetCodeSystemConceptByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System Concept%s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getCodeSystemConceptByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(codeSystemConcept)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getAllValueSets(w http.ResponseWriter, r *http.Request) {
	valueSets, err := models.GetAllValueSets(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSets)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getValueSetByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")
	id_type, err := determineIdType(id)
	if err != nil {
		customErrors.BadRequest(w, r, err, app.logger)
		return
	}

	var valueSet *xo.ValueSet
	switch id_type {
	case Oid:
		valueSet, err = rp.GetValueSetByOID(r.Context(), id)
	case Id:
		valueSet, err = rp.GetValueSetByID(r.Context(), id)
	case Unknown:
		panic("unreachable!")
	}

	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSet)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getValueSetVersionsByValueSetOID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	oid := r.PathValue("oid")

	valueSetVersions, err := rp.GetValueSetVersionByValueSetOID(r.Context(), oid)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set Versions with Value Set OID %s not found", oid)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetVersionsByValueSetOID",
				Id:     oid,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSetVersions)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getValueSetVersionByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")

	valueSetVersion, err := rp.GetValueSetVersionByID(r.Context(), id)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set Version %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetVersionsByValueSetOID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSetVersion)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getValueSetConceptByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	valueSetConcept, err := rp.GetValueSetConceptByID(r.Context(), id)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set Concept %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetConceptByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSetConcept)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getValueSetConceptsByVersionID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("valueSetVersionId")

	valueSetConcepts, err := rp.GetValueSetConceptByValueSetVersionID(r.Context(), id)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: No Value Set Concepts found for version %s", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetConceptsByVersionID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}

		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSetConcepts)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getValueSetConceptsByCodeSystemOID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	oid := r.PathValue("codeSystemOid")

	valueSetConcepts, err := rp.GetValueSetConceptsByCodeSystemOID(r.Context(), oid)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: No Value Set Concepts found for CodeSystem %s", oid)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetConceptsByCodeSystemOID",
				Id:     oid,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}

		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(valueSetConcepts)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) getAllHotTopics(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	hotTopics, err := rp.GetAllHotTopics(r.Context())
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := "Error: No Hot Topics found"
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "home: Get Hot Topics",
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}

		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	for _, t := range *hotTopics {
		// skip sending system config to the frontend
		if t.HotTopicName == "SYSTEM CONFIG" {
			continue
		}
		// format the sequence id to align with the uswds js controls
		divId := fmt.Sprintf("m-a%s", strconv.Itoa(t.Seq))

		component := components.HotTopic(t.HotTopicName, t.HotTopicDescription, divId, t.HotTopicID.String())
		err = component.Render(r.Context(), w)
		if err != nil {
			customErrors.ServerError(w, r, err, app.logger)
		}
	}
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	component := components.Home()
	err := component.Render(r.Context(), w)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) formSearch(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}
	searchTerm := r.Form["search"][0]
	searchType := r.Form["options"][0]
	app.search(w, r, searchTerm, searchType)
}

func (app *Application) directSearch(w http.ResponseWriter, r *http.Request) {
	searchType := r.URL.Query().Get("type")
	searchTerm := r.URL.Query().Get("input")
	app.search(w, r, searchTerm, searchType)
}

func (app *Application) search(w http.ResponseWriter, r *http.Request, searchTerm, searchType string) {
	rp := app.repository
	logger := app.logger

	result := &models.CodeSystemResultRow{}
	defaultPageCount := 5

	// retrieve code system
	codeSystem, err := rp.GetCodeSystemsByLikeOID(r.Context(), searchTerm)
	if err != nil || len(*codeSystem) < 1 {
		if err == nil {
			err = sql.ErrNoRows
		}
		customErrors.SearchError(w, r, err, searchTerm, logger)
		return
	}

	for _, cs := range *codeSystem {
		result.CodeSystems = append(result.CodeSystems, &cs)
	}

	if len(result.CodeSystems) <= defaultPageCount {
		defaultPageCount = len(result.CodeSystems)
	}

	result.PageCount = defaultPageCount
	result.CodeSystemsCount = strconv.Itoa(len(result.CodeSystems))

	// // retrieve concepts that are part of that code system
	// concepts, err := rp.GetCodeSystemConceptsByCodeSystemOID(r.Context(), app.db, codeSystem)
	// for _, csc := range *concepts {
	// 	result.CodeSystems = append(result.CodeSystems, &csc)
	// }
	// result.CodeSystemConcepts = concepts

	// for now
	result.CodeSystemConceptsCount = strconv.Itoa(0)

	// for now
	result.ValueSetsCount = strconv.Itoa(0)

	w.Header().Set("HX-Push-Url", fmt.Sprintf("/search?type=%s&input=%s", searchType, searchTerm))

	component := components.SearchResults(true, "Search", searchTerm, result)
	err = component.Render(r.Context(), w)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}

func (app *Application) handleBannerToggle(w http.ResponseWriter, r *http.Request) {
	action := r.PathValue("action")
	component := components.UsaBanner(action)
	err := component.Render(r.Context(), w)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}
}
