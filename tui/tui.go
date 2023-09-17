package tui

import (
	"fmt"
	"unicode"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chrisdaly3/chiclets/tui/constants"
	"github.com/chrisdaly3/chiclets/tui/styles"
)

// type view tracks what HTML to render
// either All Team info, or Team Stat Data
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

// Table Headers dependent on view
func (ui *UIModel) getHeaders() []string {
	if ui.view == homeNav {
		return HomeHeaders
	} else if ui.view == teamNav {
		return TeamHeaders
	}
	return []string{"HEADERS", "NOT", "SET", "ERROR"}
}

var HomeHeaders = []string{"ID", "Locale", "Team Name", "Division"}
var TeamHeaders = []string{"ID", "Player", "Position", "Number"}

var InitModel = UIModel{
	view:  homeNav,
	flex:  flexbox.New(0, 0).SetStyle(styles.FlexStyleWhite),
	table: table.NewTableSingleType[string](0, 0, HomeHeaders),
}

var ratio = []int{1, 10, 10, 5}
var minSize = []int{2, 5, 5, 5}

// searchTable accepts a tea.KeyMsg.Str() vand sets a columnar search query in model table.
func (ui *UIModel) searchTable(key string) {
	ind, str := ui.table.GetFilter()
	posX, _ := ui.table.GetCursorLocation()
	if posX != ind && key != "backspace" {
		ui.table.SetFilter(posX, key)
		return
	}
	if key == "backspace" {
		if len(str) == 1 {
			ui.table.UnsetFilter()
			return
		} else if len(str) > 1 {
			str = str[0 : len(str)-1]
		} else {
			return
		}
	} else {
		str += key
	}
	ui.table.SetFilter(ind, str)
}

func (ui *UIModel) Init() tea.Cmd { return nil }

func (ui *UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.flex.SetWidth(msg.Width)
		ui.flex.SetHeight(msg.Height)
		ui.table.SetWidth(msg.Width)
		ui.table.SetHeight(msg.Height - ui.flex.GetHeight())

	case constants.SelectionMessage:
		if ui.view == homeNav {
			return ui, ui.GetRosterCmd
		} else if ui.view == teamNav {
			fmt.Printf("ID: %s", ui.table.GetCursorValue())
		}

	case constants.RosterMessage:
		fmt.Printf("URL is %s: ", msg.URL)

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
			return ui, ui.selectedCmd

		case key.Matches(msg, Keybindings.Esc):
			if _, s := ui.table.GetFilter(); s != "" {
				ui.table.UnsetFilter()
			}

		case key.Matches(msg, Keybindings.Backspace):
			ui.searchTable(msg.String())

		default:
			if len(msg.String()) == 1 && unicode.IsUpper(msg.Runes[0]) {
				ui.searchTable(msg.String())
			}
		}
	}
	return ui, nil
}

func (ui *UIModel) View() string {
	ui.flex.ForceRecalculate()
	_r := ui.flex.GetRow(1)
	_c := _r.GetCell(1)
	ui.table.SetWidth(_c.GetWidth())
	ui.table.SetHeight(_c.GetHeight())
	_c.SetContent(ui.table.Render())
	return ui.flex.Render()
}
