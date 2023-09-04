package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chrisdaly3/chiclets/database"
	"github.com/chrisdaly3/chiclets/tui"
	"github.com/chrisdaly3/chiclets/tui/constants"
	"github.com/tidwall/gjson"
)

func main() {
	allTeams := getTeams()
	database.ConnectDB()

	p := tea.NewProgram(&tui.InitModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error initializing tui: %v", err)
	}

	// TODO: this is just in place for debug.
	// I will remove it as I progress and can validate data
	// in the db itself.
	Flyers := allTeams[3].Get("@values")
	fmt.Println(Flyers)
	//CreateTeamDB(allTeams)
}

func getTeams() []gjson.Result {
	var teams []gjson.Result

	response, err := http.Get(constants.TeamsURL)
	if err != nil {
		fmt.Printf("Error communicating with NHL API: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		slog.Error("Unhealthy Response", "response", response.Body)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Failure reading json response body")
		os.Exit(1)
	}

	// Check response for number of teams
	responseTeamCount := gjson.GetBytes(responseBody, "teams.#")

	// Use result of team number to iterate all team details
	for i := 0; int64(i) < responseTeamCount.Int(); i++ {
		teamIndex := fmt.Sprintf("teams.%v", i)
		gjFields := fmt.Sprintf("{%s.teamName,%s.locationName,%s.division.name}.@join", teamIndex, teamIndex, teamIndex)

		team := gjson.GetBytes(responseBody, gjFields)
		teams = append(teams, team)
	}
	return teams
}

/*func CreateTeamDB(results []gjson.Result) error {

}*/
