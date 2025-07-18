package a

type TestStruct struct {
	// Valid field with only optional marker
	// +optional
	ValidOptionalField string `json:"validOptionalField,omitempty"`

	// Valid field with only required marker
	// +required
	ValidRequiredField string `json:"validRequiredField"`

	// Conflict: default vs required
	// +default
	// +required
	DefaultVsRequiredField string `json:"defaultVsRequiredField"` // want "field DefaultVsRequiredField has conflicting markers: defaultorrequired: \\[default\\] and \\[required\\]. A field with a default value cannot be required"

	// Multiple conflicts with multiple markers in each set:
	// - optional set: +optional, +kubebuilder:validation:Optional, +k8s:optional
	// - required set: +required, +kubebuilder:validation:Required, +k8s:required
	// - default set: +default, +kubebuilder:default
	// +default
	// +kubebuilder:default
	// +required
	// +kubebuilder:validation:Required
	// +k8s:required
	MultipleConflictsMultipleMarkersField string `json:"multipleConflictsMultipleMarkersField"` // want "field MultipleConflictsMultipleMarkersField has conflicting markers: defaultorrequired: \\[default kubebuilder:default\\] and \\[k8s:required kubebuilder:validation:Required required\\]. A field with a default value cannot be required"
}
