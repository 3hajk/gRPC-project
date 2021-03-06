// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.1
// source: proto/dto.proto

package proto

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

type FetchDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchDataRequest) Reset() {
	*x = FetchDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchDataRequest) ProtoMessage() {}

func (x *FetchDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchDataRequest.ProtoReflect.Descriptor instead.
func (*FetchDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{0}
}

func (x *FetchDataRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FetchDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FetchDataResponse) Reset() {
	*x = FetchDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchDataResponse) ProtoMessage() {}

func (x *FetchDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchDataResponse.ProtoReflect.Descriptor instead.
func (*FetchDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{1}
}

func (x *FetchDataResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging  *Page `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting *Sort `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{2}
}

func (x *ListProductsRequest) GetPaging() *Page {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ListProductsRequest) GetSorting() *Sort {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skip  int64 `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{3}
}

func (x *Page) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *Page) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Direct int32  `protobuf:"varint,2,opt,name=direct,proto3" json:"direct,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{4}
}

func (x *Sort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sort) GetDirect() int32 {
	if x != nil {
		return x.Direct
	}
	return 0
}

type ListProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int32      `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	List     []*Product `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{5}
}

func (x *ListProductsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListProductsResponse) GetList() []*Product {
	if x != nil {
		return x.List
	}
	return nil
}

type StreamProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sorting *Sort `protobuf:"bytes,1,opt,name=sorting,proto3" json:"sorting,omitempty"`
}

func (x *StreamProductsRequest) Reset() {
	*x = StreamProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamProductsRequest) ProtoMessage() {}

func (x *StreamProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamProductsRequest.ProtoReflect.Descriptor instead.
func (*StreamProductsRequest) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{6}
}

func (x *StreamProductsRequest) GetSorting() *Sort {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type StreamProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *StreamProductsResponse) Reset() {
	*x = StreamProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamProductsResponse) ProtoMessage() {}

func (x *StreamProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamProductsResponse.ProtoReflect.Descriptor instead.
func (*StreamProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{7}
}

func (x *StreamProductsResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	ChangePrice int32   `protobuf:"varint,3,opt,name=changePrice,proto3" json:"changePrice,omitempty"`
	LastUpdate  uint32  `protobuf:"varint,4,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_dto_proto_rawDescGZIP(), []int{8}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetChangePrice() int32 {
	if x != nil {
		return x.ChangePrice
	}
	return 0
}

func (x *Product) GetLastUpdate() uint32 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

var File_proto_dto_proto protoreflect.FileDescriptor

var file_proto_dto_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x64, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x11,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5d, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b,
	0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x32, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x22, 0x54, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x20, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x3c, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64,
	0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x40, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x75, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x32, 0xd0, 0x01, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64,
	0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1a, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dto_proto_rawDescOnce sync.Once
	file_proto_dto_proto_rawDescData = file_proto_dto_proto_rawDesc
)

func file_proto_dto_proto_rawDescGZIP() []byte {
	file_proto_dto_proto_rawDescOnce.Do(func() {
		file_proto_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dto_proto_rawDescData)
	})
	return file_proto_dto_proto_rawDescData
}

var file_proto_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_dto_proto_goTypes = []interface{}{
	(*FetchDataRequest)(nil),       // 0: dto.FetchDataRequest
	(*FetchDataResponse)(nil),      // 1: dto.FetchDataResponse
	(*ListProductsRequest)(nil),    // 2: dto.ListProductsRequest
	(*Page)(nil),                   // 3: dto.Page
	(*Sort)(nil),                   // 4: dto.Sort
	(*ListProductsResponse)(nil),   // 5: dto.ListProductsResponse
	(*StreamProductsRequest)(nil),  // 6: dto.StreamProductsRequest
	(*StreamProductsResponse)(nil), // 7: dto.StreamProductsResponse
	(*Product)(nil),                // 8: dto.Product
}
var file_proto_dto_proto_depIdxs = []int32{
	3, // 0: dto.ListProductsRequest.paging:type_name -> dto.Page
	4, // 1: dto.ListProductsRequest.sorting:type_name -> dto.Sort
	8, // 2: dto.ListProductsResponse.list:type_name -> dto.Product
	4, // 3: dto.StreamProductsRequest.sorting:type_name -> dto.Sort
	8, // 4: dto.StreamProductsResponse.product:type_name -> dto.Product
	0, // 5: dto.ProductService.Fetch:input_type -> dto.FetchDataRequest
	2, // 6: dto.ProductService.List:input_type -> dto.ListProductsRequest
	6, // 7: dto.ProductService.Stream:input_type -> dto.StreamProductsRequest
	1, // 8: dto.ProductService.Fetch:output_type -> dto.FetchDataResponse
	5, // 9: dto.ProductService.List:output_type -> dto.ListProductsResponse
	7, // 10: dto.ProductService.Stream:output_type -> dto.StreamProductsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_dto_proto_init() }
func file_proto_dto_proto_init() {
	if File_proto_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchDataRequest); i {
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
		file_proto_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchDataResponse); i {
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
		file_proto_dto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsRequest); i {
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
		file_proto_dto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_proto_dto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
		file_proto_dto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsResponse); i {
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
		file_proto_dto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamProductsRequest); i {
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
		file_proto_dto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamProductsResponse); i {
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
		file_proto_dto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_proto_dto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dto_proto_goTypes,
		DependencyIndexes: file_proto_dto_proto_depIdxs,
		MessageInfos:      file_proto_dto_proto_msgTypes,
	}.Build()
	File_proto_dto_proto = out.File
	file_proto_dto_proto_rawDesc = nil
	file_proto_dto_proto_goTypes = nil
	file_proto_dto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	Fetch(ctx context.Context, in *FetchDataRequest, opts ...grpc.CallOption) (*FetchDataResponse, error)
	List(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	Stream(ctx context.Context, in *StreamProductsRequest, opts ...grpc.CallOption) (ProductService_StreamClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Fetch(ctx context.Context, in *FetchDataRequest, opts ...grpc.CallOption) (*FetchDataResponse, error) {
	out := new(FetchDataResponse)
	err := c.cc.Invoke(ctx, "/dto.ProductService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) List(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, "/dto.ProductService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Stream(ctx context.Context, in *StreamProductsRequest, opts ...grpc.CallOption) (ProductService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[0], "/dto.ProductService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_StreamClient interface {
	Recv() (*StreamProductsResponse, error)
	grpc.ClientStream
}

type productServiceStreamClient struct {
	grpc.ClientStream
}

func (x *productServiceStreamClient) Recv() (*StreamProductsResponse, error) {
	m := new(StreamProductsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	Fetch(context.Context, *FetchDataRequest) (*FetchDataResponse, error)
	List(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	Stream(*StreamProductsRequest, ProductService_StreamServer) error
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) Fetch(context.Context, *FetchDataRequest) (*FetchDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedProductServiceServer) List(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProductServiceServer) Stream(*StreamProductsRequest, ProductService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dto.ProductService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Fetch(ctx, req.(*FetchDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dto.ProductService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).List(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamProductsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).Stream(m, &productServiceStreamServer{stream})
}

type ProductService_StreamServer interface {
	Send(*StreamProductsResponse) error
	grpc.ServerStream
}

type productServiceStreamServer struct {
	grpc.ServerStream
}

func (x *productServiceStreamServer) Send(m *StreamProductsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _ProductService_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProductService_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _ProductService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/dto.proto",
}
