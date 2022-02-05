package hi

import (
	"bufio"
	"fmt"
	"os"
)

// Execute runs the actual program (meant to be used in main.go)
func Execute() {
	reader := bufio.NewReader(os.Stdin)
	line, something := reader.ReadString('\n')

	fmt.Println("line: ")
	fmt.Println(line)
	fmt.Println("something: ")
	fmt.Println(something)
}
