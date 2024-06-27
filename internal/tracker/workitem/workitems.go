package workitem

import "github.com/chex0v/yt-time-tracker/internal/util"

type WorkItems struct {
	Items []WorkItem `json:"workItems"`
}

func (items WorkItems) GroupByDate() map[int64][]WorkItem {

	w := make([]WorkItem, len(items.Items), cap(items.Items))
	copy(w, items.Items)

	return util.GroupByProperty(w, func(w WorkItem) int64 {
		return w.Date
	})

}
