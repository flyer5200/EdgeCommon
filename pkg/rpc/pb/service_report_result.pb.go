// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_report_result.proto

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

// 计算监控结果数量
type CountAllReportResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportNodeId int64  `protobuf:"varint,1,opt,name=reportNodeId,proto3" json:"reportNodeId,omitempty"`
	OkState      int32  `protobuf:"varint,2,opt,name=okState,proto3" json:"okState,omitempty"`
	Level        string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *CountAllReportResultsRequest) Reset() {
	*x = CountAllReportResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_report_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllReportResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllReportResultsRequest) ProtoMessage() {}

func (x *CountAllReportResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_report_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllReportResultsRequest.ProtoReflect.Descriptor instead.
func (*CountAllReportResultsRequest) Descriptor() ([]byte, []int) {
	return file_service_report_result_proto_rawDescGZIP(), []int{0}
}

func (x *CountAllReportResultsRequest) GetReportNodeId() int64 {
	if x != nil {
		return x.ReportNodeId
	}
	return 0
}

func (x *CountAllReportResultsRequest) GetOkState() int32 {
	if x != nil {
		return x.OkState
	}
	return 0
}

func (x *CountAllReportResultsRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

// 列出单页监控结果
type ListReportResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportNodeId int64  `protobuf:"varint,1,opt,name=reportNodeId,proto3" json:"reportNodeId,omitempty"`
	OkState      int32  `protobuf:"varint,2,opt,name=okState,proto3" json:"okState,omitempty"`
	Level        string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Offset       int64  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size         int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListReportResultsRequest) Reset() {
	*x = ListReportResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_report_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReportResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReportResultsRequest) ProtoMessage() {}

func (x *ListReportResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_report_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReportResultsRequest.ProtoReflect.Descriptor instead.
func (*ListReportResultsRequest) Descriptor() ([]byte, []int) {
	return file_service_report_result_proto_rawDescGZIP(), []int{1}
}

func (x *ListReportResultsRequest) GetReportNodeId() int64 {
	if x != nil {
		return x.ReportNodeId
	}
	return 0
}

func (x *ListReportResultsRequest) GetOkState() int32 {
	if x != nil {
		return x.OkState
	}
	return 0
}

func (x *ListReportResultsRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *ListReportResultsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListReportResultsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListReportResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportResults []*ReportResult `protobuf:"bytes,1,rep,name=reportResults,proto3" json:"reportResults,omitempty"`
}

func (x *ListReportResultsResponse) Reset() {
	*x = ListReportResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_report_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReportResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReportResultsResponse) ProtoMessage() {}

func (x *ListReportResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_report_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReportResultsResponse.ProtoReflect.Descriptor instead.
func (*ListReportResultsResponse) Descriptor() ([]byte, []int) {
	return file_service_report_result_proto_rawDescGZIP(), []int{2}
}

func (x *ListReportResultsResponse) GetReportResults() []*ReportResult {
	if x != nil {
		return x.ReportResults
	}
	return nil
}

// 上传报告结果
type UpdateReportResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportResults []*ReportResult `protobuf:"bytes,1,rep,name=reportResults,proto3" json:"reportResults,omitempty"`
}

func (x *UpdateReportResultsRequest) Reset() {
	*x = UpdateReportResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_report_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReportResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReportResultsRequest) ProtoMessage() {}

func (x *UpdateReportResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_report_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReportResultsRequest.ProtoReflect.Descriptor instead.
func (*UpdateReportResultsRequest) Descriptor() ([]byte, []int) {
	return file_service_report_result_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReportResultsRequest) GetReportResults() []*ReportResult {
	if x != nil {
		return x.ReportResults
	}
	return nil
}

// 查询某个对象的监控结果
type FindAllReportResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	TargetId int64  `protobuf:"varint,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
}

func (x *FindAllReportResultsRequest) Reset() {
	*x = FindAllReportResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_report_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllReportResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllReportResultsRequest) ProtoMessage() {}

func (x *FindAllReportResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_report_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllReportResultsRequest.ProtoReflect.Descriptor instead.
func (*FindAllReportResultsRequest) Descriptor() ([]byte, []int) {
	return file_service_report_result_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllReportResultsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FindAllReportResultsRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type FindAllReportResultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportResults []*ReportResult `protobuf:"bytes,1,rep,name=reportResults,proto3" json:"reportResults,omitempty"`
}

func (x *FindAllReportResultsResponse) Reset() {
	*x = FindAllReportResultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_report_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllReportResultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllReportResultsResponse) ProtoMessage() {}

func (x *FindAllReportResultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_report_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllReportResultsResponse.ProtoReflect.Descriptor instead.
func (*FindAllReportResultsResponse) Descriptor() ([]byte, []int) {
	return file_service_report_result_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllReportResultsResponse) GetReportResults() []*ReportResult {
	if x != nil {
		return x.ReportResults
	}
	return nil
}

var File_service_report_result_proto protoreflect.FileDescriptor

var file_service_report_result_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72,
	0x0a, 0x1c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x53, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x4d, 0x0a, 0x1b, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1c, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x32, 0xda, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x13,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x59, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_report_result_proto_rawDescOnce sync.Once
	file_service_report_result_proto_rawDescData = file_service_report_result_proto_rawDesc
)

func file_service_report_result_proto_rawDescGZIP() []byte {
	file_service_report_result_proto_rawDescOnce.Do(func() {
		file_service_report_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_report_result_proto_rawDescData)
	})
	return file_service_report_result_proto_rawDescData
}

var file_service_report_result_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_report_result_proto_goTypes = []interface{}{
	(*CountAllReportResultsRequest)(nil), // 0: pb.CountAllReportResultsRequest
	(*ListReportResultsRequest)(nil),     // 1: pb.ListReportResultsRequest
	(*ListReportResultsResponse)(nil),    // 2: pb.ListReportResultsResponse
	(*UpdateReportResultsRequest)(nil),   // 3: pb.UpdateReportResultsRequest
	(*FindAllReportResultsRequest)(nil),  // 4: pb.FindAllReportResultsRequest
	(*FindAllReportResultsResponse)(nil), // 5: pb.FindAllReportResultsResponse
	(*ReportResult)(nil),                 // 6: pb.ReportResult
	(*RPCCountResponse)(nil),             // 7: pb.RPCCountResponse
	(*RPCSuccess)(nil),                   // 8: pb.RPCSuccess
}
var file_service_report_result_proto_depIdxs = []int32{
	6, // 0: pb.ListReportResultsResponse.reportResults:type_name -> pb.ReportResult
	6, // 1: pb.UpdateReportResultsRequest.reportResults:type_name -> pb.ReportResult
	6, // 2: pb.FindAllReportResultsResponse.reportResults:type_name -> pb.ReportResult
	0, // 3: pb.ReportResultService.countAllReportResults:input_type -> pb.CountAllReportResultsRequest
	1, // 4: pb.ReportResultService.listReportResults:input_type -> pb.ListReportResultsRequest
	3, // 5: pb.ReportResultService.updateReportResults:input_type -> pb.UpdateReportResultsRequest
	4, // 6: pb.ReportResultService.findAllReportResults:input_type -> pb.FindAllReportResultsRequest
	7, // 7: pb.ReportResultService.countAllReportResults:output_type -> pb.RPCCountResponse
	2, // 8: pb.ReportResultService.listReportResults:output_type -> pb.ListReportResultsResponse
	8, // 9: pb.ReportResultService.updateReportResults:output_type -> pb.RPCSuccess
	5, // 10: pb.ReportResultService.findAllReportResults:output_type -> pb.FindAllReportResultsResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_report_result_proto_init() }
func file_service_report_result_proto_init() {
	if File_service_report_result_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_report_result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_report_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllReportResultsRequest); i {
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
		file_service_report_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReportResultsRequest); i {
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
		file_service_report_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReportResultsResponse); i {
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
		file_service_report_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReportResultsRequest); i {
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
		file_service_report_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllReportResultsRequest); i {
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
		file_service_report_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllReportResultsResponse); i {
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
			RawDescriptor: file_service_report_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_report_result_proto_goTypes,
		DependencyIndexes: file_service_report_result_proto_depIdxs,
		MessageInfos:      file_service_report_result_proto_msgTypes,
	}.Build()
	File_service_report_result_proto = out.File
	file_service_report_result_proto_rawDesc = nil
	file_service_report_result_proto_goTypes = nil
	file_service_report_result_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportResultServiceClient is the client API for ReportResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportResultServiceClient interface {
	// 计算监控结果数量
	CountAllReportResults(ctx context.Context, in *CountAllReportResultsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页监控结果
	ListReportResults(ctx context.Context, in *ListReportResultsRequest, opts ...grpc.CallOption) (*ListReportResultsResponse, error)
	// 上传报告结果
	UpdateReportResults(ctx context.Context, in *UpdateReportResultsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询某个对象的监控结果
	FindAllReportResults(ctx context.Context, in *FindAllReportResultsRequest, opts ...grpc.CallOption) (*FindAllReportResultsResponse, error)
}

type reportResultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportResultServiceClient(cc grpc.ClientConnInterface) ReportResultServiceClient {
	return &reportResultServiceClient{cc}
}

func (c *reportResultServiceClient) CountAllReportResults(ctx context.Context, in *CountAllReportResultsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.ReportResultService/countAllReportResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportResultServiceClient) ListReportResults(ctx context.Context, in *ListReportResultsRequest, opts ...grpc.CallOption) (*ListReportResultsResponse, error) {
	out := new(ListReportResultsResponse)
	err := c.cc.Invoke(ctx, "/pb.ReportResultService/listReportResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportResultServiceClient) UpdateReportResults(ctx context.Context, in *UpdateReportResultsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.ReportResultService/updateReportResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportResultServiceClient) FindAllReportResults(ctx context.Context, in *FindAllReportResultsRequest, opts ...grpc.CallOption) (*FindAllReportResultsResponse, error) {
	out := new(FindAllReportResultsResponse)
	err := c.cc.Invoke(ctx, "/pb.ReportResultService/findAllReportResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportResultServiceServer is the server API for ReportResultService service.
type ReportResultServiceServer interface {
	// 计算监控结果数量
	CountAllReportResults(context.Context, *CountAllReportResultsRequest) (*RPCCountResponse, error)
	// 列出单页监控结果
	ListReportResults(context.Context, *ListReportResultsRequest) (*ListReportResultsResponse, error)
	// 上传报告结果
	UpdateReportResults(context.Context, *UpdateReportResultsRequest) (*RPCSuccess, error)
	// 查询某个对象的监控结果
	FindAllReportResults(context.Context, *FindAllReportResultsRequest) (*FindAllReportResultsResponse, error)
}

// UnimplementedReportResultServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReportResultServiceServer struct {
}

func (*UnimplementedReportResultServiceServer) CountAllReportResults(context.Context, *CountAllReportResultsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllReportResults not implemented")
}
func (*UnimplementedReportResultServiceServer) ListReportResults(context.Context, *ListReportResultsRequest) (*ListReportResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReportResults not implemented")
}
func (*UnimplementedReportResultServiceServer) UpdateReportResults(context.Context, *UpdateReportResultsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReportResults not implemented")
}
func (*UnimplementedReportResultServiceServer) FindAllReportResults(context.Context, *FindAllReportResultsRequest) (*FindAllReportResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllReportResults not implemented")
}

func RegisterReportResultServiceServer(s *grpc.Server, srv ReportResultServiceServer) {
	s.RegisterService(&_ReportResultService_serviceDesc, srv)
}

func _ReportResultService_CountAllReportResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllReportResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportResultServiceServer).CountAllReportResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReportResultService/CountAllReportResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportResultServiceServer).CountAllReportResults(ctx, req.(*CountAllReportResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportResultService_ListReportResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportResultServiceServer).ListReportResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReportResultService/ListReportResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportResultServiceServer).ListReportResults(ctx, req.(*ListReportResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportResultService_UpdateReportResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReportResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportResultServiceServer).UpdateReportResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReportResultService/UpdateReportResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportResultServiceServer).UpdateReportResults(ctx, req.(*UpdateReportResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportResultService_FindAllReportResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReportResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportResultServiceServer).FindAllReportResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReportResultService/FindAllReportResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportResultServiceServer).FindAllReportResults(ctx, req.(*FindAllReportResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportResultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReportResultService",
	HandlerType: (*ReportResultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "countAllReportResults",
			Handler:    _ReportResultService_CountAllReportResults_Handler,
		},
		{
			MethodName: "listReportResults",
			Handler:    _ReportResultService_ListReportResults_Handler,
		},
		{
			MethodName: "updateReportResults",
			Handler:    _ReportResultService_UpdateReportResults_Handler,
		},
		{
			MethodName: "findAllReportResults",
			Handler:    _ReportResultService_FindAllReportResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_report_result.proto",
}
