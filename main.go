package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*
var files embed.FS

var (
	indexTemplate   = template.Must(template.New("layout.html").ParseFS(files, "templates/layout.html", "templates/index.html"))
	aboutTemplate   = template.Must(template.New("about.html").ParseFS(files, "templates/layout.html", "templates/about.html"))
	clickedTemplate = template.Must(template.New("clicked.html").ParseFS(files, "templates/clicked.html"))
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	err := clickedTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := aboutTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clicked", clickedHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
