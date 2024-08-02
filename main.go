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

// render is a helper function for rendering templates to w a http.ResponseWriter.
// If the template doesnt exist or fails to render, it returns an http 500 error.
func render(w http.ResponseWriter, templateName string, data any) {
	err := templates[templateName].Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index", nil)
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "clicked", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "about", nil)
}

func msgHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.PathValue("msg")
	render(w, "message", msg)
}

func main() {

	http.Handle("/static/", http.FileServer(http.FS(statics)))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clicked", clickedHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/message/{msg}", msgHandler)

	fmt.Println("starting server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
