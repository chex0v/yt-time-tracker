package add

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	"github.com/chex0v/yt-time-tracker/internal/util"
	view "github.com/chex0v/yt-time-tracker/internal/views/issue"
	views "github.com/chex0v/yt-time-tracker/internal/views/workitem"
	"github.com/spf13/cobra"
	"time"
)

var AddCmd = &cobra.Command{
	Use:   "add [task] [time] [message]",
	Short: "Добавить время в задачу",
	Long: `
	Добавляет время в задачу с номером или по шаблону.
	Задача(task) описывается как часть URL строки, например VUZ-01.
	Время(time) берётся в формате для YT, например 12h1m.
	Сообщение(message) комментарий к времени
	`,
	RunE: addTime,
}

func addTime(cmd *cobra.Command, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("Argument must be %d", 3)
	}

	var err error
	var wItem workitem.WorkItem
	var typeTask workitem.Type
	var types []workitem.Type
	var strTypeTask string

	date, _ := cmd.Flags().GetString("date")
	strTypeTask, err = cmd.Flags().GetString("type")

	c := config.GetConfig()

	taskNumber := c.TaskNumber(args[0])
	strTypeTask = c.TypeId(strTypeTask)

	if date != "" {
		date = util.GetDynamicDayByMacros(date)
	}

	t := args[1]
	message := args[2]
	timeValue := t

	if len(strTypeTask) <= 0 {
		types, err = progressbar.Progress(func() ([]workitem.Type, error) {
			return tracker.NewTracker().TaskTypesByTask(taskNumber)
		})

		if err != nil {
			return err
		}

		typeTask, err = view.ChoiceType(types)

		if err != nil {
			return err
		}
	} else {
		typeTask = workitem.Type{Id: strTypeTask}
	}

	wItem, err = progressbar.Progress(func() (workitem.WorkItem, error) {
		create := workitem.Create{Text: message, Duration: workitem.Duration{Presentation: timeValue}, Type: workitem.Type{Id: typeTask.Id}}

		if date != "" {
			timeParse, err := time.Parse(time.DateOnly, date)
			if err != nil {
				return workitem.WorkItem{}, err
			}
			create.Date = timeParse.UnixMilli()
		}

		return tracker.NewTracker().WorkItemAdd(taskNumber, create)
	})

	if err != nil {
		return err
	}
	fmt.Println("Время добавлено")
	views.WorkItem(wItem)
	return nil
}
