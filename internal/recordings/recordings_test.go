package recordings

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	// Arrange
	path := "testdata/TeslaUSB/RecentClips/"
	testFile := "2019-03-08_07-07-front.mp4"
	file, err := os.Open(path + testFile)
	if err != nil {
		t.Fatal(err)
	}
	info, fileErr := file.Stat()
	if fileErr != nil {
		t.Fatal(fileErr)
	}

	// Act
	c := Parse(info)

	// Assert
	if c.Name != "2019-03-08_07-07-front.mp4" {
		t.Errorf("expected the name to be %v, got %v instead", testFile, c.Name)
	}
	if c.Type != "recent" {
		t.Errorf("expected the type to be %v, got %v instead", "recent", c.Type)
	}
}
