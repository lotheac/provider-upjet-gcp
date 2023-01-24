/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationObservation struct {

	// Identifier of the app, usually {PROJECT_ID}
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// The GCS bucket code is being stored in for this app.
	CodeBucket *string `json:"codeBucket,omitempty" tf:"code_bucket,omitempty"`

	// The GCS bucket content is being stored in for this app.
	DefaultBucket *string `json:"defaultBucket,omitempty" tf:"default_bucket,omitempty"`

	// The default hostname for this app.
	DefaultHostname *string `json:"defaultHostname,omitempty" tf:"default_hostname,omitempty"`

	// The GCR domain used for storing managed Docker images for this app.
	GcrDomain *string `json:"gcrDomain,omitempty" tf:"gcr_domain,omitempty"`

	// an identifier for the resource with format {{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique name of the app, usually apps/{PROJECT_ID}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of dispatch rule blocks. Each block has a domain, path, and service field.
	URLDispatchRule []URLDispatchRuleObservation `json:"urlDispatchRule,omitempty" tf:"url_dispatch_rule,omitempty"`
}

type ApplicationParameters struct {

	// The domain to authenticate users with when using App Engine's User API.
	// +kubebuilder:validation:Optional
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	// Can be CLOUD_FIRESTORE or CLOUD_DATASTORE_COMPATIBILITY for new
	// instances.  To support old instances, the value CLOUD_DATASTORE is accepted
	// by the provider, but will be rejected by the API.
	// +kubebuilder:validation:Optional
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// A block of optional settings to configure specific App Engine features:
	// +kubebuilder:validation:Optional
	FeatureSettings []FeatureSettingsParameters `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`

	// Settings for enabling Cloud Identity Aware Proxy
	// +kubebuilder:validation:Optional
	Iap []IapParameters `json:"iap,omitempty" tf:"iap,omitempty"`

	// The location
	// to serve the app from.
	// +kubebuilder:validation:Required
	LocationID *string `json:"locationId" tf:"location_id,omitempty"`

	// The project ID to create the application under.
	// ~>NOTE: GCP only accepts project ID, not project number. If you are using number,
	// you may get a "Permission denied" error.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("project_id",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// The serving status of the app.
	// +kubebuilder:validation:Optional
	ServingStatus *string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
}

type FeatureSettingsObservation struct {
}

type FeatureSettingsParameters struct {

	// Set to false to use the legacy health check instead of the readiness
	// and liveness checks.
	// +kubebuilder:validation:Required
	SplitHealthChecks *bool `json:"splitHealthChecks" tf:"split_health_checks,omitempty"`
}

type IapObservation struct {
}

type IapParameters struct {

	// Whether the serving infrastructure will authenticate and authorize all incoming requests.
	// (default is false)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// OAuth2 client ID to use for the authentication flow.
	// +kubebuilder:validation:Required
	Oauth2ClientID *string `json:"oauth2ClientId" tf:"oauth2_client_id,omitempty"`

	// OAuth2 client secret to use for the authentication flow.
	// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	// +kubebuilder:validation:Required
	Oauth2ClientSecretSecretRef v1.SecretKeySelector `json:"oauth2ClientSecretSecretRef" tf:"-"`
}

type URLDispatchRuleObservation struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type URLDispatchRuleParameters struct {
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Allows management of an App Engine application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
