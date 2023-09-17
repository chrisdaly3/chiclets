package constants

import (
	"github.com/76creates/stickers/table"
)

type SelectionMessage struct {
	Value string
}

type RosterMessage struct {
	Table   *table.TableSingleType[string]
	Players [][]string
	Test    []string
}
