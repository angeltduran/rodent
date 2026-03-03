# Website to PDF (Go + Rod)

Simple Go CLI to convert a website URL into a PDF file using [`go-rod`](https://github.com/go-rod/rod).

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
