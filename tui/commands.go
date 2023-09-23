package tui

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chrisdaly3/chiclets/data"
	"github.com/chrisdaly3/chiclets/tui/constants"
	"github.com/tidwall/gjson"
)

// If the first column changes, you MUST account for the ID elsewhere
// NOTE: THIS IS BRITTLE AS F*CK
func (ui *UIModel) selectedCmd() tea.Msg {
	column, _ := ui.table.GetCursorLocation()
	switch column {
	case 1:
		ui.table.CursorLeft()
	case 2:
		ui.table.CursorLeft()
		ui.table.CursorLeft()
	case 3:
		ui.table.CursorLeft()
		ui.table.CursorLeft()
		ui.table.CursorLeft()
	}

	return constants.SelectionMessage{}
}

func (ui *UIModel) GetRosterCmd() tea.Msg {
	// API CALL HAPPENS HERE
	id := ui.table.GetCursorValue()
	requestPath := fmt.Sprintf(constants.ROSTERURL, id)

	response, err := http.Get(requestPath)
	if err != nil {
		fmt.Printf("Error communicating with NHL API: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		slog.Error("Unhealthy Response", "response:", response.Body, "statusCode:", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Read Error", "err:", err)
		os.Exit(1)
	}

	// ROSTER COLLECTION from NHL API
	TotalPlayerCount := gjson.GetBytes(responseBody, "teams.0.roster.roster.#")
	var allPlayers []gjson.Result
	var PlayerRows [][]string

	for i := 0; int64(i) < TotalPlayerCount.Int(); i++ {
		playerFields := fmt.Sprintf(
			"{teams.0.roster.roster.%v.person.id,teams.0.roster.roster.%v.person.fullName,teams.0.roster.roster.%v.position.name,teams.0.roster.roster.%v.jerseyNumber}.@join",
			i, // PlayerId
			i, // PlayerName
			i, // PlayerPosition
			i, // PlayerNumber
		)
		player := gjson.GetBytes(responseBody, playerFields)
		allPlayers = append(allPlayers, player)
	}
	for _, p := range allPlayers {
		player := p.Map()
		playerColumn := []string{player["id"].String(), player["fullName"].String(), player["name"].String(), player["jerseyNumber"].String()}
		PlayerRows = append(PlayerRows, playerColumn)
	}

	playerTable := NewTeamTable(PlayerRows)

	// Returns a message to the UI to update the table view
	return constants.RosterMessage{Table: playerTable}
}

func GetLeagueCmd() tea.Msg {
	leagueTable := NewLeagueTable(data.TeamsTable)
	return constants.LeagueMessage{Table: leagueTable}
}

func (ui *UIModel) GetPlayerCmd() tea.Msg {
	return constants.PlayerMessage{}
}
