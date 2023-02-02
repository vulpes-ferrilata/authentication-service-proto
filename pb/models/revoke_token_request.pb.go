// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.3
// source: protobuf/models/revoke_token_request.proto

package models

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

type RevokeTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *RevokeTokenRequest) Reset() {
	*x = RevokeTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_models_revoke_token_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTokenRequest) ProtoMessage() {}

func (x *RevokeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_models_revoke_token_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTokenRequest.ProtoReflect.Descriptor instead.
func (*RevokeTokenRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_models_revoke_token_request_proto_rawDescGZIP(), []int{0}
}

func (x *RevokeTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_protobuf_models_revoke_token_request_proto protoreflect.FileDescriptor

var file_protobuf_models_revoke_token_request_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x75, 0x6c,
	0x70, 0x65, 0x73, 0x2d, 0x66, 0x65, 0x72, 0x72, 0x69, 0x6c, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_models_revoke_token_request_proto_rawDescOnce sync.Once
	file_protobuf_models_revoke_token_request_proto_rawDescData = file_protobuf_models_revoke_token_request_proto_rawDesc
)

func file_protobuf_models_revoke_token_request_proto_rawDescGZIP() []byte {
	file_protobuf_models_revoke_token_request_proto_rawDescOnce.Do(func() {
		file_protobuf_models_revoke_token_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_models_revoke_token_request_proto_rawDescData)
	})
	return file_protobuf_models_revoke_token_request_proto_rawDescData
}

var file_protobuf_models_revoke_token_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_models_revoke_token_request_proto_goTypes = []interface{}{
	(*RevokeTokenRequest)(nil), // 0: models.RevokeTokenRequest
}
var file_protobuf_models_revoke_token_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_models_revoke_token_request_proto_init() }
func file_protobuf_models_revoke_token_request_proto_init() {
	if File_protobuf_models_revoke_token_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_models_revoke_token_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeTokenRequest); i {
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
			RawDescriptor: file_protobuf_models_revoke_token_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_models_revoke_token_request_proto_goTypes,
		DependencyIndexes: file_protobuf_models_revoke_token_request_proto_depIdxs,
		MessageInfos:      file_protobuf_models_revoke_token_request_proto_msgTypes,
	}.Build()
	File_protobuf_models_revoke_token_request_proto = out.File
	file_protobuf_models_revoke_token_request_proto_rawDesc = nil
	file_protobuf_models_revoke_token_request_proto_goTypes = nil
	file_protobuf_models_revoke_token_request_proto_depIdxs = nil
}
