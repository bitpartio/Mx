package main

import (
	"time"

	. "github.com/bitpartio/Mx/elements"
	. "github.com/bitpartio/Mx/utils"
)

func buildHtmlTestPage() string {
	doctype := Doctype()

	// Whitespace is meaningful in text within <pre> and <code>
	// Placing it in a function is cleaner
	code := preformattedCode()

	body := Body(BodyProps{
		InnerHTML: Stack(
			Del(DelProps{
				Datetime:  time.Now(),
				InnerHTML: "Deleted text.",
			}),
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
			H2(HProps{InnerHTML: "Basic block level elements"}),
			P(PProps{
				InnerHTML: Stack(
					"This is a normal paragraph (",
					Code(CodeProps{InnerHTML: "p"}),
					" element). To add some length to it, let us mention that this page was primarily written for testing the effect of ",
					Strong(StrongProps{InnerHTML: "user style sheets"}),
					". You can use it for various other purposes as well, like just checking how your browser displays various HTML elements.",
				),
			}),
			P(PProps{
				InnerHTML: Stack(
					"This is another paragraph. ",
					Mark(MarkProps{InnerHTML: "I think it needs to be added that the set of elements tested is not exhaustive in any sense."}),
					" I have selected those elements for which it can make sense to write user style sheet rules, in my opinion.",
				),
			}),
			Div(DivProps{
				InnerHTML: Stack(
					"This is a ",
					Code(CodeProps{InnerHTML: "div"}),
					" element. Authors may use such elements instead of paragraph markup for various reasons. (End of ",
					Code(CodeProps{InnerHTML: "div"}),
					".)",
				),
			}),
			Blockquote(BlockquoteProps{
				InnerHTML: P(PProps{
					InnerHTML: Stack(
						"This is a ",
						I(IProps{InnerHTML: "block quotation"}),
						" containing a single paragraph. Well, not quite, since this is not ",
						Em(EmProps{InnerHTML: "really"}),
						" quoted text, but I hope you understand the point. After all, this page does not use HTML markup very normally anyway.",
					),
				}),
			}),
			P(PProps{InnerHTML: "The following contains links to the Comm-244 home page"}),
			P(PProps{
				InnerHTML: Stack(
					A(AProps{
						Href:      "http://web.simmons.edu/~grovesd/comm244",
						InnerHTML: "Comm-244 Website",
					}),
					", ",
					A(AProps{
						Href:      "http://web.simmons.edu/~grovesd/comm244/week3",
						InnerHTML: "Week 3 page for class",
					}),
				),
			}),
			H2(HProps{InnerHTML: "Lists"}),
			P(PProps{
				InnerHTML: Stack(
					"This is a paragraph before an ",
					Strong(StrongProps{InnerHTML: "unordered"}),
					" list (",
					Code(CodeProps{InnerHTML: "ul"}),
					`). Note that the spacing between a paragraph and a list before or after that is hard to tune in a user style sheet. You can't guess which paragraphs are logically related to a list, e.g. as a "list header".`,
				),
			}),
			Ul(UlProps{
				InnerHTML: Stack(
					Li(LiProps{InnerHTML: " One."}),
					Li(LiProps{InnerHTML: " Two."}),
					Li(LiProps{InnerHTML: " Three. Well, probably this list item should be longer. Note that for short items lists look better if they are compactly presented, whereas for long items, it would be better to have more vertical spacing between items."}),
					Li(LiProps{InnerHTML: " Four. This is the last item in this list Let us terminate the list now without making any more fuss about it."}),
				),
			}),
			P(PProps{
				InnerHTML: Stack(
					"This is a paragraph before a ",
					Strong(StrongProps{InnerHTML: "ordered"}),
					" list (",
					Code(CodeProps{InnerHTML: "ol"}),
					`). Note that the spacing between a paragraph and a list before or after that is hard to tune in a user style sheet. You can't guess which paragraphs are logically related to a list, e.g. as a "list header".`,
				),
			}),
			Ol(OlProps{
				InnerHTML: Stack(
					Li(LiProps{InnerHTML: " One."}),
					Li(LiProps{InnerHTML: " Two."}),
					Li(LiProps{InnerHTML: " Three. Well, probably this list item should be longer. Note that if items are short, lists look better if they are compactly presented, whereas for long items, it would be better to have more vertical spacing between items."}),
					Li(LiProps{InnerHTML: " Four. This is the last item in this list. Let us terminate the list now without making any more fuss about it."}),
				),
			}),
			P(PProps{
				InnerHTML: Stack(
					"This is a paragraph before a ",
					Strong(StrongProps{InnerHTML: "definition"}),
					" list (",
					Code(CodeProps{InnerHTML: "dl"}),
					"). In principle, such a list should consist of ",
					Em(EmProps{InnerHTML: "terms"}),
					" and associated definitions. But many authors use ",
					Code(CodeProps{InnerHTML: "dl"}),
					` elements for fancy "layout" things. Usually the effect is not `,
					Em(EmProps{InnerHTML: "too"}),
					" bad, if you design user style sheet rules for ",
					Code(CodeProps{InnerHTML: "dl"}),
					" which are suitable for real definition lists.",
				),
			}),
			Dl(DlProps{InnerHTML: Stack(
				Dt(DtProps{InnerHTML: " recursion "}),
				Dd(DdProps{InnerHTML: " see recursion "}),
				Dt(DtProps{InnerHTML: " recursion, indirect"}),
				Dd(DdProps{InnerHTML: " see indirect recursion"}),
				Dt(DtProps{InnerHTML: " indirect recursion"}),
				Dd(DdProps{InnerHTML: " see recursion, indirect"}),
				Dt(DtProps{InnerHTML: " term"}),
				Dd(DdProps{InnerHTML: " a word or other expression taken into specific use in a well-defined meaning, which is often defined rather rigorously, even formally, and may differ quite a lot from an everyday meaning"}),
			),
			}),
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
