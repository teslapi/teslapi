package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ClipsPageData struct {
	Title string
	Clips []clip
}

type clip struct {
	Name           string    `json:"name"`            // e.g. 2019-03-04_17-20-front.mp4
	Type           string    `json:"type"`            // could be recent or saved
	Camera         string    `json:"camera"`          // e.g. front, right, left
	FileLocation   string    `json:"file_location"`   // the root path of the file on disk, even if it is remote
	RemoteLocation string    `json:"remote_location"` // the URL or path to the file on the cloud storage provider
	Uploaded       bool      `json:"uploaded"`        // if the clip has been uploaded or not
	FileTimestamp  time.Time `json:"file_timestamp"`  // timestamp of the file
	DateUploaded   time.Time `json:"date_uploaded"`   // the date and time that the clip was uploaded to the cloud storage provider
}

func main() {
	// scan the directory for clips
	files := []os.FileInfo{}
	err := filepath.Walk("./storage/TeslaUSB", func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			files = append(files, info)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// convert the files into clips
	clips := []clip{}
	for _, f := range files {
		clips = append(clips, clip{
			Name: f.Name(),
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ClipsPageData{
			Title: "Clips",
			Clips: clips,
		}
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, data)
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("%v", err)
	}
}
