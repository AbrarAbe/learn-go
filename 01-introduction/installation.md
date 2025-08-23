# Installing Go

This guide covers the installation of Go on different operating systems.

## Prerequisites

Before installing Go, ensure you have an internet connection to download the necessary files.

## Installation Steps

### Linux

1.  **Download:** Obtain the latest Go binary distribution from the [official Go website](https://go.dev/dl/). Choose the archive appropriate for your system (e.g., `go1.22.1.linux-amd64.tar.gz`).
2.  **Extract:** Extract the archive to `/usr/local/go`. You'll need administrative privileges:
    ```bash
    sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
    ```
    *(Replace `go1.22.1.linux-amd64.tar.gz` with the actual downloaded filename.)*
3.  **Configure Environment Variables:** Edit your shell configuration file (e.g., `~/.bashrc`, `~/.zshrc`, `~/.profile`) and add the following lines:
    ```bash
    export GOROOT=/usr/local/go
    export PATH=$GOROOT/bin:$PATH
    ```
4.  **Apply Changes:** Source your configuration file:
    ```bash
    source ~/.bashrc # Or your respective shell configuration file
    ```

### macOS

1.  **Download:** Download the macOS installer package (`.pkg` file) from the [official Go website](https://go.dev/dl/).
2.  **Run Installer:** Execute the downloaded `.pkg` file and follow the on-screen instructions. The installer typically handles environment variable setup automatically.

### Windows

1.  **Download:** Download the Windows installer (`.msi` file) from the [official Go website](https://go.dev/dl/).
2.  **Run Installer:** Execute the installer and follow the prompts. Ensure you select the option to add Go to your system's PATH during installation. The installer usually configures the environment variables for you.

## Verification

After installation, you can verify it by opening your terminal or command prompt and running:

```bash
go version
```
This command should display the installed Go version.