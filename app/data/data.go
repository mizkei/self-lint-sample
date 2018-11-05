package data

import (
	"fmt"
)

// Data is data.
type Data struct {
	ID string
}

func (d Data) String() string {
	if d.ID == "" {
		panic("empty ID")
	}
	return fmt.Sprintf("id:%s", d.ID)
}
