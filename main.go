package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var templateDir embed.FS

//go:embed static
var staticsDir embed.FS

const rootPath = "./files"

var templates = map[string]*template.Template{
	"index": template.Must(template.ParseFS(templateDir, "templates/layout.html", "templates/index.html")),
	"fb":    template.Must(template.ParseFS(templateDir, "templates/layout.html", "templates/fp.html")),
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["index"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func filePathHandler(w http.ResponseWriter, r *http.Request) {
	currPath := r.PathValue("path")
	err := templates["fb"].Execute(w, currPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.Handle("/static/", http.FileServer(http.FS(staticsDir)))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/fb/{path...}", filePathHandler)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
