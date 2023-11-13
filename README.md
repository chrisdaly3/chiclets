# Chiclets
- A hockey stats TUI (project is still a work in progress)
- View all your favorite teams and players in the terminal with ease.
- **NOTE: AS OF `11.13.23` THIS PROJECT IS IN THE PROCESS OF BEING FIXED**
    - the NHL dropped support for the old API so in the meantime I need to fix a few of the calls this app relies on

## In Use:
```
go install github.com/chrisdaly3/chiclets@latest
chiclets
```

## TODO:
- [x] ~~Implement flexbox functionality~~
- [x] ~~Sort by columns on CAPS layer~~
- [x] ~~Call NHL API for initial table values~~
- [x] ~~Design and configure stats by team view~~
- [x] ~~Populate stats view with subsequent api calls depending on selected team.~~
- [x] ~~Override Table Style Defaults~~
- [x] ~~Set up back button from nested tables~~
- [x] ~~Display player information on side flex boxes (row 1).~~
- [x] ~~Display Team's stats on top row (row 0) center cell, prior game on left cell, and next upcoming game on right cell~~
- [ ] Config for favorite teams
- [ ] Show popup at player stats to allow season selection
- [ ] Add improved logging
- [ ] Clean up project structure (specifically data directory) 
- [ ] Fix ALL API calls and data parsing. League table is done, need to fix roster tables next
