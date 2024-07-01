package report

import (
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	views "github.com/chex0v/yt-time-tracker/internal/views/workitem"
	"github.com/spf13/cobra"
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

	items, err := progressbar.Progress(func() (workitem.WorkItems, error) {
		return yt.MyWorkItemByDate(date)
	})
	if err != nil {
		return err
	}

	views.WorkItems(items)

	return nil
}
