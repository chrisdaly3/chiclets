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

func (ui *UIModel) GetTeamInfoCmd() tea.Msg {
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
	prevGame := ui.GetPreviousCmd()

	// Returns a message to the UI to update the table view
	return constants.TeamInfoMessage{Table: playerTable, TeamPriorGame: prevGame}
}

func (ui *UIModel) GetPreviousCmd() map[string]gjson.Result {
	id := ui.table.GetCursorValue()
	requestPath := fmt.Sprintf(constants.LASTGAME, id)

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

	previousGame := gjson.GetBytes(responseBody, "{teams.0.previousGameSchedule.dates.0.date,teams.0.previousGameSchedule.dates.0.games.0.teams,teams.0.previousGameSchedule.dates.0.games.0.venue.name}.@join")
	awayTeam := gjson.Result.Get(previousGame, "{teams.away.team.name,teams.away.score}").Map()
	homeTeam := gjson.Result.Get(previousGame, "{teams.home.team.name,teams.home.score}").Map()

	prevGameMap := make(map[string]gjson.Result)
	prevGameMap["date"] = previousGame.Map()["date"]
	prevGameMap["away"] = awayTeam["name"]
	prevGameMap["awayScore"] = awayTeam["score"]
	prevGameMap["home"] = homeTeam["name"]
	prevGameMap["homeScore"] = homeTeam["score"]

	return prevGameMap
}

//TODO: Setup next game and team record calls + parsers
/*func (ui *UIModel) GetNextCmd() tea.Msg {
	id := ui.table.GetCursorValue()
	requestPath := fmt.Sprintf(constants.NEXTGAME, id)

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

}

func (ui *UIModel) GetRecordCmd() tea.Msg {
	id := ui.table.GetCursorValue()
	requestPath := fmt.Sprintf(constants.RECORDURL, id)

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

}*/

func GetLeagueCmd() tea.Msg {
	leagueTable := NewLeagueTable(data.TeamsTable)
	return constants.LeagueMessage{Table: leagueTable}
}

func (ui *UIModel) GetPlayerCmd() tea.Msg {
	id := ui.table.GetCursorValue()
	stats := getPlayerStats(id)
	return constants.PlayerMessage{Player: stats}
}

// getPlayerStats is a helper function to obtain stats for the selected player
func getPlayerStats(id string) map[string]gjson.Result {
	requestPath := fmt.Sprintf(constants.PLAYERURL, id)

	response, err := http.Get(requestPath)
	if err != nil {
		fmt.Printf("Error communicating with the NHL API: %v", err)
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

	//PLAYER STAT COLLECTION from NHL API
	playerResponse := gjson.GetBytes(responseBody, "stats.0.splits.0.stat.@pretty")
	return playerResponse.Map()
}
