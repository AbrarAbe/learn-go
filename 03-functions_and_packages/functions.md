# Defining and Calling Functions

Functions are fundamental building blocks in Go, enabling you to modularize your code, enhance reusability, and improve readability. They encapsulate a specific task, making complex programs more manageable.

## Function Definitions

A Go function is defined using the `func` keyword, followed by the function name, its parameters, and its return type(s).

**Syntax:**

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // Function body: code to be executed when the function is called
    // ...
    return value // Required if a returnType is specified
}
```

*   **`func`**: Keyword to declare a function.
*   **`functionName`**: A unique identifier for the function.
*   **`parameters`**: A comma-separated list of inputs the function accepts, each with its type. If consecutive parameters share the same type, you can omit the type for all but the last parameter (e.g., `(a, b int)` instead of `(a int, b int)`).
*   **`returnType`**: The type of the value(s) the function will send back. If a function doesn't return any value, the `returnType` is omitted. Go functions can return multiple values.
*   **`{}`**: Curly braces enclose the function body, containing the statements to be executed.

## Calling Functions

To execute a function, you simply call it by its name, followed by parentheses containing any required arguments.

**Example 1: Function with No Parameters and No Return Value**

```go
package main

import "fmt"

// greet prints a simple greeting message.
func greet() {
    fmt.Println("Hello from a function!")
}

func main() {
    greet() // Call the greet function
}
```
**Output:**
```
Hello from a function!
```

**Example 2: Function with Parameters**

```go
package main

import "fmt"

// sayHello takes a name (string) and prints a personalized greeting.
func sayHello(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    sayHello("Alice") // Pass "Alice" as an argument
    sayHello("Bob")   // Pass "Bob" as an argument
}
```
**Output:**
```
Hello, Alice
Hello, Bob
```

**Example 3: Function with a Single Return Value**

```go
package main

import "fmt"

// add takes two integers and returns their sum.
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(10, 5) // Call add and assign its return value to 'result'
    fmt.Println("Sum:", result) // Output: Sum: 15
}
```

**Example 4: Function with Multiple Return Values**

Go's ability to return multiple values is a powerful feature, often used for returning a result alongside an error.

```go
package main

import "fmt"

// divide performs integer division and returns the quotient and a boolean indicating success.
func divide(numerator, denominator int) (int, bool) {
    if denominator == 0 {
        return 0, false // Cannot divide by zero
    }
    return numerator / denominator, true
}

func main() {
    quotient, ok := divide(10, 2)
    if ok {
        fmt.Println("Quotient:", quotient) // Output: Quotient: 5
    } else {
        fmt.Println("Cannot divide by zero.")
    }

    q, success := divide(7, 0)
    if success {
        fmt.Println("Quotient:", q)
    } else {
        fmt.Println("Cannot divide by zero.") // Output: Cannot divide by zero.
    }
}
```

By mastering function definition and calling, you gain the ability to structure your Go programs logically and efficiently.