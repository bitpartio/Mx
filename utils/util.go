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

// BuildMarkup
func BuildMarkup(tag string, values map[string]interface{}) string {
	var s strings.Builder

	s.WriteString("<")
	s.WriteString(tag)
	s.WriteString(" ")

	closed := true

	for p := range values {
		if p == "innerhtml" {
			closed = false
			continue
		}
		s.WriteString("{{")
		s.WriteString(p)
		s.WriteString("}} ")
	}

	if closed {
		s.WriteString("/>")
	} else {
		s.WriteString(">{{")
		s.WriteString("innerhtml")
		s.WriteString("}}</")
		s.WriteString(tag)
		s.WriteString(">")
	}

	return s.String()
}

// Create a reference pointer for a value
// From: https://stackoverflow.com/questions/30716354/how-do-i-do-a-literal-int64-in-go
func Ptr[T any](v T) *T {
	return &v
}
