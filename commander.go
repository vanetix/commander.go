package commander

import (
	"os"
	"fmt"
	"unicode/utf8"
)

/**
 * `Option` type, contains data for a specific option
 */

type Option struct {
	Name string
	Tiny string
	Verbose string
	Description string
	Required bool
	Callback func(...string)
}

/**
 * Commander type, contains all program data
 */

type Commander struct {
	Name string
	Version string
	Options []Option
}

/**
 * Initialize a new commander with `args`
 */

func Init(name string, version string) *Commander {
	p := &Commander{
		Name: name,
		Version: version,
	}

	p.Add(&Option{
		Name: "help",
		Tiny: "-h",
		Verbose: "--help",
		Description: "display usage",
		Required: false,
		Callback: func(args ...string) {
			p.Usage()
		},
	})

	p.Add(&Option{
		Name: "version",
		Tiny: "-V",
		Verbose: "--version",
		Description: "display version",
		Required: false,
		Callback: func(args ...string) {
			fmt.Fprintf(os.Stdout, "  Version: %s\n", p.Version)
		},
	})

	return p
}

/**
 * `Parse` arguments
 */

func (commander *Commander) Parse() {
	args := commander.explode(os.Args[1:])

	for i, l := 0, len(commander.Options); i < l; i++ {
		option := commander.Options[i]

		found := false
		for j, l := 0, len(args); j < l; j++ {
			arg := args[j]

			if option.Tiny == arg || option.Verbose == arg {
				if j != l  - 1 && args[j + 1][0] != '-' {
					option.Callback(args[j + 1])
					j++
				} else {
					option.Callback()
				}
				found = true
			}
		}
		if option.Required && !found {
			// Option is required and not found
			fmt.Fprintf(os.Stderr, "%s, %s is required.", option.Tiny, option.Description)
		}
	}
}

func (commander *Commander) explode(args []string) []string {
	newargs := make([]string, 0, len(args))

	for i := range args {
		arg := args[i]
		l := utf8.RuneCountInString(arg)

		if l > 2 && arg[0] == '-' && arg[1] != '-' {
			for i := 1; i < l; i++ {
				newargs = append(newargs, "-" + string(arg[i]))
			}
		} else {
			newargs = append(newargs, arg)
		}
	}

	return newargs
}

/**
 * `Add` `option` to the commander instance
 */

func (commander *Commander) Add(options ...*Option) {
	for i := range options {
		commander.Options = append(commander.Options, *options[i])
	}
}

/**
 * Display the usage of `commander`
 */

func (commander *Commander) Usage() {
	fmt.Fprintf(os.Stderr, "\n  Usage: %s [options]\n\n", commander.Name)
	fmt.Fprintf(os.Stderr, "  Options:\n");

	options := commander.Options
	for i := range options {
		fmt.Fprintf(os.Stderr, "    %s, %s %s",
			options[i].Tiny, options[i].Verbose, options[i].Description)
	}
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

/**
 * Return the total number of arguments
 */

func (commander *Commander) Len() int {
	return len(os.Args)
}
