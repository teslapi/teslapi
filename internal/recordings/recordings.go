package recordings

import (
	"os"
	"regexp"
	"time"
)

// Clip represents a specific recording
type Clip struct {
	Name           string    `json:"name"`            // e.g. 2019-03-04_17-20-front.mp4
	Type           string    `json:"type"`            // could be recent or saved
	Camera         string    `json:"camera"`          // e.g. front, right, left
	FileLocation   string    `json:"file_location"`   // the root path of the file on disk, even if it is remote
	RemoteLocation string    `json:"remote_location"` // the URL or path to the file on the cloud storage provider
	Uploaded       bool      `json:"uploaded"`        // if the clip has been uploaded or not
	FileTimestamp  time.Time `json:"file_timestamp"`  // timestamp of the file
	DateUploaded   time.Time `json:"date_uploaded"`   // the date and time that the clip was uploaded to the cloud storage provider
}

// Parse takes a file and returns a clip
func Parse(f os.FileInfo) Clip {
	return Clip{
		Name: f.Name(),
	}
}

func getType(f os.FileInfo) string {
	return ""
}

func getCamera(f os.FileInfo) string {
	if regexp.MustCompile(`right_repeater`).MatchString(f.Name()) {
		return "right"
	}

	if regexp.MustCompile(`left_repeater`).MatchString(f.Name()) {
		return "left"
	}

	if regexp.MustCompile(`front`).MatchString(f.Name()) {
		return "front"
	}

	return ""
}
