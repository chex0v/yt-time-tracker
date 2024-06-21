package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	TaskInfoUrl = "/issues/%s?fields=id,project(id,name)"
)

type TaskInfo struct {
	ID      string  `json:"id"`
	Project Project `json:"project"`
	Name    string  `json:"name"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c Client) TaskInfo(taskNumber string) (TaskInfo, error) {

	req, err := http.NewRequest("GET", c.Url+fmt.Sprintf(TaskInfoUrl, taskNumber), nil)

	if err != nil {
		return TaskInfo{}, err
	}

	req.Header.Add("Authorization", c.headerToken())

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
		return TaskInfo{}, fmt.Errorf("status code for task info error: %d %s", res.StatusCode, res.Status)
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

	return target, nil

}
