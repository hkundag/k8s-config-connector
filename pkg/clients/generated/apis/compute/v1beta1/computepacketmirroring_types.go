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

type PacketmirroringCollectorIlb struct {
	UrlRef v1alpha1.ResourceRef `json:"urlRef"`
}

type PacketmirroringFilter struct {
	/* IP CIDR ranges that apply as filter on the source (ingress) or destination (egress) IP in the IP header. Only IPv4 is supported. If no ranges are specified, all traffic that matches the specified IPProtocols is mirrored. If neither cidrRanges nor IPProtocols is specified, all traffic is mirrored. */
	// +optional
	CidrRanges []string `json:"cidrRanges,omitempty"`

	/* Direction of traffic to mirror, either INGRESS, EGRESS, or BOTH. The default is BOTH. */
	// +optional
	Direction *string `json:"direction,omitempty"`

	/* Protocols that apply as filter on mirrored traffic. If no protocols are specified, all traffic that matches the specified CIDR ranges is mirrored. If neither cidrRanges nor IPProtocols is specified, all traffic is mirrored. */
	// +optional
	IpProtocols []string `json:"ipProtocols,omitempty"`
}

type PacketmirroringInstances struct {
	/* Immutable. Output only. Unique identifier for the instance; defined by the server. */
	// +optional
	CanonicalUrl *string `json:"canonicalUrl,omitempty"`

	// +optional
	UrlRef *v1alpha1.ResourceRef `json:"urlRef,omitempty"`
}

type PacketmirroringMirroredResources struct {
	/* A set of virtual machine instances that are being mirrored. They must live in zones contained in the same region as this packetMirroring. Note that this config will apply only to those network interfaces of the Instances that belong to the network specified in this packetMirroring. You may specify a maximum of 50 Instances. */
	// +optional
	Instances []PacketmirroringInstances `json:"instances,omitempty"`

	/* Immutable. A set of subnetworks for which traffic from/to all VM instances will be mirrored. They must live in the same region as this packetMirroring. You may specify a maximum of 5 subnetworks. */
	// +optional
	Subnetworks []PacketmirroringSubnetworks `json:"subnetworks,omitempty"`

	/* A set of mirrored tags. Traffic from/to all VM instances that have one or more of these tags will be mirrored. */
	// +optional
	Tags []string `json:"tags,omitempty"`
}

type PacketmirroringNetwork struct {
	/* Immutable. */
	UrlRef v1alpha1.ResourceRef `json:"urlRef"`
}

type PacketmirroringSubnetworks struct {
	/* Immutable. Output only. Unique identifier for the subnetwork; defined by the server. */
	// +optional
	CanonicalUrl *string `json:"canonicalUrl,omitempty"`

	/* Immutable. */
	// +optional
	UrlRef *v1alpha1.ResourceRef `json:"urlRef,omitempty"`
}

type ComputePacketMirroringSpec struct {
	/* The Forwarding Rule resource of type `loadBalancingScheme=INTERNAL` that will be used as collector for mirrored traffic. The specified forwarding rule must have `isMirroringCollector` set to true. */
	CollectorIlb PacketmirroringCollectorIlb `json:"collectorIlb"`

	/* An optional description of this resource. Provide this property when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE. */
	// +optional
	Enable *string `json:"enable,omitempty"`

	/* Filter for mirrored traffic. If unspecified, all traffic is mirrored. */
	// +optional
	Filter *PacketmirroringFilter `json:"filter,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored. */
	MirroredResources PacketmirroringMirroredResources `json:"mirroredResources"`

	/* Immutable. Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network. */
	Network PacketmirroringNetwork `json:"network"`

	/* The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535. */
	// +optional
	Priority *int `json:"priority,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type PacketmirroringCollectorIlbStatus struct {
	/* Output only. Unique identifier for the forwarding rule; defined by the server. */
	// +optional
	CanonicalUrl *string `json:"canonicalUrl,omitempty"`
}

type PacketmirroringNetworkStatus struct {
	/* Output only. Unique identifier for the network; defined by the server. */
	// +optional
	CanonicalUrl *string `json:"canonicalUrl,omitempty"`
}

type ComputePacketMirroringStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputePacketMirroring's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// +optional
	CollectorIlb *PacketmirroringCollectorIlbStatus `json:"collectorIlb,omitempty"`

	/* Output only. The unique identifier for the resource. This identifier is defined by the server. */
	// +optional
	Id *int `json:"id,omitempty"`

	// +optional
	Network *PacketmirroringNetworkStatus `json:"network,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* URI of the region where the packetMirroring resides. */
	// +optional
	Region *string `json:"region,omitempty"`

	/* Server-defined URL for the resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputePacketMirroring is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputePacketMirroring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputePacketMirroringSpec   `json:"spec,omitempty"`
	Status ComputePacketMirroringStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputePacketMirroringList contains a list of ComputePacketMirroring
type ComputePacketMirroringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputePacketMirroring `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputePacketMirroring{}, &ComputePacketMirroringList{})
}
