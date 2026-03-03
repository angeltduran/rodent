# rodent

simple E2E testing library. The purpose of it is not to download and install any external libraries/frameworks as it should run as is, specially useful for containers / pods to not overload them with dependencies / libraries, configure, setup and supply chain attacks. You could use puppeteer or playwright for your tests, but I believe E2E should be either done through command line flags or bash scripts with a simple binary. I Have left an example page with some basic common use cases where you could expand your E2E tests

## Run binary
## Requirements
- A Chromium-based browser available to launch/download
- download binary from /bin folder and run
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
