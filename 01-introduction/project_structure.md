# Go Project Structure

Go doesn't enforce a strict project directory layout, but common conventions exist to help organize Go projects, especially with the advent of Go Modules.

## Go Modules

Go Modules are the standard for dependency management and project structure in modern Go development. They allow projects to be self-contained, independent of the `GOPATH` environment variable.

## Common Directory Layout

A typical Go project might include the following structure:

```
myproject/
├── go.mod        # Module definition and dependencies
├── go.sum        # Checksums for dependencies
├── main.go       # Main application entry point (for simple projects)
├── cmd/          # Contains subdirectories for multiple main applications
│   ├── app1/
│   │   └── main.go
│   └── app2/
│       └── main.go
├── internal/     # Private application and library code, not importable by external projects
│   └── ...
├── pkg/          # Public library code intended for external use
│   └── ...
├── api/          # API definitions (e.g., OpenAPI specifications)
│   └── ...
├── web/          # Web assets, templates, static files
│   └── ...
└── Makefile      # Optional: for build automation tasks
```

## Key Files and Directories:

*   **`go.mod`**: Essential for Go Modules. It defines the module's path and lists its direct dependencies.
*   **`go.sum`**: Contains the cryptographic checksums of dependencies, ensuring their integrity.
*   **`cmd/`**: Used when a project has multiple distinct executable applications. Each subdirectory here corresponds to a separate executable.
*   **`internal/`**: Code within this directory is considered private to the module. It cannot be imported by other Go projects outside of this module. This is useful for encapsulating implementation details.
*   **`pkg/`**: Conventionally used for library code that is intended to be shared and imported by other projects. However, many developers now prefer to place reusable packages directly under the module root or in directories that mirror their import paths (e.g., `myproject/mypackage/`).

Understanding and adhering to these conventions can significantly improve the maintainability and clarity of your Go projects.