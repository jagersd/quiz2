package config

import (
	"fmt"
	"log"
	"quiz2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Roachconn(migrate bool) {
	dsn := getRoachConn()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	} else {
		fmt.Println("connected to database!")
	}

	if migrate {
		fmt.Println("Executing database migration")
		database.AutoMigrate(
			&models.Question{},
			&models.Aquiz{},
			&models.Option{},
			&models.Subject{},
			&models.Result{},
			&models.Qtype{},
		)
	}

	DB = database

}
