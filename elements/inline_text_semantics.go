package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#inline_text_semantics

import (
	. "github.com/bitpartio/Mx/utils"
)

func init() {
	AOptions = aOptions{
		Referrerpolicy: referrerpolicyOptions{
			NoReferrer:          referrerpolicyOptionNoReferrer,
			NoReferrerDowngrade: referrerpolicyOptionNoReferrerDowngrade,
			Origin:              referrerpolicyOptionOrigin,
			CrossOrigin:         referrerpolicyOptionCrossOrigin,
			SameOrigin:          referrerpolicyOptionSameOrigin,
			StrictOrigin:        referrerpolicyOptionStrictOrigin,
			StrictCrossOrigin:   referrerpolicyOptionStrictCrossOrigin,
			Unsafe:              referrerpolicyOptionUnsafe,
		},
		Rel: aRelOptions{
			Alternate:  aRelOptionAlternate,
			Author:     aRelOptionAuthor,
			Bookmark:   aRelOptionBookmark,
			External:   aRelOptionExternal,
			Help:       aRelOptionHelp,
			License:    aRelOptionLicense,
			Next:       aRelOptionNext,
			NoFollow:   aRelOptionNoFollow,
			NoOpener:   aRelOptionNoOpener,
			NoReferrer: aRelOptionNoReferrer,
			Prev:       aRelOptionPrev,
			Search:     aRelOptionSearch,
			Tag:        aRelOptionTag,
		},
	}
}

/*
 * Together with its href attribute, creates a hyperlink to web pages,
 * files, email addresses, locations in the same page, or anything else
 * a URL can address.
 */
func A(props AProps) string {
	var referrerpolicy string
	if props.Referrerpolicy != nil {
		referrerpolicy = BuildProp("referrerpolicy", props.Referrerpolicy().String())
	}

	var rel string
	if len(props.Rel) > 0 {
		relStrings := make([]string, len(props.Rel))
		for k, rel := range props.Rel {
			relStrings[k] = rel().String()
		}

		rel = BuildPropList("rel", relStrings)
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"download":       BuildProp("download", props.Download),
		"href":           BuildProp("href", props.Href),
		"hreflang":       BuildProp("hreflang", props.Hreflang),
		"ping":           BuildPropList("ping", props.Ping),
		"referrerpolicy": referrerpolicy,
		"rel":            rel,
		"target":         BuildProp("target", props.Target),
		"type":           BuildProp("type", props.Type),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<a {{global}} {{download}} {{href}} {{hreflang}} {{ping}} {{referrerpolicy}} {{rel}} {{target}} {{type}}>{{innerhtml}}</a>`)

	s := Render(t, values)
	return s
}

type AProps struct {
	GlobalProps

	Download       string
	Href           string
	Hreflang       string // Limited values but too complex for enum. Ref: https://datatracker.ietf.org/doc/html/rfc5646
	Ping           []string
	Referrerpolicy func() referrerpolicyOption
	Rel            []func() aRelOption
	Target         string
	Type           string // Limited values but too complex for enum. Ref: https://www.iana.org/assignments/media-types/media-types.xhtml

	InnerHTML string
}

type aOptions struct {
	Referrerpolicy referrerpolicyOptions
	Rel            aRelOptions
}

var AOptions aOptions

/* Referrerpolicy */
type referrerpolicyOption struct{ string }

func (o referrerpolicyOption) String() string { return o.string }

func referrerpolicyOptionNoReferrer() referrerpolicyOption {
	return referrerpolicyOption{"no-referrer"}
}

func referrerpolicyOptionNoReferrerDowngrade() referrerpolicyOption {
	return referrerpolicyOption{"no-referrer-when-downgrade"}
}

func referrerpolicyOptionOrigin() referrerpolicyOption {
	return referrerpolicyOption{"origin"}
}

func referrerpolicyOptionCrossOrigin() referrerpolicyOption {
	return referrerpolicyOption{"origin-when-cross-origin"}
}

func referrerpolicyOptionSameOrigin() referrerpolicyOption {
	return referrerpolicyOption{"same-origin"}
}

func referrerpolicyOptionStrictOrigin() referrerpolicyOption {
	return referrerpolicyOption{"strict-origin"}
}

func referrerpolicyOptionStrictCrossOrigin() referrerpolicyOption {
	return referrerpolicyOption{"same-origin-when-cross-origin"}
}

func referrerpolicyOptionUnsafe() referrerpolicyOption {
	return referrerpolicyOption{"unsafe-url"}
}

type referrerpolicyOptions struct {
	NoReferrer          func() referrerpolicyOption
	NoReferrerDowngrade func() referrerpolicyOption
	Origin              func() referrerpolicyOption
	CrossOrigin         func() referrerpolicyOption
	SameOrigin          func() referrerpolicyOption
	StrictOrigin        func() referrerpolicyOption
	StrictCrossOrigin   func() referrerpolicyOption
	Unsafe              func() referrerpolicyOption
}

/* Rel */
type aRelOption struct{ string }

func (o aRelOption) String() string { return o.string }

func aRelOptionAlternate() aRelOption {
	return aRelOption{"alternate"}
}

func aRelOptionAuthor() aRelOption {
	return aRelOption{"author"}
}

func aRelOptionBookmark() aRelOption {
	return aRelOption{"bookmark"}
}
func aRelOptionExternal() aRelOption {
	return aRelOption{"external"}
}
func aRelOptionHelp() aRelOption {
	return aRelOption{"help"}
}
func aRelOptionLicense() aRelOption {
	return aRelOption{"license"}
}
func aRelOptionNext() aRelOption {
	return aRelOption{"next"}
}
func aRelOptionNoFollow() aRelOption {
	return aRelOption{"nofollow"}
}
func aRelOptionNoOpener() aRelOption {
	return aRelOption{"noopener"}
}
func aRelOptionNoReferrer() aRelOption {
	return aRelOption{"noreferrer"}
}
func aRelOptionPrev() aRelOption {
	return aRelOption{"prev"}
}
func aRelOptionSearch() aRelOption {
	return aRelOption{"search"}
}
func aRelOptionTag() aRelOption {
	return aRelOption{"tag"}
}

type aRelOptions struct {
	Alternate  func() aRelOption
	Author     func() aRelOption
	Bookmark   func() aRelOption
	External   func() aRelOption
	Help       func() aRelOption
	License    func() aRelOption
	Next       func() aRelOption
	NoFollow   func() aRelOption
	NoOpener   func() aRelOption
	NoReferrer func() aRelOption
	Prev       func() aRelOption
	Search     func() aRelOption
	Tag        func() aRelOption
}

type ARelOptions []func() aRelOption

/*
 * Represents an abbreviation or acronym.
 */
type AbbrProps struct {
	GlobalProps

	InnerHTML string
}

func Abbr(props AbbrProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<abbr {{global}}>{{innerhtml}}</abbr>`)

	s := Render(t, values)
	return s
}

