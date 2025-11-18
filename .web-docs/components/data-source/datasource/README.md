  Include a short description about the data source. This is a good place
  to call out what the data source does, and any requirements for the given
  data source environment. See https://www.packer.io/docs/data-source/amazon-ami
-->

The `host-info` data source automatically detects the host operating system, version, 
and CPU architecture where Packer is running. This is useful for creating dynamic build 
configurations that adapt to the build environment, such as generating OS-specific image 
names or conditionally running platform-specific provisioners.

This data source requires no configuration and has no external dependencies beyond the 
Packer SDK.


<!-- Data source Configuration Fields -->

## Configuration Reference

This data source does not require any configuration parameters. It automatically detects 
the host system information when executed.


## Output Reference

The `host-info` data source exports the following attributes:

- `os_type` (string) - The host operating system name (e.g., `darwin`, `linux`, `windows`).

- `version` (string) - The operating system version or release number.

- `architecture` (string) - The CPU architecture (e.g., `amd64`, `arm64`, `386`).

- `platform` (string) - Detailed platform information (e.g., `darwin`, `ubuntu`, `rhel`).

- `family` (string) - The operating system family classification (e.g., `standalone`, `debian`, `rhel`).


## Example Usage

### Basic Usage

Detect the current host operating system and use it in build configuration:

```hcl
data "host-info" "current" {
  # No configuration required - automatically detects host OS
}

locals {
  # Use detected values to create dynamic image names
  image_name = "my-app-${data.host-info.current.os_type}-${data.host-info.current.architecture}"
}

source "docker" "example" {
  image  = "ubuntu:22.04"
  commit = true
}

build {
  sources = ["source.docker.example"]
  
  provisioner "shell" {
    inline = [
      "echo 'Building on ${data.host-info.current.os_type}'",
      "echo 'Host architecture: ${data.host-info.current.architecture}'",
      "echo 'Platform: ${data.host-info.current.platform}'"
    ]
  }
  
  post-processor "docker-tag" {
    repository = "myapp"
    tags       = [local.image_name]
  }
}
```

### Conditional Provisioning

Use host OS information to conditionally run provisioners:

```hcl
data "host-info" "current" {}

source "null" "example" {
  communicator = "none"
}

build {
  sources = ["source.null.example"]
  
  # Only run on macOS
  provisioner "shell-local" {
    only   = data.host-info.current.os_type == "darwin" ? ["null.example"] : []
    inline = ["echo 'Running on macOS'"]
  }
  
  # Only run on Linux
  provisioner "shell-local" {
    only   = data.host-info.current.os_type == "linux" ? ["null.example"] : []
    inline = ["echo 'Running on Linux'"]
  }
}
```

### Dynamic Variable Selection

Select build variables based on the host operating system:

```hcl
data "host-info" "current" {}

locals {
  builder_config = {
    darwin = {
      vm_name = "macos-builder"
      cpus    = 4
    }
    linux = {
      vm_name = "linux-builder"
      cpus    = 8
    }
    windows = {
      vm_name = "windows-builder"
      cpus    = 4
    }
  }
  
  current_config = local.builder_config[data.host-info.current.os_type]
}

source "virtualbox-iso" "example" {
  vm_name       = local.current_config.vm_name
  cpus          = local.current_config.cpus
  iso_url       = "https://example.com/os.iso"
  iso_checksum  = "sha256:..."
  ssh_username  = "packer"
  ssh_password  = "packer"
  shutdown_command = "shutdown -h now"
}

build {
  sources = ["source.virtualbox-iso.example"]
}
```
