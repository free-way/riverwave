// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maps-definitions.proto

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

type Map struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	AreaName             string   `protobuf:"bytes,2,opt,name=AreaName" json:"AreaName,omitempty"`
	AreaCoordinates      string   `protobuf:"bytes,3,opt,name=AreaCoordinates" json:"AreaCoordinates,omitempty"`
	Flags                []*Flag  `protobuf:"bytes,4,rep,name=Flags" json:"Flags,omitempty"`
	EventId              int64    `protobuf:"varint,5,opt,name=EventId" json:"EventId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}
func (*Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{0}
}
func (m *Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map.Unmarshal(m, b)
}
func (m *Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map.Marshal(b, m, deterministic)
}
func (dst *Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map.Merge(dst, src)
}
func (m *Map) XXX_Size() int {
	return xxx_messageInfo_Map.Size(m)
}
func (m *Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Map proto.InternalMessageInfo

func (m *Map) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Map) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *Map) GetAreaCoordinates() string {
	if m != nil {
		return m.AreaCoordinates
	}
	return ""
}

func (m *Map) GetFlags() []*Flag {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *Map) GetEventId() int64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

type Flag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Coordinates          string   `protobuf:"bytes,3,opt,name=Coordinates" json:"Coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Flag) Reset()         { *m = Flag{} }
func (m *Flag) String() string { return proto.CompactTextString(m) }
func (*Flag) ProtoMessage()    {}
func (*Flag) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{1}
}
func (m *Flag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flag.Unmarshal(m, b)
}
func (m *Flag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flag.Marshal(b, m, deterministic)
}
func (dst *Flag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flag.Merge(dst, src)
}
func (m *Flag) XXX_Size() int {
	return xxx_messageInfo_Flag.Size(m)
}
func (m *Flag) XXX_DiscardUnknown() {
	xxx_messageInfo_Flag.DiscardUnknown(m)
}

var xxx_messageInfo_Flag proto.InternalMessageInfo

func (m *Flag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Flag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Flag) GetCoordinates() string {
	if m != nil {
		return m.Coordinates
	}
	return ""
}

