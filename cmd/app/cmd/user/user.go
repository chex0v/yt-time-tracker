package user

import (
	"github.com/chex0v/yt-time-tracker/internal/config"
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

func userInfo(cmd *cobra.Command, args []string) error {

	config := config.GetConfig()

	client := tracker.NewClient(config.ApiUrl, config.Token)

	s := progressbar.NewProgressBar()

	s.Start()
	user, err := client.MyUserInfo()
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
