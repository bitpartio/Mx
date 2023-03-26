package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#demarcating_edits

import . "github.com/bitpartio/Mx/utils"

/*
 * Represents a range of text that has been added to a document. You can use
 * the <del> element to similarly represent a range of text that has been
 * deleted from the document.
 */
type DelProps struct {
	GlobalProps

	InnerHTML string
}

func Del(props DelProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<del {{global}}>{{innerhtml}}</del>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a range of text that has been deleted from a document. This
 * can be used when rendering "track changes" or source code diff
 * information, for example. The <ins> element can be used for the opposite
 * purpose: to indicate text that has been added to the document.
 */
type InsProps struct {
	GlobalProps

	InnerHTML string
}

func Ins(props InsProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<ins {{global}}>{{innerhtml}}</ins>`)

	s := Render(t, values)
	return s
}
