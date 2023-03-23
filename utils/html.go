package utils

import (
	"regexp"

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
	// TODO
	// Possibly throw a warning that Clean is not for production use
	// and is only to for development
	// Write tests to make sure it doesn't mess up html
	// Should cover most cases, but not all
	c := gohtml.Format(s)

	// Remove whitespace at the end of tags
	r1 := regexp.MustCompile(`\s+>`)
	o := r1.ReplaceAllString(c, `>`)

	// Remove whitespace after tags
	r2 := regexp.MustCompile(`(<\w+)([ ]{2,})(.*>)`)
	o = r2.ReplaceAllString(o, "$1 $3")

	// Remove whitespace between values
	r3 := regexp.MustCompile(`('|")([ ]{2,})(.*>)`)
	o = r3.ReplaceAllString(o, "$1 $3")

	return o
}
