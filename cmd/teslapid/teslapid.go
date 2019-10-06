package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/teslapi/teslapi/internal/recordings"
)

type clipsPageData struct {
	Title string
	Clips []recordings.Clip
}

func main() {
	// scan the directory for clips
	clips, err := scan("./storage/TeslaUSB")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		tmpl.Execute(w, clipsPageData{
			Title: "Clips",
			Clips: clips,
		})
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func scan(dir string) ([]recordings.Clip, error) {
	clips := []recordings.Clip{}
	err := filepath.Walk(dir, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			// determine the type
			clipType := "saved"
			if regexp.MustCompile(`RecentClips`).MatchString(root) {
				clipType = "recent"
			}

			// get the camera that recorded the clip
			camera := "front"
			if regexp.MustCompile(`right_repeater`).MatchString(info.Name()) {
				camera = "right"
			}
			if regexp.MustCompile(`left_repeater`).MatchString(info.Name()) {
				camera = "left"
			}

			clips = append(clips, recordings.Clip{
				Name:          info.Name(),
				Type:          clipType,
				Camera:        camera,
				FileLocation:  root,
				FileTimestamp: info.ModTime(),
				Uploaded:      false,
			})
		}

		return nil
	})

	if err != nil {
		return clips, err
	}

	return clips, err
}
