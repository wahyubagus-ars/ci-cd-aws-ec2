package main

import (
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"             // Import gorilla/mux for route matching
	"github.com/stretchr/testify/assert" // for assertions
)

func TestGreeting(t *testing.T) {
	// Define expected data for the template
	expectedData := map[string]interface{}{
		"title": "Greeting",
		"name":  "World",
	}

	// Create a mock response writer to capture the response
	recorder := httptest.NewRecorder()

	// Define a mock request with the name parameter
	req, err := http.NewRequest(http.MethodGet, "/John", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock template that always returns nil error (replace with your actual template parsing logic)
	mockTemplate := template.New("index.html")
	var tmpl bytes.Buffer
	mockTemplate.Execute(&tmpl, expectedData) // Simulate template execution

	// Define a handler function that uses the mock template
	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl.WriteTo(w) // Write the pre-rendered template to the mock response writer
	}

	// Create a router and register the handler for the route
	router := mux.NewRouter()
	router.HandleFunc("/{name}", handler)

	// Serve the request using the matched route handler
	router.ServeHTTP(recorder, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code, "Unexpected status code")

	// Assert the response body contains the expected data
	expectedBody := "" // Replace with expected output based on your template
	assert.Contains(t, recorder.Body.String(), expectedBody, "Unexpected response body")
}
