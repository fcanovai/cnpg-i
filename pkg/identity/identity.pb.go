// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.2
// source: proto/identity.proto

package identity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginCapability_Service_Type int32

const (
	PluginCapability_Service_TYPE_UNSPECIFIED PluginCapability_Service_Type = 0
	// TYPE_OPERATOR_SERVICE indicated that the Plugin provider RPCs
	// for the Operator service.
	// The presence of this capability determines whether the CO will
	// attempt to invoke the REQUIRED Operator RPCs, as well
	// as specific RPCs as indicated by GetCapabilities.
	PluginCapability_Service_TYPE_OPERATOR_SERVICE PluginCapability_Service_Type = 1
	// TYPE_WAL_SERVICE indicates that the Plugin provides RPCs for
	// the WAL service. Plugins MAY provide this capability.
	// The presence of this capability determines whether the CO will
	// attempt to invoke the REQUIRED WALService RPCs, as well
	// as specific RPCs as indicated by GetCapabilities.
	PluginCapability_Service_TYPE_WAL_SERVICE PluginCapability_Service_Type = 2
	// TYPE_BACKUP_SERVICE indicates that the Plugin provides RPCs for
	// the Backup service.
	// The presence of this capability determines whether the CO will
	// attempt to invoke the REQUIRED Backup Service RPCs, as well
	// as specific RPCs as indicated by GetCapabilities.
	PluginCapability_Service_TYPE_BACKUP_SERVICE PluginCapability_Service_Type = 3
	// TYPE_LIFECYCLE_SERVICE indicates that the Plugin provides RPCs for
	// the Lifecycle service.
	PluginCapability_Service_TYPE_LIFECYCLE_SERVICE PluginCapability_Service_Type = 4
	// TYPE_RECONCILER_HOOKS indicates that the Plugin provides RPCs to
	// enhance the behavior of the reconcilers
	PluginCapability_Service_TYPE_RECONCILER_HOOKS PluginCapability_Service_Type = 5
)

// Enum value maps for PluginCapability_Service_Type.
var (
	PluginCapability_Service_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_OPERATOR_SERVICE",
		2: "TYPE_WAL_SERVICE",
		3: "TYPE_BACKUP_SERVICE",
		4: "TYPE_LIFECYCLE_SERVICE",
		5: "TYPE_RECONCILER_HOOKS",
	}
	PluginCapability_Service_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":       0,
		"TYPE_OPERATOR_SERVICE":  1,
		"TYPE_WAL_SERVICE":       2,
		"TYPE_BACKUP_SERVICE":    3,
		"TYPE_LIFECYCLE_SERVICE": 4,
		"TYPE_RECONCILER_HOOKS":  5,
	}
)

func (x PluginCapability_Service_Type) Enum() *PluginCapability_Service_Type {
	p := new(PluginCapability_Service_Type)
	*p = x
	return p
}

func (x PluginCapability_Service_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginCapability_Service_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_identity_proto_enumTypes[0].Descriptor()
}

func (PluginCapability_Service_Type) Type() protoreflect.EnumType {
	return &file_proto_identity_proto_enumTypes[0]
}

func (x PluginCapability_Service_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginCapability_Service_Type.Descriptor instead.
func (PluginCapability_Service_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{4, 0, 0}
}

type GetPluginMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPluginMetadataRequest) Reset() {
	*x = GetPluginMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginMetadataRequest) ProtoMessage() {}

func (x *GetPluginMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetPluginMetadataRequest) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{0}
}

type GetPluginMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name MUST follow domain name notation format
	// (https://tools.ietf.org/html/rfc1035#section-2.3.1). It SHOULD
	// include the plugin's host company name and the plugin name,
	// to minimize the possibility of collisions. It MUST be 63
	// characters or less, beginning and ending with an alphanumeric
	// character ([a-z0-9A-Z]) with dashes (-), dots (.), and
	// alphanumerics between. This field is REQUIRED.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This field is REQUIRED. Value of this field is opaque.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// A name to display for the plugin. This field is REQUIRED.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A description for the plugin. This field is REQUIRED.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// URL of the home page of the plugin project.
	ProjectUrl string `protobuf:"bytes,5,opt,name=project_url,json=projectUrl,proto3" json:"project_url,omitempty"`
	// URL of the source code repository for the plugin project.
	RepositoryUrl string `protobuf:"bytes,6,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty"`
	// License of the plugin. This field is REQUIRED.
	License string `protobuf:"bytes,7,opt,name=license,proto3" json:"license,omitempty"`
	// URL of the license of the plugin. This field is REQUIRED.
	LicenseUrl string `protobuf:"bytes,8,opt,name=license_url,json=licenseUrl,proto3" json:"license_url,omitempty"`
	// Maturity level (alpha, beta, stable)
	Maturity string `protobuf:"bytes,9,opt,name=maturity,proto3" json:"maturity,omitempty"`
	// Provider/vendor of the plugin, e.g. an organization
	Vendor string `protobuf:"bytes,10,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// This field is OPTIONAL. Values are opaque.
	Manifest map[string]string `protobuf:"bytes,11,rep,name=manifest,proto3" json:"manifest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPluginMetadataResponse) Reset() {
	*x = GetPluginMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginMetadataResponse) ProtoMessage() {}

