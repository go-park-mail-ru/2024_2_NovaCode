package content

import (
	"fmt"
	"net/http"
)

var allowedImagesContentTypes = map[string]string{
	"image/jpg":     "jpg",
	"image/jpeg":    "jpeg",
	"image/png":     "png",
	"image/webp":    "webp",
	"image/svg+xml": "svg",
}

func IsImage(fileBytes []byte) (string, error) {
	contentType := http.DetectContentType(fileBytes)

	extension, ok := allowedImagesContentTypes[contentType]
	if !ok {
		return "", fmt.Errorf("content type is not allowed: %v", contentType)
	}

	return extension, nil
}
