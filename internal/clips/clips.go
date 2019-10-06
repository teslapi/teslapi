package clips

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/teslapi/teslapi/internal/parse"
)

const dbPath = "../../storage/db"

// Clip is a recording from the USB drive, this can be a local file or remote file
type Clip struct {
	ID             string    `json:"id"`              // the id of the file, generated using the hash of the filename
	Type           string    `json:"type"`            // could be recent or saved
	Name           string    `json:"name"`            // e.g. 2019-03-04_17-20-front.mp4
	Camera         string    `json:"camera"`          // e.g. front, right, left
	FileLocation   string    `json:"file_location"`   // the root path of the file on disk, even if it is remote
	RemoteLocation string    `json:"remote_location"` // the URL or path to the file on the cloud storage provider
	Uploaded       bool      `json:"uploaded"`        // if the clip has been uploaded or not
	FileTimestamp  time.Time `json:"file_timestamp"`  // timestamp of the file
	DateUploaded   time.Time `json:"date_uploaded"`   // the date and time that the clip was uploaded to the cloud storage provider
}

// Import will import all of the recordings into a local database
func Import(path string) error {
	files := []os.FileInfo{}
	filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			files = append(files, info)
		}

		return nil
	})

	for _, f := range files {
		// get the clip type
		_, err := Save(Clip{
			Name:          f.Name(),
			Camera:        parse.CameraFromRecording(f.Name()),
			FileTimestamp: f.ModTime(),
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// Save takes a clip and persists to the database
func Save(c Clip) (Clip, error) {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return c, errors.New("save has not been implemented yet")
}

// Handle will handle the API response
func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Upload will upload an individual clip to a remote URL
func Upload(c Clip) error {
	return errors.New("upload has not been implemented")
}
