package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teslapi/teslapi/internal/recordings"
)

type ClipsPageData struct {
	Title string
	Clips []recordings.Clip
}

func main() {
	// scan the directory for clips
	clips := []recordings.Clip{}
	err := filepath.Walk("./storage/TeslaUSB", func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			clips = append(clips, recordings.Clip{
				Name: info.Name(),
			})
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		tmpl.Execute(w, ClipsPageData{
			Title: "Clips",
			Clips: clips,
		})
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("%v", err)
	}
}
