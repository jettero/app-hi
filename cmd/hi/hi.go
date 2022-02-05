package hi

import (
	"bufio"
	"io"
	"os"

	"github.com/jettero/app-hi/pkg/patprint"
)

// Execute runs the actual program (meant to be used in main.go)
func Execute() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		line, something := reader.ReadString('\n')
		patterns := patprint.ProcessPatterns(os.Args)

		if something == io.EOF {
			if len(line) > 0 {
				patprint.PrintLine(patterns, line)
			}
			break
		}
		patprint.PrintLine(patterns, line)
	}
}
