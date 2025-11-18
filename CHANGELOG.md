# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - Unreleased

### Added
- Initial release of the Host Info plugin
- `host-info` data source for detecting host operating system information
- Support for detecting:
  - Operating system name (os)
  - OS version/release (version)
  - CPU architecture (architecture)
  - Platform information (platform)
  - OS family classification (family)
- Comprehensive documentation and examples
- Unit tests and acceptance tests
- Cross-platform support (Linux, macOS, Windows, FreeBSD, etc.)

### Dependencies
- github.com/hashicorp/packer-plugin-sdk v0.6.1
- github.com/shirou/gopsutil/v3 v3.24.5
- Go 1.23.2+

[0.1.0]: https://github.com/Stromweld/packer-plugin-host-info/releases/tag/v0.1.0

