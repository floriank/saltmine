package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Version contains the current version injected via LD_FLAGS
// and derived from the git repository
var Version = "No version string injected"
var db *gorm.DB

func init() {
	db = Connect("./saltmine.db")
}

func main() {
	fmt.Println("Running saltmine Version %s", Version)
}
