# Contributing

Thank you for your interest in contributing to go-qianyi-sdk! We welcome contributions from the community.

## Getting Started

1. Fork the repository.
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/go-qianyi-sdk.git
   ```
3. Create a branch for your changes:
   ```bash
   git checkout -b feat/your-feature-name
   ```

## Development

### Prerequisites

- Go 1.22 or later
- golangci-lint (optional, for linting)

### Commands

```bash
make build      # Build the module
make test       # Run tests with race detection and coverage
make lint       # Run golangci-lint
make vet        # Run go vet
make fmt        # Format code
make ci         # Full CI pipeline (fmt → vet → lint → test → build)
```

### Code Style

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Run `make fmt` before committing.
- Ensure `make ci` passes locally before submitting a PR.

## Testing

- All new code must have corresponding tests.
- Use the mock HTTP client pattern in `client_test.go` for HTTP-level tests.
- Run `make test` and ensure no regressions.

## Pull Request Process

1. Update CHANGELOG.md with your changes under the `[Unreleased]` section.
2. Ensure all CI checks pass (lint, test on Go 1.22+).
3. Open a PR against the `main` branch.
4. Maintainers will review your PR within 3-5 business days.

## Commit Messages

Write [conventional commits](https://www.conventionalcommits.org/):

```
feat: add support for marketplace API
fix: handle null response in inventory query
docs: update README with new service endpoints
refactor: extract generic response helper
```

## Report Issues

Open a [GitHub Issue](https://github.com/QuoVadis86/go-qianyi-sdk/issues/new/choose) with:

- A clear description of the problem
- Steps to reproduce (if applicable)
- Go version and environment details

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
