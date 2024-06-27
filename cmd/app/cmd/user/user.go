package user

import (
	"github.com/chex0v/yt-time-tracker/internal/progressbar"
	"github.com/chex0v/yt-time-tracker/internal/tracker"
	"github.com/cheynewallace/tabby"
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

	s := progressbar.NewProgressBar()
	yt := tracker.NewTracker()
	s.Start()
	user, err := yt.MyUserInfo()
	s.Stop()

	if err != nil {
		log.Fatal(err)
	}

	t := tabby.New()

	t.AddLine("ID: ", user.Id)
	t.AddLine("Email: ", user.Email)
	t.AddLine("Login: ", user.Login)
	t.AddLine("Full name: ", user.FullName)
	t.AddLine("Is online: ", user.Online)

	t.Print()

	return nil
}
