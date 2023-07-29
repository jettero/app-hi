//go:build debug
// +build debug

package dfmt

import "fmt"

func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
