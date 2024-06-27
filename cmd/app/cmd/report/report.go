package report

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
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

func reportByToday(cmd *cobra.Command, _ []string) error {
	var date string
	date, _ = cmd.Flags().GetString("date")

	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}

	yt := tracker.NewTracker()

	_, err := progressbar.Progress(func() (workitem.WorkItems, error) {
		items, err := yt.MyWorkItemByDate(date)
		if err != nil {
			return workitem.WorkItems{}, err
		}
		viewWorkItems(items)
		return items, nil
	})

	return err
}

func viewWorkItems(workItems workitem.WorkItems) {
	t := tabby.New()

	groupByData := workItems.GroupByDate()

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
	t.Print()
}
