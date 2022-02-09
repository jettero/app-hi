package main

import (
	"fmt"
	"os"
	"strings"

	c "github.com/jettero/app-hi/pkg/colors"
)

func main() {
	a := strings.Join(os.Args[1:], " ")
	fmt.Println(a, "→", c.Color(a)+strings.Join(c.FixColor(a), " ")+c.Color("reset"))
}
