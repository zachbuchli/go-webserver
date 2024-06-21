package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var files embed.FS

//go:embed static
var statics embed.FS

var templates = map[string]*template.Template{
	"index":   template.Must(template.ParseFS(files, "templates/layout.html", "templates/index.html")),
	"about":   template.Must(template.ParseFS(files, "templates/layout.html", "templates/about.html")),
	"message": template.Must(template.ParseFS(files, "templates/layout.html", "templates/message.html")),
	"clicked": template.Must(template.ParseFS(files, "templates/clicked.html")),
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["index"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["clicked"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["about"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func msgHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.PathValue("msg")
	err := templates["message"].Execute(w, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/static/", http.FileServer(http.FS(statics)))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clicked", clickedHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/message/{msg}", msgHandler)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
