package commander

import (
	"os"
	"testing"
)

func TestArgData(t *testing.T) {
	exec := false

        doWork := func(s ...string) {
		if len(s) < 1 || len(s) > 1 {
			t.Error("Got an invalid number of arguments");
		} else {
			exec = true
		}
	}

	prog := Init("test", "0.0.1")
	prog.Add(&Option{
		Name: "work",
		Tiny: "-w",
		Verbose: "--work",
		Description: "do work",
		Required: true,
		Callback: doWork,
	})

	// Mock arguments to parse
	os.Args = []string{"bogus", "-w", "work"}
	prog.Parse()

	if !exec {
		t.Error("Did not execute `doWork()`")
	}
}

func TestArgs(t *testing.T) {
	exec := 0

	doWork := func(s ...string) {
		exec++
	}

	prog := Init("test", "0.0.1")
	prog.Add(&Option{
		Name: "work",
		Tiny: "-w",
		Verbose: "--work",
		Description: "do work",
		Required: true,
		Callback: doWork,
	})

	prog.Add(&Option{
		Name: "more work",
		Tiny: "-m",
		Verbose: "--more",
		Description: "do work",
		Required: true,
		Callback: doWork,
	})

	// Mock arguments to parse
	os.Args = []string{"bogus", "-w", "-m"}
	prog.Parse()

	if exec != 2 {
		t.Error("Did not run all callbacks")
	}
}
