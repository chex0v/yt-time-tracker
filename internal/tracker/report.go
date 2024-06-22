package tracker

import (
	"encoding/json"
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"io"
	"log"
	"net/http"
)

const (
	WorkItemsPath = "/workItems"
)

func (c Client) MyWorkItemByDate(date string) ([]WorkItem, error) {

	u, err := c.MyUserInfo()

	if err != nil {
		log.Fatal(err)
	}

	url := c.Url + WorkItemsPath + fmt.Sprintf("?startDate=%s&endDate=%s&author=%s&fields=id,author(id,name),creator(id,name),created,date,duration(minutes,presentation),text,issue(id,summary,idReadable)", date, date, u.Id)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", c.headerToken())

	if err != nil {
		log.Fatal(err)
	}

	p := progressbar.NewProgressBar()
	p.Start()
	res, err := c.HTTPClient.Do(req)
	p.Stop()

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []WorkItem{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []WorkItem{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var target []WorkItem

	err = json.Unmarshal(body, &target)

	if err != nil {
		return []WorkItem{}, err
	}

	return target, nil
}
