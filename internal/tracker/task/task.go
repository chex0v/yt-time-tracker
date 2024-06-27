package task

import "github.com/chex0v/yt-time-tracker/internal/tracker/project"

type Task struct {
	ID      string          `json:"id"`
	Project project.Project `json:"project"`
	Name    string          `json:"name"`
}
