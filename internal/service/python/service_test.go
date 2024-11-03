package python_test

import (
	"io"
	"strings"
	"testing"

	"github.com/aethiopicuschan/machina/internal/service/common"
	"github.com/aethiopicuschan/machina/internal/service/python"
	"github.com/stretchr/testify/assert"
)

func TestPythonService(t *testing.T) {
	testcases := []struct {
		src    string
		expect []common.Bundle
	}{
		{
			src: `
				@app.route("/api/livestream/<int:livestream_id>/reaction", methods=["GET, "POST"])
				@app.route("/api/user/<string:username>/livestream/<int:livestream_id>", methods=["GET"])
			`,
			expect: []common.Bundle{
				{
					Method: "GET",
					Regexp: "/api/livestream/[0-9]+/reaction",
					Name:   "/api/livestream/:livestream_id/reaction",
				},
				{
					Method: "POST",
					Regexp: "/api/livestream/[0-9]+/reaction",
					Name:   "/api/livestream/:livestream_id/reaction",
				},
				{
					Method: "GET",
					Regexp: "/api/user/[0-9a-zA-Z]+/livestream/[0-9]+",
					Name:   "/api/user/:username/livestream/:livestream_id",
				},
			},
		},
	}

	for _, testcase := range testcases {
		s := &python.PythonService{}
		err := s.Load(io.NopCloser(strings.NewReader(testcase.src)))
		assert.NoError(t, err)
		got, err := s.GenerateBundles()
		assert.NoError(t, err)
		assert.Equal(t, testcase.expect, got)
		s.Close()
	}
}
