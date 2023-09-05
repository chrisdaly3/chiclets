package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB(resp NHLResponse) (int64, error) {
	teams := &resp.Teams
	db, err := gorm.Open(sqlite.Open("NHL.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	log.Println("DB Connection Successful")
	file, err := os.Create("gorm.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newDBLogger := logger.New(
		log.New(file, "", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Enable color
		},
	)

	db.Logger = newDBLogger
	log.Println("Migrating Models")
	if err := db.AutoMigrate(&Team{}, &Conference{}, &Franchise{}, &TimeZone{}, &Division{}, &Venue{}); err != nil {
		panic("Failed to migrate DB models")
	}
	log.Println("Models migrated successfully")

	result := db.Create(teams)
	if result.Error != nil {
		return 0, result.Error
	}

	Database = DbInstance{Db: db}

	return db.RowsAffected, nil

}