func (x *GetPluginMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetPluginMetadataResponse) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{1}
}

func (x *GetPluginMetadataResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetProjectUrl() string {
	if x != nil {
		return x.ProjectUrl
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetRepositoryUrl() string {
	if x != nil {
		return x.RepositoryUrl
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetLicenseUrl() string {
	if x != nil {
		return x.LicenseUrl
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetMaturity() string {
	if x != nil {
		return x.Maturity
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *GetPluginMetadataResponse) GetManifest() map[string]string {
	if x != nil {
		return x.Manifest
	}
	return nil
}

type GetPluginCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPluginCapabilitiesRequest) Reset() {
	*x = GetPluginCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginCapabilitiesRequest) ProtoMessage() {}

func (x *GetPluginCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*GetPluginCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{2}
}

type GetPluginCapabilitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the capabilities that the controller service supports. This
	// field is OPTIONAL.
	Capabilities []*PluginCapability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
}

func (x *GetPluginCapabilitiesResponse) Reset() {
	*x = GetPluginCapabilitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginCapabilitiesResponse) ProtoMessage() {}

func (x *GetPluginCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginCapabilitiesResponse.ProtoReflect.Descriptor instead.
func (*GetPluginCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{3}
}

func (x *GetPluginCapabilitiesResponse) GetCapabilities() []*PluginCapability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type PluginCapability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*PluginCapability_Service_
	Type isPluginCapability_Type `protobuf_oneof:"type"`
}

func (x *PluginCapability) Reset() {
	*x = PluginCapability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginCapability) ProtoMessage() {}

func (x *PluginCapability) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginCapability.ProtoReflect.Descriptor instead.
func (*PluginCapability) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{4}
}

func (m *PluginCapability) GetType() isPluginCapability_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *PluginCapability) GetService() *PluginCapability_Service {
	if x, ok := x.GetType().(*PluginCapability_Service_); ok {
		return x.Service
	}
	return nil
}

type isPluginCapability_Type interface {
	isPluginCapability_Type()
}

type PluginCapability_Service_ struct {
	Service *PluginCapability_Service `protobuf:"bytes,1,opt,name=service,proto3,oneof"`
}

func (*PluginCapability_Service_) isPluginCapability_Type() {}

type ProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProbeRequest) Reset() {
	*x = ProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeRequest) ProtoMessage() {}

func (x *ProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeRequest.ProtoReflect.Descriptor instead.
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{5}
}

type ProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is OPTIONAL. If not present, the caller SHALL assume
	// that the plugin is in a ready state and is accepting calls to its
	// Controller and/or Node services (according to the plugin's reported
	// capabilities).
	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ProbeResponse) Reset() {
	*x = ProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeResponse) ProtoMessage() {}

func (x *ProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeResponse.ProtoReflect.Descriptor instead.
func (*ProbeResponse) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{6}
}

func (x *ProbeResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

type PluginCapability_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PluginCapability_Service_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.identity.v1.PluginCapability_Service_Type" json:"type,omitempty"`
}

func (x *PluginCapability_Service) Reset() {
	*x = PluginCapability_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginCapability_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginCapability_Service) ProtoMessage() {}

func (x *PluginCapability_Service) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginCapability_Service.ProtoReflect.Descriptor instead.
func (*PluginCapability_Service) Descriptor() ([]byte, []int) {
	return file_proto_identity_proto_rawDescGZIP(), []int{4, 0}
}

func (x *PluginCapability_Service) GetType() PluginCapability_Service_Type {
	if x != nil {
		return x.Type
	}
	return PluginCapability_Service_TYPE_UNSPECIFIED
}

var File_proto_identity_proto protoreflect.FileDescriptor

var file_proto_identity_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xda, 0x03, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x56, 0x0a,
	0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x68, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6e, 0x70, 0x67,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xd5, 0x02, 0x0a,
	0x10, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x47, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0xef, 0x01, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x9d, 0x01, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49,
	0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10,
	0x04, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x43,
	0x49, 0x4c, 0x45, 0x52, 0x5f, 0x48, 0x4f, 0x4f, 0x4b, 0x53, 0x10, 0x05, 0x42, 0x06, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x32, 0xc8, 0x02, 0x0a, 0x08,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x70, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x2e,
	0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6e, 0x70,
	0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x12, 0x1f, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x2d, 0x70, 0x67, 0x2f, 0x63, 0x6e, 0x70, 0x67, 0x2d, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_identity_proto_rawDescOnce sync.Once
	file_proto_identity_proto_rawDescData = file_proto_identity_proto_rawDesc
)

func file_proto_identity_proto_rawDescGZIP() []byte {
	file_proto_identity_proto_rawDescOnce.Do(func() {
		file_proto_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_identity_proto_rawDescData)
	})
	return file_proto_identity_proto_rawDescData
}

