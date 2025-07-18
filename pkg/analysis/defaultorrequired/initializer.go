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
package defaultorrequired

import (
	"golang.org/x/tools/go/analysis"

	"k8s.io/apimachinery/pkg/util/validation/field"

	"sigs.k8s.io/kube-api-linter/pkg/analysis/conflictingmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/registry"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

func init() {
	registry.DefaultRegistry().RegisterLinter(Initializer())
}

const name = "defaultorrequired"

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewInitializer(
		name,
		initAnalyzer(),
		true,
	)
}

// initAnalyzer returns the initialized Analyzer.
func initAnalyzer() *analysis.Analyzer {
	cfg := &conflictingmarkers.ConflictingMarkersConfig{
		ConflictSets: []conflictingmarkers.ConflictSet{
			{
				Name:        "defaultorrequired",
				SetA:        []string{markersconsts.DefaultMarker, markersconsts.KubebuilderDefaultMarker},
				SetB:        []string{markersconsts.RequiredMarker, markersconsts.KubebuilderRequiredMarker, markersconsts.K8sRequiredMarker},
				Description: "A field with a default value cannot be required",
			},
		},
	}

	conflictInit, ok := conflictingmarkers.Initializer().(initializer.ConfigurableAnalyzerInitializer)
	if !ok {
		panic("conflictingmarkers initializer is not configurable")
	}

	// If the config we have provided is not valid, panic, as this is a bug in our config above.
	if errs := conflictInit.ValidateConfig(cfg, field.NewPath("defaultorrequired")); len(errs) > 0 {
		panic(errs.ToAggregate())
	}

	analyzer, err := conflictInit.Init(cfg)
	if err != nil {
		panic(err)
	}

	return analyzer
}
