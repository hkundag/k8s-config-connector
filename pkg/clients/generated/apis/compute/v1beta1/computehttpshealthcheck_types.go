// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeHTTPSHealthCheckSpec struct {
	/* How often (in seconds) to send a health check. The default value is 5
	seconds. */
	// +optional
	CheckIntervalSec *int `json:"checkIntervalSec,omitempty"`

	/* An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* A so-far unhealthy instance will be marked healthy after this many
	consecutive successes. The default value is 2. */
	// +optional
	HealthyThreshold *int `json:"healthyThreshold,omitempty"`

	/* The value of the host header in the HTTPS health check request. If
	left empty (default value), the public IP on behalf of which this
	health check is performed will be used. */
	// +optional
	Host *string `json:"host,omitempty"`

	/* The TCP port number for the HTTPS health check request.
	The default value is 443. */
	// +optional
	Port *int `json:"port,omitempty"`

	/* The request path of the HTTPS health check request.
	The default value is /. */
	// +optional
	RequestPath *string `json:"requestPath,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* How long (in seconds) to wait before claiming failure.
	The default value is 5 seconds.  It is invalid for timeoutSec to have
	greater value than checkIntervalSec. */
	// +optional
	TimeoutSec *int `json:"timeoutSec,omitempty"`

	/* A so-far healthy instance will be marked unhealthy after this many
	consecutive failures. The default value is 2. */
	// +optional
	UnhealthyThreshold *int `json:"unhealthyThreshold,omitempty"`
}

type ComputeHTTPSHealthCheckStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeHTTPSHealthCheck's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeHTTPSHealthCheck is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeHTTPSHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeHTTPSHealthCheckSpec   `json:"spec,omitempty"`
	Status ComputeHTTPSHealthCheckStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeHTTPSHealthCheckList contains a list of ComputeHTTPSHealthCheck
type ComputeHTTPSHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeHTTPSHealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeHTTPSHealthCheck{}, &ComputeHTTPSHealthCheckList{})
}
