# Copyright (c) 2025 Corey Hemminger
# SPDX-License-Identifier: Apache-2.0

data "host-info" "test" {
  # No configuration required
}

locals {
  host_os      = data.host-info.test.os_type
  host_version = data.host-info.test.version
  host_arch    = data.host-info.test.architecture
  host_platform = data.host-info.test.platform
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo Host OS: ${local.host_os}",
      "echo Host Version: ${local.host_version}",
      "echo Host Architecture: ${local.host_arch}",
      "echo Host Platform: ${local.host_platform}",
    ]
  }
}
