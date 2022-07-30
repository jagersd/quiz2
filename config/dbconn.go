package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDbConn() string {
	type DbConn struct {
		driver string
		user   string
		pass   string
		dsn    string
		dbname string
	}

	cfile, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatal(err)
	}

	c := DbConn{
		cfile.Section("dbconn").Key("driver").String(),
		cfile.Section("dbconn").Key("user").String(),
		cfile.Section("dbconn").Key("pass").String(),
		cfile.Section("dbconn").Key("dsn").String(),
		cfile.Section("dbconn").Key("dbname").String(),
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.user, c.pass, c.dsn, c.dbname)
}

func getRoachConn() string {
	cfile, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatal(err)
	}
	return cfile.Section("dbconn").Key("roachstring").String()
}
