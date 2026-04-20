# Go SDK

Official Go client for CustomKeys secrets management.

[![Go Reference](https://pkg.go.dev/badge/github.com/cksxe/sdk-go.svg)](https://pkg.go.dev/github.com/cksxe/sdk-go)

## Install

```bash
go get github.com/cksxe/sdk-go
```

## Quick Start

```go
package main

import (
    "log"
    "os"

    customkeys "github.com/cksxe/sdk-go"
)

func main() {
    client, err := customkeys.New(
        customkeys.WithToken(os.Getenv("CUSTOMKEYS_TOKEN")),
        customkeys.WithEnv(os.Getenv("CUSTOMKEYS_ENV_ID")),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    dbURL, ok := client.Get("DATABASE_URL")
    if ok {
        log.Println("Connected:", dbURL)
    }
}
```

## Documentation

Full docs at [customkeys.superxepic.dev/docs/sdk-go](https://customkeys.superxepic.dev/docs/sdk-go)

## Security

- TLS 1.2+ enforced
- Token sanitization (rejects control chars)
- Response body size limits (10 MB)
- Memory zeroing on Close()
- No secrets written to disk

## License

MIT
