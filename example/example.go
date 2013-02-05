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

	help := &commander.Option{
		Name: "help",
		Tiny: "-h",
		Verbose: "--help",
		Description: "display usage",
		Required: false,
		Callback: func(...string) {
			prog.Usage()
		},
	}

	version := &commander.Option{
		Name: "version",
		Tiny: "-v",
		Verbose: "--version",
		Description: "display version",
		Required: false,
		Callback: func(...string) {
			fmt.Printf("  Version: %s\n", prog.Version)
		},
	}

	prog.Add(help)
	prog.Add(version)
	prog.Parse()
}
