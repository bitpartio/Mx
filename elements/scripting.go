package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#scripting

import . "github.com/bitpartio/gomx/utils"

/*
 * Container element to use with either the canvas scripting API or the
 * WebGL API to draw graphics and animations.
 */
type CanvasProps struct {
	GlobalProps
	InnerHTML string
}

func Canvas(props CanvasProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <canvas {{lang}} {{global}}>
		  {{innerhtml}}
		</canvas>
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a section of HTML to be inserted if a script type on the page
 * is unsupported or if scripting is currently turned off in the browser.
 */
type NoscriptProps struct {
	GlobalProps
	InnerHTML string
}

func Noscript(props NoscriptProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <noscript {{lang}} {{global}}>
		  {{innerhtml}}
		</noscript>
  `)

	s := Render(t, values)
	return s
}

/*
 * Used to embed executable code or data; this is typically used to embed
 * or refer to JavaScript code. The <script> element can also be used with
 * other languages, such as WebGL's GLSL shader programming language and
 * JSON.
 */
type ScriptProps struct {
	GlobalProps
	InnerHTML string
}

func Script(props ScriptProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <script {{lang}} {{global}}>
		  {{innerhtml}}
		</script>
  `)

	s := Render(t, values)
	return s
}
