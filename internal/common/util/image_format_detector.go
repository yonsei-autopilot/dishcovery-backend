package util

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func DetectImageFormat(file io.Reader) (string, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return "", fmt.Errorf("failed to read image data: %w", err)
	}

	_, format, err := image.DecodeConfig(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", fmt.Errorf("failed to decode image format: %w", err)
	}

	return format, nil
}
