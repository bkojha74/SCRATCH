package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpLogger(t *testing.T) {
	// Create a buffer to capture log output
	var logBuffer bytes.Buffer

	// Create a mock HTTP handler for testing
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	// Create a request for testing
	req := httptest.NewRequest("GET", "https://example.com/test", nil)

	// Create an instance of the HttpLogger middleware with the logBuffer as the writer
	loggerMiddleware := HttpLogger(&logBuffer, mockHandler)

	// Serve the HTTP request with the middleware
	loggerMiddleware.ServeHTTP(httptest.NewRecorder(), req)

	// Check if the logging output matches the expected log statement
	expectedLog := "GET /test\n"
	assert.Equal(t, expectedLog, logBuffer.String(), "Logged output should match the expected log statement")
}
