package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path"
)

const (
	Port = ":8080"
)

func Greeting(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var filepath = path.Join("", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Greeting",
		"name":  ps.ByName("name"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/:name", Greeting)

	http.ListenAndServe(Port, router)
}