type FlagsResponse struct {
	Flags                []*Flag  `protobuf:"bytes,1,rep,name=Flags" json:"Flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagsResponse) Reset()         { *m = FlagsResponse{} }
func (m *FlagsResponse) String() string { return proto.CompactTextString(m) }
func (*FlagsResponse) ProtoMessage()    {}
func (*FlagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{2}
}
func (m *FlagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagsResponse.Unmarshal(m, b)
}
func (m *FlagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagsResponse.Marshal(b, m, deterministic)
}
func (dst *FlagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagsResponse.Merge(dst, src)
}
func (m *FlagsResponse) XXX_Size() int {
	return xxx_messageInfo_FlagsResponse.Size(m)
}
func (m *FlagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlagsResponse proto.InternalMessageInfo

func (m *FlagsResponse) GetFlags() []*Flag {
	if m != nil {
		return m.Flags
	}
	return nil
}

type MapsResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Maps                 []*Map   `protobuf:"bytes,3,rep,name=Maps" json:"Maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapsResponse) Reset()         { *m = MapsResponse{} }
func (m *MapsResponse) String() string { return proto.CompactTextString(m) }
func (*MapsResponse) ProtoMessage()    {}
func (*MapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{3}
}
func (m *MapsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapsResponse.Unmarshal(m, b)
}
func (m *MapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapsResponse.Marshal(b, m, deterministic)
}
func (dst *MapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapsResponse.Merge(dst, src)
}
func (m *MapsResponse) XXX_Size() int {
	return xxx_messageInfo_MapsResponse.Size(m)
}
func (m *MapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MapsResponse proto.InternalMessageInfo

func (m *MapsResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MapsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MapsResponse) GetMaps() []*Map {
	if m != nil {
		return m.Maps
	}
	return nil
}

type MapResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Map                  *Map     `protobuf:"bytes,3,opt,name=Map" json:"Map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapResponse) Reset()         { *m = MapResponse{} }
func (m *MapResponse) String() string { return proto.CompactTextString(m) }
func (*MapResponse) ProtoMessage()    {}
func (*MapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{4}
}
func (m *MapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapResponse.Unmarshal(m, b)
}
func (m *MapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapResponse.Marshal(b, m, deterministic)
}
func (dst *MapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapResponse.Merge(dst, src)
}
func (m *MapResponse) XXX_Size() int {
	return xxx_messageInfo_MapResponse.Size(m)
}
func (m *MapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MapResponse proto.InternalMessageInfo

func (m *MapResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MapResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MapResponse) GetMap() *Map {
	if m != nil {
		return m.Map
	}
	return nil
}

type AddFlagRequest struct {
	MapId                int64    `protobuf:"varint,1,opt,name=MapId" json:"MapId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Coordinates          string   `protobuf:"bytes,3,opt,name=Coordinates" json:"Coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFlagRequest) Reset()         { *m = AddFlagRequest{} }
func (m *AddFlagRequest) String() string { return proto.CompactTextString(m) }
func (*AddFlagRequest) ProtoMessage()    {}
func (*AddFlagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{5}
}
func (m *AddFlagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFlagRequest.Unmarshal(m, b)
}
func (m *AddFlagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFlagRequest.Marshal(b, m, deterministic)
}
func (dst *AddFlagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFlagRequest.Merge(dst, src)
}
func (m *AddFlagRequest) XXX_Size() int {
	return xxx_messageInfo_AddFlagRequest.Size(m)
}
func (m *AddFlagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFlagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFlagRequest proto.InternalMessageInfo

func (m *AddFlagRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *AddFlagRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddFlagRequest) GetCoordinates() string {
	if m != nil {
		return m.Coordinates
	}
	return ""
}

type DeleteFlagRequest struct {
	MapId                int64    `protobuf:"varint,1,opt,name=MapId" json:"MapId,omitempty"`
	FlagId               int64    `protobuf:"varint,2,opt,name=FlagId" json:"FlagId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFlagRequest) Reset()         { *m = DeleteFlagRequest{} }
func (m *DeleteFlagRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFlagRequest) ProtoMessage()    {}
func (*DeleteFlagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{6}
}
func (m *DeleteFlagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFlagRequest.Unmarshal(m, b)
}
func (m *DeleteFlagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFlagRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteFlagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFlagRequest.Merge(dst, src)
}
func (m *DeleteFlagRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFlagRequest.Size(m)
}
func (m *DeleteFlagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFlagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFlagRequest proto.InternalMessageInfo

func (m *DeleteFlagRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *DeleteFlagRequest) GetFlagId() int64 {
	if m != nil {
		return m.FlagId
	}
	return 0
}

type GetFlagsRequest struct {
	MapId                int64    `protobuf:"varint,1,opt,name=MapId" json:"MapId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlagsRequest) Reset()         { *m = GetFlagsRequest{} }
func (m *GetFlagsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlagsRequest) ProtoMessage()    {}
func (*GetFlagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_maps_definitions_b26118dd225d6995, []int{7}
}
func (m *GetFlagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlagsRequest.Unmarshal(m, b)
}
func (m *GetFlagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlagsRequest.Marshal(b, m, deterministic)
}
func (dst *GetFlagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlagsRequest.Merge(dst, src)
}
func (m *GetFlagsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlagsRequest.Size(m)
}
func (m *GetFlagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlagsRequest proto.InternalMessageInfo

func (m *GetFlagsRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func init() {
	proto.RegisterType((*Map)(nil), "definitions.Map")
	proto.RegisterType((*Flag)(nil), "definitions.Flag")
	proto.RegisterType((*FlagsResponse)(nil), "definitions.FlagsResponse")
	proto.RegisterType((*MapsResponse)(nil), "definitions.MapsResponse")
	proto.RegisterType((*MapResponse)(nil), "definitions.MapResponse")
	proto.RegisterType((*AddFlagRequest)(nil), "definitions.AddFlagRequest")
	proto.RegisterType((*DeleteFlagRequest)(nil), "definitions.DeleteFlagRequest")
	proto.RegisterType((*GetFlagsRequest)(nil), "definitions.GetFlagsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Maps service

type MapsClient interface {
	CreateMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*GeneralResponse, error)
	GetMaps(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MapsResponse, error)
	EditMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*GeneralResponse, error)
	DeleteMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*GeneralResponse, error)
	AddFlag(ctx context.Context, in *AddFlagRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	DeleteFlag(ctx context.Context, in *DeleteFlagRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	GetFlags(ctx context.Context, in *GetFlagsRequest, opts ...grpc.CallOption) (*FlagsResponse, error)
}

type mapsClient struct {
	cc *grpc.ClientConn
}

func NewMapsClient(cc *grpc.ClientConn) MapsClient {
	return &mapsClient{cc}
}

func (c *mapsClient) CreateMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/CreateMap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) GetMaps(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MapsResponse, error) {
	out := new(MapsResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/GetMaps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) EditMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/EditMap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) DeleteMap(ctx context.Context, in *Map, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/DeleteMap", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) AddFlag(ctx context.Context, in *AddFlagRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/AddFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) DeleteFlag(ctx context.Context, in *DeleteFlagRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/DeleteFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapsClient) GetFlags(ctx context.Context, in *GetFlagsRequest, opts ...grpc.CallOption) (*FlagsResponse, error) {
	out := new(FlagsResponse)
	err := grpc.Invoke(ctx, "/definitions.Maps/GetFlags", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Maps service

type MapsServer interface {
	CreateMap(context.Context, *Map) (*GeneralResponse, error)
	GetMaps(context.Context, *Empty) (*MapsResponse, error)
	EditMap(context.Context, *Map) (*GeneralResponse, error)
	DeleteMap(context.Context, *Map) (*GeneralResponse, error)
	AddFlag(context.Context, *AddFlagRequest) (*GeneralResponse, error)
	DeleteFlag(context.Context, *DeleteFlagRequest) (*GeneralResponse, error)
	GetFlags(context.Context, *GetFlagsRequest) (*FlagsResponse, error)
}

func RegisterMapsServer(s *grpc.Server, srv MapsServer) {
	s.RegisterService(&_Maps_serviceDesc, srv)
}

func _Maps_CreateMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Map)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).CreateMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/CreateMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).CreateMap(ctx, req.(*Map))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_GetMaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).GetMaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/GetMaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).GetMaps(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_EditMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Map)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).EditMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/EditMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).EditMap(ctx, req.(*Map))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_DeleteMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Map)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).DeleteMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/DeleteMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).DeleteMap(ctx, req.(*Map))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_AddFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).AddFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/AddFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).AddFlag(ctx, req.(*AddFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_DeleteFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).DeleteFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/DeleteFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).DeleteFlag(ctx, req.(*DeleteFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maps_GetFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapsServer).GetFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/definitions.Maps/GetFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapsServer).GetFlags(ctx, req.(*GetFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "definitions.Maps",
	HandlerType: (*MapsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMap",
			Handler:    _Maps_CreateMap_Handler,
		},
		{
			MethodName: "GetMaps",
			Handler:    _Maps_GetMaps_Handler,
		},
		{
			MethodName: "EditMap",
			Handler:    _Maps_EditMap_Handler,
		},
		{
			MethodName: "DeleteMap",
			Handler:    _Maps_DeleteMap_Handler,
		},
		{
			MethodName: "AddFlag",
			Handler:    _Maps_AddFlag_Handler,
		},
		{
			MethodName: "DeleteFlag",
			Handler:    _Maps_DeleteFlag_Handler,
		},
		{
			MethodName: "GetFlags",
			Handler:    _Maps_GetFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maps-definitions.proto",
}

func init() {
	proto.RegisterFile("maps-definitions.proto", fileDescriptor_maps_definitions_b26118dd225d6995)
}

var fileDescriptor_maps_definitions_b26118dd225d6995 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x3e, 0x9a, 0x49, 0x69, 0xe9, 0x08, 0x45, 0xc6, 0x20, 0x64, 0xad, 0x90, 0xea,
	0x0b, 0x3d, 0x94, 0x0b, 0x02, 0xf5, 0x10, 0x95, 0x34, 0xb2, 0x54, 0x73, 0x30, 0x57, 0x2e, 0x0b,
	0x3b, 0x2d, 0x96, 0x12, 0x7b, 0xf1, 0x6e, 0x91, 0xf8, 0x33, 0xfc, 0x31, 0xfe, 0x0c, 0xda, 0xb5,
	0xdd, 0xd8, 0x31, 0x6a, 0x2a, 0xdf, 0xf6, 0xed, 0xbc, 0x9d, 0xf7, 0x66, 0x5e, 0x62, 0x98, 0x6f,
	0xb8, 0x54, 0x6f, 0x05, 0xdd, 0xa4, 0x59, 0xaa, 0xd3, 0x3c, 0x53, 0x67, 0xb2, 0xc8, 0x75, 0x8e,
	0xb3, 0xc6, 0x95, 0x7f, 0xa8, 0x7e, 0xf0, 0x82, 0x44, 0x59, 0x62, 0x7f, 0x1c, 0x70, 0x63, 0x2e,
	0xf1, 0x08, 0x06, 0x91, 0xf0, 0x9c, 0xc0, 0x09, 0xdd, 0x64, 0x10, 0x09, 0xf4, 0xe1, 0x60, 0x51,
	0x10, 0xff, 0xcc, 0x37, 0xe4, 0x0d, 0x02, 0x27, 0x9c, 0x26, 0xf7, 0x18, 0x43, 0x38, 0x36, 0xe7,
	0xcb, 0x3c, 0x2f, 0x44, 0x9a, 0x71, 0x4d, 0xca, 0x73, 0x2d, 0x65, 0xf7, 0x1a, 0x4f, 0x61, 0x74,
	0xb5, 0xe6, 0xb7, 0xca, 0x1b, 0x06, 0x6e, 0x38, 0x3b, 0x3f, 0x39, 0x6b, 0x7a, 0x33, 0x95, 0xa4,
	0xac, 0xa3, 0x07, 0x93, 0xe5, 0x2f, 0xca, 0x74, 0x24, 0xbc, 0x91, 0xf5, 0x50, 0x43, 0x76, 0x0d,
	0x43, 0x43, 0xe9, 0x18, 0x44, 0x18, 0x36, 0xcc, 0xd9, 0x33, 0x06, 0x30, 0xeb, 0x9a, 0x6a, 0x5e,
	0xb1, 0xf7, 0xf0, 0xd4, 0x0a, 0x26, 0xa4, 0x64, 0x9e, 0x29, 0xda, 0x3a, 0x74, 0x1e, 0x76, 0xc8,
	0x6e, 0xe0, 0x30, 0xe6, 0x72, 0xfb, 0x70, 0x0e, 0xe3, 0x2f, 0x9a, 0xeb, 0x3b, 0x55, 0x79, 0xaa,
	0x90, 0x99, 0x24, 0x26, 0xa5, 0xf8, 0x6d, 0x6d, 0xad, 0x86, 0xf8, 0x06, 0x86, 0xa6, 0x83, 0xe7,
	0x5a, 0xa5, 0x67, 0x2d, 0xa5, 0x98, 0xcb, 0xc4, 0x56, 0xd9, 0x77, 0x98, 0x19, 0xd0, 0x5f, 0x86,
	0xd9, 0x40, 0xed, 0xf0, 0xff, 0x53, 0x31, 0x45, 0xf6, 0x15, 0x8e, 0x16, 0x42, 0xd8, 0xf1, 0xe8,
	0xe7, 0x1d, 0x29, 0x8d, 0xcf, 0x61, 0x14, 0x73, 0x79, 0xbf, 0xe1, 0x12, 0xf4, 0x5c, 0xf2, 0x02,
	0x4e, 0x3e, 0xd1, 0x9a, 0x34, 0xed, 0x17, 0x98, 0xc3, 0xd8, 0x90, 0x22, 0x61, 0x25, 0xdc, 0xa4,
	0x42, 0xec, 0x14, 0x8e, 0x57, 0xa4, 0xab, 0xa8, 0x1e, 0x68, 0x70, 0xfe, 0xd7, 0x2d, 0xb7, 0x8a,
	0x17, 0x30, 0xbd, 0x2c, 0x88, 0x6b, 0x32, 0xbf, 0xe6, 0xce, 0xd8, 0xfe, 0xab, 0xd6, 0xcd, 0x8a,
	0x32, 0x2a, 0xf8, 0xba, 0xde, 0x32, 0x7b, 0x82, 0x1f, 0x60, 0xb2, 0x22, 0x6d, 0x3b, 0x61, 0x8b,
	0xba, 0xdc, 0x48, 0xfd, 0xdb, 0x7f, 0xb1, 0xdb, 0x50, 0x35, 0xde, 0x7e, 0x84, 0xc9, 0x52, 0xa4,
	0xba, 0x9f, 0xf0, 0x05, 0x4c, 0xcb, 0x65, 0xf5, 0x7b, 0x7e, 0x05, 0x93, 0x2a, 0x49, 0x7c, 0xd9,
	0xa2, 0xb6, 0xf3, 0xdd, 0xdb, 0xe7, 0x1a, 0x60, 0x9b, 0x19, 0xbe, 0x6e, 0xb1, 0x3b, 0x61, 0x3e,
	0xc2, 0xd5, 0x41, 0x1d, 0x1f, 0xee, 0x72, 0x5b, 0xa9, 0xfa, 0x7e, 0xe7, 0x0f, 0xd7, 0xd8, 0xec,
	0xb7, 0xb1, 0xfd, 0x48, 0xbd, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x71, 0x4e, 0x59, 0xd9,
	0x04, 0x00, 0x00,
}
