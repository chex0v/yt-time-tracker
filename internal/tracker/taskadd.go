package tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
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
	Date     int64        `json:"date,omitempty"`
}

func (c Client) WorkItemAdd(taskNumber string, taskAdd WorkItemCreate) (WorkItem, error) {

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(taskAdd)
	if err != nil {
		return WorkItem{}, err
	}

	req, err := http.NewRequest("POST", c.Url+fmt.Sprintf(AddTracker, taskNumber), payloadBuf)

	if err != nil {
		return WorkItem{}, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {c.headerToken()},
	}
	s := progressbar.NewProgressBar()
	s.Start()
	res, err := c.HTTPClient.Do(req)
	s.Stop()

	if err != nil {
		return WorkItem{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return WorkItem{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WorkItem{}, err
	}

	target := WorkItem{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return WorkItem{}, err
	}

	return target, nil
}
