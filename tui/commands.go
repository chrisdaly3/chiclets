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
	forwardsCount := gjson.GetBytes(responseBody, "forwards.#")
	defensemenCount := gjson.GetBytes(responseBody, "defensemen.#")
	goaliesCount := gjson.GetBytes(responseBody, "goalies.#")
	var allPlayers []gjson.Result
	var PlayerRows [][]string

	for i := 0; int64(i) < forwardsCount.Int(); i++ {
		playerFields := fmt.Sprintf(
			"{forwards.%d.id,\"firstName\":forwards.%d.firstName.default,\"lastName\":forwards.%d.lastName.default,forwards.%d.positionCode,forwards.%d.sweaterNumber}",
			i, // PlayerId
			i, // PlayerFirstName
			i, // PlayerLastName
			i, // PlayerPosition
			i, // PlayerNumber
		)
		player := gjson.GetBytes(responseBody, playerFields)
		allPlayers = append(allPlayers, player)
	}
	for i := 0; int64(i) < defensemenCount.Int(); i++ {
		playerFields := fmt.Sprintf(
			"{defensemen.%d.id,\"firstName\":defensemen.%d.firstName.default,\"lastName\":defensemen.%d.lastName.default,defensemen.%d.positionCode,defensemen.%d.sweaterNumber}",
			i, // PlayerId
			i, // PlayerFirstName
			i, // PlayerLastName
			i, // PlayerPosition
			i, // PlayerNumber
		)
		player := gjson.GetBytes(responseBody, playerFields)
		allPlayers = append(allPlayers, player)
	}
	for i := 0; int64(i) < goaliesCount.Int(); i++ {
		playerFields := fmt.Sprintf(
			"{goalies.%d.id,\"firstName\":goalies.%d.firstName.default,\"lastName\":goalies.%d.lastName.default,goalies.%d.positionCode,goalies.%d.sweaterNumber}",
			i, // PlayerId
			i, // PlayerFirstName
			i, // PlayerLastName
			i, // PlayerPosition
			i, // PlayerNumber
		)
		player := gjson.GetBytes(responseBody, playerFields)
		allPlayers = append(allPlayers, player)
	}
	for _, p := range allPlayers {
		player := p.Map()
		playerColumn := []string{player["id"].String(), player["firstName"].String() + " " + player["lastName"].String(), player["positionCode"].String(), player["sweaterNumber"].String()}
		PlayerRows = append(PlayerRows, playerColumn)
	}

	playerTable := NewTeamTable(PlayerRows)
	teamStandings := ui.GetRecordCmd()

	// Returns a message to the UI to update the table view
	return constants.TeamInfoMessage{Table: playerTable, TeamStats: teamStandings}
}

func (ui *UIModel) GetRecordCmd() map[string]gjson.Result {
	id := ui.table.GetCursorValue()
	requestPath := constants.TEAMSURL

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

	records := gjson.GetBytes(responseBody, fmt.Sprintf("standings.#(teamAbbrev.default==%v)", id))
	name := gjson.Result.Get(records, "teamName.default")
	points := gjson.Result.Get(records, "points")
	gamesPlayed := gjson.Result.Get(records, "gamesPlayed")
	divRank := gjson.Result.Get(records, "divisionSequence")
	streakCode := gjson.Result.Get(records, "streakCode")
	streakCount := gjson.Result.Get(records, "streakCount")
	regWins := gjson.Result.Get(records, "regulationWins")
	row := gjson.Result.Get(records, "regulationPlusOtWins")
	goalsFor := gjson.Result.Get(records, "goalFor")
	goalsAgainst := gjson.Result.Get(records, "goalAgainst")

	recordsMap := make(map[string]gjson.Result)
	recordsMap["teamName"] = name
	recordsMap["points"] = points
	recordsMap["gamesPlayed"] = gamesPlayed
	recordsMap["divRank"] = divRank
	recordsMap["streakCode"] = streakCode
	recordsMap["streakCount"] = streakCount
	recordsMap["regWins"] = regWins
	recordsMap["row"] = row
	recordsMap["goalsFor"] = goalsFor
	recordsMap["goalsAgainst"] = goalsAgainst

	return recordsMap
}

func GetLeagueCmd() tea.Msg {
	leagueTable := NewLeagueTable(data.TeamsTable)
	return constants.LeagueMessage{Table: leagueTable}
}

func (ui *UIModel) GetPlayerCmd() tea.Msg {
	id := ui.table.GetCursorValue()
	name, stats := getPlayerStats(id)
	return constants.PlayerMessage{PlayerName: name, Player: stats}
}

// getPlayerStats is a helper function to obtain stats for the selected player
func getPlayerStats(id string) (name, stats map[string]gjson.Result) {
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
	playerNameResponse := gjson.GetBytes(responseBody, "{\"firstName\":firstName.default,\"lastName\":lastName.default}")
	playerStatsResponse := gjson.GetBytes(responseBody, "featuredStats.regularSeason.subSeason")
	return playerNameResponse.Map(), playerStatsResponse.Map()
}
