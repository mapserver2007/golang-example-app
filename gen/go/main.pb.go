// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: main.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetUsersAndItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*GetUserResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Items []*GetItemResponse `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUsersAndItemsResponse) Reset() {
	*x = GetUsersAndItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersAndItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersAndItemsResponse) ProtoMessage() {}

func (x *GetUsersAndItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersAndItemsResponse.ProtoReflect.Descriptor instead.
func (*GetUsersAndItemsResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *GetUsersAndItemsResponse) GetUsers() []*GetUserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetUsersAndItemsResponse) GetItems() []*GetItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type PostUsersAndItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*PostUserRequest `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Items []*PostItemRequest `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PostUsersAndItemsRequest) Reset() {
	*x = PostUsersAndItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostUsersAndItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUsersAndItemsRequest) ProtoMessage() {}

func (x *PostUsersAndItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUsersAndItemsRequest.ProtoReflect.Descriptor instead.
func (*PostUsersAndItemsRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *PostUsersAndItemsRequest) GetUsers() []*PostUserRequest {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *PostUsersAndItemsRequest) GetItems() []*PostItemRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x7a, 0x0a, 0x18, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x41, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0xed, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x72, 0x0a, 0x11,
	0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x0a, 0x5a, 0x08, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_main_proto_goTypes = []interface{}{
	(*GetUsersAndItemsResponse)(nil), // 0: example.GetUsersAndItemsResponse
	(*PostUsersAndItemsRequest)(nil), // 1: example.PostUsersAndItemsRequest
	(*GetUserResponse)(nil),          // 2: example.GetUserResponse
	(*GetItemResponse)(nil),          // 3: example.GetItemResponse
	(*PostUserRequest)(nil),          // 4: example.PostUserRequest
	(*PostItemRequest)(nil),          // 5: example.PostItemRequest
	(*empty.Empty)(nil),              // 6: google.protobuf.Empty
	(*SimpleApiResponse)(nil),        // 7: example.SimpleApiResponse
}
var file_main_proto_depIdxs = []int32{
	2, // 0: example.GetUsersAndItemsResponse.users:type_name -> example.GetUserResponse
	3, // 1: example.GetUsersAndItemsResponse.items:type_name -> example.GetItemResponse
	4, // 2: example.PostUsersAndItemsRequest.users:type_name -> example.PostUserRequest
	5, // 3: example.PostUsersAndItemsRequest.items:type_name -> example.PostItemRequest
	6, // 4: example.MainService.GetUsersAndItems:input_type -> google.protobuf.Empty
	1, // 5: example.MainService.PostUsersAndItems:input_type -> example.PostUsersAndItemsRequest
	0, // 6: example.MainService.GetUsersAndItems:output_type -> example.GetUsersAndItemsResponse
	7, // 7: example.MainService.PostUsersAndItems:output_type -> example.SimpleApiResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	file_user_proto_init()
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersAndItemsResponse); i {
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
		file_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostUsersAndItemsRequest); i {
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
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MainServiceClient is the client API for MainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MainServiceClient interface {
	GetUsersAndItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetUsersAndItemsResponse, error)
	PostUsersAndItems(ctx context.Context, in *PostUsersAndItemsRequest, opts ...grpc.CallOption) (*SimpleApiResponse, error)
}

type mainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMainServiceClient(cc grpc.ClientConnInterface) MainServiceClient {
	return &mainServiceClient{cc}
}

func (c *mainServiceClient) GetUsersAndItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetUsersAndItemsResponse, error) {
	out := new(GetUsersAndItemsResponse)
	err := c.cc.Invoke(ctx, "/example.MainService/GetUsersAndItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) PostUsersAndItems(ctx context.Context, in *PostUsersAndItemsRequest, opts ...grpc.CallOption) (*SimpleApiResponse, error) {
	out := new(SimpleApiResponse)
	err := c.cc.Invoke(ctx, "/example.MainService/PostUsersAndItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServiceServer is the server API for MainService service.
type MainServiceServer interface {
	GetUsersAndItems(context.Context, *empty.Empty) (*GetUsersAndItemsResponse, error)
	PostUsersAndItems(context.Context, *PostUsersAndItemsRequest) (*SimpleApiResponse, error)
}

// UnimplementedMainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMainServiceServer struct {
}

func (*UnimplementedMainServiceServer) GetUsersAndItems(context.Context, *empty.Empty) (*GetUsersAndItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersAndItems not implemented")
}
func (*UnimplementedMainServiceServer) PostUsersAndItems(context.Context, *PostUsersAndItemsRequest) (*SimpleApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUsersAndItems not implemented")
}

func RegisterMainServiceServer(s *grpc.Server, srv MainServiceServer) {
	s.RegisterService(&_MainService_serviceDesc, srv)
}

func _MainService_GetUsersAndItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).GetUsersAndItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.MainService/GetUsersAndItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).GetUsersAndItems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_PostUsersAndItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUsersAndItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).PostUsersAndItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.MainService/PostUsersAndItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).PostUsersAndItems(ctx, req.(*PostUsersAndItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.MainService",
	HandlerType: (*MainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsersAndItems",
			Handler:    _MainService_GetUsersAndItems_Handler,
		},
		{
			MethodName: "PostUsersAndItems",
			Handler:    _MainService_PostUsersAndItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
