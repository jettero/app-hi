///usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"os"
	"strings"

	c "github.com/jettero/app-hi/pkg/colors"
)

// this nonsense is just whatever I happened to be working on last
func main() {
	a := strings.Join(os.Args[1:], " ")
	if a == "" {
		a = "yellow on blue"
	}

	fmt.Println(a, "→", c.Color(a, strings.Join(c.FixColor(a), " "))+c.Color("reset", ""))
}
