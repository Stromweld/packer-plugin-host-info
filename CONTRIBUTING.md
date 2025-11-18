# Contributing to Packer Plugin Host Info

Thank you for your interest in contributing to the Packer Host Info plugin!

## Development Environment

### Prerequisites

- [Go](https://golang.org/doc/install) >= 1.20
- [Packer](https://www.packer.io/downloads) >= 1.10.2
- [packer-plugin-sdk](https://github.com/hashicorp/packer-plugin-sdk) >= v0.5.2

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Stromweld/packer-plugin-host-info
   cd packer-plugin-host-info
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Install packer-sdc tool:
   ```bash
   go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@latest
   ```

## Building

### Quick Build

```bash
make build
```

Or manually:

```bash
go build -o packer-plugin-host-info
```

### Development Build

Build and install locally for testing:

```bash
make dev
```

This will:
1. Build the plugin with development version
2. Install it to your local Packer plugins directory

## Testing

### Unit Tests

Run all unit tests:

```bash
make test
```

Or manually:

```bash
go test ./... -v
```

### Acceptance Tests

Run acceptance tests (requires Packer to be installed):

```bash
make testacc
```

Or manually:

```bash
PACKER_ACC=1 go test -count 1 -v ./datasource/hostinfo -timeout=120m
```

### Test Coverage

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Code Generation

The plugin uses code generation for HCL2 specifications. After modifying struct tags or adding new configuration fields:

```bash
make generate
```

Or manually:

```bash
cd datasource/hostinfo
go generate
```

## Making Changes

### Adding New Fields to Output

1. Update the `DatasourceOutput` struct in `datasource/hostinfo/data.go`
2. Update the `Execute()` method to populate the new field
3. Run `make generate` to update generated files
4. Update tests in `datasource/hostinfo/data_test.go`
5. Update documentation in `.web-docs/`

Example:

```go
type DatasourceOutput struct {
    OS           string `mapstructure:"os"`
    Version      string `mapstructure:"version"`
    Architecture string `mapstructure:"architecture"`
    Platform     string `mapstructure:"platform"`
    Family       string `mapstructure:"family"`
    NewField     string `mapstructure:"new_field"` // Add new field
}
```

### Code Style

- Follow standard Go conventions and idioms
- Run `go fmt` before committing
- Run `go vet` to catch common mistakes
- Add comments for exported functions and types

### Commit Messages

Use clear, descriptive commit messages:

```
Add new field for kernel version detection

- Added kernel_version to DatasourceOutput
- Updated documentation
- Added tests for new field
```

## Documentation

### Update Documentation

All user-facing changes should include documentation updates:

1. **README.md** - High-level overview and quick start
2. **docs/README.md** - Plugin documentation index
3. **docs/datasources/info.mdx** - Data source reference
4. **USAGE.md** - Detailed usage examples
5. **example/** - Working examples

### Generate Web Documentation

```bash
make generate
```

This generates documentation for the Packer website.

## Pull Requests

### Before Submitting

- [ ] Tests pass (`make test`)
- [ ] Code is formatted (`go fmt ./...`)
- [ ] Generated code is up to date (`make generate`)
- [ ] Documentation is updated
- [ ] Examples work with the changes

### PR Guidelines

1. Create a descriptive PR title
2. Reference any related issues
3. Describe the changes and motivation
4. Include any breaking changes in the description
5. Add/update tests for new functionality

## Release Process

Releases are automated via GitHub Actions when a new tag is pushed:

1. Update version in `version/version.go`
2. Update CHANGELOG.md
3. Create and push a tag:
   ```bash
   git tag v0.1.0
   git push origin v0.1.0
   ```

## Getting Help

- Open an issue for bug reports or feature requests
- Check existing issues before creating a new one
- Provide as much context as possible

## License

By contributing, you agree that your contributions will be licensed under the project's license (MPL-2.0).

