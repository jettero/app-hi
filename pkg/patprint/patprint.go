package patprint

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type pattern struct {
	pattern string
	color   string
	matcher *regexp.Regexp
}

type annotation struct {
	color string
	start int
	stop  int
}

func generateAnnotations(color string, indices [][]int) []annotation {
	var ret []annotation
	for i := 0; i < len(indices); i++ {
		start, stop := indices[i][0], indices[i][1]
		did_something := false
		if len(ret) > 0 {
			if ret[len(ret)-1].stop == start {
				ret[len(ret)-1].stop = stop
				did_something = true
			}
		}
		if !did_something {
			ret = append(ret, annotation{color: color, start: start, stop: stop})
		}
	}
	return ret
}

func ProcessPatterns(args []string) []pattern {
	var patterns []pattern

	for i := 1; i < len(args); i += 2 {
		p := pattern{
			pattern: args[i],
			matcher: regexp.MustCompile(args[i]),
			color:   args[i+1],
		}
		patterns = append(patterns, p)
	}

	return patterns
}

func stripLineEndings(line string) string {
	l := 0
	for len(line) != l {
		l = len(line)
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		// if we removed something we'll loop through here again in case there's more
		// otherwise fall through, we're done
	}
	return line
}

func printRuler(line string) {
	rline := []rune(line)

	ruler_l0 := ""
	ruler_l1 := ""

	a := 0
	for i := 0; i < len(rline); i++ {
		v := strconv.Itoa(a)
		if (a % 10) == 0 {
			d := len([]rune(ruler_l1)) - len(ruler_l0)
			if d > 0 {
				ruler_l0 += strings.Repeat(" ", d)
			}
			ruler_l0 += v
			ruler_l1 += "↓"

		} else {
			ruler_l1 += v[len(v)-1:]
		}

		s := string(rline[i : i+1])
		a += len(s)
	}
	fmt.Println()
	if len(ruler_l0) > 0 {
		fmt.Println(ruler_l0)
	}
	if len(ruler_l1) > 0 {
		fmt.Println(ruler_l1)
	}
}

func PrintLine(patterns []pattern, line string) {
	line = stripLineEndings(line)

	if len(line) < 1 {
		return
	}

	for i := 0; i < len(patterns); i++ {
		// Go documentation (perhaps only regexp documentation) is a pile of
		// shit. When you go a `go doc regexp.Regxep.FindAllStringIndex` yo'll
		// never find this 'n' param anywhere. If you read the entire `go doc
		// regexp.Regexp` section, you still won't find it.
		//
		// You have to guess to read the entire package help and happen to
		// notice the brief desctiption in a small-ish paragraph in the middle
		// of the overview section.
		//
		// … the -1 means we want them all, not 0 or 1 or 2 …
		indices := patterns[i].matcher.FindAllStringIndex(line, -1)
		annotations := generateAnnotations(patterns[i].color, indices)
		fmt.Printf("pat: %v, annotations: %v\n", patterns[i], annotations)
	}

	printRuler(line)
	fmt.Println(line)
}
