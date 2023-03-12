package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#image_and_multimedia

import . "github.com/bitpartio/gomx/utils"

/*
 * Defines an area inside an image map that has predefined clickable areas.
 * An image map allows geometric areas on an image to be associated
 * with hyperlink.
 */
type AreaProps struct {
	GlobalProps
}

func Area(props AreaProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <area {{global}} />
  `)

	s := Render(t, values)
	return s
}

/*
 * Used to embed sound content in documents. It may contain one or more
 * audio sources, represented using the src attribute or the source element:
 * the browser will choose the most suitable one. It can also be the
 * destination for streamed media, using a MediaStream.
 */
type AudioProps struct {
	GlobalProps
	InnerHTML string
}

func Audio(props AudioProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <audio {{global}}>
		  {{innerhtml}}
		</audio>
  `)

	s := Render(t, values)
	return s
}

/*
 * Embeds an image into the document.
 */
type ImgProps struct {
	GlobalProps
}

func Img(props ImgProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <img {{global}} />
  `)

	s := Render(t, values)
	return s
}

/*
 * Used with <area> elements to define an image map (a clickable link area).
 */
type MapProps struct {
	GlobalProps
	InnerHTML string
}

func Map(props MapProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <map {{global}}>
		  {{innerhtml}}
		</map>
  `)

	s := Render(t, values)
	return s
}

/*
 * Used as a child of the media elements, audio and video. It lets you
 * specify timed text tracks (or time-based data), for example to
 * automatically handle subtitles. The tracks are formatted in WebVTT
 * format (.vtt files)—Web Video Text Tracks.
 */
type TrackProps struct {
	GlobalProps
}

func Track(props TrackProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <track {{global}} />
  `)

	s := Render(t, values)
	return s
}

/*
 * Embeds a media player which supports video playback into the document.
 * You can use <video> for audio content as well, but the audio element may
 * provide a more appropriate user experience.
 */
type VideoProps struct {
	GlobalProps
	InnerHTML string
}

func Video(props VideoProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <video {{global}}>
		  {{innerhtml}}
		</video>
  `)

	s := Render(t, values)
	return s
}
