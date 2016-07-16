package main

import (
	"github.com/atclate/goa-exercise/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "/goa_exercise")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(models.Cache{})
	db.AutoMigrate(models.Source{})
	database = db
}
