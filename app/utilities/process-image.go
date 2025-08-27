package utilities

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

const (
	uploadRoot = "./uploads"
)

var allowedMIMEs = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/gif":  true,
}

func ProcessImage(file multipart.File) ([]byte, error) {
	buf, err := io.ReadAll(file)

	if err != nil {
		return []byte{}, errors.New("failed to read file")
	}

	mime := bimg.DetermineImageTypeName(buf)
	if mime == "" {
		return []byte{}, errors.New("unsupported or unrecognized image")
	}
	normalized := "image/" + strings.ToLower(mime)

	if !allowedMIMEs[normalized] {
		return []byte{}, errors.New("unsupported image type")
	}

	opts := bimg.Options{
		Width:         1080,
		Height:        1080,
		Crop:          true,
		Gravity:       bimg.GravityCentre,
		Enlarge:       false,
		StripMetadata: true,
		NoProfile:     true,
		Type:          bimg.WEBP,
		Quality:       90,
		Lossless:      false,
		Compression:   6,
	}

	img := bimg.NewImage(buf)

	processedImg, err := img.Process(opts)

	if err != nil {
		return []byte{}, errors.New("image processing failed")
	}

	return processedImg, nil
}

func GenerateImageFilename(postId, imageIndex int, attachmentType string) string {
	var filename string

	switch attachmentType {
	case "image":
		filename = fmt.Sprintf("%d_%s.webp", postId, attachmentType)
	case "carousel":
		filename = fmt.Sprintf("%d_%s_%d.webp", postId, attachmentType, imageIndex)
	}

	return filename
}

func SaveImage(filename string, file []byte) error {
	fullpath := filepath.Join(uploadRoot, filename)

	err := os.WriteFile(fullpath, file, 0o644)

	if err != nil {
		return errors.New("failed to save file")
	}

	return nil
}
