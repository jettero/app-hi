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

	// convert "a(ab)aba" "purple" to the patterns PrintLine uses
	patterns := patprint.ProcessPatterns(os.Args)

	for true {
		line, something := reader.ReadString('\n')
		patprint.PrintLine(patterns, line)
		if something == io.EOF {
			break
		}
	}
}
