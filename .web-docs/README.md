<!--
  Include a short overview about the plugin.

  This document is a great location for creating a table of contents for each
  of the components the plugin may provide. This document should load automatically
  when navigating to the docs directory for a plugin.

-->

The Host Info plugin provides a data source for detecting the host operating system, version, 
and CPU architecture where Packer is running. This is useful for creating dynamic build 
configurations that adapt to the build environment.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    hostinfo = {
      source  = "github.com/Stromweld/host-info"
      version = "~> 0.1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/Stromweld/host-info
```

### Components

The Host Info plugin provides the following component:

#### Data Sources

- [host-info](/packer/integrations/stromweld/host-info/latest/components/data-source/datasource) - Automatically detects the host operating system, version, CPU architecture, platform, and OS family where Packer is running.

