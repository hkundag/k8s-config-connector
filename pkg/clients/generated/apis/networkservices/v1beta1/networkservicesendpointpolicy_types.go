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

type EndpointpolicyEndpointMatcher struct {
	/* The matcher is based on node metadata presented by xDS clients. */
	// +optional
	MetadataLabelMatcher *EndpointpolicyMetadataLabelMatcher `json:"metadataLabelMatcher,omitempty"`
}

type EndpointpolicyMetadataLabelMatcher struct {
	/* Specifies how matching should be done. Supported values are: MATCH_ANY: At least one of the Labels specified in the matcher should match the metadata presented by xDS client. MATCH_ALL: The metadata presented by the xDS client should contain all of the labels specified here. The selection is determined based on the best match. For example, suppose there are three EndpointPolicy resources P1, P2 and P3 and if P1 has a the matcher as MATCH_ANY , P2 has MATCH_ALL , and P3 has MATCH_ALL . If a client with label connects, the config from P1 will be selected. If a client with label connects, the config from P2 will be selected. If a client with label connects, the config from P3 will be selected. If there is more than one best match, (for example, if a config P4 with selector exists and if a client with label connects), an error will be thrown. Possible values: METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED, MATCH_ANY, MATCH_ALL */
	// +optional
	MetadataLabelMatchCriteria *string `json:"metadataLabelMatchCriteria,omitempty"`

	/* The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria This list can have at most 64 entries. The list can be empty if the match criteria is MATCH_ANY, to specify a wildcard match (i.e this matches any client). */
	// +optional
	MetadataLabels []EndpointpolicyMetadataLabels `json:"metadataLabels,omitempty"`
}

type EndpointpolicyMetadataLabels struct {
	/* Required. Label name presented as key in xDS Node Metadata. */
	LabelName string `json:"labelName"`

	/* Required. Label value presented as value corresponding to the above key, in xDS Node Metadata. */
	LabelValue string `json:"labelValue"`
}

type EndpointpolicyTrafficPortSelector struct {
	/* Optional. A list of ports. Can be port numbers or port range (example, specifies all ports from 80 to 90, including 80 and 90) or named ports or * to specify all ports. If the list is empty, all ports are selected. */
	// +optional
	Ports []string `json:"ports,omitempty"`
}

type NetworkServicesEndpointPolicySpec struct {
	// +optional
	AuthorizationPolicyRef *v1alpha1.ResourceRef `json:"authorizationPolicyRef,omitempty"`

	// +optional
	ClientTlsPolicyRef *v1alpha1.ResourceRef `json:"clientTlsPolicyRef,omitempty"`

	/* Optional. A free-text description of the resource. Max length 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Required. A matcher that selects endpoints to which the policies should be applied. */
	EndpointMatcher EndpointpolicyEndpointMatcher `json:"endpointMatcher"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	ServerTlsPolicyRef *v1alpha1.ResourceRef `json:"serverTlsPolicyRef,omitempty"`

	/* Optional. Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports. */
	// +optional
	TrafficPortSelector *EndpointpolicyTrafficPortSelector `json:"trafficPortSelector,omitempty"`

	/* Required. The type of endpoint config. This is primarily used to validate the configuration. Possible values: ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED, SIDECAR_PROXY, GRPC_SERVER */
	Type string `json:"type"`
}

type NetworkServicesEndpointPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkServicesEndpointPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The timestamp when the resource was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. The timestamp when the resource was updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesEndpointPolicy is the Schema for the networkservices API
// +k8s:openapi-gen=true
type NetworkServicesEndpointPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesEndpointPolicySpec   `json:"spec,omitempty"`
	Status NetworkServicesEndpointPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesEndpointPolicyList contains a list of NetworkServicesEndpointPolicy
type NetworkServicesEndpointPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesEndpointPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesEndpointPolicy{}, &NetworkServicesEndpointPolicyList{})
}
