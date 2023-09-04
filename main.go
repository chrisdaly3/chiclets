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
	database.ConnectDB()
	allTeams := getTeams()

	p := tea.NewProgram(&tui.InitModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error initializing tui: %v", err)
	}
	fmt.Println(allTeams)
}

func getTeams() gjson.Result {

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

	teams := gjson.Get(string(responseBody), "teams.#.name")
	return teams
}
