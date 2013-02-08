# commander.go [![Build Status](https://travis-ci.org/vanetix/commander.go.png?branch=master)](https://travis-ci.org/vanetix/commander.go)
*More unix like argument parsing for go.* Heavily inspired by: [commander.c](https://github.com/visionmedia/commander.c)

## Installing
`go get github.com/vanetix/commander.go`

## Usage
```go
func doWork(args ...string) {
	// Do some work
}

func main() {
	program := commander.Init("Program", "0.1.3")

	commander.Add(&commander.Option{
	    Required: false,
	    Name: "Work",
	    Tiny: "-w",
	    Verbose: "--work",
	    Description: "do some work",
	    Callback: doWork,
	})

	commander.Parse()
}
```

## Release history
- 0.0.1 - Initial stable version

## License (MIT)
Copyright (c) 2013 Matt McFarland

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
