package tui

import (
	"fmt"
	//"runtime"
	"unicode"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chrisdaly3/chiclets/data"
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
	view           view
	flex           *flexbox.FlexBox
	table          *table.TableSingleType[string]
	statsDisplayed bool
}

var HomeHeaders = []string{"ID", "Locale", "Team Name", "Division"}
var TeamHeaders = []string{"ID", "Player", "Position", "Number"}

var InitModel = UIModel{
	view:           homeNav,
	flex:           NewFlex(),
	table:          NewLeagueTable(data.TeamsTable),
	statsDisplayed: false,
}

var ratio = []int{35, 65, 50, 50}
var minSize = []int{5, 5, 5, 5}

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
			return ui, ui.GetTeamInfoCmd
		} else if ui.view == teamNav {
			return ui, ui.GetPlayerCmd
		}

	case constants.TeamInfoMessage:
		//runtime.Breakpoint()
		if ui.view == homeNav {
			ui.view = teamNav
			ui.table = msg.Table
		}

		//update top flexbox rows with Team Information for the season
		var teamStatsMessage = fmt.Sprintf("\nTEAM STATS\n\n%s\n\nPoints: %v   |   Division Rank: %v\nGames Played: %v   |   Current Streak: %v\nRegulation Wins: %v   |   ROW: %v\nGoals For: %v   |   Goals Against: %v",
			msg.TeamStats["teamName"],
			msg.TeamStats["points"],
			msg.TeamStats["divRank"],
			msg.TeamStats["gamesPlayed"],
			msg.TeamStats["streak"],
			msg.TeamStats["regWins"],
			msg.TeamStats["row"],
			msg.TeamStats["goalsFor"],
			msg.TeamStats["goalsAgainst"],
		)
		var lastGameMessage = fmt.Sprintf("\nPREVIOUS GAME\n\nDate: %v\n\nAway Team: %v\n\tScore: %v\n\nHome Team: %v\n\tScore: %v",
			msg.TeamPriorGame["date"],
			msg.TeamPriorGame["away"],
			msg.TeamPriorGame["awayScore"].Int(),
			msg.TeamPriorGame["home"],
			msg.TeamPriorGame["homeScore"].Int(),
		)
		var nextGameMessage = fmt.Sprintf("\nNEXT GAME\n\nDate: %v\n\nAway Team: %v\n\tWins: %v\tLosses: %v\tOTL: %v\n\nHome Team: %v\n\tWins: %v\tLosses: %v\tOTL: %v",
			msg.TeamNextGame["date"],
			msg.TeamNextGame["away"],
			msg.TeamNextGame["awayWins"].Int(),
			msg.TeamNextGame["awayLosses"].Int(),
			msg.TeamNextGame["awayOTL"].Int(),
			msg.TeamNextGame["home"],
			msg.TeamNextGame["homeWins"].Int(),
			msg.TeamNextGame["homeLosses"].Int(),
			msg.TeamNextGame["homeOTL"].Int(),
		)

		seasonCell := ui.flex.GetRow(0).GetCell(1)
		seasonCell.SetContent(teamStatsMessage)
		priorGameCell := ui.flex.GetRow(0).GetCell(0)
		priorGameCell.SetContent(lastGameMessage)
		nextGameCell := ui.flex.GetRow(0).GetCell(2)
		nextGameCell.SetContent(nextGameMessage)

	case constants.LeagueMessage:
		if ui.view == teamNav {
			ui.view = homeNav
			ui.table = msg.Table
		}
		// Unset the stat blocks displayed for the team and players
		// team row
		ui.flex.GetRow(0).GetCell(0).SetContent("")
		ui.flex.GetRow(0).GetCell(1).SetContent(constants.HOMEHEADER)
		ui.flex.GetRow(0).GetCell(2).SetContent("")
		//player row
		ui.flex.GetRow(1).GetCell(0).SetContent("")
		ui.flex.GetRow(1).GetCell(2).SetContent("")

	case constants.PlayerMessage:
		statCell := ui.flex.GetRow(1).GetCell(2)

		//Player Stat block content
		var playerStatsMessage = fmt.Sprintf("\nCURRENT SEASON\n\nGames Played: %v\n\nGoals: %v\n\nAssists: %v\n\n+/-: %v\n\nPenalty Minutes: %v\n\nPower Play Goals: %v\n\nPower Play Pts: %v\n\nShort-hand Goals: %v\n\nShort-hand Points: %v\n\nOvertime Goals: %v\n\nGame-winning Goals: %v\n\nShots: %v\n\nShot Pct: %.3f",
			msg.Player["games"].Int(),
			msg.Player["goals"].Int(),
			msg.Player["assists"].Int(),
			msg.Player["plusMinus"].Int(),
			msg.Player["pim"].Int(),
			msg.Player["powerPlayGoals"].Int(),
			msg.Player["powerPlayPoints"].Int(),
			msg.Player["shortHandedGoals"].Int(),
			msg.Player["shortHandedPoints"].Int(),
			msg.Player["overTimeGoals"].Int(),
			msg.Player["gameWinningGoals"].Int(),
			msg.Player["shots"].Int(),
			msg.Player["shotPct"].Float())

		//Goalie Stat block content
		var goalieStatsMessage = fmt.Sprintf("\nCURRENT SEASON\n\nGames Played: %v\n\nShots Against: %v\n\nSaves: %v\n\nGA: %v\n\nSave Pct: %.3f\n\nGoals Allowed/Gm: %.3f\n\nShutouts: %v",
			msg.Player["games"].Int(),
			msg.Player["shotsAgainst"].Int(),
			msg.Player["saves"].Int(),
			msg.Player["goalsAgainst"].Int(),
			msg.Player["savePercentage"].Float(),
			msg.Player["goalAgainstAverage"].Float(),
			msg.Player["shutouts"].Int(),
		)

		if ui.statsDisplayed == false {
			ui.statsDisplayed = true
			if _, ok := msg.Player["saves"]; ok {
				statCell.SetContent(goalieStatsMessage)
			} else {
				statCell.SetContent(playerStatsMessage)
			}

			// TODO: Set up all player stats and details for current season.
			// Display career stats on left frame? perhaps
		} else if ui.statsDisplayed == true {
			ui.statsDisplayed = false
			statCell.SetContent("")
			if _, ok := msg.Player["saves"]; ok {
				statCell.SetContent(goalieStatsMessage)
			} else {
				statCell.SetContent(playerStatsMessage)
			}
		}

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

		case key.Matches(msg, Keybindings.Previous):
			return ui, GetLeagueCmd

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
	ui.table.SetStyles(styles.TableStyles)
	ui.table.SetWidth(_c.GetWidth())
	ui.table.SetHeight(_c.GetHeight())
	_c.SetContent(ui.table.Render())
	return ui.flex.Render()
}
