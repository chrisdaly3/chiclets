package tui

import (
	"fmt"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chrisdaly3/chiclets/data"
	"github.com/chrisdaly3/chiclets/tui/constants"
)

// view keeps track
// of what we're "focused" on
type view int

const (
	homeNav view = iota
	teamNav
)

type UIModel struct {
	view  view
	flex  *flexbox.FlexBox
	table *table.TableSingleType[string]
}

var headers = []string{"Locale", "Team Name", "Division"}
var InitModel = UIModel{
	view:  homeNav,
	flex:  flexbox.New(0, 0),
	table: table.NewTableSingleType[string](0, 0, headers),
}
var ratio = []int{4, 6, 4}
var minSize = []int{5, 4, 8}

func (ui *UIModel) PopulateTable() {
	ui.table.SetRatio(ratio).SetMinWidth(minSize)
	ui.table.AddRows(data.TeamsTable)
}

func (ui *UIModel) selected() tea.Msg {
	return constants.SelectionMessage{}
}

func (ui *UIModel) Init() tea.Cmd { return nil }

func (ui *UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.flex.SetWidth(msg.Width)
		ui.table.SetWidth(msg.Width)
		ui.table.SetHeight(msg.Height - ui.flex.GetHeight())

		// TODO: Either find a way to access values on position,
		// OR uncover a better way to handle the selected item for stats
		// on the teamNav view after selection. Might require some adjustments
		// to the initial API response
	case constants.SelectionMessage:
		column, row := ui.table.GetCursorLocation()
		fmt.Print(column, row)
	// Add All Keybindings here
	case tea.KeyMsg:
		// might need to add conditional logic for view / state consideration here
		switch {
		case key.Matches(msg, Keybindings.Quit):
			return ui, tea.Quit

		case key.Matches(msg, Keybindings.Down):
			ui.table.CursorDown()

		case key.Matches(msg, Keybindings.Up):
			ui.table.CursorUp()

		case key.Matches(msg, Keybindings.Left):
			ui.table.CursorLeft()

		case key.Matches(msg, Keybindings.Right):
			ui.table.CursorRight()

		case key.Matches(msg, Keybindings.Select):
			return ui, ui.selected
		}
	}
	return ui, nil
}

func (ui *UIModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, ui.flex.Render(), ui.table.Render())
}
