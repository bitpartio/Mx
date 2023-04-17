package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#document_metadata

import (
	. "github.com/bitpartio/Mx/utils"
)

/*
 * Specifies the base URL to use for all relative URLs in a document. There
 * can be only one such element in a document.
 */
type BaseProps struct {
	GlobalProps

	Href   string
	Target string
}

func Base(props BaseProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"href":   BuildProp("href", props.Href),
		"target": BuildProp("target", props.Target),
	}

	m := BuildMarkup("base", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Contains machine-readable information (metadata) about the document,
 * like its title, scripts, and style sheets.
 */
type HeadProps struct {
	GlobalProps

	InnerHTML string
}

func Head(props HeadProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("head", values)

	t := Mx(m)

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
		"global": BuildGlobalProps(props.GlobalProps),
	}

	m := BuildMarkup("link", values)

	t := Mx(m)

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
		"global": BuildGlobalProps(props.GlobalProps),

		"charset": BuildProp("charset", props.Charset),
	}

	m := BuildMarkup("meta", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Contains style information for a document, or part of a document. It
 * contains CSS, which is applied to the contents of the document containing
 * this element.
 */
type StyleProps struct {
	GlobalProps

	InnerHTML string
}

func Style(props StyleProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("style", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Defines the document's title that is shown in a browser's title bar or a
 * page's tab. It only contains text; tags within the element are ignored.
 */
type TitleProps struct {
	GlobalProps

	InnerHTML string
}

func Title(props TitleProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("title", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}
