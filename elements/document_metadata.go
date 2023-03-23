package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#document_metadata

import (
	. "github.com/bitpartio/gomx/utils"
)

/*
 * Specifies the base URL to use for all relative URLs in a document. There
 * can be only one such element in a document.
 */
type BaseProps struct {
	GlobalProps
}

func Base(props BaseProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<base {{global}} />`)

	s := Render(t, values)
	return s
}

/*
 * Contains machine-readable information (metadata) about the document,
 * like its title, scripts, and style sheets.
 */
type HeadProps struct {
	InnerHTML string
	GlobalProps
}

func Head(props HeadProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<head {{global}}>{{innerhtml}}</head>`)

	s := Render(t, values)
	return s
}

/*
 * Specifies relationships between the current document and an external
 * resource. This element is most commonly used to link to CSS, but is also
 * used to establish site icons (both "favicon" style icons and icons for
 * the home screen and apps on mobile devices) among other things.
 */
type LinkProps struct {
	GlobalProps
}

func Link(props LinkProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<link {{global}} />`)

	s := Render(t, values)
	return s
}

/*
 * Represents metadata that cannot be represented by other HTML meta-related
 * elements, like <base>, <link>, <script>, <style> and <title>.
 */
type MetaProps struct {
	GlobalProps
	Charset string
}

func Meta(props MetaProps) string {
	values := map[string]interface{}{
		"global":  RenderGlobalProps(props.GlobalProps),
		"charset": BuildProp("charset", props.Charset),
	}

	t := Mx(`<meta {{global}} {{charset}} />`)

	s := Render(t, values)
	return s
}

/*
 * Contains style information for a document, or part of a document. It
 * contains CSS, which is applied to the contents of the document containing
 * this element.
 */
type StyleProps struct {
	InnerHTML string
	GlobalProps
}

func Style(props StyleProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<style {{global}}>{{innerhtml}}</style>`)

	s := Render(t, values)
	return s
}

/*
 * Defines the document's title that is shown in a browser's title bar or a
 * page's tab. It only contains text; tags within the element are ignored.
 */
type TitleProps struct {
	InnerHTML string
	GlobalProps
}

func Title(props TitleProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`<title {{global}}>{{innerhtml}}</title>`)

	s := Render(t, values)
	return s
}
