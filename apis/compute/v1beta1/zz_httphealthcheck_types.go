// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HTTPHealthCheckInitParameters_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HTTPHealthCheckObservation_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/httpHealthChecks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HTTPHealthCheckParameters_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	// +kubebuilder:validation:Optional
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	// +kubebuilder:validation:Optional
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

// HTTPHealthCheckSpec defines the desired state of HTTPHealthCheck
type HTTPHealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPHealthCheckParameters_2 `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HTTPHealthCheckInitParameters_2 `json:"initProvider,omitempty"`
}

// HTTPHealthCheckStatus defines the observed state of HTTPHealthCheck.
type HTTPHealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPHealthCheckObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HTTPHealthCheck is the Schema for the HTTPHealthChecks API. An HttpHealthCheck resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HTTPHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HTTPHealthCheckSpec   `json:"spec"`
	Status            HTTPHealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPHealthCheckList contains a list of HTTPHealthChecks
type HTTPHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPHealthCheck `json:"items"`
}

// Repository type metadata.
var (
	HTTPHealthCheck_Kind             = "HTTPHealthCheck"
	HTTPHealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPHealthCheck_Kind}.String()
	HTTPHealthCheck_KindAPIVersion   = HTTPHealthCheck_Kind + "." + CRDGroupVersion.String()
	HTTPHealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(HTTPHealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTPHealthCheck{}, &HTTPHealthCheckList{})
}
