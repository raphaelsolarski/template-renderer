package renderer

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RasterizeSvgToPng(svg string) ([]byte, error) {
	cmd := exec.Command("rsvg-convert", "-f", "png")
	cmd.Stdin = bytes.NewReader([]byte(svg))

	var pngOutput bytes.Buffer
	cmd.Stdout = &pngOutput

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("conversion ended with error: %s", err)
	}

	copied := make([]byte, pngOutput.Len())
	copy(copied, pngOutput.Bytes())
	return copied, nil
}
