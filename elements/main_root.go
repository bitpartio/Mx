package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#main_root

import . "github.com/bitpartio/gomx/utils"

/*
 * Represents the root (top-level element) of an HTML document, so it is
 * also referred to as the root element. All other elements must be
 * descendants of this element.
 */
type HTMLProps struct {
	GlobalProps

	Lang string

	InnerHTML string
}

func HTML(props HTMLProps) string {
	lang := BuildProp("lang", props.Lang)

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"lang": lang,

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<html {{lang}} {{global}}>{{innerhtml}}</html>`)

	s := Render(t, values)
	return s
}
