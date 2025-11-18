# Packer Plugin Host Info - Quick Start Guide

## Overview

This Packer plugin provides a data source that automatically detects the host operating system, version, and CPU architecture where Packer is running.

## Installation

### Build from Source

```bash
cd /Users/corey.hemminger/github/personal/other/packer-plugin-host-info
go build -o packer-plugin-host-info
packer plugins install --path packer-plugin-host-info github.com/Stromweld/host-info
```

## Usage

### Basic Example

Create a Packer template file (e.g., `example.pkr.hcl`):

```hcl
packer {
  required_plugins {
    host-info = {
      version = ">= v0.1.0"
      source  = "github.com/Stromweld/host-info"
    }
  }
}

data "host-info" "current" {
  # No configuration required - automatically detects host OS
}

locals {
  # Access the detected values
  os_name  = data.host-info.current.os
  os_ver   = data.host-info.current.version
  cpu_arch = data.host-info.current.architecture
  platform = data.host-info.current.platform
  family   = data.host-info.current.family
}

# Example: Display the values
build {
  sources = ["null.example"]
  
  provisioner "shell-local" {
    inline = [
      "echo 'Building on: ${local.os_name} ${local.os_ver}'",
      "echo 'Architecture: ${local.cpu_arch}'",
      "echo 'Platform: ${local.platform}'",
      "echo 'Family: ${local.family}'"
    ]
  }
}
```

## Data Source Outputs

| Output | Type | Description | Example Values |
|--------|------|-------------|----------------|
| `os` | string | Operating system name | `darwin`, `linux`, `windows` |
| `version` | string | OS version/release | `26.1`, `22.04`, `10.0.19045` |
| `architecture` | string | CPU architecture | `amd64`, `arm64`, `386` |
| `platform` | string | Platform information | `darwin`, `ubuntu`, `rhel` |
| `family` | string | OS family | `Standalone Workstation`, `debian`, `rhel` |

## Real-World Examples

### Conditional Docker Image Building

```hcl
data "host-info" "current" {}

source "docker" "app" {
  image  = "ubuntu:22.04"
  commit = true
  changes = [
    "LABEL builder.os=${data.host-info.current.os}",
    "LABEL builder.arch=${data.host-info.current.architecture}",
    "LABEL builder.date=${timestamp()}"
  ]
}
```

### OS-Specific Build Paths

```hcl
data "host-info" "current" {}

locals {
  artifact_path = data.host-info.current.os == "windows" ? "C:/artifacts" : "/tmp/artifacts"
}
```

### Architecture-Specific Image Selection

```hcl
data "host-info" "current" {}

source "amazon-ebs" "app" {
  ami_name = "myapp-${data.host-info.current.architecture}-{{timestamp}}"
  
  source_ami_filter {
    filters = {
      architecture = data.host-info.current.architecture
    }
  }
}
```

## Testing

Run the unit tests:

```bash
cd datasource/scaffolding
go test -v
```

Expected output:
```
=== RUN   TestDatasource_Execute
    data_test.go:45: Host OS: darwin
    data_test.go:46: Architecture: arm64
    data_test.go:47: Version: 26.1
    data_test.go:48: Platform: darwin
    data_test.go:49: Family: Standalone Workstation
--- PASS: TestDatasource_Execute (0.03s)
```

## Implementation Details

The plugin uses:
- Go's `runtime` package for OS and architecture detection
- `github.com/shirou/gopsutil/v3/host` for detailed host information
- Packer Plugin SDK for integration with Packer

## Supported Platforms

- **Operating Systems**: Linux, macOS (Darwin), Windows, FreeBSD, etc.
- **Architectures**: amd64, arm64, 386, arm, and more

## Troubleshooting

If you encounter issues:

1. Verify Go dependencies are installed:
   ```bash
   go mod download
   go mod tidy
   ```

2. Rebuild the plugin:
   ```bash
   go build -v
   ```

3. Check that the plugin binary exists:
   ```bash
   ls -lh packer-plugin-host-info
   ```

4. Run tests to verify functionality:
   ```bash
   go test ./... -v
   ```

