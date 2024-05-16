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

type ReservationAssignmentInitParameters struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	Assignee *string `json:"assignee,omitempty" tf:"assignee,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	JobType *string `json:"jobType,omitempty" tf:"job_type,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ReservationAssignmentObservation struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	Assignee *string `json:"assignee,omitempty" tf:"assignee,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	JobType *string `json:"jobType,omitempty" tf:"job_type,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Output only. The resource name of the assignment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The reservation for the resource
	Reservation *string `json:"reservation,omitempty" tf:"reservation,omitempty"`

	// Assignment will remain in PENDING state if no active capacity commitment is present. It will become ACTIVE when some capacity commitment becomes active. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ReservationAssignmentParameters struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	// +kubebuilder:validation:Optional
	Assignee *string `json:"assignee,omitempty" tf:"assignee,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	// +kubebuilder:validation:Optional
	JobType *string `json:"jobType,omitempty" tf:"job_type,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The reservation for the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Reservation
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Reservation *string `json:"reservation,omitempty" tf:"reservation,omitempty"`

	// Reference to a Reservation in bigquery to populate reservation.
	// +kubebuilder:validation:Optional
	ReservationRef *v1.Reference `json:"reservationRef,omitempty" tf:"-"`

	// Selector for a Reservation in bigquery to populate reservation.
	// +kubebuilder:validation:Optional
	ReservationSelector *v1.Selector `json:"reservationSelector,omitempty" tf:"-"`
}

// ReservationAssignmentSpec defines the desired state of ReservationAssignment
type ReservationAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationAssignmentParameters `json:"forProvider"`
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
	InitProvider ReservationAssignmentInitParameters `json:"initProvider,omitempty"`
}

// ReservationAssignmentStatus defines the observed state of ReservationAssignment.
type ReservationAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ReservationAssignment is the Schema for the ReservationAssignments API. The BigqueryReservation Assignment resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ReservationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assignee) || (has(self.initProvider) && has(self.initProvider.assignee))",message="spec.forProvider.assignee is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.jobType) || (has(self.initProvider) && has(self.initProvider.jobType))",message="spec.forProvider.jobType is a required parameter"
	Spec   ReservationAssignmentSpec   `json:"spec"`
	Status ReservationAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationAssignmentList contains a list of ReservationAssignments
type ReservationAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReservationAssignment `json:"items"`
}

// Repository type metadata.
var (
	ReservationAssignment_Kind             = "ReservationAssignment"
	ReservationAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReservationAssignment_Kind}.String()
	ReservationAssignment_KindAPIVersion   = ReservationAssignment_Kind + "." + CRDGroupVersion.String()
	ReservationAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ReservationAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ReservationAssignment{}, &ReservationAssignmentList{})
}
