package python

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/aethiopicuschan/machina/internal/service/common"
)

type PythonService struct {
	code string
}

func (s *PythonService) Load(r io.ReadCloser) (err error) {
	defer r.Close()
	code, err := io.ReadAll(r)
	if err != nil {
		return
	}
	s.code = string(code)
	return
}

func (s *PythonService) GenerateBundles() (bundles []common.Bundle, err error) {
	re := regexp.MustCompile(`@app\.route\(\s*[\s\S]*?\)`)
	re2 := regexp.MustCompile(`@app.route\("([^"]+)",methods=\[([^\]]+)\]`)
	reNone := regexp.MustCompile(`<(\w+)>`)
	reStr := regexp.MustCompile(`<string:(\w+)>`)
	reInt := regexp.MustCompile(`<int:(\w+)>`)
	matches := re.FindAllString(s.code, -1)

	for _, match := range matches {
		// Windows style line endings
		trimmed := strings.ReplaceAll(match, "\r\n", "")
		// Unix style line endings
		trimmed = strings.ReplaceAll(trimmed, "\n", "")
		// Classic Mac style line endings
		trimmed = strings.ReplaceAll(trimmed, "\r", "")
		// Remove all spaces and tabs
		trimmed = strings.ReplaceAll(trimmed, " ", "")
		trimmed = strings.ReplaceAll(trimmed, "\t", "")
		// Remove all newlines
		trimmed = strings.ReplaceAll(trimmed, ",)", ")")

		// Find the path and methods
		matches := re2.FindStringSubmatch(trimmed)
		if len(matches) < 3 {
			err = fmt.Errorf("invalid route: %s", trimmed)
			return
		}
		path := matches[1]
		methods := strings.Split(matches[2], ",")
		for i, method := range methods {
			methods[i] = strings.ReplaceAll(method, "\"", "")
		}

		// Replace the path with regex and param
		replacedRegex := reNone.ReplaceAllString(path, "[0-9a-zA-Z]+")
		replacedParam := reNone.ReplaceAllString(path, ":$1")
		replacedRegex = reStr.ReplaceAllString(replacedRegex, "[0-9a-zA-Z]+")
		replacedParam = reStr.ReplaceAllString(replacedParam, ":$1")
		replacedRegex = reInt.ReplaceAllString(replacedRegex, "[0-9]+")
		replacedParam = reInt.ReplaceAllString(replacedParam, ":$1")

		// Create a bundle for each method
		for _, method := range methods {
			if strings.Contains(replacedParam, ":") {
				bundles = append(bundles, common.Bundle{
					Method: method,
					Regexp: replacedRegex,
					Name:   replacedParam,
				})
			}
		}
	}

	return
}

func (s *PythonService) Close() (err error) {
	return
}
