package main

import (
	"fmt"
	"os"
	"testing"

	. "github.com/bitpartio/Mx/utils"
)

var staticPage string

func init() {
	staticPageBytes, err := os.ReadFile("html-test-page.html")
	if err != nil {
		fmt.Println("Could not read html test page:", err)
	} else {
		staticPage = string(staticPageBytes)
	}
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
