package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/utils"
)

func UploadImage2IPFS(imageName []byte) (*models.ImageWithCID, error) {
	url := fmt.Sprintf("http://%s/image", utils.Multiplatform2ipfsURL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(imageName))
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
			return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: Error when reading erorr message: %s", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Non-OK HTTP status from the api with status code %d: %s", resp.StatusCode, apiErr)
	}

	var data models.ImageWithCID
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
