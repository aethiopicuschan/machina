package common

import "fmt"

type General struct {
	RankingCount   int
	SlowCount      int
	ShowStddev     bool
	ShowStatusCode bool
	ShowBytes      bool
	Percentiles    []float64
}

func (g General) String() string {
	percentiles := ""
	for i, percentile := range g.Percentiles {
		percentiles += fmt.Sprintf("%0.1f", percentile)
		if i != len(g.Percentiles)-1 {
			percentiles += ", "
		}
	}
	result := fmt.Sprintf("ranking_count = %d\nslow_count = %d\nshow_stddev = %t\nshow_status_code = %t\nshow_bytes = %t\npercentiles = [%s]\n", g.RankingCount, g.SlowCount, g.ShowStddev, g.ShowStatusCode, g.ShowBytes, percentiles)

	return result
}
