# Changelog

## [1.0.0] - 2026-06-10

### Added
- Initial stable release of go-qianyi-sdk.
- Full API coverage: Shop, SKU, Order, Refund, Warehouse, Inventory, ASN, ODO, Adjust, Purchase, Logistics, Report, CustomerField (66+ methods).
- `context.Context` support across all API methods for cancellation, deadlines, and tracing.
- Generic response helpers (`doList`, `doSingle`, `doAction`) reducing boilerplate by ~60%.
- CI/CD pipeline: GitHub Actions with multi-version Go testing, golangci-lint, and coverage reporting.
- Comprehensive test suite: client tests, response parsing, error handling, and mock HTTP tests.
- Makefile with `test`, `lint`, `build`, `coverage`, `ci` targets.
- `.golangci.yml` with 40+ linters configured.
- User-Agent header (`go-qianyi-sdk/1.0`) on all requests.
- CHANGELOG, CONTRIBUTING, and issue/PR templates.

### Changed
- **Breaking**: All service methods now accept `context.Context` as the first parameter.
- **Breaking**: `Client.Do` now takes `context.Context` as the first parameter.
- `strings.Builder` replaced with `bytes.Buffer` for multipart body construction.
- JSON marshal errors and multipart write errors are now properly handled (previously silent).
- `QueryShippingInfo`, `QueryPickupStatus`, `QueryOrderDocument` now return `(json.RawMessage, error)` instead of `error` (previously silent data loss).

### Fixed
- Unrealistic `go 1.26.1` downgraded to `go 1.22.0` (stable, well-tested baseline).
- `go.sum` now committed for reproducible builds.

### Removed
- All `_ = json.Marshal(...)` and `_ = w.WriteField(...)` silent error suppression patterns.
