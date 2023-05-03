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
	ImgOptions = imgOptions{
		Crossorigin: crossoriginOptions{
			Anonymous:   crossoriginOptionAnonymous,
			Credentials: crossoriginOptionCredentials,
		},
		Decoding: decodingOptions{
			Sync:  decodingOptionSync,
			Async: decodingOptionAsync,
			Auto:  decodingOptionAuto,
		},
		Fetchpriority: fetchpriorityOptions{
			High: fetchpriorityOptionHigh,
			Low:  fetchpriorityOptionLow,
			Auto: fetchpriorityOptionAuto,
		},
		Loading: loadingOptions{
			Eager: loadingOptionEager,
			Lazy:  loadingOptionLazy,
		},
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
	}
	TrackOptions = trackOptions{
		Kind: kindOptions{
			Subtitles:    kindOptionSubtitles,
			Captions:     kindOptionCaptions,
			Descriptions: kindOptionDescriptions,
			Chapters:     kindOptionChapters,
			Metadata:     kindOptionMetadata,
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
	Coords         func() coordsOption
	Download       string
	Href           string
	Hreflang       string // Limited values but too complex for enum. Ref: https://datatracker.ietf.org/doc/html/rfc5646
	Ping           []string
	Referrerpolicy func() referrerpolicyOption
	Shape          func() shapeOption
	Target         string
}

func Area(props AreaProps) string {
	var coords string
	if props.Coords != nil {
		coords = BuildProp("coords", props.Coords().String())
	}
	var referrerpolicy string
	if props.Referrerpolicy != nil {
		referrerpolicy = BuildProp("referrerpolicy", props.Referrerpolicy().String())
	}
	var shape string
	if props.Shape != nil {
		referrerpolicy = BuildProp("shape", props.Shape().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"alt":            BuildProp("alt", props.Alt),
		"download":       BuildProp("download", props.Download),
		"coords":         coords,
		"href":           BuildProp("href", props.Href),
		"hreflang":       BuildProp("hreflang", props.Hreflang),
		"ping":           BuildPropListWithSpaces("ping", props.Ping),
		"referrerpolicy": referrerpolicy,
		"shape":          shape,
		"target":         BuildProp("target", props.Target),
	}

	m := BuildMarkup("area", values)

	t := Mx(m)

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
	Crossorigin           func() crossoriginOption
	Disableremoteplayback bool
	Loop                  bool
	Muted                 bool
	Preload               func() preloadOption
	Src                   string

	InnerHTML string
}

func Audio(props AudioProps) string {
	var controlslist string
	if props.Controlslist != nil {
		controlslist = BuildProp("shape", props.Controlslist().String())
	}
	var crossorigin string
	if props.Crossorigin != nil {
		crossorigin = BuildProp("crossorigin", props.Crossorigin().String())
	}
	var preload string
	if props.Preload != nil {
		preload = BuildProp("preload", props.Preload().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"autoplay":              BuildBooleanProp("autoplay", props.Autoplay),
		"controls":              BuildBooleanProp("controls", props.Controls),
		"controlslist":          controlslist,
		"crossorigin":           crossorigin,
		"disableremoteplayback": BuildBooleanProp("disableremoteplayback", props.Disableremoteplayback),
		"loop":                  BuildBooleanProp("loop", props.Loop),
		"muted":                 BuildBooleanProp("muted", props.Muted),
		"preload":               preload,
		"src":                   BuildProp("src", props.Src),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("audio", values)

	t := Mx(m)

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

/* Audio Options */
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
	Alt            string
	Crossorigin    func() crossoriginOption
	Decoding       func() decodingOption
	Elementtiming  string
	Fetchpriority  func() fetchpriorityOption
	Height         *int
	Ismap          bool
	Loading        func() loadingOption
	Referrerpolicy func() referrerpolicyOption
	Sizes          []string
	Src            string
	Srcset         []string
	Width          *int
	Usemap         string

	GlobalProps
}

func Img(props ImgProps) string {
	var crossorigin string
	if props.Crossorigin != nil {
		crossorigin = BuildProp("crossorigin", props.Crossorigin().String())
	}
	var decoding string
	if props.Decoding != nil {
		decoding = BuildProp("decoding", props.Decoding().String())
	}
	var fetchpriority string
	if props.Fetchpriority != nil {
		fetchpriority = BuildProp("fetchpriority", props.Fetchpriority().String())
	}
	var loading string
	if props.Loading != nil {
		loading = BuildProp("loading", props.Loading().String())
	}
	var referrerpolicy string
	if props.Referrerpolicy != nil {
		referrerpolicy = BuildProp("referrerpolicy", props.Referrerpolicy().String())
	}

	values := map[string]interface{}{
		"global":         BuildGlobalProps(props.GlobalProps),
		"alt":            BuildProp("alt", props.Alt),
		"crossorigin":    crossorigin,
		"decoding":       decoding,
		"elementtiming":  BuildProp("elementtiming", props.Elementtiming),
		"fetchpriority":  fetchpriority,
		"height":         BuildIntProp("height", props.Height),
		"ismap":          BuildBooleanProp("ismap", props.Ismap),
		"loading":        loading,
		"referrerpolicy": referrerpolicy,
		"sizes":          BuildPropListWithCommas("referrerpolicy", props.Sizes),
		"src":            BuildProp("src", props.Src),
		"srcset":         BuildPropListWithCommas("srcset", props.Srcset),
		"width":          BuildIntProp("width", props.Width),
		"usemap":         BuildProp("usemap", props.Usemap),
	}

	m := BuildMarkup("img", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* decoding */
type decodingOption struct{ string }

func (o decodingOption) String() string { return o.string }

func decodingOptionSync() decodingOption {
	return decodingOption{"sync"}
}

func decodingOptionAsync() decodingOption {
	return decodingOption{"async"}
}

func decodingOptionAuto() decodingOption {
	return decodingOption{"auto"}
}

type decodingOptions struct {
	Sync  func() decodingOption
	Async func() decodingOption
	Auto  func() decodingOption
}

/* fetchpriority */
type fetchpriorityOption struct{ string }

func (o fetchpriorityOption) String() string { return o.string }

func fetchpriorityOptionHigh() fetchpriorityOption {
	return fetchpriorityOption{"high"}
}

func fetchpriorityOptionLow() fetchpriorityOption {
	return fetchpriorityOption{"low"}
}

func fetchpriorityOptionAuto() fetchpriorityOption {
	return fetchpriorityOption{"auto"}
}

type fetchpriorityOptions struct {
	High func() fetchpriorityOption
	Low  func() fetchpriorityOption
	Auto func() fetchpriorityOption
}

/* Img Options */
type imgOptions struct {
	Crossorigin    crossoriginOptions
	Decoding       decodingOptions
	Fetchpriority  fetchpriorityOptions
	Loading        loadingOptions
	Referrerpolicy referrerpolicyOptions
}

var ImgOptions imgOptions

/*
 * used with <area> elements to define an image map (a clickable link area).
 */
type mapProps struct {
	GlobalProps

	Name string

	InnerHTML string
}

func Map(props mapProps) string {
	values := map[string]interface{}{
		"global": buildGlobalProps(props.GlobalProps),

		"name": BuildProp("name", props.Name),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("map", values)

	t := Mx(m)

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

	Default bool
	Kind    func() kindOption
	Label   string
	Src     string
	Srclang string
}

func Track(props TrackProps) string {
	var kind string
	if props.Kind != nil {
		kind = BuildProp("kind", props.Kind().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"kind":    kind,
		"label":   BuildProp("label", props.Label),
		"src":     BuildProp("src", props.Src),
		"srclang": BuildProp("srclang", props.Srclang),
	}

	m := BuildMarkup("track", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* kind */
type kindOption struct{ string }

func (o kindOption) String() string { return o.string }

func kindOptionSubtitles() kindOption {
	return kindOption{"subtitles"}
}

func kindOptionCaptions() kindOption {
	return kindOption{"captions"}
}

func kindOptionDescriptions() kindOption {
	return kindOption{"descriptions"}
}

func kindOptionChapters() kindOption {
	return kindOption{"chapters"}
}

func kindOptionMetadata() kindOption {
	return kindOption{"metadata"}
}

type kindOptions struct {
	Subtitles    func() kindOption
	Captions     func() kindOption
	Descriptions func() kindOption
	Chapters     func() kindOption
	Metadata     func() kindOption
}

/* Track Options */
type trackOptions struct {
	Kind kindOptions
}

var TrackOptions trackOptions

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

	m := BuildMarkup("video", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}
