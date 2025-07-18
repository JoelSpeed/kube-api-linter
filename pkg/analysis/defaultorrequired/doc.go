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
/*
The `defaultorrequired` linter checks that fields that are marked as required, do not have a default value.

Any field marked with any of  `+required`, `+kubebuilder:validation:Required`, `+k8s:required`,
must not be marked with any of `+default`, `+kubebuilder:default`.

Defaulting occurs after validation, and so fields that are marked as required can not be defaulted.
*/
package defaultorrequired
