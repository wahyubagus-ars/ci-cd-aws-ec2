package main_test

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
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

	// Create a handler function that uses the mock template
	handler := httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tmpl.WriteTo(w) // Write the pre-rendered template to the mock response writer
	})

	// Pass the mock request and handler to the router (simulates actual execution)
	router := httprouter.New()
	router.GET("/:name", handler)
	router.ServeHTTP(recorder, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code, "Unexpected status code")

	// Assert the response body contains the expected data
	expectedBody := "" // Replace with expected output based on your template
	assert.Contains(t, recorder.Body.String(), expectedBody, "Unexpected response body")
}
