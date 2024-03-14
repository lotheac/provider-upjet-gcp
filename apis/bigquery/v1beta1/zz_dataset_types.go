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

type AccessDatasetInitParameters struct {

	// The dataset this entry applies to
	// Structure is documented below.
	Dataset []DatasetDatasetInitParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Which resources in the dataset this entry applies to. Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	TargetTypes []*string `json:"targetTypes,omitempty" tf:"target_types,omitempty"`
}

type AccessDatasetObservation struct {

	// The dataset this entry applies to
	// Structure is documented below.
	Dataset []DatasetDatasetObservation `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Which resources in the dataset this entry applies to. Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	TargetTypes []*string `json:"targetTypes,omitempty" tf:"target_types,omitempty"`
}

type AccessDatasetParameters struct {

	// The dataset this entry applies to
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Dataset []DatasetDatasetParameters `json:"dataset" tf:"dataset,omitempty"`

	// Which resources in the dataset this entry applies to. Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	// +kubebuilder:validation:Optional
	TargetTypes []*string `json:"targetTypes" tf:"target_types,omitempty"`
}

type AccessInitParameters struct {

	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	// Structure is documented below.
	Dataset []AccessDatasetInitParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// An email address of a Google Group to grant access to.
	GroupByEmail *string `json:"groupByEmail,omitempty" tf:"group_by_email,omitempty"`

	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: allUsers
	IAMMember *string `json:"iamMember,omitempty" tf:"iam_member,omitempty"`

	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles
	// are supported. Predefined roles that have equivalent basic roles
	// are swapped by the API to their basic counterparts. See
	// official docs.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// A routine from a different dataset to grant access to. Queries
	// executed against that routine will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that routine is updated by any user, access to the routine
	// needs to be granted again via an update operation.
	// Structure is documented below.
	Routine []RoutineInitParameters `json:"routine,omitempty" tf:"routine,omitempty"`

	// A special group to grant access to. Possible values include:
	SpecialGroup *string `json:"specialGroup,omitempty" tf:"special_group,omitempty"`

	// An email address of a user to grant access to. For example:
	// fred@example.com
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("email",true)
	UserByEmail *string `json:"userByEmail,omitempty" tf:"user_by_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate userByEmail.
	// +kubebuilder:validation:Optional
	UserByEmailRef *v1.Reference `json:"userByEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate userByEmail.
	// +kubebuilder:validation:Optional
	UserByEmailSelector *v1.Selector `json:"userByEmailSelector,omitempty" tf:"-"`

	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View []ViewInitParameters `json:"view,omitempty" tf:"view,omitempty"`
}

type AccessObservation struct {

	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	// Structure is documented below.
	Dataset []AccessDatasetObservation `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// An email address of a Google Group to grant access to.
	GroupByEmail *string `json:"groupByEmail,omitempty" tf:"group_by_email,omitempty"`

	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: allUsers
	IAMMember *string `json:"iamMember,omitempty" tf:"iam_member,omitempty"`

	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles
	// are supported. Predefined roles that have equivalent basic roles
	// are swapped by the API to their basic counterparts. See
	// official docs.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// A routine from a different dataset to grant access to. Queries
	// executed against that routine will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that routine is updated by any user, access to the routine
	// needs to be granted again via an update operation.
	// Structure is documented below.
	Routine []RoutineObservation `json:"routine,omitempty" tf:"routine,omitempty"`

	// A special group to grant access to. Possible values include:
	SpecialGroup *string `json:"specialGroup,omitempty" tf:"special_group,omitempty"`

	// An email address of a user to grant access to. For example:
	// fred@example.com
	UserByEmail *string `json:"userByEmail,omitempty" tf:"user_by_email,omitempty"`

	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View []ViewObservation `json:"view,omitempty" tf:"view,omitempty"`
}

