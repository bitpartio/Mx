package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#forms

import . "github.com/bitpartio/gomx/utils"

/*
 * An interactive element activated by a user with a mouse, keyboard,
 * finger, voice command, or other assistive technology. Once activated,
 * it then performs an action, such as submitting a form or opening a
 * dialog.
 */
type ButtonProps struct {
	GlobalProps
	InnerHTML string
	Disabled  bool
}

func Button(props ButtonProps) string {
	disabled := BuildBooleanProp("disabled", props.Disabled)

	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,
		"disabled":  disabled,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <button {{global}} {{disabled}}>{{innerhtml}}</button>
  `)

	s := Render(t, values)
	return s
}

/*
 * Contains a set of <option> elements that represent the permissible or
 * recommended options available to choose from within other controls.
 */
type DatalistProps struct {
	GlobalProps
	InnerHTML string
}

func Datalist(props DatalistProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <datalist {{global}}>
		  {{innerhtml}}
		</datalist>
  `)

	s := Render(t, values)
	return s
}

/*
 * Used to group several controls as well as labels (<label>) within a web
 * form.
 */
type FieldsetProps struct {
	GlobalProps
	InnerHTML string
}

func Fieldset(props FieldsetProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <fieldset {{global}}>
		  {{innerhtml}}
		</fieldset>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a document section containing interactive controls for
 * submitting information.
 */
type FormProps struct {
	GlobalProps
	InnerHTML string
}

func Form(props FormProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <form {{global}}>
		  {{innerhtml}}
		</form>
  `)

	s := Render(t, values)
	return s
}

/*
 * Used to create interactive controls for web-based forms in order to
 * accept data from the user; a wide variety of types of input data and
 * control widgets are available, depending on the device and user agent.
 * The <input> element is one of the most powerful and complex in all of
 * HTML due to the sheer number of combinations of input types and
 * attributes.
 */
type InputProps struct {
	GlobalProps
}

func Input(props InputProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <input {{global}} />
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a caption for an item in a user interface.
 */
type LabelProps struct {
	GlobalProps
	InnerHTML string
}

func Label(props LabelProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <label {{global}}>
		  {{innerhtml}}
		</label>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a caption for the content of its parent <fieldset>.
 */
type LegendProps struct {
	GlobalProps
	InnerHTML string
}

func Legend(props LegendProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <legend {{global}}>
		  {{innerhtml}}
		</legend>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents either a scalar value within a known range or
 * a fractional value.
 */
type MeterProps struct {
	GlobalProps
	InnerHTML string
}

func Meter(props MeterProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <meter {{global}}>
		  {{innerhtml}}
		</meter>
  `)

	s := Render(t, values)
	return s
}

/*
 * Creates a grouping of options within a <select> element.
 */
type OptgroupProps struct {
	GlobalProps
	InnerHTML string
}

func Optgroup(props OptgroupProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <optgroup {{global}}>
		  {{innerhtml}}
		</optgroup>
  `)

	s := Render(t, values)
	return s
}

/*
 * Used to define an item contained in a select, an <optgroup>, or a
 * <datalist> element. As such, <option> can represent menu items in
 * popups and other lists of items in an HTML document.
 */
type OptionProps struct {
	GlobalProps
	InnerHTML string
}

func Option(props OptionProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <option {{global}}>{{innerhtml}}</option>
  `)

	s := Render(t, values)
	return s
}

/*
 * Container element into which a site or app can inject the results
 * of a calculation or the outcome of a user action.
 */
type OutputProps struct {
	GlobalProps
	InnerHTML string
}

func Output(props OutputProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <output {{global}}>{{innerhtml}}</output>
  `)

	s := Render(t, values)
	return s
}

/*
 * Displays an indicator showing the completion progress of a task,
 * typically displayed as a progress bar.
 */
type ProgressProps struct {
	GlobalProps
	InnerHTML string
}

func Progress(props ProgressProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <progress {{global}}>{{innerhtml}}</progress>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a control that provides a menu of options.
 */
type SelectProps struct {
	GlobalProps
	InnerHTML string
}

func Select(props SelectProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <select {{global}}>
		  {{innerhtml}}
		</select>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a multi-line plain-text editing control, useful when you
 * want to allow users to enter a sizeable amount of free-form text, for
 * example a comment on a review or feedback form.
 */
type TextareaProps struct {
	GlobalProps
	InnerHTML string
}

func Textarea(props TextareaProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <textarea {{global}}>
		  {{innerhtml}}
		</textarea>
  `)

	s := Render(t, values)
	return s
}
