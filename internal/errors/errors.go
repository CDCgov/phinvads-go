package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/CDCgov/phinvads-go/internal/ui/components"
)

// Request Errors
type RequestError struct {
	Err    error
	Msg    string
	Method string
	Uri    string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("Request Error: %s", e.Msg)
}

var ErrInvalidId = fmt.Errorf("invalid identifier")

func BadRequest(w http.ResponseWriter, r *http.Request, err error, logger *slog.Logger) {
	logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// Database Errors
type DatabaseError struct {
	Err    error
	Msg    string
	Method string
	Id     string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Msg: %s. Err: %s", e.Msg, e.Err)
}

func (e *DatabaseError) NoRows(w http.ResponseWriter, r *http.Request, err error, logger *slog.Logger) {
	LogError(w, r, err, e.Msg, logger)
	http.Error(w, e.Msg, http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter, r *http.Request, err error, logger *slog.Logger) {
	LogError(w, r, err, "Server Error", logger)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func LogError(w http.ResponseWriter, r *http.Request, err error, errrorText string, logger *slog.Logger) {
	logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
}

func SearchError(w http.ResponseWriter, r *http.Request, err error, searchTerm string, logger *slog.Logger) {
	if errors.Is(err, sql.ErrNoRows) {
		errorString := fmt.Sprintf("Error: Code System %s not found", searchTerm)
		dbErr := &DatabaseError{
			Err:    err,
			Msg:    errorString,
			Method: "getCodeSystemById",
			Id:     searchTerm,
		}

		msg := fmt.Sprintf("Method %s returned an error while retrieving %s: %s", dbErr.Method, dbErr.Id, dbErr.Msg)
		LogError(w, r, dbErr.Err, msg, logger)

		component := components.Error("Search", dbErr.Msg)
		err := component.Render(r.Context(), w)
		if err != nil {
			LogError(w, r, err, "Failed to render search error component", logger)
		}
	} else {
		LogError(w, r, err, http.StatusText(http.StatusInternalServerError), logger)

		component := components.Error("search", err.Error())
		err := component.Render(r.Context(), w)
		if err != nil {
			LogError(w, r, err, "Failed to render search error component", logger)
		}
	}
}

// func (app *Application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }
