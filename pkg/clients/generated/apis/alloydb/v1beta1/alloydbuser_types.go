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

type UserValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type AlloyDBUserSpec struct {
	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* List of database roles this database user has. */
	// +optional
	DatabaseRoles []string `json:"databaseRoles,omitempty"`

	/* Password for this database user. */
	// +optional
	Password *UserPassword `json:"password,omitempty"`

	/* Immutable. Optional. The userId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]. */
	UserType string `json:"userType"`
}

type AlloyDBUserStatus struct {
	/* Conditions represent the latest available observations of the
	   AlloyDBUser's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlloyDBUser is the Schema for the alloydb API
// +k8s:openapi-gen=true
type AlloyDBUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBUserSpec   `json:"spec,omitempty"`
	Status AlloyDBUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlloyDBUserList contains a list of AlloyDBUser
type AlloyDBUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlloyDBUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlloyDBUser{}, &AlloyDBUserList{})
}
