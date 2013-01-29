package commander

import (
	"testing"
)

func TestParse(t *testing.T) {
	commander := Commander{
		Name: "commander",
		Version: "1.0.0",
	}

	commander.Add(Option{
		Name: "Help",
		Tiny: 'h',
		Verbose: "help",
		Description: "Display usage",
		Callback: commander.Usage,
		Required: true,
	})

	commander.Parse()
}
