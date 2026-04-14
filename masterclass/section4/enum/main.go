package main

import "fmt"

// Define multiple constants together
const (
	HOST = "localhost"
	PORT = 8080
)
const MAX, MIN = 100, 0

// 0 based indexing by default
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 1 based indexing like this
const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// Edit the 0 based indexing to custom incrementing numbers
type ByteSize float64

// The 1 << n is a bit shift — it means
// "shift the number 1 left by n bits",
// which equals 2ⁿ. So:
//
//	1 << 10 = 2¹⁰ = 1024
//	1 << 20 = 2²⁰ = 1,048,576
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Having LogLevel type defined in this way
// helps you compare types easily
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// UseLogLevel demonstrates using the LogLevel type
func UseLogLevel(level LogLevel) {
	// Do something with level
}

const (
	stuff, dblStuff = iota, iota * 2
	random, dblRandom
	hello, dblHello
)

func main() {
	// No enunm keyword
	// But provides mechanism to create enumerated constants
	// Using the iota identifier
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(January, February, March, April, May, June, July, August, September, October, November, December)

	// See ByteSize constants
	for _, size := range []ByteSize{KB, MB, GB, TB, PB, EB, ZB, YB} {
		fmt.Printf("%e\n", size)
	}

	// level variable requires LogLevel type
	var level LogLevel = Debug

	// String() method provides human readable output
	fmt.Println(level)

	// Compare level variable type easily
	fmt.Println(level == Debug)

	// Practice function call with LogLevel type
	UseLogLevel(level)

	// Define multiple constants with iota
	fmt.Println(stuff, dblStuff)
	fmt.Println(random, dblRandom)
	fmt.Println(hello, dblHello)
}
