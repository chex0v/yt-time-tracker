package cmd

import (
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/add"
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/report"
	taskInfo "github.com/chex0v/yt-time-tracker/cmd/app/cmd/taskinfo"
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/user"
	"github.com/spf13/cobra"
	"os"
)

var (
	task     string
	date     string
	dateDo   string
	typeTask string
)
var rootCmd = &cobra.Command{
	Use:   "ytt",
	Short: "Для работы с yt",
	Long:  `Инструмент для работы с верменем на yt`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addCommands() {
	reportCmd := report.MyReportByTodayCmd
	addCmd := add.AddCmd
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(taskInfo.TaskInfoCmd)
	rootCmd.AddCommand(user.MyUserInfoCmd)
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().StringVarP(&date, "date", "d", "", "Date of report. Format YYYY-MM-DD")
	reportCmd.Flags().StringVarP(&dateDo, "date-to", "t", "", "Date end of report. Format YYYY-MM-DD")
	addCmd.Flags().StringVarP(&task, "date", "d", "", "Date for add. Format YYYY-MM-DD")
	addCmd.Flags().StringVarP(&typeTask, "type", "t", "", "Type of task. Id of type task")
}
func init() {
	addCommands()
}
