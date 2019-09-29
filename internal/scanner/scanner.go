package scanner

import (
	"os"
	"path/filepath"
)

// Recording represents a recording
type Recording struct {
	Name string
}

// Scan a directory and return all of the files
func Scan(path string) []Recording {
	files := []Recording{}

	filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false && info.Name() != ".DS_Store" {
			files = append(files, Recording{Name: info.Name()})
		}

		return nil
	})

	return files
}
