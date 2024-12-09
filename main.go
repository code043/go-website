package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/skills", home)
	http.HandleFunc("/about", home)
	http.HandleFunc("/contact", home)

	// Start
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}
func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}
func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html")
}
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)

}
