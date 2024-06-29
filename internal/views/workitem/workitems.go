package views

import (
	"fmt"
	"github.com/chex0v/yt-time-tracker/internal/tracker/workitem"
	"github.com/cheynewallace/tabby"
	"log"
	"sort"
	"strconv"
	"time"
)

func WorkItems(workItems workitem.WorkItems) {
	t := tabby.New()

	groupByData := workItems.GroupByDate()

	keys := make([]int64, 0, len(groupByData))
	for k := range groupByData {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	totalM := 0
	for _, d := range keys {
		tm := time.Unix(d/1000, 0)
		t.AddLine("Дата: ", tm.Format(time.DateOnly), "")
		for _, item := range groupByData[d] {
			t.AddLine(item.Id, item.Author.Name, item.Duration.Presentation, item.Text, item.Issue.Link())
			totalM += item.Duration.Minutes
		}
	}
	d, err := time.ParseDuration(strconv.Itoa(totalM) + "m")

	if err != nil {
		log.Fatal(err)
	}
	t.AddLine("-------------")
	t.AddLine("Итого: ", fmt.Sprintf("%.2f часов", d.Hours()))
	t.Print()
}
