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

var templates = template.Must(template.ParseGlob("templates/*"))

var (
	indexTemplate = template.Must(template.New("layout.html").ParseFS(files, "templates/layout.html", "templates/index.html"))
	aboutTemplate = template.Must(template.New("layout.html").ParseFS(files, "templates/layout.html", "templates/about.html"))
)

type Msg struct {
	Msg string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	newMsg := "World"
	err := indexTemplate.Execute(w, &Msg{Msg: newMsg})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "clicked.html", nil)
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
