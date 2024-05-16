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

type GRPCHealthCheckInitParameters struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type GRPCHealthCheckObservation struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type GRPCHealthCheckParameters struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	// +kubebuilder:validation:Optional
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type HTTPHealthCheckInitParameters struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPHealthCheckObservation struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPHealthCheckParameters struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPSHealthCheckInitParameters struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPSHealthCheckObservation struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPSHealthCheckParameters struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HealthCheckInitParameters struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	GRPCHealthCheck *GRPCHealthCheckInitParameters `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPHealthCheck *HTTPHealthCheckInitParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPSHealthCheck *HTTPSHealthCheckInitParameters `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	Http2HealthCheck *Http2HealthCheckInitParameters `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	LogConfig *HealthCheckLogConfigInitParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A nested object resource
	// Structure is documented below.
	SSLHealthCheck *SSLHealthCheckInitParameters `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	TCPHealthCheck *TCPHealthCheckInitParameters `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckLogConfigInitParameters struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type HealthCheckLogConfigObservation struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type HealthCheckLogConfigParameters struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type HealthCheckObservation struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	GRPCHealthCheck *GRPCHealthCheckObservation `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPHealthCheck *HTTPHealthCheckObservation `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPSHealthCheck *HTTPSHealthCheckObservation `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	Http2HealthCheck *Http2HealthCheckObservation `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/healthChecks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	LogConfig *HealthCheckLogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A nested object resource
	// Structure is documented below.
	SSLHealthCheck *SSLHealthCheckObservation `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A nested object resource
	// Structure is documented below.
	TCPHealthCheck *TCPHealthCheckObservation `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckParameters struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	// +kubebuilder:validation:Optional
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GRPCHealthCheck *GRPCHealthCheckParameters `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPHealthCheck *HTTPHealthCheckParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPSHealthCheck *HTTPSHealthCheckParameters `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Http2HealthCheck *Http2HealthCheckParameters `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig *HealthCheckLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SSLHealthCheck *SSLHealthCheckParameters `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TCPHealthCheck *TCPHealthCheckParameters `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

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

type Http2HealthCheckInitParameters struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type Http2HealthCheckObservation struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type Http2HealthCheckParameters struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type SSLHealthCheckInitParameters struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type SSLHealthCheckObservation struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type SSLHealthCheckParameters struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type TCPHealthCheckInitParameters struct {

	// The TCP port number for the TCP health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type TCPHealthCheckObservation struct {

	// The TCP port number for the TCP health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type TCPHealthCheckParameters struct {

	// The TCP port number for the TCP health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

// HealthCheckSpec defines the desired state of HealthCheck
type HealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthCheckParameters `json:"forProvider"`
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
	InitProvider HealthCheckInitParameters `json:"initProvider,omitempty"`
}

// HealthCheckStatus defines the observed state of HealthCheck.
type HealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HealthCheck is the Schema for the HealthChecks API. Health Checks determine whether instances are responsive and able to do work.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthCheckSpec   `json:"spec"`
	Status            HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheckList contains a list of HealthChecks
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheck `json:"items"`
}

// Repository type metadata.
var (
	HealthCheck_Kind             = "HealthCheck"
	HealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthCheck_Kind}.String()
	HealthCheck_KindAPIVersion   = HealthCheck_Kind + "." + CRDGroupVersion.String()
	HealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(HealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}
