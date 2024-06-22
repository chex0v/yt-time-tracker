package report

import (
	"fmt"
	taskInfo "github.com/chex0v/yt-time-tracker/cmd/app/cmd/taskinfo"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"log"
	"sort"
	"strconv"
	"time"
)

var MyReportByTodayCmd = &cobra.Command{
	Use:   "report",
	Short: "Информация о учёте времени за сегодня",
	Long: `
	Получение времени учтённого за сегодня
	`,
	RunE: reportByToday,
}

func reportByToday(cmd *cobra.Command, args []string) error {
	var err error
	var date string

	date, _ = cmd.Flags().GetString("date")

	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}

	config := config.GetConfig()

	client := tracker.NewClient(config.ApiUrl, config.Token)

	s := progressbar.NewProgressBar()

	s.Start()
	workItems, err := client.MyWorkItemByDate(date)
	s.Stop()

	if err != nil {
		log.Fatal(err)
	}

	t := tabby.New()

	groupByData := taskInfo.GroupByProperty(workItems, func(i tracker.WorkItem) int64 {
		return i.Date
	})

	keys := make([]int64, 0, len(groupByData))
	for k := range groupByData {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	totalM := 0
	for _, d := range keys {
		tm := time.Unix(d/1000, 0)
		t.AddLine("Дата: ", tm.Format(time.DateOnly), "")
		for _, item := range groupByData[d] {
			t.AddLine(item.Id, item.Author.Name, item.Duration.Presentation, item.Text, item.Issue.Link())
			totalM += item.Duration.Minutes
		}
	}
	d, err := time.ParseDuration(strconv.Itoa(totalM) + "m")

	if err != nil {
		log.Fatal(err)
	}
	t.AddLine("-------------")
	t.AddLine("Итого: ", fmt.Sprintf("%.2f часов", d.Hours()))
	s.Stop()
	t.Print()

	return nil
}
