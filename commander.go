package commander

import (
	"os"
	"fmt"
)

/**
 * `Option` type, contains data for a specific option
 */

type Option struct {
	Name string
	Tiny rune
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
	args := os.Args

	for i := range args {
		fmt.Println(args[i])
	}
}

/**
 * `Add` `option` to the commander instance
 */

func (commander *Commander) Add(option *Option) {
	append(commander.Options, option)
}

/**
 * Display the usage of `commander`
 */

func (commander *Commander) Usage() {
	fmt.Fprintf(os.Stderr, "\n  Usage: %s [options]\n\n", commander.Name)
	fmt.Fprintf(os.Stderr, "  Options:\n");

	options := &commander.Options
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
