package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#forms

import (
	"time"

	. "github.com/bitpartio/Mx/utils"
)

func init() {
	Input = inputTypes{
		Button:        InputButton,
		Checkbox:      InputCheckbox,
		Color:         InputColor,
		Date:          InputDate,
		DatetimeLocal: InputDatetimeLocal,
		Email:         InputEmail,
		File:          InputFile,
		Hidden:        InputHidden,
		Image:         InputImage,
		Month:         InputMonth,
		Number:        InputNumber,
		Password:      InputPassword,
		Radio:         InputRadio,
		Range:         InputRange,
		Reset:         InputReset,
		Search:        InputSearch,
		Submit:        InputSubmit,
		Tel:           InputTel,
		Text:          InputText,
		Time:          InputTime,
		Url:           InputUrl,
		Week:          InputWeek,
	}
	ButtonOptions = buttonOptions{
		Formenctype: formenctypeOptions{
			Url:       formenctypeOptionUrl,
			Multipart: formenctypeOptionMultipart,
			Text:      formenctypeOptionText,
		},
		Formmethod: formmethodOptions{
			Get:  formmethodOptionGet,
			Post: formmethodOptionPost,
		},
		Type: typeOptions{
			Button: typeOptionButton,
			Reset:  typeOptionReset,
			Submit: typeOptionSubmit,
		},
	}
	FormOptions = formOptions{
		Autocomplete: autocompleteFormOptions{
			On:  autocompleteFormOptionOn,
			Off: autocompleteFormOptionOff,
		},
		Method: methodOptions{
			Dialog: methodOptionDialog,
			Get:    methodOptionGet,
			Post:   methodOptionPost,
		},
		Rel: formRelOptions{
			External:   formRelOptionExternal,
			Help:       formRelOptionHelp,
			License:    formRelOptionLicense,
			Next:       formRelOptionNext,
			NoFollow:   formRelOptionNoFollow,
			NoOpener:   formRelOptionNoOpener,
			NoReferrer: formRelOptionNoReferrer,
			Opener:     formRelOptionOpener,
			Prev:       formRelOptionPrev,
			Search:     formRelOptionSearch,
		},
	}
	InputOptions = inputOptions{
		Autocomplete: autocompleteInputOptions{
			Off:                      autocompleteInputOptionOff,
			On:                       autocompleteInputOptionOn,
			Name:                     autocompleteInputOptionName,
			HonorificPrefix:          autocompleteInputOptionHonorificPrefix,
			GivenName:                autocompleteInputOptionGivenName,
			AdditionalName:           autocompleteInputOptionAdditionalName,
			FamilyName:               autocompleteInputOptionFamilyName,
			HonorificSuffix:          autocompleteInputOptionHonorificSuffix,
			Nickname:                 autocompleteInputOptionNickname,
			Email:                    autocompleteInputOptionEmail,
			Username:                 autocompleteInputOptionUsername,
			NewPassword:              autocompleteInputOptionNewPassword,
			CurrentPassword:          autocompleteInputOptionCurrentPassword,
			OneTimeCode:              autocompleteInputOptionOneTimeCode,
			OrganizationTitle:        autocompleteInputOptionOrganizationTitle,
			Organization:             autocompleteInputOptionOrganization,
			StreetAddress:            autocompleteInputOptionStreetAddress,
			AddressLine1:             autocompleteInputOptionAddressLine1,
			AddressLine2:             autocompleteInputOptionAddressLine2,
			AddressLine3:             autocompleteInputOptionAddressLine3,
			AddressLevel1:            autocompleteInputOptionAddressLevel1,
			AddressLevel2:            autocompleteInputOptionAddressLevel2,
			AddressLevel3:            autocompleteInputOptionAddressLevel3,
			AddressLevel4:            autocompleteInputOptionAddressLevel4,
			Country:                  autocompleteInputOptionCountry,
			CountryName:              autocompleteInputOptionCountryName,
			PostalCode:               autocompleteInputOptionPostalCode,
			CreditCardName:           autocompleteInputOptionCreditCardName,
			CreditCardGivenName:      autocompleteInputOptionCreditCardGivenName,
			CreditCardAdditionalName: autocompleteInputOptionCreditCardAdditionalName,
			CreditCardFamilyName:     autocompleteInputOptionCreditCardFamilyName,
			CreditCardNumber:         autocompleteInputOptionCreditCardNumber,
			CreditCardExp:            autocompleteInputOptionCreditCardExp,
			CreditCardExpMonth:       autocompleteInputOptionCreditCardExpMonth,
			CreditCardExpYear:        autocompleteInputOptionCreditCardExpYear,
			CreditCardSecurityCode:   autocompleteInputOptionCreditCardSecurityCode,
			CreditCardType:           autocompleteInputOptionCreditCardType,
			TransactionCurrency:      autocompleteInputOptionTransactionCurrency,
			TransactionAmount:        autocompleteInputOptionTransactionAmount,
			Language:                 autocompleteInputOptionLanguage,
			Birthday:                 autocompleteInputOptionBirthday,
			BirthdayDay:              autocompleteInputOptionBirthdayDay,
			BirthdayMonth:            autocompleteInputOptionBirthdayMonth,
			BirthdayYear:             autocompleteInputOptionBirthdayYear,
			Sex:                      autocompleteInputOptionSex,
			Telephone:                autocompleteInputOptionTelephone,
			TelephoneCountryCode:     autocompleteInputOptionTelephoneCountryCode,
			TelephoneNational:        autocompleteInputOptionTelephoneNational,
			TelephoneAreaCode:        autocompleteInputOptionTelephoneAreaCode,
			TelephoneLocal:           autocompleteInputOptionTelephoneLocal,
			TelephoneExtension:       autocompleteInputOptionTelephoneExtension,
			IMPP:                     autocompleteInputOptionIMPP,
			Url:                      autocompleteInputOptionUrl,
			Photo:                    autocompleteInputOptionPhoto,
		},
		Capture: captureOptions{
			Env:  captureOptionEnv,
			User: captureOptionUser,
		},
	}
}

