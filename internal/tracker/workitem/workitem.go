package workitem

import (
	"github.com/chex0v/yt-time-tracker/internal/tracker/issue"
	"github.com/chex0v/yt-time-tracker/internal/tracker/user"
)

type Create struct {
	Text     string   `json:"text"`
	Duration Duration `json:"duration"`
	Type     Type     `json:"type"`
	Date     int64    `json:"date,omitempty"`
}

type WorkItem struct {
	Duration    Duration    `json:"duration"`
	Type        Type        `json:"type"`
	Date        int64       `json:"date,omitempty"`
	Created     int64       `json:"created"`
	Creator     user.User   `json:"creator"`
	Author      user.User   `json:"author"`
	Id          string      `json:"id"`
	TextPreview string      `json:"textPreview,omitempty"`
	Text        string      `json:"text,omitempty"`
	Issue       issue.Issue `json:"issue"`
}
