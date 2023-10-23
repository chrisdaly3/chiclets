package constants

const (
	BASEURL      = "https://statsapi.web.nhl.com/api/v1"
	TEAMSURL     = "https://statsapi.web.nhl.com/api/v1/teams"
	SCHEDULEURL  = "https://statsapi.web.nhl.com/api/v1/schedule"
	ROSTERURL    = "https://statsapi.web.nhl.com/api/v1/teams/%s?expand=team.roster"
	STANDINGSURL = "https://statsapi.web.nhl.com/api/v1/standings"
	LASTGAME     = "https://statsapi.web.nhl.com/api/v1/teams/%s?expand=team.schedule.previous"
	NEXTGAME     = "https://statsapi.web.nhl.com/api/v1/teams/%s?expand=team.schedule.next"
	PLAYERURL    = "https://statsapi.web.nhl.com/api/v1/people/%s/stats?stats=statsSingleSeason&season=20232024"
)

const (
	HOMEHEADER = "\n\n\nWelcome\nTo\nChiclets!\n\nFind Your Favorite Team\nTo Get Started"
)
