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

type AdvertisedIPRangesInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	Range *string `json:"range,omitempty" tf:"range,omitempty"`
}

type AdvertisedIPRangesObservation struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	Range *string `json:"range,omitempty" tf:"range,omitempty"`
}

type AdvertisedIPRangesParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	// +kubebuilder:validation:Optional
	Range *string `json:"range" tf:"range,omitempty"`
}

type BGPInitParameters struct {

	// User-specified flag to indicate which mode to use for advertisement.
	// Default value is DEFAULT.
	// Possible values are: DEFAULT, CUSTOM.
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`

	// User-specified list of prefix groups to advertise in custom mode.
	// This field can only be populated if advertiseMode is CUSTOM and
	// is advertised to all peers of the router. These groups will be
	// advertised in addition to any specified prefixes. Leave this field
	// blank to advertise no custom groups.
	// This enum field has the one valid value: ALL_SUBNETS
	AdvertisedGroups []*string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`

	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is CUSTOM and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIPRanges []AdvertisedIPRangesInitParameters `json:"advertisedIpRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`

	// Local BGP Autonomous System Number (ASN). Must be an RFC6996
	// private ASN, either 16-bit or 32-bit. The value will be fixed for
	// this router resource. All VPN tunnels that link to this router
	// will have the same local ASN.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// Explicitly specifies a range of valid BGP Identifiers for this Router.
	// It is provided as a link-local IPv4 range (from 169.254.0.0/16), of
	// size at least /30, even if the BGP sessions are over IPv6. It must
	// not overlap with any IPv4 BGP session ranges. Other vendors commonly
	// call this router ID.
	IdentifierRange *string `json:"identifierRange,omitempty" tf:"identifier_range,omitempty"`

	// The interval in seconds between BGP keepalive messages that are sent
	// to the peer. Hold time is three times the interval at which keepalive
	// messages are sent, and the hold time is the maximum number of seconds
	// allowed to elapse between successive keepalive messages that BGP
	// receives from a peer.
	// BGP will use the smaller of either the local hold time value or the
	// peer's hold time value as the hold time for the BGP connection
	// between the two peers. If set, this value must be between 20 and 60.
	// The default is 20.
	KeepaliveInterval *float64 `json:"keepaliveInterval,omitempty" tf:"keepalive_interval,omitempty"`
}

type BGPObservation struct {

	// User-specified flag to indicate which mode to use for advertisement.
	// Default value is DEFAULT.
	// Possible values are: DEFAULT, CUSTOM.
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`

	// User-specified list of prefix groups to advertise in custom mode.
	// This field can only be populated if advertiseMode is CUSTOM and
	// is advertised to all peers of the router. These groups will be
	// advertised in addition to any specified prefixes. Leave this field
	// blank to advertise no custom groups.
	// This enum field has the one valid value: ALL_SUBNETS
	AdvertisedGroups []*string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`

	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is CUSTOM and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIPRanges []AdvertisedIPRangesObservation `json:"advertisedIpRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`

	// Local BGP Autonomous System Number (ASN). Must be an RFC6996
	// private ASN, either 16-bit or 32-bit. The value will be fixed for
	// this router resource. All VPN tunnels that link to this router
	// will have the same local ASN.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// Explicitly specifies a range of valid BGP Identifiers for this Router.
	// It is provided as a link-local IPv4 range (from 169.254.0.0/16), of
	// size at least /30, even if the BGP sessions are over IPv6. It must
	// not overlap with any IPv4 BGP session ranges. Other vendors commonly
	// call this router ID.
	IdentifierRange *string `json:"identifierRange,omitempty" tf:"identifier_range,omitempty"`

	// The interval in seconds between BGP keepalive messages that are sent
	// to the peer. Hold time is three times the interval at which keepalive
	// messages are sent, and the hold time is the maximum number of seconds
	// allowed to elapse between successive keepalive messages that BGP
	// receives from a peer.
	// BGP will use the smaller of either the local hold time value or the
	// peer's hold time value as the hold time for the BGP connection
	// between the two peers. If set, this value must be between 20 and 60.
	// The default is 20.
	KeepaliveInterval *float64 `json:"keepaliveInterval,omitempty" tf:"keepalive_interval,omitempty"`
}

