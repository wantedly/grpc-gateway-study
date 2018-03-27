// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

/*
Package api_pb is a generated protocol buffer package.

It is generated from these files:
	book.proto

It has these top-level messages:
	Book
	ListBooksRequest
	ListBooksResponse
	GetBookRequest
	CreateBookRequest
	UpdateBookRequest
	DeleteBookRequest
*/
package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type Book struct {
	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Book) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type ListBooksRequest struct {
}

func (m *ListBooksRequest) Reset()                    { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()               {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListBooksResponse struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *ListBooksResponse) Reset()                    { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()               {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type GetBookRequest struct {
	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *GetBookRequest) Reset()                    { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()               {}
func (*GetBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type CreateBookRequest struct {
	Book *Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
}

func (m *CreateBookRequest) Reset()                    { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()               {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type UpdateBookRequest struct {
	Book *Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
}

func (m *UpdateBookRequest) Reset()                    { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()               {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type DeleteBookRequest struct {
	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *DeleteBookRequest) Reset()                    { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()               {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "com.github.wantedly.grpc_gateway_study.api.Book")
	proto.RegisterType((*ListBooksRequest)(nil), "com.github.wantedly.grpc_gateway_study.api.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "com.github.wantedly.grpc_gateway_study.api.ListBooksResponse")
	proto.RegisterType((*GetBookRequest)(nil), "com.github.wantedly.grpc_gateway_study.api.GetBookRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "com.github.wantedly.grpc_gateway_study.api.CreateBookRequest")
	proto.RegisterType((*UpdateBookRequest)(nil), "com.github.wantedly.grpc_gateway_study.api.UpdateBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "com.github.wantedly.grpc_gateway_study.api.DeleteBookRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := grpc.Invoke(ctx, "/com.github.wantedly.grpc_gateway_study.api.BookService/ListBooks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/com.github.wantedly.grpc_gateway_study.api.BookService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/com.github.wantedly.grpc_gateway_study.api.BookService/CreateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/com.github.wantedly.grpc_gateway_study.api.BookService/UpdateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/com.github.wantedly.grpc_gateway_study.api.BookService/DeleteBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceServer interface {
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*google_protobuf1.Empty, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.wantedly.grpc_gateway_study.api.BookService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.wantedly.grpc_gateway_study.api.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.wantedly.grpc_gateway_study.api.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.wantedly.grpc_gateway_study.api.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.wantedly.grpc_gateway_study.api.BookService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.wantedly.grpc_gateway_study.api.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}

func init() { proto.RegisterFile("book.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x89, 0xde, 0x9b, 0xdb, 0x9e, 0x42, 0x69, 0x07, 0xff, 0x94, 0xa8, 0x58, 0xe2, 0xa6,
	0x8a, 0x4c, 0x24, 0xee, 0x8a, 0x6e, 0x6a, 0x55, 0x04, 0x57, 0x15, 0x17, 0xea, 0x22, 0x4c, 0x9a,
	0x31, 0x0e, 0x6d, 0x33, 0x63, 0x66, 0x62, 0x09, 0xd2, 0x8d, 0x6f, 0x20, 0xdd, 0xb8, 0xf6, 0x95,
	0x7c, 0x05, 0x9f, 0xc1, 0xb5, 0xcc, 0x24, 0xb5, 0xad, 0xa1, 0x60, 0x02, 0x2e, 0x93, 0x39, 0xe7,
	0x9b, 0xdf, 0x9c, 0xef, 0xe3, 0x00, 0x84, 0x9c, 0x2f, 0xb0, 0x48, 0xb9, 0xe2, 0xe8, 0xde, 0x9c,
	0xaf, 0x70, 0xcc, 0xd4, 0x87, 0x2c, 0xc4, 0x6b, 0x92, 0x28, 0x1a, 0x2d, 0x73, 0x1c, 0xa7, 0x62,
	0x1e, 0xc4, 0x44, 0xd1, 0x35, 0xc9, 0x03, 0xa9, 0xb2, 0x28, 0xc7, 0x44, 0x30, 0xe7, 0x66, 0xcc,
	0x79, 0xbc, 0xa4, 0x1e, 0x11, 0xcc, 0x23, 0x49, 0xc2, 0x15, 0x51, 0x8c, 0x27, 0xb2, 0x50, 0x72,
	0x6e, 0x94, 0xa7, 0xe6, 0x2b, 0xcc, 0xde, 0x7b, 0x74, 0x25, 0x54, 0x5e, 0x1c, 0xba, 0xb7, 0xe1,
	0x6c, 0xc2, 0xf9, 0x02, 0x5d, 0x87, 0x0b, 0x7d, 0x79, 0xc0, 0xa2, 0x81, 0x35, 0xb4, 0x46, 0xed,
	0x99, 0xad, 0x3f, 0x5f, 0x44, 0x2e, 0x82, 0xde, 0x4b, 0x26, 0x95, 0x2e, 0x92, 0x33, 0xfa, 0x31,
	0xa3, 0x52, 0xb9, 0xef, 0xa0, 0x7f, 0xf0, 0x4f, 0x0a, 0x9e, 0x48, 0x8a, 0x9e, 0xc1, 0xb9, 0x6e,
	0x91, 0x03, 0x6b, 0x78, 0x79, 0xd4, 0xf1, 0x1f, 0xe0, 0x7f, 0x7f, 0x00, 0xd6, 0x4a, 0xb3, 0xa2,
	0xdd, 0xbd, 0x0b, 0xdd, 0xe7, 0xd4, 0x68, 0x97, 0xd7, 0x9d, 0x66, 0x7b, 0x03, 0xfd, 0x27, 0x29,
	0x25, 0x8a, 0x1e, 0x56, 0x4f, 0xe1, 0x4c, 0x1f, 0x9b, 0xd2, 0x26, 0x18, 0xa6, 0x5b, 0x4b, 0xbf,
	0x16, 0xd1, 0x7f, 0x91, 0xbe, 0x0f, 0xfd, 0x29, 0x5d, 0xd2, 0x63, 0xe9, 0x53, 0x6f, 0xf4, 0x7f,
	0x9d, 0x43, 0x47, 0x17, 0xbe, 0xa2, 0xe9, 0x27, 0x36, 0xa7, 0xe8, 0x9b, 0x05, 0xed, 0x3f, 0xc3,
	0x47, 0x8f, 0xea, 0x30, 0xfc, 0xed, 0xa3, 0xf3, 0xb8, 0x61, 0x77, 0xe1, 0xb8, 0xdb, 0xfd, 0xf2,
	0xe3, 0xe7, 0xf6, 0x52, 0x0b, 0xd9, 0x9e, 0x71, 0x0e, 0x7d, 0xb5, 0xe0, 0xa2, 0xb4, 0x0e, 0x8d,
	0xeb, 0x48, 0x1f, 0xfb, 0xed, 0xd4, 0x1e, 0xac, 0x3b, 0x30, 0x24, 0x08, 0xf5, 0x0a, 0x12, 0xef,
	0x73, 0x39, 0xcb, 0x0d, 0xda, 0x5a, 0x00, 0xfb, 0x8c, 0xa0, 0x5a, 0x2f, 0xae, 0x64, 0xab, 0x01,
	0xd9, 0x15, 0x43, 0xd6, 0x75, 0xcb, 0x19, 0x8d, 0x4d, 0x04, 0xd0, 0x77, 0x0b, 0x60, 0x1f, 0xaf,
	0x7a, 0x54, 0x95, 0x58, 0x36, 0xa0, 0xba, 0x63, 0xa8, 0x6e, 0xf9, 0x57, 0x0f, 0xe7, 0x85, 0x77,
	0x43, 0x2b, 0x21, 0x37, 0x00, 0xfb, 0x9c, 0xd6, 0x63, 0xac, 0xe4, 0xdb, 0xb9, 0x86, 0x8b, 0x2d,
	0x84, 0x77, 0x5b, 0x08, 0x3f, 0xd5, 0x5b, 0x68, 0xe7, 0x9c, 0x5f, 0x71, 0x6e, 0xd2, 0x7a, 0x6b,
	0x13, 0xc1, 0x02, 0x11, 0x86, 0xb6, 0xe9, 0x79, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xc9,
	0xca, 0x7c, 0x1f, 0x05, 0x00, 0x00,
}