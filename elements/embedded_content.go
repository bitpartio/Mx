package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#embedded_content

import . "github.com/bitpartio/Mx/utils"

func init() {
	IframeOptions = iframeOptions{
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
		Sandbox: sandboxOptions{
			Downlaods:                       sandboxOptionDownlaods,
			DownlaodsWithoutUserInteraction: sandboxOptionDownlaodsWithoutUserInteraction,
			Forms:                           sandboxOptionForms,
			Modals:                          sandboxOptionModals,
			OrientationLock:                 sandboxOptionOrientationLock,
			PointerLock:                     sandboxOptionPointerLock,
			Popups:                          sandboxOptionPopups,
			PopupsToEscapeSandbox:           sandboxOptionPopupsToEscapeSandbox,
			Presentation:                    sandboxOptionPresentation,
			SameOrigin:                      sandboxOptionSameOrigin,
			Scripts:                         sandboxOptionScripts,
			StorageAccessByUserActivation:   sandboxOptionStorageAccessByUserActivation,
			TopNavigation:                   sandboxOptionTopNavigation,
			TopNavigationByUserActivation:   sandboxOptionTopNavigationByUserActivation,
			TopNavigationToCustomProtocols:  sandboxOptionTopNavigationToCustomProtocols,
		},
	}
}

/*
 * Embeds external content at the specified point in the document. This
 * content is provided by an external application or other source of
 * interactive content such as a browser plug-in.
 */
type EmbedProps struct {
	GlobalProps

	Height *int
	Src    string
	Type   string
	Width  *int
}

func Embed(props EmbedProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"height": BuildIntProp("height", props.Height),
		"src":    BuildProp("src", props.Src),
		"type":   BuildProp("type", props.Type),
		"width":  BuildIntProp("width", props.Width),
	}

	m := BuildMarkup("embed", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents a nested browsing context, embedding another HTML page
 * into the current one.
 */
type IframeProps struct {
	GlobalProps

	Allow          string
	Credentialless bool
	Csp            string
	Height         *int
	Loading        func() loadingOption
	Name           string
	Referrerpolicy func() referrerpolicyOption
	Sandbox        []func() sandboxOption
	Src            string
	Srcdoc         string
	Width          *int

	InnerHTML string
}

func Iframe(props IframeProps) string {
	var loading string
	if props.Loading != nil {
		loading = props.Loading().String()
	}

	var referrerpolicy string
	if props.Referrerpolicy != nil {
		referrerpolicy = props.Referrerpolicy().String()
	}

	var sandbox string
	if len(props.Sandbox) > 0 {
		sandboxStrings := make([]string, len(props.Sandbox))
		for k, sandbox := range props.Sandbox {
			sandboxStrings[k] = sandbox().String()
		}

		sandbox = BuildPropListWithSpaces("rel", sandboxStrings)
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"allow":          BuildProp("allow", props.Allow),
		"credentialless": BuildBooleanProp("credentialless", props.Credentialless),
		"csp":            BuildProp("csp", props.Csp),
		"height":         BuildIntProp("height", props.Height),
		"loading":        loading,
		"name":           BuildProp("name", props.Name),
		"referrerpolicy": referrerpolicy,
		"sandbox":        sandbox,
		"src":            BuildProp("src", props.Src),
		"srcdoc":         BuildProp("srcdoc", props.Srcdoc),
		"width":          BuildIntProp("width", props.Width),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("iframe", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Sandbox */
type sandboxOption struct{ string }

func (o sandboxOption) String() string { return o.string }

func sandboxOptionDownlaods() sandboxOption {
	return sandboxOption{"allow-downloads"}
}

func sandboxOptionDownlaodsWithoutUserInteraction() sandboxOption {
	return sandboxOption{"allow-downloads-without-user-interaction"}
}

func sandboxOptionForms() sandboxOption {
	return sandboxOption{"allow-forms"}
}

func sandboxOptionModals() sandboxOption {
	return sandboxOption{"allow-modals"}
}

func sandboxOptionOrientationLock() sandboxOption {
	return sandboxOption{"allow-orientation-lock"}
}

func sandboxOptionPointerLock() sandboxOption {
	return sandboxOption{"allow-pointer-lock"}
}

func sandboxOptionPopups() sandboxOption {
	return sandboxOption{"allow-popups"}
}

func sandboxOptionPopupsToEscapeSandbox() sandboxOption {
	return sandboxOption{"allow-popups-to-escape-sandbox"}
}

func sandboxOptionPresentation() sandboxOption {
	return sandboxOption{"allow-presentation"}
}

func sandboxOptionSameOrigin() sandboxOption {
	return sandboxOption{"allow-same-origin"}
}

func sandboxOptionScripts() sandboxOption {
	return sandboxOption{"allow-scripts"}
}

func sandboxOptionStorageAccessByUserActivation() sandboxOption {
	return sandboxOption{"allow-storage-access-by-user-activation"}
}

func sandboxOptionTopNavigation() sandboxOption {
	return sandboxOption{"allow-top-navigation"}
}

func sandboxOptionTopNavigationByUserActivation() sandboxOption {
	return sandboxOption{"allow-top-navigation-by-custom-protocols"}
}

func sandboxOptionTopNavigationToCustomProtocols() sandboxOption {
	return sandboxOption{"allow-top-navigation-to-custom-protocols"}
}

type sandboxOptions struct {
	Downlaods                       func() sandboxOption
	DownlaodsWithoutUserInteraction func() sandboxOption
	Forms                           func() sandboxOption
	Modals                          func() sandboxOption
	OrientationLock                 func() sandboxOption
	PointerLock                     func() sandboxOption
	Popups                          func() sandboxOption
	PopupsToEscapeSandbox           func() sandboxOption
	Presentation                    func() sandboxOption
	SameOrigin                      func() sandboxOption
	Scripts                         func() sandboxOption
	StorageAccessByUserActivation   func() sandboxOption
	TopNavigation                   func() sandboxOption
	TopNavigationByUserActivation   func() sandboxOption
	TopNavigationToCustomProtocols  func() sandboxOption
}

type SandboxOptions []func() sandboxOption

type iframeOptions struct {
	Loading        loadingOptions
	Referrerpolicy referrerpolicyOptions
	Sandbox        sandboxOptions
}

var IframeOptions iframeOptions

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

	m := BuildMarkup("object", values)

	t := Mx(m)

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

	m := BuildMarkup("picture", values)

	t := Mx(m)

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

	m := BuildMarkup("source", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}
