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
	Tiny byte
	Verbose string
	Description string
	Required bool
	Callback func()
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
 * `Parse` arguments
 */

func (commander *Commander) Parse() {
	args := commander.explode(os.Args[1:])
	l := len(args)

	for i := range commander.Options {
		option := commander.Options[i]

		for j := range args {
			arg := args[j]

			if utf8.RuneCountInString(arg) < 2 {
				prog.Usage()
			}

			if option.Tiny == arg || option.Verbose == arg {
				if l - 1 > j && !args[j + 1][0] == '-' {
					option.Callback(args[j + 1])
				} else {
					option.Callback()
				}
			}
		}
	}

}

func (commander *Commander) explode(args []string) []string {
	newargs := make([]string, len(args))

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

func (commander *Commander) Add(option *Option) {
	commander.Options = append(commander.Options, *option)
}

/**
 * Display the usage of `commander`
 */

func (commander *Commander) Usage() {
	fmt.Fprintf(os.Stderr, "\n  Usage: %s [options]\n\n", commander.Name)
	fmt.Fprintf(os.Stderr, "  Options:\n");

	options := commander.Options
	for i := range options {
		fmt.Fprintf(os.Stderr, "    -%c, --%s %s",
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
