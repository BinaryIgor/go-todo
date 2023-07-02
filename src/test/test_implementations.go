package test

import "time"

type TestClock struct {
	now time.Time
}

func NewTestClock() TestClock {
	return TestClock{time.Now().UTC()}
}

func (t *TestClock) SetTime(time time.Time) {
	t.now = time
}

func (t *TestClock) AddTime(duration time.Duration) {
	t.now = t.now.Add(duration)
}

func (t TestClock) Now() time.Time {
	return t.now
}
