package main

import "fmt"

func main() {
	// Variable declaration
	// but not initialized
	// i.e. memory hasn't been allocated for it yet
	// but it depends, compiler can choose to allocate
	// memory and give it a default value
	// in Go, "zero value" is assigned to variables
	// that have not beeen initialized
	var greeting string // zero-value is an empty string ""

	// This is the initialization, now we allocated memory
	// variable is pointing to this location of data in memory
	greeting = "Hello friend"
	fmt.Println(greeting)

	// Can declare & init for ints
	var count int
	count = 10
	fmt.Println(count)

	// Can declare and init for bools
	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	// Can declare multiple variables same line
	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Println(firstName, lastName)

	// Short variable declaration & initialization same time (preferred)
	// Go compiler infers the type from the assigned value
	email := "test@test.com"
	fmt.Println(email)

	age := 24
	fmt.Println(age)

	// Redundant but acceptable syntax it seems
	var year int = 2025
	fmt.Println(year)

	// Acceptable and not redundant but I suppose less clean than :=
	var month = "May"
	fmt.Println(month)
}
