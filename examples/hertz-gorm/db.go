package main

import (
	sqlcommentergorm "github.com/a631807682/sqlcommenter-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func dbInit() {
	dbDSN := "sqlcommenter:sqlcommenter@tcp(localhost:8910)/sqlcommenter?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&Test{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Test{})
	if err != nil {
		panic(err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)
	db.Use(sqlcommentergorm.New(&sqlcommentergorm.Config{
		Application: "example-hertz-gorm",
	}))
	DB = db
}

type Test struct {
	gorm.Model
	Name string
}
