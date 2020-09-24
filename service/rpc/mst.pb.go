// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: service/rpc/mst.proto

package rpc

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

type MSTChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	High  []byte `protobuf:"bytes,3,opt,name=high,proto3" json:"high,omitempty"`
}

func (x *MSTChild) Reset() {
	*x = MSTChild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_mst_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSTChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSTChild) ProtoMessage() {}

func (x *MSTChild) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_mst_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSTChild.ProtoReflect.Descriptor instead.
func (*MSTChild) Descriptor() ([]byte, []int) {
	return file_service_rpc_mst_proto_rawDescGZIP(), []int{0}
}

func (x *MSTChild) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *MSTChild) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *MSTChild) GetHigh() []byte {
	if x != nil {
		return x.High
	}
	return nil
}

type MSTNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level    uint32      `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Low      []byte      `protobuf:"bytes,2,opt,name=low,proto3" json:"low,omitempty"`
	Children []*MSTChild `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *MSTNode) Reset() {
	*x = MSTNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_mst_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSTNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSTNode) ProtoMessage() {}

func (x *MSTNode) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_mst_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSTNode.ProtoReflect.Descriptor instead.
func (*MSTNode) Descriptor() ([]byte, []int) {
	return file_service_rpc_mst_proto_rawDescGZIP(), []int{1}
}

func (x *MSTNode) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MSTNode) GetLow() []byte {
	if x != nil {
		return x.Low
	}
	return nil
}

func (x *MSTNode) GetChildren() []*MSTChild {
	if x != nil {
		return x.Children
	}
	return nil
}

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeHash []byte `protobuf:"bytes,1,opt,name=nodeHash,proto3" json:"nodeHash,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_mst_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_mst_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_mst_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodeRequest) GetNodeHash() []byte {
	if x != nil {
		return x.NodeHash
	}
	return nil
}

var File_service_rpc_mst_proto protoreflect.FileDescriptor

var file_service_rpc_mst_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x76, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x22, 0x46, 0x0a, 0x08,
	0x4d, 0x53, 0x54, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x69, 0x67, 0x68, 0x22, 0x6c, 0x0a, 0x07, 0x4d, 0x53, 0x54, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x75, 0x6c, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x53, 0x54, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x32, 0x5c, 0x0a, 0x0a, 0x4d, 0x53, 0x54, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x76, 0x75, 0x6c, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x76, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x53, 0x54, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x75, 0x6c,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_rpc_mst_proto_rawDescOnce sync.Once
	file_service_rpc_mst_proto_rawDescData = file_service_rpc_mst_proto_rawDesc
)

func file_service_rpc_mst_proto_rawDescGZIP() []byte {
	file_service_rpc_mst_proto_rawDescOnce.Do(func() {
		file_service_rpc_mst_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rpc_mst_proto_rawDescData)
	})
	return file_service_rpc_mst_proto_rawDescData
}

var file_service_rpc_mst_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_rpc_mst_proto_goTypes = []interface{}{
	(*MSTChild)(nil),       // 0: vulture.service.rpc.MSTChild
	(*MSTNode)(nil),        // 1: vulture.service.rpc.MSTNode
	(*GetNodeRequest)(nil), // 2: vulture.service.rpc.GetNodeRequest
}
var file_service_rpc_mst_proto_depIdxs = []int32{
	0, // 0: vulture.service.rpc.MSTNode.children:type_name -> vulture.service.rpc.MSTChild
	2, // 1: vulture.service.rpc.MSTService.GetNode:input_type -> vulture.service.rpc.GetNodeRequest
	1, // 2: vulture.service.rpc.MSTService.GetNode:output_type -> vulture.service.rpc.MSTNode
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_rpc_mst_proto_init() }
func file_service_rpc_mst_proto_init() {
	if File_service_rpc_mst_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_rpc_mst_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSTChild); i {
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
		file_service_rpc_mst_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSTNode); i {
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
		file_service_rpc_mst_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
			RawDescriptor: file_service_rpc_mst_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_rpc_mst_proto_goTypes,
		DependencyIndexes: file_service_rpc_mst_proto_depIdxs,
		MessageInfos:      file_service_rpc_mst_proto_msgTypes,
	}.Build()
	File_service_rpc_mst_proto = out.File
	file_service_rpc_mst_proto_rawDesc = nil
	file_service_rpc_mst_proto_goTypes = nil
	file_service_rpc_mst_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MSTServiceClient is the client API for MSTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MSTServiceClient interface {
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*MSTNode, error)
}

type mSTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMSTServiceClient(cc grpc.ClientConnInterface) MSTServiceClient {
	return &mSTServiceClient{cc}
}

func (c *mSTServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*MSTNode, error) {
	out := new(MSTNode)
	err := c.cc.Invoke(ctx, "/vulture.service.rpc.MSTService/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MSTServiceServer is the server API for MSTService service.
type MSTServiceServer interface {
	GetNode(context.Context, *GetNodeRequest) (*MSTNode, error)
}

// UnimplementedMSTServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMSTServiceServer struct {
}

func (*UnimplementedMSTServiceServer) GetNode(context.Context, *GetNodeRequest) (*MSTNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}

func RegisterMSTServiceServer(s *grpc.Server, srv MSTServiceServer) {
	s.RegisterService(&_MSTService_serviceDesc, srv)
}

func _MSTService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSTServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vulture.service.rpc.MSTService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSTServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MSTService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vulture.service.rpc.MSTService",
	HandlerType: (*MSTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNode",
			Handler:    _MSTService_GetNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/rpc/mst.proto",
}
