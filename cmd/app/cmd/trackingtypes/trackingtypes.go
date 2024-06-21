package trackingtypes

import (
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"log"
)

var TrackingTypesCmd = &cobra.Command{
	Use:   "types",
	Short: "Список типов",
	Long: `
	Получаем информацию о типах при добавлении времени
	`,
	RunE: trackingTypes,
}

func trackingTypes(cmd *cobra.Command, args []string) error {

	config := config.GetConfig()

	client := tracker.NewClient(config.ApiUrl, config.Token)
	task, err := cmd.Flags().GetString("task")

	if err != nil {
		log.Fatal(err)
	}
	s := progressbar.NewProgressBar()
	s.Start()
	var types []tracker.WorkItemType

	if task != "" {
		types, err = client.TaskTypesByTask(task)
	} else {
		types, err = client.TaskType()
	}

	s.Stop()

	if err != nil {
		log.Fatal(err)
	}

	t := tabby.New()

	t.AddHeader("Id", "Name")

	for _, workType := range types {
		t.AddLine(workType.Id, workType.Name)
	}
	t.Print()

	return nil
}
