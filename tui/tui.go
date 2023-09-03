package tui

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// var headers will eventually be replaced by api call stats
var headers = []string{"header1", "header2", "header3"}
var InitModel = UIModel{view: homeNav, flex: flexbox.New(0, 0), table: table.NewTableSingleType[string](0, 0, headers)}

func (ui *UIModel) Init() tea.Cmd { return nil }

func (ui *UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.flex.SetWidth(msg.Width)
		ui.table.SetWidth(msg.Width)
		ui.table.SetHeight(msg.Height - ui.flex.GetHeight())

	// Add All Keybindings here
	case tea.KeyMsg:
		// might need to add conditional logic for view / state consideration here
		switch {
		case key.Matches(msg, Keybindings.Quit):
			return ui, tea.Quit
		}
	}
	return ui, nil
}

func (ui *UIModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, ui.flex.Render(), ui.table.Render())
}
