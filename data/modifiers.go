package data

import (
	"github.com/tidwall/gjson"
)

var TeamsTable = sliceTeams(AllTeams) //TeamsTable is what we want to pass to table creation for the team

func sliceTeams(gj []gjson.Result) [][]string {

	var teamRows [][]string

	for _, t := range gj {
		team := t.Map()
		teamColumn := []string{team["locationName"].Str, team["teamName"].Str, team["name"].Str}
		teamRows = append(teamRows, teamColumn)
	}
	return teamRows

}
