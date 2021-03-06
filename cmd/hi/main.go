package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"text/tabwriter"

	c "github.com/jettero/app-hi/pkg/colors"
	"github.com/jettero/app-hi/pkg/patprint"
	"github.com/spf13/pflag"
)

var examples [][]string = [][]string{
	[]string{"blue on_green", "black on_yellow", "black on_cyan", "black on_white"},
	[]string{"nc_cwd", "nc_curs", "mc_pwd", "todo"},
	[]string{"nc_pwd", "mc_cwd", "mc_curs", "white on_black"},
	[]string{"white on_red", "white on_blue", "white on_magenta", "pitch on white"},
	[]string{"nc_exe", "mc_exec", "mc_dir", "mc_file"},
	[]string{"alert", "nc_file", "nc_dir", "nc_exec"},
	[]string{"dire", "mc_exe", "black", "red"},
	[]string{"green", "blue", "magenta", "purple"},
	[]string{"gray", "ocean", "lightblue", "grey"},
	[]string{"brown", "blood", "orange", "bold black"},
	[]string{"bold red", "bold green", "bold yellow", "bold blue"},
	[]string{"bold magenta", "bold cyan", "bold white", "yellow"},
	[]string{"cyan", "white", "umber", "pitch"},
	[]string{"sky", "lime", "pink", "coal"},
	[]string{"violet", "normal"},
}

func PrintHelp(exitval int) {
	b := bytes.NewBufferString("\nUSAGE: hi [<regex> <color>]+\n")
	pflag.CommandLine.SetOutput(b)
	pflag.PrintDefaults()

	fmt.Fprintf(b, "\nExample colors (not an exhaustive list):\n")

	// NOTE to self: In the examples, tabwriter uses tabs which
	// absofuckinglutely suck ass (like the rest of tabs in Golang and any
	// other language). You will be 100% completely and totally unable to at
	// any point fucking like ever get this table to line up at all using '\t'
	// as the pad character -- it just will fucking not work.
	//
	// Spaces are 100% automatic. Ignoring examples.
	//
	// later:
	//   Turns out, the tabwriter doc actually uses spaces in the examples
	//   (mostly anyway) seems like they know tabs suck too. Perhaps "tab
	//   writer" refers to the use of tab as a delimiter (appropriate use)
	//   rather than the thing being written as something the terminal should
	//   manage (stupid inconsistent broken use).
	w := tabwriter.NewWriter(b, 0, 0, 1, ' ', 0)

	var items []string
	for _, row := range examples {
		fmt.Fprintf(w, "."+strings.Join(row, ".\t.")+".\n")
		for _, item := range row {
			items = append(items, item)
		}
	}

	w.Flush()

	sort.Slice(items, func(i, j int) bool {
		return len(items[i]) > len(items[j])
	})
	RST := c.Color("reset", "")
	s := b.String()
	for _, item := range items {
		s = strings.Replace(s, fmt.Sprintf(".%s.", item), fmt.Sprintf("%s%s", c.Color(item, item), RST), -1)
	}

	fmt.Println(s)

	if exitval >= 0 {
		os.Exit(exitval)
	}
}

func pflagErrorUsage() {
	PrintHelp(-1)
}

func main() {
	pflag.Usage = pflagErrorUsage
	Args := ProcessConfigAndArgs()

	if len(Args) == 0 {
		PrintHelp(0)
	}

	if (len(Args) % 2) != 0 {
		os.Stderr.WriteString("ERROR: odd number of arguments")
		PrintHelp(1)
	}

	if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice != 0 {
		os.Stderr.WriteString("ERROR: stdin appears to be an interactive shell. Pipe something in instead")
		PrintHelp(1)
	}

	patterns := patprint.ProcessPatterns(Args)
	reader := bufio.NewReader(os.Stdin)

	for true {
		line, err := reader.ReadString('\n')
		patprint.PrintLine(patterns, line)
		if err == io.EOF {
			break
		}
	}

	// We're exiting anyway, so I'm not sure it's necessary to go free those
	// regex memories, but it won't hurt either.
	patprint.PostProcessPatterns(patterns)
}
