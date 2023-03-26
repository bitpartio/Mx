package elements

import . "github.com/bitpartio/Mx/utils"

// Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#web_components

/*
 * Part of the Web Components technology suite, this element is a placeholder
 * inside a web component that you can fill with your own markup, which lets
 * you create separate DOM trees and present them together.
 */
type SlotProps struct {
	GlobalProps

	Name string

	InnerHTML string
}

func Slot(props SlotProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"name": BuildProp("name", props.Name),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<slot {{name}} {{global}}>{{innerhtml}}</slot>`)

	s := Render(t, values)
	return s
}

/*
 * A mechanism for holding HTML that is not to be rendered immediately when a
 * page is loaded but may be instantiated subsequently during runtime using
 * JavaScript.
 */
type TemplateProps struct {
	GlobalProps

	InnerHTML string
}

func Template(props TemplateProps) string {
	values := map[string]interface{}{
		"global": BuildGlobalProps(props.GlobalProps),

		"innerhtml": props.InnerHTML,
	}

	t := Mx(`<template {{global}}>{{innerhtml}}</template>`)

	s := Render(t, values)
	return s
}
