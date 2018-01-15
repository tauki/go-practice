// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filesend.proto

/*
Package protobuffer is a generated protocol buffer package.

It is generated from these files:
	filesend.proto

It has these top-level messages:
	File
	Files
	Resp
*/
package protobuffer

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

type File struct {
	File     []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName" json:"fileName,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *File) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *File) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type Files struct {
	Files []*File `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
}

func (m *Files) Reset()                    { *m = Files{} }
func (m *Files) String() string            { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()               {}
func (*Files) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Files) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type Resp struct {
	Resp string `protobuf:"bytes,1,opt,name=resp" json:"resp,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resp) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func init() {
	proto.RegisterType((*File)(nil), "protobuffer.File")
	proto.RegisterType((*Files)(nil), "protobuffer.Files")
	proto.RegisterType((*Resp)(nil), "protobuffer.Resp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FileSender service

type FileSenderClient interface {
	SendFile(ctx context.Context, in *Files, opts ...grpc.CallOption) (*Resp, error)
}

type fileSenderClient struct {
	cc *grpc.ClientConn
}

func NewFileSenderClient(cc *grpc.ClientConn) FileSenderClient {
	return &fileSenderClient{cc}
}

func (c *fileSenderClient) SendFile(ctx context.Context, in *Files, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/protobuffer.FileSender/SendFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileSender service

type FileSenderServer interface {
	SendFile(context.Context, *Files) (*Resp, error)
}

func RegisterFileSenderServer(s *grpc.Server, srv FileSenderServer) {
	s.RegisterService(&_FileSender_serviceDesc, srv)
}

func _FileSender_SendFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Files)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSenderServer).SendFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuffer.FileSender/SendFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSenderServer).SendFile(ctx, req.(*Files))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileSender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuffer.FileSender",
	HandlerType: (*FileSenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFile",
			Handler:    _FileSender_SendFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filesend.proto",
}

func init() { proto.RegisterFile("filesend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xcc, 0x49,
	0x2d, 0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x53, 0x49, 0xa5,
	0x69, 0x69, 0xa9, 0x45, 0x4a, 0x66, 0x5c, 0x2c, 0x6e, 0x99, 0x39, 0xa9, 0x42, 0x42, 0x5c, 0x2c,
	0x20, 0x65, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x90, 0x14, 0x17, 0x07, 0x88,
	0xf6, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0x0c, 0xb8,
	0x58, 0x41, 0xfa, 0x8a, 0x85, 0xd4, 0xb9, 0x58, 0xc1, 0xe6, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70,
	0x1b, 0x09, 0xea, 0x21, 0x99, 0xae, 0x07, 0x52, 0x12, 0x04, 0x91, 0x57, 0x92, 0xe2, 0x62, 0x09,
	0x4a, 0x2d, 0x2e, 0x00, 0xd9, 0x54, 0x94, 0x5a, 0x5c, 0x00, 0xb6, 0x89, 0x33, 0x08, 0xcc, 0x36,
	0x72, 0xe4, 0xe2, 0x02, 0x29, 0x0d, 0x4e, 0xcd, 0x4b, 0x49, 0x2d, 0x12, 0x32, 0xe6, 0xe2, 0x00,
	0xb1, 0x20, 0xee, 0xc2, 0x30, 0xaf, 0x58, 0x0a, 0xd5, 0x0e, 0x90, 0xa1, 0x4a, 0x0c, 0x49, 0x6c,
	0x60, 0x31, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x80, 0x38, 0x8b, 0xee, 0x00, 0x00,
	0x00,
}
