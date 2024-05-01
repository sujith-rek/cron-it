package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// DB is the database connection
var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}




