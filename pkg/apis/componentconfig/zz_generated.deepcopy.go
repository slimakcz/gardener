// +build !ignore_autogenerated

/*
Copyright 2018 The Gardener Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package componentconfig

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionConfiguration) DeepCopyInto(out *ClientConnectionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionConfiguration.
func (in *ClientConnectionConfiguration) DeepCopy() *ClientConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ClientConnection = in.ClientConnection
	if in.GardenerClientConnection != nil {
		in, out := &in.GardenerClientConnection, &out.GardenerClientConnection
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClientConnectionConfiguration)
			**out = **in
		}
	}
	in.Controller.DeepCopyInto(&out.Controller)
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ControllerManagerImagesConfiguration, len(*in))
		copy(*out, *in)
	}
	out.LeaderElection = in.LeaderElection
	out.Metrics = in.Metrics
	out.Server = in.Server
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerControllerConfiguration) DeepCopyInto(out *ControllerManagerControllerConfiguration) {
	*out = *in
	out.HealthCheckPeriod = in.HealthCheckPeriod
	in.Reconciliation.DeepCopyInto(&out.Reconciliation)
	if in.WatchNamespace != nil {
		in, out := &in.WatchNamespace, &out.WatchNamespace
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerControllerConfiguration.
func (in *ControllerManagerControllerConfiguration) DeepCopy() *ControllerManagerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerImagesConfiguration) DeepCopyInto(out *ControllerManagerImagesConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerImagesConfiguration.
func (in *ControllerManagerImagesConfiguration) DeepCopy() *ControllerManagerImagesConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerImagesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerReconciliationConfiguration) DeepCopyInto(out *ControllerReconciliationConfiguration) {
	*out = *in
	out.ResyncPeriod = in.ResyncPeriod
	if in.RetryDuration != nil {
		in, out := &in.RetryDuration, &out.RetryDuration
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerReconciliationConfiguration.
func (in *ControllerReconciliationConfiguration) DeepCopy() *ControllerReconciliationConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerReconciliationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfiguration) DeepCopyInto(out *MetricsConfiguration) {
	*out = *in
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfiguration.
func (in *MetricsConfiguration) DeepCopy() *MetricsConfiguration {
	if in == nil {
		return nil
	}
	out := new(MetricsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}
