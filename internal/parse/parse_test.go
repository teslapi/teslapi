package parse

import (
	"testing"
)

func TestCameraFromRecording(t *testing.T) {
	testCases := []struct {
		desc     string
		filename string
		expected string
	}{
		{
			desc:     "can parse the right repeater camera",
			filename: "2019-03-08_07-09-right_repeater.mp4",
			expected: "right",
		},
		{
			desc:     "can parse the left repeater camera",
			filename: "2019-03-08_07-09-left_repeater.mp4",
			expected: "left",
		},
		{
			desc:     "can parse the front camera",
			filename: "2019-03-08_07-09-front.mp4",
			expected: "front",
		},
		{
			desc:     "if it does not match at all, it returns empty",
			filename: "2019-03-08_07-09-somethingelse.mp4",
			expected: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			camera := CameraFromRecording(tC.filename)

			if camera != tC.expected {
				t.Errorf("expected camera to be %v, got %v instead", tC.expected, camera)
			}
		})
	}
}
