package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/teslapi/teslapi/cmd/teslapid/api"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", serveTemplate)
	api.Routes()

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Path
	if t == "/" {
		t = "index.html"
	}

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(t))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
