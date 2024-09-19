package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"

	"github.com/CDCgov/phinvads-fhir/internal/app/fhir/r5"
	"github.com/CDCgov/phinvads-fhir/internal/database/models/xo"
	customErrors "github.com/CDCgov/phinvads-fhir/internal/errors"
)

func (app *Application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: OK"))
}

func (app *Application) codesystem(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")
	id_type, err := determineIdType(id)
	if err != nil {
		customErrors.BadRequest(w, r, err, app.logger)
		return
	}

	var codeSystem *xo.CodeSystem
	if id_type == "oid" {
		codeSystem, err = rp.GetCodeSystemByOID(r.Context(), id)
	} else {
		codeSystem, err = rp.GetCodeSystemByID(r.Context(), id)
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

	fhirCodeSystem, err := r5.SerializeCodeSystemToFhir(codeSystem)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}

	marshaller, err := jsonformat.NewMarshaller(false, "", "", fhirversion.R4)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}

	fhirJson, err := marshaller.MarshalResource(fhirCodeSystem)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
	}

	w.Write(fhirJson)
}
