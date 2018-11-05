package model

import (
	"fmt"
	"time"

	"github.com/mizkei/self-lint-sample/app/data"
	"github.com/mizkei/self-lint-sample/app/test"
	atime "github.com/mizkei/self-lint-sample/app/time"
)

// Greet returns greeting message.
func Greet(data data.Data) string {
	now := atime.Now()

	var msg string
	switch {
	case now.Hour() < 12:
		msg = "good morning"
	case now.Hour() < 18:
		msg = "good afternoon"
	default:
		msg = "good evening"
	}

	return fmt.Sprintf("(%s), %s.", data.ID, msg)
}

// GreetIllegal is ...
func GreetIllegal() string {
	atime.SetTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local))
	return Greet(test.MakeTestData())
}
