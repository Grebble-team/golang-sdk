package helpers

import (
	"net/http"
)

func GetFileContentType(data []byte) (string, error) {
	contentType := http.DetectContentType(data)

	return contentType, nil
}