var file_proto_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_identity_proto_goTypes = []interface{}{
	(PluginCapability_Service_Type)(0),    // 0: cnpgi.identity.v1.PluginCapability.Service.Type
	(*GetPluginMetadataRequest)(nil),      // 1: cnpgi.identity.v1.GetPluginMetadataRequest
	(*GetPluginMetadataResponse)(nil),     // 2: cnpgi.identity.v1.GetPluginMetadataResponse
	(*GetPluginCapabilitiesRequest)(nil),  // 3: cnpgi.identity.v1.GetPluginCapabilitiesRequest
	(*GetPluginCapabilitiesResponse)(nil), // 4: cnpgi.identity.v1.GetPluginCapabilitiesResponse
	(*PluginCapability)(nil),              // 5: cnpgi.identity.v1.PluginCapability
	(*ProbeRequest)(nil),                  // 6: cnpgi.identity.v1.ProbeRequest
	(*ProbeResponse)(nil),                 // 7: cnpgi.identity.v1.ProbeResponse
	nil,                                   // 8: cnpgi.identity.v1.GetPluginMetadataResponse.ManifestEntry
	(*PluginCapability_Service)(nil),      // 9: cnpgi.identity.v1.PluginCapability.Service
}
var file_proto_identity_proto_depIdxs = []int32{
	8, // 0: cnpgi.identity.v1.GetPluginMetadataResponse.manifest:type_name -> cnpgi.identity.v1.GetPluginMetadataResponse.ManifestEntry
	5, // 1: cnpgi.identity.v1.GetPluginCapabilitiesResponse.capabilities:type_name -> cnpgi.identity.v1.PluginCapability
	9, // 2: cnpgi.identity.v1.PluginCapability.service:type_name -> cnpgi.identity.v1.PluginCapability.Service
	0, // 3: cnpgi.identity.v1.PluginCapability.Service.type:type_name -> cnpgi.identity.v1.PluginCapability.Service.Type
	1, // 4: cnpgi.identity.v1.Identity.GetPluginMetadata:input_type -> cnpgi.identity.v1.GetPluginMetadataRequest
	3, // 5: cnpgi.identity.v1.Identity.GetPluginCapabilities:input_type -> cnpgi.identity.v1.GetPluginCapabilitiesRequest
	6, // 6: cnpgi.identity.v1.Identity.Probe:input_type -> cnpgi.identity.v1.ProbeRequest
	2, // 7: cnpgi.identity.v1.Identity.GetPluginMetadata:output_type -> cnpgi.identity.v1.GetPluginMetadataResponse
	4, // 8: cnpgi.identity.v1.Identity.GetPluginCapabilities:output_type -> cnpgi.identity.v1.GetPluginCapabilitiesResponse
	7, // 9: cnpgi.identity.v1.Identity.Probe:output_type -> cnpgi.identity.v1.ProbeResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_identity_proto_init() }
func file_proto_identity_proto_init() {
	if File_proto_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginMetadataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginCapabilitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginCapabilitiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginCapability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_identity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginCapability_Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_identity_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PluginCapability_Service_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_identity_proto_goTypes,
		DependencyIndexes: file_proto_identity_proto_depIdxs,
		EnumInfos:         file_proto_identity_proto_enumTypes,
		MessageInfos:      file_proto_identity_proto_msgTypes,
	}.Build()
	File_proto_identity_proto = out.File
	file_proto_identity_proto_rawDesc = nil
	file_proto_identity_proto_goTypes = nil
	file_proto_identity_proto_depIdxs = nil
}
