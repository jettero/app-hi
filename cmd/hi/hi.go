package hi

import (
	"fmt"
	"os"
)

func Execute() {
	fmt.Println("hi")
	for i := 1; i < 3; i++ {
		fmt.Println(os.Args[i])
	}
}
