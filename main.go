package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob(filepath.Join("templates/* .html")))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandle).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
