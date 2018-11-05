package model_test

import (
	"testing"
	"time"

	"github.com/mizkei/self-lint-sample/app/data"
	"github.com/mizkei/self-lint-sample/app/model"
	mtime "github.com/mizkei/self-lint-sample/app/time"
)

func TestGreet(t *testing.T) {
	defer mtime.ResetTime()
	for name, tc := range map[string]struct {
		time   time.Time
		data   data.Data
		expect string
	}{
		"morning": {
			time:   parseDatetime(t, "2018-11-04 11:59:59"),
			data:   data.Data{ID: "1234"},
			expect: "(1234), good morning.",
		},
		"evening": {
			time:   parseDatetime(t, "2018-11-04 18:00:00"),
			data:   data.Data{ID: "12345"},
			expect: "(12345), good evening.",
		},
		"afternoon": {
			time:   parseDatetime(t, "2018-11-04 17:59:59"),
			data:   data.Data{ID: "123456"},
			expect: "(123456), good afternoon.",
		},
	} {
		t.Run(name, func(t *testing.T) {
			mtime.SetTime(tc.time)
			res := model.Greet(tc.data)
			if res != tc.expect {
				t.Fatalf("fail, got:%s, want:%s", res, tc.expect)
			}
		})
	}
}

func parseDatetime(t *testing.T, v string) time.Time {
	t.Helper()
	tm, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		t.Fatal(err)
	}
	return tm
}
