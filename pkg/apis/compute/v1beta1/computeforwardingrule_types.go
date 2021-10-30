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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ForwardingruleFilterLabels struct {
	/* Immutable. Name of metadata label.

	The name can have a maximum length of 1024 characters and must be at least 1 character long. */
	Name string `json:"name"`

	/* Immutable. The value of the label must match the specified value.

	value can have a maximum length of 1024 characters. */
	Value string `json:"value"`
}

type ForwardingruleIpAddress struct {
	/*  */
	// +optional
	AddressRef *v1alpha1.ResourceRef `json:"addressRef,omitempty"`

	/*  */
	// +optional
	Ip *string `json:"ip,omitempty"`
}

type ForwardingruleMetadataFilters struct {
	/* Immutable. The list of label value pairs that must match labels in the provided metadata based on `filterMatchCriteria`

	This list must not be empty and can have at the most 64 entries. */
	FilterLabels []ForwardingruleFilterLabels `json:"filterLabels"`

	/* Immutable. Specifies how individual `filterLabel` matches within the list of `filterLabel`s contribute towards the overall `metadataFilter` match.

	Supported values are:

	*   MATCH_ANY: At least one of the `filterLabels` must have a matching label in the provided metadata.
	*   MATCH_ALL: All `filterLabels` must have matching labels in the provided metadata. Possible values: NOT_SET, MATCH_ALL, MATCH_ANY */
	FilterMatchCriteria string `json:"filterMatchCriteria"`
}

type ForwardingruleTarget struct {
	/*  */
	// +optional
	TargetGRPCProxyRef *v1alpha1.ResourceRef `json:"targetGRPCProxyRef,omitempty"`

	/*  */
	// +optional
	TargetHTTPProxyRef *v1alpha1.ResourceRef `json:"targetHTTPProxyRef,omitempty"`

	/*  */
	// +optional
	TargetHTTPSProxyRef *v1alpha1.ResourceRef `json:"targetHTTPSProxyRef,omitempty"`

	/*  */
	// +optional
	TargetSSLProxyRef *v1alpha1.ResourceRef `json:"targetSSLProxyRef,omitempty"`

	/*  */
	// +optional
	TargetTCPProxyRef *v1alpha1.ResourceRef `json:"targetTCPProxyRef,omitempty"`

	/*  */
	// +optional
	TargetVPNGatewayRef *v1alpha1.ResourceRef `json:"targetVPNGatewayRef,omitempty"`
}

