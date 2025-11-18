// Copyright (c) 2025 Corey Hemminger
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	hostInfoData "github.com/Stromweld/packer-plugin-host-info/datasource/hostinfo"
	hostInfoVersion "github.com/Stromweld/packer-plugin-host-info/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("host-info", new(hostInfoData.Datasource))
	pps.SetVersion(hostInfoVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
