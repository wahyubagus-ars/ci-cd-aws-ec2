package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path"
)

const (
	Port = ":8080"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var filepath = path.Join("", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Greeting",
		"name":  vars["name"],
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{name}", Greeting)

	http.ListenAndServe(Port, router)
}
