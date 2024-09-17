package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CDCgov/phinvads-fhir/internal/app"
)

func TestCommonHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock HTTP handler that we can pass to our commonHeaders
	// middleware, which writes a 200 status code and an "OK" response body.
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	// Pass the mock HTTP handler to our commonHeaders middleware. Because
	// commonHeaders *returns* a http.Handler we can call its ServeHTTP()
	// method, passing in the http.ResponseRecorder and dummy http.Request to
	// execute it.
	app.commonHeaders(next).ServeHTTP(rr, r)
	rs := rr.Result()
	// Check that the middleware has correctly set the Content-Security-Policy header on the response.
	expectedValue := "default-src 'self'; style-src 'self'"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expectedValue)
	// Check that the middleware has correctly set the Referrer-Policy header on the response.
	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, rs.Header.Get("Referrer-Policy"), expectedValue)
	// Check that the middleware has correctly set the X-Content-Type-Options header on the response.
	expectedValue = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expectedValue)
	// Check that the middleware has correctly set the X-Frame-Options header on the response.
	expectedValue = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), expectedValue)
	// Check that the middleware has correctly set the X-XSS-Protection header on the response
	expectedValue = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expectedValue)
	// Check that the middleware has correctly set the Server header on the response.
	expectedValue = "Go"
	assert.Equal(t, rs.Header.Get("Server"), expectedValue)
	// Check that the middleware has correctly called the next handler in line and the response status code and body are as expected.
	assert.Equal(t, rs.StatusCode, http.StatusOK)
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
