package taskInfo

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"log"
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
	s.Stop()
	if err != nil {
		log.Fatal(err)
	}

	t := tabby.New()
	t.AddHeader("Id", "Author", "Date", "Duration")

	for _, item := range items.Items {
		tm := time.Unix(item.Date, 0)
		t.AddLine(item.Id, item.Author.Name, tm.Format(time.ANSIC), item.Duration.Presentation)
	}
	t.Print()

	return nil
}
