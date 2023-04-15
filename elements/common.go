package elements

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

/* Crossorigin */
type crossoriginOption struct{ string }

func (o crossoriginOption) String() string { return o.string }

func crossoriginOptionAnonymous() crossoriginOption {
	return crossoriginOption{"anonymous"}
}

func crossoriginOptionCredentials() crossoriginOption {
	return crossoriginOption{"use-credentials"}
}

type crossoriginOptions struct {
	Anonymous   func() crossoriginOption
	Credentials func() crossoriginOption
}

/* Loading */
type loadingOption struct{ string }

func (o loadingOption) String() string { return o.string }

func loadingOptionEager() loadingOption {
	return loadingOption{"eager"}
}

func loadingOptionLazy() loadingOption {
	return loadingOption{"lazy"}
}

type loadingOptions struct {
	Eager func() loadingOption
	Lazy  func() loadingOption
}
