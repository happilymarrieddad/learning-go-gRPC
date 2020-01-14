// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package grpc

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

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "auth.LoginReply")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1c, 0xb8, 0x78, 0x7c, 0xf2,
	0xd3, 0x33, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x53, 0x73,
	0x13, 0x33, 0x73, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x29, 0x2e, 0x8e,
	0x82, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x26, 0xb0, 0x04, 0x9c, 0xaf, 0xa4, 0xc4,
	0xc5, 0x05, 0x35, 0xa1, 0x20, 0xa7, 0x12, 0xa4, 0xbf, 0x24, 0x3f, 0x3b, 0x35, 0x0f, 0xa6, 0x1f,
	0xcc, 0x31, 0xb2, 0xe4, 0x62, 0x0b, 0x33, 0x74, 0x2c, 0x2d, 0xc9, 0x10, 0xd2, 0xe7, 0x62, 0x05,
	0xab, 0x16, 0x12, 0xd2, 0x03, 0xbb, 0x05, 0xd9, 0x72, 0x29, 0x01, 0x14, 0xb1, 0x82, 0x9c, 0x4a,
	0x25, 0x06, 0x27, 0xb6, 0x28, 0x96, 0xf4, 0xa2, 0x82, 0xe4, 0x24, 0x36, 0xb0, 0xab, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x29, 0x74, 0xc0, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// V1AuthClient is the client API for V1Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type V1AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type v1AuthClient struct {
	cc *grpc.ClientConn
}

func NewV1AuthClient(cc *grpc.ClientConn) V1AuthClient {
	return &v1AuthClient{cc}
}

func (c *v1AuthClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/auth.V1Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1AuthServer is the server API for V1Auth service.
type V1AuthServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

// UnimplementedV1AuthServer can be embedded to have forward compatible implementations.
type UnimplementedV1AuthServer struct {
}

func (*UnimplementedV1AuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterV1AuthServer(s *grpc.Server, srv V1AuthServer) {
	s.RegisterService(&_V1Auth_serviceDesc, srv)
}

func _V1Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.V1Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _V1Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.V1Auth",
	HandlerType: (*V1AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _V1Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
