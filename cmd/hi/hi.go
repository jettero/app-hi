package hi

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type pattern struct {
	pattern string
	color   string
	matcher *regexp.Regexp
}

func printLine(patterns []pattern, line string) {
	for i := 0; i < len(patterns); i++ {
		// Go documentation is a pile of shit. When you go a
		// `go doc regexp.Regxep.FindAllStringIndex` yo'll never find this 'n'
		// param anywhere. If you read the entire `go doc regexp.Regexp`
		// section, you still won't find it.
		//
		// You have to guess to read the entire package help and happen to
		// notice the brief desctiption in a small-ish paragraph in the middle
		// of the overview section.
		//
		// … the -1 means we want them all, not 0 or 1 or 2 …
		indices := patterns[i].matcher.FindAllStringIndex(line, -1)
		fmt.Printf("pat: %v, indices: %v\n", patterns[i], indices)
	}
	fmt.Println(line)
}

// Execute runs the actual program (meant to be used in main.go)
func Execute() {
	reader := bufio.NewReader(os.Stdin)
	var patterns []pattern

	for i := 1; i < len(os.Args); i += 2 {
		p := pattern{
			pattern: os.Args[i],
			matcher: regexp.MustCompile(os.Args[i]),
			color:   os.Args[i+1],
		}
		patterns = append(patterns, p)
	}

	for true {
		line, something := reader.ReadString('\n')

		if something == io.EOF {
			if len(line) > 0 {
				printLine(patterns, line)
			}
			break
		}
		printLine(patterns, line)
	}
}
