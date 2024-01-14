package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := Person{
			Name: "Bruce Wayne",
			Age:  34,
			City: "Gotham City",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedBody := `{"name":"Bruce Wayne","age":34,"city":"Gotham City"}`
	actualBody := strings.TrimSpace(rr.Body.String())

	if actualBody != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", actualBody, expectedBody)
	}
}
