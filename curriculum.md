# Go Programming Tutorial

This tutorial will guide you through the Go programming language. We'll cover the fundamentals and gradually increase the complexity.

## Curriculum:

1.  **Introduction to Go & Setting Up Your Environment**
    *   [x] What is Go? Why use it?
    *   [x] Installing Go (and verifying the installation).
    *   [x] Your first "Hello, World!" program.
    *   [x] Understanding Go's project structure.

2.  **Basic Syntax & Data Types**
    *   [x] Variables, Constants, and Data Types (int, float, string, bool).
    *   [x] Operators (arithmetic, comparison, logical).
    *   [x] Control Flow (if/else, for loops).

3.  **Functions & Packages**
    *   [ ] Defining and calling functions.
    *   [ ] Function parameters and return values.
    *   [ ] Introduction to packages and importing.
    *   [ ] Creating and using custom packages.

4.  **Data Structures: Arrays, Slices, and Maps**
    *   [ ] Arrays and Slices: differences and use cases.
    *   [ ] Maps (dictionaries): creating, accessing, and modifying.

5.  **Pointers**
    *   [ ] Understanding pointers and their use.
    *   [ ] Working with pointers and dereferencing.

6.  **Structs & Methods**
    *   [ ] Defining structs and accessing fields.
    *   [ ] Methods associated with structs.

7.  **Interfaces**
    *   [ ] Understanding interfaces and their role in Go.
    *   [ ] Implementing interfaces.

8.  **Concurrency: Goroutines and Channels**
    *   [ ] Introduction to concurrency in Go.
    *   [ ] Goroutines and how to launch them.
    *   [ ] Channels for communication between goroutines.

9.  **Error Handling**
    *   [ ] Understanding errors in Go.
    *   [ ] Handling errors using `if err != nil`.
    *   [ ] Custom error types.

10. **Testing**
    *   [ ] Writing unit tests in Go.
    *   [ ] Using the `testing` package.

11. **Working with Files**
    *   [ ] Reading and writing files.
    *   [ ] File operations using the `os` and `io/ioutil` packages.

12. **JSON Handling**
    *   [ ] Encoding and decoding JSON.
    *   [ ] Working with the `encoding/json` package.

13. **Modules & Dependency Management**
    *   [ ] Introduction to Go modules.
    *   [ ] Managing dependencies.

14. **Project: Project Based Learning**
    *   [ ] We will build a project based on your interests.

## Lesson 1: Introduction to Go & Setting Up Your Environment

### 1. What is Go? Why Use It?

Go, also known as Golang, is an open-source programming language developed by Google. It's designed to be:

*   **Efficient:** It compiles to machine code, making it fast.
*   **Concurrent:** Built-in support for concurrency makes it easy to write programs that can do multiple things at once.
*   **Simple:** Go has a clean and straightforward syntax, making it easy to learn and use.
*   **Scalable:** Suitable for both small and large projects.
*   **Cross-Platform:**  You can compile Go code for various operating systems (Windows, macOS, Linux, etc.).

Why use Go?

*   **Backend Development:** Web servers, APIs, microservices.
*   **Cloud Computing:** Kubernetes, Docker, and other cloud infrastructure tools are written in Go.
*   **System Programming:** Network tools, command-line utilities.
*   **DevOps:** Automation, infrastructure management.

### 2. Installing Go

Follow these steps based on your operating system:

*   **Linux:**
    1.  Download the Go binary from the official Go website: [https://go.dev/dl/](https://go.dev/dl/) (Choose the appropriate version for your system)
    2.  Extract the archive to `/usr/local/go`: `sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz` (Replace `go1.22.1.linux-amd64.tar.gz` with the actual file name)
    3.  Set the `GOROOT` and `PATH` environment variables (edit your `.bashrc`, `.zshrc`, or similar):
        ```bash
        export GOROOT=/usr/local/go
        export PATH=$GOROOT/bin:$PATH
        ```
    4.  Source your shell configuration: `source ~/.bashrc` or `source ~/.zshrc`.

*   **macOS:**
    1.  Download the Go installer from the official Go website: [https://go.dev/dl/](https://go.dev/dl/)
    2.  Run the installer.
    3.  The installer usually sets up the environment variables automatically.  You can confirm by checking your shell configuration file (e.g., `.zshrc` or `.bash_profile`) for the `GOROOT` and `PATH` settings.

*   **Windows:**
    1.  Download the Go installer from the official Go website: [https://go.dev/dl/](https://go.dev/dl/)
    2.  Run the installer and follow the prompts. Make sure you check the box to add Go to your PATH.
    3.  The installer should set up the environment variables automatically.

### 3. Verifying the Installation

Open your terminal or command prompt and type:

```bash
go version