/*
 * Used to draw the reader's attention to the element's contents, which are
 * not otherwise granted special importance. This was formerly known as the
 * Boldface element, and most browsers still draw the text in boldface.
 * However, you should not use <b> for styling text or granting importance.
 * If you wish to create boldface text, you should use the CSS font-weight
 * property. If you wish to indicate an element is of special importance,
 * you should use the strong element.
 */
type BProps struct {
	GlobalProps

	InnerHTML string
}

func B(props BProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<b {{global}}>{{innerhtml}}</b>`)

	s := Render(t, values)
	return s
}

/*
 * Tells the browser's bidirectional algorithm to treat the text it contains
 * in isolation from its surrounding text. It's particularly useful when a
 * website dynamically inserts some text and doesn't know the directionality
 * of the text being inserted.
 */
type BdiProps struct {
	GlobalProps

	InnerHTML string
}

func Bdi(props BdiProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<bdi {{global}}>{{innerhtml}}</bdi>`)

	s := Render(t, values)
	return s
}

/*
 * Overrides the current directionality of text, so that the text within is
 * rendered in a different direction.
 */
type BdoProps struct {
	GlobalProps

	InnerHTML string
}

func Bdo(props BdoProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<bdo {{global}}>{{innerhtml}}</bdo>`)

	s := Render(t, values)
	return s
}

/*
 * Produces a line break in text (carriage-return). It is useful for writing
 * a poem or an address, where the division of lines is significant.
 */
type BrProps struct {
	GlobalProps
}

func Br(props BrProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<br {{global}} />`)

	s := Render(t, values)
	return s
}

/*
 * Used to mark up the title of a cited creative work. The reference may be
 * in an abbreviated form according to context-appropriate conventions
 * related to citation metadata.
 */
