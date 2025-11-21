package times

import (
	"math/rand"
	"time"
)

func SimulateTime(min int, max int) time.Duration {
	return time.Duration(rand.Intn(max-min)+min) * time.Millisecond
}
