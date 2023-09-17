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

// NewTeamTable fills the UI table with data
// dependent on SelectionMessage
func NewTeamTable() *table.TableSingleType[string] {
	var Roster = table.NewTableSingleType[string](0, 0, TeamHeaders)
	Roster.SetRatio(ratio).SetMinWidth(minSize)
	var RosterRows = [][]string{[]string{"player1ID", "Player1Name", "Player1Position", "Player1Number"},
		[]string{"player2ID", "Player2Name", "Player2Position", "Player2Number"}}
	//TODO: AddRows based on Team Roster API response
	Roster.AddRows(RosterRows)
	return Roster
}
