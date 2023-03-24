package main

import (
	"testing"

	. "github.com/bitpartio/gomx/utils"
)

func BenchmarkRender(b *testing.B) {
	for n := 0; n < b.N; n++ {
		render()
	}
}

func render() {
	page := buildHtmlTestPage()

	_ = Minify(page)
}
