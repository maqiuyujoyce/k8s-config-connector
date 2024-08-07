//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContent) DeepCopyInto(out *FileContent) {
	*out = *in
	in.Content.DeepCopyInto(&out.Content)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContent.
func (in *FileContent) DeepCopy() *FileContent {
	if in == nil {
		return nil
	}
	out := new(FileContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmConfiguration) DeepCopyInto(out *HelmConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmConfiguration.
func (in *HelmConfiguration) DeepCopy() *HelmConfiguration {
	if in == nil {
		return nil
	}
	out := new(HelmConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmConfigurationList) DeepCopyInto(out *HelmConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmConfigurationList.
func (in *HelmConfigurationList) DeepCopy() *HelmConfigurationList {
	if in == nil {
		return nil
	}
	out := new(HelmConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmConfigurationSpec) DeepCopyInto(out *HelmConfigurationSpec) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
	in.DefaultValues.DeepCopyInto(&out.DefaultValues)
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]FileContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CRDs != nil {
		in, out := &in.CRDs, &out.CRDs
		*out = make([]FileContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmConfigurationSpec.
func (in *HelmConfigurationSpec) DeepCopy() *HelmConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(HelmConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmConfigurationStatus) DeepCopyInto(out *HelmConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmConfigurationStatus.
func (in *HelmConfigurationStatus) DeepCopy() *HelmConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(HelmConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
