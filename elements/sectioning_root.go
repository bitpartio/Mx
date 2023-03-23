package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#sectioning_root

import . "github.com/bitpartio/gomx/utils"

/*
 * represents the content of an HTML document. There can be only one such
 * element in a document.
 */
type BodyProps struct {
	GlobalProps
	InnerHTML string
}

func Body(props BodyProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<body {{global}}>{{innerhtml}}</body>`)

	s := Render(t, values)
	return s
}
