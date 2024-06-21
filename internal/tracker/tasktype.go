package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WorkItemType struct {
	AutoAttached bool   `json:"autoAttached"`
	Name         string `json:"name"`
	Id           string `json:"id"`
}

const (
	TaskType       = "/admin/timeTrackingSettings/workItemTypes?fields=id,name,url,autoAttached"
	TaskTypeByTask = "/admin/projects/%s/timeTrackingSettings/workItemTypes?fields=id,name,url,autoAttached"
)

func (c Client) TaskType() ([]WorkItemType, error) {
	req, err := http.NewRequest("GET", c.Url+TaskType, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.headerToken())

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var target []WorkItemType

	err = json.Unmarshal(body, &target)

	if err != nil {
		return nil, err
	}

	return target, nil
}

func (c Client) TaskTypesByTask(taskNumber string) ([]WorkItemType, error) {
	tInfo, err := c.TaskInfo(taskNumber)
	if err != nil {
		return nil, err
	}

	pId := tInfo.Project.ID
	if pId == "" {
		return nil, fmt.Errorf("project not found for task %s", taskNumber)
	}

	req, err := http.NewRequest("GET", c.Url+fmt.Sprintf(TaskTypeByTask, pId), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.headerToken())

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var target []WorkItemType
	err = json.Unmarshal(body, &target)

	if err != nil {
		return nil, err
	}

	return target, nil
}
