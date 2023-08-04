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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AppconnectionApplicationEndpoint struct {
	/* Hostname or IP address of the remote application endpoint. */
	Host string `json:"host"`

	/* Port of the remote application endpoint. */
	Port int `json:"port"`
}

type AppconnectionGateway struct {
	/* AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}. */
	AppGateway string `json:"appGateway"`

	/* Ingress port reserved on the gateways for this AppConnection, if not specified or zero, the default port is 19443. */
	// +optional
	IngressPort *int `json:"ingressPort,omitempty"`

	/* The type of hosting used by the gateway. Refer to
	https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	for a list of possible values. */
	// +optional
	Type *string `json:"type,omitempty"`

	/* Server-defined URI for this resource. */
	// +optional
	Uri *string `json:"uri,omitempty"`
}

type BeyondCorpAppConnectionSpec struct {
	/* Address of the remote application endpoint for the BeyondCorp AppConnection. */
	ApplicationEndpoint AppconnectionApplicationEndpoint `json:"applicationEndpoint"`

	/* List of AppConnectors that are authorised to be associated with this AppConnection. */
	// +optional
	Connectors []string `json:"connectors,omitempty"`

	/* An arbitrary user-provided name for the AppConnection. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Gateway used by the AppConnection. */
	// +optional
	Gateway *AppconnectionGateway `json:"gateway,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The region of the AppConnection. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The type of network connectivity used by the AppConnection. Refer to
	https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
	for a list of possible values. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type BeyondCorpAppConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   BeyondCorpAppConnection's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BeyondCorpAppConnection is the Schema for the beyondcorp API
// +k8s:openapi-gen=true
type BeyondCorpAppConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeyondCorpAppConnectionSpec   `json:"spec,omitempty"`
	Status BeyondCorpAppConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BeyondCorpAppConnectionList contains a list of BeyondCorpAppConnection
type BeyondCorpAppConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BeyondCorpAppConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BeyondCorpAppConnection{}, &BeyondCorpAppConnectionList{})
}