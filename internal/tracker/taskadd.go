package tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	AddTracker = "/issues/%s/timeTracking/workItems"
)

type TypeDuration struct {
	Id string `json:"id"`
}
type WorkItemCreate struct {
	Text     string       `json:"text"`
	Duration Duration     `json:"duration"`
	Type     TypeDuration `json:"type"`
}

func (c Client) TaskAdd(taskNumber string, taskAdd WorkItemCreate) (TaskInfo, error) {

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(taskAdd)
	if err != nil {
		return TaskInfo{}, err
	}

	req, err := http.NewRequest("POST", c.Url+fmt.Sprintf(AddTracker, taskNumber), payloadBuf)

	if err != nil {
		return TaskInfo{}, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {c.headerToken()},
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return TaskInfo{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return TaskInfo{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return TaskInfo{}, err
	}

	target := TaskInfo{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return TaskInfo{}, err
	}

	return TaskInfo{}, nil
}
