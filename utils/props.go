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
func BuildPropList(name string, props []string, joiner string) string {
	if len(props) > 0 {
		prop := strings.Join(props, joiner)
		var s strings.Builder
		s.WriteString(name)
		s.WriteString(`="`)
		s.WriteString(prop)
		s.WriteString(`"`)
		return s.String()
	}

	return ""
}

// BuildPropListWithCommas
func BuildPropListWithCommas(name string, props []string) string {
	return BuildPropList(name, props, ", ")
}

// BuildPropListWithSpaces
func BuildPropListWithSpaces(name string, props []string) string {
	return BuildPropList(name, props, " ")
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

// BuildChronosProp
func BuildChronosProp(name string, prop time.Time, format string) string {
	var s strings.Builder
	s.WriteString(name)
	s.WriteString(`="`)
	s.WriteString(prop.Format(format))
	s.WriteString(`"`)
	return s.String()
}

// BuildDateProp
func BuildDateProp(name string, prop time.Time) string {
	return BuildChronosProp(name, prop, "2006-01-02")
}

// BuildDateTimeProp
func BuildDateTimeProp(name string, prop time.Time) string {
	return BuildChronosProp(name, prop, time.RFC3339)
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
