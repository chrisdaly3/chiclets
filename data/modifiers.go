package data

import (
	"github.com/tidwall/gjson"
)

var TeamsTable = sliceTeams(AllTeams) //TeamsTable is what we want to pass to table creation for the team

func sliceTeams(gj []gjson.Result) [][]string {

	var teamRows [][]string

	for _, t := range gj {
		teamAcronym := gjson.Get(t.Raw, "0").String()
		teamName := gjson.Get(t.Raw, "1").String()
		teamPoints := gjson.Get(t.Raw, "2").String()
		teamDivision := gjson.Get(t.Raw, "3").String()
		teamColumn := []string{teamAcronym, teamName, teamPoints, teamDivision}
		teamRows = append(teamRows, teamColumn)
	}
	return teamRows

}
