package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// Version contains the current version injected via LD_FLAGS
// and derived from the git repository
var Version = "No version string injected"
var db *gorm.DB

func init() {
	db = Connect("./saltmine.db")
}

func main() {
	router := NewRouter()
	fmt.Printf("Running saltmine version %s\n", Version)
	log.Fatal(http.ListenAndServe(":8081", router))
}
