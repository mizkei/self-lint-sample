package test

import (
	"strconv"
	"time"

	"github.com/mizkei/self-lint-sample/app/data"
)

// MakeTestData makes data.Data.
func MakeTestData() data.Data {
	return data.Data{ID: strconv.FormatInt(time.Now().Unix(), 10)}
}
