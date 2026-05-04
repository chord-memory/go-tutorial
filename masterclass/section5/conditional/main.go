package main

import (
	"fmt"
	"time"
)

func main() {
	tmp := 25
	if tmp > 30 {
		fmt.Println("tmp is greater than 30")
	} else {
		fmt.Println("tmp is less than 30")
	}

	// map example
	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}
	// ok = does the key exist in the map
	// hasAccess = the value of the key
	if hasAccess, ok := userAccess["jane"]; ok && hasAccess {
		fmt.Println("Jane has access:", ok, hasAccess)
	}
	// Can store variables but unnecessary if only used in conditional
	hasAccess, ok := userAccess["john"]
	if ok && hasAccess {
		fmt.Println("John has access:", ok, hasAccess)
	} else {
		fmt.Println("John does not have access:", ok, hasAccess)
	}

	// switch statements

	score := 85

	// switch checking expression
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	// switch checking equality
	day := "Sunday"
	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// store variable in switch context
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// type switch
	checkType := func(i interface{}) { // i can hold any type (interface{})
		switch v := i.(type) { // type switch: extracts the concrete type of i (only valid in switch statement)
		case int: // if i holds an int...
			fmt.Printf("int: %d\n", v)
		case string: // if i holds a string...
			fmt.Printf("string: %s\n", v)
		case bool: // if i holds a bool...
			fmt.Printf("bool: %t\n", v)
		default: // if i holds any other type...
			fmt.Printf("unknown: %T\n", v)
		}
	}

	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)

	// type assertion (can only be done on interface{} types)
	var i interface{} = 42 // i holds 42, but Go only knows it as interface{}
	// You can't do math on i directly — Go won't let you
	// i + 1  // compile error: invalid operation
	v, ok := i.(int) // "I believe this is an int, give it to me as one"
	// now v is 42 as a real int, ok is true

	// Safe form (two return values) (see above)
	if ok {
		fmt.Println(v + 1) // safe to use v as int
	}

	// Unsafe form (one return value) — panics if wrong
	v_unsafe := i.(int) // crashes at runtime if i isn't actually an int
	fmt.Println(v_unsafe + 1)

	var i_hello interface{} = "hello"
	v_int, ok_int := i_hello.(int)    // ok = false, v = 0 (zero value for int)
	v_str, ok_str := i_hello.(string) // ok = true,  v = "hello"
	fmt.Println(v_int, ok_int)
	fmt.Println(v_str, ok_str)
}
