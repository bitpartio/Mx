package utils

import (
	"strings"

	"github.com/valyala/fasttemplate"
)

// Mx
func Mx(s string) *Tpl {
	t := fasttemplate.New(s, "{{", "}}")
	return t
}

// Stack
func Stack(elements ...string) string {
	if len(elements) == 0 {
		return ""
	}

	var s strings.Builder

	for _, element := range elements {
		s.WriteString(element)
	}

	return s.String()
}

// Render
func Render(t *Tpl, v map[string]interface{}) string {
	s := t.ExecuteString(v)
	return s
}
