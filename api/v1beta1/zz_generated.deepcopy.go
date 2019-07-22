// +build !ignore_autogenerated

/*
Copyright 2019 FinTech Studios, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// autogenerated by controller-gen object, do not modify manually

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPDeploymentTarget) DeepCopyInto(out *VPDeploymentTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPDeploymentTarget.
func (in *VPDeploymentTarget) DeepCopy() *VPDeploymentTarget {
	if in == nil {
		return nil
	}
	out := new(VPDeploymentTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPDeploymentTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPDeploymentTargetList) DeepCopyInto(out *VPDeploymentTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPDeploymentTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPDeploymentTargetList.
func (in *VPDeploymentTargetList) DeepCopy() *VPDeploymentTargetList {
	if in == nil {
		return nil
	}
	out := new(VPDeploymentTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPDeploymentTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPDeploymentTargetMetadata) DeepCopyInto(out *VPDeploymentTargetMetadata) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
	if in.ModifiedAt != nil {
		in, out := &in.ModifiedAt, &out.ModifiedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPDeploymentTargetMetadata.
func (in *VPDeploymentTargetMetadata) DeepCopy() *VPDeploymentTargetMetadata {
	if in == nil {
		return nil
	}
	out := new(VPDeploymentTargetMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPDeploymentTargetObjectSpec) DeepCopyInto(out *VPDeploymentTargetObjectSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPDeploymentTargetObjectSpec.
func (in *VPDeploymentTargetObjectSpec) DeepCopy() *VPDeploymentTargetObjectSpec {
	if in == nil {
		return nil
	}
	out := new(VPDeploymentTargetObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPDeploymentTargetSpec) DeepCopyInto(out *VPDeploymentTargetSpec) {
	*out = *in
	out.Kubernetes = in.Kubernetes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPDeploymentTargetSpec.
func (in *VPDeploymentTargetSpec) DeepCopy() *VPDeploymentTargetSpec {
	if in == nil {
		return nil
	}
	out := new(VPDeploymentTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPKubernetesTarget) DeepCopyInto(out *VPKubernetesTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPKubernetesTarget.
func (in *VPKubernetesTarget) DeepCopy() *VPKubernetesTarget {
	if in == nil {
		return nil
	}
	out := new(VPKubernetesTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPNamespace) DeepCopyInto(out *VPNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPNamespace.
func (in *VPNamespace) DeepCopy() *VPNamespace {
	if in == nil {
		return nil
	}
	out := new(VPNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPNamespaceList) DeepCopyInto(out *VPNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPNamespaceList.
func (in *VPNamespaceList) DeepCopy() *VPNamespaceList {
	if in == nil {
		return nil
	}
	out := new(VPNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPNamespaceMetadata) DeepCopyInto(out *VPNamespaceMetadata) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
	if in.ModifiedAt != nil {
		in, out := &in.ModifiedAt, &out.ModifiedAt
		*out = new(v1.Timestamp)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPNamespaceMetadata.
func (in *VPNamespaceMetadata) DeepCopy() *VPNamespaceMetadata {
	if in == nil {
		return nil
	}
	out := new(VPNamespaceMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPNamespaceSpec) DeepCopyInto(out *VPNamespaceSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPNamespaceSpec.
func (in *VPNamespaceSpec) DeepCopy() *VPNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(VPNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPNamespaceStatus) DeepCopyInto(out *VPNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPNamespaceStatus.
func (in *VPNamespaceStatus) DeepCopy() *VPNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(VPNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}
