//main_test.go

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHandler(t *testing.T) {
	// expected body
	expected := gin.H{
		"message": "fortnite",
	}

	router := SetupRouter()
	response := performRequest(router, "GET", "/epic")

	assert.Equal(t, http.StatusOK, response.Code)

	// Convert the JSON response to a map
	var jsonResponse map[string]string
	err := json.Unmarshal([]byte(response.Body.Bytes()), &jsonResponse)
	assert.Nil(t, err)

	value, exists := jsonResponse["message"]

	// Make some assertions on the correctness of the response.
	assert.True(t, exists)
	assert.Equal(t, expected["message"], value)
}
