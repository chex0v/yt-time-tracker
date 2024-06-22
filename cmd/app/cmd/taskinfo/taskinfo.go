package taskInfo

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"log"
	"sort"
	"time"
)

var TaskInfoCmd = &cobra.Command{
	Use:   "info [task]",
	Short: "Узнать информацию о задаче",
	Long: `
	Получаем информацию о задаче
	`,
	RunE: taskInfo,
}

func taskInfo(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Argument must be %d", 1)
	}

	config := config.GetConfig()

	taskNumber := args[0]

	client := tracker.NewClient(config.ApiUrl, config.Token)

	s := progressbar.NewProgressBar()
	s.Start()
	info, err := client.TaskInfo(taskNumber)
	tInfo := tabby.New()
	tInfo.AddLine(fmt.Sprintf("Id: %s, Project: %s", info.ID, info.Project.Name))
	tInfo.AddLine()
	tInfo.Print()
	if err != nil {
		log.Fatal(err)
	}
	items, err := client.TaskTackerInfo(taskNumber)
	if err != nil {
		log.Fatal(err)
	}

	s = progressbar.NewProgressBar()

	t := tabby.New()
	t.AddHeader("", "")
	groupByData := GroupByProperty(items.Items, func(i tracker.WorkItem) int64 {
		return i.Date
	})

	keys := make([]int64, 0, len(groupByData))
	for k := range groupByData {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, d := range keys {
		tm := time.Unix(d/1000, 0)
		t.AddLine("Дата: ", tm.Format(time.DateOnly), "")
		for _, item := range groupByData[d] {
			t.AddLine(item.Id, item.Author.Name, item.Duration.Presentation)
		}
	}
	s.Stop()
	t.Print()

	return nil
}

func GroupByProperty[T any, K comparable](items []T, getProperty func(T) K) map[K][]T {
	grouped := make(map[K][]T)

	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}
