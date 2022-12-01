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

type StorageBucketAccessControlSpec struct {
	/* Reference to the bucket. */
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* Immutable. The entity holding the permission, in one of the following forms:
	user-userId
	user-email
	group-groupId
	group-email
	domain-domain
	project-team-projectId
	allUsers
	allAuthenticatedUsers
	Examples:
	The user liz@example.com would be user-liz@example.com.
	The group example@googlegroups.com would be
	group-example@googlegroups.com.
	To refer to all members of the Google Apps for Business domain
	example.com, the entity would be domain-example.com. */
	Entity string `json:"entity"`

	/* The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]. */
	// +optional
	Role *string `json:"role,omitempty"`
}

type StorageBucketAccessControlStatus struct {
	/* Conditions represent the latest available observations of the
	   StorageBucketAccessControl's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The domain associated with the entity. */
	// +optional
	Domain *string `json:"domain,omitempty"`

	/* The email address associated with the entity. */
	// +optional
	Email *string `json:"email,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageBucketAccessControl is the Schema for the storage API
// +k8s:openapi-gen=true
type StorageBucketAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageBucketAccessControlSpec   `json:"spec,omitempty"`
	Status StorageBucketAccessControlStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageBucketAccessControlList contains a list of StorageBucketAccessControl
type StorageBucketAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageBucketAccessControl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageBucketAccessControl{}, &StorageBucketAccessControlList{})
}
