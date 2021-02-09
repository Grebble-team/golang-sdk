package helpers

import (
	"encoding/json"
	"net/http"
)

func GetFileContentType(data []byte) (string, error) {
	contentType := http.DetectContentType(data)

	return contentType, nil
}

func MapAttributesContentType(attr map[string]string, o interface{}) error {
	result, err := json.Marshal(attr)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &o)
	if err != nil {
		return err
	}

	return nil
}
