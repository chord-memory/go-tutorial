package main

import "fmt"

func main() {

	// Only one type of loop: for

	// C-style
	// prints 0 thru 9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while-type
	// prints 3 thru 1
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	// infinite loop
	// prints 0 thru 4
	counter := 0
	for {
		fmt.Println("counter:", counter)
		counter++
		if counter == 5 {
			break
		}
	}

	// skipping
	// prints odd numbers 1 thru 9
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// array
	items := [3]string{"Go", "Python", "JavaScript"}
	for i := 0; i < len(items); i++ {
		fmt.Println(i, items[i])
	}
	for index, value := range items {
		fmt.Println(index, value)
	}
	for _, value := range items {
		fmt.Println(value)
	}
	for index, _ := range items {
		fmt.Println(items[index])
	}

}
