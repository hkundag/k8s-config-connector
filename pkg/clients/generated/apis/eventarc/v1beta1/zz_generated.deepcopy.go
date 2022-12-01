//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventarcTrigger) DeepCopyInto(out *EventarcTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventarcTrigger.
func (in *EventarcTrigger) DeepCopy() *EventarcTrigger {
	if in == nil {
		return nil
	}
	out := new(EventarcTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventarcTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventarcTriggerList) DeepCopyInto(out *EventarcTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventarcTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventarcTriggerList.
func (in *EventarcTriggerList) DeepCopy() *EventarcTriggerList {
	if in == nil {
		return nil
	}
	out := new(EventarcTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventarcTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventarcTriggerSpec) DeepCopyInto(out *EventarcTriggerSpec) {
	*out = *in
	if in.ChannelRef != nil {
		in, out := &in.ChannelRef, &out.ChannelRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	in.Destination.DeepCopyInto(&out.Destination)
	if in.MatchingCriteria != nil {
		in, out := &in.MatchingCriteria, &out.MatchingCriteria
		*out = make([]TriggerMatchingCriteria, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(TriggerTransport)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventarcTriggerSpec.
func (in *EventarcTriggerSpec) DeepCopy() *EventarcTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(EventarcTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventarcTriggerStatus) DeepCopyInto(out *EventarcTriggerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.ResourceConditions != nil {
		in, out := &in.ResourceConditions, &out.ResourceConditions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(TriggerTransportStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventarcTriggerStatus.
func (in *EventarcTriggerStatus) DeepCopy() *EventarcTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(EventarcTriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerCloudRunService) DeepCopyInto(out *TriggerCloudRunService) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	out.ServiceRef = in.ServiceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerCloudRunService.
func (in *TriggerCloudRunService) DeepCopy() *TriggerCloudRunService {
	if in == nil {
		return nil
	}
	out := new(TriggerCloudRunService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerDestination) DeepCopyInto(out *TriggerDestination) {
	*out = *in
	if in.CloudFunctionRef != nil {
		in, out := &in.CloudFunctionRef, &out.CloudFunctionRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.CloudRunService != nil {
		in, out := &in.CloudRunService, &out.CloudRunService
		*out = new(TriggerCloudRunService)
		(*in).DeepCopyInto(*out)
	}
	if in.Gke != nil {
		in, out := &in.Gke, &out.Gke
		*out = new(TriggerGke)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkflowRef != nil {
		in, out := &in.WorkflowRef, &out.WorkflowRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerDestination.
func (in *TriggerDestination) DeepCopy() *TriggerDestination {
	if in == nil {
		return nil
	}
	out := new(TriggerDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerGke) DeepCopyInto(out *TriggerGke) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerGke.
func (in *TriggerGke) DeepCopy() *TriggerGke {
	if in == nil {
		return nil
	}
	out := new(TriggerGke)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerMatchingCriteria) DeepCopyInto(out *TriggerMatchingCriteria) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerMatchingCriteria.
func (in *TriggerMatchingCriteria) DeepCopy() *TriggerMatchingCriteria {
	if in == nil {
		return nil
	}
	out := new(TriggerMatchingCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerPubsub) DeepCopyInto(out *TriggerPubsub) {
	*out = *in
	if in.TopicRef != nil {
		in, out := &in.TopicRef, &out.TopicRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerPubsub.
func (in *TriggerPubsub) DeepCopy() *TriggerPubsub {
	if in == nil {
		return nil
	}
	out := new(TriggerPubsub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerPubsubStatus) DeepCopyInto(out *TriggerPubsubStatus) {
	*out = *in
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerPubsubStatus.
func (in *TriggerPubsubStatus) DeepCopy() *TriggerPubsubStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerPubsubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTransport) DeepCopyInto(out *TriggerTransport) {
	*out = *in
	if in.Pubsub != nil {
		in, out := &in.Pubsub, &out.Pubsub
		*out = new(TriggerPubsub)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTransport.
func (in *TriggerTransport) DeepCopy() *TriggerTransport {
	if in == nil {
		return nil
	}
	out := new(TriggerTransport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTransportStatus) DeepCopyInto(out *TriggerTransportStatus) {
	*out = *in
	if in.Pubsub != nil {
		in, out := &in.Pubsub, &out.Pubsub
		*out = new(TriggerPubsubStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTransportStatus.
func (in *TriggerTransportStatus) DeepCopy() *TriggerTransportStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerTransportStatus)
	in.DeepCopyInto(out)
	return out
}
