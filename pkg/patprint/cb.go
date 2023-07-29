package patprint

import (
	"strconv"
	"strings"

	"encoding/json"
	"github.com/jettero/app-hi/pkg/dfmt"
	"go.elara.ws/pcre"
)

func matched(b bool) int32 {
	if b {
		dfmt.Printf(" match=yes\n")
		return 0
	}
	dfmt.Printf(" match=no\n")
	return 1
}

func MahCallback(cb *pcre.CalloutBlock) int32 {
	// NOTE: There's some DEBUG_HI Print items elsewhere in this file, but for
	// callouts -- which could get executed frequently -- I chose to only
	// enable them at compile time when needed

	if cb.CalloutNumber == 0 {
		// TODO: I'd really prefer to only parse the callback comparison float
		// once instead of every single time the callout gets called out. I
		// have ideas on how to get there, but nothing worth writing down. I
		// wanna get it this all working first then circle back to it.
		scs := strings.Split(cb.CalloutString, " ")
		midx := len(cb.Substrings) - 1
		if len(scs) == 2 && midx >= 0 {
			j, err := json.Marshal(&cb)
			if err == nil {
				dfmt.Printf("cb: %s\n", j)
			}
			dfmt.Printf("scs: %+v midx: %d", scs, midx)
			cb_v, err := strconv.ParseFloat(scs[1], 64)
			if err == nil {
				dfmt.Printf(" cb_v=%0.2f ", cb_v)
				ss_m, err := strconv.ParseFloat(cb.Substrings[midx], 64)
				if err == nil {
					dfmt.Printf(" ss_m=%0.2f", ss_m)
					switch scs[0] {
					case ">":
						dfmt.Printf(" op=%s", scs[0])
						return matched(ss_m > cb_v)
					case "<":
						dfmt.Printf(" op=%s", scs[0])
						return matched(ss_m < cb_v)
					case ">=":
						dfmt.Printf(" op=%s", scs[0])
						return matched(ss_m >= cb_v)
					case "<=":
						dfmt.Printf(" op=%s", scs[0])
						return matched(ss_m <= cb_v)
					}
				} else {
					dfmt.Printf(" reject(%s)", cb.Substrings[midx])
				}
			}
			dfmt.Printf("\n")
		}
	}
	return 0
}
