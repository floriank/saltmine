package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func init() {
	db, _ = gorm.Open("sqlite3", "./saltmine_test.db")
}

func reset(db *gorm.DB) {
	db.DropTableIfExists(&Project{})
	db.CreateTable(&Project{})

	db.DropTableIfExists(&Ticket{})
	db.CreateTable(&Ticket{})
}
