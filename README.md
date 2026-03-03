# rodent

simple E2E testing library

## Run binary
## Requirements
- A Chromium-based browser available to launch/download
download binary from /bin folder and run
```bash
./rodent -url "https://example.com" -out "example_dot_com.pdf" -timeout 45s
```
## Run from source
## Requirements

- Go 1.22+
- A Chromium-based browser available for Rod to launch/download

## Setup

```bash
go mod tidy
```

## Usage 

```bash
go run . -url "https://example.com" -out "example.pdf"
```

Optional flags:

- `-timeout` (default: `30s`)

Example:

```bash
go run . -url "https://golang.org" -out "golang.pdf" -timeout 45s
```
