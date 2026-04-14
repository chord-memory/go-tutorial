package main

import "fmt"

// Can declare const/var outside of func
// but cannot do anything else. See err below
// when I had `TEST = TEST + "test"`
/*
jordan@Jordans-MBP go-tutorial % go run ./masterclass/section4/value
# github.com/chord-memory/go-tutorial/masterclass/section4/value
masterclass/section4/value/main.go:10:1: syntax error: non-declaration statement outside function body
masterclass/section4/value/main.go:13:1: syntax error: non-declaration statement outside function body
jordan@Jordans-MBP go-tutorial %
*/
const HOST string = "localhost"

func main() {

	// Host is a const so cannot change it. Will see the below error if changed
	// HOST = HOST + ":8000"
	/*
		jordan@Jordans-MBP go-tutorial % go run ./masterclass/section4/value
		# github.com/chord-memory/go-tutorial/masterclass/section4/value
		masterclass/section4/value/main.go:13:2: cannot assign to HOST (neither addressable nor a map index expression)
		jordan@Jordans-MBP go-tutorial %
	*/

	// TEST is a var not const so it can be edited
	var TEST string = "test"
	TEST = TEST + "test"
	fmt.Println(TEST)

	// I suppose if var is omitted tho then it is var not const
	stuff := "random"
	fmt.Println(stuff)
	// Ok only omit var if doing : but can omit type whenever value is included
	// I guess this is still short hand (declare and init same line)
	// but not as clean as :=
	var hello = "world"
	fmt.Println(hello)
	// And I suppose you can omit type for const since redundant
	const thing = "whatever"
	fmt.Println(thing)

	// Two different sizes of floats

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 5.2
	fmt.Println(rate)

	// Type can be inferred
	const PORT = 8080
	fmt.Println(PORT)
}
