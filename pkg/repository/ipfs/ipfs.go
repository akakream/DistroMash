package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/akakream/DistroMash/models"
	"github.com/akakream/DistroMash/pkg/repository/crdt"
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

	var data models.ImageWithCID
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func PinCid(cid string) error {
	url := fmt.Sprintf("http://%s/pin/%s", utils.Multiplatform2ipfsURL, cid)

	req, err := http.NewRequest("POST", url, nil)
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
		return fmt.Errorf(
			"Non-OK HTTP status from the api with status code %d",
			resp.StatusCode,
		)
	}

	return nil
}

type JobResult struct {
	Data  models.ImageWithCID
	Error error
}

func LogPostResult(postResultChan <-chan JobResult) {
	result := <-postResultChan
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println(result.Data)
	}
}

func AsyncPostImage(postResultChan chan<- JobResult, imageNameTag []byte) {
	var result JobResult

	cidNameTagPair, err := UploadImage2IPFS(imageNameTag)
	if err != nil {
		result.Error = err
	}

	// Add the cid to the CRDT key value store
	nameTag := cidNameTagPair.Name + ":" + cidNameTagPair.Tag
	crdtPayload, err := json.Marshal(models.Crdt{Key: nameTag, Value: cidNameTagPair.Cid})
	if err != nil {
		result.Error = err
	}
	err = crdt.PostCrdtKeyValue(crdtPayload)
	if err != nil {
		result.Error = err
	}

	// Add the reverse lookup
	reverseCrdtPayload, err := json.Marshal(
		models.Crdt{Key: cidNameTagPair.Cid, Value: nameTag},
	)
	if err != nil {
		result.Error = err
	}
	err = crdt.PostCrdtKeyValue(reverseCrdtPayload)
	if err != nil {
		result.Error = err
	}

	if result.Error == nil {
		result.Data = models.ImageWithCID{
			Name: cidNameTagPair.Name,
			Tag:  cidNameTagPair.Tag,
			Cid:  cidNameTagPair.Cid,
		}
	}

	postResultChan <- result
}
