// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.19.1
// source: api/address_api.proto

package address_api

import (
	proto "github.com/golang/protobuf/proto"
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

type GetUserAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserAddressRequest) Reset() {
	*x = GetUserAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_address_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAddressRequest) ProtoMessage() {}

func (x *GetUserAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_address_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAddressRequest.ProtoReflect.Descriptor instead.
func (*GetUserAddressRequest) Descriptor() ([]byte, []int) {
	return file_api_address_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserAddressRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetUserAddressResponse) Reset() {
	*x = GetUserAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_address_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAddressResponse) ProtoMessage() {}

func (x *GetUserAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_address_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAddressResponse.ProtoReflect.Descriptor instead.
func (*GetUserAddressResponse) Descriptor() ([]byte, []int) {
	return file_api_address_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserAddressResponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	ZipCode string `protobuf:"bytes,3,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_address_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_api_address_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_api_address_api_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

var File_api_address_api_proto protoreflect.FileDescriptor

var file_api_address_api_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x48, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x53, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x6b, 0x66, 0x6d, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_address_api_proto_rawDescOnce sync.Once
	file_api_address_api_proto_rawDescData = file_api_address_api_proto_rawDesc
)

func file_api_address_api_proto_rawDescGZIP() []byte {
	file_api_address_api_proto_rawDescOnce.Do(func() {
		file_api_address_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_address_api_proto_rawDescData)
	})
	return file_api_address_api_proto_rawDescData
}

var file_api_address_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_address_api_proto_goTypes = []interface{}{
	(*GetUserAddressRequest)(nil),  // 0: GetUserAddressRequest
	(*GetUserAddressResponse)(nil), // 1: GetUserAddressResponse
	(*Address)(nil),                // 2: Address
}
var file_api_address_api_proto_depIdxs = []int32{
	2, // 0: GetUserAddressResponse.address:type_name -> Address
	0, // 1: AddressService.GetUserAddress:input_type -> GetUserAddressRequest
	1, // 2: AddressService.GetUserAddress:output_type -> GetUserAddressResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_address_api_proto_init() }
func file_api_address_api_proto_init() {
	if File_api_address_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_address_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAddressRequest); i {
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
		file_api_address_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAddressResponse); i {
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
		file_api_address_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_api_address_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_address_api_proto_goTypes,
		DependencyIndexes: file_api_address_api_proto_depIdxs,
		MessageInfos:      file_api_address_api_proto_msgTypes,
	}.Build()
	File_api_address_api_proto = out.File
	file_api_address_api_proto_rawDesc = nil
	file_api_address_api_proto_goTypes = nil
	file_api_address_api_proto_depIdxs = nil
}
