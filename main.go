package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
}
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)

}