type AccessParameters struct {

	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Dataset []AccessDatasetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// An email address of a Google Group to grant access to.
	// +kubebuilder:validation:Optional
	GroupByEmail *string `json:"groupByEmail,omitempty" tf:"group_by_email,omitempty"`

	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: allUsers
	// +kubebuilder:validation:Optional
	IAMMember *string `json:"iamMember,omitempty" tf:"iam_member,omitempty"`

	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles
	// are supported. Predefined roles that have equivalent basic roles
	// are swapped by the API to their basic counterparts. See
	// official docs.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// A routine from a different dataset to grant access to. Queries
	// executed against that routine will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that routine is updated by any user, access to the routine
	// needs to be granted again via an update operation.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Routine []RoutineParameters `json:"routine,omitempty" tf:"routine,omitempty"`

	// A special group to grant access to. Possible values include:
	// +kubebuilder:validation:Optional
	SpecialGroup *string `json:"specialGroup,omitempty" tf:"special_group,omitempty"`

	// An email address of a user to grant access to. For example:
	// fred@example.com
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	UserByEmail *string `json:"userByEmail,omitempty" tf:"user_by_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate userByEmail.
	// +kubebuilder:validation:Optional
	UserByEmailRef *v1.Reference `json:"userByEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate userByEmail.
	// +kubebuilder:validation:Optional
	UserByEmailSelector *v1.Selector `json:"userByEmailSelector,omitempty" tf:"-"`

	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	View []ViewParameters `json:"view,omitempty" tf:"view,omitempty"`
}

