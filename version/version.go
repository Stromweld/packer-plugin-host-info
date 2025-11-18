// Copyright (c) 2025 Corey Hemminger
// SPDX-License-Identifier: Apache-2.0

package version

import "github.com/hashicorp/packer-plugin-sdk/version"

var (
	Version           = "0.0.2"
	VersionPrerelease = "dev"
	VersionMetadata   = ""
	PluginVersion     = version.NewPluginVersion(Version, VersionPrerelease, VersionMetadata)
)
