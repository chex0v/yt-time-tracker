package progressbar

import (
	"github.com/briandowns/spinner"
	"time"
)

func NewProgressBar() *spinner.Spinner {
	return spinner.New(spinner.CharSets[9], 100*time.Millisecond)
}

type funcProgress[T any] func() (T, error)

func Progress[T any](fn funcProgress[T]) (T, error) {
	s := NewProgressBar()
	s.Start()
	result, err := fn()
	s.Stop()
	return result, err
}
