package elements

// Ref: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#table_content

import . "github.com/bitpartio/gomx/utils"

/*
 * Specifies the caption (or title) of a table.
 */
type CaptionProps struct {
	GlobalProps
	InnerHTML string
}

func Caption(props CaptionProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <caption {{global}}>
		  {{innerhtml}}
		</caption>
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a column within a table and is used for defining common semantics
 * on all common cells. It is generally found within a <colgroup> element.
 */
type ColProps struct {
	GlobalProps
}

func Col(props ColProps) string {
	values := map[string]interface{}{
		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <col {{global}} />
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a group of columns within a table.
 */
type ColgroupProps struct {
	GlobalProps
	InnerHTML string
}

func Colgroup(props ColgroupProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <colgroup {{global}}>
		  {{innerhtml}}
		</colgroup>
  `)

	s := Render(t, values)
	return s
}

/*
 * Represents tabular data — that is, information presented in a
 * two-dimensional table comprised of rows and columns of cells containing
 * data.
 */
type TableProps struct {
	GlobalProps
	InnerHTML string
}

func Table(props TableProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <table {{global}}>
		  {{innerhtml}}
		</table>
  `)

	s := Render(t, values)
	return s
}

/*
 * Encapsulates a set of table rows (<tr> elements), indicating that they
 * comprise the body of the table (<table>).
 */
type TbodyProps struct {
	GlobalProps
	InnerHTML string
}

func Tbody(props TbodyProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <tbody {{global}}>
		  {{innerhtml}}
		</tbody>
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a cell of a table that contains data. It participates in the
 * table model.
 */
type TdProps struct {
	GlobalProps
	InnerHTML string
}

func Td(props TdProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <td {{global}}>{{innerhtml}}</td>
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a set of rows summarizing the columns of the table.
 */
type TfootProps struct {
	GlobalProps
	InnerHTML string
}

func Tfoot(props TfootProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <tfoot {{global}}>
		  {{innerhtml}}
		</tfoot> `)

	s := Render(t, values)
	return s
}

/*
 * Defines a cell as header of a group of table cells. The exact nature of
 * this group is defined by the scope and headers attributes.
 */
type ThProps struct {
	GlobalProps
	InnerHTML string
}

func Th(props ThProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <th {{global}}>{{innerhtml}}</th>
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a set of rows defining the head of the columns of the table.
 */
type TheadProps struct {
	GlobalProps
	InnerHTML string
}

func Thead(props TheadProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <thead {{global}}>
		  {{innerhtml}}
		</thead>
  `)

	s := Render(t, values)
	return s
}

/*
 * Defines a row of cells in a table. The row's cells can then be
 * established using a mix of <td> (data cell) and <th> (header cell)
 * elements.
 */
type TrProps struct {
	GlobalProps
	InnerHTML string
}

func Tr(props TrProps) string {
	values := map[string]interface{}{
		"innerhtml": props.InnerHTML,

		"global": RenderGlobalProps(props.GlobalProps),
	}

	t := Mx(`
    <tr {{global}}>
		  {{innerhtml}}
		</tr>
  `)

	s := Render(t, values)
	return s
}
