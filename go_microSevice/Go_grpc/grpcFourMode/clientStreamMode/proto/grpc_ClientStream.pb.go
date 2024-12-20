// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: grpc_ClientStream.proto

package proto

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

type SquareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquareRequest) Reset() {
	*x = SquareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_ClientStream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRequest) ProtoMessage() {}

func (x *SquareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_ClientStream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRequest.ProtoReflect.Descriptor instead.
func (*SquareRequest) Descriptor() ([]byte, []int) {
	return file_grpc_ClientStream_proto_rawDescGZIP(), []int{0}
}

func (x *SquareRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SquareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SquareResponse) Reset() {
	*x = SquareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_ClientStream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareResponse) ProtoMessage() {}

func (x *SquareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_ClientStream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareResponse.ProtoReflect.Descriptor instead.
func (*SquareResponse) Descriptor() ([]byte, []int) {
	return file_grpc_ClientStream_proto_rawDescGZIP(), []int{1}
}

func (x *SquareResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_grpc_ClientStream_proto protoreflect.FileDescriptor

var file_grpc_ClientStream_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x27, 0x0a, 0x0d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x4c, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_ClientStream_proto_rawDescOnce sync.Once
	file_grpc_ClientStream_proto_rawDescData = file_grpc_ClientStream_proto_rawDesc
)

func file_grpc_ClientStream_proto_rawDescGZIP() []byte {
	file_grpc_ClientStream_proto_rawDescOnce.Do(func() {
		file_grpc_ClientStream_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_ClientStream_proto_rawDescData)
	})
	return file_grpc_ClientStream_proto_rawDescData
}

var file_grpc_ClientStream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_ClientStream_proto_goTypes = []interface{}{
	(*SquareRequest)(nil),  // 0: proto.SquareRequest
	(*SquareResponse)(nil), // 1: proto.SquareResponse
}
var file_grpc_ClientStream_proto_depIdxs = []int32{
	0, // 0: proto.Calculator.SquareBatch:input_type -> proto.SquareRequest
	1, // 1: proto.Calculator.SquareBatch:output_type -> proto.SquareResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_ClientStream_proto_init() }
func file_grpc_ClientStream_proto_init() {
	if File_grpc_ClientStream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_ClientStream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRequest); i {
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
		file_grpc_ClientStream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareResponse); i {
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
			RawDescriptor: file_grpc_ClientStream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_ClientStream_proto_goTypes,
		DependencyIndexes: file_grpc_ClientStream_proto_depIdxs,
		MessageInfos:      file_grpc_ClientStream_proto_msgTypes,
	}.Build()
	File_grpc_ClientStream_proto = out.File
	file_grpc_ClientStream_proto_rawDesc = nil
	file_grpc_ClientStream_proto_goTypes = nil
	file_grpc_ClientStream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	SquareBatch(ctx context.Context, opts ...grpc.CallOption) (Calculator_SquareBatchClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) SquareBatch(ctx context.Context, opts ...grpc.CallOption) (Calculator_SquareBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[0], "/proto.Calculator/SquareBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorSquareBatchClient{stream}
	return x, nil
}

type Calculator_SquareBatchClient interface {
	Send(*SquareRequest) error
	CloseAndRecv() (*SquareResponse, error)
	grpc.ClientStream
}

type calculatorSquareBatchClient struct {
	grpc.ClientStream
}

func (x *calculatorSquareBatchClient) Send(m *SquareRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorSquareBatchClient) CloseAndRecv() (*SquareResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SquareResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	SquareBatch(Calculator_SquareBatchServer) error
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) SquareBatch(Calculator_SquareBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method SquareBatch not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_SquareBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).SquareBatch(&calculatorSquareBatchServer{stream})
}

type Calculator_SquareBatchServer interface {
	SendAndClose(*SquareResponse) error
	Recv() (*SquareRequest, error)
	grpc.ServerStream
}

type calculatorSquareBatchServer struct {
	grpc.ServerStream
}

func (x *calculatorSquareBatchServer) SendAndClose(m *SquareResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorSquareBatchServer) Recv() (*SquareRequest, error) {
	m := new(SquareRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SquareBatch",
			Handler:       _Calculator_SquareBatch_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_ClientStream.proto",
}
