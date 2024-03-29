package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Connect provides a simple function to connect to the database used
func Connect(filename string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filename)

	if err != nil {
		log.Fatalf("could not connect to database: %s", err.Error())
	}

	return &db
}
