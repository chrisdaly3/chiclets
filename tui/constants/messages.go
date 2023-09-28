package constants

import (
	"github.com/76creates/stickers/table"
	"github.com/tidwall/gjson"
)

type SelectionMessage struct {
	Value string
}

type RosterMessage struct {
	Table *table.TableSingleType[string]
}

type LeagueMessage struct {
	Table *table.TableSingleType[string]
}

type PlayerMessage struct {
	Player map[string]gjson.Result
}
