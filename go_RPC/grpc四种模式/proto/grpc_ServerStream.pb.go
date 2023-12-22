// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: grpc_ServerStream.proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_ServerStream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_ServerStream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_ServerStream_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_ServerStream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_ServerStream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_ServerStream_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_grpc_ServerStream_proto protoreflect.FileDescriptor

var file_grpc_ServerStream_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32,
	0x3c, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x2b, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_ServerStream_proto_rawDescOnce sync.Once
	file_grpc_ServerStream_proto_rawDescData = file_grpc_ServerStream_proto_rawDesc
)

func file_grpc_ServerStream_proto_rawDescGZIP() []byte {
	file_grpc_ServerStream_proto_rawDescOnce.Do(func() {
		file_grpc_ServerStream_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_ServerStream_proto_rawDescData)
	})
	return file_grpc_ServerStream_proto_rawDescData
}

var file_grpc_ServerStream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_ServerStream_proto_goTypes = []interface{}{
	(*Request)(nil),      // 0: Request
	(*FileResponse)(nil), // 1: FileResponse
}
var file_grpc_ServerStream_proto_depIdxs = []int32{
	0, // 0: ServiceStream.DownLoadFile:input_type -> Request
	1, // 1: ServiceStream.DownLoadFile:output_type -> FileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_ServerStream_proto_init() }
func file_grpc_ServerStream_proto_init() {
	if File_grpc_ServerStream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_ServerStream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_ServerStream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
			RawDescriptor: file_grpc_ServerStream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_ServerStream_proto_goTypes,
		DependencyIndexes: file_grpc_ServerStream_proto_depIdxs,
		MessageInfos:      file_grpc_ServerStream_proto_msgTypes,
	}.Build()
	File_grpc_ServerStream_proto = out.File
	file_grpc_ServerStream_proto_rawDesc = nil
	file_grpc_ServerStream_proto_goTypes = nil
	file_grpc_ServerStream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceStreamClient is the client API for ServiceStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceStreamClient interface {
	DownLoadFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (ServiceStream_DownLoadFileClient, error)
}

type serviceStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceStreamClient(cc grpc.ClientConnInterface) ServiceStreamClient {
	return &serviceStreamClient{cc}
}

func (c *serviceStreamClient) DownLoadFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (ServiceStream_DownLoadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceStream_serviceDesc.Streams[0], "/ServiceStream/DownLoadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceStreamDownLoadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceStream_DownLoadFileClient interface {
	Recv() (*FileResponse, error)
	grpc.ClientStream
}

type serviceStreamDownLoadFileClient struct {
	grpc.ClientStream
}

func (x *serviceStreamDownLoadFileClient) Recv() (*FileResponse, error) {
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceStreamServer is the server API for ServiceStream service.
type ServiceStreamServer interface {
	DownLoadFile(*Request, ServiceStream_DownLoadFileServer) error
}

// UnimplementedServiceStreamServer can be embedded to have forward compatible implementations.
type UnimplementedServiceStreamServer struct {
}

func (*UnimplementedServiceStreamServer) DownLoadFile(*Request, ServiceStream_DownLoadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownLoadFile not implemented")
}

func RegisterServiceStreamServer(s *grpc.Server, srv ServiceStreamServer) {
	s.RegisterService(&_ServiceStream_serviceDesc, srv)
}

func _ServiceStream_DownLoadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceStreamServer).DownLoadFile(m, &serviceStreamDownLoadFileServer{stream})
}

type ServiceStream_DownLoadFileServer interface {
	Send(*FileResponse) error
	grpc.ServerStream
}

type serviceStreamDownLoadFileServer struct {
	grpc.ServerStream
}

func (x *serviceStreamDownLoadFileServer) Send(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ServiceStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceStream",
	HandlerType: (*ServiceStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownLoadFile",
			Handler:       _ServiceStream_DownLoadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_ServerStream.proto",
}
