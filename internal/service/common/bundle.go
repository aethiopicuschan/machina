package common

import (
	"fmt"
)

type Bundle struct {
	Method string
	Regexp string
	Name   string
}

func (b Bundle) String() string {
	return fmt.Sprintf("[[Bundle]]\nregexp = '^(%s) %s'\nname = '%s'\n", b.Method, b.Regexp, b.Name)
}
