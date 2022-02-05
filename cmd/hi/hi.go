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
		fmt.Printf("pat: %v\n", patterns[i])
		line = patterns[i].matcher.ReplaceAllStringFunc(line, func(m string) string {
			return fmt.Sprintf("[%s]%s[/%s]", patterns[i].color, m, patterns[i].color)
		})
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
