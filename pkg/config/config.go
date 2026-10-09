/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

// GolangCIConfig is the complete configuration for the KAL
// linter when built as an integration into golangci-lint.
type GolangCIConfig struct {
	// Preset defines the preset profile to be used for the linter config.
	// This will be used to set the default values for other configuration
	// including but not limited to which linters are enabled/disabled and
	// the configuration for individual linters.
	// Valid values are "Native" and "CustomResource".
	// When not specified, the default value is "Native".
	Preset ConfigPreset `mapstructure:"preset"`

	// Linters allows the user to configure which linters should,
	// and should not be enabled.
	Linters Linters `mapstructure:"linters"`

	// LintersConfig contains configuration for individual linters.
	LintersConfig LintersConfig `mapstructure:"lintersConfig"`
}

// ConfigPreset is the preset profile to be used for the linter config.
type ConfigPreset string

const (
	// ConfigPresetNative is the preset profile for native types.
	ConfigPresetNative ConfigPreset = "Native"

	// ConfigPresetCustomResource is the preset profile for custom resource types.
	ConfigPresetCustomResource ConfigPreset = "CustomResource"
)
