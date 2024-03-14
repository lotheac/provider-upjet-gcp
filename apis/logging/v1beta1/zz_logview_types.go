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

type LogViewInitParameters struct {

	// Describes this view.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type LogViewObservation struct {

	// The bucket of the resource
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Output only. The creation timestamp of the view.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Describes this view.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// an identifier for the resource with format {{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The parent of the resource.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Output only. The last update timestamp of the view.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type LogViewParameters struct {

	// The bucket of the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/logging/v1beta1.ProjectBucketConfig
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a ProjectBucketConfig in logging to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a ProjectBucketConfig in logging to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Describes this view.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The parent of the resource.
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

// LogViewSpec defines the desired state of LogView
type LogViewSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogViewParameters `json:"forProvider"`
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
	InitProvider LogViewInitParameters `json:"initProvider,omitempty"`
}

// LogViewStatus defines the observed state of LogView.
type LogViewStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogViewObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogView is the Schema for the LogViews API. Describes a view over log entries in a bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type LogView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogViewSpec   `json:"spec"`
	Status            LogViewStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogViewList contains a list of LogViews
type LogViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogView `json:"items"`
}

// Repository type metadata.
var (
	LogView_Kind             = "LogView"
	LogView_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogView_Kind}.String()
	LogView_KindAPIVersion   = LogView_Kind + "." + CRDGroupVersion.String()
	LogView_GroupVersionKind = CRDGroupVersion.WithKind(LogView_Kind)
)

func init() {
	SchemeBuilder.Register(&LogView{}, &LogViewList{})
}
