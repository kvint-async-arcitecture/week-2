// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: auth/api/rpc/service.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_api_rpc_service_proto protoreflect.FileDescriptor

var file_auth_api_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70,
	0x63, 0x1a, 0x1b, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x48,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_api_rpc_service_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),  // 0: rpc.RegisterRequest
	(*RegisterResponse)(nil), // 1: rpc.RegisterResponse
}
var file_auth_api_rpc_service_proto_depIdxs = []int32{
	0, // 0: rpc.AuthService.Register:input_type -> rpc.RegisterRequest
	1, // 1: rpc.AuthService.Register:output_type -> rpc.RegisterResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_api_rpc_service_proto_init() }
func file_auth_api_rpc_service_proto_init() {
	if File_auth_api_rpc_service_proto != nil {
		return
	}
	file_auth_api_rpc_register_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_api_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_api_rpc_service_proto_goTypes,
		DependencyIndexes: file_auth_api_rpc_service_proto_depIdxs,
	}.Build()
	File_auth_api_rpc_service_proto = out.File
	file_auth_api_rpc_service_proto_rawDesc = nil
	file_auth_api_rpc_service_proto_goTypes = nil
	file_auth_api_rpc_service_proto_depIdxs = nil
}