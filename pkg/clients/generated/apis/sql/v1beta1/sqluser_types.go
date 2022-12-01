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

type UserPassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *UserValueFrom `json:"valueFrom,omitempty"`
}

type UserPasswordPolicy struct {
	/* Number of failed attempts allowed before the user get locked. */
	// +optional
	AllowedFailedAttempts *int `json:"allowedFailedAttempts,omitempty"`

	/* If true, the check that will lock user after too many failed login attempts will be enabled. */
	// +optional
	EnableFailedAttemptsCheck *bool `json:"enableFailedAttemptsCheck,omitempty"`

	/* If true, the user must specify the current password before changing the password. This flag is supported only for MySQL. */
	// +optional
	EnablePasswordVerification *bool `json:"enablePasswordVerification,omitempty"`

	/* Password expiration duration with one week grace period. */
	// +optional
	PasswordExpirationDuration *string `json:"passwordExpirationDuration,omitempty"`

	// +optional
	Status []UserStatus `json:"status,omitempty"`
}

type UserStatus struct {
	/* If true, user does not have login privileges. */
	// +optional
	Locked *bool `json:"locked,omitempty"`

	/* Password expiration duration with one week grace period. */
	// +optional
	PasswordExpirationTime *string `json:"passwordExpirationTime,omitempty"`
}

type UserValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type SQLUserSpec struct {
	/* Immutable. The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL instances. Can be an IP address. Changing this forces a new resource to be created. */
	// +optional
	Host *string `json:"host,omitempty"`

	InstanceRef v1alpha1.ResourceRef `json:"instanceRef"`

	/* The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to
	either CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT. */
	// +optional
	Password *UserPassword `json:"password,omitempty"`

	// +optional
	PasswordPolicy *UserPasswordPolicy `json:"passwordPolicy,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The user type. It determines the method to authenticate the user during login.
	The default is the database's built-in user type. Flags include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT". */
	// +optional
	Type *string `json:"type,omitempty"`
}

type UserSqlServerUserDetailsStatus struct {
	/* If the user has been disabled. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* The server roles for this user in the database. */
	// +optional
	ServerRoles []string `json:"serverRoles,omitempty"`
}

type SQLUserStatus struct {
	/* Conditions represent the latest available observations of the
	   SQLUser's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SqlServerUserDetails []UserSqlServerUserDetailsStatus `json:"sqlServerUserDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SQLUser is the Schema for the sql API
// +k8s:openapi-gen=true
type SQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLUserSpec   `json:"spec,omitempty"`
	Status SQLUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SQLUserList contains a list of SQLUser
type SQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SQLUser{}, &SQLUserList{})
}
