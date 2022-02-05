package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")
	for i := 1; i < 3; i++ {
		fmt.Println(os.Args[i])
	}
}
