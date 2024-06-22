package add

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
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

func viewTaskTypes(types []tracker.WorkItemType) (tracker.WorkItemType, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | red }}",
		Active:   "\U0001F32D {{ .Id }} ({{ .Name }})",
		Inactive: "  {{ .Id }} ({{ .Name }})",
		Selected: "\U0001F336 {{ .Id | red | cyan }} {{ .Name }}",
		Details: `
--------- Тип ----------
{{ "Name:" | faint }}	{{ .Name }}`,
	}

	prompt := promptui.Select{
		Label:     "Нужно выбрать тип времени",
		Size:      len(types),
		Items:     types,
		Templates: templates,
	}

	indexType, _, err := prompt.Run()

	if err != nil {
		return tracker.WorkItemType{}, err
	}

	return types[indexType], nil
}

func addTime(cmd *cobra.Command, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("Argument must be %d", 3)
	}

	taskNumber := args[0]
	t := args[1]
	message := args[2]

	timeValue := t

	config := config.GetConfig()

	clientTracker := tracker.NewClient(config.ApiUrl, config.Token)

	types, err := clientTracker.TaskTypesByTask(taskNumber)

	if err != nil {
		log.Fatal(err)
	}

	typeTask, err := viewTaskTypes(types)

	if err != nil {
		log.Fatal(err)
	}

	create := tracker.WorkItemCreate{Text: message, Duration: tracker.Duration{Presentation: timeValue}, Type: tracker.TypeDuration{Id: typeTask.Id}}

	wItem := tracker.WorkItem{}
	wItem, err = clientTracker.WorkItemAdd(taskNumber, create)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Время добавлено. id: %s", wItem.Id)
	return nil
}
