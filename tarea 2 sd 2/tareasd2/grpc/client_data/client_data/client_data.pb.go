//
//desde el directorio "grpc", ejecutar el comando:
//protoc --go_out=plugins=grpc:client_data client_data/client_data.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: client_data/client_data.proto

package client_data

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

type UploadResCode int32

const (
	UploadResCode_Unknown UploadResCode = 0
	UploadResCode_Ok      UploadResCode = 1
	UploadResCode_Failed  UploadResCode = 2
)

// Enum value maps for UploadResCode.
var (
	UploadResCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	UploadResCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x UploadResCode) Enum() *UploadResCode {
	p := new(UploadResCode)
	*p = x
	return p
}

func (x UploadResCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadResCode) Descriptor() protoreflect.EnumDescriptor {
	return file_client_data_client_data_proto_enumTypes[0].Descriptor()
}

func (UploadResCode) Type() protoreflect.EnumType {
	return &file_client_data_client_data_proto_enumTypes[0]
}

func (x UploadResCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadResCode.Descriptor instead.
func (UploadResCode) EnumDescriptor() ([]byte, []int) {
	return file_client_data_client_data_proto_rawDescGZIP(), []int{0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	ChunkId int64  `protobuf:"varint,2,opt,name=ChunkId,proto3" json:"ChunkId,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_data_client_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_client_data_client_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_client_data_client_data_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetChunkId() int64 {
	if x != nil {
		return x.ChunkId
	}
	return 0
}

type UploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//	*UploadReq_FileName
	//	*UploadReq_DataChunk
	Req isUploadReq_Req `protobuf_oneof:"Req"`
}

func (x *UploadReq) Reset() {
	*x = UploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_data_client_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReq) ProtoMessage() {}

func (x *UploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_data_client_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReq.ProtoReflect.Descriptor instead.
func (*UploadReq) Descriptor() ([]byte, []int) {
	return file_client_data_client_data_proto_rawDescGZIP(), []int{1}
}

func (m *UploadReq) GetReq() isUploadReq_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *UploadReq) GetFileName() string {
	if x, ok := x.GetReq().(*UploadReq_FileName); ok {
		return x.FileName
	}
	return ""
}

func (x *UploadReq) GetDataChunk() *Chunk {
	if x, ok := x.GetReq().(*UploadReq_DataChunk); ok {
		return x.DataChunk
	}
	return nil
}

type isUploadReq_Req interface {
	isUploadReq_Req()
}

type UploadReq_FileName struct {
	FileName string `protobuf:"bytes,1,opt,name=FileName,proto3,oneof"`
}

type UploadReq_DataChunk struct {
	DataChunk *Chunk `protobuf:"bytes,2,opt,name=DataChunk,proto3,oneof"`
}

func (*UploadReq_FileName) isUploadReq_Req() {}

func (*UploadReq_DataChunk) isUploadReq_Req() {}

type UploadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResCode UploadResCode `protobuf:"varint,1,opt,name=ResCode,proto3,enum=client_data.UploadResCode" json:"ResCode,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *UploadRes) Reset() {
	*x = UploadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_data_client_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRes) ProtoMessage() {}

func (x *UploadRes) ProtoReflect() protoreflect.Message {
	mi := &file_client_data_client_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRes.ProtoReflect.Descriptor instead.
func (*UploadRes) Descriptor() ([]byte, []int) {
	return file_client_data_client_data_proto_rawDescGZIP(), []int{2}
}

func (x *UploadRes) GetResCode() UploadResCode {
	if x != nil {
		return x.ResCode
	}
	return UploadResCode_Unknown
}

func (x *UploadRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DownloadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//	*DownloadReq_FileName
	//	*DownloadReq_ChunkId
	Req isDownloadReq_Req `protobuf_oneof:"Req"`
}

func (x *DownloadReq) Reset() {
	*x = DownloadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_data_client_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReq) ProtoMessage() {}

func (x *DownloadReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_data_client_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReq.ProtoReflect.Descriptor instead.
func (*DownloadReq) Descriptor() ([]byte, []int) {
	return file_client_data_client_data_proto_rawDescGZIP(), []int{3}
}

func (m *DownloadReq) GetReq() isDownloadReq_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *DownloadReq) GetFileName() string {
	if x, ok := x.GetReq().(*DownloadReq_FileName); ok {
		return x.FileName
	}
	return ""
}

func (x *DownloadReq) GetChunkId() int64 {
	if x, ok := x.GetReq().(*DownloadReq_ChunkId); ok {
		return x.ChunkId
	}
	return 0
}

type isDownloadReq_Req interface {
	isDownloadReq_Req()
}

type DownloadReq_FileName struct {
	FileName string `protobuf:"bytes,1,opt,name=FileName,proto3,oneof"`
}

type DownloadReq_ChunkId struct {
	ChunkId int64 `protobuf:"varint,2,opt,name=ChunkId,proto3,oneof"`
}

func (*DownloadReq_FileName) isDownloadReq_Req() {}

func (*DownloadReq_ChunkId) isDownloadReq_Req() {}

var File_client_data_client_data_proto protoreflect.FileDescriptor

var file_client_data_client_data_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x05,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x09, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x22,
	0x5b, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07,
	0x52, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x52, 0x65, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x0b,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x2a, 0x30, 0x0a, 0x0d,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0x92,
	0x01, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a,
	0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x42, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x18, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_data_client_data_proto_rawDescOnce sync.Once
	file_client_data_client_data_proto_rawDescData = file_client_data_client_data_proto_rawDesc
)

func file_client_data_client_data_proto_rawDescGZIP() []byte {
	file_client_data_client_data_proto_rawDescOnce.Do(func() {
		file_client_data_client_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_data_client_data_proto_rawDescData)
	})
	return file_client_data_client_data_proto_rawDescData
}

var file_client_data_client_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_data_client_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_client_data_client_data_proto_goTypes = []interface{}{
	(UploadResCode)(0),  // 0: client_data.UploadResCode
	(*Chunk)(nil),       // 1: client_data.Chunk
	(*UploadReq)(nil),   // 2: client_data.UploadReq
	(*UploadRes)(nil),   // 3: client_data.UploadRes
	(*DownloadReq)(nil), // 4: client_data.DownloadReq
}
var file_client_data_client_data_proto_depIdxs = []int32{
	1, // 0: client_data.UploadReq.DataChunk:type_name -> client_data.Chunk
	0, // 1: client_data.UploadRes.ResCode:type_name -> client_data.UploadResCode
	2, // 2: client_data.ClientData.UploadFile:input_type -> client_data.UploadReq
	4, // 3: client_data.ClientData.DownloadFile:input_type -> client_data.DownloadReq
	3, // 4: client_data.ClientData.UploadFile:output_type -> client_data.UploadRes
	1, // 5: client_data.ClientData.DownloadFile:output_type -> client_data.Chunk
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_client_data_client_data_proto_init() }
func file_client_data_client_data_proto_init() {
	if File_client_data_client_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_data_client_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_client_data_client_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReq); i {
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
		file_client_data_client_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRes); i {
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
		file_client_data_client_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReq); i {
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
	file_client_data_client_data_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UploadReq_FileName)(nil),
		(*UploadReq_DataChunk)(nil),
	}
	file_client_data_client_data_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*DownloadReq_FileName)(nil),
		(*DownloadReq_ChunkId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_data_client_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_data_client_data_proto_goTypes,
		DependencyIndexes: file_client_data_client_data_proto_depIdxs,
		EnumInfos:         file_client_data_client_data_proto_enumTypes,
		MessageInfos:      file_client_data_client_data_proto_msgTypes,
	}.Build()
	File_client_data_client_data_proto = out.File
	file_client_data_client_data_proto_rawDesc = nil
	file_client_data_client_data_proto_goTypes = nil
	file_client_data_client_data_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientDataClient is the client API for ClientData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientDataClient interface {
	//
	//Cliente envía el nombre de archivo a un DataNode, seguido de una secuencia de chunks
	//Luego de que el DataNode haga "lo que tiene que hacer" (distribuír los chunks entre los datanodes), le responde con un UploadRes
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (ClientData_UploadFileClient, error)
	//
	//Cliente envía el nombre de archivo a un DataNode, seguido de los ids de los chunks que el DataNode tiene de ese archivo (según lo que le dijo el DataNode)
	//El DataNode le retorna todos los chunks que tiene almacenado
	//El Cliente debe llamar a esta función con los tres DataNode, y luego ordenarlos en su propio proceso
	DownloadFile(ctx context.Context, opts ...grpc.CallOption) (ClientData_DownloadFileClient, error)
}

