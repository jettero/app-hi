package patprint

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	c "github.com/jettero/app-hi/pkg/colors"
	"github.com/rubrikinc/go-pcre" // pcre.*
	// The built in regexp.* in Golang is absolutely awful.
	//
	// The author's failed quest to understand backtracking lead him to write a
	// stupid dumb NFA on DFA O(n) library without any backtracking. The O(n)
	// goal is worth persuing I think, but modern RE engine authors leave it to
	// the programmers to avoid time complexity bombs in their regular
	// expressions.
	//
	// In addition to being stupid and anachronistic the package is also
	// provably less performant than pcre, and (I'm told) completely falls over
	// trying to deal with a megabyte of data — though I didn't follow the
	// particulars of that story, so perhaps it's not as bad as I'm claiming
	// here.
	//
	// Regardless, even if the all of the above is completely false; perl,
	// python, pike, ruby, php, grep, sed, awk, vim and even shell programmers
	// are going to expect modern regular expressions with backtracking, so Go
	// can go climb up its own ass and cram it in harder with all these
	//
	// FuckingTabCharacters
)

type pattern struct {
	pattern string
	color   string
	matcher pcre.Regexp
}

type annotation struct {
	color string
	start int
	stop  int
}

func generateAnnotations(color string, matches []pcre.Match) []annotation {
	var ret []annotation
	for _, m := range matches {
		start, stop := m.Loc[0], m.Loc[1]
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

	for i := 0; i+1 < len(args); i += 2 {
		re, err := pcre.CompileJIT(args[i], 0, 0)
		// CompileJIT does a Compile and a Study. It's expensive, but probably
		// makes the text processing much faster.
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("ERROR compiling pattern \"%s\": %v\n", args[i], err))
			continue
		}
		p := pattern{
			pattern: args[i],
			matcher: re,
			color:   args[i+1],
		}
		patterns = append(patterns, p)
	}

	return patterns
}

func PostProcessPatterns(patterns []pattern) {
	for _, p := range patterns {
		p.matcher.FreeRegexp()
	}
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

func PrintRuler(line string) {
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

func combineAnnotationsStack(annotations_stack [][]annotation) []annotation {
	min := int(^uint(0) >> 1)
	max := 0

	// flatten/collect all the annoatations from the stack
	var annotations []annotation
	for _, asi := range annotations_stack {
		for _, a := range asi {
			if a.start < min {
				min = a.start
			}
			if a.stop > max {
				max = a.stop
			}
			annotations = append(annotations, a)
		}
	}

	// prepend a reset code to the beginning of the stack
	// in python we could just
	//   annotations.insert(0, annotation(color="reset", ...))
	// I wonder if there's a more concise way to say this in golang, cuz the
	// following sucks. (I don't need it anyway, but I'm leaving it here until
	// I get sick of looking at it or find an answer to the conciseness
	// question)
	// annotations = append([]annotation{annotation{color: "reset", start: min, stop: max}}, annotations...)

	var combined []annotation
	for i := min; i < max; i++ {
		var winner annotation
		for _, a := range annotations {
			if i >= a.start && i < a.stop {
				winner = a
			}
		}

		if i >= winner.start && i < winner.stop {
			if len(combined) == 0 {
				if winner.color != "reset" {
					combined = append(combined, annotation{color: winner.color, start: i, stop: i + 1})
				}
			} else {
				r := &combined[len(combined)-1]
				if r.color == winner.color && r.stop == i {
					r.stop++
				} else {
					combined = append(combined, annotation{color: winner.color, start: i, stop: i + 1})
				}
			}
		}
	}

	return combined
}

func fakeColor(color string, words string) string {
	return fmt.Sprintf("[%s]%", strings.Join(c.FixColor(color, words), " "), words)
}

func ColorizeLine(line string, annotations []annotation) string {
	var pos int = 0
	var ret string = ""

	cf := c.Color
	if os.Getenv("DEBUG_HI") == "1" || os.Getenv("DEBUG_HI_PATPRINT") == "1" || os.Getenv("DEBUG_HI_MARKUP") == "1" {
		cf = fakeColor
	}
	RST := cf("reset", "")

	for _, a := range annotations {
		if a.start > pos {
			ret += line[pos:a.start]
			pos = a.start
		}

		ret += cf(a.color, line[a.start:a.stop])
		ret += RST
		pos = a.stop
	}

	if pos < len(line) {
		ret += line[pos:]
	}

	return ret
}

func PrintLine(patterns []pattern, line string) {
	line = stripLineEndings(line)

	if len(line) < 1 {
		return
	}

	var annotations_stack [][]annotation

	for i := 0; i < len(patterns); i++ {
		matches, err := patterns[i].matcher.FindAll(line, 0)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("ERROR matching with pattern \"%s\": %v\n",
				patterns[i].pattern, err))
			continue
		}

		a := generateAnnotations(patterns[i].color, matches)
		if os.Getenv("DEBUG_HI") == "1" || os.Getenv("DEBUG_HI_PATPRINT") == "1" {
			fmt.Printf("pat: %v => %v\n", patterns[i], a)
		}
		annotations_stack = append(annotations_stack, a)
	}

	combined := combineAnnotationsStack(annotations_stack)

	if os.Getenv("DEBUG_HI") == "1" || os.Getenv("DEBUG_HI_PATPRINT") == "1" {
		fmt.Printf("combineAnnotationsStack():\n\tannotations_satck = %v\n\tcombined = %v\n",
			annotations_stack, combined)
		PrintRuler(line)
	}

	line = ColorizeLine(line, combined)

	fmt.Println(line)
}
