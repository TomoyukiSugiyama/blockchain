// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: proto/blockchain.proto

package proto

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

type TransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount        int32                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	mi := &file_proto_blockchain_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TransactionRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransactionRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransactionReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionReply) Reset() {
	*x = TransactionReply{}
	mi := &file_proto_blockchain_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionReply) ProtoMessage() {}

func (x *TransactionReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionReply.ProtoReflect.Descriptor instead.
func (*TransactionReply) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_blockchain_proto protoreflect.FileDescriptor

var file_proto_blockchain_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4c, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x3e, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x54, 0x72, 0x75, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_blockchain_proto_rawDescOnce sync.Once
	file_proto_blockchain_proto_rawDescData = file_proto_blockchain_proto_rawDesc
)

func file_proto_blockchain_proto_rawDescGZIP() []byte {
	file_proto_blockchain_proto_rawDescOnce.Do(func() {
		file_proto_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blockchain_proto_rawDescData)
	})
	return file_proto_blockchain_proto_rawDescData
}

var file_proto_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_blockchain_proto_goTypes = []any{
	(*TransactionRequest)(nil), // 0: TransactionRequest
	(*TransactionReply)(nil),   // 1: TransactionReply
}
var file_proto_blockchain_proto_depIdxs = []int32{
	0, // 0: Blockchain.ExecuteTrunsaction:input_type -> TransactionRequest
	1, // 1: Blockchain.ExecuteTrunsaction:output_type -> TransactionReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_blockchain_proto_init() }
func file_proto_blockchain_proto_init() {
	if File_proto_blockchain_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blockchain_proto_goTypes,
		DependencyIndexes: file_proto_blockchain_proto_depIdxs,
		MessageInfos:      file_proto_blockchain_proto_msgTypes,
	}.Build()
	File_proto_blockchain_proto = out.File
	file_proto_blockchain_proto_rawDesc = nil
	file_proto_blockchain_proto_goTypes = nil
	file_proto_blockchain_proto_depIdxs = nil
}
