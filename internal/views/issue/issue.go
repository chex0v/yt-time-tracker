package view

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/tracker/issue"
	"github.com/cheynewallace/tabby"
)

func Issue(issue issue.Issue) {
	tInfo := tabby.New()
	tInfo.AddLine(fmt.Sprintf("Id: %s, Project: %s", issue.ID, issue.Project.Name))
	tInfo.AddLine("")
	tInfo.Print()
}
