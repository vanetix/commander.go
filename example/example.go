package main

import (
	"fmt"
	"github.com/vanetix/commander.go"
)

func main() {
	prog := &commander.Commander{
		Name: "example",
		Version: "0.0.1",
	}

	option := &commander.Option{
		Name: "help",
		Tiny: "-h",
		Verbose: "--help",
		Description: "display usage",
		Required: true,
		Callback: func(...string) {
			prog.Usage()
		},
	}

	prog.Add(option)

	fmt.Println(prog)
	prog.Parse()
}
