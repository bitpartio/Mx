package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#embedded_content

import . "github.com/bitpartio/gomx/utils"

/*
 * Embeds external content at the specified point in the document. This
 * content is provided by an external application or other source of
 * interactive content such as a browser plug-in.
 */
type EmbedProps struct {
	GlobalProps
}

func Embed(props EmbedProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<embed {{global}} />`)

	s := Render(t, values)
	return s
}

/*
 * Represents a nested browsing context, embedding another HTML page
 * into the current one.
 */
type IframeProps struct {
	GlobalProps

	InnerHTML string
}

func Iframe(props IframeProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<iframe {{global}}>{{innerhtml}}</iframe>`)

	s := Render(t, values)
	return s
}

/*
 * Represents an external resource, which can be treated as an image, a
 * nested browsing context, or a resource to be handled by a plugin.
 */
type ObjectProps struct {
	GlobalProps

	InnerHTML string
}

func Object(props ObjectProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<object {{global}}>{{innerhtml}}</object>`)

	s := Render(t, values)
	return s
}

/*
 * Contains zero or more <source> elements and one <img> element to offer
 * alternative versions of an image for different display/device scenarios.
 */
type PictureProps struct {
	GlobalProps

	InnerHTML string
}

func Picture(props PictureProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<picture {{global}}>{{innerhtml}}</picture>`)

	s := Render(t, values)
	return s
}

/*
 * Specifies multiple media resources for the picture, the audio element,
 * or the video element. It is a void element, meaning that it has no
 * content and does not have a closing tag. It is commonly used to offer the
 * same media content in multiple file formats in order to provide
 * compatibility with a broad range of browsers given their differing
 * support for image file formats and media file formats.
 */
type SourceProps struct {
	GlobalProps
}

func Source(props SourceProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<source {{global}} />`)

	s := Render(t, values)
	return s
}
