package user

import (
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/chex0v/yt-time-tracker/internal/tracker/user"
	view "github.com/chex0v/yt-time-tracker/internal/views/user"
	"github.com/spf13/cobra"
	"log"
)

var MyUserInfoCmd = &cobra.Command{
	Use:   "me",
	Short: "Информация о моём профиле",
	Long: `
	Получаем информацию о текущем профиле пользователя
	`,
	RunE: userInfo,
}

func userInfo(*cobra.Command, []string) error {

	u, err := progressbar.Progress(func() (user.User, error) {
		return tracker.NewTracker().MyUserInfo()
	})
	if err != nil {
		log.Fatal(err)
	}

	view.User(u)
	return nil
}
