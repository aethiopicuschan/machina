package common_test

import (
	"testing"

	"github.com/aethiopicuschan/machina/internal/service/common"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testcases := []struct {
		src    common.Parse
		expect string
	}{
		{
			src: common.Parse{
				LogFormat:     `LogFormat`,
				RequestIndex:  5,
				StatusIndex:   6,
				BytesIndex:    7,
				DurationIndex: 10,
			},
			expect: "log_format = 'LogFormat'\nrequest_index = 5\nstatus_index = 6\nbytes_index = 7\nduration_index = 10\n",
		},
	}

	for _, testcase := range testcases {
		got := testcase.src.String()
		assert.Equal(t, testcase.expect, got)
	}
}
