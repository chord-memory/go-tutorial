package main

import "fmt"

type Point struct{ X, Y int }

// Functions must be declared at package level, not inside main
func (p Point) String() string {
	return fmt.Sprintf("Point[X=%d, Y=%d]", p.X, p.Y)
}

func main() {
	var t any

	// Println adds newline character at end
	// by default and spaces between args

	// Printf offers string formatting
	// using format specifiers
	// e.g. %d for integers, %s for strings, %v for default format, %T for type
	// to substitute variables into a string

	name := "Go"
	version := 1.22
	fmt.Printf("Hello, %s version %.2f\n", name, version)
	// Output: Hello, Go version 1.22

	fmt.Printf("Hello World\n")
	fmt.Printf("Hello " + "World\n") // string concatenation
	fmt.Println(1 + 1)
	fmt.Println(3.14)
	fmt.Println(true, false)

	// Nil formatters:
	fmt.Printf("%v\n", t)  // print the value human readable: <nil>
	fmt.Printf("%#v\n", t) // same output as %v when value is nil

	// Slice formatters:
	// A slice is like an extendable array
	fmt.Printf("%v\n", []int{1, 2, 3})  // print the slice human readable: [1 2 3]
	fmt.Printf("%#v\n", []int{1, 2, 3}) // go syntax representation of the slice: []int{1, 2, 3}
	fmt.Printf("%T\n", []int{1, 2, 3})  // type of the slice: []int

	// Struct formatters:
	// A struct is a collection of key & value pairs
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "John", Age: 30}
	fmt.Printf("%v\n", user)  // print the struct human readable: {John 30}
	fmt.Printf("%#v\n", user) // go syntax representation of the struct: main.User{Name:"John", Age:30}
	fmt.Printf("%+v\n", user) // print the struct human readable with field names: {Name:John Age:30}
	fmt.Printf("%T\n", user)  // type of the struct: main.User

	// Array & Map formatters mirror Slice & Struct formatters

	// Integer formatters:
	number := 100
	fmt.Printf("%d\n", number)
	fmt.Printf("%8d\n", number)  // width allocated to 8 digits
	fmt.Printf("%08d\n", number) // zero-pad the number to 8 digits
	fmt.Printf("%-8d\n", number) // left-align the number to 8 digits
	fmt.Printf("%+8d\n", number) // include sign
	fmt.Printf("%T\n", number)   // type of the number: int
	fmt.Printf("%v\n", number)   // default format for the number: 100
	fmt.Printf("%#v\n", number)  // go syntax representation of the number: 100

	// Float formatters:
	pi := 3.1415926
	fmt.Printf("%f\n", pi)   // fixed to 6 decimal places: 3.141593
	fmt.Printf("%.2f\n", pi) // fixed to specified decimal places: 3.14
	fmt.Printf("%e\n", pi)   // scientific notation: 3.141593e+00
	fmt.Printf("%g\n", pi)   // whichever is shorter, %f or %e: 3.1415926
	fmt.Printf("%T\n", pi)   // type of the float: float64
	fmt.Printf("%#v\n", pi)  // go syntax representation of the float: 3.1415926

	// String formatters:
	hello := "Hello"
	fmt.Printf("%s\n", hello)   // string: Hello
	fmt.Printf("%q\n", hello)   // quoted string: "Hello"
	fmt.Printf("%x\n", hello)   // hexadecimal string: 48656c6c6f
	fmt.Printf("%X\n", hello)   // hexadecimal string with uppercase letters: 48656c6c6f
	fmt.Printf("%T\n", hello)   // type of the string: string
	fmt.Printf("%#v\n", hello)  // go syntax representation of the string: "Hello"
	fmt.Printf("%8v\n", hello)  // width allocated to 8 characters: "     Hello"
	fmt.Printf("%-8v\n", hello) // width allocated to 8 characters, left-align: "Hello     "

	number = 123456789
	fmt.Printf("%e\n", float64(number)) // scientific notation: 1.234568e+08

	// TODO: understand this

	// Errors: Most errors implement fmt.Stringer (String() string), so you can print with %v or %s;
	// use %q to quote their output if needed.
	err := fmt.Errorf("this is an error")
	fmt.Printf("%v\n", err) // this is an error
	fmt.Printf("%s\n", err) // this is an error
	fmt.Printf("%q\n", err) // "this is an error"

	// Custom formatting: types can implement the fmt.Stringer interface for %v/%s,
	// or fmt.Formatter for more control (and GoStringer for %#v).
	pt := Point{X: 1, Y: 2}
	fmt.Printf("%v\n", pt) // Point[X=1, Y=2]

}
