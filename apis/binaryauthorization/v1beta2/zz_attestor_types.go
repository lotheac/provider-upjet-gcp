// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttestationAuthorityNoteInitParameters struct {

	// The resource name of a ATTESTATION_AUTHORITY Note, created by the
	// user. If the Note is in a different project from the Attestor, it
	// should be specified in the format projects/*/notes/* (or the legacy
	// providers/*/notes/*). This field may not be updated.
	// An attestation by this attestor is stored as a Container Analysis
	// ATTESTATION_AUTHORITY Occurrence that names a container image
	// and that links to this Note.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/containeranalysis/v1beta2.Note
	NoteReference *string `json:"noteReference,omitempty" tf:"note_reference,omitempty"`

	// Reference to a Note in containeranalysis to populate noteReference.
	// +kubebuilder:validation:Optional
	NoteReferenceRef *v1.Reference `json:"noteReferenceRef,omitempty" tf:"-"`

	// Selector for a Note in containeranalysis to populate noteReference.
	// +kubebuilder:validation:Optional
	NoteReferenceSelector *v1.Selector `json:"noteReferenceSelector,omitempty" tf:"-"`

	// Public keys that verify attestations signed by this attestor. This
	// field may be updated.
	// If this field is non-empty, one of the specified public keys must
	// verify that an attestation was signed by this attestor for the
	// image specified in the admission request.
	// If this field is empty, this attestor always returns that no valid
	// attestations exist.
	// Structure is documented below.
	PublicKeys []PublicKeysInitParameters `json:"publicKeys,omitempty" tf:"public_keys,omitempty"`
}

type AttestationAuthorityNoteObservation struct {

	// (Output)
	// This field will contain the service account email address that
	// this Attestor will use as the principal when querying Container
	// Analysis. Attestor administrators must grant this service account
	// the IAM role needed to read attestations from the noteReference in
	// Container Analysis (containeranalysis.notes.occurrences.viewer).
	// This email address is fixed for the lifetime of the Attestor, but
	// callers should not make any other assumptions about the service
	// account email; future versions may use an email based on a
	// different naming pattern.
	DelegationServiceAccountEmail *string `json:"delegationServiceAccountEmail,omitempty" tf:"delegation_service_account_email,omitempty"`

	// The resource name of a ATTESTATION_AUTHORITY Note, created by the
	// user. If the Note is in a different project from the Attestor, it
	// should be specified in the format projects/*/notes/* (or the legacy
	// providers/*/notes/*). This field may not be updated.
	// An attestation by this attestor is stored as a Container Analysis
	// ATTESTATION_AUTHORITY Occurrence that names a container image
	// and that links to this Note.
	NoteReference *string `json:"noteReference,omitempty" tf:"note_reference,omitempty"`

	// Public keys that verify attestations signed by this attestor. This
	// field may be updated.
	// If this field is non-empty, one of the specified public keys must
	// verify that an attestation was signed by this attestor for the
	// image specified in the admission request.
	// If this field is empty, this attestor always returns that no valid
	// attestations exist.
	// Structure is documented below.
	PublicKeys []PublicKeysObservation `json:"publicKeys,omitempty" tf:"public_keys,omitempty"`
}

type AttestationAuthorityNoteParameters struct {

	// The resource name of a ATTESTATION_AUTHORITY Note, created by the
	// user. If the Note is in a different project from the Attestor, it
	// should be specified in the format projects/*/notes/* (or the legacy
	// providers/*/notes/*). This field may not be updated.
	// An attestation by this attestor is stored as a Container Analysis
	// ATTESTATION_AUTHORITY Occurrence that names a container image
	// and that links to this Note.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/containeranalysis/v1beta2.Note
	// +kubebuilder:validation:Optional
	NoteReference *string `json:"noteReference,omitempty" tf:"note_reference,omitempty"`

	// Reference to a Note in containeranalysis to populate noteReference.
	// +kubebuilder:validation:Optional
	NoteReferenceRef *v1.Reference `json:"noteReferenceRef,omitempty" tf:"-"`

	// Selector for a Note in containeranalysis to populate noteReference.
	// +kubebuilder:validation:Optional
	NoteReferenceSelector *v1.Selector `json:"noteReferenceSelector,omitempty" tf:"-"`

	// Public keys that verify attestations signed by this attestor. This
	// field may be updated.
	// If this field is non-empty, one of the specified public keys must
	// verify that an attestation was signed by this attestor for the
	// image specified in the admission request.
	// If this field is empty, this attestor always returns that no valid
	// attestations exist.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PublicKeys []PublicKeysParameters `json:"publicKeys,omitempty" tf:"public_keys,omitempty"`
}

