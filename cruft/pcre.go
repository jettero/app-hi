///usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"os"

	"github.com/rubrikinc/go-pcre"
)

func main() {
	re, err := pcre.CompileJIT("(ab)", 0, 0)
	fmt.Println("wtf", print)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	m := re.NewMatcher()
	if m.MatchString("Bababooey", 0) {
		fmt.Println("matched!", m.Groups())
		os.Exit(0)
	} else {
		fmt.Println("not matched")
		os.Exit(1)
	}
}
