// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.0
// source: Prod.proto

package pbfiles

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

//枚举
type ProdArea int32

const (
	ProdArea_A ProdArea = 0
	ProdArea_B ProdArea = 1
	ProdArea_C ProdArea = 2
)

// Enum value maps for ProdArea.
var (
	ProdArea_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	ProdArea_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x ProdArea) Enum() *ProdArea {
	p := new(ProdArea)
	*p = x
	return p
}

func (x ProdArea) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProdArea) Descriptor() protoreflect.EnumDescriptor {
	return file_Prod_proto_enumTypes[0].Descriptor()
}

func (ProdArea) Type() protoreflect.EnumType {
	return &file_Prod_proto_enumTypes[0]
}

func (x ProdArea) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProdArea.Descriptor instead.
func (ProdArea) EnumDescriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{0}
}

type ProdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	Area   ProdArea `protobuf:"varint,2,opt,name=area,proto3,enum=pbfiles.ProdArea" json:"area,omitempty"`
}

func (x *ProdRequest) Reset() {
	*x = ProdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRequest) ProtoMessage() {}

func (x *ProdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest.ProtoReflect.Descriptor instead.
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{0}
}

func (x *ProdRequest) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *ProdRequest) GetArea() ProdArea {
	if x != nil {
		return x.Area
	}
	return ProdArea_A
}

type ProdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdStock int32 `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
}

func (x *ProdResponse) Reset() {
	*x = ProdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponse) ProtoMessage() {}

func (x *ProdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse.ProtoReflect.Descriptor instead.
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{1}
}

func (x *ProdResponse) GetProdStock() int32 {
	if x != nil {
		return x.ProdStock
	}
	return 0
}

//---
type ProdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ProdListRequest) Reset() {
	*x = ProdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdListRequest) ProtoMessage() {}

func (x *ProdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdListRequest.ProtoReflect.Descriptor instead.
func (*ProdListRequest) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{2}
}

func (x *ProdListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ProdList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId     int32   `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdName   string  `protobuf:"bytes,2,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	IsSubsidy  bool    `protobuf:"varint,3,opt,name=is_subsidy,json=isSubsidy,proto3" json:"is_subsidy,omitempty"`
	IsIntegral int32   `protobuf:"varint,4,opt,name=is_integral,json=isIntegral,proto3" json:"is_integral,omitempty"`
	Price      float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProdList) Reset() {
	*x = ProdList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdList) ProtoMessage() {}

func (x *ProdList) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdList.ProtoReflect.Descriptor instead.
func (*ProdList) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{3}
}

func (x *ProdList) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *ProdList) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *ProdList) GetIsSubsidy() bool {
	if x != nil {
		return x.IsSubsidy
	}
	return false
}

func (x *ProdList) GetIsIntegral() int32 {
	if x != nil {
		return x.IsIntegral
	}
	return 0
}

func (x *ProdList) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

//返回数组
type ProdListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdList []*ProdList `protobuf:"bytes,1,rep,name=prodList,proto3" json:"prodList,omitempty"`
}

func (x *ProdListResponse) Reset() {
	*x = ProdListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Prod_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdListResponse) ProtoMessage() {}

func (x *ProdListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Prod_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdListResponse.ProtoReflect.Descriptor instead.
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return file_Prod_proto_rawDescGZIP(), []int{4}
}

func (x *ProdListResponse) GetProdList() []*ProdList {
	if x != nil {
		return x.ProdList
	}
	return nil
}

var File_Prod_proto protoreflect.FileDescriptor

var file_Prod_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x50, 0x72, 0x6f, 0x64,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a,
	0x0b, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70,
	0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0x2d, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x25, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73,
	0x69, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x62,
	0x73, 0x69, 0x64, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x73, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x10, 0x50,
	0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x1f,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x41, 0x72, 0x65, 0x61, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10,
	0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x02, 0x32,
	0x8f, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x60, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x63, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f,
	0x7b, 0x73, 0x69, 0x7a, 0x65, 0x7d, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Prod_proto_rawDescOnce sync.Once
	file_Prod_proto_rawDescData = file_Prod_proto_rawDesc
)

func file_Prod_proto_rawDescGZIP() []byte {
	file_Prod_proto_rawDescOnce.Do(func() {
		file_Prod_proto_rawDescData = protoimpl.X.CompressGZIP(file_Prod_proto_rawDescData)
	})
	return file_Prod_proto_rawDescData
}

var file_Prod_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Prod_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Prod_proto_goTypes = []interface{}{
	(ProdArea)(0),            // 0: pbfiles.ProdArea
	(*ProdRequest)(nil),      // 1: pbfiles.ProdRequest
	(*ProdResponse)(nil),     // 2: pbfiles.ProdResponse
	(*ProdListRequest)(nil),  // 3: pbfiles.ProdListRequest
	(*ProdList)(nil),         // 4: pbfiles.ProdList
	(*ProdListResponse)(nil), // 5: pbfiles.ProdListResponse
	(*ProdModel)(nil),        // 6: pbfiles.ProdModel
}
var file_Prod_proto_depIdxs = []int32{
	0, // 0: pbfiles.ProdRequest.area:type_name -> pbfiles.ProdArea
	4, // 1: pbfiles.ProdListResponse.prodList:type_name -> pbfiles.ProdList
	1, // 2: pbfiles.ProdService.GetProdStock:input_type -> pbfiles.ProdRequest
	3, // 3: pbfiles.ProdService.GetProdList:input_type -> pbfiles.ProdListRequest
	1, // 4: pbfiles.ProdService.GetProdInfo:input_type -> pbfiles.ProdRequest
	2, // 5: pbfiles.ProdService.GetProdStock:output_type -> pbfiles.ProdResponse
	5, // 6: pbfiles.ProdService.GetProdList:output_type -> pbfiles.ProdListResponse
	6, // 7: pbfiles.ProdService.GetProdInfo:output_type -> pbfiles.ProdModel
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Prod_proto_init() }
func file_Prod_proto_init() {
	if File_Prod_proto != nil {
		return
	}
	file_module_Prod_Module_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Prod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRequest); i {
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
		file_Prod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponse); i {
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
		file_Prod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdListRequest); i {
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
		file_Prod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdList); i {
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
		file_Prod_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdListResponse); i {
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
			RawDescriptor: file_Prod_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Prod_proto_goTypes,
		DependencyIndexes: file_Prod_proto_depIdxs,
		EnumInfos:         file_Prod_proto_enumTypes,
		MessageInfos:      file_Prod_proto_msgTypes,
	}.Build()
	File_Prod_proto = out.File
	file_Prod_proto_rawDesc = nil
	file_Prod_proto_goTypes = nil
	file_Prod_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
	GetProdList(ctx context.Context, in *ProdListRequest, opts ...grpc.CallOption) (*ProdListResponse, error)
	GetProdInfo(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdModel, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/pbfiles.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdList(ctx context.Context, in *ProdListRequest, opts ...grpc.CallOption) (*ProdListResponse, error) {
	out := new(ProdListResponse)
	err := c.cc.Invoke(ctx, "/pbfiles.ProdService/GetProdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdInfo(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdModel, error) {
	out := new(ProdModel)
	err := c.cc.Invoke(ctx, "/pbfiles.ProdService/GetProdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
	GetProdList(context.Context, *ProdListRequest) (*ProdListResponse, error)
	GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}
func (*UnimplementedProdServiceServer) GetProdList(context.Context, *ProdListRequest) (*ProdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdList not implemented")
}
func (*UnimplementedProdServiceServer) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdInfo not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfiles.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfiles.ProdService/GetProdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdList(ctx, req.(*ProdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfiles.ProdService/GetProdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdInfo(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbfiles.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
		{
			MethodName: "GetProdList",
			Handler:    _ProdService_GetProdList_Handler,
		},
		{
			MethodName: "GetProdInfo",
			Handler:    _ProdService_GetProdInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Prod.proto",
}
