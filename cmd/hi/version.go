// populated via version-izer.sh
//
// NOTE to self:
// the linker time names of these vars can be found like this:
//
//   go tool nm ./hi | grep -E "BuildTime|CommitHash"

package main

import (
	"fmt"
	"runtime/debug"
)

var (
	Version    = "who fucking knows"
	CommitHash = "there's no way to embed this in a go program without a makefile or something"
	BuildTime  = "you got me man, use the makefile, go build sucks"
)

func PrintBuildVersion() {
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("CommitHash: %s\n", CommitHash)
	fmt.Printf("BuildTime:  %s\n", BuildTime)

	if bi, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("\n")
		fmt.Printf("BuildInfo.Main.Path:    %s\n", (*bi).Main.Path)
		fmt.Printf("BuildInfo.Main.Version: %s\n", (*bi).Main.Version)
		fmt.Printf("BuildInfo.Main.Sum:     %s\n", (*bi).Main.Sum)
	}
}
