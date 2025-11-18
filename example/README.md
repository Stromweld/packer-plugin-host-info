# Host Info Plugin Example

This folder contains a complete working example demonstrating the use of the Packer Host Info plugin.

## Overview

The example shows how to:
- Use the `host-info` data source to detect host system information
- Access OS, version, architecture, platform, and family data
- Use detected values in locals and build configurations
- Apply host information as Docker image labels
- Create platform-aware builds

## Files

- **data.pkr.hcl** - Defines the `host-info` data source
- **variables.pkr.hcl** - Shows how to use the detected values in locals
- **build.pkr.hcl** - Complete build example with Docker

## Usage

### Install the Plugin

First, build and install the plugin:

```bash
cd ..
go build -o packer-plugin-host-info
packer plugins install --path packer-plugin-host-info github.com/Stromweld/host-info
```

### Initialize Packer

```bash
cd example
packer init .
```

### Validate the Configuration

```bash
packer validate .
```

### Build

```bash
packer build .
```

## What It Does

The example will:
1. Detect your host operating system information
2. Display the detected values (OS, version, architecture, platform, family)
3. Build a Docker image tagged with your host platform information
4. Apply metadata labels showing the build environment

## Expected Output

You'll see output similar to:

```
Building on: darwin 26.1
Architecture: arm64
Platform: darwin
Family: Standalone Workstation
OS-Arch Combo: darwin-arm64
```

## Customization

You can modify the example to:
- Use different builders (AWS, Azure, GCP, etc.)
- Apply conditional logic based on your host OS
- Create OS-specific build artifacts
- Tag images with build metadata
