# Control Flow in Go

Control flow statements dictate the order in which code is executed, allowing programs to make decisions and repeat actions. Go provides `if`/`else` statements for conditional execution and a versatile `for` loop for iteration.

## Conditional Statements (`if`, `else if`, `else`)

`if` statements execute a block of code if a specified condition is true. `else if` allows checking multiple conditions sequentially, and `else` provides a default path.

*   Conditions must evaluate to a boolean (`true` or `false`).
*   Parentheses `()` around conditions are not required, but curly braces `{}` are mandatory for the code blocks.

**Syntax:**

```go
if condition1 {
    // Code block if condition1 is true
} else if condition2 {
    // Code block if condition1 is false and condition2 is true
} else {
    // Code block if all preceding conditions are false
}
```

**Example:**

```go
package main

import "fmt"

func main() {
	temperature := 25

	if temperature > 30 {
		fmt.Println("It's hot outside!")
	} else if temperature > 20 {
		fmt.Println("It's warm outside.")
	} else {
		fmt.Println("It's cool outside.")
	}
}
```

**`if` with a Short Statement:**
A short statement can be included before the condition. Variables declared in this statement are scoped to the `if` block.

```go
if value := 10 * 2; value > 15 {
    fmt.Println("Value is greater than 15:", value) // Value is greater than 15: 20
}
```

## Loops (`for`)

Go has a single looping construct: the `for` loop, which can be used in various forms.

### 1. C-style `for` loop

This is a traditional loop with initialization, condition, and post-statement.

**Syntax:**

```go
for initialization; condition; post {
    // Code to execute
}
```

**Example:**

```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Iteration:", i)
	}
}
```

### 2. `while`-style `for` loop

By omitting initialization, condition, and post-statements, the `for` loop can act like a `while` loop when a condition is provided.

**Syntax:**

```go
for condition {
    // Code to execute as long as condition is true
}
```

**Example:**

```go
package main

import "fmt"

func main() {
	count := 0
	for count < 3 {
		fmt.Println("Count:", count)
		count++
	}
}
```

### 3. Infinite `for` loop

An infinite loop can be created by omitting all statements. Use `break` to exit.

**Syntax:**

```go
for {
    // Code to execute indefinitely
    if conditionToBreak {
        break // Exit the loop
    }
}
```

### 4. `for...range` loop

Iterates over elements in collections like slices, arrays, strings, maps, and channels. It provides the index and value (or key and value).

**Syntax (for slices/arrays):**

```go
for index, value := range collection {
    // Use index and value
}
```

**Example:**

```go
package main

import "fmt"

func main() {
	colors := []string{"Red", "Green", "Blue"}
	for i, color := range colors {
		// %d is a format specifier for integers, printing the index 'i'
		// %s is a format specifier for strings, printing the color 'color'
		fmt.Printf("Index %d: %s\n", i, color)
	}

    // Ignoring index
    for _, color := range colors {
        fmt.Println("Color:", color)
    }
}