type clientDataClient struct {
	cc grpc.ClientConnInterface
}

func NewClientDataClient(cc grpc.ClientConnInterface) ClientDataClient {
	return &clientDataClient{cc}
}

func (c *clientDataClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (ClientData_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientData_serviceDesc.Streams[0], "/client_data.ClientData/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientDataUploadFileClient{stream}
	return x, nil
}

type ClientData_UploadFileClient interface {
	Send(*UploadReq) error
	CloseAndRecv() (*UploadRes, error)
	grpc.ClientStream
}

type clientDataUploadFileClient struct {
	grpc.ClientStream
}

func (x *clientDataUploadFileClient) Send(m *UploadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientDataUploadFileClient) CloseAndRecv() (*UploadRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientDataClient) DownloadFile(ctx context.Context, opts ...grpc.CallOption) (ClientData_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientData_serviceDesc.Streams[1], "/client_data.ClientData/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientDataDownloadFileClient{stream}
	return x, nil
}

type ClientData_DownloadFileClient interface {
	Send(*DownloadReq) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type clientDataDownloadFileClient struct {
	grpc.ClientStream
}

func (x *clientDataDownloadFileClient) Send(m *DownloadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientDataDownloadFileClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientDataServer is the server API for ClientData service.
type ClientDataServer interface {
	//
	//Cliente envía el nombre de archivo a un DataNode, seguido de una secuencia de chunks
	//Luego de que el DataNode haga "lo que tiene que hacer" (distribuír los chunks entre los datanodes), le responde con un UploadRes
	UploadFile(ClientData_UploadFileServer) error
	//
	//Cliente envía el nombre de archivo a un DataNode, seguido de los ids de los chunks que el DataNode tiene de ese archivo (según lo que le dijo el DataNode)
	//El DataNode le retorna todos los chunks que tiene almacenado
	//El Cliente debe llamar a esta función con los tres DataNode, y luego ordenarlos en su propio proceso
	DownloadFile(ClientData_DownloadFileServer) error
}

// UnimplementedClientDataServer can be embedded to have forward compatible implementations.
type UnimplementedClientDataServer struct {
}

func (*UnimplementedClientDataServer) UploadFile(ClientData_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedClientDataServer) DownloadFile(ClientData_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}

func RegisterClientDataServer(s *grpc.Server, srv ClientDataServer) {
	s.RegisterService(&_ClientData_serviceDesc, srv)
}

func _ClientData_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientDataServer).UploadFile(&clientDataUploadFileServer{stream})
}

type ClientData_UploadFileServer interface {
	SendAndClose(*UploadRes) error
	Recv() (*UploadReq, error)
	grpc.ServerStream
}

type clientDataUploadFileServer struct {
	grpc.ServerStream
}

func (x *clientDataUploadFileServer) SendAndClose(m *UploadRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientDataUploadFileServer) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientData_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientDataServer).DownloadFile(&clientDataDownloadFileServer{stream})
}

type ClientData_DownloadFileServer interface {
	Send(*Chunk) error
	Recv() (*DownloadReq, error)
	grpc.ServerStream
}

type clientDataDownloadFileServer struct {
	grpc.ServerStream
}

func (x *clientDataDownloadFileServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientDataDownloadFileServer) Recv() (*DownloadReq, error) {
	m := new(DownloadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClientData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client_data.ClientData",
	HandlerType: (*ClientDataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _ClientData_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _ClientData_DownloadFile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "client_data/client_data.proto",
}