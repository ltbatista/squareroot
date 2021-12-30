// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: squareroot/squarepb/square.proto

package squarepb

import (
	context "context"
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

type SquareRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquareRootRequest) Reset() {
	*x = SquareRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareroot_squarepb_square_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootRequest) ProtoMessage() {}

func (x *SquareRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_squareroot_squarepb_square_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootRequest.ProtoReflect.Descriptor instead.
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return file_squareroot_squarepb_square_proto_rawDescGZIP(), []int{0}
}

func (x *SquareRootRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SquareRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberRoot float64 `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
}

func (x *SquareRootResponse) Reset() {
	*x = SquareRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareroot_squarepb_square_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootResponse) ProtoMessage() {}

func (x *SquareRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_squareroot_squarepb_square_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootResponse.ProtoReflect.Descriptor instead.
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return file_squareroot_squarepb_square_proto_rawDescGZIP(), []int{1}
}

func (x *SquareRootResponse) GetNumberRoot() float64 {
	if x != nil {
		return x.NumberRoot
	}
	return 0
}

var File_squareroot_squarepb_square_proto protoreflect.FileDescriptor

var file_squareroot_squarepb_square_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x74, 0x32, 0x5a,
	0x0a, 0x11, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x19, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f,
	0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_squareroot_squarepb_square_proto_rawDescOnce sync.Once
	file_squareroot_squarepb_square_proto_rawDescData = file_squareroot_squarepb_square_proto_rawDesc
)

func file_squareroot_squarepb_square_proto_rawDescGZIP() []byte {
	file_squareroot_squarepb_square_proto_rawDescOnce.Do(func() {
		file_squareroot_squarepb_square_proto_rawDescData = protoimpl.X.CompressGZIP(file_squareroot_squarepb_square_proto_rawDescData)
	})
	return file_squareroot_squarepb_square_proto_rawDescData
}

var file_squareroot_squarepb_square_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_squareroot_squarepb_square_proto_goTypes = []interface{}{
	(*SquareRootRequest)(nil),  // 0: square.SquareRootRequest
	(*SquareRootResponse)(nil), // 1: square.SquareRootResponse
}
var file_squareroot_squarepb_square_proto_depIdxs = []int32{
	0, // 0: square.SquareRootService.SquareRoot:input_type -> square.SquareRootRequest
	1, // 1: square.SquareRootService.SquareRoot:output_type -> square.SquareRootResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_squareroot_squarepb_square_proto_init() }
func file_squareroot_squarepb_square_proto_init() {
	if File_squareroot_squarepb_square_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_squareroot_squarepb_square_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootRequest); i {
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
		file_squareroot_squarepb_square_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootResponse); i {
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
			RawDescriptor: file_squareroot_squarepb_square_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_squareroot_squarepb_square_proto_goTypes,
		DependencyIndexes: file_squareroot_squarepb_square_proto_depIdxs,
		MessageInfos:      file_squareroot_squarepb_square_proto_msgTypes,
	}.Build()
	File_squareroot_squarepb_square_proto = out.File
	file_squareroot_squarepb_square_proto_rawDesc = nil
	file_squareroot_squarepb_square_proto_goTypes = nil
	file_squareroot_squarepb_square_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SquareRootServiceClient is the client API for SquareRootService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquareRootServiceClient interface {
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type squareRootServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSquareRootServiceClient(cc grpc.ClientConnInterface) SquareRootServiceClient {
	return &squareRootServiceClient{cc}
}

func (c *squareRootServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/square.SquareRootService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareRootServiceServer is the server API for SquareRootService service.
type SquareRootServiceServer interface {
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedSquareRootServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSquareRootServiceServer struct {
}

func (*UnimplementedSquareRootServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterSquareRootServiceServer(s *grpc.Server, srv SquareRootServiceServer) {
	s.RegisterService(&_SquareRootService_serviceDesc, srv)
}

func _SquareRootService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareRootServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/square.SquareRootService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareRootServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SquareRootService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "square.SquareRootService",
	HandlerType: (*SquareRootServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquareRoot",
			Handler:    _SquareRootService_SquareRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squareroot/squarepb/square.proto",
}
