// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dto.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FetchRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db91a37befd03e7a, []int{0}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type FetchResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db91a37befd03e7a, []int{1}
}

func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (m *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(m, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ListRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db91a37befd03e7a, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListResponse struct {
	Page                 uint64     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             uint64     `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total                uint64     `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	List                 []*Product `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Error                string     `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db91a37befd03e7a, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListResponse) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListResponse) GetList() []*Product {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Product struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                string   `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	ChangePrice          uint64   `protobuf:"varint,3,opt,name=changePrice,proto3" json:"changePrice,omitempty"`
	Last                 uint64   `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_db91a37befd03e7a, []int{4}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Product) GetChangePrice() uint64 {
	if m != nil {
		return m.ChangePrice
	}
	return 0
}

func (m *Product) GetLast() uint64 {
	if m != nil {
		return m.Last
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchRequest)(nil), "dto.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "dto.FetchResponse")
	proto.RegisterType((*ListRequest)(nil), "dto.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "dto.ListResponse")
	proto.RegisterType((*Product)(nil), "dto.Product")
}

func init() {
	proto.RegisterFile("proto/dto.proto", fileDescriptor_db91a37befd03e7a)
}

var fileDescriptor_db91a37befd03e7a = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0xf8, 0xda, 0x82, 0x0c, 0x18, 0x61, 0x62, 0x4c, 0xc3, 0xa9, 0x69, 0x62, 0xc2, 0x45,
	0x34, 0x78, 0xf4, 0xe6, 0xc1, 0x93, 0x07, 0x52, 0x6e, 0xde, 0x6a, 0x19, 0x60, 0x93, 0xc2, 0xd6,
	0xdd, 0xc1, 0x83, 0xbf, 0xc1, 0x1f, 0x6d, 0x76, 0x76, 0xd1, 0x3d, 0x75, 0xde, 0xcc, 0x7b, 0xd3,
	0xf7, 0x66, 0xe1, 0xaa, 0x33, 0x9a, 0xf5, 0xfd, 0x86, 0xf5, 0x42, 0x2a, 0x4c, 0x36, 0xac, 0xcb,
	0x02, 0xc6, 0x2f, 0xc4, 0xcd, 0xbe, 0xa2, 0x8f, 0x13, 0x59, 0xc6, 0x09, 0x24, 0x27, 0xd3, 0xe6,
	0xbd, 0xa2, 0x37, 0x1f, 0x56, 0xae, 0x2c, 0x6f, 0xe1, 0x32, 0x30, 0x6c, 0xa7, 0x8f, 0x96, 0xf0,
	0x1a, 0x32, 0x32, 0x46, 0x9b, 0x40, 0xf2, 0xa0, 0x7c, 0x82, 0xd1, 0xab, 0xb2, 0x7c, 0xde, 0x73,
	0x03, 0x7d, 0xbd, 0xdd, 0x5a, 0x62, 0x61, 0xa5, 0x55, 0x40, 0x4e, 0xdc, 0xaa, 0x83, 0xe2, 0xfc,
	0xbf, 0xb4, 0x3d, 0x28, 0xbf, 0x7b, 0x30, 0xf6, 0xea, 0xf0, 0x0f, 0x84, 0xb4, 0xab, 0x77, 0x14,
	0xc4, 0x52, 0xe3, 0x0c, 0x2e, 0xdc, 0x77, 0xad, 0xbe, 0x28, 0xa8, 0x7f, 0xb1, 0x5b, 0xcb, 0x9a,
	0xeb, 0x36, 0x4f, 0xfc, 0x5a, 0x01, 0x58, 0x40, 0xda, 0x2a, 0xcb, 0x79, 0x5a, 0x24, 0xf3, 0xd1,
	0x72, 0xbc, 0x70, 0xd9, 0x57, 0x46, 0x6f, 0x4e, 0x0d, 0x57, 0x32, 0xf9, 0xcb, 0x92, 0xc5, 0x59,
	0x14, 0x0c, 0x02, 0xcd, 0x19, 0x39, 0xd6, 0x07, 0x0a, 0x59, 0xa5, 0x76, 0xa2, 0xce, 0xa8, 0xc6,
	0xbb, 0x18, 0x56, 0x1e, 0x60, 0x01, 0xa3, 0x66, 0x5f, 0x1f, 0x77, 0xb4, 0x92, 0x99, 0x37, 0x12,
	0xb7, 0xdc, 0xae, 0xb6, 0x16, 0x3b, 0x12, 0xca, 0xd5, 0x4b, 0x05, 0xfd, 0x35, 0x99, 0x4f, 0x32,
	0xf8, 0x00, 0x99, 0xdc, 0x19, 0xa7, 0xe2, 0x33, 0x7e, 0x95, 0x19, 0xc6, 0x2d, 0x7f, 0xa2, 0xf2,
	0x1f, 0xde, 0x41, 0xea, 0x8e, 0x86, 0x13, 0x99, 0x46, 0xd7, 0x9f, 0x4d, 0xa3, 0xce, 0x99, 0xfe,
	0x3c, 0x78, 0xcb, 0xe4, 0xe1, 0xdf, 0xfb, 0xf2, 0x79, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xaf,
	0x01, 0x89, 0x96, 0x12, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/dto.Server/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/dto.Server/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedServerServer can be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (*UnimplementedServerServer) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedServerServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dto.Server/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dto.Server/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Server_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Server_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dto.proto",
}
