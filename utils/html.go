package utils

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/yosssi/gohtml"
)

var min *minify.M

func init() {
	min = minify.New()

	min.AddFunc("text/html", html.Minify)
}

/*
 * Minify HTML
 */
func Minify(s string) string {
	o, err := min.String("text/html", s)
	if err != nil {
		return s
	}
	return o
}

/*
 * Clean HTML to readable format
 */
func Clean(s string) string {
	return gohtml.Format(s)
}
