package constants

import (
	"github.com/76creates/stickers/table"
	"github.com/tidwall/gjson"
)

type SelectionMessage struct {
	Value string
}

type TeamInfoMessage struct {
	Table     *table.TableSingleType[string]
	TeamStats map[string]gjson.Result
}

type LeagueMessage struct {
	Table *table.TableSingleType[string]
}

type PlayerMessage struct {
	PlayerName map[string]gjson.Result
	Player     map[string]gjson.Result
}
