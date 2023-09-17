package tui

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chrisdaly3/chiclets/data"
	"github.com/chrisdaly3/chiclets/tui/constants"
	"io"
	"log/slog"
	"net/http"
	"os"
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

	var Roster data.RosterResponse
	json.Unmarshal(responseBody, &Roster)

	//FIX: TestData for now. Should eventually populate with
	// newTeamTable data
	return constants.RosterMessage{Test: Roster.Teams[0].Name}
}
