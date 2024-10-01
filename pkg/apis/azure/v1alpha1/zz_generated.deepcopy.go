//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilitySet) DeepCopyInto(out *AvailabilitySet) {
	*out = *in
	if in.CountFaultDomains != nil {
		in, out := &in.CountFaultDomains, &out.CountFaultDomains
		*out = new(int32)
		**out = **in
	}
	if in.CountUpdateDomains != nil {
		in, out := &in.CountUpdateDomains, &out.CountUpdateDomains
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilitySet.
func (in *AvailabilitySet) DeepCopy() *AvailabilitySet {
	if in == nil {
		return nil
	}
	out := new(AvailabilitySet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureResource) DeepCopyInto(out *AzureResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureResource.
func (in *AzureResource) DeepCopy() *AzureResource {
	if in == nil {
		return nil
	}
	out := new(AzureResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketConfig) DeepCopyInto(out *BackupBucketConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.CloudConfiguration != nil {
		in, out := &in.CloudConfiguration, &out.CloudConfiguration
		*out = new(CloudConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketConfig.
func (in *BackupBucketConfig) DeepCopy() *BackupBucketConfig {
	if in == nil {
		return nil
	}
	out := new(BackupBucketConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupBucketConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudConfiguration) DeepCopyInto(out *CloudConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudConfiguration.
func (in *CloudConfiguration) DeepCopy() *CloudConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudConfiguration)
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
	if in.CountUpdateDomains != nil {
		in, out := &in.CountUpdateDomains, &out.CountUpdateDomains
		*out = make([]DomainCount, len(*in))
		copy(*out, *in)
	}
	if in.CountFaultDomains != nil {
		in, out := &in.CountFaultDomains, &out.CountFaultDomains
		*out = make([]DomainCount, len(*in))
		copy(*out, *in)
	}
	if in.MachineImages != nil {
		in, out := &in.MachineImages, &out.MachineImages
		*out = make([]MachineImages, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MachineTypes != nil {
		in, out := &in.MachineTypes, &out.MachineTypes
		*out = make([]MachineType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudConfiguration != nil {
		in, out := &in.CloudConfiguration, &out.CloudConfiguration
		*out = new(CloudConfiguration)
		**out = **in
	}
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
func (in *ControlPlaneConfig) DeepCopyInto(out *ControlPlaneConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.CloudControllerManager != nil {
		in, out := &in.CloudControllerManager, &out.CloudControllerManager
		*out = new(CloudControllerManagerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
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
func (in *DNSRecordConfig) DeepCopyInto(out *DNSRecordConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.CloudConfiguration != nil {
		in, out := &in.CloudConfiguration, &out.CloudConfiguration
		*out = new(CloudConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordConfig.
func (in *DNSRecordConfig) DeepCopy() *DNSRecordConfig {
	if in == nil {
		return nil
	}
	out := new(DNSRecordConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecordConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolume) DeepCopyInto(out *DataVolume) {
	*out = *in
	if in.ImageRef != nil {
		in, out := &in.ImageRef, &out.ImageRef
		*out = new(Image)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolume.
func (in *DataVolume) DeepCopy() *DataVolume {
	if in == nil {
		return nil
	}
	out := new(DataVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnosticsProfile) DeepCopyInto(out *DiagnosticsProfile) {
	*out = *in
	if in.StorageURI != nil {
		in, out := &in.StorageURI, &out.StorageURI
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnosticsProfile.
func (in *DiagnosticsProfile) DeepCopy() *DiagnosticsProfile {
	if in == nil {
		return nil
	}
	out := new(DiagnosticsProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainCount) DeepCopyInto(out *DomainCount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainCount.
func (in *DomainCount) DeepCopy() *DomainCount {
	if in == nil {
		return nil
	}
	out := new(DomainCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityConfig) DeepCopyInto(out *IdentityConfig) {
	*out = *in
	if in.ACRAccess != nil {
		in, out := &in.ACRAccess, &out.ACRAccess
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityConfig.
func (in *IdentityConfig) DeepCopy() *IdentityConfig {
	if in == nil {
		return nil
	}
	out := new(IdentityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityStatus) DeepCopyInto(out *IdentityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityStatus.
func (in *IdentityStatus) DeepCopy() *IdentityStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.URN != nil {
		in, out := &in.URN, &out.URN
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.CommunityGalleryImageID != nil {
		in, out := &in.CommunityGalleryImageID, &out.CommunityGalleryImageID
		*out = new(string)
		**out = **in
	}
	if in.SharedGalleryImageID != nil {
		in, out := &in.SharedGalleryImageID, &out.SharedGalleryImageID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureConfig) DeepCopyInto(out *InfrastructureConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(ResourceGroup)
		**out = **in
	}
	in.Networks.DeepCopyInto(&out.Networks)
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityConfig)
		(*in).DeepCopyInto(*out)
	}
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
func (in *InfrastructureState) DeepCopyInto(out *InfrastructureState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ManagedItems != nil {
		in, out := &in.ManagedItems, &out.ManagedItems
		*out = make([]AzureResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureState.
func (in *InfrastructureState) DeepCopy() *InfrastructureState {
	if in == nil {
		return nil
	}
	out := new(InfrastructureState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureStatus) DeepCopyInto(out *InfrastructureStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Networks.DeepCopyInto(&out.Networks)
	out.ResourceGroup = in.ResourceGroup
	if in.AvailabilitySets != nil {
		in, out := &in.AvailabilitySets, &out.AvailabilitySets
		*out = make([]AvailabilitySet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RouteTables != nil {
		in, out := &in.RouteTables, &out.RouteTables
		*out = make([]RouteTable, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]SecurityGroup, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityStatus)
		**out = **in
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
func (in *MachineImage) DeepCopyInto(out *MachineImage) {
	*out = *in
	if in.AcceleratedNetworking != nil {
		in, out := &in.AcceleratedNetworking, &out.AcceleratedNetworking
		*out = new(bool)
		**out = **in
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = new(string)
		**out = **in
	}
	if in.SkipMarketplaceAgreement != nil {
		in, out := &in.SkipMarketplaceAgreement, &out.SkipMarketplaceAgreement
		*out = new(bool)
		**out = **in
	}
	in.ImageRef.DeepCopyInto(&out.ImageRef)
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
	if in.URN != nil {
		in, out := &in.URN, &out.URN
		*out = new(string)
		**out = **in
	}
	if in.SkipMarketplaceAgreement != nil {
		in, out := &in.SkipMarketplaceAgreement, &out.SkipMarketplaceAgreement
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.CommunityGalleryImageID != nil {
		in, out := &in.CommunityGalleryImageID, &out.CommunityGalleryImageID
		*out = new(string)
		**out = **in
	}
	if in.SharedGalleryImageID != nil {
		in, out := &in.SharedGalleryImageID, &out.SharedGalleryImageID
		*out = new(string)
		**out = **in
	}
	if in.AcceleratedNetworking != nil {
		in, out := &in.AcceleratedNetworking, &out.AcceleratedNetworking
		*out = new(bool)
		**out = **in
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
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
func (in *MachineType) DeepCopyInto(out *MachineType) {
	*out = *in
	if in.AcceleratedNetworking != nil {
		in, out := &in.AcceleratedNetworking, &out.AcceleratedNetworking
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineType.
func (in *MachineType) DeepCopy() *MachineType {
	if in == nil {
		return nil
	}
	out := new(MachineType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatGatewayConfig) DeepCopyInto(out *NatGatewayConfig) {
	*out = *in
	if in.IdleConnectionTimeoutMinutes != nil {
		in, out := &in.IdleConnectionTimeoutMinutes, &out.IdleConnectionTimeoutMinutes
		*out = new(int32)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(int32)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]PublicIPReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatGatewayConfig.
func (in *NatGatewayConfig) DeepCopy() *NatGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(NatGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	in.VNet.DeepCopyInto(&out.VNet)
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = new(string)
		**out = **in
	}
	if in.NatGateway != nil {
		in, out := &in.NatGateway, &out.NatGateway
		*out = new(NatGatewayConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceEndpoints != nil {
		in, out := &in.ServiceEndpoints, &out.ServiceEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]Zone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	in.VNet.DeepCopyInto(&out.VNet)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPReference) DeepCopyInto(out *PublicIPReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPReference.
func (in *PublicIPReference) DeepCopy() *PublicIPReference {
	if in == nil {
		return nil
	}
	out := new(PublicIPReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroup) DeepCopyInto(out *ResourceGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroup.
func (in *ResourceGroup) DeepCopy() *ResourceGroup {
	if in == nil {
		return nil
	}
	out := new(ResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTable.
func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.ManagedDefaultStorageClass != nil {
		in, out := &in.ManagedDefaultStorageClass, &out.ManagedDefaultStorageClass
		*out = new(bool)
		**out = **in
	}
	if in.ManagedDefaultVolumeSnapshotClass != nil {
		in, out := &in.ManagedDefaultVolumeSnapshotClass, &out.ManagedDefaultVolumeSnapshotClass
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNet) DeepCopyInto(out *VNet) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = new(string)
		**out = **in
	}
	if in.DDosProtectionPlanID != nil {
		in, out := &in.DDosProtectionPlanID, &out.DDosProtectionPlanID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNet.
func (in *VNet) DeepCopy() *VNet {
	if in == nil {
		return nil
	}
	out := new(VNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VNetStatus) DeepCopyInto(out *VNetStatus) {
	*out = *in
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VNetStatus.
func (in *VNetStatus) DeepCopy() *VNetStatus {
	if in == nil {
		return nil
	}
	out := new(VNetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmoDependency) DeepCopyInto(out *VmoDependency) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmoDependency.
func (in *VmoDependency) DeepCopy() *VmoDependency {
	if in == nil {
		return nil
	}
	out := new(VmoDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfig) DeepCopyInto(out *WorkerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.NodeTemplate != nil {
		in, out := &in.NodeTemplate, &out.NodeTemplate
		*out = new(extensionsv1alpha1.NodeTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.DiagnosticsProfile != nil {
		in, out := &in.DiagnosticsProfile, &out.DiagnosticsProfile
		*out = new(DiagnosticsProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.DataVolumes != nil {
		in, out := &in.DataVolumes, &out.DataVolumes
		*out = make([]DataVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfig.
func (in *WorkerConfig) DeepCopy() *WorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerStatus) DeepCopyInto(out *WorkerStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.MachineImages != nil {
		in, out := &in.MachineImages, &out.MachineImages
		*out = make([]MachineImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VmoDependencies != nil {
		in, out := &in.VmoDependencies, &out.VmoDependencies
		*out = make([]VmoDependency, len(*in))
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
func (in *Zone) DeepCopyInto(out *Zone) {
	*out = *in
	if in.ServiceEndpoints != nil {
		in, out := &in.ServiceEndpoints, &out.ServiceEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NatGateway != nil {
		in, out := &in.NatGateway, &out.NatGateway
		*out = new(ZonedNatGatewayConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Zone.
func (in *Zone) DeepCopy() *Zone {
	if in == nil {
		return nil
	}
	out := new(Zone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZonedNatGatewayConfig) DeepCopyInto(out *ZonedNatGatewayConfig) {
	*out = *in
	if in.IdleConnectionTimeoutMinutes != nil {
		in, out := &in.IdleConnectionTimeoutMinutes, &out.IdleConnectionTimeoutMinutes
		*out = new(int32)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]ZonedPublicIPReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZonedNatGatewayConfig.
func (in *ZonedNatGatewayConfig) DeepCopy() *ZonedNatGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(ZonedNatGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZonedPublicIPReference) DeepCopyInto(out *ZonedPublicIPReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZonedPublicIPReference.
func (in *ZonedPublicIPReference) DeepCopy() *ZonedPublicIPReference {
	if in == nil {
		return nil
	}
	out := new(ZonedPublicIPReference)
	in.DeepCopyInto(out)
	return out
}
