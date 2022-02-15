// populated via version-izer.sh
//
// NOTE to self:
// the linker time names of these vars can be found like this:
//
//   go tool nm ./hi | grep -E "BuildTime|CommitHash"

package main

import (
	"fmt"
)

var (
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

func PrintBuildVersion() {
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("CommitHash: %s\n", CommitHash)
	fmt.Printf("BuildTime:  %s\n", BuildTime)
}
