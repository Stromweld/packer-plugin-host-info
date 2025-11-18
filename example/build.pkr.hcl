# Copyright (c) 2025 Corey Hemminger
# SPDX-License-Identifier: Apache-2.0

packer {
  required_plugins {
    hostinfo = {
      version = ">=v0.1.0"
      source  = "github.com/Stromweld/host-info"
    }
    docker = {
      version = ">=v1.0.0"
      source  = "github.com/hashicorp/docker"
    }
  }
}

# This example demonstrates using the host-info data source to:
# 1. Tag Docker images with build environment metadata
# 2. Make OS-specific decisions in your builds
# 3. Create platform-aware artifact naming

source "docker" "example" {
  image  = "ubuntu:22.04"
  commit = true
  changes = [
    "LABEL builder.os=${local.host_os}",
    "LABEL builder.version=${local.host_version}",
    "LABEL builder.architecture=${local.host_arch}",
    "LABEL builder.platform=${local.host_platform}",
    "LABEL builder.timestamp=${local.build_timestamp}",
  ]
}

build {
  name = "hostos-example"
  sources = ["source.docker.example"]

  provisioner "shell" {
    inline = [
      "echo 'Building on: ${local.host_os} ${local.host_version}'",
      "echo 'Architecture: ${local.host_arch}'",
      "echo 'Platform: ${local.host_platform}'",
      "echo 'Family: ${local.host_family}'",
      "echo 'OS-Arch Combo: ${local.os_arch_combo}'",
    ]
  }

  post-processor "docker-tag" {
    repository = "example/hostos-demo"
    tags = [
      "latest",
      "${local.os_arch_combo}",
      "${local.build_timestamp}",
    ]
  }
}
