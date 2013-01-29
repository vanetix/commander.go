package commander

import (
	"os"
	"fmt"
)

/**
 * `Option` type, contains data for a specific option
 */

type Option struct {
	Required bool
	Name string
	Tiny byte
	Verbose string
	Description string
	Callback *func
}

/**
 * Commander type, contains all program data
 */

type Commander struct {
	Name string
	Version string
	Options []Option
	Args []string
}

/**
 * `Parse` arguments
 */

func (commander *Commander) Parse() {}

/**
 * `Add` `option` to the commander instance
 */

func (commander *Commander) Add(option *Option) {}

/**
 * Display the usage of `commander`
 */

func (commander *Commander) Usage() {}

/**
 * Return the total number of arguments
 */

func (commander *Commander) NArgs() int {
	return len((*commander).args)
}
