package patprint

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/zeebo/xxh3"
)

var ColorTable = map[string]string{
	"clear":      "0",
	"unbold":     "0",
	"bold":       "1",
	"black":      "30",
	"red":        "31",
	"green":      "32",
	"yellow":     "33",
	"blue":       "34",
	"magenta":    "35",
	"cyan":       "36",
	"white":      "37",
	"on_black":   "40",
	"on_red":     "41",
	"on_green":   "42",
	"on_yellow":  "43",
	"on_blue":    "44",
	"on_magenta": "45",
	"on_cyan":    "46",
	"on_white":   "47",
	"on_default": "49",
}

var NickTable = map[string]string{
	"normal":     "clear",
	"unbold":     "clear",
	"on_nothing": "on_default",

	"blood":     "red",
	"umber":     "bold red",
	"sky":       "bold blue",
	"ocean":     "cyan",
	"lightblue": "cyan",
	"cyan":      "bold cyan",
	"lime":      "bold green",
	"orange":    "yellow",
	"brown":     "yellow",
	"yellow":    "bold yellow",
	"purple":    "magenta",
	"violet":    "bold magenta",
	"pink":      "bold magenta",
	"pitch":     "bold black",
	"coal":      "bold black",
	"grey":      "white",
	"gray":      "white",
	"white":     "bold white",

	"dire":  "bold yellow on_red",
	"alert": "bold yellow on_red",
	"todo":  "black on_yellow",

	"nc_dir":  "bold white on_blue",
	"nc_file": "bold white on_blue",
	"nc_exe":  "bold green on_blue",
	"nc_exec": "bold green on_blue",
	"nc_curs": "black on_cyan",
	"nc_pwd":  "black on_white",
	"nc_cwd":  "black on_white",

	"mc_dir":  "bold white on_blue",
	"mc_file": "bold white on_blue",
	"mc_exe":  "bold green on_blue",
	"mc_exec": "bold green on_blue",
	"mc_curs": "black on_cyan",
	"mc_pwd":  "black on_white",
	"mc_cwd":  "black on_white",
}

var ickyTable = map[[3]int]bool{
	// This exclusion list can be pretty subjective Is it truely invisible like
	// \x1b[31;41m (red on red)?  or is it just hard to read like \x1b[0;34;40m
	// (dark blue on black)?
	//
	// Anything combinations that looked a little hard to read on my own
	// personal loonix xfce-terminal or on my fancy macos iterm2 console --
	// ECMA-48 colors look quite different between the two by the way -- I here
	// marked as "icky"; ymmv, but I'm trying to describe a generalized icky
	// factor for dark backgrounds here.
	//
	// I also assumed (not always true) black background (40) is very close or
	// identical to default background (49, often transparent) and therefore
	// turn any 40 bg into 49 bg. Black backgrounds can stand and even look
	// interesting on transparent terminals where the underlying background is
	// not black, but I'm assuming this is commonly not the case.

	// 40 black
	[3]int{0, 30, 40}: true,
	[3]int{0, 34, 40}: true,
	// 41 red
	[3]int{0, 31, 41}: true,
	[3]int{0, 35, 41}: true,
	[3]int{1, 30, 41}: true,
	[3]int{1, 34, 41}: true,
	[3]int{1, 35, 41}: true,
	// 42 green
	[3]int{0, 32, 42}: true,
	[3]int{0, 36, 42}: true,
	[3]int{0, 37, 42}: true,
	[3]int{1, 32, 42}: true,
	[3]int{1, 33, 42}: true,
	[3]int{1, 36, 42}: true,
	[3]int{1, 37, 42}: true,
	// 43 orange
	[3]int{0, 33, 43}: true, // almost visible in iterm2
	[3]int{0, 36, 43}: true,
	[3]int{0, 37, 43}: true,
	[3]int{1, 31, 43}: true,
	//[3]int{1, 34, 43}: true, // borderline
	[3]int{1, 35, 43}: true,
	// 44 blue
	[3]int{0, 34, 44}: true,
	[3]int{0, 30, 44}: true,
	// 45 magenta
	[3]int{0, 35, 45}: true,
	[3]int{0, 31, 45}: true,
	[3]int{1, 30, 45}: true, // horrible on iterm2, sorta ok in xfce
	[3]int{1, 34, 45}: true, // horrible on iterm2, sorta ok in xfce
	//[3]int{1, 35, 45}: true, // borderline
	// 46 cyan
	[3]int{0, 36, 46}: true,
	[3]int{0, 32, 46}: true,
	[3]int{0, 33, 46}: true,
	[3]int{0, 37, 46}: true,
	// 47 white
	[3]int{0, 37, 47}: true,
	[3]int{0, 32, 47}: true,
	[3]int{0, 33, 47}: true, // my iterm2 theme seems to treat non-bold yellow as an actual dim yellow
	[3]int{0, 36, 47}: true,
	[3]int{1, 32, 47}: true,
	[3]int{1, 36, 47}: true,
}

func FixColor(color string) []string {
	color = strings.ToLower(color)
	sep, _ := regexp.Compile(`[^a-z]+`)
	nc, _ := regexp.Compile(`[nm]c\s+([a-z]+)`)
	on, _ := regexp.Compile(`on\s+([a-z]+)`)
	un, _ := regexp.Compile(`un\s+bold`)
	color = sep.ReplaceAllString(color, " ")
	color = nc.ReplaceAllString(color, "mc_$1")
	color = on.ReplaceAllString(color, "on_$1")
	color = un.ReplaceAllString(color, "unbold")

	var ret []string
	for _, f := range strings.Fields(color) {
		if g := NickTable[f]; len(g) > 0 {
			ret = append(ret, strings.Fields(g)...)
		} else {
			ret = append(ret, f)
		}
	}

	return ret
}

func Color(color string, words string) string {
	var ret []string

	if color == "unique" || color == "hash" {
		return UniqueColorForString(words) + words
	}

	for _, f := range FixColor(color) {
		if v := ColorTable[f]; len(v) > 0 {
			ret = append(ret, v)
		}
	}

	return "\x1b[" + strings.Join(ret, ";") + "m" + words
}

var MostColorsTable []string

func buildMostColorsTable() {
	for bb := 0; bb < 2; bb++ {
		for bg := 40; bg < 48; bg++ {
			for fg := 30; fg < 38; fg++ {

				if !ickyTable[[3]int{bb, fg, bg}] {
					if bg == 40 {
						// We assume default background for most of this program; on most color
						// changes we reset with \x1b[0m anyway, so we needn't print the black
						// background we assume to be default (see note in ickyTable).
						MostColorsTable = append(MostColorsTable, fmt.Sprintf("\x1b[%d;%dm", bb, fg))
					} else {
						MostColorsTable = append(MostColorsTable, fmt.Sprintf("\x1b[%d;%d;%dm", bb, fg, bg))
					}
				}
			}
		}
	}
}

func UniqueColorForString(x string) string {
	if len(MostColorsTable) < 1 {
		buildMostColorsTable()
	}

	sum := int(xxh3.HashString(x) % uint64(len(MostColorsTable)))

	return MostColorsTable[sum]
}
