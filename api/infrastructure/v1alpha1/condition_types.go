package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConditionType is a valid value for Condition.Type.
type ConditionType string

// Common ConditionTypes used by Cluster API objects.
const (
	// ReadyCondition defines the Ready condition type that summarizes the operational state of the vcluster API object.
	ReadyCondition ConditionType = "Ready"

	// ControlPlaneInitializedCondition defines the initialized condition type if the vcluster is reachable.
	ControlPlaneInitializedCondition ConditionType = "ControlPlaneInitialized"

	// KubeconfigReadyCondition defines the ready condition type if the vcluster kubeconfig was written.
	KubeconfigReadyCondition ConditionType = "KubeconfigReady"

	// HelmChartDeployedCondition defines the helm chart deployed condition type that defines if the helm chart was deployed correctly.
	HelmChartDeployedCondition ConditionType = "HelmChartDeployed"

	// InfrastructureClusterSyncedCondition defines the infrastructure cluster synced condition type that defines if the infrastructure cluster was patched correctly.
	InfrastructureClusterSyncedCondition ConditionType = "InfrastructureClusterSynced"
)

// ConditionSeverity expresses the severity of a Condition Type failing.
type ConditionSeverity string

// Condition severity levels
const (
	// ConditionSeverityError specifies that a condition with `Status=False` is an error.
	ConditionSeverityError ConditionSeverity = "Error"

	// ConditionSeverityWarning specifies that a condition with `Status=False` is a warning.
	ConditionSeverityWarning ConditionSeverity = "Warning"

	// ConditionSeverityInfo specifies that a condition with `Status=False` is informative.
	ConditionSeverityInfo ConditionSeverity = "Info"

	// ConditionSeverityNone should apply only to conditions with `Status=True`.
	ConditionSeverityNone ConditionSeverity = ""
)

// Conditions is an array of conditions
type Conditions []Condition

// Condition defines an observation of a Cluster API resource operational state.
type Condition struct {
	// Type of condition in CamelCase or in foo.example.com/CamelCase.
	// Many .condition.type values are consistent across resources like Available, but because arbitrary conditions
	// can be useful (see .node.status.conditions), the ability to deconflict is important.
	// +required
	Type ConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	// +required
	Status corev1.ConditionStatus `json:"status"`

	// Severity provides an explicit classification of Reason code, so the users or machines can immediately
	// understand the current situation and act accordingly.
	// The Severity field MUST be set only when Status=False.
	// +optional
	Severity ConditionSeverity `json:"severity,omitempty"`

	// Last time the condition transitioned from one status to another.
	// This should be when the underlying condition changed. If that is not known, then using the time when
	// the API field changed is acceptable.
	// +required
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// The reason for the condition's last transition in CamelCase.
	// The specific API may choose whether this field is considered a guaranteed API.
	// This field may not be empty.
	// +optional
	Reason string `json:"reason,omitempty"`

	// A human readable message indicating details about the transition.
	// This field may be empty.
	// +optional
	Message string `json:"message,omitempty"`
}
