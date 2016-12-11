package utils

import (
	"github.com/pascallimeux/urmmongo/utils/log"
	"time"
)

type Timer struct {
	Start   time.Time
	Elapsed time.Duration
}

func (t *Timer) StartTimer() {
	t.Start = time.Now()
}

func (t *Timer) GetTimer() time.Duration {
	t.Elapsed = time.Since(t.Start)
	return t.Elapsed
}

func (t *Timer) LogElapsed(location, label string) {
	t.Elapsed = time.Since(t.Start)
	log.Info(location, "processing ", label, " in ", t.Elapsed.String())
}
