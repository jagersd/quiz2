package cli

import (
	"flag"
	"fmt"
	"quiz2/config"
	"quiz2/models"
)

func CheckFlags() bool {
	addData := flag.Bool("d", false, "Add default db data")
	dbMigrate := flag.Bool("m", false, "Migrate to new database")
	flag.Parse()
	if *addData == true {
		addDefaultData()
		return true
	} else if *dbMigrate == true {
		config.Roachconn(true)
		return true
	}

	return false
}

func addDefaultData() {
	config.Roachconn(false)

	fmt.Println("adding items to database")

	qtypes := []models.Qtype{
		{Description: "Multiple Choice"},
		{Description: "Open Question"},
	}
	config.DB.Create(&qtypes)
}
