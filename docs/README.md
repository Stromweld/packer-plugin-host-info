# Packer Plugin Host Info

The Host Info plugin provides a data source that automatically detects information about the host operating system where Packer is running.

## Components

### Data Sources

- [info](datasources/info.mdx) - Detects host operating system, version, architecture, platform, and OS family

## Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    host-info = {
      source  = "github.com/Stromweld/host-info"
      version = ">=0.1.0"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/Stromweld/host-info
```

## Usage

```hcl
data "host-info" "current" {
  # No configuration required
}

locals {
  build_platform = "${data.host-info.current.os}-${data.host-info.current.architecture}"
}
```

### Components

#### Data Sources

- [data source](/packer/integrations/hashicorp/scaffolding/latest/components/datasource/datasource-name) - The scaffolding data source is used to
  export scaffolding data.

