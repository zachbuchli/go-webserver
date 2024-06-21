package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type Msg struct {
	Msg string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	newMsg := "World"
	err := templates.ExecuteTemplate(w, "index.html", &Msg{Msg: newMsg})
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
	err := templates.ExecuteTemplate(w, "about.html", nil)
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
