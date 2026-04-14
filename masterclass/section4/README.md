# Section 4 Summary

## Print

`Println` adds newline character at end by default and spaces between args
```
fmt.Println(true, false)
```
`Printf` offers string formatting using format specifiers to substitute variables into a string
e.g. %d for integers, %s for strings, %f for float, %t for bool, %v for default format, %T for type, %p for pointer
```
fmt.Printf("Hello, %s version %.2f\n", name, version)
```
*Primary formatters*:
| Specifier | Description |
|-----------|-------------|
| `%+v`     | Like `%v`; for structs, includes field names |
| `%#v`     | Go-syntax style (types/literals where it applies) |
| `%8v`     | Width allocated |
| `%-v`     | Left-align, makes sense with width allocated |
| `%0d`     | Zero-pad for numbers |

Many more types exist for integers, floats, strings / byte slices, and pointers.

Note there are 2 sizes of floats: `float32` and `float64`

## Variables

Variable declaration (without initialization)
"zero value" is assigned to uninitialized variabled
```
var greeting string  // zero value is empty string ""
```
Variable initialization, now we assign a value
```
greeting = "Hello friend"
```
Can declare multiple variables on the same line
```
var firstName, lastName string
firstName = "John"
lastName = "Doe"
```
Short variable delcaration and initialization at the same time
```
email := "test@test.com"
```
Redundant but acceptable syntax
```
var year int = 2025
```
Acceptable and not redundant but less clean than `:=`
```
var month = "May"
```
Note: can declare/initialize variable outside of function

## Constants

Constants are immutable and cannot be reassigned
```
const HOST string = "localhost"
```
The variable type may be inferenced
```
const PORT = 8080
```
Can define multiple constants together
```
const (
	HOST = "localhost"
	PORT = 8080
)
const MAX, MIN = 100, 0
```
Note: can declare/initialize const outside of function

## Enum

Create an enum by: (1) defining a custom type, (2) declaring typed constants with `iota` as the incrementor, and (3) adding a `String()` method
```
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
```
NOTE: `String()` is like `__str__` in Python

Change the iota values to custom incrementing numbers
```
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
```
To better understand bit shifting in general, most relevant for bitmask flags:
```
1 << 0:  0000 0001 = 1   (2⁰)
1 << 1:  0000 0010 = 2   (2¹)
1 << 2:  0000 0100 = 4   (2²)
1 << 3:  0000 1000 = 8   (2³)
1 << 10: 0100 0000 0000 = 1024 (2¹⁰)
```
`1 << n` means "take the number 1 and shift its bits left by n positions."　Each left shift is equivalent to multiplying by 2.　So 1 << 10 = 2¹⁰ = 1024.

## Logger

Another way to implement String method via map:
* We are attaching a `String` method to `LogLevel` type
* This is a method receiver, so called via the `LogLevel` variable like:
* `level.String()` or `fmt.Printf("%s", level)` (fmt auo-detects this method)
* Akin to attaching methods to classes in other laguages, you attach methods to types
* Satisfies fmt.Stringer interface, which fmt uses for %s and %v formatting
```
var levelNames = map[LogLevel]string{
	Debug:   "DEBUG",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR",
}

func (l LogLevel) String() string {
	if name, ok := levelNames[l]; ok {
		return name
	}
	return "UNKNOWN"
}
```

NOTE: Go has no classes. Instead:
* Structs (`type Person struct { Name string }`) — collections of fields, like classes but without inheritance
* Any type can have methods — you attached `String()` to `LogLevel`, which is just an `int` underneath
* Composition over inheritance — embed structs within structs instead of extending classes
* Interfaces — define behavior, satisfied implicitly (no `implements` keyword)

Therefore, attachment of `String()` to `LogLevel` achieves OOP-like encapsulation without classes.