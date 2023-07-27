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

type JobResult struct {
    Data models.ImageWithCID
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

func AsyncPostImage(postResultChan chan<- JobResult, imageTag []byte) {
    var result JobResult

	cidTagPair, err := UploadImage2IPFS(imageTag)
	if err != nil {
        result.Error = err
	}
    
    // Add the cid to the CRDT key value store
    crdtPayload, err := json.Marshal(models.Crdt{Key: cidTagPair.Name, Value: cidTagPair.Cid})
    if err != nil {
        result.Error = err
    }
    err = crdt.PostCrdtKeyValue(crdtPayload)
    if err != nil {
        result.Error = err
    }

    // Add the reverse lookup
    reverseCrdtPayload, err := json.Marshal(models.Crdt{Key: cidTagPair.Cid, Value: cidTagPair.Name})
    if err != nil {
        result.Error = err
    }
    err = crdt.PostCrdtKeyValue(reverseCrdtPayload)
    if err != nil {
        result.Error = err
    }

    if result.Error == nil {
        result.Data = models.ImageWithCID{
            Name: cidTagPair.Name,
            Cid: cidTagPair.Cid,
        }
    }

    postResultChan <- result
}
