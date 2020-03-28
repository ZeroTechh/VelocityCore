// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email-srv.proto

package email_srv

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

type EmailData struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailData) Reset()         { *m = EmailData{} }
func (m *EmailData) String() string { return proto.CompactTextString(m) }
func (*EmailData) ProtoMessage()    {}
func (*EmailData) Descriptor() ([]byte, []int) {
	return fileDescriptor_47cb0f489c1ca7ac, []int{0}
}

func (m *EmailData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailData.Unmarshal(m, b)
}
func (m *EmailData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailData.Marshal(b, m, deterministic)
}
func (m *EmailData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailData.Merge(m, src)
}
func (m *EmailData) XXX_Size() int {
	return xxx_messageInfo_EmailData.Size(m)
}
func (m *EmailData) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailData.DiscardUnknown(m)
}

var xxx_messageInfo_EmailData proto.InternalMessageInfo

func (m *EmailData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailData) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_47cb0f489c1ca7ac, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmailData)(nil), "EmailData")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("email-srv.proto", fileDescriptor_47cb0f489c1ca7ac) }

var fileDescriptor_47cb0f489c1ca7ac = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x2d, 0x2e, 0x2a, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe5, 0xe2, 0x74,
	0x05, 0x09, 0xb9, 0x24, 0x96, 0x24, 0x0a, 0x89, 0x70, 0xb1, 0x82, 0xe5, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x26, 0xb0,
	0x20, 0x98, 0xad, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0x69, 0x64, 0x00, 0x62, 0x80,
	0x54, 0xa9, 0x73, 0xf1, 0x07, 0xa7, 0xe6, 0xa5, 0x04, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0x42, 0x84,
	0xb8, 0xf4, 0xe0, 0x46, 0x4b, 0xb1, 0xe9, 0x81, 0xd5, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x2d, 0x36,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x4e, 0xb1, 0xa7, 0x8b, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClient interface {
	SendSimpleEmail(ctx context.Context, in *EmailData, opts ...grpc.CallOption) (*Empty, error)
}

type emailClient struct {
	cc *grpc.ClientConn
}

func NewEmailClient(cc *grpc.ClientConn) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) SendSimpleEmail(ctx context.Context, in *EmailData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Email/SendSimpleEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
type EmailServer interface {
	SendSimpleEmail(context.Context, *EmailData) (*Empty, error)
}

// UnimplementedEmailServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServer struct {
}

func (*UnimplementedEmailServer) SendSimpleEmail(ctx context.Context, req *EmailData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSimpleEmail not implemented")
}

func RegisterEmailServer(s *grpc.Server, srv EmailServer) {
	s.RegisterService(&_Email_serviceDesc, srv)
}

func _Email_SendSimpleEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).SendSimpleEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Email/SendSimpleEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).SendSimpleEmail(ctx, req.(*EmailData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Email_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSimpleEmail",
			Handler:    _Email_SendSimpleEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email-srv.proto",
}
