package db

import (
	"cronbackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB(config *Config) {
	dsn := config.ConnectionString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if config.Migrate {
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			panic("failed to migrate database")
		}

		err = db.AutoMigrate(&models.ScheduleCluster{})
		if err != nil {
			panic("failed to migrate database")
		}

		err = db.AutoMigrate(&models.ScheduleJob{})
		if err != nil {
			panic("failed to migrate database")
		}

		err = db.AutoMigrate(&models.CheckJobs{})
		if err != nil {
			panic("failed to migrate database")
		}
	}

	DB = db
}
