package main

import (
	"fmt"
	"log"
	"net/http"
)

// Version contains the current version injected via LD_FLAGS
// and derived from the git repository
var Version = "No version string injected"

func main() {
	router := NewRouter()
	fmt.Printf("Running saltmine version %s\n", Version)
	log.Fatal(http.ListenAndServe(":8080", router))
}
