// Copyright (c) 2025 Corey Hemminger
// SPDX-License-Identifier: Apache-2.0

package hostinfo

import (
	"testing"
)

func TestDatasource_Execute(t *testing.T) {
	d := &Datasource{}

	// Execute the data source
	result, err := d.Execute()
	if err != nil {
		t.Fatalf("Execute() returned error: %v", err)
	}

	if result.IsNull() {
		t.Fatal("Execute() returned null value")
	}

	// Verify the result has the expected structure
	resultMap := result.AsValueMap()

	expectedFields := []string{"os_type", "version", "architecture", "platform", "family"}
	for _, field := range expectedFields {
		if _, exists := resultMap[field]; !exists {
			t.Errorf("Expected field %q not found in result", field)
		}
	}

	// Verify OS field is not empty
	osValue := resultMap["os_type"]
	if osValue.IsNull() || osValue.AsString() == "" {
		t.Error("OS field should not be empty")
	}

	// Verify Architecture field is not empty
	archValue := resultMap["architecture"]
	if archValue.IsNull() || archValue.AsString() == "" {
		t.Error("Architecture field should not be empty")
	}

	t.Logf("Host OS: %s", osValue.AsString())
	t.Logf("Architecture: %s", archValue.AsString())
	t.Logf("Version: %s", resultMap["version"].AsString())
	t.Logf("Platform: %s", resultMap["platform"].AsString())
	t.Logf("Family: %s", resultMap["family"].AsString())
}

func TestDatasource_Configure(t *testing.T) {
	d := &Datasource{}

	// Configure should work with no inputs since Config is empty
	err := d.Configure()
	if err != nil {
		t.Fatalf("Configure() returned error: %v", err)
	}
}
