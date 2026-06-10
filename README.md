# QERP (千易) Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/QuoVadis86/go-qianyi-sdk.svg)](https://pkg.go.dev/github.com/QuoVadis86/go-qianyi-sdk)
[![Go Version](https://img.shields.io/github/go-mod/go-version/qy-info/go-qianyi-sdk)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/qy-info/go-qianyi-sdk)](LICENSE)

A Go SDK for the [QERP Open API](https://open.qianyierp.com/en-US/) (千易ERP开放平台).

## Features

- **MD5 Signing** — Automatic signature generation for every request
- **All Endpoints** — Shop, SKU, Order, Refund, Warehouse, Inventory, ASN, ODO, Adjust, Purchase, Logistics, Reports
- **Idiomatic Go** — Clean API with typed request/response structures
- **Environment Support** — Test, domestic production, overseas production

## Installation

```bash
go get github.com/QuoVadis86/go-qianyi-sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/QuoVadis86/go-qianyi-sdk"
)

func main() {
    sdk := qianyi.NewSDK("your-app-id", "your-app-secret")
    sdk.TestEnv() // Use test environment

    // List shops
    shops, total, err := sdk.Shop.QueryList(1, 10, "", "", "", "")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Shops: %d total\n", total)
    for _, shop := range shops {
        fmt.Printf("  %s (%s)\n", shop.Name, shop.Platform)
    }
}
```

## Environments

```go
sdk.TestEnv()                    // gerp-test1.800best.com
// default:                       www.qianyierp.com (domestic)
qianyi.WithBaseURL("https://asia.qianyierp.com") // overseas
```

## Services

| Service | File | Operations |
|---------|------|-----------|
| `sdk.Shop` | `shop.go` | Query shop list |
| `sdk.Sku` | `sku.go` | Query, create, update, enable/disable, query sys SKU |
| `sdk.Order` | `order.go` | Create, cancel, query list, query number list |
| `sdk.Refund` | `refund.go` | Create, cancel, query list |
| `sdk.Warehouse` | `warehouse.go` | Query warehouse list |
| `sdk.Inventory` | `inventory.go` | Query inventory V2, query inventory logs |
| `sdk.Asn` | `asn.go` | Create, query list, close inbound orders |
| `sdk.Odo` | `odo.go` | Query list, close outbound orders |
| `sdk.Adjust` | `adjust.go` | Query list, create adjustment orders |
| `sdk.Purchase` | `purchase.go` | Query list, create, update purchase orders |
| `sdk.Logistics` | `logistics.go` | Query channels, query tracking |
| `sdk.Report` | `report.go` | Query report list, query sales report |
| `sdk.CustomerField` | `customerfield.go` | Query custom fields |

## Authentication

All API calls require an `appId` and `appSecret` obtained from Qianyi ERP:
Settings → Three-party Application Integration → Get appId and appSecret

### Signature Algorithm

1. Sort parameters by name (ASCII ascending)
2. Concatenate as `key1=value1key2=value2...appSecret`
3. MD5 hash the result → 32-char lowercase hex string

## Error Handling

```go
result, err := sdk.Shop.QueryList(1, 10, "", "", "", "")
if err != nil {
    if apiErr, ok := err.(*qianyi.APIError); ok {
        fmt.Printf("API error [%s]: %s\n", apiErr.ErrorCode, apiErr.Message)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

## License

MIT
