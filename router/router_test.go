package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestInitDB(t *testing.T) {
	godotenv.Load("../.env")

	// Test initialization of DBConfig
	dbConfig := initDB()

	// Assert that the returned value is not nil
	assert.NotNil(t, dbConfig)

	// Assert that the DBConfig is initialized successfully
	assert.NotNil(t, dbConfig.DB)
}

func TestNewRouter(t *testing.T) {
	godotenv.Load("../.env")

	// Create a new request to test the router
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router using the NewRouter function
	router := NewRouter()

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check if the status code is as expected
	assert.Equal(t, http.StatusOK, rr.Code)
}
