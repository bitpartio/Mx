package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#image_and_multimedia

import (
	"strconv"
	"strings"

	. "github.com/bitpartio/Mx/utils"
)

func init() {
	AreaOptions = areaOptions{
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
		Rel: areaRelOptions{
			Alternate:  areaRelOptionAlternate,
			Author:     areaRelOptionAuthor,
			Bookmark:   areaRelOptionBookmark,
			Help:       areaRelOptionHelp,
			License:    areaRelOptionLicense,
			Next:       areaRelOptionNext,
			NoFollow:   areaRelOptionNoFollow,
			NoReferrer: areaRelOptionNoReferrer,
			Prefetch:   areaRelOptionPrefetch,
			Prev:       areaRelOptionPrev,
			Search:     areaRelOptionSearch,
			Tag:        areaRelOptionTag,
		},
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
	AudioOptions = audioOptions{
		Controlslist: controlslistOptions{
			NoDownload:       controlslistOptionNodownload,
			NoFullscreen:     controlslistOptionNofullscreen,
			NoRemotePlayback: controlslistOptionNoremoteplayback,
		},
		Crossorigin: crossoriginOptions{
			Anonymous:   crossoriginOptionAnonymous,
			Credentials: crossoriginOptionCredentials,
		},
		Preload: preloadOptions{
			None:     preloadOptionNone,
			Metadata: preloadOptionMetadata,
			Auto:     preloadOptionAuto,
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

	Alt            string
	Coords         coordsOption
	Download       string
	Href           string
	Hreflang       string // Limited values but too complex for enum. Ref: https://datatracker.ietf.org/doc/html/rfc5646
	Ping           []string
	Referrerpolicy func() referrerpolicyOption
	Shape          func() shapeOption
	Target         string
}

func Area(props AreaProps) string {
	var referrerpolicy string
	if props.Referrerpolicy != nil {
		referrerpolicy = BuildProp("referrerpolicy", props.Referrerpolicy().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"alt":            BuildProp("alt", props.Alt),
		"download":       BuildProp("download", props.Download),
		"coords":         BuildProp("coords", props.Coords.String()),
		"href":           BuildProp("href", props.Href),
		"hreflang":       BuildProp("hreflang", props.Hreflang),
		"ping":           BuildPropList("ping", props.Ping),
		"referrerpolicy": referrerpolicy,
		"shape":          BuildProp("shape", props.Shape().String()),
		"target":         BuildProp("target", props.Target),
	}

	t := Mx(`<area {{global}} {{alt}} {{coords}} {{href}} {{hreflang}} {{ping}} {{shape}} {{target}}/>`)

	s := Render(t, values)
	return s
}

type areaOptions struct {
	Coords         coordsOptions
	Referrerpolicy referrerpolicyOptions
	Rel            areaRelOptions
	Shape          shapeOptions
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

/* Rel */
type areaRelOption struct{ string }

func (o areaRelOption) String() string { return o.string }

func areaRelOptionAlternate() areaRelOption {
	return areaRelOption{"alternate"}
}

func areaRelOptionAuthor() areaRelOption {
	return areaRelOption{"author"}
}

func areaRelOptionBookmark() areaRelOption {
	return areaRelOption{"bookmark"}
}
func areaRelOptionHelp() areaRelOption {
	return areaRelOption{"help"}
}
func areaRelOptionLicense() areaRelOption {
	return areaRelOption{"license"}
}
func areaRelOptionNext() areaRelOption {
	return areaRelOption{"next"}
}
func areaRelOptionNoFollow() areaRelOption {
	return areaRelOption{"nofollow"}
}
func areaRelOptionNoReferrer() areaRelOption {
	return areaRelOption{"noreferrer"}
}
func areaRelOptionPrefetch() areaRelOption {
	return areaRelOption{"prefetch"}
}
func areaRelOptionPrev() areaRelOption {
	return areaRelOption{"prev"}
}
func areaRelOptionSearch() areaRelOption {
	return areaRelOption{"search"}
}
func areaRelOptionTag() areaRelOption {
	return areaRelOption{"tag"}
}

type areaRelOptions struct {
	Alternate  func() areaRelOption
	Author     func() areaRelOption
	Bookmark   func() areaRelOption
	Help       func() areaRelOption
	License    func() areaRelOption
	Next       func() areaRelOption
	NoFollow   func() areaRelOption
	NoReferrer func() areaRelOption
	Prev       func() areaRelOption
	Prefetch   func() areaRelOption
	Search     func() areaRelOption
	Tag        func() areaRelOption
}

type AreaRelOptions []func() areaRelOption

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

	Autoplay              bool
	Controls              bool
	Controlslist          func() controlslistOption
	Disableremoteplayback bool
	Loop                  bool
	Muted                 bool
	Preload               func() preloadOption
	Src                   string

	InnerHTML string
}

func Audio(props AudioProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"autoplay":              BuildBooleanProp("autoplay", props.Autoplay),
		"controls":              BuildBooleanProp("controls", props.Controls),
		"controlslist":          BuildProp("controlslist", props.Controlslist().String()),
		"disableremoteplayback": BuildBooleanProp("disableremoteplayback", props.Disableremoteplayback),
		"loop":                  BuildBooleanProp("loop", props.Loop),
		"muted":                 BuildBooleanProp("muted", props.Muted),
		"preload":               BuildProp("preload", props.Preload().String()),
		"src":                   BuildProp("src", props.Src),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<audio {{global}} {{autoplay}} {{controls}} {{controlslist}} {{disableremoteplayback}} {{loop}} {{muted}} {{preload}} {{src}}>{{innerhtml}}</audio>`)

	s := Render(t, values)
	return s
}

/* Controlslist */
type controlslistOption struct{ string }

func (o controlslistOption) String() string { return o.string }

func controlslistOptionNodownload() controlslistOption {
	return controlslistOption{"nodownload"}
}

func controlslistOptionNofullscreen() controlslistOption {
	return controlslistOption{"nofullscreen"}
}

func controlslistOptionNoremoteplayback() controlslistOption {
	return controlslistOption{"noremoteplayback"}
}

type controlslistOptions struct {
	NoDownload       func() controlslistOption
	NoFullscreen     func() controlslistOption
	NoRemotePlayback func() controlslistOption
}

/* Preload */
type preloadOption struct{ string }

func (o preloadOption) String() string { return o.string }

func preloadOptionNone() preloadOption {
	return preloadOption{"none"}
}

func preloadOptionMetadata() preloadOption {
	return preloadOption{"metadata"}
}

func preloadOptionAuto() preloadOption {
	return preloadOption{"auto"}
}

type preloadOptions struct {
	None     func() preloadOption
	Metadata func() preloadOption
	Auto     func() preloadOption
}

// * Audio Options */
type audioOptions struct {
	Controlslist controlslistOptions
	Crossorigin  crossoriginOptions
	Preload      preloadOptions
}

var AudioOptions audioOptions

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
