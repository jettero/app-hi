package patprint

import (
	"regexp"
	"strings"
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

func FixColor(color string) []string {
	nc, _ := regexp.Compile(`[nm]c\s+([a-z]+)`)
	on, _ := regexp.Compile(`on\s+([a-z]+)`)
	un, _ := regexp.Compile(`un\s+bold`)
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

func Color(color string) string {
	var ret []string
	for _, f := range FixColor(color) {
		if v := ColorTable[f]; len(v) > 0 {
			ret = append(ret, v)
		}
	}

	return "\x1b[" + strings.Join(ret, ";") + "m"
}
