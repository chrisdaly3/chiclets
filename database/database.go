package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormTable holds the gorm DB accessors for each table.
type GormTable struct {
	teamDB   *gorm.DB
	playerDB *gorm.DB
}

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("NHL.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	log.Println("DB Connection Successful")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating Models")
	if err := db.AutoMigrate(&Team{}, &Player{}); err != nil {
		panic("Failed to migrate DB models")
	}
	log.Println("Models migrated successfully, ")
}
