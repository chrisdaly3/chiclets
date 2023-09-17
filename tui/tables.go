package tui

import (
	"fmt"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/chrisdaly3/chiclets/data"
	"github.com/chrisdaly3/chiclets/tui/styles"
)

// NewLeagueTable fills the UIModel table with teams data
func (ui *UIModel) NewLeagueTable() {
	ui.table.SetRatio(ratio).SetMinWidth(minSize)
	ui.table.AddRows(data.TeamsTable)
	r1 := ui.flex.NewRow().AddCells(
		flexbox.NewCell(3, 2),
		flexbox.NewCell(6, 3).SetStyle(styles.FlexStyleBackground),
		flexbox.NewCell(3, 2),
	)

	r2 := ui.flex.NewRow().AddCells(
		flexbox.NewCell(5, 8).SetStyle(styles.FlexStyleOrange),
		flexbox.NewCell(10, 8).SetStyle(styles.FlexStyleTableBackground),
		flexbox.NewCell(5, 8).SetStyle(styles.FlexStyleOrange),
	)

	r3 := ui.flex.NewRow().AddCells(
		flexbox.NewCell(3, 2),
		flexbox.NewCell(6, 3).SetStyle(styles.FlexStyleIce).
			SetContent(fmt.Sprintf(HelpText,
				Keybindings.Up.Help().Key,
				Keybindings.Down.Help().Key,
				Keybindings.Left.Help().Key,
				Keybindings.Right.Help().Key,
				Keybindings.Select.Help().Key,
				Keybindings.Search.Help().Key,
				Keybindings.Esc.Help().Key,
				Keybindings.Quit.Help().Key)),
		flexbox.NewCell(3, 2),
	)
	flexRows := []*flexbox.Row{r1, r2, r3}
	ui.flex.AddRows(flexRows)
}

// NewTeamTable is a helper function called by the GetRosterCmd
// to populate a new table with Team Roster data
func NewTeamTable(rows [][]string) *table.TableSingleType[string] {
	var RosterTable = table.NewTableSingleType[string](0, 0, TeamHeaders)
	RosterTable.SetRatio(ratio).SetMinWidth(minSize)
	RosterTable.AddRows(rows)
	return RosterTable
}