type ComputeForwardingRuleSpec struct {
	/* Immutable. This field is used along with the `backend_service` field for internal load balancing or with the `target` field for internal TargetInstance. This field cannot be used with `port` or `portRange` fields. When the load balancing scheme is `INTERNAL` and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule. */
	// +optional
	AllPorts *bool `json:"allPorts,omitempty"`

	/* This field is used along with the `backend_service` field for internal load balancing or with the `target` field for internal TargetInstance. If the field is set to `TRUE`, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer. */
	// +optional
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty"`

	/* A ComputeBackendService to receive the matched traffic. This is
	used only for internal load balancing. */
	// +optional
	BackendServiceRef *v1alpha1.ResourceRef `json:"backendServiceRef,omitempty"`

	/* Immutable. An optional description of this resource. Provide this property when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The IP address that this forwarding rule is serving on behalf of.

	Addresses are restricted based on the forwarding rule's load
	balancing scheme (EXTERNAL or INTERNAL) and scope (global or
	regional).

	When the load balancing scheme is EXTERNAL, for global forwarding
	rules, the address must be a global IP, and for regional forwarding
	rules, the address must live in the same region as the forwarding
	rule. If this field is empty, an ephemeral IPv4 address from the
	same scope (global or regional) will be assigned. A regional
	forwarding rule supports IPv4 only. A global forwarding rule
	supports either IPv4 or IPv6.

	When the load balancing scheme is INTERNAL, this can only be an RFC
	1918 IP address belonging to the network/subnet configured for the
	forwarding rule. By default, if this field is empty, an ephemeral
	internal IP address will be automatically allocated from the IP
	range of the subnet or network configured for this forwarding rule. */
	// +optional
	IpAddress *ForwardingruleIpAddress `json:"ipAddress,omitempty"`

	/* Immutable. The IP protocol to which this rule applies. For protocol forwarding, valid options are `TCP`, `UDP`, `ESP`, `AH`, `SCTP` or `ICMP`. For Internal TCP/UDP Load Balancing, the load balancing scheme is `INTERNAL`, and one of `TCP` or `UDP` are valid. For Traffic Director, the load balancing scheme is `INTERNAL_SELF_MANAGED`, and only `TCP`is valid. For Internal HTTP(S) Load Balancing, the load balancing scheme is `INTERNAL_MANAGED`, and only `TCP` is valid. For HTTP(S), SSL Proxy, and TCP Proxy Load Balancing, the load balancing scheme is `EXTERNAL` and only `TCP` is valid. For Network TCP/UDP Load Balancing, the load balancing scheme is `EXTERNAL`, and one of `TCP` or `UDP` is valid. */
	// +optional
	IpProtocol *string `json:"ipProtocol,omitempty"`

	/* Immutable. The IP Version that will be used by this forwarding rule. Valid options are `IPV4` or `IPV6`. This can only be specified for an external global forwarding rule. Possible values: UNSPECIFIED_VERSION, IPV4, IPV6 */
	// +optional
	IpVersion *string `json:"ipVersion,omitempty"`

	/* Immutable. Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a `PacketMirroring` rule applies to them. This can only be set to true for load balancers that have their `loadBalancingScheme` set to `INTERNAL`. */
	// +optional
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty"`

	/* Immutable. Specifies the forwarding rule type.

	*   `EXTERNAL` is used for:
	    *   Classic Cloud VPN gateways
	    *   Protocol forwarding to VMs from an external IP address
	    *   The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP
	*   `INTERNAL` is used for:
	    *   Protocol forwarding to VMs from an internal IP address
	    *   Internal TCP/UDP load balancers
	*   `INTERNAL_MANAGED` is used for:
	    *   Internal HTTP(S) load balancers
	*   `INTERNAL_SELF_MANAGED` is used for:
	    *   Traffic Director

	For more information about forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values: INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL */
	// +optional
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	/* Location represents the geographical location of the ComputeForwardingRule. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of [xDS](https://github.com/envoyproxy/data-plane-api/blob/master/XDS_PROTOCOL.md) compliant clients. In their xDS requests to Loadbalancer, xDS clients present [node metadata](https://github.com/envoyproxy/data-plane-api/search?q=%22message+Node%22+in%3A%2Fenvoy%2Fapi%2Fv2%2Fcore%2Fbase.proto&). If a match takes place, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. `TargetHttpProxy`, `UrlMap`) referenced by the `ForwardingRule` will not be visible to those proxies.

	For each `metadataFilter` in this list, if its `filterMatchCriteria` is set to MATCH_ANY, at least one of the `filterLabel`s must match the corresponding label provided in the metadata. If its `filterMatchCriteria` is set to MATCH_ALL, then all of its `filterLabel`s must match with corresponding labels provided in the metadata.

	`metadataFilters` specified here will be applifed before those specified in the `UrlMap` that this `ForwardingRule` references.

	`metadataFilters` only applies to Loadbalancers that have their loadBalancingScheme set to `INTERNAL_SELF_MANAGED`. */
	// +optional
	MetadataFilters []ForwardingruleMetadataFilters `json:"metadataFilters,omitempty"`

	/* This field is not used for external load balancing. For internal
	load balancing, this field identifies the network that the load
	balanced IP should belong to for this forwarding rule. If this
	field is not specified, the default network will be used. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. This signifies the networking tier used for configuring this load balancer and can only take the following values: `PREMIUM`, `STANDARD`. For regional ForwardingRule, the valid values are `PREMIUM` and `STANDARD`. For GlobalForwardingRule, the valid value is `PREMIUM`. If this field is not specified, it is assumed to be `PREMIUM`. If `IPAddress` is specified, this value must be equal to the networkTier of the Address. */
	// +optional
	NetworkTier *string `json:"networkTier,omitempty"`

	/* Immutable. When the load balancing scheme is `EXTERNAL`, `INTERNAL_SELF_MANAGED` and `INTERNAL_MANAGED`, you can specify a `port_range`. Use with a forwarding rule that points to a target proxy or a target pool. Do not use with a forwarding rule that points to a backend service. This field is used along with the `target` field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when `IPProtocol` is `TCP`, `UDP`, or `SCTP`, only packets addressed to ports in the specified range will be forwarded to `target`. Forwarding rules with the same `[IPAddress, IPProtocol]` pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:

	*   TargetHttpProxy: 80, 8080
	*   TargetHttpsProxy: 443
	*   TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	*   TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	*   TargetVpnGateway: 500, 4500

	@pattern: d+(?:-d+)? */
	// +optional
	PortRange *string `json:"portRange,omitempty"`

	/* Immutable. This field is used along with the `backend_service` field for internal load balancing. When the load balancing scheme is `INTERNAL`, a list of ports can be configured, for example, ['80'], ['8000','9000']. Only packets addressed to these ports are forwarded to the backends configured with the forwarding rule. If the forwarding rule's loadBalancingScheme is INTERNAL, you can specify ports in one of the following ways: * A list of up to five ports, which can be non-contiguous * Keyword `ALL`, which causes the forwarding rule to forward traffic on any port of the forwarding rule's protocol. @pattern: d+(?:-d+)? For more information, refer to [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications). */
	// +optional
	Ports []string `json:"ports,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing. */
	// +optional
	ServiceLabel *string `json:"serviceLabel,omitempty"`

	/* The subnetwork that the load balanced IP should belong to for this
	forwarding rule. This field is only used for internal load
	balancing.

	If the network specified is in auto subnet mode, this field is
	optional. However, if the network is in custom subnet mode, a
	subnetwork must be specified. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`

	/* The target resource to receive the matched traffic. The forwarded
	traffic must be of a type appropriate to the target object. For
	INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	are valid. */
	// +optional
	Target *ForwardingruleTarget `json:"target,omitempty"`
}

type ComputeForwardingRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeForwardingRule's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* [Output Only] Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* Used internally during label updates. */
	LabelFingerprint string `json:"labelFingerprint,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* [Output Only] Server-defined URL for the resource. */
	SelfLink string `json:"selfLink,omitempty"`
	/* [Output Only] The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing. */
	ServiceName string `json:"serviceName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeForwardingRule is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeForwardingRuleSpec   `json:"spec,omitempty"`
	Status ComputeForwardingRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeForwardingRuleList contains a list of ComputeForwardingRule
type ComputeForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeForwardingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeForwardingRule{}, &ComputeForwardingRuleList{})
}