type BGPParameters struct {

	// User-specified flag to indicate which mode to use for advertisement.
	// Default value is DEFAULT.
	// Possible values are: DEFAULT, CUSTOM.
	// +kubebuilder:validation:Optional
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`

	// User-specified list of prefix groups to advertise in custom mode.
	// This field can only be populated if advertiseMode is CUSTOM and
	// is advertised to all peers of the router. These groups will be
	// advertised in addition to any specified prefixes. Leave this field
	// blank to advertise no custom groups.
	// This enum field has the one valid value: ALL_SUBNETS
	// +kubebuilder:validation:Optional
	AdvertisedGroups []*string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`

	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is CUSTOM and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AdvertisedIPRanges []AdvertisedIPRangesParameters `json:"advertisedIpRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`

	// Local BGP Autonomous System Number (ASN). Must be an RFC6996
	// private ASN, either 16-bit or 32-bit. The value will be fixed for
	// this router resource. All VPN tunnels that link to this router
	// will have the same local ASN.
	// +kubebuilder:validation:Optional
	Asn *float64 `json:"asn" tf:"asn,omitempty"`

	// Explicitly specifies a range of valid BGP Identifiers for this Router.
	// It is provided as a link-local IPv4 range (from 169.254.0.0/16), of
	// size at least /30, even if the BGP sessions are over IPv6. It must
	// not overlap with any IPv4 BGP session ranges. Other vendors commonly
	// call this router ID.
	// +kubebuilder:validation:Optional
	IdentifierRange *string `json:"identifierRange,omitempty" tf:"identifier_range,omitempty"`

	// The interval in seconds between BGP keepalive messages that are sent
	// to the peer. Hold time is three times the interval at which keepalive
	// messages are sent, and the hold time is the maximum number of seconds
	// allowed to elapse between successive keepalive messages that BGP
	// receives from a peer.
	// BGP will use the smaller of either the local hold time value or the
	// peer's hold time value as the hold time for the BGP connection
	// between the two peers. If set, this value must be between 20 and 60.
	// The default is 20.
	// +kubebuilder:validation:Optional
	KeepaliveInterval *float64 `json:"keepaliveInterval,omitempty" tf:"keepalive_interval,omitempty"`
}

type RouterInitParameters struct {

	// BGP information specific to this router.
	// Structure is documented below.
	BGP *BGPInitParameters `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `json:"encryptedInterconnectRouter,omitempty" tf:"encrypted_interconnect_router,omitempty"`

	// A reference to the network to which this router belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type RouterObservation struct {

	// BGP information specific to this router.
	// Structure is documented below.
	BGP *BGPObservation `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `json:"encryptedInterconnectRouter,omitempty" tf:"encrypted_interconnect_router,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/routers/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A reference to the network to which this router belongs.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router resides.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RouterParameters struct {

	// BGP information specific to this router.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BGP *BGPParameters `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	// +kubebuilder:validation:Optional
	EncryptedInterconnectRouter *bool `json:"encryptedInterconnectRouter,omitempty" tf:"encrypted_interconnect_router,omitempty"`

	// A reference to the network to which this router belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router resides.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// RouterSpec defines the desired state of Router
type RouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterParameters `json:"forProvider"`
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
	InitProvider RouterInitParameters `json:"initProvider,omitempty"`
}

// RouterStatus defines the observed state of Router.
type RouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Router is the Schema for the Routers API. Represents a Router resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Router struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterSpec   `json:"spec"`
	Status            RouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterList contains a list of Routers
type RouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Router `json:"items"`
}

// Repository type metadata.
var (
	Router_Kind             = "Router"
	Router_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Router_Kind}.String()
	Router_KindAPIVersion   = Router_Kind + "." + CRDGroupVersion.String()
	Router_GroupVersionKind = CRDGroupVersion.WithKind(Router_Kind)
)

func init() {
	SchemeBuilder.Register(&Router{}, &RouterList{})
}
