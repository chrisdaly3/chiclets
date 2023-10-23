package tui

import (
	"fmt"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/chrisdaly3/chiclets/tui/constants"
	"github.com/chrisdaly3/chiclets/tui/styles"
)

// NewFlex creates a flexbox and adds Styles
func NewFlex() *flexbox.FlexBox {
	flex := flexbox.New(0, 0).StylePassing(false)

	r1 := flex.NewRow().AddCells(
		flexbox.NewCell(4, 4).SetStyle(styles.FlexStyleBackground),
		flexbox.NewCell(8, 4).SetStyle(styles.FlexStyleBackground).
			SetContent(constants.HOMEHEADER),
		flexbox.NewCell(4, 4).SetStyle(styles.FlexStyleBackground),
	)

	r2 := flex.NewRow().AddCells(
		flexbox.NewCell(1, 8).SetStyle(styles.FlexStyleTable),
		flexbox.NewCell(6, 8).SetStyle(styles.FlexStyleTable),
		flexbox.NewCell(5, 8).SetStyle(styles.FlexStyleBackgroundNoBorder),
		flexbox.NewCell(1, 8).SetStyle(styles.FlexStyleTable),
	)

	r3 := flex.NewRow().AddCells(
		flexbox.NewCell(2, 3).SetStyle(styles.FlexStyleIce).
			SetContent(fmt.Sprintf(HelpText,
				Keybindings.Up.Help().Key,
				Keybindings.Down.Help().Key,
				Keybindings.Left.Help().Key,
				Keybindings.Right.Help().Key,
				Keybindings.Select.Help().Key,
				Keybindings.Previous.Help().Key,
				Keybindings.Search.Help().Key,
				Keybindings.Esc.Help().Key,
				Keybindings.Quit.Help().Key)),
	)

	flexRows := []*flexbox.Row{r1, r2, r3}
	flex.AddRows(flexRows)
	return flex
}

// NewLeagueTable fills the UIModel table with teams data
func NewLeagueTable(rows [][]string) *table.TableSingleType[string] {
	var LeagueTable = table.NewTableSingleType[string](0, 0, HomeHeaders)
	LeagueTable.SetRatio(ratio).SetMinWidth(minSize)
	LeagueTable.AddRows(rows)
	return LeagueTable
}

// NewTeamTable is a helper function called by the GetTeamInfoCmd
// to populate a new table with Team Roster data
func NewTeamTable(rows [][]string) *table.TableSingleType[string] {
	var RosterTable = table.NewTableSingleType[string](0, 0, TeamHeaders)
	RosterTable.SetRatio(ratio).SetMinWidth(minSize)
	RosterTable.AddRows(rows)
	return RosterTable
}
