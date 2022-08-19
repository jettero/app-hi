///usr/bin/env go run $0 $@ ; exit $?

package main

import (
	"fmt"
	"os"

	"go.arsenm.dev/pcre"
)

func main() {
	re, err := pcre.Compile("ab(ab)ab")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	matches := re.FindStringSubmatch("ababab")
	if len(matches) > 0 {
		fmt.Println("matched!", matches)
		os.Exit(0)
	} else {
		fmt.Println("not matched")
		os.Exit(1)
	}
}
