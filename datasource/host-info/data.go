// Copyright (c) 2025 Corey Hemminger
// SPDX-License-Identifier: Apache-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package hostinfo

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	// This data source doesn't require any configuration
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	OsType       string `mapstructure:"os_type"`
	Version      string `mapstructure:"version"`
	Architecture string `mapstructure:"architecture"`
	Platform     string `mapstructure:"platform"`
	Family       string `mapstructure:"family"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	// Get host information using gopsutil
	info, err := host.Info()
	if err != nil {
		return cty.NullVal(cty.EmptyObject), fmt.Errorf("failed to get host information: %w", err)
	}

	// Get CPU architecture from runtime
	arch := runtime.GOARCH

	// Normalize OS name
	osName := strings.ToLower(runtime.GOOS)

	output := DatasourceOutput{
		OsType:       osName,
		Version:      info.PlatformVersion,
		Architecture: arch,
		Platform:     info.Platform,
		Family:       info.PlatformFamily,
	}

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
