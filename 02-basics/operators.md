# Operators in Go

Operators are symbols used to perform operations on values (operands). Go supports various types of operators for arithmetic, comparison, logical operations, and assignments.

## Arithmetic Operators

Used for mathematical calculations.

*   `+` : Addition
*   `-` : Subtraction
*   `*` : Multiplication
*   `/` : Division (integer division for integers, floating-point division for floats)
*   `%` : Modulo (remainder of integer division)

**Example:**
```go
package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("a + b =", a + b)     // Output: a + b = 13
	fmt.Println("a - b =", a - b)     // Output: a - b = 7
	fmt.Println("a * b =", a * b)     // Output: a * b = 30
	fmt.Println("a / b =", a / b)     // Output: a / b = 3 (integer division)
	fmt.Println("a %% b =", a % b)    // Output: a % b = 1
}
```

## Comparison (Relational) Operators

Used to compare two values. They return a boolean (`true` or `false`).

*   `==` : Equal to
*   `!=` : Not equal to
*   `>`  : Greater than
*   `<`  : Less than
*   `>=` : Greater than or equal to
*   `<=` : Less than or equal to

**Example:**
```go
package main

import "fmt"

func main() {
	x := 15
	y := 20

	fmt.Println("x == y:", x == y) // Output: x == y: false
	fmt.Println("x != y:", x != y) // Output: x != y: true
	fmt.Println("x < y:", x < y)   // Output: x < y: true
}
```

## Logical Operators

Used to combine or modify boolean conditions.

*   `&&` : Logical AND (true if both operands are true)
*   `||` : Logical OR (true if at least one operand is true)
*   `!`  : Logical NOT (reverses the boolean value)

**Example:**
```go
package main

import "fmt"

func main() {
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
}
```

## Assignment Operators

Used to assign values to variables.

*   `=` : Simple assignment
*   `+=` : Add and assign (e.g., `a += b` is `a = a + b`)
*   `-=` : Subtract and assign
*   `*=` : Multiply and assign
*   `/=` : Divide and assign
*   `%=` : Modulo and assign

**Example:**
```go
package main

import "fmt"

func main() {
	score := 100
	score -= 10 // score is now 90
	score /= 2  // score is now 45
	fmt.Println("Final score:", score) // Output: Final score: 45
}
```

## Bitwise Operators (Brief Mention)

Go also supports bitwise operators (`&`, `|`, `^`, `&^`, `<<`, `>>`) which operate on the binary representations of numbers. These are typically used for low-level operations and data manipulation.

Understanding these operators is fundamental for controlling program flow and performing calculations in Go.