type CiteProps struct {
	GlobalProps

	InnerHTML string
}

func Cite(props CiteProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<cite {{global}}>{{innerhtml}}</cite>`)

	s := Render(t, values)
	return s
}

/*
 * Displays its contents styled in a fashion intended to indicate that the
 * text is a short fragment of computer code. By default, the content text
 * is displayed using the user agent default monospace font.
 */
type CodeProps struct {
	GlobalProps

	InnerHTML string
}

func Code(props CodeProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<code {{global}}>{{innerhtml}}</code>`)

	s := Render(t, values)
	return s
}

/*
 * Links a given piece of content with a machine-readable translation. If
 * the content is time- or date-related, the time element must be used.
 */
type DataProps struct {
	GlobalProps

	InnerHTML string
}

func Data(props DataProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<data {{global}}>{{innerhtml}}</data>`)

	s := Render(t, values)
	return s
}

/*
 * Used to indicate the term being defined within the context of a
 * definition phrase or sentence. The ancestor <p> element, the
 * <dt>/<dd> pairing, or the nearest section ancestor of the <dfn> element,
 * is considered to be the definition of the term.
 */
type DfnProps struct {
	GlobalProps

	InnerHTML string
}

func Dfn(props DfnProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<dfn {{global}}>{{innerhtml}}</dfn>`)

	s := Render(t, values)
	return s
}

/*
 * Marks text that has stress emphasis. The <em> element can be nested,
 * with each level of nesting indicating a greater degree of emphasis.
 */
type EmProps struct {
	GlobalProps

	InnerHTML string
}

func Em(props EmProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<em {{global}}>{{innerhtml}}</em>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a range of text that is set off from the normal text for some
 * reason, such as idiomatic text, technical terms, taxonomical designations,
 * among others. Historically, these have been presented using italicized
 * type, which is the original source of the <i> naming of this element.
 */
type IProps struct {
	GlobalProps

	InnerHTML string
}

func I(props IProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<i {{global}}>{{innerhtml}}</i>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a span of inline text denoting textual user input from a
 * keyboard, voice input, or any other text entry device. By convention,
 * the user agent defaults to rendering the contents of a <kbd> element
 * using its default monospace font, although this is not mandated by
 * the HTML standard.
 */
type KbdProps struct {
	GlobalProps

	InnerHTML string
}

func Kbd(props KbdProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<kbd {{global}}>{{innerhtml}}</kbd>`)

	s := Render(t, values)
	return s
}

/*
 * Represents text which is marked or highlighted for reference or
 * notation purposes due to the marked passage's relevance in the
 * enclosing context.
 */
type MarkProps struct {
	GlobalProps

	InnerHTML string
}

func Mark(props MarkProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<mark {{global}}>{{innerhtml}}</mark>`)

	s := Render(t, values)
	return s
}

/*
 * Indicates that the enclosed text is a short inline quotation. Most
 * modern browsers implement this by surrounding the text in quotation
 * marks. This element is intended for short quotations that don't require
 * paragraph breaks; for long quotations use the <blockquote> element.
 */
type QProps struct {
	GlobalProps

	InnerHTML string
}

func Q(props QProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<q {{global}}>{{innerhtml}}</q>`)

	s := Render(t, values)
	return s
}

/*
 * Used to provide fall-back parentheses for browsers that do not support
 * display of ruby annotations using the <ruby> element. One <rp> element
 * should enclose each of the opening and closing parentheses that wrap
 * the <rt> element that contains the annotation's text.
 */
type RpProps struct {
	GlobalProps

	InnerHTML string
}

func Rp(props RpProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<rp {{global}}>{{innerhtml}}</rp>`)

	s := Render(t, values)
	return s
}

/*
 * Specifies the ruby text component of a ruby annotation, which is used to
 * provide pronunciation, translation, or transliteration information for
 * East Asian typography. The <rt> element must always be contained within
 * a <ruby> element.
 */
type RtProps struct {
	GlobalProps

	InnerHTML string
}

func Rt(props RtProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<rt {{global}}>{{innerhtml}}</rt>`)

	s := Render(t, values)
	return s
}

/*
 * Represents small annotations that are rendered above, below, or next to
 * base text, usually used for showing the pronunciation of East Asian
 * characters. It can also be used for annotating other kinds of text, but
 * this usage is less common.
 */
type RubyProps struct {
	GlobalProps

	InnerHTML string
}

func Ruby(props RubyProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<ruby {{global}}>{{innerhtml}}</ruby>`)

	s := Render(t, values)
	return s
}

