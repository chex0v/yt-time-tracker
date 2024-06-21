package cmd

import (
	"os"

	add "github.com/chex0v/yt-time-tracker/cmd/app/cmd/add"
	taskInfo "github.com/chex0v/yt-time-tracker/cmd/app/cmd/taskinfo"
	trackingtypes "github.com/chex0v/yt-time-tracker/cmd/app/cmd/trackingtypes"
	"github.com/spf13/cobra"
)

var (
	task string
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
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(taskInfo.TaskInfoCmd)
	rootCmd.AddCommand(trackingtypes.TrackingTypesCmd)
	cmd.Flags().StringVarP(&task, "task", "t", "", "Task number")
}
func init() {
	addCommands()
}
