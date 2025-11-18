# Packer Plugin Host Info

A Packer plugin that provides a data source to detect the host operating system, version, and CPU architecture where Packer is running.

## Features

This plugin provides a data source that automatically detects:
- **OS**: The operating system name (e.g., `darwin`, `linux`, `windows`)
- **Version**: The OS version/release number
- **Architecture**: The CPU architecture (e.g., `amd64`, `arm64`)
- **Platform**: Detailed platform information
- **Family**: OS family classification

## Usage

Add the plugin to your Packer template:

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
  # No configuration required
}

locals {
  # Use the detected values in your build
  image_name = "my-image-${data.host-info.current.os_type}-${data.host-info.current.architecture}"
}
```

## Data Source Output

The `host-info` data source provides the following output attributes:

- `os_type` (string): The host operating system (e.g., `darwin`, `linux`, `windows`)
- `version` (string): The OS version/release
- `architecture` (string): The CPU architecture (e.g., `amd64`, `arm64`, `386`)
- `platform` (string): Platform information (e.g., `darwin`, `ubuntu`, `rhel`)
- `family` (string): OS family (e.g., `standalone`, `debian`, `rhel`)

## Example

See the [example](example) directory for a complete working example.

In this repository you will also find a pre-defined GitHub Action configuration for the release workflow
(`.goreleaser.yml` and `.github/workflows/release.yml`). The release workflow configuration makes sure the GitHub
release artifacts are created with the correct binaries and naming conventions.

Please see the [GitHub template repository documentation](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template)
for how to create a new repository from this template on GitHub.

## Packer plugin projects

Here's a non exaustive list of Packer plugins that you can checkout:

* [github.com/hashicorp/packer-plugin-docker](https://github.com/hashicorp/packer-plugin-docker)
* [github.com/exoscale/packer-plugin-exoscale](https://github.com/exoscale/packer-plugin-exoscale)
* [github.com/sylviamoss/packer-plugin-comment](https://github.com/sylviamoss/packer-plugin-comment)
* [github.com/hashicorp/packer-plugin-hashicups](https://github.com/hashicorp/packer-plugin-hashicups)

Looking at their code will give you good examples.

## Installation

### From Source

1. Clone this GitHub repository locally:
   ```shell
   git clone https://github.com/Stromweld/packer-plugin-host-info
   cd packer-plugin-host-info
   ```

2. Build the plugin binary:
   ```shell 
   go build -ldflags="-X github.com/Stromweld/packer-plugin-host-info/version.VersionPrerelease=dev" -o packer-plugin-host-info
   ```

3. Install the compiled plugin:
   ```shell
   packer plugins install --path packer-plugin-host-info github.com/Stromweld/host-info
   ```

### Build on *nix systems
Unix like systems with the make, sed, and grep commands installed can use the `make dev` to execute the build from source steps. 

### Build on Windows Powershell
The preferred solution for building on Windows are steps 2-4 listed above.
If you would prefer to script the building process you can use the following as a guide

```powershell
$MODULE_NAME = (Get-Content go.mod | Where-Object { $_ -match "^module"  }) -replace 'module ',''
$FQN = $MODULE_NAME -replace 'packer-plugin-',''
go build -ldflags="-X $MODULE_NAME/version.VersionPrerelease=dev" -o packer-plugin-host-info.exe
packer plugins install --path packer-plugin-host-info.exe $FQN
```

## Running Acceptance Tests

Make sure to install the plugin locally using the steps in [Build from source](#build-from-source).

Once everything needed is set up, run:
```
PACKER_ACC=1 go test -count 1 -v ./... -timeout=120m
```

This will run the acceptance tests for all plugins in this set.

## Registering Plugin as Packer Integration

Partner and community plugins can be hard to find if a user doesn't know what 
they are looking for. To assist with plugin discovery Packer offers an integration
portal at https://developer.hashicorp.com/packer/integrations to list known integrations 
that work with the latest release of Packer. 

Registering a plugin as an integration requires [metadata configuration](./metadata.hcl) within the plugin
repository and approval by the Packer team. To initiate the process of registering your 
plugin as a Packer integration refer to the [Developing Plugins](https://developer.hashicorp.com/packer/docs/plugins/creation#registering-plugins) page.

# Requirements

-	[packer-plugin-sdk](https://github.com/hashicorp/packer-plugin-sdk) >= v0.5.2
-	[Go](https://golang.org/doc/install) >= 1.20

## Packer Compatibility
This plugin is compatible with Packer >= v1.10.2
