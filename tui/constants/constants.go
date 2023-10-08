package constants

const (
	BASEURL      = "https://statsapi.web.nhl.com/api/v1"
	TEAMSURL     = "https://statsapi.web.nhl.com/api/v1/teams"
	SCHEDULEURL  = "https://statsapi.web.nhl.com/api/v1/schedule"
	ROSTERURL    = "https://statsapi.web.nhl.com/api/v1/teams/%s?expand=team.roster"
	STANDINGSURL = "https://statsapi.web.nhl.com/api/v1/standings"
	LASTGAME     = "https://statsapi.web.nhl.com/api/v1/teams/%s?expand=team.schedule.previous"
	NEXTGAME     = "https://statsapi.web.nhl.com/api/v1/teams/%s?expand=team.schedule.next"
	PLAYERURL    = "https://statsapi.web.nhl.com/api/v1/people/%s/stats?stats=statsSingleSeason&season=20222023"
)

const (
	HOMEHEADER = `
   ______   __  __   ____   ______   __       ______   ______   _____
  / ____/  / / / /  /  _/  / ____/  / /      / ____/  /_  __/  / ___/
 / /      / /_/ /   / /   / /      / /      / __/      / /     \__ \ 
/ /___   / __  /  _/ /   / /___   / /___   / /___     / /     ___/ / 
\____/  /_/ /_/  /___/   \____/  /_____/  /_____/    /_/     /____/  
                                                                     
  `
)
