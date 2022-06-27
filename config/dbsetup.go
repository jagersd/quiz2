package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"quiz2/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := GetDbConn()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&models.Question{},
		&models.Aquiz{},
		&models.Option{},
		&models.Subject{},
		&models.Result{},
		&models.Qtype{},
	)

	DB = database
}
