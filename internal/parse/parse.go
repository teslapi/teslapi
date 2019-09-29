package parse

import "regexp"

// CameraFromRecording will take the name of the file and identify the camera
// that recorded the video
func CameraFromRecording(filename string) string {
	if regexp.MustCompile(`right_repeater`).MatchString(filename) {
		return "right"
	}

	if regexp.MustCompile(`left_repeater`).MatchString(filename) {
		return "left"
	}

	if regexp.MustCompile(`front`).MatchString(filename) {
		return "front"
	}

	return ""
}
