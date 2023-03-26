package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#content_sectioning

import . "github.com/bitpartio/gomx/utils"

/*
 * Indicates that the enclosed HTML provides contact information for a
 * person or people, or for an organization.
 */
type AddressProps struct {
	GlobalProps

	InnerHTML string
}

func Address(props AddressProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<address {{global}}>{{innerhtml}}</address>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a self-contained composition in a document, page,
 * application, or site, which is intended to be independently
 * distributable or reusable (e.g., in syndication). Examples include:
 * a forum post, a magazine or newspaper article, or a blog entry,
 * a product card, a user-submitted comment, an interactive widget or
 * gadget, or any other independent item of content.
 */
type ArticleProps struct {
	GlobalProps

	InnerHTML string
}

func Article(props ArticleProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<article {{global}}>{{innerhtml}}</article>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a portion of a document whose content is only indirectly
 * related to the document's main content. Asides are frequently presented
 * as sidebars or call-out boxes.
 */
type AsideProps struct {
	GlobalProps

	InnerHTML string
}

func Aside(props AsideProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<aside {{global}}>{{innerhtml}}</aside>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a footer for its nearest ancestor sectioning content or
 * sectioning root element. A <footer> typically contains information
 * about the author of the section, copyright data or links to related
 * documents.
 */
type FooterProps struct {
	GlobalProps

	InnerHTML string
}

func Footer(props FooterProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<footer {{global}}>{{innerhtml}}</footer>`)

	s := Render(t, values)
	return s
}

/*
 * Represents introductory content, typically a group of introductory or
 * navigational aids. It may contain some heading elements but also a logo,
 * a search form, an author name, and other elements.
 */
type HeaderProps struct {
	GlobalProps

	InnerHTML string
}

func Header(props HeaderProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<header {{global}}>{{innerhtml}}</header>`)

	s := Render(t, values)
	return s
}

/*
 * Represent six levels of section headings. <h1> is the highest section
 * level and <h6> is the lowest.
 */
type HProps struct {
	GlobalProps

	InnerHTML string
}

func H1(props HProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<h1 {{global}}>{{innerhtml}}</h1>`)

	s := Render(t, values)
	return s
}

func H2(props HProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<h2 {{global}}>{{innerhtml}}</h2>`)

	s := Render(t, values)
	return s
}

func H3(props HProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<h3 {{global}}>{{innerhtml}}</h3>`)

	s := Render(t, values)
	return s
}

func H4(props HProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<h4 {{global}}>{{innerhtml}}</h4>`)

	s := Render(t, values)
	return s
}

func H5(props HProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<h5 {{global}}>{{innerhtml}}</h5>`)

	s := Render(t, values)
	return s
}

func H6(props HProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<h6 {{global}}>{{innerhtml}}</h6>`)

	s := Render(t, values)
	return s
}

/*
 * Represents the dominant content of the body of a document. The main
 * content area consists of content that is directly related to or
 * expands upon the central topic of a document, or the central
 * functionality of an application.
 */
type MainProps struct {
	GlobalProps

	InnerHTML string
}

func Main(props MainProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<main {{global}}>{{innerhtml}}</main>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a section of a page whose purpose is to provide navigation
 * links, either within the current document or to other documents. Common
 * examples of navigation sections are menus, tables of contents, and
 * indexes
 */
type NavProps struct {
	GlobalProps

	InnerHTML string
}

func Nav(props NavProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<nav {{global}}>{{innerhtml}}</nav>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a generic standalone section of a document, which doesn't
 * have a more specific semantic element to represent it. Sections should
 * always have a heading, with very few exceptions.
 */
type SectionProps struct {
	GlobalProps

	InnerHTML string
}

func Section(props SectionProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<section {{global}}>{{innerhtml}}</section>`)

	s := Render(t, values)
	return s
}
