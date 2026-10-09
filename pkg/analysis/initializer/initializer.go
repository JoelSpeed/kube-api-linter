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
package initializer

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

// ConfigurableInitializerFunc is a function that initializes an Analyzer.
type InitializerFunc[T any] func(config.ConfigPreset, *T) (*analysis.Analyzer, error)

// ValidateFunc is a function that validates the configuration for an Analyzer.
type ValidateFunc[T any] func(*T, *field.Path) field.ErrorList

// AnalyzerInitializer is used to initialize analyzers.
type AnalyzerInitializer interface {
	// Name returns the name of the analyzer initialized by this initializer.
	Name() string

	// Init returns the newly initialized analyzer.
	// It will be passed the complete LintersConfig and is expected to rely only on its own configuration.
	Init(config.ConfigPreset, any) (*analysis.Analyzer, error)

	// Default determines whether the initializer initializes an analyzer that should be
	// on by default, or not.
	DefaultNative() bool

	// DefaultCustomResource determines whether the initializer initializes an analyzer that should be
	// on by default for custom resource types, or not.
	DefaultCustomResource() bool
}

// ConfigurableAnalyzerInitializer is an analyzer initializer that also has a configuration.
// This means it can validate its config.
type ConfigurableAnalyzerInitializer interface {
	AnalyzerInitializer

	// ConfigType returns the type of the config for the linter.
	ConfigType() any

	// ValidateConfig will be called during the config validation phase and is used to validate
	// the provided config for the linter.
	ValidateConfig(any, *field.Path) field.ErrorList
}

// NewInitializer construct a new initializer for initializing an Analyzer.
func NewInitializer(name string, analyzer *analysis.Analyzer, defaultNative, defaultCustomResource bool) AnalyzerInitializer {
	return initializer[any]{
		name:                  name,
		analyzer:              analyzer,
		defaultNative:         defaultNative,
		defaultCustomResource: defaultCustomResource,
	}
}

// NewConfigurableInitializer constructs a new initializer for initializing a
// configurable Analyzer.
func NewConfigurableInitializer[T any](name string, initFunc InitializerFunc[T], defaultNative, defaultCustomResource bool, validateFunc ValidateFunc[T]) ConfigurableAnalyzerInitializer {
	return configurableInitializer[T]{
		initializer: initializer[T]{
			name:                  name,
			defaultNative:         defaultNative,
			defaultCustomResource: defaultCustomResource,
		},
		initFunc:     initFunc,
		validateFunc: validateFunc,
	}
}

type initializer[T any] struct {
	name                  string
	analyzer              *analysis.Analyzer
	defaultNative         bool
	defaultCustomResource bool
}

// Name returns the name of the initializer.
func (i initializer[T]) Name() string {
	return i.name
}

// Init returns a newly initialized analyzer.
func (i initializer[T]) Init(_ config.ConfigPreset, _ any) (*analysis.Analyzer, error) {
	return i.analyzer, nil
}

// Default determines whether this initializer should be enabled by default or not.
func (i initializer[T]) DefaultNative() bool {
	return i.defaultNative
}

// DefaultCustomResource determines whether this initializer should be enabled by default for custom resource types or not.
func (i initializer[T]) DefaultCustomResource() bool {
	return i.defaultCustomResource
}

type configurableInitializer[T any] struct {
	initializer[T]

	initFunc     InitializerFunc[T]
	validateFunc ValidateFunc[T]
}

// Init returns a newly initialized analyzer.
func (i configurableInitializer[T]) Init(preset config.ConfigPreset, cfg any) (*analysis.Analyzer, error) {
	cfgT, ok := cfg.(*T)
	if !ok {
		return nil, fmt.Errorf("failed to initialize analyzer: %w", NewIncorrectTypeError(cfg))
	}

	return i.initFunc(preset, cfgT)
}

// ConfigType returns the type of the config for the linter.
func (i configurableInitializer[T]) ConfigType() any {
	return new(T)
}

// ValidateConfig validates the configuration for the initializer.
func (i configurableInitializer[T]) ValidateConfig(cfg any, fld *field.Path) field.ErrorList {
	cfgT, ok := cfg.(*T)
	if !ok {
		return field.ErrorList{field.InternalError(fld, NewIncorrectTypeError(cfg))}
	}

	return i.validateFunc(cfgT, fld)
}
