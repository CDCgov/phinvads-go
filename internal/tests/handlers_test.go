package tests

import (
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	app := newTestApplication(t)
	//create test server with routes from the app, then shut the test server down after the test runs.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/api")

	Equal(t, code, http.StatusOK)
	Equal(t, body, "OK")
}
