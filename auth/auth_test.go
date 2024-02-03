package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApiKey(t *testing.T) {
	// Test case 1: Valid Authorization header
	headers1 := http.Header{"Authorization": []string{"ApiKey my-api-key"}}
	apiKey1, err1 := GetApiKey(headers1)
	assert.NoError(t, err1, "Error should be nil")
	assert.Equal(t, "my-api-key", apiKey1, "API key should match")

	// Test case 2: No Authorization header
	headers2 := http.Header{}
	apiKey2, err2 := GetApiKey(headers2)
	assert.Error(t, err2, "Error should not be nil")
	assert.Equal(t, "", apiKey2, "API key should be empty")

	// Test case 3: Malformed Authorization header
	headers3 := http.Header{"Authorization": []string{"InvalidFormat"}}
	apiKey3, err3 := GetApiKey(headers3)
	assert.Error(t, err3, "Error should not be nil")
	assert.Equal(t, "", apiKey3, "API key should be empty")

	// Test case 4: Malformed first part of Authorization header
	headers4 := http.Header{"Authorization": []string{"InvalidType my-api-key"}}
	apiKey4, err4 := GetApiKey(headers4)
	assert.Error(t, err4, "Error should not be nil")
	assert.Equal(t, "", apiKey4, "API key should be empty")
}
