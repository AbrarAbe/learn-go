package main

import "fmt"

func main() {
	var age int = 30
	var name string = "Abrar"
	var isUnemployed = "true"
	var phi float64 = 3.141592653589793 // get the exact number back - 3.141592653589793
	// var phi float32 = 3.141592653589793 // do not get the exact number back - 3.1415927

	fmt.Println("your age is:", age)
	fmt.Println("your name is:", name)
	fmt.Println("still unemployed:", isUnemployed)
	fmt.Println("phi:", phi)

	count := 10      // Go infers the type as int
	message := "Hi"  // Go infers the type as string
	isValid := false // Go infers the type as bool

	fmt.Println("Count:", count)
	fmt.Println("Message:", message)
	fmt.Println("Is Valid:", isValid)

	// Declare multiple variables in a single line
	var x, y, z int = 1, 2, 3

	// Declare multiple variables with inferred types
	a, b, c := 4, 5, 6

	fmt.Println("x:", x, "y:", y, "z:", z)
	fmt.Println("a:", a, "b:", b, "c:", c)

	// Declare multiple variables with inferred types and initial values
	d, e, f := 7, 8, 9

	fmt.Println("d:", d, "e:", e, "f:", f)
	fmt.Printf("d: %d e: %d f: %d\n", d, e, f) // Printf allows for formatted output with type specifiers. %d is used for integers, and \n adds a newline character at the end.
}
