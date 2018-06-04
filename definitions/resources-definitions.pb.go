// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources-definitions.proto

package definitions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateResourceParams struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResourceParams) Reset()         { *m = CreateResourceParams{} }
func (m *CreateResourceParams) String() string { return proto.CompactTextString(m) }
func (*CreateResourceParams) ProtoMessage()    {}
func (*CreateResourceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_definitions_fee3dc46cae993e1, []int{0}
}
func (m *CreateResourceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResourceParams.Unmarshal(m, b)
}
func (m *CreateResourceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResourceParams.Marshal(b, m, deterministic)
}
func (dst *CreateResourceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResourceParams.Merge(dst, src)
}
func (m *CreateResourceParams) XXX_Size() int {
	return xxx_messageInfo_CreateResourceParams.Size(m)
}
func (m *CreateResourceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResourceParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResourceParams proto.InternalMessageInfo

func (m *CreateResourceParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateResourceParams) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type EditResourceParams struct {
	ResourceId           int32    `protobuf:"varint,1,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditResourceParams) Reset()         { *m = EditResourceParams{} }
func (m *EditResourceParams) String() string { return proto.CompactTextString(m) }
func (*EditResourceParams) ProtoMessage()    {}
func (*EditResourceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_definitions_fee3dc46cae993e1, []int{1}
}
func (m *EditResourceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditResourceParams.Unmarshal(m, b)
}
func (m *EditResourceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditResourceParams.Marshal(b, m, deterministic)
}
func (dst *EditResourceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditResourceParams.Merge(dst, src)
}
func (m *EditResourceParams) XXX_Size() int {
	return xxx_messageInfo_EditResourceParams.Size(m)
}
func (m *EditResourceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EditResourceParams.DiscardUnknown(m)
}

var xxx_messageInfo_EditResourceParams proto.InternalMessageInfo

func (m *EditResourceParams) GetResourceId() int32 {
	if m != nil {
		return m.ResourceId
	}
	return 0
}

func (m *EditResourceParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EditResourceParams) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type DeleteResourceParams struct {
	ResourceId           int32    `protobuf:"varint,1,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResourceParams) Reset()         { *m = DeleteResourceParams{} }
func (m *DeleteResourceParams) String() string { return proto.CompactTextString(m) }
func (*DeleteResourceParams) ProtoMessage()    {}
func (*DeleteResourceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_definitions_fee3dc46cae993e1, []int{2}
}
func (m *DeleteResourceParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResourceParams.Unmarshal(m, b)
}
func (m *DeleteResourceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResourceParams.Marshal(b, m, deterministic)
}
func (dst *DeleteResourceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResourceParams.Merge(dst, src)
}
func (m *DeleteResourceParams) XXX_Size() int {
	return xxx_messageInfo_DeleteResourceParams.Size(m)
}
func (m *DeleteResourceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResourceParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResourceParams proto.InternalMessageInfo

func (m *DeleteResourceParams) GetResourceId() int32 {
	if m != nil {
		return m.ResourceId
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateResourceParams)(nil), "definitions.CreateResourceParams")
	proto.RegisterType((*EditResourceParams)(nil), "definitions.EditResourceParams")
	proto.RegisterType((*DeleteResourceParams)(nil), "definitions.DeleteResourceParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Resources service

type ResourcesClient interface {
	GetResources(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GeneralResponse, error)
	AddResource(ctx context.Context, in *CreateResourceParams, opts ...grpc.CallOption) (*GeneralResponse, error)
	EditResource(ctx context.Context, in *EditResourceParams, opts ...grpc.CallOption) (*GeneralResponse, error)
	DeleteResource(ctx context.Context, in *DeleteResourceParams, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type resourcesClient struct {
	cc *grpc.ClientConn
}

func NewResourcesClient(cc *grpc.ClientConn) ResourcesClient {
	return &resourcesClient{cc}
}

func (c *resourcesClient) GetResources(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Resources/GetResources", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) AddResource(ctx context.Context, in *CreateResourceParams, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Resources/AddResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) EditResource(ctx context.Context, in *EditResourceParams, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Resources/EditResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) DeleteResource(ctx context.Context, in *DeleteResourceParams, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Resources/DeleteResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Resources service

type ResourcesServer interface {
	GetResources(context.Context, *Empty) (*GeneralResponse, error)
	AddResource(context.Context, *CreateResourceParams) (*GeneralResponse, error)
	EditResource(context.Context, *EditResourceParams) (*GeneralResponse, error)
	DeleteResource(context.Context, *DeleteResourceParams) (*GeneralResponse, error)
}

func RegisterResourcesServer(s *grpc.Server, srv ResourcesServer) {
	s.RegisterService(&_Resources_serviceDesc, srv)
}

func _Resources_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Resources/GetResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetResources(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_AddResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).AddResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Resources/AddResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).AddResource(ctx, req.(*CreateResourceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_EditResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditResourceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).EditResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Resources/EditResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).EditResource(ctx, req.(*EditResourceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Resources/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).DeleteResource(ctx, req.(*DeleteResourceParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Resources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "definitions.Resources",
	HandlerType: (*ResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResources",
			Handler:    _Resources_GetResources_Handler,
		},
		{
			MethodName: "AddResource",
			Handler:    _Resources_AddResource_Handler,
		},
		{
			MethodName: "EditResource",
			Handler:    _Resources_EditResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _Resources_DeleteResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources-definitions.proto",
}

func init() {
	proto.RegisterFile("resources-definitions.proto", fileDescriptor_resources_definitions_fee3dc46cae993e1)
}

var fileDescriptor_resources_definitions_fee3dc46cae993e1 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xed, 0xfa, 0x01, 0x9d, 0x5d, 0x3c, 0x0c, 0x3d, 0x94, 0x2a, 0x54, 0xf7, 0xe4, 0xc5,
	0x1e, 0xf4, 0xe0, 0x4d, 0xf0, 0xa3, 0x14, 0x4f, 0x96, 0xf5, 0x2c, 0x12, 0xcd, 0x48, 0x03, 0xdd,
	0x64, 0xc9, 0xc4, 0x83, 0xbf, 0xd0, 0xbf, 0x25, 0xdd, 0x1a, 0x9b, 0xb4, 0x0b, 0xb9, 0xcd, 0xbe,
	0x9b, 0x79, 0xf2, 0xf2, 0x10, 0x38, 0xb1, 0xc4, 0xe6, 0xcb, 0x7e, 0x10, 0x5f, 0x4a, 0xfa, 0x54,
	0x5a, 0x39, 0x65, 0x34, 0x4f, 0x1a, 0x6b, 0x9c, 0xc1, 0x3c, 0x88, 0x46, 0x05, 0x2f, 0x84, 0x25,
	0xb9, 0xfe, 0x55, 0xde, 0xc2, 0xe0, 0xc1, 0x92, 0x70, 0x54, 0xfd, 0xed, 0xcf, 0x85, 0x15, 0x35,
	0x23, 0xc2, 0x81, 0x16, 0x35, 0x0d, 0x7b, 0x67, 0xbd, 0x8b, 0x7e, 0xd5, 0xce, 0xab, 0xac, 0x11,
	0x6e, 0x31, 0xcc, 0xd6, 0xd9, 0x6a, 0x2e, 0x5f, 0x01, 0xa7, 0x52, 0xb9, 0xad, 0xed, 0x31, 0xe4,
	0xbe, 0xcf, 0x9b, 0x92, 0x2d, 0xe4, 0xb0, 0x02, 0x1f, 0x3d, 0xc9, 0x7f, 0x7c, 0xd6, 0x81, 0xdf,
	0x0f, 0xf0, 0x37, 0x30, 0x78, 0xa4, 0x25, 0xed, 0xd4, 0x4b, 0x5d, 0x70, 0xf5, 0x93, 0x41, 0xdf,
	0xef, 0x30, 0xde, 0x43, 0x31, 0x23, 0xb7, 0xf9, 0xc6, 0x49, 0x28, 0x69, 0x5a, 0x37, 0xee, 0x7b,
	0x74, 0x1a, 0x65, 0x33, 0xd2, 0x64, 0xc5, 0xb2, 0x22, 0x6e, 0x8c, 0x66, 0x2a, 0xf7, 0x70, 0x0e,
	0xf9, 0x9d, 0x94, 0x9e, 0x81, 0xe7, 0xd1, 0xf1, 0x2e, 0x87, 0x49, 0xe2, 0x33, 0x14, 0xa1, 0x3b,
	0x1c, 0xc7, 0xad, 0x76, 0xb4, 0x26, 0x81, 0x2f, 0x70, 0x1c, 0xdb, 0xda, 0x6a, 0xd9, 0xa5, 0x32,
	0x05, 0x7d, 0x3f, 0x6a, 0x1f, 0xca, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x88, 0xbe,
	0x5d, 0x62, 0x02, 0x00, 0x00,
}