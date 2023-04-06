package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#interactive_elements

import . "github.com/bitpartio/Mx/utils"

/*
 * Creates a disclosure widget in which information is visible only when the
 * widget is toggled into an "open" state. A summary or label must be
 * provided using the <summary> element.
 */
type DetailsProps struct {
	GlobalProps

	Open bool

	InnerHTML string
}

func Details(props DetailsProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"open": BuildBooleanProp("open", props.Open),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<details {{global}} {{open}}>{{innerhtml}}</details>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a dialog box or other interactive component, such as a
 * dismissible alert, inspector, or subwindow.
 */
type DialogProps struct {
	GlobalProps

	InnerHTML string
}

func Dialog(props DialogProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<dialog {{global}}>{{innerhtml}}</dialog>`)

	s := Render(t, values)
	return s
}

/*
 * Specifies a summary, caption, or legend for a details element's
 * disclosure box. Clicking the <summary> element toggles the state of the
 * parent <details> element open and closed.
 */
type SummaryProps struct {
	GlobalProps

	InnerHTML string
}

func Summary(props SummaryProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<summary {{global}}>{{innerhtml}}</summary>`)

	s := Render(t, values)
	return s
}
