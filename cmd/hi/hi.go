package hi

import (
	"fmt"
	"os"
)

func Execute() {
	fmt.Println("hi")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
