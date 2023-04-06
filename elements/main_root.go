package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#main_root

import . "github.com/bitpartio/Mx/utils"

/*
 * Represents the root (top-level element) of an HTML document, so it is
 * also referred to as the root element. All other elements must be
 * descendants of this element.
 */
type HTMLProps struct {
	GlobalProps

	Lang  string
	Xmlns string

	InnerHTML string
}

func HTML(props HTMLProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"lang":  BuildProp("lang", props.Lang),
		"xmlns": BuildProp("xmlns", props.Xmlns),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<html {{global}} {{lang}} {{xmlns}}>{{innerhtml}}</html>`)

	s := Render(t, values)
	return s
}
