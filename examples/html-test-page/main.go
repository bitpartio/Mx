package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/bitpartio/gomx/elements"
	. "github.com/bitpartio/gomx/utils"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

var min *minify.M

func init() {
	min = minify.New()

	min.AddFunc("text/html", html.Minify)
}

func main() {
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

	hr := Hr(HrProps{GlobalProps{}})
	hello := B(BProps{InnerHTML: "Bold Hello"})
	header := Header(
		HeaderProps{
			GlobalProps: GlobalProps{
				Data: DataValues{"current": "this is the current page"},
			},
			InnerHTML: button,
		},
	)
	slot := Slot(SlotProps{InnerHTML: "Slot Hello", Name: "slot-name"})

	children := Stack(
		header,
		hr,
		hello,
		slot,
	)

	output := HTML(
		HTMLProps{
			InnerHTML: children,
			Lang:      "en",
		},
	)
	s, _ := min.String("text/html", output)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, s)
	})

	http.HandleFunc("/basic", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, _ := ReadTemplate("templates/html.html")
		body := "Hello, World"
		v := Values{"body": body}
		b := Render(t, v)
		fmt.Fprintf(w, b)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
