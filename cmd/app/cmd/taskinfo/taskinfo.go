package taskInfo

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/chex0v/yt-time-tracker/internal/tracker/issue"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	view "github.com/chex0v/yt-time-tracker/internal/views/issue"
	views "github.com/chex0v/yt-time-tracker/internal/views/workitem"
	"github.com/spf13/cobra"
)

var TaskInfoCmd = &cobra.Command{
	Use:   "time [task]",
	Short: "Узнать информацию о зафиксированном времени в задаче",
	Long: `
	Получаем информацию о зафиксированном времени в задаче
	`,
	RunE: taskInfo,
}

func taskInfo(_ *cobra.Command, args []string) error {
	var err error
	var task issue.Issue
	var workItems workitem.WorkItems

	if len(args) < 1 {
		return fmt.Errorf("Argument must be %d", 1)
	}

	c := config.GetConfig()
	taskNumber := c.TaskNumber(args[0])

	task, err = progressbar.Progress(func() (issue.Issue, error) {
		yt := tracker.NewTracker()
		return yt.TaskInfo(taskNumber)
	})
	view.Issue(task)

	if err != nil {
		return err
	}

	workItems, err = progressbar.Progress(func() (workitem.WorkItems, error) {
		yt := tracker.NewTracker()
		return yt.TaskTackerInfo(taskNumber)
	})
	if err != nil {
		return err
	}

	views.WorkItems(workItems)

	return nil
}
