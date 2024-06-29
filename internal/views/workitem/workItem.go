package views

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	"github.com/cheynewallace/tabby"
)

func WorkItem(item workitem.WorkItem) {
	t := tabby.New()

	t.AddLine(fmt.Sprintf("Id: %s", item.Id))
}
