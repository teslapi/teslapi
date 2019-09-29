package scanner

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"time"

	"github.com/teslapi/teslapi/internal/parse"
)

// Recording represents a recording
type Recording struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	Camera    string    `json:"camera"`
}

// Scan a directory and return all of the files
func Scan(path string) []Recording {
	files := []Recording{}

	filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// hash the filename for an ID
		hash := md5.Sum([]byte(info.Name()))
		id := hex.EncodeToString(hash[:])

		// determine which camera this is for
		camera := parse.CameraFromRecording(info.Name())

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			files = append(files, Recording{
				ID:        id,
				Name:      info.Name(),
				Timestamp: info.ModTime(),
				Camera:    camera,
			})
		}

		return nil
	})

	return files
}
