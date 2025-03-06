// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/chunk_server_metadata_service.proto

package chunkserver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChunkServer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChunkServer) Reset() {
	*x = ChunkServer{}
	mi := &file_proto_chunk_server_metadata_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkServer) ProtoMessage() {}

func (x *ChunkServer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunk_server_metadata_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkServer.ProtoReflect.Descriptor instead.
func (*ChunkServer) Descriptor() ([]byte, []int) {
	return file_proto_chunk_server_metadata_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ChunkServerRegisterReq struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	MonitorAddress   string                 `protobuf:"bytes,1,opt,name=monitor_address,json=monitorAddress,proto3" json:"monitor_address,omitempty"`
	StreamingAddress string                 `protobuf:"bytes,2,opt,name=streaming_address,json=streamingAddress,proto3" json:"streaming_address,omitempty"`
	Space            int64                  `protobuf:"varint,3,opt,name=space,proto3" json:"space,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ChunkServerRegisterReq) Reset() {
	*x = ChunkServerRegisterReq{}
	mi := &file_proto_chunk_server_metadata_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkServerRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkServerRegisterReq) ProtoMessage() {}

func (x *ChunkServerRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunk_server_metadata_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkServerRegisterReq.ProtoReflect.Descriptor instead.
func (*ChunkServerRegisterReq) Descriptor() ([]byte, []int) {
	return file_proto_chunk_server_metadata_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkServerRegisterReq) GetMonitorAddress() string {
	if x != nil {
		return x.MonitorAddress
	}
	return ""
}

func (x *ChunkServerRegisterReq) GetStreamingAddress() string {
	if x != nil {
		return x.StreamingAddress
	}
	return ""
}

func (x *ChunkServerRegisterReq) GetSpace() int64 {
	if x != nil {
		return x.Space
	}
	return 0
}

type RegisterResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	mi := &file_proto_chunk_server_metadata_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chunk_server_metadata_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_proto_chunk_server_metadata_service_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_chunk_server_metadata_service_proto protoreflect.FileDescriptor

var file_proto_chunk_server_metadata_service_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x84, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0x68, 0x0a, 0x1a, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x28, 0x5a, 0x26,
	0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_chunk_server_metadata_service_proto_rawDescOnce sync.Once
	file_proto_chunk_server_metadata_service_proto_rawDescData []byte
)

func file_proto_chunk_server_metadata_service_proto_rawDescGZIP() []byte {
	file_proto_chunk_server_metadata_service_proto_rawDescOnce.Do(func() {
		file_proto_chunk_server_metadata_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_chunk_server_metadata_service_proto_rawDesc), len(file_proto_chunk_server_metadata_service_proto_rawDesc)))
	})
	return file_proto_chunk_server_metadata_service_proto_rawDescData
}

var file_proto_chunk_server_metadata_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_chunk_server_metadata_service_proto_goTypes = []any{
	(*ChunkServer)(nil),            // 0: chunkserver.ChunkServer
	(*ChunkServerRegisterReq)(nil), // 1: chunkserver.ChunkServerRegisterReq
	(*RegisterResp)(nil),           // 2: chunkserver.RegisterResp
}
var file_proto_chunk_server_metadata_service_proto_depIdxs = []int32{
	1, // 0: chunkserver.ChunkServerRegisterService.Register:input_type -> chunkserver.ChunkServerRegisterReq
	2, // 1: chunkserver.ChunkServerRegisterService.Register:output_type -> chunkserver.RegisterResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chunk_server_metadata_service_proto_init() }
func file_proto_chunk_server_metadata_service_proto_init() {
	if File_proto_chunk_server_metadata_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_chunk_server_metadata_service_proto_rawDesc), len(file_proto_chunk_server_metadata_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chunk_server_metadata_service_proto_goTypes,
		DependencyIndexes: file_proto_chunk_server_metadata_service_proto_depIdxs,
		MessageInfos:      file_proto_chunk_server_metadata_service_proto_msgTypes,
	}.Build()
	File_proto_chunk_server_metadata_service_proto = out.File
	file_proto_chunk_server_metadata_service_proto_goTypes = nil
	file_proto_chunk_server_metadata_service_proto_depIdxs = nil
}
