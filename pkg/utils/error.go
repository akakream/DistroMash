package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetErrorFromResponse(response *http.Response) (string, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var bodyJson apiNonOKResponse
	if err := json.Unmarshal(body, &bodyJson); err != nil {
		log.Println(err)
		return "", err
	}

	return bodyJson.Err, nil
}

type apiNonOKResponse struct {
	Err    string
	Status int
}