/*
 * Renders text with a strikethrough, or a line through it. Use the <s>
 * element to represent things that are no longer relevant or no longer
 * accurate. However, <s> is not appropriate when indicating document
 * edits; for that, use the del and ins elements, as appropriate.
 */
type SProps struct {
	GlobalProps

	InnerHTML string
}

func S(props SProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<s {{global}}>{{innerhtml}}</s>`)

	s := Render(t, values)
	return s
}

/*
 * Used to enclose inline text which represents sample (or quoted) output
 * from a computer program. Its contents are typically rendered using the
 * browser's default monospaced font (such as Courier or Lucida Console).
 */
type SampProps struct {
	GlobalProps

	InnerHTML string
}

func Samp(props SampProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<samp {{global}}>{{innerhtml}}</samp>`)

	s := Render(t, values)
	return s
}

/*
 * Represents side-comments and small print, like copyright and legal text,
 * independent of its styled presentation. By default, it renders text
 * within it one font-size smaller, such as from small to x-small.
 */
type SmallProps struct {
	GlobalProps

	InnerHTML string
}

func Small(props SmallProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<small {{global}}>{{innerhtml}}</small>`)

	s := Render(t, values)
	return s
}

/*
 * A generic inline container for phrasing content, which does not
 * inherently represent anything. It can be used to group elements for
 * styling purposes (using the class or id attributes), or because they
 * share attribute values, such as lang. It should be used only when no
 * other semantic element is appropriate. <span> is very much like a div
 * element, but div is a block-level element whereas a <span> is an inline
 * element.
 */
type SpanProps struct {
	GlobalProps

	InnerHTML string
}

func Span(props SpanProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<span {{global}}>{{innerhtml}}</span>`)

	s := Render(t, values)
	return s
}

/*
 * Indicates that its contents have strong importance, seriousness, or
 * urgency. Browsers typically render the contents in bold type.
 */
type StrongProps struct {
	GlobalProps

	InnerHTML string
}

func Strong(props StrongProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<strong {{global}}>{{innerhtml}}</strong>`)

	s := Render(t, values)
	return s
}

/*
 * Specifies inline text which should be displayed as subscript for solely
 * typographical reasons. Subscripts are typically rendered with a lowered
 * baseline using smaller text.
 */
type SubProps struct {
	GlobalProps

	InnerHTML string
}

func Sub(props SubProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<sub {{global}}>{{innerhtml}}</sub>`)

	s := Render(t, values)
	return s
}

/*
 * Specifies inline text which is to be displayed as superscript for solely
 * typographical reasons. Superscripts are usually rendered with a raised
 * baseline using smaller text.
 */
type SupProps struct {
	GlobalProps

	InnerHTML string
}

func Sup(props SupProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<sup {{global}}>{{innerhtml}}</sup>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a specific period in time. It may include the datetime
 * attribute to translate dates into machine-readable format, allowing for
 * better search engine results or custom features such as reminders.
 */
type TimeProps struct {
	GlobalProps

	InnerHTML string
}

func Time(props TimeProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<time {{global}}>{{innerhtml}}</time>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a span of inline text which should be rendered in a way that
 * indicates that it has a non-textual annotation. This is rendered by
 * default as a simple solid underline, but may be altered using CSS.
 */
type UProps struct {
	GlobalProps

	InnerHTML string
}

func U(props UProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<u {{global}}>{{innerhtml}}</u>`)

	s := Render(t, values)
	return s
}

/*
 * Represents the name of a variable in a mathematical expression or a
 * programming context. It's typically presented using an italicized version
 * of the current typeface, although that behavior is browser-dependent.
 */
type VarProps struct {
	GlobalProps

	InnerHTML string
}

func Var(props VarProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<var {{global}}>{{innerhtml}}</var>`)

	s := Render(t, values)
	return s
}

/*
 * Represents a word break opportunity—a position within text where the
 * browser may optionally break a line, though its line-breaking rules
 * would not otherwise create a break at that location.
 */
type WbrProps struct {
	GlobalProps
}

func Wbr(props WbrProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<wbr {{global}} />`)

	s := Render(t, values)
	return s
}
