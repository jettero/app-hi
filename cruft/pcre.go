///usr/bin/env go run $0 $@ ; exit $?

package main

import (
	"fmt"
	"os"

	"go.elara.ws/pcre"
)

func main() {
	re, err := pcre.Compile("ab(ab)ab")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	smatches := re.FindStringSubmatch("abababababab")
	fmt.Println("s-matches:", smatches)

	imatches := re.FindStringSubmatchIndex("abababababab")
	fmt.Println("i-matches:", imatches)

	amatches := re.FindAllStringSubmatchIndex("abababababab", -1)
	fmt.Println("a-matches:", amatches)

	nsmatches := re.FindStringSubmatch("xxxxx")
	fmt.Println("!s-matches:", nsmatches)

	nimatches := re.FindStringSubmatchIndex("xxxxx")
	fmt.Println("!i-matches:", nimatches)

	namatches := re.FindAllStringSubmatchIndex("xxxxx", -1)
	fmt.Println("!a-matches:", namatches)

	os.Exit(0)
}
