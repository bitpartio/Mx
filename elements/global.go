package elements

import (
	"strings"

	. "github.com/bitpartio/Mx/utils"
)

// Better dev experience once Golang supports seamless instantiation.
// Go Issue #9859 - https://github.com/golang/go/issues/9859

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes

/*
 * Props
 *   All properties collected into a single struct
 */
type GlobalProps struct {
	Accesskey             []rune
	Aria                  AriaRoles
	Htmx                  HtmxProps
	Autocapitalize        func() autocapitalizeOption
	Autofocus             bool
	Class                 []string
	Contenteditable       bool
	Data                  DataValues
	Dir                   func() dirOption
	Draggable             bool
	Enterkeyhint          string
	Exportparts           []string
	Hidden                func() hiddenOption
	ID                    string
	Inert                 string
	Inputmode             string
	Is                    string
	Itemid                string
	Itemprop              string
	Itemref               string
	Itemscope             string
	Itemtype              string
	Lang                  string
	Nonce                 string
	Part                  string
	Slot                  string
	Spellcheck            string
	Style                 string
	Tabindex              string
	Title                 string
	Translate             string
	Virtualkeyboardpolicy string
	// Events
	Onabort             string
	Onautocomplete      string
	Onautocompleteerror string
	Onblur              string
	Oncancel            string
	Oncanplay           string
	Oncanplaythrough    string
	Onchange            string
	Onclick             string
	Onclose             string
	Oncontextmenu       string
	Oncuechange         string
	Ondblclick          string
	Ondrag              string
	Ondragend           string
	Ondragenter         string
	Ondragleave         string
	Ondragover          string
	Ondragstart         string
	Ondrop              string
	Ondurationchange    string
	Onemptied           string
	Onended             string
	Onerror             string
	Onfocus             string
	Oninput             string
	Oninvalid           string
	Onkeydown           string
	Onkeypress          string
	Onkeyup             string
	Onload              string
	Onloadeddata        string
	Onloadedmetadata    string
	Onloadstart         string
	Onmousedown         string
	Onmouseenter        string
	Onmouseleave        string
	Onmousemove         string
	Onmouseout          string
	Onmouseover         string
	Onmouseup           string
	Onmousewheel        string
	Onpause             string
	Onplay              string
	Onplaying           string
	Onprogress          string
	Onratechange        string
	Onreset             string
	Onresize            string
	Onscroll            string
	Onseeked            string
	Onseeking           string
	Onselect            string
	Onshow              string
	Onsort              string
	Onstalled           string
	Onsubmit            string
	Onsuspend           string
	Ontimeupdate        string
	Ontoggle            string
	Onvolumechange      string
	Onwaiting           string
}

/*
 * Options
 *   Encapsulate options into predetermined functions
 */
type globalOptions struct {
	Autocapitalize autocapitalizeOptions
	Dir            dirOptions
	Hidden         hiddenOptions
}

var GlobalOptions globalOptions

// Initialize
func init() {
	GlobalOptions = globalOptions{
		Autocapitalize: autocapitalizeOptions{
			On:         autocapitalizeOptionOn,
			Sentences:  autocapitalizeOptionSentences,
			Off:        autocapitalizeOptionOff,
			None:       autocapitalizeOptionNone,
			Words:      autocapitalizeOptionWords,
			Characters: autocapitalizeOptionCharacters,
		},
		Dir: dirOptions{
			Ltr:  dirOptionLtr,
			Rtr:  dirOptionRtr,
			Auto: dirOptionAuto,
		},
		Hidden: hiddenOptions{
			Hidden:     hiddenOptionHidden,
			UntilFound: hiddenOptionUntilFound,
		},
	}
}

/*
 * Autocapitalize
 */
type autocapitalizeOption struct{ string }

func (o autocapitalizeOption) String() string { return o.string }

func autocapitalizeOptionOn() autocapitalizeOption {
	return autocapitalizeOption{"on"}
}

func autocapitalizeOptionSentences() autocapitalizeOption {
	return autocapitalizeOption{"sentences"}
}

func autocapitalizeOptionOff() autocapitalizeOption {
	return autocapitalizeOption{"off"}
}

func autocapitalizeOptionNone() autocapitalizeOption {
	return autocapitalizeOption{"none"}
}

func autocapitalizeOptionWords() autocapitalizeOption {
	return autocapitalizeOption{"words"}
}

func autocapitalizeOptionCharacters() autocapitalizeOption {
	return autocapitalizeOption{"characters"}
}

type autocapitalizeOptions struct {
	On         func() autocapitalizeOption
	Sentences  func() autocapitalizeOption
	Off        func() autocapitalizeOption
	None       func() autocapitalizeOption
	Words      func() autocapitalizeOption
	Characters func() autocapitalizeOption
}

/*
 * Dir
 */

type dirOption struct{ string }

func (o dirOption) String() string { return o.string }

func dirOptionLtr() dirOption {
	return dirOption{"ltr"}
}

func dirOptionRtr() dirOption {
	return dirOption{"rtr"}
}

func dirOptionAuto() dirOption {
	return dirOption{"auto"}
}

type dirOptions struct {
	Ltr  func() dirOption
	Rtr  func() dirOption
	Auto func() dirOption
}

/*
 * Hidden
 */

type hiddenOption struct{ string }

func (o hiddenOption) String() string { return o.string }

func hiddenOptionHidden() hiddenOption {
	return hiddenOption{"hidden"}
}

func hiddenOptionUntilFound() hiddenOption {
	return hiddenOption{"until-found"}
}

func hiddenOptionAuto() hiddenOption {
	return hiddenOption{"auto"}
}

type hiddenOptions struct {
	Hidden     func() hiddenOption
	UntilFound func() hiddenOption
}

/*
 * buildGlobalProps
 */
func buildGlobalProps(props GlobalProps) map[string]interface{} {
	values := make(map[string]interface{})

	// id
	values["id"] = BuildProp("id", props.ID)

	// class
	values["class"] = BuildPropList("class", props.Class)

	// autocapitalize
	if props.Autocapitalize != nil {
		values["autocapitalize"] = BuildProp("autocapitalize", props.Autocapitalize().String())
	}

	// dir
	if props.Dir != nil {
		values["dir"] = BuildProp("dir", props.Dir().String())
	}

	// hidden
	if props.Hidden != nil {
		values["hidden"] = BuildProp("hidden", props.Hidden().String())
	}

	values["data"] = BuildDataValues(props.Data)
	values["aria"] = BuildAriaRoles(props.Aria)
	values["htmx"] = BuildHtmxProps(props.Htmx)

	return values
}

/*
 * BuildGlobalProps
 */
func BuildGlobalProps(props GlobalProps) string {
	values := buildGlobalProps(props)

	tags := []string{
		"{{id}}",
		"{{class}}",
		"{{autocapitalize}}",
		"{{dir}}",
		"{{hidden}}",
		"{{data}}",
		"{{aria}}",
		"{{htmx}}",
	}

	t := Mx(strings.Join(tags[:], " "))

	s := Render(t, values)
	return s
}
