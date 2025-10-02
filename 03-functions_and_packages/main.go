package main

import (
	"fmt"
	"learn-go/03-functions_and_packages/utils"
)

// simple function (no parameter, no return value)
func greet() {
	fmt.Println("Hello from a function!")
}

// function with paramater (no return value)
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// fuction with return value
func add(a int, b int) int {
	return a + b
}

// function with multiple return value
func calculate(a, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func main() {
	// calling the great function
	greet()
	sayHello("Abrar")
	result := add(5, 3) // calling the function and return value
	fmt.Println("Sum:", result)
	s, d := calculate(10, 4) // assigning multiple return value
	fmt.Println("Sum", s, "Difference:", d)

	// Using our custom package
	upperName := utils.ToUpper("abrar")
	fmt.Println("Uppercase name:", upperName)
}
