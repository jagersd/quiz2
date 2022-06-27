package cli

import (
	"flag"
	"fmt"
	"quiz2/config"
	"quiz2/models"
)

func CheckFlags() bool {
	addData := flag.Bool("d", false, "Add default db data")
	flag.Parse()
	if *addData == true {
		addDefaultData()
		return true
	}
	return false
}

func addDefaultData() {
	config.ConnectDatabase()

	fmt.Println("adding items to database")

	qtypes := []models.Qtype{
		{Description: "Multiple Choice"},
		{Description: "Open Question"},
	}
	config.DB.Create(&qtypes)
}