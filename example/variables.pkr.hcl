# Copyright (c) 2025 Corey Hemminger
# SPDX-License-Identifier: Apache-2.0

locals {
  # Host OS information from the host-info data source
  host_os          = data.host-info.current.os_type
  host_version     = data.host-info.current.version
  host_arch        = data.host-info.current.architecture
  host_platform    = data.host-info.current.platform
  host_family      = data.host-info.current.family

  # Useful derived values
  os_arch_combo    = "${local.host_os}-${local.host_arch}"
  build_timestamp  = formatdate("YYYY-MM-DD-hhmm", timestamp())
}
