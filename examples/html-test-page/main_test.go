package main

import (
	"testing"

	. "github.com/bitpartio/Mx/utils"
)

var staticPage string

func init() {
	staticPage, _ = RenderTemplate("html-test-page.html")
}

func BenchmarkRender(b *testing.B) {
	for n := 0; n < b.N; n++ {
		render()
	}
}

func render() {
	page := buildHtmlTestPage()

	_ = Minify(page)
}

func BenchmarkRenderStatic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		renderStatic()
	}
}

func renderStatic() {
	_ = Minify(staticPage)
}
