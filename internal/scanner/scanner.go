package scanner

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
)

// Recording represents a recording
type Recording struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Scan a directory and return all of the files
func Scan(path string) []Recording {
	files := []Recording{}

	filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		hash := md5.Sum([]byte(info.Name()))
		id := hex.EncodeToString(hash[:])

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			files = append(files, Recording{
				ID:   id,
				Name: info.Name(),
			})
		}

		return nil
	})

	return files
}
