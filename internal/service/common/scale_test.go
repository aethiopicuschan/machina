package common_test

import (
	"testing"

	"github.com/aethiopicuschan/machina/internal/service/common"
	"github.com/stretchr/testify/assert"
)

func TestScale(t *testing.T) {
	testcases := []struct {
		src    common.Scale
		expect string
	}{
		{
			src: common.Scale{
				Scale:          10,
				EffectiveDigit: 5,
			},
			expect: "scale = 10\neffective_digit = 5\n",
		},
		{
			src: common.Scale{
				Scale:          -5,
				EffectiveDigit: 3,
			},
			expect: "scale = -5\neffective_digit = 3\n",
		},
	}

	for _, testcase := range testcases {
		got := testcase.src.String()
		assert.Equal(t, testcase.expect, got)
	}
}
