package mock

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CreateHttpRequest(method string, url string, body interface{}) (*http.Request, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	stream := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(method, url, stream)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
