package elements

import . "github.com/bitpartio/gomx/utils"

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#text_content
// Updated: 2023-03-08

/*
 * Indicates that the enclosed text is an extended quotation. Usually, this
 * is rendered visually by indentation. A URL for the source of the
 * quotation may be given using the cite attribute, while a text
 * representation of the source can be given using the <cite> element.
 */
type BlockquoteProps struct {
	GlobalProps
	InnerHTML string
}

func Blockquote(props BlockquoteProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <blockquote {{global}}>
		  {{innerhtml}}
		</blockquote>
  `)

	s := Render(t, values)
	return s
}

/*
 * Provides the description, definition, or value for the preceding
 * term (<dt>) in a description list (<dl>).
 */
type DdProps struct {
	GlobalProps
	InnerHTML string
}

func Dd(props DdProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <dd {{global}}>{{innerhtml}}</dd>
  `)

	s := Render(t, values)
	return s
}

/*
 * The generic container for flow content. It has no effect on the content
 * or layout until styled in some way using CSS (e.g., styling is directly
 * applied to it, or some kind of layout model like flexbox is applied to
 * its parent element).
 */
type DivProps struct {
	GlobalProps
	InnerHTML string
}

func Div(props DivProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <div {{global}}>
		  {{innerhtml}}
		</div>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a description list. The element encloses a list of groups of
 * terms (specified using the <dt> element) and descriptions (provided by
 * <dd> elements). Common uses for this element are to implement a glossary
 * or to display metadata (a list of key-value pairs).
 */
type DlProps struct {
	GlobalProps
	InnerHTML string
}

func Dl(props DlProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <dl {{global}}>
		  {{innerhtml}}
		</dl>
  `)

	s := Render(t, values)
	return s
}

/*
 * Specifies a term in a description or definition list, and as such must be
 * used inside a <dl> element. It is usually followed by a <dd> element;
 * however, multiple <dt> elements in a row indicate several terms that
 * are all defined by the immediate next <dd> element.
 */
type DtProps struct {
	GlobalProps
	InnerHTML string
}

func Dt(props DtProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <dt {{global}}>{{innerhtml}}</dt>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a caption or legend describing the rest of the contents of its
 * parent <figure> element.
 */
type FigcaptionProps struct {
	GlobalProps
	InnerHTML string
}

func Figcaption(props FigcaptionProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <figcaption {{global}}>{{innerhtml}}</figcaption>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents self-contained content, potentially with an optional caption,
 * which is specified using the <figcaption> element. The figure, its
 * caption, and its contents are referenced as a single unit.
 */
type FigureProps struct {
	GlobalProps
	InnerHTML string
}

func Figure(props FigureProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <figure {{global}}>
		  {{innerhtml}}
		</figure>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a thematic break between paragraph-level elements: for
 * example, a change of scene in a story, or a shift of topic within a
 * section.
 */
type HrProps struct {
	GlobalProps
}

func Hr(props HrProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <hr {{global}} />
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents an item in a list. It must be contained in a parent element:
 * an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>).
 * In menus and unordered lists, list items are usually displayed using
 * bullet points. In ordered lists, they are usually displayed with an
 * ascending counter on the left, such as a number or letter.
 */
type LiProps struct {
	GlobalProps
	InnerHTML string
}

func Li(props LiProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <li {{global}}>{{innerhtml}}</li>
  `)

	s := Render(t, values)
	return s
}

/*
 * A semantic alternative to <ul>, but treated by browsers (and exposed
 * through the accessibility tree) as no different than <ul>. It represents
 * an unordered list of items (which are represented by <li> elements).
 */
type MenuProps struct {
	GlobalProps
	InnerHTML string
}

func Menu(props MenuProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <menu {{global}}>
		  {{innerhtml}}
		</menu>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents an ordered list of items — typically rendered as a numbered
 * list.
 */
type OlProps struct {
	GlobalProps
	InnerHTML string
}

func Ol(props OlProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <ol {{global}}>
		  {{innerhtml}}
		</ol>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents a paragraph. Paragraphs are usually represented in visual
 * media as blocks of text separated from adjacent blocks by blank lines
 * and/or first-line indentation, but HTML paragraphs can be any structural
 * grouping of related content, such as images or form fields.
 */
type PProps struct {
	GlobalProps
	InnerHTML string
}

func P(props PProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <p {{global}}>
		  {{innerhtml}}
		</p>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents preformatted text which is to be presented exactly as
 * written in the HTML file. The text is typically rendered using a
 * non-proportional, or monospaced, font. Whitespace inside this element
 * is displayed as written.
 */
type PreProps struct {
	GlobalProps
	InnerHTML string
}

func Pre(props PreProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <pre {{global}}>
		  {{innerhtml}}
		</pre>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents an unordered list of items, typically rendered as a bulleted
 * list.
 */
type UlProps struct {
	GlobalProps
	InnerHTML string
}

func Ul(props UlProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <ul {{global}}>
		  {{innerhtml}}
		</ul>
  `)

	s := Render(t, values)
	return s
}
