// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_http_fastcgi.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 创建Fastcgi服务
type CreateHTTPFastcgiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOn            bool   `protobuf:"varint,1,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Address         string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ParamsJSON      []byte `protobuf:"bytes,3,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	ReadTimeoutJSON []byte `protobuf:"bytes,4,opt,name=readTimeoutJSON,proto3" json:"readTimeoutJSON,omitempty"`
	ConnTimeoutJSON []byte `protobuf:"bytes,5,opt,name=connTimeoutJSON,proto3" json:"connTimeoutJSON,omitempty"`
	PoolSize        int32  `protobuf:"varint,6,opt,name=poolSize,proto3" json:"poolSize,omitempty"`
	PathInfoPattern string `protobuf:"bytes,7,opt,name=pathInfoPattern,proto3" json:"pathInfoPattern,omitempty"`
}

func (x *CreateHTTPFastcgiRequest) Reset() {
	*x = CreateHTTPFastcgiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPFastcgiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPFastcgiRequest) ProtoMessage() {}

func (x *CreateHTTPFastcgiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPFastcgiRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPFastcgiRequest) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHTTPFastcgiRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *CreateHTTPFastcgiRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateHTTPFastcgiRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *CreateHTTPFastcgiRequest) GetReadTimeoutJSON() []byte {
	if x != nil {
		return x.ReadTimeoutJSON
	}
	return nil
}

func (x *CreateHTTPFastcgiRequest) GetConnTimeoutJSON() []byte {
	if x != nil {
		return x.ConnTimeoutJSON
	}
	return nil
}

func (x *CreateHTTPFastcgiRequest) GetPoolSize() int32 {
	if x != nil {
		return x.PoolSize
	}
	return 0
}

func (x *CreateHTTPFastcgiRequest) GetPathInfoPattern() string {
	if x != nil {
		return x.PathInfoPattern
	}
	return ""
}

type CreateHTTPFastcgiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFastcgiId int64 `protobuf:"varint,1,opt,name=httpFastcgiId,proto3" json:"httpFastcgiId,omitempty"`
}

func (x *CreateHTTPFastcgiResponse) Reset() {
	*x = CreateHTTPFastcgiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPFastcgiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPFastcgiResponse) ProtoMessage() {}

func (x *CreateHTTPFastcgiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPFastcgiResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPFastcgiResponse) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHTTPFastcgiResponse) GetHttpFastcgiId() int64 {
	if x != nil {
		return x.HttpFastcgiId
	}
	return 0
}

