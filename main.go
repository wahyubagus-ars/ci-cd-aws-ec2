package main

import (
	"encoding/json"
	"net/http"
)

const (
	Port = ":8080"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	person := Person{
		Name: "Bruce Wayne",
		Age:  34,
		City: "Gotham City",
	}

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(person)
	})

	http.ListenAndServe(Port, nil)
}
