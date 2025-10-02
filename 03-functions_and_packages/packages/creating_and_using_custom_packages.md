# Creating and Using Custom Packages

While Go's standard library is extensive, the true power of packages comes from organizing your own code into logical, reusable modules. This guide walks you through creating a custom package and using it in your main application.

## Core Concepts

*   **Package as a Directory:** In Go, a package is typically represented by a directory containing one or more `.go` files. All files within that directory must declare the same package name.
*   **Exported vs. Unexported:** Go uses a simple rule to control visibility:
    *   If an identifier (like a function, type, or variable) starts with a **capital letter**, it is **exported** (public) and can be accessed by other packages.
    *   If it starts with a **lowercase letter**, it is **unexported** (private) and can only be accessed from within its own package.

## Step-by-Step Example

Let's illustrate this with the `utils` package created in the `learn-go/03-functions_and_packages` directory.

### 1. Creating the Package

First, we created a directory to house our package:
`learn-go/03-functions_and_packages/utils/`

Inside this directory, we created a file named `helpers.go`:
```learn-go/03-functions_and_packages/utils/helpers.go#L1-16
package utils

import "strings"

// ToUpper converts a string to uppercase.
// This function is exported because its name starts with a capital letter.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// toLower converts a string to lowercase.
// This function is unexported because its name starts with a lowercase letter.
// It can only be used by other functions within the 'utils' package.
func toLower(s string) string {
	return strings.ToLower(s)
}
```
*   `package utils` at the top declares all code in this file belongs to the `utils` package.
*   `ToUpper` is exported and can be used by `main.go`.
*   `toLower` is unexported and is only for internal use within the `utils` package.

### 2. Using the Custom Package

Next, we modified `main.go` to import and use our new package.

The `import` path is derived from the module name defined in `go.mod` (`learn-go`) plus the path to the package directory.

```learn-go/03-functions_and_packages/main.go#L1-39
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
```
*   We added `import "learn-go/03-functions_and_packages/utils"` to make our package's functionality available.
*   We then called the exported function using `utils.ToUpper("abrar")`. The syntax is `packageName.FunctionName()`.

This process of creating and importing packages is fundamental to building modular and maintainable applications in Go.