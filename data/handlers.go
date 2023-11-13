package data

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/chrisdaly3/chiclets/tui/constants"
	"github.com/tidwall/gjson"
)

var AllTeams = getTeams()

func getTeams() []gjson.Result {
	var teams []gjson.Result

	response, err := http.Get(constants.TEAMSURL)
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

	// Check response for number of teams
	responseTeamCount := gjson.GetBytes(responseBody, "standings.#")

	// Use result of team number to iterate all team details
	for i := 0; int64(i) < responseTeamCount.Int(); i++ {
		teamIndex := fmt.Sprintf("standings.%v", i)
		gjFields := fmt.Sprintf("{%s.teamAbbrev.default,%s.teamName.default,%s.points,%s.divisionName}.@values", teamIndex, teamIndex, teamIndex, teamIndex)

		team := gjson.GetBytes(responseBody, gjFields)
		teams = append(teams, team)
	}
	return teams
}