// 修改Fastcgi服务
type UpdateHTTPFastcgiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFastcgiId   int64  `protobuf:"varint,1,opt,name=httpFastcgiId,proto3" json:"httpFastcgiId,omitempty"`
	IsOn            bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Address         string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ParamsJSON      []byte `protobuf:"bytes,4,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	ReadTimeoutJSON []byte `protobuf:"bytes,5,opt,name=readTimeoutJSON,proto3" json:"readTimeoutJSON,omitempty"`
	ConnTimeoutJSON []byte `protobuf:"bytes,6,opt,name=connTimeoutJSON,proto3" json:"connTimeoutJSON,omitempty"`
	PoolSize        int32  `protobuf:"varint,7,opt,name=poolSize,proto3" json:"poolSize,omitempty"`
	PathInfoPattern string `protobuf:"bytes,8,opt,name=pathInfoPattern,proto3" json:"pathInfoPattern,omitempty"`
}

func (x *UpdateHTTPFastcgiRequest) Reset() {
	*x = UpdateHTTPFastcgiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPFastcgiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPFastcgiRequest) ProtoMessage() {}

func (x *UpdateHTTPFastcgiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPFastcgiRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPFastcgiRequest) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHTTPFastcgiRequest) GetHttpFastcgiId() int64 {
	if x != nil {
		return x.HttpFastcgiId
	}
	return 0
}

func (x *UpdateHTTPFastcgiRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *UpdateHTTPFastcgiRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateHTTPFastcgiRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *UpdateHTTPFastcgiRequest) GetReadTimeoutJSON() []byte {
	if x != nil {
		return x.ReadTimeoutJSON
	}
	return nil
}

func (x *UpdateHTTPFastcgiRequest) GetConnTimeoutJSON() []byte {
	if x != nil {
		return x.ConnTimeoutJSON
	}
	return nil
}

func (x *UpdateHTTPFastcgiRequest) GetPoolSize() int32 {
	if x != nil {
		return x.PoolSize
	}
	return 0
}

func (x *UpdateHTTPFastcgiRequest) GetPathInfoPattern() string {
	if x != nil {
		return x.PathInfoPattern
	}
	return ""
}

// 获取Fastcgi详情
type FindEnabledHTTPFastcgiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFastcgiId int64 `protobuf:"varint,1,opt,name=httpFastcgiId,proto3" json:"httpFastcgiId,omitempty"`
}

func (x *FindEnabledHTTPFastcgiRequest) Reset() {
	*x = FindEnabledHTTPFastcgiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFastcgiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFastcgiRequest) ProtoMessage() {}

func (x *FindEnabledHTTPFastcgiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFastcgiRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFastcgiRequest) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledHTTPFastcgiRequest) GetHttpFastcgiId() int64 {
	if x != nil {
		return x.HttpFastcgiId
	}
	return 0
}

type FindEnabledHTTPFastcgiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFastcgi *HTTPFastcgi `protobuf:"bytes,1,opt,name=httpFastcgi,proto3" json:"httpFastcgi,omitempty"`
}

func (x *FindEnabledHTTPFastcgiResponse) Reset() {
	*x = FindEnabledHTTPFastcgiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFastcgiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFastcgiResponse) ProtoMessage() {}

func (x *FindEnabledHTTPFastcgiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFastcgiResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFastcgiResponse) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledHTTPFastcgiResponse) GetHttpFastcgi() *HTTPFastcgi {
	if x != nil {
		return x.HttpFastcgi
	}
	return nil
}

// 获取Fastcgi配置
type FindEnabledHTTPFastcgiConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFastcgiId int64 `protobuf:"varint,1,opt,name=httpFastcgiId,proto3" json:"httpFastcgiId,omitempty"`
}

func (x *FindEnabledHTTPFastcgiConfigRequest) Reset() {
	*x = FindEnabledHTTPFastcgiConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFastcgiConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFastcgiConfigRequest) ProtoMessage() {}

func (x *FindEnabledHTTPFastcgiConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFastcgiConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFastcgiConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledHTTPFastcgiConfigRequest) GetHttpFastcgiId() int64 {
	if x != nil {
		return x.HttpFastcgiId
	}
	return 0
}

type FindEnabledHTTPFastcgiConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpFastcgiJSON []byte `protobuf:"bytes,1,opt,name=httpFastcgiJSON,proto3" json:"httpFastcgiJSON,omitempty"`
}

func (x *FindEnabledHTTPFastcgiConfigResponse) Reset() {
	*x = FindEnabledHTTPFastcgiConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_fastcgi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFastcgiConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFastcgiConfigResponse) ProtoMessage() {}

func (x *FindEnabledHTTPFastcgiConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_fastcgi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFastcgiConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFastcgiConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_fastcgi_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledHTTPFastcgiConfigResponse) GetHttpFastcgiJSON() []byte {
	if x != nil {
		return x.HttpFastcgiJSON
	}
	return nil
}

var File_service_http_fastcgi_proto protoreflect.FileDescriptor

var file_service_http_fastcgi_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63,
	0x67, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x22, 0x41, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63,
	0x67, 0x69, 0x49, 0x64, 0x22, 0xa8, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61,
	0x73, 0x74, 0x63, 0x67, 0x69, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f,
	0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53,
	0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6f,
	0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x6f,
	0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22,
	0x45, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73,
	0x74, 0x63, 0x67, 0x69, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70,
	0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x0b,
	0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x22, 0x4b, 0x0a, 0x23, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61,
	0x73, 0x74, 0x63, 0x67, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67,
	0x69, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x46,
	0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63,
	0x67, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x46,
	0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0xfd, 0x02, 0x0a, 0x12, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69,
	0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x1c, 0x66, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67,
	0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74,
	0x63, 0x67, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_fastcgi_proto_rawDescOnce sync.Once
	file_service_http_fastcgi_proto_rawDescData = file_service_http_fastcgi_proto_rawDesc
)

func file_service_http_fastcgi_proto_rawDescGZIP() []byte {
	file_service_http_fastcgi_proto_rawDescOnce.Do(func() {
		file_service_http_fastcgi_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_fastcgi_proto_rawDescData)
	})
	return file_service_http_fastcgi_proto_rawDescData
}

var file_service_http_fastcgi_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_http_fastcgi_proto_goTypes = []interface{}{
	(*CreateHTTPFastcgiRequest)(nil),             // 0: pb.CreateHTTPFastcgiRequest
	(*CreateHTTPFastcgiResponse)(nil),            // 1: pb.CreateHTTPFastcgiResponse
	(*UpdateHTTPFastcgiRequest)(nil),             // 2: pb.UpdateHTTPFastcgiRequest
	(*FindEnabledHTTPFastcgiRequest)(nil),        // 3: pb.FindEnabledHTTPFastcgiRequest
	(*FindEnabledHTTPFastcgiResponse)(nil),       // 4: pb.FindEnabledHTTPFastcgiResponse
	(*FindEnabledHTTPFastcgiConfigRequest)(nil),  // 5: pb.FindEnabledHTTPFastcgiConfigRequest
	(*FindEnabledHTTPFastcgiConfigResponse)(nil), // 6: pb.FindEnabledHTTPFastcgiConfigResponse
	(*HTTPFastcgi)(nil),                          // 7: pb.HTTPFastcgi
	(*RPCSuccess)(nil),                           // 8: pb.RPCSuccess
}
var file_service_http_fastcgi_proto_depIdxs = []int32{
	7, // 0: pb.FindEnabledHTTPFastcgiResponse.httpFastcgi:type_name -> pb.HTTPFastcgi
	0, // 1: pb.HTTPFastcgiService.createHTTPFastcgi:input_type -> pb.CreateHTTPFastcgiRequest
	2, // 2: pb.HTTPFastcgiService.updateHTTPFastcgi:input_type -> pb.UpdateHTTPFastcgiRequest
	3, // 3: pb.HTTPFastcgiService.findEnabledHTTPFastcgi:input_type -> pb.FindEnabledHTTPFastcgiRequest
	5, // 4: pb.HTTPFastcgiService.findEnabledHTTPFastcgiConfig:input_type -> pb.FindEnabledHTTPFastcgiConfigRequest
	1, // 5: pb.HTTPFastcgiService.createHTTPFastcgi:output_type -> pb.CreateHTTPFastcgiResponse
	8, // 6: pb.HTTPFastcgiService.updateHTTPFastcgi:output_type -> pb.RPCSuccess
	4, // 7: pb.HTTPFastcgiService.findEnabledHTTPFastcgi:output_type -> pb.FindEnabledHTTPFastcgiResponse
	6, // 8: pb.HTTPFastcgiService.findEnabledHTTPFastcgiConfig:output_type -> pb.FindEnabledHTTPFastcgiConfigResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_http_fastcgi_proto_init() }
func file_service_http_fastcgi_proto_init() {
	if File_service_http_fastcgi_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_http_fastcgi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_fastcgi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPFastcgiRequest); i {
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
		file_service_http_fastcgi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPFastcgiResponse); i {
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
		file_service_http_fastcgi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPFastcgiRequest); i {
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
		file_service_http_fastcgi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFastcgiRequest); i {
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
		file_service_http_fastcgi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFastcgiResponse); i {
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
		file_service_http_fastcgi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFastcgiConfigRequest); i {
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
		file_service_http_fastcgi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFastcgiConfigResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_http_fastcgi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_fastcgi_proto_goTypes,
		DependencyIndexes: file_service_http_fastcgi_proto_depIdxs,
		MessageInfos:      file_service_http_fastcgi_proto_msgTypes,
	}.Build()
	File_service_http_fastcgi_proto = out.File
	file_service_http_fastcgi_proto_rawDesc = nil
	file_service_http_fastcgi_proto_goTypes = nil
	file_service_http_fastcgi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPFastcgiServiceClient is the client API for HTTPFastcgiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPFastcgiServiceClient interface {
	// 创建Fastcgi
	CreateHTTPFastcgi(ctx context.Context, in *CreateHTTPFastcgiRequest, opts ...grpc.CallOption) (*CreateHTTPFastcgiResponse, error)
	// 修改Fastcgi
	UpdateHTTPFastcgi(ctx context.Context, in *UpdateHTTPFastcgiRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 获取Fastcgi详情
	FindEnabledHTTPFastcgi(ctx context.Context, in *FindEnabledHTTPFastcgiRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiResponse, error)
	// 获取Fastcgi配置
	FindEnabledHTTPFastcgiConfig(ctx context.Context, in *FindEnabledHTTPFastcgiConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiConfigResponse, error)
}

type hTTPFastcgiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPFastcgiServiceClient(cc grpc.ClientConnInterface) HTTPFastcgiServiceClient {
	return &hTTPFastcgiServiceClient{cc}
}

func (c *hTTPFastcgiServiceClient) CreateHTTPFastcgi(ctx context.Context, in *CreateHTTPFastcgiRequest, opts ...grpc.CallOption) (*CreateHTTPFastcgiResponse, error) {
	out := new(CreateHTTPFastcgiResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/createHTTPFastcgi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFastcgiServiceClient) UpdateHTTPFastcgi(ctx context.Context, in *UpdateHTTPFastcgiRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/updateHTTPFastcgi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFastcgiServiceClient) FindEnabledHTTPFastcgi(ctx context.Context, in *FindEnabledHTTPFastcgiRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiResponse, error) {
	out := new(FindEnabledHTTPFastcgiResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/findEnabledHTTPFastcgi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFastcgiServiceClient) FindEnabledHTTPFastcgiConfig(ctx context.Context, in *FindEnabledHTTPFastcgiConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiConfigResponse, error) {
	out := new(FindEnabledHTTPFastcgiConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/findEnabledHTTPFastcgiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPFastcgiServiceServer is the server API for HTTPFastcgiService service.
type HTTPFastcgiServiceServer interface {
	// 创建Fastcgi
	CreateHTTPFastcgi(context.Context, *CreateHTTPFastcgiRequest) (*CreateHTTPFastcgiResponse, error)
	// 修改Fastcgi
	UpdateHTTPFastcgi(context.Context, *UpdateHTTPFastcgiRequest) (*RPCSuccess, error)
	// 获取Fastcgi详情
	FindEnabledHTTPFastcgi(context.Context, *FindEnabledHTTPFastcgiRequest) (*FindEnabledHTTPFastcgiResponse, error)
	// 获取Fastcgi配置
	FindEnabledHTTPFastcgiConfig(context.Context, *FindEnabledHTTPFastcgiConfigRequest) (*FindEnabledHTTPFastcgiConfigResponse, error)
}

// UnimplementedHTTPFastcgiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPFastcgiServiceServer struct {
}

func (*UnimplementedHTTPFastcgiServiceServer) CreateHTTPFastcgi(context.Context, *CreateHTTPFastcgiRequest) (*CreateHTTPFastcgiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPFastcgi not implemented")
}
func (*UnimplementedHTTPFastcgiServiceServer) UpdateHTTPFastcgi(context.Context, *UpdateHTTPFastcgiRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPFastcgi not implemented")
}
func (*UnimplementedHTTPFastcgiServiceServer) FindEnabledHTTPFastcgi(context.Context, *FindEnabledHTTPFastcgiRequest) (*FindEnabledHTTPFastcgiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPFastcgi not implemented")
}
func (*UnimplementedHTTPFastcgiServiceServer) FindEnabledHTTPFastcgiConfig(context.Context, *FindEnabledHTTPFastcgiConfigRequest) (*FindEnabledHTTPFastcgiConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPFastcgiConfig not implemented")
}

func RegisterHTTPFastcgiServiceServer(s *grpc.Server, srv HTTPFastcgiServiceServer) {
	s.RegisterService(&_HTTPFastcgiService_serviceDesc, srv)
}

func _HTTPFastcgiService_CreateHTTPFastcgi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPFastcgiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).CreateHTTPFastcgi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/CreateHTTPFastcgi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).CreateHTTPFastcgi(ctx, req.(*CreateHTTPFastcgiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFastcgiService_UpdateHTTPFastcgi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPFastcgiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).UpdateHTTPFastcgi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/UpdateHTTPFastcgi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).UpdateHTTPFastcgi(ctx, req.(*UpdateHTTPFastcgiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFastcgiService_FindEnabledHTTPFastcgi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPFastcgiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/FindEnabledHTTPFastcgi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgi(ctx, req.(*FindEnabledHTTPFastcgiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFastcgiService_FindEnabledHTTPFastcgiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPFastcgiConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/FindEnabledHTTPFastcgiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgiConfig(ctx, req.(*FindEnabledHTTPFastcgiConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPFastcgiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPFastcgiService",
	HandlerType: (*HTTPFastcgiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPFastcgi",
			Handler:    _HTTPFastcgiService_CreateHTTPFastcgi_Handler,
		},
		{
			MethodName: "updateHTTPFastcgi",
			Handler:    _HTTPFastcgiService_UpdateHTTPFastcgi_Handler,
		},
		{
			MethodName: "findEnabledHTTPFastcgi",
			Handler:    _HTTPFastcgiService_FindEnabledHTTPFastcgi_Handler,
		},
		{
			MethodName: "findEnabledHTTPFastcgiConfig",
			Handler:    _HTTPFastcgiService_FindEnabledHTTPFastcgiConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_fastcgi.proto",
}
