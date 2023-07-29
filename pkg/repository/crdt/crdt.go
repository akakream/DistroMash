package crdt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/utils"
)

func PostCrdtKeyValue(keyValuePair []byte) error {
	url := fmt.Sprintf("http://%s/crdt", utils.Libp2pURL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(keyValuePair))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return fmt.Errorf(
				"Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s",
				resp.StatusCode,
				err.Error(),
			)
		}
		return fmt.Errorf(
			"Non-OK HTTP status from the api with status code %d: %s",
			resp.StatusCode,
			apiErr,
		)
	}

	var data models.Crdt
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	return nil
}

func GetCrdtValue(key string) (*models.Crdt, error) {
	url := fmt.Sprintf("http://%s/crdt/%s", utils.Libp2pURL, key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf(
				"Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s",
				resp.StatusCode,
				err.Error(),
			)
		}
		return nil, fmt.Errorf(
			"Non-OK HTTP status from the api with status code %d: %s",
			resp.StatusCode,
			apiErr,
		)
	}

	var data models.Crdt
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func GetCrdtList() ([]models.Crdt, error) {
	url := fmt.Sprintf("http://%s/crdt", utils.Libp2pURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf(
				"Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s",
				resp.StatusCode,
				err.Error(),
			)
		}
		return nil, fmt.Errorf(
			"Non-OK HTTP status from the api with status code %d: %s",
			resp.StatusCode,
			apiErr,
		)
	}

	var data []models.Crdt
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DeleteCrdtKeyValue(key string) error {
	url := fmt.Sprintf("http://%s/crdt/%s", utils.Libp2pURL, key)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		apiErr, err := utils.GetErrorFromResponse(resp)
		if err != nil {
			return fmt.Errorf(
				"Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s",
				resp.StatusCode,
				err.Error(),
			)
		}
		return fmt.Errorf(
			"Non-OK HTTP status from the api with status code %d: %s",
			resp.StatusCode,
			apiErr,
		)
	}

	return nil
}
