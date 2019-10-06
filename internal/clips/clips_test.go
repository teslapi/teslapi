package clips

import (
	"testing"
)

func TestImport(t *testing.T) {
	// Arrange
	filepath := "../../testdata/TeslaUSB"

	// Act
	err := Import(filepath)

	// Assert
	if err != nil {
		t.Error(err)
	}
}
