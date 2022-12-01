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
func (in *BillingBudgetsBudget) DeepCopyInto(out *BillingBudgetsBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetsBudget.
func (in *BillingBudgetsBudget) DeepCopy() *BillingBudgetsBudget {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetsBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingBudgetsBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetsBudgetList) DeepCopyInto(out *BillingBudgetsBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BillingBudgetsBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetsBudgetList.
func (in *BillingBudgetsBudgetList) DeepCopy() *BillingBudgetsBudgetList {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetsBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingBudgetsBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetsBudgetSpec) DeepCopyInto(out *BillingBudgetsBudgetSpec) {
	*out = *in
	if in.AllUpdatesRule != nil {
		in, out := &in.AllUpdatesRule, &out.AllUpdatesRule
		*out = new(BudgetAllUpdatesRule)
		(*in).DeepCopyInto(*out)
	}
	in.Amount.DeepCopyInto(&out.Amount)
	out.BillingAccountRef = in.BillingAccountRef
	if in.BudgetFilter != nil {
		in, out := &in.BudgetFilter, &out.BudgetFilter
		*out = new(BudgetBudgetFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ThresholdRules != nil {
		in, out := &in.ThresholdRules, &out.ThresholdRules
		*out = make([]BudgetThresholdRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetsBudgetSpec.
func (in *BillingBudgetsBudgetSpec) DeepCopy() *BillingBudgetsBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetsBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingBudgetsBudgetStatus) DeepCopyInto(out *BillingBudgetsBudgetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingBudgetsBudgetStatus.
func (in *BillingBudgetsBudgetStatus) DeepCopy() *BillingBudgetsBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(BillingBudgetsBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetAllUpdatesRule) DeepCopyInto(out *BudgetAllUpdatesRule) {
	*out = *in
	if in.DisableDefaultIamRecipients != nil {
		in, out := &in.DisableDefaultIamRecipients, &out.DisableDefaultIamRecipients
		*out = new(bool)
		**out = **in
	}
	if in.MonitoringNotificationChannels != nil {
		in, out := &in.MonitoringNotificationChannels, &out.MonitoringNotificationChannels
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.PubsubTopicRef != nil {
		in, out := &in.PubsubTopicRef, &out.PubsubTopicRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.SchemaVersion != nil {
		in, out := &in.SchemaVersion, &out.SchemaVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetAllUpdatesRule.
func (in *BudgetAllUpdatesRule) DeepCopy() *BudgetAllUpdatesRule {
	if in == nil {
		return nil
	}
	out := new(BudgetAllUpdatesRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetAmount) DeepCopyInto(out *BudgetAmount) {
	*out = *in
	if in.LastPeriodAmount != nil {
		in, out := &in.LastPeriodAmount, &out.LastPeriodAmount
		*out = new(BudgetLastPeriodAmount)
		**out = **in
	}
	if in.SpecifiedAmount != nil {
		in, out := &in.SpecifiedAmount, &out.SpecifiedAmount
		*out = new(BudgetSpecifiedAmount)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetAmount.
func (in *BudgetAmount) DeepCopy() *BudgetAmount {
	if in == nil {
		return nil
	}
	out := new(BudgetAmount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetBudgetFilter) DeepCopyInto(out *BudgetBudgetFilter) {
	*out = *in
	if in.CalendarPeriod != nil {
		in, out := &in.CalendarPeriod, &out.CalendarPeriod
		*out = new(string)
		**out = **in
	}
	if in.CreditTypes != nil {
		in, out := &in.CreditTypes, &out.CreditTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreditTypesTreatment != nil {
		in, out := &in.CreditTypesTreatment, &out.CreditTypesTreatment
		*out = new(string)
		**out = **in
	}
	if in.CustomPeriod != nil {
		in, out := &in.CustomPeriod, &out.CustomPeriod
		*out = new(BudgetCustomPeriod)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]BudgetLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subaccounts != nil {
		in, out := &in.Subaccounts, &out.Subaccounts
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetBudgetFilter.
func (in *BudgetBudgetFilter) DeepCopy() *BudgetBudgetFilter {
	if in == nil {
		return nil
	}
	out := new(BudgetBudgetFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetCustomPeriod) DeepCopyInto(out *BudgetCustomPeriod) {
	*out = *in
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(BudgetEndDate)
		(*in).DeepCopyInto(*out)
	}
	in.StartDate.DeepCopyInto(&out.StartDate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetCustomPeriod.
func (in *BudgetCustomPeriod) DeepCopy() *BudgetCustomPeriod {
	if in == nil {
		return nil
	}
	out := new(BudgetCustomPeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetEndDate) DeepCopyInto(out *BudgetEndDate) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(int)
		**out = **in
	}
	if in.Month != nil {
		in, out := &in.Month, &out.Month
		*out = new(int)
		**out = **in
	}
	if in.Year != nil {
		in, out := &in.Year, &out.Year
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetEndDate.
func (in *BudgetEndDate) DeepCopy() *BudgetEndDate {
	if in == nil {
		return nil
	}
	out := new(BudgetEndDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetLabels) DeepCopyInto(out *BudgetLabels) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetLabels.
func (in *BudgetLabels) DeepCopy() *BudgetLabels {
	if in == nil {
		return nil
	}
	out := new(BudgetLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetLastPeriodAmount) DeepCopyInto(out *BudgetLastPeriodAmount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetLastPeriodAmount.
func (in *BudgetLastPeriodAmount) DeepCopy() *BudgetLastPeriodAmount {
	if in == nil {
		return nil
	}
	out := new(BudgetLastPeriodAmount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpecifiedAmount) DeepCopyInto(out *BudgetSpecifiedAmount) {
	*out = *in
	if in.CurrencyCode != nil {
		in, out := &in.CurrencyCode, &out.CurrencyCode
		*out = new(string)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(int)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpecifiedAmount.
func (in *BudgetSpecifiedAmount) DeepCopy() *BudgetSpecifiedAmount {
	if in == nil {
		return nil
	}
	out := new(BudgetSpecifiedAmount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetStartDate) DeepCopyInto(out *BudgetStartDate) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(int)
		**out = **in
	}
	if in.Month != nil {
		in, out := &in.Month, &out.Month
		*out = new(int)
		**out = **in
	}
	if in.Year != nil {
		in, out := &in.Year, &out.Year
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetStartDate.
func (in *BudgetStartDate) DeepCopy() *BudgetStartDate {
	if in == nil {
		return nil
	}
	out := new(BudgetStartDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetThresholdRules) DeepCopyInto(out *BudgetThresholdRules) {
	*out = *in
	if in.SpendBasis != nil {
		in, out := &in.SpendBasis, &out.SpendBasis
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetThresholdRules.
func (in *BudgetThresholdRules) DeepCopy() *BudgetThresholdRules {
	if in == nil {
		return nil
	}
	out := new(BudgetThresholdRules)
	in.DeepCopyInto(out)
	return out
}
