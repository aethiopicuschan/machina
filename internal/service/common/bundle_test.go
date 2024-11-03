package common_test

import (
	"testing"

	"github.com/aethiopicuschan/machina/internal/service/common"
	"github.com/stretchr/testify/assert"
)

func TestBundle(t *testing.T) {
	testcases := []struct {
		src    common.Bundle
		expect string
	}{
		{
			src: common.Bundle{
				Method: "GET",
				Regexp: "/api/livestream/[0-9]+/reaction",
				Name:   "/api/livestream/:livestream_id/reaction",
			},
			expect: "[[Bundle]]\nregexp = '^(GET) /api/livestream/[0-9]+/reaction'\nname = '/api/livestream/:livestream_id/reaction'\n",
		},
	}

	for _, testcase := range testcases {
		got := testcase.src.String()
		assert.Equal(t, testcase.expect, got)
	}
}
