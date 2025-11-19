# GitHub Copilot Instructions for packer-plugin-host-info

## Project Overview

This is a **Packer plugin** that provides a data source for detecting the host operating system, version, and CPU architecture where Packer is running. It's written in **Go** and follows the HashiCorp Packer Plugin SDK conventions.

**Project Details:**
- **Name:** packer-plugin-host-info
- **Type:** Packer Data Source Plugin
- **Language:** Go 1.25.3
- **SDK:** hashicorp/packer-plugin-sdk v0.6.4
- **License:** Apache-2.0
- **Author:** Corey Hemminger
- **Repository:** github.com/Stromweld/packer-plugin-host-info

## Architecture

### Project Structure
```
.
├── main.go                      # Plugin entry point
├── datasource/
│   └── host-info/
│       ├── data.go             # Main data source implementation
│       ├── data.hcl2spec.go    # Generated HCL2 spec
│       ├── data_test.go        # Unit tests
│       └── data_acc_test.go    # Acceptance tests
├── version/
│   └── version.go              # Version information
├── example/                     # Example Packer templates
├── docs/                        # Plugin documentation
├── .web-docs/                   # Web documentation assets
└── GNUmakefile                 # Build automation
```

### Key Components

1. **Main Entry Point** (`main.go`):
   - Registers the plugin with Packer's plugin system
   - Uses `plugin.NewSet()` to create a plugin set
   - Registers the datasource as `plugin.DEFAULT_NAME`
   - Sets plugin version from `version.PluginVersion`

2. **Datasource Implementation** (`datasource/host-info/data.go`):
   - **Config**: Empty struct (no configuration required)
   - **Datasource**: Main struct containing config
   - **DatasourceOutput**: Contains detected OS information fields
   - **Methods**:
     - `ConfigSpec()`: Returns HCL2 spec for configuration
     - `Configure()`: Configures the datasource
     - `OutputSpec()`: Returns HCL2 spec for output
     - `Execute()`: Core logic that detects host information

3. **Version Management** (`version/version.go`):
   - Version: "0.0.2"
   - VersionPrerelease: "dev" (overridden during builds)
   - Uses HashiCorp's plugin versioning system

## Development Guidelines

### Code Style and Conventions

1. **Go Best Practices**:
   - Follow standard Go formatting (use `gofmt` or `goimports`)
   - Use meaningful variable and function names
   - Add comments for exported types and functions
   - Keep functions small and focused

2. **Package Naming**:
   - Main datasource package is `hostinfo` (no hyphen)
   - Import alias for datasource: `hostInfoData`
   - Import alias for version: `hostInfoVersion`

3. **Error Handling**:
   - Always check and return errors appropriately
   - Use `fmt.Errorf` with `%w` for error wrapping
   - Provide contextual error messages

4. **Copyright Headers**:
   - Every `.go` file MUST include:
     ```go
     // Copyright (c) 2025 Corey Hemminger
     // SPDX-License-Identifier: Apache-2.0
     ```
   - This is enforced by `.copywrite.hcl`

### Packer Plugin SDK Patterns

1. **Data Source Interface**:
   - Implement `ConfigSpec()`, `Configure()`, `OutputSpec()`, and `Execute()`
   - Use `mapstructure` tags for HCL2 mapping
   - Generate HCL2 specs with `//go:generate packer-sdc mapstructure-to-hcl2`

2. **HCL2 Integration**:
   - Use `hcldec.ObjectSpec` for specifications
   - Convert outputs using `hcl2helper.HCL2ValueFromConfig()`
   - Return `cty.Value` from `Execute()`

3. **Configuration Decoding**:
   - Use `config.Decode()` from packer-plugin-sdk
   - Handle multiple raw configuration inputs

### Testing Requirements

1. **Unit Tests** (`data_test.go`):
   - Test `Configure()` method
   - Test `Execute()` method
   - Verify all output fields are present
   - Check that values are not null or empty
   - Log detected values for debugging

2. **Acceptance Tests** (`data_acc_test.go`):
   - Run with `PACKER_ACC=1` environment variable
   - Test real-world plugin behavior
   - Use `make testacc` to run acceptance tests

3. **Test Commands**:
   - Unit tests: `make test`
   - Acceptance tests: `make testacc`
   - Plugin compatibility: `make plugin-check`

### Build and Release Process

1. **Development Build**:
   ```bash
   make dev
   ```
   - Builds binary with `-dev` prerelease tag
   - Installs plugin locally for testing
   - Binary name: `packer-plugin-host-info`

2. **Production Build**:
   ```bash
   make build
   ```
   - Creates production binary without dev tag

3. **Code Generation**:
   ```bash
   make generate
   ```
   - Runs `go generate` to create HCL2 specs
   - Generates documentation from docs source
   - Compiles web documentation

