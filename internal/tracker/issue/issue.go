package issue

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/config"
	"github.com/chex0v/yt-time-tracker/internal/tracker/project"
	"log"
	"net/url"
)

type Issue struct {
	ID         string          `json:"id"`
	Summary    string          `json:"summary"`
	IdReadable string          `json:"idReadable"`
	Project    project.Project `json:"project,omitempty"`
	Name       string          `json:"name,omitempty"`
}

func (i Issue) Link() string {
	c := config.GetConfig()
	u, err := url.Parse(c.ApiUrl)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s://%s/issue/%s/", u.Scheme, u.Host, i.IdReadable)
}
