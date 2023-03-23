package main

import (
	. "github.com/bitpartio/gomx/elements"
	. "github.com/bitpartio/gomx/utils"
)

func buildHtmlTestPage() string {
	doctype := Doctype()

	// Whitespace is meaningful in text within <pre> and <code>
	// Placing it in a function is cleaner
	code := preformattedCode()

	body := Body(BodyProps{
		InnerHTML: Stack(
			H1(HProps{InnerHTML: "Testing display of HTML elements"}),
			P(PProps{InnerHTML: Stack(
				"This page contains a bunch of HTML Elements and text. You can copy the source code and use it test out various CSS Properties. For testing purposes, you may use ",
				A(AProps{
					Href:      "css-linking#internal",
					InnerHTML: Dfn(DfnProps{InnerHTML: "internal styles"}),
				}),
				". Recall that these CSS rules are placed in between the ",
				Code(CodeProps{InnerHTML: "head"}),
				" tags using the following format:",
			)},
			),
			Pre(PreProps{
				InnerHTML: Code(CodeProps{
					InnerHTML: code,
				}),
			}),
			H1(HProps{InnerHTML: "This is 1st level heading"}),
			P(PProps{InnerHTML: "This is a test paragraph."}),
			H2(HProps{InnerHTML: "This is 2nd level heading"}),
			P(PProps{InnerHTML: "This is a test paragraph."}),
			H3(HProps{InnerHTML: "This is 3rd level heading"}),
			P(PProps{InnerHTML: "This is a test paragraph."}),
			H4(HProps{InnerHTML: "This is 4th level heading"}),
			P(PProps{InnerHTML: "This is a test paragraph."}),
			H5(HProps{InnerHTML: "This is 5th level heading"}),
			P(PProps{InnerHTML: "This is a test paragraph."}),
			H6(HProps{InnerHTML: "This is 6th level heading"}),
			P(PProps{InnerHTML: "This is a test paragraph."}),
		)},
	)

	html := HTML(HTMLProps{
		Lang: "en",
		InnerHTML: Head(HeadProps{
			InnerHTML: Stack(
				Meta(MetaProps{Charset: "UTF-8"}),
				Title(TitleProps{InnerHTML: "HTML Page for Testing CSS"}),
				body,
			),
		}),
	})

	output := Stack(
		doctype,
		html,
	)

	return output
}

func preformattedCode() string {
	return `&lt;style type=&quot;text/css&quot;&gt;
  selector {
    property: value
  }
&lt;/style&gt;`
}
