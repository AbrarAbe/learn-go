package main

import "fmt"

func main() {
	// Arithmetic operators
	a := 10
	b := 5

	// Arithmatic oprators
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	/// Comparison operators
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// Logical operators
	x := true
	y := false
	fmt.Println("x && y:", x && y) // AND
	fmt.Println("x || y:", x || y) // OR
	fmt.Println("!x:", !x)         // NOT
}