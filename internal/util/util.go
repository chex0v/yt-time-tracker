package util

import "time"

var (
	TODAY     = []string{"today", "now"}
	YESTERDAY = []string{"yesterday", "last"}
	TOMORROW  = []string{"tomorrow", "next"}
)

func GroupByProperty[T any, K comparable](items []T, getProperty func(T) K) map[K][]T {
	grouped := make(map[K][]T)

	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}

func GetDynamicDayByMacros(macros string) string {
	hashDate := make(map[string]time.Time, len(TODAY)+len(YESTERDAY)+len(TOMORROW))
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)

	for _, m := range TODAY {
		hashDate[m] = today
	}
	for _, m := range YESTERDAY {
		hashDate[m] = yesterday
	}
	for _, m := range TOMORROW {
		hashDate[m] = tomorrow
	}
	v, ok := hashDate[macros]

	if ok {
		return v.Format(time.DateOnly)
	}

	return macros
}
