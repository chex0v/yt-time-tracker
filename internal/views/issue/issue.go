package view

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/tracker/issue"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	"github.com/cheynewallace/tabby"
	"github.com/manifoldco/promptui"
)

func Issue(issue issue.Issue) {
	tInfo := tabby.New()
	tInfo.AddLine(fmt.Sprintf("Id: %s, Project: %s", issue.ID, issue.Project.Name))
	tInfo.AddLine("")
	tInfo.Print()
}

func ChoiceType(types []workitem.Type) (workitem.Type, error) {

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
		return workitem.Type{}, err
	}

	return types[indexType], nil
}
