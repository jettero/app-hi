package patprint

import (
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

	// note that we can't use "on red" here, we only get one pass through FixColor
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

var uniqueColorsTable = []string{
	"red",
	"green",
	"blue",
	"magenta",
	"ocean",
	"grey",
	"coal",
	"umber",
	"lime",
	"yellow",
	"sky",
	"violet",
	"cyan",
	"white",

	// note that we can't use "on red" here, we only get one pass through FixColor
	"coal on_blue",
	"umber on_blue",
	"white on_blue",
	"cyan on_blue",
	// "yellow on_blue", // yellow on blue is readable, but indistinguishable from white on blue
	// "lime on_blue", // same with cyan and lime

	"yellow on_red",
	"white on_red",
	"white on_magenta",
	"pitch on_white",
	"white on_black",
	"black on_yellow",
	"blue on_green",
}

func FixColor(color string, words string) []string {
	color = strings.ToLower(color)
	sep, _ := regexp.Compile(`[^a-z]+`)
	nc, _ := regexp.Compile(`[nm]c\s+([a-z]+)`)
	on, _ := regexp.Compile(`on\s+([a-z]+)`)
	un, _ := regexp.Compile(`un\s+bold`)
	color = sep.ReplaceAllString(color, " ")
	color = nc.ReplaceAllString(color, "mc_$1")
	color = on.ReplaceAllString(color, "on_$1")
	color = un.ReplaceAllString(color, "unbold")

	fields := strings.Fields(color)

	var ret []string
	for _, f := range fields {
		if f == "unique" || f == "hash" {
			sum := xxh3.HashString(words)
			idx := int(sum % uint64(len(uniqueColorsTable)))
			ret = append(ret, strings.Fields(uniqueColorsTable[idx])...)
			// as soon as we hit unique, we take over processing of the whole color token set
			break
		}
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

	for _, f := range FixColor(color, words) {
		if v := ColorTable[f]; len(v) > 0 {
			ret = append(ret, v)
		}
	}

	return "\x1b[" + strings.Join(ret, ";") + "m" + words
}
