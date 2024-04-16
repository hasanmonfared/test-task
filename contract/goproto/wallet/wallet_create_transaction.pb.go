// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contract/protobuf/wallet/wallet_create_transaction.proto

package wallet

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

type Type int32

const (
	Type_DEPOSIT    Type = 0
	Type_WITHDRAWAL Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "DEPOSIT",
		1: "WITHDRAWAL",
	}
	Type_value = map[string]int32{
		"DEPOSIT":    0,
		"WITHDRAWAL": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_protobuf_wallet_wallet_create_transaction_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_contract_protobuf_wallet_wallet_create_transaction_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescGZIP(), []int{0}
}

type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Type   Type    `protobuf:"varint,2,opt,name=type,proto3,enum=wallet.Type" json:"type,omitempty"`
	Amount float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTransactionRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateTransactionRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_DEPOSIT
}

func (x *CreateTransactionRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTransactionResponse) Reset() {
	*x = CreateTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionResponse) ProtoMessage() {}

func (x *CreateTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescGZIP(), []int{1}
}

var File_contract_protobuf_wallet_wallet_create_transaction_proto protoreflect.FileDescriptor

var file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x22, 0x68, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x23, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x10, 0x01, 0x32, 0x65,
	0x0a, 0x0f, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescOnce sync.Once
	file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescData = file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDesc
)

func file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescGZIP() []byte {
	file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescOnce.Do(func() {
		file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescData)
	})
	return file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDescData
}

var file_contract_protobuf_wallet_wallet_create_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contract_protobuf_wallet_wallet_create_transaction_proto_goTypes = []interface{}{
	(Type)(0),                         // 0: wallet.Type
	(*CreateTransactionRequest)(nil),  // 1: wallet.CreateTransactionRequest
	(*CreateTransactionResponse)(nil), // 2: wallet.CreateTransactionResponse
}
var file_contract_protobuf_wallet_wallet_create_transaction_proto_depIdxs = []int32{
	0, // 0: wallet.CreateTransactionRequest.type:type_name -> wallet.Type
	1, // 1: wallet.PresenceService.GetPresence:input_type -> wallet.CreateTransactionRequest
	2, // 2: wallet.PresenceService.GetPresence:output_type -> wallet.CreateTransactionResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_protobuf_wallet_wallet_create_transaction_proto_init() }
func file_contract_protobuf_wallet_wallet_create_transaction_proto_init() {
	if File_contract_protobuf_wallet_wallet_create_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionRequest); i {
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
		file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionResponse); i {
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
			RawDescriptor: file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_protobuf_wallet_wallet_create_transaction_proto_goTypes,
		DependencyIndexes: file_contract_protobuf_wallet_wallet_create_transaction_proto_depIdxs,
		EnumInfos:         file_contract_protobuf_wallet_wallet_create_transaction_proto_enumTypes,
		MessageInfos:      file_contract_protobuf_wallet_wallet_create_transaction_proto_msgTypes,
	}.Build()
	File_contract_protobuf_wallet_wallet_create_transaction_proto = out.File
	file_contract_protobuf_wallet_wallet_create_transaction_proto_rawDesc = nil
	file_contract_protobuf_wallet_wallet_create_transaction_proto_goTypes = nil
	file_contract_protobuf_wallet_wallet_create_transaction_proto_depIdxs = nil
}
