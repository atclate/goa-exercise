//go:generate goagen bootstrap -d github.com/atclate/goa-exercise/design

package main

import (
	"fmt"

	"github.com/atclate/goa-exercise/app"
	"github.com/atclate/goa-exercise/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	sdb *models.SourceDB
	cdb *models.CacheDB
)

func init() {
	db, err := gorm.Open("mysql", "root:root@/goa_exercise?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		sdb = models.NewSourceDB(db)
		cdb = models.NewCacheDB(db)
	} else {
		panic(fmt.Sprintf("Unable to initialize database. Error: %v\n", err.Error()))
	}
}

func main() {
	// Create service
	service := goa.New("GoWorkshop")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "cache" controller
	c := NewCacheController(service)
	app.MountCacheController(service, c)
	// Mount "public" controller
	c2 := NewPublicController(service)
	app.MountPublicController(service, c2)
	// Mount "source" controller
	c3 := NewSourceController(service)
	app.MountSourceController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
