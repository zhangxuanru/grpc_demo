// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.0
// source: module/Order.Module.proto

package pbfiles

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//主订单
type OrderMain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int32                `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo    string               `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId     int32                `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderMoney float32              `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`
	OrderTime  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`
}

func (x *OrderMain) Reset() {
	*x = OrderMain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_Order_Module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderMain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderMain) ProtoMessage() {}

func (x *OrderMain) ProtoReflect() protoreflect.Message {
	mi := &file_module_Order_Module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderMain.ProtoReflect.Descriptor instead.
func (*OrderMain) Descriptor() ([]byte, []int) {
	return file_module_Order_Module_proto_rawDescGZIP(), []int{0}
}

func (x *OrderMain) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderMain) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *OrderMain) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderMain) GetOrderMoney() float32 {
	if x != nil {
		return x.OrderMoney
	}
	return 0
}

func (x *OrderMain) GetOrderTime() *timestamp.Timestamp {
	if x != nil {
		return x.OrderTime
	}
	return nil
}

var File_module_Order_Module_proto protoreflect.FileDescriptor

var file_module_Order_Module_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x61, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x3b, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_module_Order_Module_proto_rawDescOnce sync.Once
	file_module_Order_Module_proto_rawDescData = file_module_Order_Module_proto_rawDesc
)

func file_module_Order_Module_proto_rawDescGZIP() []byte {
	file_module_Order_Module_proto_rawDescOnce.Do(func() {
		file_module_Order_Module_proto_rawDescData = protoimpl.X.CompressGZIP(file_module_Order_Module_proto_rawDescData)
	})
	return file_module_Order_Module_proto_rawDescData
}

var file_module_Order_Module_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_module_Order_Module_proto_goTypes = []interface{}{
	(*OrderMain)(nil),           // 0: pbfiles.OrderMain
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_module_Order_Module_proto_depIdxs = []int32{
	1, // 0: pbfiles.OrderMain.order_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_module_Order_Module_proto_init() }
func file_module_Order_Module_proto_init() {
	if File_module_Order_Module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_module_Order_Module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderMain); i {
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
			RawDescriptor: file_module_Order_Module_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_module_Order_Module_proto_goTypes,
		DependencyIndexes: file_module_Order_Module_proto_depIdxs,
		MessageInfos:      file_module_Order_Module_proto_msgTypes,
	}.Build()
	File_module_Order_Module_proto = out.File
	file_module_Order_Module_proto_rawDesc = nil
	file_module_Order_Module_proto_goTypes = nil
	file_module_Order_Module_proto_depIdxs = nil
}
