package trackingtypes

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

var TrackingTypesCmd = &cobra.Command{
	Use:   "types",
	Short: "Список типов",
	Long: `
	Получаем информацию о типах при добавлении времени
	`,
	RunE: trackingTypes,
}

type workItemType struct {
	AutoAttached bool   `json:"autoAttached"`
	Name         string `json:"name"`
	Id           string `json:"id"`
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

func trackingTypes(cmd *cobra.Command, args []string) error {

	config := config.GetConfig()

	var bearer = "Bearer " + config.Token

	client := NewClient("https://"+config.ApiUrl+"/api", config.Token)

	req, err := http.NewRequest("GET", client.BaseURL+"/admin/timeTrackingSettings/workItemTypes?fields=id,name,url,autoAttached", nil)

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

	target := []workItemType{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range target {
		fmt.Println(item)
	}

	return nil
}
