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
	nextGame := ui.GetNextCmd()
	teamStandings := ui.GetRecordCmd()

	// Returns a message to the UI to update the table view
	return constants.TeamInfoMessage{Table: playerTable, TeamPriorGame: prevGame, TeamNextGame: nextGame, TeamStats: teamStandings}
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

// TODO: Setup next game and team record calls + parsers
func (ui *UIModel) GetNextCmd() map[string]gjson.Result {
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

	upcomingGame := gjson.GetBytes(responseBody, "{teams.0.nextGameSchedule.dates.0.date,teams.0.nextGameSchedule.dates.0.games.0.teams,teams.0.nextGameSchedule.dates.0.games.0.venue.name}.@join")
	awayTeam := gjson.Result.Get(upcomingGame, "{teams.away.team.name,teams.away.leagueRecord.wins,teams.away.leagueRecord.losses,teams.away.leagueRecord.ot}").Map()
	homeTeam := gjson.Result.Get(upcomingGame, "{teams.home.team.name,teams.home.leagueRecord.wins,teams.home.leagueRecord.losses,teams.home.leagueRecord.ot}").Map()

	nextGameMap := make(map[string]gjson.Result)
	nextGameMap["date"] = upcomingGame.Map()["date"]
	nextGameMap["away"] = awayTeam["name"]
	nextGameMap["awayWins"] = awayTeam["wins"]
	nextGameMap["awayLosses"] = awayTeam["losses"]
	nextGameMap["awayOTL"] = awayTeam["ot"]
	nextGameMap["home"] = homeTeam["name"]
	nextGameMap["homeWins"] = homeTeam["wins"]
	nextGameMap["homeLosses"] = homeTeam["losses"]
	nextGameMap["homeOTL"] = homeTeam["ot"]

	return nextGameMap

}

func (ui *UIModel) GetRecordCmd() map[string]gjson.Result {
	id := ui.table.GetCursorValue()
	requestPath := constants.STANDINGSURL

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

	records := gjson.GetBytes(responseBody, fmt.Sprintf("records.#.teamRecords.#(team.id == %v )", id))
	name := gjson.Result.Get(records, "0.team.name")
	points := gjson.Result.Get(records, "0.points")
	gamesPlayed := gjson.Result.Get(records, "0.gamesPlayed")
	divRank := gjson.Result.Get(records, "0.divisionRank")
	streak := gjson.Result.Get(records, "0.streak.streakCode")
	regWins := gjson.Result.Get(records, "0.regulationWins")
	row := gjson.Result.Get(records, "0.row")
	goalsFor := gjson.Result.Get(records, "0.goalsScored")
	goalsAgainst := gjson.Result.Get(records, "0.goalsAgainst")

	recordsMap := make(map[string]gjson.Result)
	recordsMap["teamName"] = name
	recordsMap["points"] = points
	recordsMap["gamesPlayed"] = gamesPlayed
	recordsMap["divRank"] = divRank
	recordsMap["streak"] = streak
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
