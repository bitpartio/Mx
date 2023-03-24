package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#forms

import . "github.com/bitpartio/gomx/utils"

func init() {
	ButtonOptions = buttonOptions{
		Formenctype: formenctypeOptions{
			Url:       formenctypeOptionUrl,
			Multipart: formenctypeOptionMultipart,
			Text:      formenctypeOptionText,
		},
		Formmethod: formmethodOptions{
			Get:  formmethodOptionGet,
			Post: formmethodOptionPost,
		},
		Type: typeOptions{
			Button: typeOptionButton,
			Reset:  typeOptionReset,
			Submit: typeOptionSubmit,
		},
	}
}

/*
 * An interactive element activated by a user with a mouse, keyboard,
 * finger, voice command, or other assistive technology. Once activated,
 * it then performs an action, such as submitting a form or opening a
 * dialog.
 */
func Button(props ButtonProps) string {
	autofocus := BuildBooleanProp("autofocus", props.Autofocus)
	disabled := BuildBooleanProp("disabled", props.Disabled)
	form := BuildProp("form", props.Form)
	formaction := BuildProp("formaction", props.Formaction)
	var formenctype string
	if props.Formenctype != nil {
		formenctype = BuildProp("formenctype", props.Formenctype().String())
	}
	var formmethod string
	if props.Formmethod != nil {
		formmethod = BuildProp("formmethod", props.Formmethod().String())
	}
	formnovalidate := BuildBooleanProp("formnovalidate", props.Formnovalidate)
	formtarget := BuildProp("formtarget", props.Formtarget)
	name := BuildProp("name", props.Name)
	var typeOf string
	if props.Type != nil {
		typeOf = BuildProp("type", props.Type().String())
	}
	value := BuildProp("value", props.Value)

	values := map[string]interface{}{
		"innerhtml":      props.InnerHTML,
		"autofocus":      autofocus,
		"disabled":       disabled,
		"form":           form,
		"formaction":     formaction,
		"formenctype":    formenctype,
		"formmethod":     formmethod,
		"formnovalidate": formnovalidate,
		"formtarget":     formtarget,
		"name":           name,
		"type":           typeOf,
		"value":          value,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<button {{global}} {{autofocus}} {{disabled}} {{form}} {{formaction}} {{formenctype}} {{formmethod}} {{formnovalidate}} {{formtarget}} {{name}} {{type}} {{value}}>{{innerhtml}}</button>`)

	s := Render(t, values)
	return s
}

type ButtonProps struct {
	GlobalProps
	InnerHTML      string
	Autofocus      bool
	Disabled       bool
	Form           string
	Formaction     string
	Formenctype    func() formenctypeOption
	Formmethod     func() formmethodOption
	Formnovalidate bool
	Formtarget     string
	Name           string
	Type           func() typeOption
	Value          string
}

type buttonOptions struct {
	Formenctype formenctypeOptions
	Formmethod  formmethodOptions
	Type        typeOptions
}

var ButtonOptions buttonOptions

/* Formenctype */
type formenctypeOption struct{ string }

func (o formenctypeOption) String() string { return o.string }

func formenctypeOptionUrl() formenctypeOption {
	return formenctypeOption{"application/x-www-form-urlencoded"}
}

func formenctypeOptionMultipart() formenctypeOption {
	return formenctypeOption{"multipart/form-data"}
}

func formenctypeOptionText() formenctypeOption {
	return formenctypeOption{"text/plain"}
}

type formenctypeOptions struct {
	Url       func() formenctypeOption
	Multipart func() formenctypeOption
	Text      func() formenctypeOption
}

/* Formmethod */
type formmethodOption struct{ string }

func (o formmethodOption) String() string { return o.string }

func formmethodOptionGet() formmethodOption {
	return formmethodOption{"get"}
}

func formmethodOptionPost() formmethodOption {
	return formmethodOption{"post"}
}

type formmethodOptions struct {
	Get  func() formmethodOption
	Post func() formmethodOption
}

/* Type */
type typeOption struct{ string }

func (o typeOption) String() string { return o.string }

func typeOptionButton() typeOption {
	return typeOption{"button"}
}

func typeOptionReset() typeOption {
	return typeOption{"reset"}
}

func typeOptionSubmit() typeOption {
	return typeOption{"submit"}
}

type typeOptions struct {
	Button func() typeOption
	Reset  func() typeOption
	Submit func() typeOption
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

	t := Mx(`<datalist {{global}}>{{innerhtml}}</datalist>`)

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

	t := Mx(`<fieldset {{global}}>{{innerhtml}}</fieldset>`)

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

	t := Mx(`<form {{global}}>{{innerhtml}}</form>`)

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

	t := Mx(`<input {{global}} />`)

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

	t := Mx(`<label {{global}}>{{innerhtml}}</label>`)

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

	t := Mx(`<legend {{global}}>{{innerhtml}}</legend>`)

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

	t := Mx(`<meter {{global}}>{{innerhtml}}</meter>`)

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

	t := Mx(`<optgroup {{global}}>{{innerhtml}}</optgroup>`)

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

	t := Mx(`<option {{global}}>{{innerhtml}}</option>`)

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

	t := Mx(`<output {{global}}>{{innerhtml}}</output>`)

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

	t := Mx(`<progress {{global}}>{{innerhtml}}</progress>`)

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

	t := Mx(`<select {{global}}>{{innerhtml}}</select>`)

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

	t := Mx(`<textarea {{global}}>{{innerhtml}}</textarea>`)

	s := Render(t, values)
	return s
}
