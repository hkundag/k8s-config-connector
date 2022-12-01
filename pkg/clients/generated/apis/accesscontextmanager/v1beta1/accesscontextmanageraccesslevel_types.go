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

type AccesslevelBasic struct {
	/* How the conditions list should be combined to determine if a request
	is granted this AccessLevel. If AND is used, each Condition in
	conditions must be satisfied for the AccessLevel to be applied. If
	OR is used, at least one Condition in conditions must be satisfied
	for the AccessLevel to be applied. Default value: "AND" Possible values: ["AND", "OR"]. */
	// +optional
	CombiningFunction *string `json:"combiningFunction,omitempty"`

	/* A set of requirements for the AccessLevel to be granted. */
	Conditions []AccesslevelConditions `json:"conditions"`
}

type AccesslevelConditions struct {
	/* Device specific restrictions, all restrictions must hold for
	the Condition to be true. If not specified, all devices are
	allowed. */
	// +optional
	DevicePolicy *AccesslevelDevicePolicy `json:"devicePolicy,omitempty"`

	/* A list of CIDR block IP subnetwork specification. May be IPv4
	or IPv6.
	Note that for a CIDR IP address block, the specified IP address
	portion must be properly truncated (i.e. all the host bits must
	be zero) or the input is considered malformed. For example,
	"192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	is not. The originating IP of a request must be in one of the
	listed subnets in order for this Condition to be true.
	If empty, all IP addresses are allowed. */
	// +optional
	IpSubnetworks []string `json:"ipSubnetworks,omitempty"`

	// +optional
	Members []AccesslevelMembers `json:"members,omitempty"`

	/* Whether to negate the Condition. If true, the Condition becomes
	a NAND over its non-empty fields, each field must be false for
	the Condition overall to be satisfied. Defaults to false. */
	// +optional
	Negate *bool `json:"negate,omitempty"`

	/* The request must originate from one of the provided
	countries/regions.
	Format: A valid ISO 3166-1 alpha-2 code. */
	// +optional
	Regions []string `json:"regions,omitempty"`

	// +optional
	RequiredAccessLevels []v1alpha1.ResourceRef `json:"requiredAccessLevels,omitempty"`
}

type AccesslevelCustom struct {
	/* Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	This page details the objects and attributes that are used to the build the CEL expressions for
	custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec. */
	Expr AccesslevelExpr `json:"expr"`
}

type AccesslevelDevicePolicy struct {
	/* A list of allowed device management levels.
	An empty list allows all management levels. Possible values: ["MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"]. */
	// +optional
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty"`

	/* A list of allowed encryptions statuses.
	An empty list allows all statuses. Possible values: ["ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"]. */
	// +optional
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty"`

	/* A list of allowed OS versions.
	An empty list allows all types and all versions. */
	// +optional
	OsConstraints []AccesslevelOsConstraints `json:"osConstraints,omitempty"`

	/* Whether the device needs to be approved by the customer admin. */
	// +optional
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty"`

	/* Whether the device needs to be corp owned. */
	// +optional
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty"`

	/* Whether or not screenlock is required for the DevicePolicy
	to be true. Defaults to false. */
	// +optional
	RequireScreenLock *bool `json:"requireScreenLock,omitempty"`
}

type AccesslevelExpr struct {
	/* Description of the expression. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Textual representation of an expression in Common Expression Language syntax. */
	Expression string `json:"expression"`

	/* String indicating the location of the expression for error reporting, e.g. a file name and a position in the file. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Title for the expression, i.e. a short string describing its purpose. */
	// +optional
	Title *string `json:"title,omitempty"`
}

type AccesslevelMembers struct {
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	// +optional
	User *string `json:"user,omitempty"`
}

type AccesslevelOsConstraints struct {
	/* The minimum allowed OS version. If not set, any version
	of this OS satisfies the constraint.
	Format: "major.minor.patch" such as "10.5.301", "9.2.1". */
	// +optional
	MinimumVersion *string `json:"minimumVersion,omitempty"`

	/* The operating system type of the device. Possible values: ["OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"]. */
	OsType string `json:"osType"`

	/* If you specify DESKTOP_CHROME_OS for osType, you can optionally include requireVerifiedChromeOs to require Chrome Verified Access. */
	// +optional
	RequireVerifiedChromeOs *bool `json:"requireVerifiedChromeOs,omitempty"`
}

type AccessContextManagerAccessLevelSpec struct {
	/* The AccessContextManagerAccessPolicy this
	AccessContextManagerAccessLevel lives in. */
	AccessPolicyRef v1alpha1.ResourceRef `json:"accessPolicyRef"`

	/* A set of predefined conditions for the access level and a combining function. */
	// +optional
	Basic *AccesslevelBasic `json:"basic,omitempty"`

	/* Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	See CEL spec at: https://github.com/google/cel-spec. */
	// +optional
	Custom *AccesslevelCustom `json:"custom,omitempty"`

	/* Description of the AccessLevel and its use. Does not affect behavior. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Human readable title. Must be unique within the Policy. */
	Title string `json:"title"`
}

type AccessContextManagerAccessLevelStatus struct {
	/* Conditions represent the latest available observations of the
	   AccessContextManagerAccessLevel's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerAccessLevel is the Schema for the accesscontextmanager API
// +k8s:openapi-gen=true
type AccessContextManagerAccessLevel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerAccessLevelSpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessLevelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerAccessLevelList contains a list of AccessContextManagerAccessLevel
type AccessContextManagerAccessLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerAccessLevel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessLevel{}, &AccessContextManagerAccessLevelList{})
}
