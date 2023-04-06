package utils

import (
	"strconv"
	"strings"
	"time"
)

// BuildProp
func BuildProp(name, prop string) string {
	if prop != "" {
		var s strings.Builder
		s.WriteString(name)
		s.WriteString(`="`)
		s.WriteString(prop)
		s.WriteString(`"`)
		return s.String()
	}

	return ""
}

// BuildPropList
func BuildPropList(name string, props []string) string {
	if len(props) > 0 {
		prop := strings.Join(props, " ")
		var s strings.Builder
		s.WriteString(name)
		s.WriteString(`="`)
		s.WriteString(prop)
		s.WriteString(`"`)
		return s.String()
	}

	return ""
}

// BuildBooleanProp
func BuildBooleanProp(name string, prop bool) (s string) {
	if prop == true {
		s = name
	}
	return s
}

// BuildNumberProp
func BuildNumberProp(name string, prop int) string {
	var s strings.Builder
	s.WriteString(name)
	s.WriteString(`="`)
	s.WriteString(strconv.Itoa(prop))
	s.WriteString(`"`)
	return s.String()
}

// BuildDateProp
func BuildDateProp(name string, prop time.Time) string {

	var s strings.Builder
	s.WriteString(name)
	s.WriteString(`="`)
	s.WriteString(prop.Format(time.RFC3339))
	s.WriteString(`"`)
	return s.String()
}

// BuildProps
func BuildProps(prefix string, attr map[string]string) string {
	if len(attr) > 0 {
		var attrBuilder strings.Builder
		for key, d := range attr {
			var keyBuilder strings.Builder
			keyBuilder.WriteString(prefix)
			keyBuilder.WriteString(key)
			fullKey := keyBuilder.String()
			attrBuilder.WriteString(fullKey)
			attrBuilder.WriteString(`="`)
			attrBuilder.WriteString(d)
			attrBuilder.WriteString(`" `)
		}
		return attrBuilder.String()
	}

	return ""
}

// BuildDataValues
func BuildDataValues(data DataValues) string {
	return BuildProps("data-", data)
}

// BuildAriaRoles
func BuildAriaRoles(aria AriaRoles) string {
	return BuildProps("aria-", aria)
}

// BuildHtmxProps
func BuildHtmxProps(hx HtmxProps) string {
	return BuildProps("hx-", hx)
}

// BuildID
func BuildID(id string) string {
	s := ""
	if id != "" {
		var b strings.Builder
		b.WriteString(`id="`)
		b.WriteString(id)
		b.WriteString(`"`)
		s = b.String()
		return s
	}

	return ""
}
