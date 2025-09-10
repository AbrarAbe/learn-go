# Variables, Constants, and Data Types

Understanding variables, constants, and data types is fundamental to programming in Go. This section covers their declaration, usage, and Go's built-in data types.

## Variables

Variables are named storage locations that hold values. In Go, variables must be declared with a specific type, and their values can be changed.

### Declaration

You can declare variables using the `var` keyword:

```go
var variableName type
var age int
var name string = "Bob" // Declaration with initialization
```

Go also supports **short variable declaration** using the `:=` operator, which infers the type:

```go
count := 10       // count is an int
message := "Hi"   // message is a string
isValid := false  // isValid is a bool
```
The `:=` operator can only be used inside functions.

### Zero Values

If a variable is declared but not initialized, it will be assigned its **zero value**:

*   `0` for numeric types (int, float, etc.)
*   `false` for boolean types
*   `""` (empty string) for string types
*   `nil` for pointers, slices, maps, channels, interfaces, and functions

## Constants

Constants are identifiers that represent fixed values. Unlike variables, their values cannot be changed after they are defined. They must be initialized at compile time.

### Declaration

Constants are declared using the `const` keyword:

```go
const constantName type = value
const Pi float64 = 3.14159
const Greeting string = "Hello"
```

You can also declare constants without explicitly specifying the type, and Go will infer it:

```go
const Version = "1.0" // Version is inferred as string
const Count = 100     // Count is inferred as int
```

## Basic Data Types

Go has several built-in data types:

### Numeric Types

*   **Integers:**
    *   `int`: The default integer type. Its size (32 or 64 bits) depends on the system architecture.
    *   `int8`, `int16`, `int32`, `int64`: Signed integers of specific sizes.
    *   `uint`, `uint8`, `uint16`, `uint32`, `uint64`: Unsigned integers of specific sizes. `uint` is the unsigned counterpart to `int`.
    *   `byte`: Alias for `uint8`. Used for byte values.
    *   `rune`: Alias for `int32`. Used for Unicode code points.
*   **Floating-Point Numbers:**
    *   `float32`: Single-precision floating-point number.
    *   `float64`: Double-precision floating-point number (the default and generally preferred for precision).
*   **Complex Numbers:**
    *   `complex64`, `complex128`: For complex numbers.

### Boolean Type

*   **`bool`**: Represents `true` or `false`.

### String Type

*   **`string`**: Represents a sequence of characters. Strings in Go are immutable.

### Composite Types (Brief Mention)

Go also supports composite types like Arrays, Slices, Structs, Maps, Pointers, Functions, Channels, and Interfaces, which will be covered in later lessons.

By mastering these fundamental data types and the concepts of variables and constants, you build a strong base for writing Go programs.