/*
 * An interactive element activated by a user with a mouse, keyboard,
 * finger, voice command, or other assistive technology. Once activated,
 * it then performs an action, such as submitting a form or opening a
 * dialog.
 */
type ButtonProps struct {
	GlobalProps

	Autofocus      bool
	Disabled       bool
	Form           string
	Formaction     string
	Formenctype    func() formenctypeOption
	Formmethod     func() formmethodOption
	Formnovalidate bool
	Formtarget     string
	Name           string
	Type           func() typeOption
	Value          string

	InnerHTML string
}

func Button(props ButtonProps) string {
	var formenctype string
	if props.Formenctype != nil {
		formenctype = BuildProp("formenctype", props.Formenctype().String())
	}
	var formmethod string
	if props.Formmethod != nil {
		formmethod = BuildProp("formmethod", props.Formmethod().String())
	}
	var typeOf string
	if props.Type != nil {
		typeOf = BuildProp("type", props.Type().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"autofocus":      BuildBooleanProp("autofocus", props.Autofocus),
		"disabled":       BuildBooleanProp("disabled", props.Disabled),
		"form":           BuildProp("form", props.Form),
		"formaction":     BuildProp("formaction", props.Formaction),
		"formenctype":    formenctype,
		"formmethod":     formmethod,
		"formnovalidate": BuildBooleanProp("formnovalidate", props.Formnovalidate),
		"formtarget":     BuildProp("formtarget", props.Formtarget),
		"name":           BuildProp("name", props.Name),
		"type":           typeOf,
		"value":          BuildProp("value", props.Value),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("button", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

type buttonOptions struct {
	Formenctype formenctypeOptions
	Formmethod  formmethodOptions
	Type        typeOptions
}

var ButtonOptions buttonOptions

/* Formenctype */
type formenctypeOption struct{ string }

func (o formenctypeOption) String() string { return o.string }

func formenctypeOptionUrl() formenctypeOption {
	return formenctypeOption{"application/x-www-form-urlencoded"}
}

func formenctypeOptionMultipart() formenctypeOption {
	return formenctypeOption{"multipart/form-data"}
}

func formenctypeOptionText() formenctypeOption {
	return formenctypeOption{"text/plain"}
}

type formenctypeOptions struct {
	Url       func() formenctypeOption
	Multipart func() formenctypeOption
	Text      func() formenctypeOption
}

/* Formmethod */
type formmethodOption struct{ string }

func (o formmethodOption) String() string { return o.string }

func formmethodOptionGet() formmethodOption {
	return formmethodOption{"get"}
}

func formmethodOptionPost() formmethodOption {
	return formmethodOption{"post"}
}

type formmethodOptions struct {
	Get  func() formmethodOption
	Post func() formmethodOption
}

/* Type */
type typeOption struct{ string }

func (o typeOption) String() string { return o.string }

func typeOptionButton() typeOption {
	return typeOption{"button"}
}

func typeOptionReset() typeOption {
	return typeOption{"reset"}
}

func typeOptionSubmit() typeOption {
	return typeOption{"submit"}
}

type typeOptions struct {
	Button func() typeOption
	Reset  func() typeOption
	Submit func() typeOption
}

/*
 * Contains a set of <option> elements that represent the permissible or
 * recommended options available to choose from within other controls.
 */
type DatalistProps struct {
	GlobalProps

	InnerHTML string
}

func Datalist(props DatalistProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("datalist", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Used to group several controls as well as labels (<label>) within a web
 * form.
 */
type FieldsetProps struct {
	GlobalProps

	Disabled bool
	Form     string
	Name     string

	InnerHTML string
}

func Fieldset(props FieldsetProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"disabled": BuildBooleanProp("disabled", props.Disabled),
		"form":     BuildProp("form", props.Form),
		"name":     BuildProp("name", props.Name),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("fieldset", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents a document section containing interactive controls for
 * submitting information.
 */
type FormProps struct {
	GlobalProps

	AcceptCharset string
	Action        string
	Autocomplete  func() autocompleteFormOption
	Enctype       string
	Method        func() methodOption
	Novalidate    bool
	Name          string
	Rel           func() formRelOption
	Target        string

	InnerHTML string
}

func Form(props FormProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}
	var method string
	if props.Method != nil {
		method = BuildProp("method", props.Method().String())
	}
	var rel string
	if props.Rel != nil {
		rel = BuildProp("rel", props.Rel().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"accept-charset": BuildProp("accept-charset", props.AcceptCharset),
		"action":         BuildProp("action", props.Action),
		"autocomplete":   autocomplete,
		"enctype":        BuildProp("enctype", props.Enctype),
		"method":         method,
		"novalidate":     BuildBooleanProp("novalidate", props.Novalidate),
		"name":           BuildProp("name", props.Name),
		"rel":            rel,
		"target":         BuildProp("target", props.Target),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("form", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

type formOptions struct {
	Autocomplete autocompleteFormOptions
	Method       methodOptions
	Rel          formRelOptions
}

var FormOptions formOptions

/* Autocomplete */
type autocompleteFormOption struct{ string }

func (o autocompleteFormOption) String() string { return o.string }

func autocompleteFormOptionOff() autocompleteFormOption {
	return autocompleteFormOption{"off"}
}

func autocompleteFormOptionOn() autocompleteFormOption {
	return autocompleteFormOption{"on"}
}

type autocompleteFormOptions struct {
	Off func() autocompleteFormOption
	On  func() autocompleteFormOption
}

/* Method */
type methodOption struct{ string }

func (o methodOption) String() string { return o.string }

func methodOptionDialog() methodOption {
	return methodOption{"dialog"}
}

func methodOptionGet() methodOption {
	return methodOption{"get"}
}

func methodOptionPost() methodOption {
	return methodOption{"post"}
}

type methodOptions struct {
	Dialog func() methodOption
	Get    func() methodOption
	Post   func() methodOption
}

/* Rel */
type formRelOption struct{ string }

func (o formRelOption) String() string { return o.string }

func formRelOptionExternal() formRelOption {
	return formRelOption{"external"}
}
func formRelOptionHelp() formRelOption {
	return formRelOption{"help"}
}
func formRelOptionLicense() formRelOption {
	return formRelOption{"license"}
}
func formRelOptionNext() formRelOption {
	return formRelOption{"next"}
}
func formRelOptionNoFollow() formRelOption {
	return formRelOption{"nofollow"}
}
func formRelOptionNoOpener() formRelOption {
	return formRelOption{"noopener"}
}
func formRelOptionNoReferrer() formRelOption {
	return formRelOption{"noreferrer"}
}
func formRelOptionOpener() formRelOption {
	return formRelOption{"opener"}
}
func formRelOptionPrev() formRelOption {
	return formRelOption{"prev"}
}
func formRelOptionSearch() formRelOption {
	return formRelOption{"search"}
}

type formRelOptions struct {
	External   func() formRelOption
	Help       func() formRelOption
	License    func() formRelOption
	Next       func() formRelOption
	NoFollow   func() formRelOption
	NoOpener   func() formRelOption
	NoReferrer func() formRelOption
	Opener     func() formRelOption
	Prev       func() formRelOption
	Search     func() formRelOption
}

type FormRelOptions []func() formRelOption

/*
 * Used to create interactive controls for web-based forms in order to
 * accept data from the user; a wide variety of types of input data and
 * control widgets are available, depending on the device and user agent.
 * The <input> element is one of the most powerful and complex in all of
 * HTML due to the sheer number of combinations of input types and
 * attributes.
 */
type inputTypes struct {
	Button        func(props InputButtonProps) string
	Checkbox      func(props InputCheckboxProps) string
	Color         func(props InputColorProps) string
	Date          func(props InputDateProps) string
	DatetimeLocal func(props InputDatetimeLocalProps) string
	Email         func(props InputEmailProps) string
	File          func(props InputFileProps) string
	Hidden        func(props InputHiddenProps) string
	Image         func(props InputImageProps) string
	Month         func(props InputMonthProps) string
	Number        func(props InputNumberProps) string
	Password      func(props InputPasswordProps) string
	Radio         func(props InputRadioProps) string
	Range         func(props InputRangeProps) string
	Reset         func(props InputResetProps) string
	Search        func(props InputSearchProps) string
	Submit        func(props InputSubmitProps) string
	Tel           func(props InputTelProps) string
	Text          func(props InputTextProps) string
	Time          func(props InputTimeProps) string
	Url           func(props InputUrlProps) string
	Week          func(props InputWeekProps) string
}

var Input inputTypes

type inputOptions struct {
	Autocomplete autocompleteInputOptions
	Capture      captureOptions
}

var InputOptions inputOptions

/* Input Button */
type InputButtonProps struct {
	GlobalProps

	Disabled bool
	Form     string
	Name     string
	Value    string
}

func InputButton(props InputButtonProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "button"),

		"disabled": BuildBooleanProp("disabled", props.Disabled),
		"form":     BuildProp("form", props.Form),
		"name":     BuildProp("name", props.Name),
		"value":    BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Autocomplete */
type autocompleteInputOption struct{ string }

func (o autocompleteInputOption) String() string { return o.string }

func autocompleteInputOptionOff() autocompleteInputOption {
	return autocompleteInputOption{"off"}
}
func autocompleteInputOptionOn() autocompleteInputOption {
	return autocompleteInputOption{"on"}
}
func autocompleteInputOptionName() autocompleteInputOption {
	return autocompleteInputOption{"name"}
}
func autocompleteInputOptionHonorificPrefix() autocompleteInputOption {
	return autocompleteInputOption{"honorific-prefix"}
}
func autocompleteInputOptionGivenName() autocompleteInputOption {
	return autocompleteInputOption{"given-name"}
}
func autocompleteInputOptionAdditionalName() autocompleteInputOption {
	return autocompleteInputOption{"additional-name"}
}
func autocompleteInputOptionFamilyName() autocompleteInputOption {
	return autocompleteInputOption{"family-name"}
}
func autocompleteInputOptionHonorificSuffix() autocompleteInputOption {
	return autocompleteInputOption{"honorific-suffix"}
}
func autocompleteInputOptionNickname() autocompleteInputOption {
	return autocompleteInputOption{"nickname"}
}
func autocompleteInputOptionEmail() autocompleteInputOption {
	return autocompleteInputOption{"email"}
}
func autocompleteInputOptionUsername() autocompleteInputOption {
	return autocompleteInputOption{"username"}
}
func autocompleteInputOptionNewPassword() autocompleteInputOption {
	return autocompleteInputOption{"new-password"}
}
func autocompleteInputOptionCurrentPassword() autocompleteInputOption {
	return autocompleteInputOption{"current-password"}
}
func autocompleteInputOptionOneTimeCode() autocompleteInputOption {
	return autocompleteInputOption{"one-time-code"}
}
func autocompleteInputOptionOrganizationTitle() autocompleteInputOption {
	return autocompleteInputOption{"organization-title"}
}
func autocompleteInputOptionOrganization() autocompleteInputOption {
	return autocompleteInputOption{"organization"}
}
func autocompleteInputOptionStreetAddress() autocompleteInputOption {
	return autocompleteInputOption{"street-address"}
}
func autocompleteInputOptionAddressLine1() autocompleteInputOption {
	return autocompleteInputOption{"address-line-1"}
}
func autocompleteInputOptionAddressLine2() autocompleteInputOption {
	return autocompleteInputOption{"address-line-2"}
}
func autocompleteInputOptionAddressLine3() autocompleteInputOption {
	return autocompleteInputOption{"address-line-3"}
}
func autocompleteInputOptionAddressLevel1() autocompleteInputOption {
	return autocompleteInputOption{"address-level-1"}
}
func autocompleteInputOptionAddressLevel2() autocompleteInputOption {
	return autocompleteInputOption{"address-level-2"}
}
func autocompleteInputOptionAddressLevel3() autocompleteInputOption {
	return autocompleteInputOption{"address-level-3"}
}
func autocompleteInputOptionAddressLevel4() autocompleteInputOption {
	return autocompleteInputOption{"address-level-4"}
}
func autocompleteInputOptionCountry() autocompleteInputOption {
	return autocompleteInputOption{"country"}
}
func autocompleteInputOptionCountryName() autocompleteInputOption {
	return autocompleteInputOption{"country-name"}
}
func autocompleteInputOptionPostalCode() autocompleteInputOption {
	return autocompleteInputOption{"postal-code"}
}
func autocompleteInputOptionCreditCardName() autocompleteInputOption {
	return autocompleteInputOption{"cc-name"}
}
func autocompleteInputOptionCreditCardGivenName() autocompleteInputOption {
	return autocompleteInputOption{"cc-given-name"}
}
func autocompleteInputOptionCreditCardAdditionalName() autocompleteInputOption {
	return autocompleteInputOption{"cc-additional-name"}
}
func autocompleteInputOptionCreditCardFamilyName() autocompleteInputOption {
	return autocompleteInputOption{"cc-family-name"}
}
func autocompleteInputOptionCreditCardNumber() autocompleteInputOption {
	return autocompleteInputOption{"cc-number"}
}
func autocompleteInputOptionCreditCardExp() autocompleteInputOption {
	return autocompleteInputOption{"cc-exp"}
}
func autocompleteInputOptionCreditCardExpMonth() autocompleteInputOption {
	return autocompleteInputOption{"cc-exp-month"}
}
func autocompleteInputOptionCreditCardExpYear() autocompleteInputOption {
	return autocompleteInputOption{"cc-exp-year"}
}
func autocompleteInputOptionCreditCardSecurityCode() autocompleteInputOption {
	return autocompleteInputOption{"cc-csc"}
}
func autocompleteInputOptionCreditCardType() autocompleteInputOption {
	return autocompleteInputOption{"cc-type"}
}
func autocompleteInputOptionTransactionCurrency() autocompleteInputOption {
	return autocompleteInputOption{"transaction-currency"}
}
func autocompleteInputOptionTransactionAmount() autocompleteInputOption {
	return autocompleteInputOption{"transaction-amount"}
}
func autocompleteInputOptionLanguage() autocompleteInputOption {
	return autocompleteInputOption{"language"}
}
func autocompleteInputOptionBirthday() autocompleteInputOption {
	return autocompleteInputOption{"bday"}
}
func autocompleteInputOptionBirthdayDay() autocompleteInputOption {
	return autocompleteInputOption{"bday-day"}
}
func autocompleteInputOptionBirthdayMonth() autocompleteInputOption {
	return autocompleteInputOption{"bday-month"}
}
func autocompleteInputOptionBirthdayYear() autocompleteInputOption {
	return autocompleteInputOption{"bday-year"}
}
func autocompleteInputOptionSex() autocompleteInputOption {
	return autocompleteInputOption{"sex"}
}
func autocompleteInputOptionTelephone() autocompleteInputOption {
	return autocompleteInputOption{"tel"}
}
func autocompleteInputOptionTelephoneCountryCode() autocompleteInputOption {
	return autocompleteInputOption{"tel-country-code"}
}
func autocompleteInputOptionTelephoneNational() autocompleteInputOption {
	return autocompleteInputOption{"tel-national"}
}
func autocompleteInputOptionTelephoneAreaCode() autocompleteInputOption {
	return autocompleteInputOption{"tel-area-code"}
}
func autocompleteInputOptionTelephoneLocal() autocompleteInputOption {
	return autocompleteInputOption{"tel-local"}
}
func autocompleteInputOptionTelephoneExtension() autocompleteInputOption {
	return autocompleteInputOption{"tel-extension"}
}
func autocompleteInputOptionIMPP() autocompleteInputOption {
	return autocompleteInputOption{"impp"}
}
func autocompleteInputOptionUrl() autocompleteInputOption {
	return autocompleteInputOption{"url"}
}
func autocompleteInputOptionPhoto() autocompleteInputOption {
	return autocompleteInputOption{"photo"}
}

type autocompleteInputOptions struct {
	Off                      func() autocompleteInputOption
	On                       func() autocompleteInputOption
	Name                     func() autocompleteInputOption
	HonorificPrefix          func() autocompleteInputOption
	GivenName                func() autocompleteInputOption
	AdditionalName           func() autocompleteInputOption
	FamilyName               func() autocompleteInputOption
	HonorificSuffix          func() autocompleteInputOption
	Nickname                 func() autocompleteInputOption
	Email                    func() autocompleteInputOption
	Username                 func() autocompleteInputOption
	NewPassword              func() autocompleteInputOption
	CurrentPassword          func() autocompleteInputOption
	OneTimeCode              func() autocompleteInputOption
	OrganizationTitle        func() autocompleteInputOption
	Organization             func() autocompleteInputOption
	StreetAddress            func() autocompleteInputOption
	AddressLine1             func() autocompleteInputOption
	AddressLine2             func() autocompleteInputOption
	AddressLine3             func() autocompleteInputOption
	AddressLevel1            func() autocompleteInputOption
	AddressLevel2            func() autocompleteInputOption
	AddressLevel3            func() autocompleteInputOption
	AddressLevel4            func() autocompleteInputOption
	Country                  func() autocompleteInputOption
	CountryName              func() autocompleteInputOption
	PostalCode               func() autocompleteInputOption
	CreditCardName           func() autocompleteInputOption
	CreditCardGivenName      func() autocompleteInputOption
	CreditCardAdditionalName func() autocompleteInputOption
	CreditCardFamilyName     func() autocompleteInputOption
	CreditCardNumber         func() autocompleteInputOption
	CreditCardExp            func() autocompleteInputOption
	CreditCardExpMonth       func() autocompleteInputOption
	CreditCardExpYear        func() autocompleteInputOption
	CreditCardSecurityCode   func() autocompleteInputOption
	CreditCardType           func() autocompleteInputOption
	TransactionCurrency      func() autocompleteInputOption
	TransactionAmount        func() autocompleteInputOption
	Language                 func() autocompleteInputOption
	Birthday                 func() autocompleteInputOption
	BirthdayDay              func() autocompleteInputOption
	BirthdayMonth            func() autocompleteInputOption
	BirthdayYear             func() autocompleteInputOption
	Sex                      func() autocompleteInputOption
	Telephone                func() autocompleteInputOption
	TelephoneCountryCode     func() autocompleteInputOption
	TelephoneNational        func() autocompleteInputOption
	TelephoneAreaCode        func() autocompleteInputOption
	TelephoneLocal           func() autocompleteInputOption
	TelephoneExtension       func() autocompleteInputOption
	IMPP                     func() autocompleteInputOption
	Url                      func() autocompleteInputOption
	Photo                    func() autocompleteInputOption
}

/* Input Checkbox */
type InputCheckboxProps struct {
	GlobalProps

	Checked  bool
	Disabled bool
	Form     string
	Name     string
	Value    string
}

func InputCheckbox(props InputCheckboxProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "checkbox"),

		"checked":  BuildBooleanProp("checked", props.Checked),
		"disabled": BuildBooleanProp("disabled", props.Disabled),
		"form":     BuildProp("form", props.Form),
		"name":     BuildProp("name", props.Name),
		"value":    BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Color */
type InputColorProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Name         string
	Value        string
}

func InputColor(props InputColorProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "color"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"name":         BuildProp("name", props.Name),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Date */
type InputDateProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Max          time.Time
	Min          time.Time
	Name         string
	Readonly     bool
	Required     bool
	Step         *int
	Value        string
}

func InputDate(props InputDateProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "date"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"max":          BuildDateProp("max", props.Max),
		"min":          BuildDateProp("min", props.Min),
		"name":         BuildProp("name", props.Name),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"step":         BuildIntProp("step", props.Step),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input DatetimeLocal */
type InputDatetimeLocalProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Max          time.Time
	Min          time.Time
	Name         string
	Readonly     bool
	Required     bool
	Step         *int
	Value        string
}

func InputDatetimeLocal(props InputDatetimeLocalProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "datetime-local"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"max":          BuildDateTimeProp("max", props.Max),
		"min":          BuildDateTimeProp("max", props.Min),
		"name":         BuildProp("name", props.Name),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"step":         BuildIntProp("step", props.Step),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Email */
type InputEmailProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Maxlength    *int
	Minlength    *int
	Multiple     bool
	Name         string
	Pattern      string
	Placeholder  string
	Readonly     bool
	Required     bool
	Size         *int
	Value        string
}

func InputEmail(props InputEmailProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "email"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"maxlength":    BuildIntProp("maxlength", props.Maxlength),
		"minlength":    BuildIntProp("minlength", props.Minlength),
		"multiple":     BuildBooleanProp("multiple", props.Multiple),
		"name":         BuildProp("name", props.Name),
		"pattern":      BuildProp("pattern", props.Pattern),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"size":         BuildIntProp("size", props.Size),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input File */
type InputFileProps struct {
	GlobalProps

	Accept       []string
	Autocomplete func() autocompleteInputOption
	Capture      func() captureOption
	Disabled     bool
	Form         string
	List         string
	Multiple     bool
	Name         string
	Readonly     bool
	Required     bool
	Value        string
}

func InputFile(props InputFileProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	var capture string
	if props.Capture != nil {
		capture = BuildProp("capture", props.Capture().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "file"),

		"accept":       BuildPropListWithCommas("accept", props.Accept),
		"autocomplete": autocomplete,
		"capture":      capture,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"multiple":     BuildBooleanProp("multiple", props.Multiple),
		"name":         BuildProp("name", props.Name),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Capture */
type captureOption struct{ string }

func (o captureOption) String() string { return o.string }

func captureOptionEnv() captureOption {
	return captureOption{"environment"}
}

func captureOptionUser() captureOption {
	return captureOption{"user"}
}

type captureOptions struct {
	Env  func() captureOption
	User func() captureOption
}

/* Input Hidden */
type InputHiddenProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	Name         string
	Value        string
}

func InputHidden(props InputHiddenProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "hidden"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"name":         BuildProp("name", props.Form),
		"value":        BuildProp("value", props.Form),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Image */
type InputImageProps struct {
	GlobalProps

	Alt            string
	Autocomplete   func() autocompleteInputOption
	Disabled       bool
	Form           string
	Formaction     string
	Formenctype    func() formenctypeOption
	Formmethod     func() formmethodOption
	Formnovalidate bool
	Formtarget     string
	Height         *int
	List           string
	Name           string
	Readonly       bool
	Required       bool
	Src            string
	Width          *int
}

func InputImage(props InputImageProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	var formenctype string
	if props.Formenctype != nil {
		formenctype = BuildProp("formenctype", props.Formenctype().String())
	}

	var formmethod string
	if props.Formmethod != nil {
		formmethod = BuildProp("formmethod", props.Formmethod().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "image"),

		"alt":            BuildProp("alt", props.Alt),
		"autocomplete":   autocomplete,
		"disabled":       BuildBooleanProp("disabled", props.Disabled),
		"form":           BuildProp("form", props.Form),
		"formaction":     BuildProp("formaction", props.Form),
		"formenctype":    formenctype,
		"formmethod":     formmethod,
		"formnovalidate": BuildBooleanProp("formnovalidate", props.Formnovalidate),
		"formtarget":     BuildProp("formtarget", props.Formtarget),
		"height":         BuildIntProp("height", props.Height),
		"list":           BuildProp("list", props.List),
		"name":           BuildProp("name", props.Name),
		"readonly":       BuildBooleanProp("readonly", props.Readonly),
		"required":       BuildBooleanProp("required", props.Required),
		"src":            BuildProp("src", props.Src),
		"width":          BuildIntProp("width", props.Width),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Month */
type InputMonthProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Max          time.Time
	Min          time.Time
	Name         string
	Readonly     bool
	Required     bool
	Step         *int
	Value        string
}

func InputMonth(props InputMonthProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "month"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"max":          BuildDateMonthProp("max", props.Max),
		"min":          BuildDateMonthProp("max", props.Min),
		"name":         BuildProp("name", props.Name),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"step":         BuildIntProp("step", props.Step),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Number */
type InputNumberProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Max          *float64
	Min          *float64
	Name         string
	Placeholder  string
	Readonly     bool
	Required     bool
	Step         *float64
	Value        *float64
}

func InputNumber(props InputNumberProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "number"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"max":          BuildFloatProp("max", props.Max),
		"min":          BuildFloatProp("min", props.Min),
		"name":         BuildProp("name", props.Name),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"step":         BuildFloatProp("step", props.Step),
		"value":        BuildFloatProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Password */
type InputPasswordProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	File         string
	Maxlength    *int
	Minlength    *int
	Name         string
	Pattern      string
	Placeholder  string
	Readonly     bool
	Required     bool
	Size         *int
	Value        string
}

func InputPassword(props InputPasswordProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "password"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"file":         BuildProp("file", props.File),
		"maxlength":    BuildIntProp("maxlength", props.Maxlength),
		"minlength":    BuildIntProp("minlength", props.Minlength),
		"name":         BuildProp("name", props.Name),
		"pattern":      BuildProp("pattern", props.Pattern),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"size":         BuildIntProp("size", props.Size),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Radio */
type InputRadioProps struct {
	GlobalProps

	Checked  bool
	Disabled bool
	Form     string
	Name     string
	Required bool
	Value    string
}

func InputRadio(props InputRadioProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "radio"),

		"checked":  BuildBooleanProp("checked", props.Checked),
		"disabled": BuildBooleanProp("disabled", props.Disabled),
		"form":     BuildProp("form", props.Form),
		"name":     BuildProp("name", props.Name),
		"required": BuildBooleanProp("required", props.Required),
		"value":    BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Range */
type InputRangeProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Max          *float64
	Min          *float64
	Name         string
	Step         *float64
	Value        *float64
}

func InputRange(props InputRangeProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "range"),

		"autocomplate": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"max":          BuildFloatProp("max", props.Max),
		"min":          BuildFloatProp("min", props.Min),
		"name":         BuildProp("name", props.Name),
		"step":         BuildFloatProp("step", props.Step),
		"value":        BuildFloatProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Reset */
type InputResetProps struct {
	GlobalProps

	Disabled bool
	Form     string
	Name     string
	Value    string
}

func InputReset(props InputResetProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "reset"),

		"disabled": BuildBooleanProp("disabled", props.Disabled),
		"form":     BuildProp("form", props.Form),
		"name":     BuildProp("name", props.Name),
		"value":    BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Search */
type InputSearchProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Dirname      string
	Disabled     bool
	Form         string
	List         string
	Maxlength    *int
	Minlength    *int
	Name         string
	Pattern      string
	Placeholder  string
	Readonly     bool
	Required     bool
	Size         *int
	Value        string
}

func InputSearch(props InputSearchProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "search"),

		"autocomplete": autocomplete,
		"dirname":      BuildProp("dirname", props.Dirname),
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"maxlength":    BuildIntProp("maxlength", props.Maxlength),
		"minlength":    BuildIntProp("minlength", props.Minlength),
		"name":         BuildProp("name", props.Name),
		"pattern":      BuildProp("pattern", props.Pattern),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"size":         BuildIntProp("size", props.Size),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Submit */
type InputSubmitProps struct {
	GlobalProps

	Disabled       bool
	Form           string
	Formaction     string
	Formenctype    func() formenctypeOption
	Formmethod     func() formmethodOption
	Formnovalidate bool
	Formtarget     string
	Name           string
	Value          string
}

func InputSubmit(props InputSubmitProps) string {
	var formenctype string
	if props.Formenctype != nil {
		formenctype = BuildProp("formenctype", props.Formenctype().String())
	}

	var formmethod string
	if props.Formmethod != nil {
		formmethod = BuildProp("formmethod", props.Formmethod().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "submit"),

		"disabled":       BuildBooleanProp("disabled", props.Disabled),
		"form":           BuildProp("form", props.Form),
		"formaction":     BuildProp("formaction", props.Formaction),
		"formenctype":    formenctype,
		"formmethod":     formmethod,
		"formnovalidate": BuildBooleanProp("formnovalidate", props.Formnovalidate),
		"formtarget":     BuildProp("formtarget", props.Formtarget),
		"name":           BuildProp("name", props.Name),
		"value":          BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Tel */
type InputTelProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Maxlength    *int
	Minlength    *int
	Name         string
	Pattern      string
	Placeholder  string
	Readonly     bool
	Required     bool
	Size         *int
	Value        string
}

func InputTel(props InputTelProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "tel"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"maxlength":    BuildIntProp("maxlength", props.Maxlength),
		"minlength":    BuildIntProp("minlength", props.Minlength),
		"name":         BuildProp("name", props.Name),
		"pattern":      BuildProp("pattern", props.Pattern),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"size":         BuildIntProp("size", props.Size),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Text */
type InputTextProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Dirname      string
	Disabled     bool
	Form         string
	List         string
	Maxlength    *int
	Minlength    *int
	Name         string
	Pattern      string
	Placeholder  string
	Readonly     bool
	Required     bool
	Size         *int
	Value        string
}

func InputText(props InputTextProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "text"),

		"autocomplete": autocomplete,
		"dirname":      BuildProp("dirname", props.Dirname),
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"maxlength":    BuildIntProp("maxlength", props.Maxlength),
		"minlength":    BuildIntProp("minlength", props.Minlength),
		"name":         BuildProp("name", props.Name),
		"pattern":      BuildProp("pattern", props.Pattern),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"size":         BuildIntProp("size", props.Size),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Time */
type InputTimeProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Max          time.Time
	Min          time.Time
	Name         string
	Readonly     bool
	Required     bool
	Step         *int
	Value        time.Time
}

func InputTime(props InputTimeProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "time"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"max":          BuildTimeProp("max", props.Max),
		"min":          BuildTimeProp("min", props.Min),
		"name":         BuildProp("name", props.Name),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"step":         BuildIntProp("step", props.Step),
		"value":        BuildTimeProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Url */
type InputUrlProps struct {
	GlobalProps

	Autocomplete func() autocompleteInputOption
	Disabled     bool
	Form         string
	List         string
	Maxlength    *int
	Minlength    *int
	Name         string
	Pattern      string
	Placeholder  string
	Readonly     bool
	Required     bool
	Size         *int
	Value        string
}

func InputUrl(props InputUrlProps) string {
	var autocomplete string
	if props.Autocomplete != nil {
		autocomplete = BuildProp("autocomplete", props.Autocomplete().String())
	}

	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "url"),

		"autocomplete": autocomplete,
		"disabled":     BuildBooleanProp("disabled", props.Disabled),
		"form":         BuildProp("form", props.Form),
		"list":         BuildProp("list", props.List),
		"maxlength":    BuildIntProp("maxlength", props.Maxlength),
		"minlength":    BuildIntProp("minlength", props.Minlength),
		"name":         BuildProp("name", props.Name),
		"pattern":      BuildProp("pattern", props.Pattern),
		"placeholder":  BuildProp("placeholder", props.Placeholder),
		"readonly":     BuildBooleanProp("readonly", props.Readonly),
		"required":     BuildBooleanProp("required", props.Required),
		"size":         BuildIntProp("size", props.Size),
		"value":        BuildProp("value", props.Value),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/* Input Week */
type InputWeekProps struct {
	GlobalProps
}

func InputWeek(props InputWeekProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"type": BuildProp("type", "week"),
	}

	m := BuildMarkup("input", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents a caption for an item in a user interface.
 */
type LabelProps struct {
	GlobalProps

	InnerHTML string
}

func Label(props LabelProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("label", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents a caption for the content of its parent <fieldset>.
 */
type LegendProps struct {
	GlobalProps

	InnerHTML string
}

func Legend(props LegendProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("legend", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents either a scalar value within a known range or
 * a fractional value.
 */
type MeterProps struct {
	GlobalProps

	InnerHTML string
}

func Meter(props MeterProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("meter", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Creates a grouping of options within a <select> element.
 */
type OptgroupProps struct {
	GlobalProps

	InnerHTML string
}

func Optgroup(props OptgroupProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("optgroup", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Used to define an item contained in a select, an <optgroup>, or a
 * <datalist> element. As such, <option> can represent menu items in
 * popups and other lists of items in an HTML document.
 */
type OptionProps struct {
	GlobalProps

	InnerHTML string
}

func Option(props OptionProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("option", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Container element into which a site or app can inject the results
 * of a calculation or the outcome of a user action.
 */
type OutputProps struct {
	GlobalProps

	InnerHTML string
}

func Output(props OutputProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("output", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Displays an indicator showing the completion progress of a task,
 * typically displayed as a progress bar.
 */
type ProgressProps struct {
	GlobalProps

	InnerHTML string
}

func Progress(props ProgressProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("progress", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents a control that provides a menu of options.
 */
type SelectProps struct {
	GlobalProps

	InnerHTML string
}

func Select(props SelectProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("select", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}

/*
 * Represents a multi-line plain-text editing control, useful when you
 * want to allow users to enter a sizeable amount of free-form text, for
 * example a comment on a review or feedback form.
 */
type TextareaProps struct {
	GlobalProps

	InnerHTML string
}

func Textarea(props TextareaProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	m := BuildMarkup("textarea", values)

	t := Mx(m)

	s := Render(t, values)
	return s
}
