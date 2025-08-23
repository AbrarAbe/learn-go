# Verifying Go Installation

This guide explains how to verify that Go has been successfully installed on your system.

## Verification Steps

1.  **Open Terminal/Command Prompt:** Launch your system's command-line interface (Terminal on macOS/Linux, Command Prompt or PowerShell on Windows).

2.  **Run the `go version` Command:** Type the following command and press Enter:
    ```bash
    go version
    ```

3.  **Check the Output:**
    *   **Success:** If Go is installed correctly, you will see output similar to this (the version number and system details may vary):
        ```
        go version go1.22.1 linux/amd64
        ```
    *   **Failure:** If you receive an error message such as "command not found," it indicates that the Go executable is not recognized by your system. This often means the `PATH` environment variable was not set up correctly during installation, or the installation itself was incomplete.

## Troubleshooting Common Issues

If the verification command fails, please revisit the installation instructions for your operating system and ensure that the Go binary directory (`$GOROOT/bin`) has been added to your system's `PATH` environment variable. You might need to restart your terminal or even your computer for the changes to take effect.