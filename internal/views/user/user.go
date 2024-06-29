package view

import (
	user2 "github.com/chex0v/yt-time-tracker/internal/tracker/user"
	"github.com/cheynewallace/tabby"
)

func User(u user2.User) {
	t := tabby.New()

	t.AddLine("ID: ", u.Id)
	t.AddLine("Email: ", u.Email)
	t.AddLine("Login: ", u.Login)
	t.AddLine("Full name: ", u.FullName)
	t.AddLine("Is online: ", u.Online)

	t.Print()
}
