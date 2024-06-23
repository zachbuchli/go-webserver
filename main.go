package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed templates
var templateDir embed.FS

//go:embed static
var staticsDir embed.FS

const rootPath = "files"

type FbFile struct {
	Name string
	Path string
}

type Fb struct {
	CurrentPath string
	SubPaths    []FbFile
	IsDirectory bool
}

var templates = map[string]*template.Template{
	"index": template.Must(template.ParseFS(templateDir, "templates/layout.html", "templates/index.html")),
	"fb":    template.Must(template.ParseFS(templateDir, "templates/layout.html", "templates/fb.html")),
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["index"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func filePathHandler(w http.ResponseWriter, r *http.Request) {
	currPath := r.PathValue("path")
	fullPath := filepath.Join(rootPath, currPath)
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		// going to assume this means file not found.
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	fb := Fb{CurrentPath: currPath}
	if fileInfo.IsDir() {
		fb.IsDirectory = true
		dir, err := os.Open(fullPath)
		defer dir.Close()
		if err != nil {
			http.Error(w, "error opening dir", http.StatusInternalServerError)
			return
		}
		files, err := dir.Readdir(-1)
		if err != nil {
			http.Error(w, "error reading dir files", http.StatusInternalServerError)
			return
		}
		subs := make([]FbFile, len(files))
		for i, f := range files {
			subs[i].Path = currPath + "/" + f.Name()
			subs[i].Name = f.Name()
		}
		fb.SubPaths = subs
	}

	err = templates["fb"].Execute(w, fb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	http.Handle("/static/", http.FileServer(http.FS(staticsDir)))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/fb/{path...}", filePathHandler)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
