package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#image_and_multimedia

import (
	"strconv"
	"strings"

	. "github.com/bitpartio/Mx/utils"
)

func init() {
	AreaOptions = areaOptions{
		Shape: shapeOptions{
			Rect:   shapeOptionRect,
			Circle: shapeOptionCircle,
			Poly:   shapeOptionPoly,
		},
		Coords: coordsOptions{
			Rect:   coordsOptionRect,
			Circle: coordsOptionCircle,
			Poly:   coordsOptionPoly,
		},
	}
}

/*
 * Defines an area inside an image map that has predefined clickable areas.
 * An image map allows geometric areas on an image to be associated
 * with hyperlink.
 */
type AreaProps struct {
	GlobalProps

	Alt    string
	Coords coordsOption
	Href   string
	Shape  func() shapeOption
}

func Area(props AreaProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"alt":    BuildProp("alt", props.Alt),
		"coords": BuildProp("coords", props.Coords.String()),
		"href":   BuildProp("href", props.Href),
		"shape":  BuildProp("shape", props.Shape().String()),
	}

	t := Mx(`<area {{global}} {{alt}} {{coords}} {{href}} {{shape}}/>`)

	s := Render(t, values)
	return s
}

type areaOptions struct {
	Coords coordsOptions
	Shape  shapeOptions
}

var AreaOptions areaOptions

/* Coords */
type coordsOption struct{ string }

func (o coordsOption) String() string { return o.string }

func coordsOptionRect(coords ...int) coordsOption {
	if len(coords) < 4 {
		return coordsOption{""}
	}

	x1s := strconv.Itoa(coords[0])
	y1s := strconv.Itoa(coords[1])
	x2s := strconv.Itoa(coords[2])
	y2s := strconv.Itoa(coords[3])

	var s strings.Builder
	s.WriteString(x1s)
	s.WriteString(",")
	s.WriteString(y1s)
	s.WriteString(",")
	s.WriteString(x2s)
	s.WriteString(",")
	s.WriteString(y2s)
	return coordsOption{s.String()}
}

func coordsOptionCircle(coords ...int) coordsOption {
	if len(coords) < 3 {
		return coordsOption{""}
	}

	x1s := strconv.Itoa(coords[0])
	y1s := strconv.Itoa(coords[1])
	rs := strconv.Itoa(coords[2])

	var s strings.Builder
	s.WriteString(x1s)
	s.WriteString(",")
	s.WriteString(y1s)
	s.WriteString(",")
	s.WriteString(rs)
	return coordsOption{s.String()}
}

func coordsOptionPoly(coords ...int) coordsOption {
	if len(coords) < 2 {
		return coordsOption{""}
	}

	var s strings.Builder
	for c, coord := range coords {
		cs := strconv.Itoa(coord)
		s.WriteString(cs)
		// Leave off , at the end
		if c < len(coords)-1 {
			s.WriteString(",")
		}
	}
	return coordsOption{s.String()}
}

type coordsOptions struct {
	Rect   func(coords ...int) coordsOption
	Circle func(coords ...int) coordsOption
	Poly   func(coords ...int) coordsOption
}

/* Shape */
type shapeOption struct{ string }

func (o shapeOption) String() string { return o.string }

func shapeOptionRect() shapeOption {
	return shapeOption{"rect"}
}

func shapeOptionCircle() shapeOption {
	return shapeOption{"circle"}
}

func shapeOptionPoly() shapeOption {
	return shapeOption{"poly"}
}

type shapeOptions struct {
	Rect   func() shapeOption
	Circle func() shapeOption
	Poly   func() shapeOption
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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<audio {{global}}>{{innerhtml}}</audio>`)

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
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<img {{global}} />`)

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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<map {{global}}>{{innerhtml}}</map>`)

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
		"global": BuildGlobalProps(props.GlobalProps),
	}

	t := Mx(`<track {{global}} />`)

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
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<video {{global}}>{{innerhtml}}</video>`)

	s := Render(t, values)
	return s
}
