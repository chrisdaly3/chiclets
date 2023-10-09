## !! As of `Oct. 8, 2023` I have discovered the NHL has created a new API !!
~~- I am unsure what this means for this project as it stands.~~
~~- I am hoping the NHL continues to support their prior API, for which this project has been designed around.~~
---> EDIT: it appears the NHL will **NOT** support the old api anymore in the new season. 
- I am tracking these changes closely, many thanks to the folks who are uncovering and documenting the endpoints throughout the community.
- I hope to get this project in working status by the time the `2023-2024` season begins, as I do not suspect the changes to be that difficult to implement, I just need to know how to call the new API.

# Chiclets
- A hockey stats TUI (work in progress)
- View all your favorite teams and players in the terminal with ease.

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
 
