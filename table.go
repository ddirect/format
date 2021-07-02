package format

import (
	"fmt"
	"strings"
)

type TableRow []interface{}
type Table []TableRow

const margin = 4

func (t Table) String() string {
	// convert to string matrix and get the max number of columns
	var st [][]string
	columns := 0
	for _, sRow := range t {
		if len(sRow) > columns {
			columns = len(sRow)
		}
		var dRow []string
		for _, sItem := range sRow {
			dRow = append(dRow, fmt.Sprintf("%v", sItem))
		}
		st = append(st, dRow)
	}
	// compute the max width for each column
	var widths []int
	for col := 0; col < columns; col++ {
		width := 0
		for _, row := range st {
			if col < len(row) {
				colWidth := len(row[col])
				if width < colWidth {
					width = colWidth
				}
			}
		}
		widths = append(widths, width+margin)
	}
	// format to string
	sb := new(strings.Builder)
	for _, row := range st {
		for col, item := range row {
			fmt.Fprintf(sb, "%*s", widths[col], item)
		}
		fmt.Fprintln(sb)
	}
	return sb.String()
}
