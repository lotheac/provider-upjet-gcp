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

type RuntimeIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type RuntimeIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type RuntimeIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type RuntimeIAMMemberInitParameters struct {
	Condition []RuntimeIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Runtime
	RuntimeName *string `json:"runtimeName,omitempty" tf:"runtime_name,omitempty"`

	// Reference to a Runtime to populate runtimeName.
	// +kubebuilder:validation:Optional
	RuntimeNameRef *v1.Reference `json:"runtimeNameRef,omitempty" tf:"-"`

	// Selector for a Runtime to populate runtimeName.
	// +kubebuilder:validation:Optional
	RuntimeNameSelector *v1.Selector `json:"runtimeNameSelector,omitempty" tf:"-"`
}

type RuntimeIAMMemberObservation struct {
	Condition []RuntimeIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	RuntimeName *string `json:"runtimeName,omitempty" tf:"runtime_name,omitempty"`
}

type RuntimeIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []RuntimeIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Runtime
	// +kubebuilder:validation:Optional
	RuntimeName *string `json:"runtimeName,omitempty" tf:"runtime_name,omitempty"`

	// Reference to a Runtime to populate runtimeName.
	// +kubebuilder:validation:Optional
	RuntimeNameRef *v1.Reference `json:"runtimeNameRef,omitempty" tf:"-"`

	// Selector for a Runtime to populate runtimeName.
	// +kubebuilder:validation:Optional
	RuntimeNameSelector *v1.Selector `json:"runtimeNameSelector,omitempty" tf:"-"`
}

// RuntimeIAMMemberSpec defines the desired state of RuntimeIAMMember
type RuntimeIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuntimeIAMMemberParameters `json:"forProvider"`
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
	InitProvider RuntimeIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// RuntimeIAMMemberStatus defines the observed state of RuntimeIAMMember.
type RuntimeIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuntimeIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RuntimeIAMMember is the Schema for the RuntimeIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RuntimeIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   RuntimeIAMMemberSpec   `json:"spec"`
	Status RuntimeIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuntimeIAMMemberList contains a list of RuntimeIAMMembers
type RuntimeIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuntimeIAMMember `json:"items"`
}

// Repository type metadata.
var (
	RuntimeIAMMember_Kind             = "RuntimeIAMMember"
	RuntimeIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuntimeIAMMember_Kind}.String()
	RuntimeIAMMember_KindAPIVersion   = RuntimeIAMMember_Kind + "." + CRDGroupVersion.String()
	RuntimeIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(RuntimeIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&RuntimeIAMMember{}, &RuntimeIAMMemberList{})
}