type AttestorInitParameters struct {

	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote *AttestationAuthorityNoteInitParameters `json:"attestationAuthorityNote,omitempty" tf:"attestation_authority_note,omitempty"`

	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type AttestorObservation struct {

	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote *AttestationAuthorityNoteObservation `json:"attestationAuthorityNote,omitempty" tf:"attestation_authority_note,omitempty"`

	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/attestors/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type AttestorParameters struct {

	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AttestationAuthorityNote *AttestationAuthorityNoteParameters `json:"attestationAuthorityNote,omitempty" tf:"attestation_authority_note,omitempty"`

	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type PkixPublicKeyInitParameters struct {

	// A PEM-encoded public key, as described in
	// https://tools.ietf.org/html/rfc7468#section-13
	PublicKeyPem *string `json:"publicKeyPem,omitempty" tf:"public_key_pem,omitempty"`

	// The signature algorithm used to verify a message against
	// a signature using this key. These signature algorithm must
	// match the structure and any object identifiers encoded in
	// publicKeyPem (i.e. this algorithm must match that of the
	// public key).
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm,omitempty"`
}

type PkixPublicKeyObservation struct {

	// A PEM-encoded public key, as described in
	// https://tools.ietf.org/html/rfc7468#section-13
	PublicKeyPem *string `json:"publicKeyPem,omitempty" tf:"public_key_pem,omitempty"`

	// The signature algorithm used to verify a message against
	// a signature using this key. These signature algorithm must
	// match the structure and any object identifiers encoded in
	// publicKeyPem (i.e. this algorithm must match that of the
	// public key).
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm,omitempty"`
}

type PkixPublicKeyParameters struct {

	// A PEM-encoded public key, as described in
	// https://tools.ietf.org/html/rfc7468#section-13
	// +kubebuilder:validation:Optional
	PublicKeyPem *string `json:"publicKeyPem,omitempty" tf:"public_key_pem,omitempty"`

	// The signature algorithm used to verify a message against
	// a signature using this key. These signature algorithm must
	// match the structure and any object identifiers encoded in
	// publicKeyPem (i.e. this algorithm must match that of the
	// public key).
	// +kubebuilder:validation:Optional
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm,omitempty"`
}

type PublicKeysInitParameters struct {

	// ASCII-armored representation of a PGP public key, as the
	// entire output by the command
	// gpg --export --armor foo@example.com (either LF or CRLF
	// line endings). When using this field, id should be left
	// blank. The BinAuthz API handlers will calculate the ID
	// and fill it in automatically. BinAuthz computes this ID
	// as the OpenPGP RFC4880 V4 fingerprint, represented as
	// upper-case hex. If id is provided by the caller, it will
	// be overwritten by the API-calculated ID.
	ASCIIArmoredPgpPublicKey *string `json:"asciiArmoredPgpPublicKey,omitempty" tf:"ascii_armored_pgp_public_key,omitempty"`

	// A descriptive comment. This field may be updated.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The ID of this public key. Signatures verified by BinAuthz
	// must include the ID of the public key that can be used to
	// verify them, and that ID must match the contents of this
	// field exactly. Additional restrictions on this field can
	// be imposed based on which public key type is encapsulated.
	// See the documentation on publicKey cases below for details.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A raw PKIX SubjectPublicKeyInfo format public key.
	// NOTE: id may be explicitly provided by the caller when using this
	// type of public key, but it MUST be a valid RFC3986 URI. If id is left
	// blank, a default one will be computed based on the digest of the DER
	// encoding of the public key.
	// Structure is documented below.
	PkixPublicKey *PkixPublicKeyInitParameters `json:"pkixPublicKey,omitempty" tf:"pkix_public_key,omitempty"`
}

type PublicKeysObservation struct {

	// ASCII-armored representation of a PGP public key, as the
	// entire output by the command
	// gpg --export --armor foo@example.com (either LF or CRLF
	// line endings). When using this field, id should be left
	// blank. The BinAuthz API handlers will calculate the ID
	// and fill it in automatically. BinAuthz computes this ID
	// as the OpenPGP RFC4880 V4 fingerprint, represented as
	// upper-case hex. If id is provided by the caller, it will
	// be overwritten by the API-calculated ID.
	ASCIIArmoredPgpPublicKey *string `json:"asciiArmoredPgpPublicKey,omitempty" tf:"ascii_armored_pgp_public_key,omitempty"`

	// A descriptive comment. This field may be updated.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The ID of this public key. Signatures verified by BinAuthz
	// must include the ID of the public key that can be used to
	// verify them, and that ID must match the contents of this
	// field exactly. Additional restrictions on this field can
	// be imposed based on which public key type is encapsulated.
	// See the documentation on publicKey cases below for details.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A raw PKIX SubjectPublicKeyInfo format public key.
	// NOTE: id may be explicitly provided by the caller when using this
	// type of public key, but it MUST be a valid RFC3986 URI. If id is left
	// blank, a default one will be computed based on the digest of the DER
	// encoding of the public key.
	// Structure is documented below.
	PkixPublicKey *PkixPublicKeyObservation `json:"pkixPublicKey,omitempty" tf:"pkix_public_key,omitempty"`
}

type PublicKeysParameters struct {

	// ASCII-armored representation of a PGP public key, as the
	// entire output by the command
	// gpg --export --armor foo@example.com (either LF or CRLF
	// line endings). When using this field, id should be left
	// blank. The BinAuthz API handlers will calculate the ID
	// and fill it in automatically. BinAuthz computes this ID
	// as the OpenPGP RFC4880 V4 fingerprint, represented as
	// upper-case hex. If id is provided by the caller, it will
	// be overwritten by the API-calculated ID.
	// +kubebuilder:validation:Optional
	ASCIIArmoredPgpPublicKey *string `json:"asciiArmoredPgpPublicKey,omitempty" tf:"ascii_armored_pgp_public_key,omitempty"`

	// A descriptive comment. This field may be updated.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The ID of this public key. Signatures verified by BinAuthz
	// must include the ID of the public key that can be used to
	// verify them, and that ID must match the contents of this
	// field exactly. Additional restrictions on this field can
	// be imposed based on which public key type is encapsulated.
	// See the documentation on publicKey cases below for details.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A raw PKIX SubjectPublicKeyInfo format public key.
	// NOTE: id may be explicitly provided by the caller when using this
	// type of public key, but it MUST be a valid RFC3986 URI. If id is left
	// blank, a default one will be computed based on the digest of the DER
	// encoding of the public key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PkixPublicKey *PkixPublicKeyParameters `json:"pkixPublicKey,omitempty" tf:"pkix_public_key,omitempty"`
}

// AttestorSpec defines the desired state of Attestor
type AttestorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttestorParameters `json:"forProvider"`
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
	InitProvider AttestorInitParameters `json:"initProvider,omitempty"`
}

// AttestorStatus defines the observed state of Attestor.
type AttestorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttestorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Attestor is the Schema for the Attestors API. An attestor that attests to container image artifacts.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Attestor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.attestationAuthorityNote) || (has(self.initProvider) && has(self.initProvider.attestationAuthorityNote))",message="spec.forProvider.attestationAuthorityNote is a required parameter"
	Spec   AttestorSpec   `json:"spec"`
	Status AttestorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttestorList contains a list of Attestors
type AttestorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attestor `json:"items"`
}

// Repository type metadata.
var (
	Attestor_Kind             = "Attestor"
	Attestor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attestor_Kind}.String()
	Attestor_KindAPIVersion   = Attestor_Kind + "." + CRDGroupVersion.String()
	Attestor_GroupVersionKind = CRDGroupVersion.WithKind(Attestor_Kind)
)

func init() {
	SchemeBuilder.Register(&Attestor{}, &AttestorList{})
}
