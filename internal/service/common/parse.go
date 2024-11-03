package common

import "fmt"

type Parse struct {
	LogFormat     string
	RequestIndex  int
	StatusIndex   int
	BytesIndex    int
	DurationIndex int
}

func (p Parse) String() string {
	return fmt.Sprintf("log_format = '%s'\nrequest_index = %d\nstatus_index = %d\nbytes_index = %d\nduration_index = %d\n", p.LogFormat, p.RequestIndex, p.StatusIndex, p.BytesIndex, p.DurationIndex)
}
