// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package testing

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_testing_UnknownObject, InType: reflect.TypeOf(&UnknownObject{})},
	}
}

// DeepCopy_testing_UnknownObject is an autogenerated deepcopy function.
func DeepCopy_testing_UnknownObject(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UnknownObject)
		out := out.(*UnknownObject)
		*out = *in
		return nil
	}
}