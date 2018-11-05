package time

import (
	stime "time"
)

var (
	defaultNow = stime.Now
	now        = defaultNow
)

// Now returns current time.
func Now() stime.Time {
	return now()
}

// SetTime set t to return value of Now.
func SetTime(t stime.Time) {
	now = func() stime.Time { return t }
}

// ResetTime reset Now.
func ResetTime() {
	now = defaultNow
}
