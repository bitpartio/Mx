package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/bitpartio/Mx/utils"
)

// Ref #2: https://github.com/cbracco/html5-test-page

func main() {
	page := buildHtmlTestPage()

	s := Minify(page)
	// s := Clean(page)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, s)
	})

	http.HandleFunc("/basic", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, _ := ReadTemplate("../../templates/html.html")
		body := "Hello, World"
		v := Values{"body": body}
		b := Render(t, v)
		fmt.Fprintf(w, b)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
