package main

import (
	"fmt"
)

// Version contains the current version injected via LD_FLAGS
// and derived from the git repository
var Version = "No version string injected"

func main() {
	fmt.Printf("Running saltmine version %s\n", Version)
}
