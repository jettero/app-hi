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
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

func PrintBuildVersion() {
	if version != "" {
		fmt.Printf("Makefile versioning:\n")
		fmt.Printf("  Version: %s\n", version)
		fmt.Printf("  Commit:  %s\n", commit)
		fmt.Printf("  Date:    %s\n", date)
		fmt.Printf("  Builder: %s\n", builtBy)

	} else {
		fmt.Printf("Makefile versioning unavailable.\n")
	}

	if bi, ok := debug.ReadBuildInfo(); ok && (*bi).Main.Version != "(devel)" {
		fmt.Printf("Go Module Proxy versioning:\n")
		fmt.Printf("  BuildInfo.Main.Path:    %s\n", (*bi).Main.Path)
		fmt.Printf("  BuildInfo.Main.Version: %s\n", (*bi).Main.Version)
		fmt.Printf("  BuildInfo.Main.Sum:     %s\n", (*bi).Main.Sum)

	} else {
		fmt.Printf("Go Module Proxy versioning unavailable.\n")
	}

}
