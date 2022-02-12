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

	RST := c.Color("reset", "")

	fmt.Println(a, "→", c.Color(a, a)+strings.Join(c.FixColor(a), " ")+RST)

	for _, s := range []string{"vim-go", "test", "test", "test1", "test2", "test3"} {
		fmt.Printf("%s%s%s\n", c.UniqueColorForString(s), s, RST)
	}

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("test%d", i)
		fmt.Printf("%s%s%s\n", c.UniqueColorForString(s), s, RST)
	}
}
