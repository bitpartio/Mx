package main

import (
	"testing"

	. "github.com/bitpartio/gomx/elements"
	. "github.com/bitpartio/gomx/utils"
)

func BenchmarkRender(b *testing.B) {
	for n := 0; n < b.N; n++ {
		render()
	}
}

func render() {
	button := Button(
		ButtonProps{
			GlobalProps: GlobalProps{
				Autocapitalize: Props.Autocapitalize.Words,
				Dir:            Props.Dir.Auto,
				Data:           DataValues{"current": "", "second": "four"},
				Aria:           AriaRoles{"current": "page"},
				Htmx:           HtmxProps{"post": "/clicked"},
			},
			InnerHTML: "hello, world",
			Disabled:  true,
		},
	)

	hr := Hr(HrProps{})
	hello := B(BProps{InnerHTML: "Bold Hello"})
	header := Header(
		HeaderProps{
			GlobalProps: GlobalProps{
				Data: DataValues{"current": "this is the current page"},
			},
			InnerHTML: button,
		},
	)

	children := Stack(
		header,
		hr,
		hello,
	)

	output := HTML(
		HTMLProps{
			InnerHTML: children,
			Lang:      "en",
		},
	)
	_, _ = min.String("text/html", output)
}
