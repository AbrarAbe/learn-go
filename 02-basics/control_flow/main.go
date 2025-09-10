package main

import "fmt"

func main() {
	// conditional statement (if, if else, else)
	score := 75

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
	// if wirh short statement
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	// fmt.Println(num) // This would cause an error: num is not defined here

	// Looping (for)
	// C style
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}
	// while style
	counter := 0
	for counter < 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// infinite loop with break
	num := 1
	for {
		fmt.Println("Number:", num)
		if num == 3 {
			break // Exit the loop when num is 3
		}
		num++
	}

	// for range loop
	colors := []string{"Red", "Green", "Blue"}
	for i, color := range colors {
		fmt.Printf("Index of %d: %s\n", i, color)
	}
	// ignoring index
	for _, color := range colors {
		fmt.Println("color", color)
	}

	// modifying slice elements using for range
	thing1 := [5]float64{1, 2, 3, 4, 5}
	for i := range thing1 {
		thing1[i] = thing1[i] * thing1[1]
	}
	fmt.Printf("The result is: %v\n", thing1)
}
