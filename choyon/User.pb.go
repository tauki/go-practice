// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User.proto

/*
Package User is a generated protocol buffer package.

It is generated from these files:
	User.proto

It has these top-level messages:
	UserRequest
	UserResponse
*/
package User

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

type UserRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserResponse struct {
	Status bool   `protobuf:"varint,10,opt,name=status" json:"status,omitempty"`
	Id     int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *UserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "UserRequest")
	proto.RegisterType((*UserResponse)(nil), "UserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/User/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "User.proto",
}

func init() { proto.RegisterFile("User.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe7, 0xe2, 0x06, 0xf1, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x26, 0xb0,
	0x20, 0x84, 0xa3, 0x94, 0xc0, 0xc5, 0x03, 0xd1, 0x58, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24,
	0xc6, 0xc5, 0x56, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0xc1, 0xa5, 0xc0, 0xa8, 0xc1, 0x11, 0x04,
	0xe5, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x80, 0xcd, 0x63, 0x0e, 0x62, 0xca, 0x4c, 0x81, 0xdb,
	0xc0, 0x84, 0xcd, 0x06, 0x66, 0x24, 0x1b, 0x8c, 0x8c, 0xb9, 0x58, 0x40, 0x36, 0x08, 0x69, 0x73,
	0x71, 0x39, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x82, 0x79, 0x3c, 0x7a, 0x48, 0xee, 0x95, 0xe2, 0xd5,
	0x43, 0x76, 0x84, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x5b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x01, 0xa9, 0x14, 0xde, 0xe4, 0x00, 0x00, 0x00,
}
