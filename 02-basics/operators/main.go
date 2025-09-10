package main

import "fmt"

func main() {
	a := 10
	b := 5

	// Arithmetic operators
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	print("--------------------------\n")

	// Comparison operators
	fmt.Println("a > b =", a > b)
	fmt.Println("a < b =", a < b)
	fmt.Println("a == b =", a == b)
	fmt.Println("a != b =", a != b)
	fmt.Println("a >= b =", a >= b)
	fmt.Println("a <= b =", a <= b)

	print("--------------------------\n")

	// Logical operators
	isSunny := true
	isWarm := false
	// Using &&
	if isSunny && isWarm {
		fmt.Println("It's a sunny and warm day.")
	} else {
		fmt.Println("It's not both sunny and warm.")
	}
	// Using ||
	if isSunny || isWarm {
		fmt.Println("It's sunny or warm (or both).")
	}
	// Using !
	isNotSunny := !isSunny // isNotSunny becomes false
	fmt.Println("Is it not sunny?", isNotSunny)

	print("--------------------------\n")

	// Assignment operators
	score := 100
	score -= 10                        // score is now 90
	score /= 2                         // score is now 45
	fmt.Println("Final score:", score) // Output: Final score: 45

	print("--------------------------\n")

	// Bitwise operators
	x := 5
	y := 3
	fmt.Println("x & y =", x&y)   // Bitwise AND
	fmt.Println("x | y =", x|y)   // Bitwise OR
	fmt.Println("x ^ y =", x^y)   // Bitwise XOR
	fmt.Println("x << 1 =", x<<1) // Left shift
	fmt.Println("x >> 1 =", x>>1) // Right shift

	print("--------------------------\n")

	// Type conversion
	var num int = 42
	var float float64 = float64(num)
	fmt.Println("Converted float:", float)

	print("--------------------------\n")
}
