// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package route

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsecureEdgeTerminationPolicyType) DeepCopyInto(out *InsecureEdgeTerminationPolicyType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsecureEdgeTerminationPolicyType.
func (in *InsecureEdgeTerminationPolicyType) DeepCopy() *InsecureEdgeTerminationPolicyType {
	if in == nil {
		return nil
	}
	out := new(InsecureEdgeTerminationPolicyType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteIngress) DeepCopyInto(out *RouteIngress) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RouteIngressCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteIngress.
func (in *RouteIngress) DeepCopy() *RouteIngress {
	if in == nil {
		return nil
	}
	out := new(RouteIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteIngressCondition) DeepCopyInto(out *RouteIngressCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteIngressCondition.
func (in *RouteIngressCondition) DeepCopy() *RouteIngressCondition {
	if in == nil {
		return nil
	}
	out := new(RouteIngressCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteIngressConditionType) DeepCopyInto(out *RouteIngressConditionType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteIngressConditionType.
func (in *RouteIngressConditionType) DeepCopy() *RouteIngressConditionType {
	if in == nil {
		return nil
	}
	out := new(RouteIngressConditionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutePort) DeepCopyInto(out *RoutePort) {
	*out = *in
	out.TargetPort = in.TargetPort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutePort.
func (in *RoutePort) DeepCopy() *RoutePort {
	if in == nil {
		return nil
	}
	out := new(RoutePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.To.DeepCopyInto(&out.To)
	if in.AlternateBackends != nil {
		in, out := &in.AlternateBackends, &out.AlternateBackends
		*out = make([]RouteTargetReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		if *in == nil {
			*out = nil
		} else {
			*out = new(RoutePort)
			**out = **in
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]RouteIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTargetReference) DeepCopyInto(out *RouteTargetReference) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTargetReference.
func (in *RouteTargetReference) DeepCopy() *RouteTargetReference {
	if in == nil {
		return nil
	}
	out := new(RouteTargetReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterShard) DeepCopyInto(out *RouterShard) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterShard.
func (in *RouterShard) DeepCopy() *RouterShard {
	if in == nil {
		return nil
	}
	out := new(RouterShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSTerminationType) DeepCopyInto(out *TLSTerminationType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSTerminationType.
func (in *TLSTerminationType) DeepCopy() *TLSTerminationType {
	if in == nil {
		return nil
	}
	out := new(TLSTerminationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WildcardPolicyType) DeepCopyInto(out *WildcardPolicyType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WildcardPolicyType.
func (in *WildcardPolicyType) DeepCopy() *WildcardPolicyType {
	if in == nil {
		return nil
	}
	out := new(WildcardPolicyType)
	in.DeepCopyInto(out)
	return out
}
