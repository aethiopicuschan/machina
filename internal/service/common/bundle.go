package common

import (
	"fmt"
	"strings"
)

type Bundle struct {
	Method string
	Regexp string
	Name   string
}

func (b Bundle) String() string {
	if !strings.Contains(b.Name, ":") {
		return ""
	}
	return fmt.Sprintf("[[Bundle]]\nregexp = '^(%s) %s'\nname = '%s'\n", b.Method, b.Regexp, b.Name)
}
