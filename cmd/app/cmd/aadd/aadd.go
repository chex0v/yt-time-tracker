package aadd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AaddCmd = &cobra.Command{
	Use:   "aadd [time] [message]",
	Short: "Добавить время в задачи",
	Long: `
	Добавляет время в задачу с номером или по шаблону.
	Время(time) берётся в формате для YT, например 12h1m.
	Сообщение(message) комментарий к времени
	`,
	RunE: addTimes,
}

func addTimes(cmd *cobra.Command, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("Argument must be %d", 3)
	}

	//var err error
	//var wItem workitem.WorkItem
	//var typeTask workitem.Type
	//var types []workitem.Type
	//var strTypeTask string
	//var taskNumbers []string
	//var date string
	//
	//date, err = cmd.Flags().GetString("date")
	//strTypeTask, err = cmd.Flags().GetString("type")
	//taskNumbers, err = cmd.Flags().GetStringSlice("tasks")
	//
	//c := config.GetConfig()
	//
	//for _, taskNumber := range taskNumbers {
	//	taskNumber = c.TaskNumber(taskNumber)
	//}
	//
	////taskNumber := c.TaskNumber(args[0])
	//strTypeTask = c.TypeId(strTypeTask)

	//t := args[0]
	//message := args[1]
	//timeValue := t

	//if len(strTypeTask) <= 0 {
	//	types, err = progressbar.Progress(func() ([]workitem.Type, error) {
	//		return tracker.NewTracker().TaskTypesByTask(taskNumber)
	//	})
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	typeTask, err = view.ChoiceType(types)
	//
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//	typeTask = workitem.Type{Id: strTypeTask}
	//}
	//
	//wItem, err = progressbar.Progress(func() (workitem.WorkItem, error) {
	//	create := workitem.Create{Text: message, Duration: workitem.Duration{Presentation: timeValue}, Type: workitem.Type{Id: typeTask.Id}}
	//
	//	if date != "" {
	//		timeParse, err := time.Parse(time.DateOnly, date)
	//		if err != nil {
	//			return workitem.WorkItem{}, err
	//		}
	//		create.Date = timeParse.UnixMilli()
	//	}
	//
	//	return tracker.NewTracker().WorkItemAdd(taskNumber, create)
	//})

	//if err != nil {
	//	return err
	//}
	//fmt.Println("Время добавлено")
	//views.WorkItem(wItem)
	//return nil
}
