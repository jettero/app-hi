package hi

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Execute runs the actual program (meant to be used in main.go)
func Execute() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		line, something := reader.ReadString('\n')

		if something == io.EOF {
			if len(line) > 0 {
				fmt.Println(line)
			}
			break
		}

		fmt.Println(line)
	}
}
