package progressbar

import (
	"github.com/briandowns/spinner"
	"time"
)

func NewProgressBar() *spinner.Spinner {
	return spinner.New(spinner.CharSets[9], 100*time.Millisecond)
}
