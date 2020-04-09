// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package vsphere

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedDHCPState) DeepCopyInto(out *AdvancedDHCPState) {
	*out = *in
	if in.LogicalSwitchID != nil {
		in, out := &in.LogicalSwitchID, &out.LogicalSwitchID
		*out = new(string)
		**out = **in
	}
	if in.ProfileID != nil {
		in, out := &in.ProfileID, &out.ProfileID
		*out = new(string)
		**out = **in
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.IPPoolID != nil {
		in, out := &in.IPPoolID, &out.IPPoolID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedDHCPState.
func (in *AdvancedDHCPState) DeepCopy() *AdvancedDHCPState {
	if in == nil {
		return nil
	}
	out := new(AdvancedDHCPState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPLoadBalancerClass) DeepCopyInto(out *CPLoadBalancerClass) {
	*out = *in
	if in.IPPoolName != nil {
		in, out := &in.IPPoolName, &out.IPPoolName
		*out = new(string)
		**out = **in
	}
	if in.TCPAppProfileName != nil {
		in, out := &in.TCPAppProfileName, &out.TCPAppProfileName
		*out = new(string)
		**out = **in
	}
	if in.UDPAppProfileName != nil {
		in, out := &in.UDPAppProfileName, &out.UDPAppProfileName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPLoadBalancerClass.
func (in *CPLoadBalancerClass) DeepCopy() *CPLoadBalancerClass {
	if in == nil {
		return nil
	}
	out := new(CPLoadBalancerClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudControllerManagerConfig) DeepCopyInto(out *CloudControllerManagerConfig) {
	*out = *in
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudControllerManagerConfig.
func (in *CloudControllerManagerConfig) DeepCopy() *CloudControllerManagerConfig {
	if in == nil {
		return nil
	}
	out := new(CloudControllerManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProfileConfig) DeepCopyInto(out *CloudProfileConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]RegionSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureDomainLabels != nil {
		in, out := &in.FailureDomainLabels, &out.FailureDomainLabels
		*out = new(FailureDomainLabels)
		**out = **in
	}
	if in.DNSServers != nil {
		in, out := &in.DNSServers, &out.DNSServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MachineImages != nil {
		in, out := &in.MachineImages, &out.MachineImages
		*out = make([]MachineImages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Constraints.DeepCopyInto(&out.Constraints)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProfileConfig.
func (in *CloudProfileConfig) DeepCopy() *CloudProfileConfig {
	if in == nil {
		return nil
	}
	out := new(CloudProfileConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudProfileConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraints) DeepCopyInto(out *Constraints) {
	*out = *in
	in.LoadBalancerConfig.DeepCopyInto(&out.LoadBalancerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraints.
func (in *Constraints) DeepCopy() *Constraints {
	if in == nil {
		return nil
	}
	out := new(Constraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneConfig) DeepCopyInto(out *ControlPlaneConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.CloudControllerManager != nil {
		in, out := &in.CloudControllerManager, &out.CloudControllerManager
		*out = new(CloudControllerManagerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LoadBalancerClasses != nil {
		in, out := &in.LoadBalancerClasses, &out.LoadBalancerClasses
		*out = make([]CPLoadBalancerClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoadBalancerSize != nil {
		in, out := &in.LoadBalancerSize, &out.LoadBalancerSize
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneConfig.
func (in *ControlPlaneConfig) DeepCopy() *ControlPlaneConfig {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailureDomainLabels) DeepCopyInto(out *FailureDomainLabels) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailureDomainLabels.
func (in *FailureDomainLabels) DeepCopy() *FailureDomainLabels {
	if in == nil {
		return nil
	}
	out := new(FailureDomainLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureConfig) DeepCopyInto(out *InfrastructureConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureConfig.
func (in *InfrastructureConfig) DeepCopy() *InfrastructureConfig {
	if in == nil {
		return nil
	}
	out := new(InfrastructureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureStatus) DeepCopyInto(out *InfrastructureStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.VsphereConfig.DeepCopyInto(&out.VsphereConfig)
	if in.CreationStarted != nil {
		in, out := &in.CreationStarted, &out.CreationStarted
		*out = new(bool)
		**out = **in
	}
	if in.NSXTInfraState != nil {
		in, out := &in.NSXTInfraState, &out.NSXTInfraState
		*out = new(NSXTInfraState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureStatus.
func (in *InfrastructureStatus) DeepCopy() *InfrastructureStatus {
	if in == nil {
		return nil
	}
	out := new(InfrastructureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerClass) DeepCopyInto(out *LoadBalancerClass) {
	*out = *in
	if in.IPPoolName != nil {
		in, out := &in.IPPoolName, &out.IPPoolName
		*out = new(string)
		**out = **in
	}
	if in.TCPAppProfileName != nil {
		in, out := &in.TCPAppProfileName, &out.TCPAppProfileName
		*out = new(string)
		**out = **in
	}
	if in.UDPAppProfileName != nil {
		in, out := &in.UDPAppProfileName, &out.UDPAppProfileName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerClass.
func (in *LoadBalancerClass) DeepCopy() *LoadBalancerClass {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfig) DeepCopyInto(out *LoadBalancerConfig) {
	*out = *in
	if in.Classes != nil {
		in, out := &in.Classes, &out.Classes
		*out = make([]LoadBalancerClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfig.
func (in *LoadBalancerConfig) DeepCopy() *LoadBalancerConfig {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineImage) DeepCopyInto(out *MachineImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineImage.
func (in *MachineImage) DeepCopy() *MachineImage {
	if in == nil {
		return nil
	}
	out := new(MachineImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineImageVersion) DeepCopyInto(out *MachineImageVersion) {
	*out = *in
	if in.GuestID != nil {
		in, out := &in.GuestID, &out.GuestID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineImageVersion.
func (in *MachineImageVersion) DeepCopy() *MachineImageVersion {
	if in == nil {
		return nil
	}
	out := new(MachineImageVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineImages) DeepCopyInto(out *MachineImages) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]MachineImageVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineImages.
func (in *MachineImages) DeepCopy() *MachineImages {
	if in == nil {
		return nil
	}
	out := new(MachineImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSXTInfraState) DeepCopyInto(out *NSXTInfraState) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.EdgeClusterRef != nil {
		in, out := &in.EdgeClusterRef, &out.EdgeClusterRef
		*out = new(Reference)
		**out = **in
	}
	if in.TransportZoneRef != nil {
		in, out := &in.TransportZoneRef, &out.TransportZoneRef
		*out = new(Reference)
		**out = **in
	}
	if in.Tier0GatewayRef != nil {
		in, out := &in.Tier0GatewayRef, &out.Tier0GatewayRef
		*out = new(Reference)
		**out = **in
	}
	if in.SNATIPPoolRef != nil {
		in, out := &in.SNATIPPoolRef, &out.SNATIPPoolRef
		*out = new(Reference)
		**out = **in
	}
	if in.Tier1GatewayRef != nil {
		in, out := &in.Tier1GatewayRef, &out.Tier1GatewayRef
		*out = new(Reference)
		**out = **in
	}
	if in.LocaleServiceRef != nil {
		in, out := &in.LocaleServiceRef, &out.LocaleServiceRef
		*out = new(Reference)
		**out = **in
	}
	if in.SegmentRef != nil {
		in, out := &in.SegmentRef, &out.SegmentRef
		*out = new(Reference)
		**out = **in
	}
	if in.SNATIPAddressAllocRef != nil {
		in, out := &in.SNATIPAddressAllocRef, &out.SNATIPAddressAllocRef
		*out = new(Reference)
		**out = **in
	}
	if in.SNATRuleRef != nil {
		in, out := &in.SNATRuleRef, &out.SNATRuleRef
		*out = new(Reference)
		**out = **in
	}
	if in.SNATIPAddress != nil {
		in, out := &in.SNATIPAddress, &out.SNATIPAddress
		*out = new(string)
		**out = **in
	}
	if in.SegmentName != nil {
		in, out := &in.SegmentName, &out.SegmentName
		*out = new(string)
		**out = **in
	}
	if in.DHCPServerConfigRef != nil {
		in, out := &in.DHCPServerConfigRef, &out.DHCPServerConfigRef
		*out = new(Reference)
		**out = **in
	}
	in.AdvancedDHCP.DeepCopyInto(&out.AdvancedDHCP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSXTInfraState.
func (in *NSXTInfraState) DeepCopy() *NSXTInfraState {
	if in == nil {
		return nil
	}
	out := new(NSXTInfraState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionSpec) DeepCopyInto(out *RegionSpec) {
	*out = *in
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.DatastoreCluster != nil {
		in, out := &in.DatastoreCluster, &out.DatastoreCluster
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]ZoneSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CaFile != nil {
		in, out := &in.CaFile, &out.CaFile
		*out = new(string)
		**out = **in
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.DNSServers != nil {
		in, out := &in.DNSServers, &out.DNSServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MachineImages != nil {
		in, out := &in.MachineImages, &out.MachineImages
		*out = make([]MachineImages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionSpec.
func (in *RegionSpec) DeepCopy() *RegionSpec {
	if in == nil {
		return nil
	}
	out := new(RegionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereConfig) DeepCopyInto(out *VsphereConfig) {
	*out = *in
	if in.ZoneConfigs != nil {
		in, out := &in.ZoneConfigs, &out.ZoneConfigs
		*out = make(map[string]ZoneConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereConfig.
func (in *VsphereConfig) DeepCopy() *VsphereConfig {
	if in == nil {
		return nil
	}
	out := new(VsphereConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerStatus) DeepCopyInto(out *WorkerStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.MachineImages != nil {
		in, out := &in.MachineImages, &out.MachineImages
		*out = make([]MachineImage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerStatus.
func (in *WorkerStatus) DeepCopy() *WorkerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneConfig) DeepCopyInto(out *ZoneConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneConfig.
func (in *ZoneConfig) DeepCopy() *ZoneConfig {
	if in == nil {
		return nil
	}
	out := new(ZoneConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpec) DeepCopyInto(out *ZoneSpec) {
	*out = *in
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.ComputeCluster != nil {
		in, out := &in.ComputeCluster, &out.ComputeCluster
		*out = new(string)
		**out = **in
	}
	if in.ResourcePool != nil {
		in, out := &in.ResourcePool, &out.ResourcePool
		*out = new(string)
		**out = **in
	}
	if in.HostSystem != nil {
		in, out := &in.HostSystem, &out.HostSystem
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.DatastoreCluster != nil {
		in, out := &in.DatastoreCluster, &out.DatastoreCluster
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpec.
func (in *ZoneSpec) DeepCopy() *ZoneSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneSpec)
	in.DeepCopyInto(out)
	return out
}
