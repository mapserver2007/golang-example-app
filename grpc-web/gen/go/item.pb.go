// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: item.proto

package proto

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetItemResponse) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetItemResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemsResponse) GetItems() []*GetItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_item_proto protoreflect.FileDescriptor

var file_item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xb8, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x50, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x67, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_proto_rawDescOnce sync.Once
	file_item_proto_rawDescData = file_item_proto_rawDesc
)

func file_item_proto_rawDescGZIP() []byte {
	file_item_proto_rawDescOnce.Do(func() {
		file_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_proto_rawDescData)
	})
	return file_item_proto_rawDescData
}

var file_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_item_proto_goTypes = []interface{}{
	(*GetItemRequest)(nil),   // 0: example.GetItemRequest
	(*GetItemResponse)(nil),  // 1: example.GetItemResponse
	(*GetItemsResponse)(nil), // 2: example.GetItemsResponse
	(*empty.Empty)(nil),      // 3: google.protobuf.Empty
}
var file_item_proto_depIdxs = []int32{
	1, // 0: example.GetItemsResponse.items:type_name -> example.GetItemResponse
	0, // 1: example.GetItemsService.GetItem:input_type -> example.GetItemRequest
	3, // 2: example.GetItemsService.GetItems:input_type -> google.protobuf.Empty
	1, // 3: example.GetItemsService.GetItem:output_type -> example.GetItemResponse
	2, // 4: example.GetItemsService.GetItems:output_type -> example.GetItemsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_item_proto_init() }
func file_item_proto_init() {
	if File_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResponse); i {
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
		file_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsResponse); i {
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
			RawDescriptor: file_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_proto_goTypes,
		DependencyIndexes: file_item_proto_depIdxs,
		MessageInfos:      file_item_proto_msgTypes,
	}.Build()
	File_item_proto = out.File
	file_item_proto_rawDesc = nil
	file_item_proto_goTypes = nil
	file_item_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GetItemsServiceClient is the client API for GetItemsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetItemsServiceClient interface {
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	GetItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetItemsResponse, error)
}

type getItemsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetItemsServiceClient(cc grpc.ClientConnInterface) GetItemsServiceClient {
	return &getItemsServiceClient{cc}
}

func (c *getItemsServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/example.GetItemsService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getItemsServiceClient) GetItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, "/example.GetItemsService/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetItemsServiceServer is the server API for GetItemsService service.
type GetItemsServiceServer interface {
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	GetItems(context.Context, *empty.Empty) (*GetItemsResponse, error)
}

// UnimplementedGetItemsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGetItemsServiceServer struct {
}

func (*UnimplementedGetItemsServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (*UnimplementedGetItemsServiceServer) GetItems(context.Context, *empty.Empty) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}

func RegisterGetItemsServiceServer(s *grpc.Server, srv GetItemsServiceServer) {
	s.RegisterService(&_GetItemsService_serviceDesc, srv)
}

func _GetItemsService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetItemsServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.GetItemsService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetItemsServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetItemsService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetItemsServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.GetItemsService/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetItemsServiceServer).GetItems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetItemsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.GetItemsService",
	HandlerType: (*GetItemsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItem",
			Handler:    _GetItemsService_GetItem_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _GetItemsService_GetItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}
