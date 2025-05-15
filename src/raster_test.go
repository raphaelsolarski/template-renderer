package renderer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRasterizeSvgToPng(t *testing.T) {
	input := string(ReadFile(t, "testdata/circle.svg"))
	expected := ReadFile(t, "testdata/circle.png")

	output, err := RasterizeSvgToPng(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}

func TestTryRasterizeInvalidSvg(t *testing.T) {
	input := `<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg">`
	output, err := RasterizeSvgToPng(input)
	assert.Nil(t, output)
	assert.Equal(t, fmt.Errorf("conversion ended with error: exit status 1"), err)
}

func ReadFile(t *testing.T, filePath string) []byte {
	expected, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	return expected
}
