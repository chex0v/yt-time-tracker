package taskInfo

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	"github.com/chex0v/yt-time-tracker/internal/util"
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

func taskInfo(_ *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Argument must be %d", 1)
	}

	taskNumber := args[0]

	yt := tracker.NewTracker()

	s := progressbar.NewProgressBar()
	s.Start()
	info, err := yt.TaskInfo(taskNumber)
	tInfo := tabby.New()
	tInfo.AddLine(fmt.Sprintf("Id: %s, Project: %s", info.ID, info.Project.Name))
	tInfo.AddLine()
	tInfo.Print()
	if err != nil {
		log.Fatal(err)
	}
	items, err := yt.TaskTackerInfo(taskNumber)
	if err != nil {
		log.Fatal(err)
	}

	s = progressbar.NewProgressBar()

	t := tabby.New()
	t.AddHeader("", "")
	groupByData := util.GroupByProperty(items.Items, func(i workitem.WorkItem) int64 {
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
			t.AddLine(item.Id, item.Author.Name, item.Duration.Presentation, item.Text)
		}
	}
	s.Stop()
	t.Print()

	return nil
}
