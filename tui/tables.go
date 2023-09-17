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
	ui.table.AddRows(data.TeamsTable).SetStylePassing(true)
	r1 := ui.flex.NewRow().AddCells(
		flexbox.NewCell(3, 5).SetStyle(styles.FlexStyleNavy),
		flexbox.NewCell(6, 5).SetStyle(styles.FlexStyleBackground),
		flexbox.NewCell(3, 5).SetStyle(styles.FlexStyleNavy),
	)

	r2 := ui.flex.NewRow().AddCells(
		flexbox.NewCell(5, 5).SetStyle(styles.FlexStyleOrange),
		flexbox.NewCell(10, 5).SetStyle(styles.FlexStyleBlank),
		flexbox.NewCell(5, 5).SetStyle(styles.FlexStyleOrange),
	)

	r3 := ui.flex.NewRow().AddCells(
		flexbox.NewCell(3, 5).SetStyle(styles.FlexStyleNavy),
		flexbox.NewCell(6, 5).SetStyle(styles.FlexStyleViolet).
			SetContent(styles.FlexStyleText.Render(fmt.Sprintf(HelpText,
				Keybindings.Up.Help().Key,
				Keybindings.Down.Help().Key,
				Keybindings.Left.Help().Key,
				Keybindings.Right.Help().Key,
				Keybindings.Select.Help().Key,
				Keybindings.Search.Help().Key,
				Keybindings.Esc.Help().Key,
				Keybindings.Quit.Help().Key))),
		flexbox.NewCell(3, 5).SetStyle(styles.FlexStyleNavy),
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
