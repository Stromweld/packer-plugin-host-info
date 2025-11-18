# Copyright (c) 2025 Corey Hemminger
# SPDX-License-Identifier: Apache-2.0

# This data source detects the host operating system, version, and architecture
# No configuration is required - it automatically detects information about the
# system where Packer is running
data "host-info" "current" {
  # No configuration required - automatically detects host OS information
}

