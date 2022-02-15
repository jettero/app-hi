package main

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"

	c "github.com/jettero/app-hi/pkg/colors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func complainAboutPatterns(name string, thing interface{}) {
	os.Stderr.WriteString(fmt.Sprintf("WARNING: \"%s\" should be a list of mappings, given: %v\n", reflect.TypeOf(thing)))
	os.Exit(1)
}

func grabPatterns(group string, v interface{}) []string {
	var asList []string
	switch patterns := v.(type) {
	case nil:
		// this is fine
	case []interface{}:
		for _, v := range patterns {
			switch m := v.(type) {
			case map[interface{}]interface{}:
				for k, v := range m {
					asList = append(asList, k.(string))
					asList = append(asList, v.(string))
					if os.Getenv("DEBUG_HI_PATTERNS_FROM_CONFIG") == "1" {
						fmt.Printf("  found pattern \"%s\" => %s\n", k.(string), v.(string))
					}
				}
			default:
				complainAboutPatterns(group, v)
			}
		}
	default:
		complainAboutPatterns(group, v)
	}
	return asList
}

func ProcessConfigAndArgs() []string {
	var conf *string = pflag.StringP("config", "c", "", "read only this specified config file (\"none\" means \"load no config files\")")
	var group *string = pflag.StringP("group", "g", "patterns",
		"the name of the group of patterns to load from config")
	var list *bool = pflag.BoolP("list", "l", false, "show the list of patterns and exit")
	var halp *bool = pflag.BoolP("help", "h", false, "show the help screen text")

	pflag.Parse()
	Args := pflag.Args()

	if *halp {
		PrintHelp(0)
	}

	var locations []string
	if len(*conf) > 0 {
		fmt.Printf("... WTF \"%s\" config\n", *conf)
		if *conf == "none" {
			locations = []string{}
		} else {
			locations = []string{*conf}
		}
	} else {
		home := os.Getenv("HOME")
		if len(home) > 0 {
			locations = append(locations, []string{
				fmt.Sprintf("%s/.app-hi", home),
				fmt.Sprintf("%s/.app-hi.yaml", home),
				fmt.Sprintf("%s/.app-hi.json", home),
				fmt.Sprintf("%s/.config/app-hi/config.yaml", home),
				fmt.Sprintf("%s/.config/app-hi/config.json", home),
			}...)
		}
		locations = append(locations, "/etc/app-hi/config.yaml")
		locations = append(locations, "/etc/app-hi/config.json")
	}

	viper.SetConfigType("yaml")

	for _, location := range locations {
		if os.Getenv("DEBUG_HI_PATTERNS_FROM_CONFIG") == "1" {
			fmt.Printf("... \"%s\" config\n", location)
		}
		v := viper.New()
		v.SetConfigFile(location)
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err == nil {
			if os.Getenv("DEBUG_HI_PATTERNS_FROM_CONFIG") == "1" {
				fmt.Printf("read \"%s\" config\n", location)
			}
			Args = append(grabPatterns(*group, v.Get(*group)), Args...)
		} else if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			os.Stderr.WriteString(fmt.Sprintf("WARNING: couldn't read config file: %v\n", err))
		}
	}

	if *list {
		fmt.Printf("known patterns:\n")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		for i := 0; i+1 < len(Args); i += 2 {
			fmt.Fprintf(w, "\t%s\t%s\n", Args[i], c.Color(Args[i+1], Args[i+1])+c.Color("reset", ""))
		}
		w.Flush()
		os.Exit(0)
	}

	return Args
}
