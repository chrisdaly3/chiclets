package database

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name   string
	Locale string
	Record string
	Roster []Player
}

type Player struct {
	gorm.Model
	FirstName string
	LastName  string
	Number    int
	Position  string
	TeamID    uint
	Team      Team
}
