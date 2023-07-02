package shared

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type utcClock struct {
}

var UtcClock = utcClock{}

func (c *utcClock) Now() time.Time {
	return time.Now().UTC()
}
