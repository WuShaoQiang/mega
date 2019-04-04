package model

import (
	"log"

	"github.com/WuShaoQiang/mega/config"
	"github.com/jinzhu/gorm"

	//Mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connet to db...")
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}
