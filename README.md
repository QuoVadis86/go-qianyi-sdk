# QERP (千易) Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/QuoVadis86/go-qianyi-sdk.svg)](https://pkg.go.dev/github.com/QuoVadis86/go-qianyi-sdk)
[![Go Version](https://img.shields.io/github/go-mod/go-version/QuoVadis86/go-qianyi-sdk)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/QuoVadis86/go-qianyi-sdk)](LICENSE)

---

> **中文** | [English](#qerp-qianyi-go-sdk)

千易ERP开放平台 Go SDK，覆盖店铺、商品、订单、库存、仓库、采购、物流等全模块 API。

---

## 特性

- **MD5 签名** — 自动为每个请求生成签名
- **全接口覆盖** — 店铺、商品、订单、库存、仓库、采购、物流、报表等
- **环境切换** — 测试环境、国内生产环境、海外生产环境一键切换
- **零依赖** — 纯标准库实现

---

## 安装

```bash
go get github.com/QuoVadis86/go-qianyi-sdk
```

---

## 快速开始

```go
package main

import (
    "fmt"
    "log"
    "github.com/QuoVadis86/go-qianyi-sdk"
)

func main() {
    sdk := qianyi.NewSDK("your-app-id", "your-app-secret")
    sdk.TestEnv() // 切换到测试环境

    shops, total, err := sdk.Shop.QueryList(1, 10, "", "", "", "")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("店铺总数: %d\n", total)
    for _, shop := range shops {
        fmt.Printf("  店铺: %s (平台: %s, 状态: %s)\n", shop.Name, shop.Platform, shop.Status)
    }
}
```

---

## 环境

| 环境 | 地址 | 用途 |
|------|------|------|
| 测试 | `gerp-test1.800best.com` | 开发调试 |
| 国内生产 | `www.qianyierp.com` | 中国区服务 |
| 海外生产 | `asia.qianyierp.com` | 全球服务 |

```go
sdk.TestEnv()                                    // 测试环境
qianyi.WithBaseURL("https://asia.qianyierp.com") // 海外生产
// 默认: www.qianyierp.com (国内生产)
```

---

## 服务列表

| 服务 | 文件 | 操作 |
|------|------|------|
| `sdk.Shop` 店铺 | `shop.go` | 查询店铺列表 |
| `sdk.Sku` 商品 | `sku.go` | 查询、创建、更新、启禁用、系统SKU查询 |
| `sdk.Order` 订单 | `order.go` | 创建、取消、查询列表、查询单号列表 |
| `sdk.Refund` 退款/退货 | `refund.go` | 创建、取消、查询列表 |
| `sdk.Warehouse` 仓库 | `warehouse.go` | 查询仓库列表 |
| `sdk.Inventory` 库存 | `inventory.go` | 查询库存(V2)、查询库存日志 |
| `sdk.Asn` 入库单 | `asn.go` | 创建、查询列表、关闭 |
| `sdk.Odo` 出库单 | `odo.go` | 查询列表、关闭 |
| `sdk.Adjust` 调整单 | `adjust.go` | 查询列表、创建 |
| `sdk.Purchase` 采购 | `purchase.go` | 查询列表、创建、更新 |
| `sdk.Logistics` 物流 | `logistics.go` | 查询渠道、查询轨迹 |
| `sdk.Report` 报表 | `report.go` | 查询报表列表、查询销售报表 |
| `sdk.CustomerField` 自定义字段 | `customerfield.go` | 查询自定义字段 |

---

## 认证

在千易ERP后台获取凭证：**设置 → 第三方应用对接**，获取 `appId` 和 `appSecret`。

### 签名算法

1. 按参数名 ASCII 升序排列
2. 拼接为 `key1=value1key2=value2...{appSecret}`
3. MD5 加密，得到 32 位小写 hex 签名

### 请求格式

- **方法**: POST
- **Content-Type**: multipart/form-data
- **公共参数**: appId, serviceType, bizParam (JSON), timestamp (毫秒), sign

### 响应格式

```json
{
    "state": "success",
    "errorCode": "",
    "errorMsg": "",
    "bizContent": "{...业务数据 JSON...}",
    "requestId": "uuid"
}
```

---

## 错误处理

```go
result, err := sdk.Shop.QueryList(1, 10, "", "", "", "")
if err != nil {
    if apiErr, ok := err.(*qianyi.APIError); ok {
        fmt.Printf("API 错误 [%s]: %s\n", apiErr.ErrorCode, apiErr.Message)
    } else {
        fmt.Printf("错误: %v\n", err)
    }
}
```

---

## English

# QERP (Qianyi) Go SDK

A Go SDK for the [QERP Open API](https://open.qianyierp.com/en-US/) — Shop, SKU, Order, Inventory, Warehouse, Purchase, Logistics and more.

### Features

- **MD5 Signing** — Automatic signature generation for every request
- **Full API Coverage** — Shop, SKU, Order, Refund, Warehouse, Inventory, ASN, ODO, Adjust, Purchase, Logistics, Report, CustomerField
- **Environment Support** — Test, domestic production, overseas production
- **Zero Dependencies** — Pure standard library implementation

### Quick Start

```go
sdk := qianyi.NewSDK("your-app-id", "your-app-secret")
sdk.TestEnv() // Switch to test environment

shops, total, _ := sdk.Shop.QueryList(1, 10, "", "", "", "")
fmt.Printf("Total shops: %d\n", total)
```

### Authentication

Get credentials from Qianyi ERP: **Settings → Three-party Application Integration**.

**Signature algorithm:**
1. Sort parameters by name (ASCII ascending)
2. Concatenate as `key1=value1key2=value2...{appSecret}`
3. MD5 hash → 32-char lowercase hex string

### Services

| Service | Endpoint | Operations |
|---------|----------|------------|
| Shop | `/api/v1/shop` | Query list |
| SKU | `/api/v1/sku` | Query, create, update, enable/disable |
| Order | `/api/v1/salesOrder` | Create, cancel, query list |
| Refund | `/api/v1/returnOrder` | Create, cancel, query list |
| Warehouse | `/api/v1/warehouse` | Query list |
| Inventory | `/api/v1/inventory` | Query V2, query logs |
| ASN | `/api/v1/asn` | Create, query, close |
| ODO | `/api/v1/odo` | Query list, close |
| Adjust | `/api/v1/adjust` | Query list, create |
| Purchase | `/api/v1/purchase` | Query list, create, update |
| Logistics | `/api/v1/logistics` | Query channels, tracking |
| Report | `/api/v1/report` | Query list, sales report |
| CustomerField | `/api/v1/customerField` | Query custom fields |

---

## License

MIT
