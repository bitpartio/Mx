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
	InnerHTML string
	Lang      string
}

func HTML(props HTMLProps) string {
	lang := BuildProp("lang", props.Lang)

	values := map[string]interface{}{
		"lang":      lang,
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <html {{lang}} {{global}}>
		  {{innerhtml}}
		</html>
  `)

	s := Render(t, values)
	return s
}
