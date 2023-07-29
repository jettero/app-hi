// vi:ft=gomod:
// vim was trying to load this file as a prolog module or something
// ?- cool(syntax).
// no.

module github.com/jettero/app-hi

go 1.20

// Elara's not happy with the way I changed this to get it working. We're
// having a dispute about the way the math should work out on the ovector
// pointers. Until that's resolved, I just published it the way I like it.
//
// I consider this temporary though.
replace go.elara.ws/pcre => github.com/jettero/golang-pcre2 v0.0.0-20230729214243-e12307ce97ca

require go.elara.ws/pcre v0.0.0-20230717141135-d1b9df80a165

require (
	github.com/hashicorp/golang-lru v0.5.4
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	modernc.org/libc v1.24.1 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.6.0 // indirect
)
