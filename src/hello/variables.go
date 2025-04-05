package main

import "fmt"

func main() {
	// Declaring variables
	var age int = 30
	var name string = "Alice"
	var isStudent bool = true
	const pi float64 = 3.14
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is student:", isStudent)
	fmt.Println("Pi:", pi)

	// Shorthand declaration (only inside functions)
	city := "New York"
	fmt.Println("City:", city)

	//Mutliple declarations
	var x, y int = 10,20
	fmt.Println("x:", x, "y:", y)
}