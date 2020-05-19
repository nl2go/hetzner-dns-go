# hetzner-dns-go: A Go library for the Hetzner DNS API

![Build](https://gitlab.com/newsletter2go/hetzner-dns-go/badges/master/pipeline.svg) ![Coverage](https://gitlab.com/newsletter2go/hetzner-dns-go/badges/master/coverage.svg)

Package hetzner-dns-go is a library for the Hetzner DNS API.

The library’s documentation is available at [GoDoc](https://godoc.org/github.com/nl2go/hetzner-dns-go),
the public API documentation is available at [dns.hetzner.com](https://dns.hetzner.com/api-docs/).

## Example

```go
package main

import (
    "fmt"
    "log"

    client "github.com/nl2go/hetzner-dns-go"
)

func main() {
    hDNSClient := client.NewAuthApiTokenClient("yourAuthAPIToken")

    zones, err := hDNSClient.ZonesGet()
    if err != nil {
        log.Fatalf("error while retrieving zones list: %s\n", err)
    }

    fmt.Println(servzonesers)
}
```
