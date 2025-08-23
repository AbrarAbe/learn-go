# Your First "Hello, World!" Program

This section guides you through creating and running your very first Go program.

## Core Concepts

*   **`package main`**: Declares the package as `main`, indicating it's an executable program. Every executable Go program must have a `main` package.
*   **`import "fmt"`**: Imports the `fmt` package, which provides functions for formatted input and output operations, such as printing to the console.
*   **`func main()`**: This is the entry point of the program. When you run an executable Go program, the `main` function in the `main` package is executed first.
*   **`fmt.Println(...)`**: A function from the `fmt` package that prints its arguments to standard output, followed by a newline character.

## Steps to Create and Run

1.  **Create a Directory**: Make a new directory for your project, for example, `hello-world`.
    ```bash
    mkdir hello-world
    cd hello-world
    ```
2.  **Create the Go File**: Inside the `hello-world` directory, create a file named `main.go`.
3.  **Add Code**: Paste the following code into `main.go`:
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
    ```
4.  **Run the Program**: Open your terminal or command prompt in the `hello-world` directory and run the program using the `go run` command:
    ```bash
    go run main.go
    ```
    You should see the output:
    ```
    Hello, World!
    ```

This simple program demonstrates the basic structure of a Go executable file and how to print output to the console.