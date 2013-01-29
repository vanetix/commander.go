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
	args := os.Args

	for i := range args {
		fmt.Println(args[i])
	}
}

/**
 * `Add` `option` to the commander instance
 */

func (commander *Commander) Add(name string, tiny byte, verbose string, 
								description string, required bool, callback func()) {
	option := new(Option)

	option.Name = name
	option.Tiny = tiny
	option.Verbose = verbose
	option.Description = description
	option.Required = required
	option.Callback = callback

	append(*commander).Options, option)
}

/**
 * Display the usage of `commander`
 */

func (commander *Commander) Usage() {}

/**
 * Return the total number of arguments
 */

func (commander *Commander) Len() int {
	return len(os.Args)
}
