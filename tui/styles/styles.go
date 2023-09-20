package styles

import (
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/lipgloss"
)

/* COLORS */
var (
	Black  = lipgloss.Color("#010101")
	Navy   = lipgloss.Color("#14213D")
	Violet = lipgloss.Color("#615789")
	Orange = lipgloss.Color("#FCA311")
	Grey   = lipgloss.Color("#A4A9AD")
	White  = lipgloss.Color("#FFFFFF")
	Red    = lipgloss.Color("#DB0A0A")
	Ice    = lipgloss.Color("#70B2C4")
)

var (
	/* FLEXBOX STYLES */
	FlexStyleBackground = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Background(Grey).
				Border(lipgloss.ThickBorder(), true).
				BorderBackground(Black, White).
				BorderForeground(Black)

	FlexStyleBackgroundNoBorder = lipgloss.NewStyle().
					Align(lipgloss.Center).
					Background(Grey).Margin(0, 1)

	FlexStyleNavy = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Background(Navy)

	FlexStyleOrange = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Background(Orange)

	FlexStyleViolet = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Background(Violet)

	FlexStyleIce = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Background(Ice).
			Foreground(Black).
			Italic(true).
			Bold(true).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(White).Margin(0, 15)

	FlexStyleBlank = lipgloss.NewStyle()
)

/* TABLE STYLES */
// Table overrides signature
// func (r *Table) SetStyles(styles map[TableStyleKey]lipgloss.Style) *Table
var TableStyles = map[table.TableStyleKey]lipgloss.Style{
	table.TableHeaderStyleKey:         TableHeaderStyle,
	table.TableRowsStyleKey:           TableRowsStyle,
	table.TableRowsSubsequentStyleKey: TableRowsSubsequentStyle,
	table.TableFooterStyleKey:         TableFooterStyle,
	table.TableCellCursorStyleKey:     TableCellCursorStyle,
	table.TableRowsCursorStyleKey:     TableRowsCursorStyle,
}

var (
	TableHeaderStyle = lipgloss.NewStyle().
				Background(Black).
				Bold(true).
				Align(lipgloss.Left).
				Italic(true).
				Foreground(Grey)

	TableRowsStyle = lipgloss.NewStyle().
			Background(White).
			Align(lipgloss.Left).
			Foreground(Black).
			Bold(true)

	TableRowsSubsequentStyle = lipgloss.NewStyle().
					Align(lipgloss.Left).
					Background(Grey).
					Foreground(Black).
					Bold(true)

	TableFooterStyle = lipgloss.NewStyle().
				Bold(true).
				Italic(true).
				Background(Black).
				Foreground(Ice)

	TableCellCursorStyle = lipgloss.NewStyle().
				Background(Red).
				Foreground(Grey)

	TableRowsCursorStyle = lipgloss.NewStyle().
				Background(Black).
				Foreground(Grey)
)
