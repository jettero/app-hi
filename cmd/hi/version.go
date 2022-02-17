// populated via version-izer.sh
//
// NOTE to self:
// the linker time names of these vars can be found like this:
//
//   go tool nm ./hi | grep -E "BuildTime|CommitHash"

package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

var (
	Version    = ""
	CommitHash = ""
	BuildTime  = ""
)

func PrintBuildVersion() {
	if Version == "" {
		fmt.Printf("--version doesn't work if you 'go run' or 'go install'. Use 'make' instead.\n")
		fmt.Printf("Here's some useless BuildInfo:\n")
	} else {
		fmt.Printf("Version:    %s\n", Version)
		fmt.Printf("CommitHash: %s\n", CommitHash)
		fmt.Printf("BuildTime:  %s\n", BuildTime)
	}

	if bi, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("\n")
		fmt.Printf("BuildInfo.Main.Path:    %s\n", (*bi).Main.Path)
		fmt.Printf("BuildInfo.Main.Version: %s\n", (*bi).Main.Version)
		fmt.Printf("BuildInfo.Main.Sum:     %s\n", (*bi).Main.Sum)
	}

	if Version == "" {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
