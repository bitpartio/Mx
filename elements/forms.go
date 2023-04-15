package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#forms

import . "github.com/bitpartio/Mx/utils"

func init() {
	Input = inputTypes{
		Button: InputButton,
	}
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
	FormOptions = formOptions{
		Autocomplete: autocompleteOptions{
			On:  autocompleteOptionOn,
			Off: autocompleteOptionOff,
		},
		Method: methodOptions{
			Dialog: methodOptionDialog,
			Get:    methodOptionGet,
			Post:   methodOptionPost,
		},
		Rel: formRelOptions{
			External:   formRelOptionExternal,
			Help:       formRelOptionHelp,
			License:    formRelOptionLicense,
			Next:       formRelOptionNext,
			NoFollow:   formRelOptionNoFollow,
			NoOpener:   formRelOptionNoOpener,
			NoReferrer: formRelOptionNoReferrer,
			Opener:     formRelOptionOpener,
			Prev:       formRelOptionPrev,
			Search:     formRelOptionSearch,
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
	var formenctype string
	if props.Formenctype != nil {
		formenctype = BuildProp("formenctype", props.Formenctype().String())
	}
	var formmethod string
	if props.Formmethod != nil {
		formmethod = BuildProp("formmethod", props.Formmethod().String())
	}
	var typeOf string
	if props.Type != nil {
		typeOf = BuildProp("type", props.Type().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"autofocus":      BuildBooleanProp("autofocus", props.Autofocus),
		"disabled":       BuildBooleanProp("disabled", props.Disabled),
		"form":           BuildProp("form", props.Form),
		"formaction":     BuildProp("formaction", props.Formaction),
		"formenctype":    formenctype,
		"formmethod":     formmethod,
		"formnovalidate": BuildBooleanProp("formnovalidate", props.Formnovalidate),
		"formtarget":     BuildProp("formtarget", props.Formtarget),
		"name":           BuildProp("name", props.Name),
		"type":           typeOf,
		"value":          BuildProp("value", props.Value),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<button {{global}} {{autofocus}} {{disabled}} {{form}} {{formaction}} {{formenctype}} {{formmethod}} {{formnovalidate}} {{formtarget}} {{name}} {{type}} {{value}}>{{innerhtml}}</button>`)

	s := Render(t, values)
	return s
}

type ButtonProps struct {
	GlobalProps

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

	InnerHTML string
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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

	Disabled bool
	Form     string
	Name     string

	InnerHTML string
}

func Fieldset(props FieldsetProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"disabled": BuildBooleanProp("disabled", props.Disabled),
		"form":     BuildProp("form", props.Form),
		"name":     BuildProp("name", props.Name),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<fieldset {{global}} {{disabled} {{form}} {{name}}>{{innerhtml}}</fieldset>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a document section containing interactive controls for
 * submitting information.
 */
type FormProps struct {
	GlobalProps

	AcceptCharset string
	Action        string
	Autocomplete  func() autocompleteOption
	Enctype       string
	Method        func() methodOption
	Novalidate    bool
	Name          string
	Rel           func() formRelOption
	Target        string

	InnerHTML string
}

func Form(props FormProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"accept-charset": BuildProp("accept-charset", props.AcceptCharset),
		"action":         BuildProp("action", props.Action),
		"autocomplete":   BuildProp("autocomplete", props.Autocomplete().String()),
		"enctype":        BuildProp("enctype", props.Enctype),
		"method":         BuildProp("method", props.Method().String()),
		"novalidate":     BuildBooleanProp("novalidate", props.Novalidate),
		"name":           BuildProp("name", props.Name),
		"rel":            BuildProp("rel", props.Rel().String()),
		"target":         BuildProp("target", props.Target),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<form {{global}} {{accept-charset}} {{action}} {{autocomplete}} {{enctype}} {{method}} {{novalidate}} {{name}} {{rel}} {{target}}>{{innerhtml}}</form>`)

	s := Render(t, values)
	return s
}

type formOptions struct {
	Autocomplete autocompleteOptions
	Method       methodOptions
	Rel          formRelOptions
}

var FormOptions formOptions

/* Autocomplete */
type autocompleteOption struct{ string }

func (o autocompleteOption) String() string { return o.string }

func autocompleteOptionOff() autocompleteOption {
	return autocompleteOption{"off"}
}

func autocompleteOptionOn() autocompleteOption {
	return autocompleteOption{"on"}
}

type autocompleteOptions struct {
	Off func() autocompleteOption
	On  func() autocompleteOption
}

/* Method */
type methodOption struct{ string }

func (o methodOption) String() string { return o.string }

func methodOptionDialog() methodOption {
	return methodOption{"dialog"}
}

func methodOptionGet() methodOption {
	return methodOption{"get"}
}

func methodOptionPost() methodOption {
	return methodOption{"post"}
}

type methodOptions struct {
	Dialog func() methodOption
	Get    func() methodOption
	Post   func() methodOption
}

/* Rel */
type formRelOption struct{ string }

func (o formRelOption) String() string { return o.string }

func formRelOptionExternal() formRelOption {
	return formRelOption{"external"}
}
func formRelOptionHelp() formRelOption {
	return formRelOption{"help"}
}
func formRelOptionLicense() formRelOption {
	return formRelOption{"license"}
}
func formRelOptionNext() formRelOption {
	return formRelOption{"next"}
}
func formRelOptionNoFollow() formRelOption {
	return formRelOption{"nofollow"}
}
func formRelOptionNoOpener() formRelOption {
	return formRelOption{"noopener"}
}
func formRelOptionNoReferrer() formRelOption {
	return formRelOption{"noreferrer"}
}
func formRelOptionOpener() formRelOption {
	return formRelOption{"opener"}
}
func formRelOptionPrev() formRelOption {
	return formRelOption{"prev"}
}
func formRelOptionSearch() formRelOption {
	return formRelOption{"search"}
}

type formRelOptions struct {
	External   func() formRelOption
	Help       func() formRelOption
	License    func() formRelOption
	Next       func() formRelOption
	NoFollow   func() formRelOption
	NoOpener   func() formRelOption
	NoReferrer func() formRelOption
	Opener     func() formRelOption
	Prev       func() formRelOption
	Search     func() formRelOption
}

type FormRelOptions []func() formRelOption

/*
 * Used to create interactive controls for web-based forms in order to
 * accept data from the user; a wide variety of types of input data and
 * control widgets are available, depending on the device and user agent.
 * The <input> element is one of the most powerful and complex in all of
 * HTML due to the sheer number of combinations of input types and
 * attributes.
 */

type inputTypes struct {
	Button        func(props InputButtonProps) string
	Checkbox      func(props InputCheckboxProps) string
	Color         func(props InputColorProps) string
	Date          func(props InputDateProps) string
	DatetimeLocal func(props InputDatetimeLocalProps) string
	Email         func(props InputEmailProps) string
	File          func(props InputFileProps) string
	Hidden        func(props InputHiddenProps) string
	Image         func(props InputImageProps) string
	Month         func(props InputMonthProps) string
	Number        func(props InputNumberProps) string
	Password      func(props InputPasswordProps) string
	Radio         func(props InputRadioProps) string
	Range         func(props InputRangeProps) string
	Reset         func(props InputResetProps) string
	Search        func(props InputSearchProps) string
	Submit        func(props InputSubmitProps) string
	Tel           func(props InputTelProps) string
	Text          func(props InputTextProps) string
	Time          func(props InputTimeProps) string
	Url           func(props InputUrlProps) string
	Week          func(props InputWeekProps) string
}

var Input inputTypes

/* Input Button */
type InputButtonProps struct {
	GlobalProps
}

func InputButton(props InputButtonProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Checkbox */
type InputCheckboxProps struct {
	GlobalProps
}

func InputCheckbox(props InputCheckboxProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Color */
type InputColorProps struct {
	GlobalProps
}

func InputColor(props InputColorProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Date */
type InputDateProps struct {
	GlobalProps
}

func InputDate(props InputDateProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input DatetimeLocal */
type InputDatetimeLocalProps struct {
	GlobalProps
}

func InputDatetimeLocal(props InputDatetimeLocalProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Email */
type InputEmailProps struct {
	GlobalProps
}

func InputEmail(props InputEmailProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input File */
type InputFileProps struct {
	GlobalProps
}

func InputFile(props InputFileProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Hidden */
type InputHiddenProps struct {
	GlobalProps
}

func InputHidden(props InputHiddenProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Image */
type InputImageProps struct {
	GlobalProps
}

func InputImage(props InputImageProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Month */
type InputMonthProps struct {
	GlobalProps
}

func InputMonth(props InputMonthProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Number */
type InputNumberProps struct {
	GlobalProps
}

func InputNumber(props InputNumberProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Password */
type InputPasswordProps struct {
	GlobalProps
}

func InputPassword(props InputPasswordProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Radio */
type InputRadioProps struct {
	GlobalProps
}

func InputRadio(props InputRadioProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Range */
type InputRangeProps struct {
	GlobalProps
}

func InputRange(props InputRangeProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Reset */
type InputResetProps struct {
	GlobalProps
}

func InputReset(props InputResetProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Search */
type InputSearchProps struct {
	GlobalProps
}

func InputSearch(props InputSearchProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Submit */
type InputSubmitProps struct {
	GlobalProps
}

func InputSubmit(props InputSubmitProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Tel */
type InputTelProps struct {
	GlobalProps
}

func InputTel(props InputTelProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Text */
type InputTextProps struct {
	GlobalProps
}

func InputText(props InputTextProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Time */
type InputTimeProps struct {
	GlobalProps
}

func InputTime(props InputTimeProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Url */
type InputUrlProps struct {
	GlobalProps
}

func InputUrl(props InputUrlProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<input {{global}} />`)

	s := Render(t, values)
	return s
}

/* Input Week */
type InputWeekProps struct {
	GlobalProps
}

func InputWeek(props InputWeekProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<textarea {{global}}>{{innerhtml}}</textarea>`)

	s := Render(t, values)
	return s
}
