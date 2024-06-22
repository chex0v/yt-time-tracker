package tracker

import (
	"encoding/json"
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	TaskTrackerUrlInfo = "/timeTracking?fields=draftWorkItem(id),enabled,workItems(created,duration(presentation,minutes),author(name),creator(name),date,id,text)"
)

type WorkItems struct {
	Items []WorkItem `json:"workItems"`
}

type Duration struct {
	Minutes      int    `json:"minutes,omitempty"`
	Presentation string `json:"presentation,omitempty"`
}

type WorkItem struct {
	Duration    Duration `json:"duration"`
	Date        int64    `json:"date"`
	Created     int64    `json:"created"`
	Creator     User     `json:"creator"`
	Author      User     `json:"author"`
	Id          string   `json:"id"`
	TextPreview string   `json:"textPreview,omitempty"`
	Text        string   `json:"text,omitempty"`
	Issue       Issue    `json:"issue"`
}

type Issue struct {
	ID         string `json:"id"`
	Summary    string `json:"summary"`
	IdReadable string `json:"idReadable"`
}

func (i Issue) Link() string {
	c := config.GetConfig()
	u, err := url.Parse(c.ApiUrl)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s://%s/issue/%s/", u.Scheme, u.Host, i.IdReadable)
}

func (c Client) TaskTackerInfo(taskNumber string) (WorkItems, error) {

	req, err := http.NewRequest("GET", c.Url+"/issues/"+taskNumber+TaskTrackerUrlInfo, nil)

	if err != nil {
		return WorkItems{}, err
	}

	req.Header.Add("Authorization", c.headerToken())

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return WorkItems{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return WorkItems{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WorkItems{}, err
	}

	target := WorkItems{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return WorkItems{}, err
	}

	return target, nil

}
