package data

/*** TEAMS RESPONSE ***/

type TeamsResponse struct {
	Teams []TeamDefault `json:"teams"`
}

type TeamDefault struct {
	Id           uint             `json:"id"`
	TeamName     string           `json:"teamName"`
	LocationName string           `json:"locationName"`
	Division     []DivisionFields `json:"division"`
}

type DivisionFields struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

/*** TEAM STATS RESPONSE ***/

type StatsResponse struct {
	Stats []StatFields `json:"stats"`
}

type StatFields struct {
	Splits []struct {
		Stat struct {
			GamesPlayed            int     `json:"gamesPlayed"`
			Wins                   int     `json:"wins"`
			Losses                 int     `json:"losses"`
			Pts                    int     `json:"pts"`
			GoalsPerGame           float64 `json:"goalsPerGame"`
			GoalsAgainstPerGame    float64 `json:"goalsAgainstPerGame"`
			PowerPlayPercentage    string  `json:"powerPlayPercentage"`
			PowerPlayGoals         float64 `json:"powerPlayGoals"`
			PowerPlayGoalsAgainst  float64 `json:"powerPlayGoalsAgainst"`
			PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
			PenaltyKillPercentage  string  `json:"penaltyKillPercentage"`
			ShotsPerGame           float64 `json:"shotsPerGame"`
			ShotsAllowed           float64 `json:"shotsAllowed"`
			FaceOffsTaken          float64 `json:"faceOffsTaken"`
			FaceOffsWon            float64 `json:"faceOffsWon"`
			FaceOffsLost           float64 `json:"faceOffsLost"`
			FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
			ShootingPctg           float64 `json:"shootingPctg"`
			SavePctg               float64 `json:"savePctg"`
		} `json:"stat"`
	}
}

/*** ROSTER RESPONSE ***/

type RosterResponse struct {
	Teams []TeamRoster `json:"teams"`
}

type TeamRoster struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Roster struct {
		Players []Player `json:"roster"`
	} `json:"roster"`
}

type Player struct {
	Person struct {
		Id       int    `json:"id"`
		FullName string `json:"fullName"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position     struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}