type DatasetDatasetInitParameters struct {

	// The ID of the dataset containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The ID of the project containing this table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type DatasetDatasetObservation struct {

	// The ID of the dataset containing this table.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The ID of the project containing this table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type DatasetDatasetParameters struct {

	// The ID of the dataset containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The ID of the project containing this table.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

type DatasetInitParameters struct {

	// An array of objects that define dataset access for one or more entities.
	// Structure is documented below.
	Access []AccessInitParameters `json:"access,omitempty" tf:"access,omitempty"`

	// Defines the default collation specification of future tables created
	// in the dataset. If a table is created in this dataset without table-level
	// default collation, then the table inherits the dataset default collation,
	// which is applied to the string fields that do not have explicit collation
	// specified. A change to this field affects only tables created afterwards,
	// and does not alter the existing tables.
	// The following values are supported:
	DefaultCollation *string `json:"defaultCollation,omitempty" tf:"default_collation,omitempty"`

	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.
	// Structure is documented below.
	DefaultEncryptionConfiguration []DefaultEncryptionConfigurationInitParameters `json:"defaultEncryptionConfiguration,omitempty" tf:"default_encryption_configuration,omitempty"`

	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs *float64 `json:"defaultPartitionExpirationMs,omitempty" tf:"default_partition_expiration_ms,omitempty"`

	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs *float64 `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms,omitempty"`

	// If set to true, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy *bool `json:"deleteContentsOnDestroy,omitempty" tf:"delete_contents_on_destroy,omitempty"`

	// A user-friendly description of the dataset
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A descriptive name for the dataset
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
	// By default, this is FALSE, which means the dataset and its table names are
	// case-sensitive. This field does not affect routine references.
	IsCaseInsensitive *bool `json:"isCaseInsensitive,omitempty" tf:"is_case_insensitive,omitempty"`

	// The labels associated with this dataset. You can use these to
	// organize and group your datasets.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The geographic location where the dataset should reside.
	// See official docs.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
	MaxTimeTravelHours *string `json:"maxTimeTravelHours,omitempty" tf:"max_time_travel_hours,omitempty"`

	// Specifies the storage billing model for the dataset.
	// Set this flag value to LOGICAL to use logical bytes for storage billing,
	// or to PHYSICAL to use physical bytes instead.
	// LOGICAL is the default if this flag isn't specified.
	StorageBillingModel *string `json:"storageBillingModel,omitempty" tf:"storage_billing_model,omitempty"`
}

type DatasetObservation struct {

	// An array of objects that define dataset access for one or more entities.
	// Structure is documented below.
	Access []AccessObservation `json:"access,omitempty" tf:"access,omitempty"`

	// The time when this dataset was created, in milliseconds since the
	// epoch.
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Defines the default collation specification of future tables created
	// in the dataset. If a table is created in this dataset without table-level
	// default collation, then the table inherits the dataset default collation,
	// which is applied to the string fields that do not have explicit collation
	// specified. A change to this field affects only tables created afterwards,
	// and does not alter the existing tables.
	// The following values are supported:
	DefaultCollation *string `json:"defaultCollation,omitempty" tf:"default_collation,omitempty"`

	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.
	// Structure is documented below.
	DefaultEncryptionConfiguration []DefaultEncryptionConfigurationObservation `json:"defaultEncryptionConfiguration,omitempty" tf:"default_encryption_configuration,omitempty"`

	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	DefaultPartitionExpirationMs *float64 `json:"defaultPartitionExpirationMs,omitempty" tf:"default_partition_expiration_ms,omitempty"`

	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	DefaultTableExpirationMs *float64 `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms,omitempty"`

	// If set to true, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	DeleteContentsOnDestroy *bool `json:"deleteContentsOnDestroy,omitempty" tf:"delete_contents_on_destroy,omitempty"`

	// A user-friendly description of the dataset
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// A hash of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A descriptive name for the dataset
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/datasets/{{dataset_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
	// By default, this is FALSE, which means the dataset and its table names are
	// case-sensitive. This field does not affect routine references.
	IsCaseInsensitive *bool `json:"isCaseInsensitive,omitempty" tf:"is_case_insensitive,omitempty"`

	// The labels associated with this dataset. You can use these to
	// organize and group your datasets.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The date when this dataset or any of its tables was last modified, in
	// milliseconds since the epoch.
	LastModifiedTime *float64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	// The geographic location where the dataset should reside.
	// See official docs.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
	MaxTimeTravelHours *string `json:"maxTimeTravelHours,omitempty" tf:"max_time_travel_hours,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Specifies the storage billing model for the dataset.
	// Set this flag value to LOGICAL to use logical bytes for storage billing,
	// or to PHYSICAL to use physical bytes instead.
	// LOGICAL is the default if this flag isn't specified.
	StorageBillingModel *string `json:"storageBillingModel,omitempty" tf:"storage_billing_model,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`
}

type DatasetParameters struct {

	// An array of objects that define dataset access for one or more entities.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Access []AccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// Defines the default collation specification of future tables created
	// in the dataset. If a table is created in this dataset without table-level
	// default collation, then the table inherits the dataset default collation,
	// which is applied to the string fields that do not have explicit collation
	// specified. A change to this field affects only tables created afterwards,
	// and does not alter the existing tables.
	// The following values are supported:
	// +kubebuilder:validation:Optional
	DefaultCollation *string `json:"defaultCollation,omitempty" tf:"default_collation,omitempty"`

	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DefaultEncryptionConfiguration []DefaultEncryptionConfigurationParameters `json:"defaultEncryptionConfiguration,omitempty" tf:"default_encryption_configuration,omitempty"`

	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	// +kubebuilder:validation:Optional
	DefaultPartitionExpirationMs *float64 `json:"defaultPartitionExpirationMs,omitempty" tf:"default_partition_expiration_ms,omitempty"`

	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	// +kubebuilder:validation:Optional
	DefaultTableExpirationMs *float64 `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms,omitempty"`

	// If set to true, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	// +kubebuilder:validation:Optional
	DeleteContentsOnDestroy *bool `json:"deleteContentsOnDestroy,omitempty" tf:"delete_contents_on_destroy,omitempty"`

	// A user-friendly description of the dataset
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A descriptive name for the dataset
	// +kubebuilder:validation:Optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
	// By default, this is FALSE, which means the dataset and its table names are
	// case-sensitive. This field does not affect routine references.
	// +kubebuilder:validation:Optional
	IsCaseInsensitive *bool `json:"isCaseInsensitive,omitempty" tf:"is_case_insensitive,omitempty"`

	// The labels associated with this dataset. You can use these to
	// organize and group your datasets.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The geographic location where the dataset should reside.
	// See official docs.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
	// +kubebuilder:validation:Optional
	MaxTimeTravelHours *string `json:"maxTimeTravelHours,omitempty" tf:"max_time_travel_hours,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the storage billing model for the dataset.
	// Set this flag value to LOGICAL to use logical bytes for storage billing,
	// or to PHYSICAL to use physical bytes instead.
	// LOGICAL is the default if this flag isn't specified.
	// +kubebuilder:validation:Optional
	StorageBillingModel *string `json:"storageBillingModel,omitempty" tf:"storage_billing_model,omitempty"`
}

type DefaultEncryptionConfigurationInitParameters struct {

	// Describes the Cloud KMS encryption key that will be used to protect destination
	// BigQuery table. The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

type DefaultEncryptionConfigurationObservation struct {

	// Describes the Cloud KMS encryption key that will be used to protect destination
	// BigQuery table. The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type DefaultEncryptionConfigurationParameters struct {

	// Describes the Cloud KMS encryption key that will be used to protect destination
	// BigQuery table. The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

type RoutineInitParameters struct {

	// The ID of the dataset containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Routine
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("dataset_id",false)
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Routine in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Routine in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The ID of the project containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Routine
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project",false)
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Routine in bigquery to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Routine in bigquery to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// The ID of the routine. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 256 characters.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Routine
	RoutineID *string `json:"routineId,omitempty" tf:"routine_id,omitempty"`

	// Reference to a Routine in bigquery to populate routineId.
	// +kubebuilder:validation:Optional
	RoutineIDRef *v1.Reference `json:"routineIdRef,omitempty" tf:"-"`

	// Selector for a Routine in bigquery to populate routineId.
	// +kubebuilder:validation:Optional
	RoutineIDSelector *v1.Selector `json:"routineIdSelector,omitempty" tf:"-"`
}

type RoutineObservation struct {

	// The ID of the dataset containing this table.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The ID of the project containing this table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The ID of the routine. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 256 characters.
	RoutineID *string `json:"routineId,omitempty" tf:"routine_id,omitempty"`
}

type RoutineParameters struct {

	// The ID of the dataset containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Routine
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("dataset_id",false)
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Routine in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Routine in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The ID of the project containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Routine
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Routine in bigquery to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Routine in bigquery to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// The ID of the routine. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 256 characters.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Routine
	// +kubebuilder:validation:Optional
	RoutineID *string `json:"routineId,omitempty" tf:"routine_id,omitempty"`

	// Reference to a Routine in bigquery to populate routineId.
	// +kubebuilder:validation:Optional
	RoutineIDRef *v1.Reference `json:"routineIdRef,omitempty" tf:"-"`

	// Selector for a Routine in bigquery to populate routineId.
	// +kubebuilder:validation:Optional
	RoutineIDSelector *v1.Selector `json:"routineIdSelector,omitempty" tf:"-"`
}

type ViewInitParameters struct {

	// The ID of the dataset containing this table.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The ID of the project containing this table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The ID of the table. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type ViewObservation struct {

	// The ID of the dataset containing this table.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The ID of the project containing this table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The ID of the table. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type ViewParameters struct {

	// The ID of the dataset containing this table.
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// The ID of the project containing this table.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// The ID of the table. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	// +kubebuilder:validation:Optional
	TableID *string `json:"tableId" tf:"table_id,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetParameters `json:"forProvider"`
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
	InitProvider DatasetInitParameters `json:"initProvider,omitempty"`
}

// DatasetStatus defines the observed state of Dataset.
type DatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Dataset is the Schema for the Datasets API. Datasets allow you to organize and control access to your tables.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasetSpec   `json:"spec"`
	Status            DatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetList contains a list of Datasets
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

// Repository type metadata.
var (
	Dataset_Kind             = "Dataset"
	Dataset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dataset_Kind}.String()
	Dataset_KindAPIVersion   = Dataset_Kind + "." + CRDGroupVersion.String()
	Dataset_GroupVersionKind = CRDGroupVersion.WithKind(Dataset_Kind)
)

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}
