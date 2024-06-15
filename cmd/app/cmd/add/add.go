package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [task] [date] [time] [message]",
	Short: "Добавить время в задачу",
	Long: `
	Добавляет время в задачу с номером или по шаблону.
	Задача(task) описывается как часть URL строки, например VUZ-01.
	Дата(date) описывается в формате d.m.Y, если год опускается берётся текущий.
	Время(time) берётся в формате для YT, например 12h1m.
	Сообщение(message) комментарий к времени
	`,
	RunE: addTime,
}

type Duration struct {
	Minutes int `json:"minutes"`
}
type TypeDuration struct {
	Id string `json:"id"`
}
type WorkItemCreate struct {
	Text     string       `json:"text"`
	Duration Duration     `json:"duration"`
	Type     TypeDuration `json:"type"`
}

func addTime(cmd *cobra.Command, args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("Argument must be %d", 4)
	}

	taskNumber := args[0]
	date := args[1]
	t := args[2]
	message := args[3]
	fmt.Println(taskNumber, date, t, message)

	timeValue, err := strconv.Atoi(t)

	if err != nil {
		log.Fatalln(err)
	}

	create := &WorkItemCreate{Text: message, Duration: Duration{Minutes: timeValue}, Type: TypeDuration{Id: "116-16"}}

	config := config.GetConfig()

	var bearer = "Bearer " + config.Token

	client := &http.Client{
		Timeout: time.Minute,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(create)

	req, err := http.NewRequest("POST", "https://"+config.ApiUrl+"/api/issues/"+taskNumber+"/timeTracking/workItems", payloadBuf)

	fmt.Println("https://"+config.ApiUrl+"/api/issues/"+taskNumber+"/timeTracking/workItems", payloadBuf)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {bearer},
	}

	res, err := client.Do(req)
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
	fmt.Println(body)

	return nil
}
