package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"quiz2/config"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := config.GetDbConn()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&Question{},
		&Aquiz{},
		&Option{},
		&Subject{},
		&Result{},
		&Qtype{},
	)

	DB = database
}
