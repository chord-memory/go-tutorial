package main

import "fmt"

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

/*
var levelNames = []string{"DEBUG", "INFO", "WARNING", "ERROR"}

func (l LogLevel) String() string {
	// if l < Debug || l > Error {
	// 	return "UNKNOWN"
	// }
	// Convert LogLevel type to int for comparison
	// Prefer this so bounds do not need to be changed
	// if LogLevels are added to const
	if l < 0 || int(l) >= len(levelNames) {
		return "UNKNOWN"
	}
	return levelNames[int(l)]
}
*/

// Personally I prefer map approach
// over the slice/array approach
// See switch approach in enum file

var levelNames = map[LogLevel]string{
	Debug:   "DEBUG",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR",
}

// We are attaching a String method to LogLevel type
// This is a method receiver, so called via the LogLevel variable like:
// `level.String()` or `fmt.Printf("%s", level)` (fmt auto-detects this method)
// Akin to attaching methods to classes in other languages, you attach methods to types
// Satisfies fmt.Stringer interface, which fmt uses for %s and %v formatting
func (l LogLevel) String() string {
	if name, ok := levelNames[l]; ok {
		return name
	}
	return "UNKNOWN"
}

func printLogLevel(level LogLevel) {
	// The String() method is called automatically when using %s
	// fmt.Printf("Log level: %d %s\n", level, level.String())
	fmt.Printf("Log level: %d %s\n", level, level)
}

// Run `go generate` to generate stringer methods
// Must install from Go tools repo
// `go install golang.org/x/tools/cmd/stringer@latest`

// TODO: Run `go generate` to generate stringer methods
// TODO: Learn more generators https://chatgpt.com/s/t_69deac2574bc81919a879e06ff791b6c

//go:generate stringer -type=Color

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {

	var Unknown LogLevel = 99

	printLogLevel(Debug)
	printLogLevel(Info)
	printLogLevel(Warning)
	printLogLevel(Error)
	printLogLevel(Unknown)

	// `type LogLevel int`
	// LogLevel is an alias for int, so we can
	// pass a int to printLogLevel without err
	printLogLevel(10)
}
