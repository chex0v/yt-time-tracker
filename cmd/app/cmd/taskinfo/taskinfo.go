package taskInfo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/spf13/cobra"
)

type WorkItems struct {
	Items []WorkItem `json:"workItems"`
}

type Duration struct {
	Minutes      int    `json:"minutes"`
	Presentation string `json:"presentation"`
}

type WorkItem struct {
	Duration Duration `json:"duration"`
	Date     int64    `json:"date"`
	Created  int64    `json:"created"`
	Creator  User     `json:"creator"`
	Author   User     `json:"author"`
	Id       string   `json:"id"`
}

type User struct {
	Name string `json:"name"`
}

var TaskInfoCmd = &cobra.Command{
	Use:   "info [task]",
	Short: "Узнать информацию о задаче",
	Long: `
	Получаем информацию о задаче
	`,
	RunE: taskInfo,
}

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(url, apiKey string) *Client {
	return &Client{
		BaseURL: url,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func taskInfo(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Argument must be %d", 1)
	}

	config := config.GetConfig()

	taskNumber := args[0]

	var bearer = "Bearer " + config.Token

	client := NewClient("https://"+config.ApiUrl+"/api", config.Token)

	req, err := http.NewRequest("GET", client.BaseURL+"/issues/"+taskNumber+"/timeTracking?fields=draftWorkItem(id),enabled,workItems(created,duration(presentation,minutes),author(name),creator(name),date,id)", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", bearer)

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	target := WorkItems{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range target.Items {

		fmt.Printf("%s: %s %s", item.Id, item.Author.Name, item.Duration.Presentation)
		fmt.Println("")
	}

	return nil
}
