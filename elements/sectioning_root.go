package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#sectioning_root

import . "github.com/bitpartio/Mx/utils"

/*
 * represents the content of an HTML document. There can be only one such
 * element in a document.
 */
type BodyProps struct {
	GlobalProps

	Onafterprint     string
	Onbeforeprint    string
	Onbeforeunload   string
	Onblur           string
	Onerror          string
	Onfocus          string
	Onhashchange     string
	Onlanguagechange string
	Onload           string
	Onmessage        string
	Onoffline        string
	Ononline         string
	Onpopstate       string
	Onredo           string
	Onresize         string
	Onstorage        string
	Onundo           string
	Onunload         string

	InnerHTML string
}

func Body(props BodyProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"onafterprint":     BuildProp("onafterprint", props.Onafterprint),
		"onbeforeprint":    BuildProp("onbeforeprint", props.Onbeforeprint),
		"onbeforeunload":   BuildProp("onbeforeunload", props.Onbeforeunload),
		"onblur":           BuildProp("onblur", props.Onblur),
		"onerror":          BuildProp("onerror", props.Onerror),
		"onfocus":          BuildProp("onfocus", props.Onfocus),
		"onhashchange":     BuildProp("onhashchange", props.Onhashchange),
		"onlanguagechange": BuildProp("onlanguagechange", props.Onlanguagechange),
		"onload":           BuildProp("onload", props.Onload),
		"onmessage":        BuildProp("onmessage", props.Onmessage),
		"onoffline":        BuildProp("onoffline", props.Onoffline),
		"ononline":         BuildProp("ononline", props.Ononline),
		"onpopstate":       BuildProp("onpopstate", props.Onpopstate),
		"onredo":           BuildProp("onredo", props.Onredo),
		"onresize":         BuildProp("onresize", props.Onresize),
		"onstorage":        BuildProp("onstorage", props.Onstorage),
		"onundo":           BuildProp("onundo", props.Onundo),
		"onunload":         BuildProp("onunload", props.Onunload),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<body {{global}} {{onafterprint}} {{onbeforeprint}} {{onbeforeunload}} {{onblur}} {{onerror}} {{onfocus}} {{onhashchange}} {{onlanguagechange}} {{onload}} {{onmessage}} {{onoffline}} {{ononline}} {{onpopstate}} {{onredo}} {{onresize}} {{onstorage}} {{onundo}} {{onunload}}>{{innerhtml}}</body>`)

	s := Render(t, values)
	return s
}
