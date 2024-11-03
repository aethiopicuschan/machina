package common_test

import (
	"testing"

	"github.com/aethiopicuschan/machina/internal/service/common"
)

func TestGeneral(t *testing.T) {
	testcases := []struct {
		src    common.General
		expect string
	}{
		{
			src: common.General{
				RankingCount:   20,
				SlowCount:      37,
				ShowStddev:     true,
				ShowStatusCode: false,
				ShowBytes:      true,
				Percentiles:    []float64{50.0, 90.0, 95.0, 99.0},
			},
			expect: "ranking_count = 20\nslow_count = 37\nshow_stddev = true\nshow_status_code = false\nshow_bytes = true\npercentiles = [50.0, 90.0, 95.0, 99.0]\n",
		},
	}

	for _, testcase := range testcases {
		got := testcase.src.String()
		if got != testcase.expect {
			t.Errorf("expect: %s, got: %s", testcase.expect, got)
		}
	}
}