4. **Release Process**:
   - Uses GoReleaser (`.goreleaser.yml`)
   - Triggered by GitHub Actions on tag push
   - Builds for multiple OS/architecture combinations:
     - darwin (macOS): amd64, arm64
     - linux: amd64, arm64
     - windows: amd64, arm64
     - freebsd: amd64, arm64
   - Binary naming: `packer-plugin-host-info_v{VERSION}_{API_VERSION}_{OS}_{ARCH}`

### Dependencies

**Primary Dependencies**:
- `github.com/hashicorp/packer-plugin-sdk` v0.6.4 - Packer plugin framework
- `github.com/hashicorp/hcl/v2` v2.24.0 - HCL2 configuration language
- `github.com/shirou/gopsutil/v3` v3.24.5 - System information gathering
- `github.com/zclconf/go-cty` v1.17.0 - Type system for HCL2

**Dev Tools**:
- `packer-sdc` - Packer Software Development Command (code generation)
- `goreleaser` - Release automation

### Key Features to Maintain

1. **Zero Configuration**:
   - The datasource requires NO user configuration
   - Automatically detects all information

2. **Output Fields** (all strings):
   - `os_type`: OS name (darwin, linux, windows)
   - `version`: OS version/release number
   - `architecture`: CPU arch (amd64, arm64, 386)
   - `platform`: Detailed platform info
   - `family`: OS family classification

3. **Cross-Platform Support**:
   - Must work on macOS, Linux, Windows, FreeBSD
   - Must support amd64 and arm64 architectures

### Documentation Standards

1. **Code Documentation**:
   - Add godoc comments for all exported types and functions
   - Include examples in documentation when helpful

2. **User Documentation** (`docs/`):
   - Maintain accurate usage examples
   - Document all output fields
   - Keep compatibility notes up to date

3. **Example Templates** (`example/`):
   - Provide working, realistic examples
   - Show integration with other plugins (e.g., Docker)
   - Demonstrate common use cases:
     - Platform-specific builds
     - Image tagging with build metadata
     - OS-aware artifact naming

### Common Tasks

1. **Adding a New Output Field**:
   - Add field to `DatasourceOutput` struct with `mapstructure` tag
   - Update `Execute()` to populate the field
   - Run `make generate` to update HCL2 specs
   - Add tests for the new field
   - Update documentation

2. **Fixing a Bug**:
   - Write a failing test first
   - Fix the issue in the code
   - Verify test passes
   - Run `make plugin-check` to ensure compatibility

3. **Updating Dependencies**:
   - Update `go.mod` with `go get -u`
   - Run `go mod tidy`
   - Test thoroughly with `make test` and `make testacc`
   - Check for breaking changes in dependencies

### GitHub Workflows

1. **Release Workflow** (`.github/workflows/release.yml`):
   - Triggered on version tags
   - Runs tests and plugin-check
   - Builds multi-platform binaries
   - Creates GitHub release

2. **Integration Notifications**:
   - Notifies on manual and tag-based releases
   - Triggers HashiCorp integration updates

3. **Example Testing** (`test-plugin-example.yml`):
   - Tests the example Packer templates
   - Ensures examples stay functional

### Important Notes

1. **Module Path**: Always use full module path in imports:
   ```go
   "github.com/Stromweld/packer-plugin-host-info/datasource/host-info"
   ```

2. **Plugin Registration**: Register as `plugin.DEFAULT_NAME` so users reference it as `host-info` in HCL

3. **Version Updates**: Update `version/version.go` before releases

4. **Backwards Compatibility**: Maintain backwards compatibility in output fields

5. **Host Information Source**: Uses `shirou/gopsutil` for OS detection and Go's `runtime` package for architecture

### When Generating Code

- **Always** include copyright headers
- **Always** add the `//go:generate` directive for types that need HCL2 specs
- **Follow** the existing code structure and patterns
- **Use** the Packer Plugin SDK helper functions
- **Test** generated code with `make plugin-check`
- **Keep** code simple and maintainable

### Common Patterns

**Creating a Data Source Execute Method**:
```go
func (d *Datasource) Execute() (cty.Value, error) {
    // 1. Gather data (from system, APIs, etc.)
    // 2. Handle errors appropriately
    // 3. Populate DatasourceOutput struct
    // 4. Convert to cty.Value using hcl2helper
    return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
```

**Error Handling Pattern**:
```go
if err != nil {
    return cty.NullVal(cty.EmptyObject), fmt.Errorf("descriptive context: %w", err)
}
```

## Support and Resources

- **Packer Plugin SDK Documentation**: https://github.com/hashicorp/packer-plugin-sdk
- **Packer Documentation**: https://www.packer.io/docs
- **Example Plugins**: See README for list of example plugins to reference
- **GoReleaser Documentation**: https://goreleaser.com

## Questions to Consider

When adding features or making changes, ask:
1. Does this maintain backwards compatibility?
2. Does this work cross-platform (macOS, Linux, Windows)?
3. Are there tests covering this change?
4. Is the documentation updated?
5. Does this follow Packer plugin conventions?
6. Will this require a version bump?
