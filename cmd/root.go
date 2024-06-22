package cmd

import (
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/add"
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/report"
	taskInfo "github.com/chex0v/yt-time-tracker/cmd/app/cmd/taskinfo"
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/trackingtypes"
	"github.com/chex0v/yt-time-tracker/cmd/app/cmd/user"
	"github.com/spf13/cobra"
	"os"
)

var (
	task string
	date string
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
	cmd := trackingtypes.TrackingTypesCmd
	reportCmd := report.MyReportByTodayCmd
	addCmd := add.AddCmd
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(taskInfo.TaskInfoCmd)
	rootCmd.AddCommand(trackingtypes.TrackingTypesCmd)
	rootCmd.AddCommand(user.MyUserInfoCmd)
	rootCmd.AddCommand(reportCmd)
	cmd.Flags().StringVarP(&task, "task", "t", "", "Task number")
	reportCmd.Flags().StringVarP(&date, "date", "d", "", "Date of report. Format YYYY-MM-DD")
	addCmd.Flags().StringVarP(&task, "date", "d", "", "Date for add. Format YYYY-MM-DD")
}
func init() {
	addCommands()
}
