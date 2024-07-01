package tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/client"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker/issue"
	"github.com/chex0v/yt-time-tracker/internal/tracker/user"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	"io"
	"log"
	"net/http"
)

const (
	UserMePath         = "/users/me?fields=id,name,login,fullName,email,online"
	WorkItemsPath      = "/workItems"
	AddTracker         = "/issues/%s/timeTracking/workItems"
	TaskInfoUrl        = "/issues/%s?fields=id,project(id,name)"
	TaskType           = "/admin/timeTrackingSettings/workItemTypes?fields=id,name,url,autoAttached"
	TaskTypeByTask     = "/admin/projects/%s/timeTrackingSettings/workItemTypes?fields=id,name,url,autoAttached"
	TaskTrackerUrlInfo = "/timeTracking?fields=draftWorkItem(id),enabled,workItems(created,duration(presentation,minutes),author(name),creator(name),date,id,text),idReadable"
)

type Tracker struct {
	Client *client.Client
}

func NewTracker() Tracker {
	appConfig := config.GetConfig()
	c := client.NewClient(appConfig.ApiUrl, appConfig.Token)

	return Tracker{Client: c}
}

func (t Tracker) MyUserInfo() (user.User, error) {

	res, err := t.Client.Do("GET", UserMePath, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return user.User{}, fmt.Errorf("status code for user info error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return user.User{}, err
	}

	target := user.User{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return user.User{}, err
	}

	return target, nil

}
func (t Tracker) MyWorkItemByDate(date, dateTo string) (workitem.WorkItems, error) {

	u, err := t.MyUserInfo()

	if err != nil {
		log.Fatal(err)
	}

	url := WorkItemsPath + fmt.Sprintf("?startDate=%s&endDate=%s&author=%s&fields=id,author(id,name),creator(id,name),created,date,duration(minutes,presentation),text,issue(id,summary,idReadable)", date, dateTo, u.Id)

	res, err := t.Client.Do("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return workitem.WorkItems{}, err
	}

	if res.StatusCode != http.StatusOK {
		return workitem.WorkItems{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var target []workitem.WorkItem

	err = json.Unmarshal(body, &target)

	if err != nil {
		return workitem.WorkItems{}, err
	}
	var workItems workitem.WorkItems

	for _, v := range target {
		workItems.Items = append(workItems.Items, v)
	}

	return workItems, nil
}
func (t Tracker) WorkItemAdd(taskNumber string, taskAdd workitem.Create) (workitem.WorkItem, error) {

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(taskAdd)
	if err != nil {
		return workitem.WorkItem{}, err
	}

	s := progressbar.NewProgressBar()
	s.Start()
	res, err := t.Client.Do("POST", fmt.Sprintf(AddTracker, taskNumber), payloadBuf)
	s.Stop()

	if err != nil {
		return workitem.WorkItem{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return workitem.WorkItem{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return workitem.WorkItem{}, err
	}

	target := workitem.WorkItem{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return workitem.WorkItem{}, err
	}

	return target, nil
}
func (t Tracker) TaskInfo(taskNumber string) (issue.Issue, error) {

	res, err := t.Client.Do("GET", fmt.Sprintf(TaskInfoUrl, taskNumber), nil)
	if err != nil {
		return issue.Issue{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return issue.Issue{}, fmt.Errorf("status code for task info error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return issue.Issue{}, err
	}

	target := issue.Issue{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return issue.Issue{}, err
	}

	return target, nil

}
func (t Tracker) TaskType() ([]workitem.Type, error) {
	res, err := t.Client.Do("GET", TaskType, nil)

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

	var target []workitem.Type

	err = json.Unmarshal(body, &target)

	if err != nil {
		return nil, err
	}

	return target, nil
}
func (t Tracker) TaskTypesByTask(taskNumber string) ([]workitem.Type, error) {
	tInfo, err := t.TaskInfo(taskNumber)
	if err != nil {
		return nil, err
	}

	pId := tInfo.Project.ID
	if pId == "" {
		return nil, fmt.Errorf("project not found for task %s", taskNumber)
	}

	res, err := t.Client.Do("GET", fmt.Sprintf(TaskTypeByTask, pId), nil)

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

	var target []workitem.Type
	err = json.Unmarshal(body, &target)

	if err != nil {
		return nil, err
	}

	return target, nil
}
func (t Tracker) TaskTackerInfo(taskNumber string) (workitem.WorkItems, error) {
	res, err := t.Client.Do("GET", "/issues/"+taskNumber+TaskTrackerUrlInfo, nil)
	if err != nil {
		return workitem.WorkItems{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return workitem.WorkItems{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return workitem.WorkItems{}, err
	}

	target := workitem.WorkItems{}

	err = json.Unmarshal(body, &target)
	for i := range target.Items {
		target.Items[i].Issue.IdReadable = taskNumber
	}

	if err != nil {
		return workitem.WorkItems{}, err
	}

	return target, nil

}
