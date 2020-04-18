// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

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

type Data struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	AccountStatus        string   `protobuf:"bytes,2,opt,name=accountStatus,proto3" json:"accountStatus,omitempty"`
	AccountCreationUTC   int64    `protobuf:"varint,3,opt,name=accountCreationUTC,proto3" json:"accountCreationUTC,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Data) GetAccountStatus() string {
	if m != nil {
		return m.AccountStatus
	}
	return ""
}

func (m *Data) GetAccountCreationUTC() int64 {
	if m != nil {
		return m.AccountCreationUTC
	}
	return 0
}

type Identifier struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Message struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "Data")
	proto.RegisterType((*Identifier)(nil), "Identifier")
	proto.RegisterType((*Message)(nil), "Message")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x1b, 0xa3, 0xdb, 0x75, 0xa4, 0x97, 0x39, 0x48, 0x50, 0x90, 0x9a, 0xf6, 0xd0, 0x53,
	0x0e, 0xfa, 0x04, 0xa5, 0x0b, 0xd2, 0x43, 0x2f, 0xd5, 0x3e, 0xc0, 0x98, 0x1d, 0x25, 0xa0, 0x1b,
	0x49, 0x66, 0xfb, 0x1a, 0xbe, 0xb2, 0x58, 0x23, 0x52, 0xb0, 0xbd, 0xcd, 0xf7, 0x87, 0xdf, 0xc0,
	0x07, 0xa3, 0xcc, 0x69, 0x1b, 0x3c, 0xbb, 0x8f, 0x14, 0x25, 0x5a, 0x81, 0xd3, 0x86, 0x84, 0xf0,
	0x12, 0xaa, 0x3e, 0x73, 0x5a, 0x36, 0x46, 0x8d, 0xd5, 0xec, 0x7c, 0x5d, 0x14, 0x4e, 0x61, 0x44,
	0xde, 0xc7, 0xbe, 0x93, 0x47, 0x21, 0xe9, 0xb3, 0x39, 0xd9, 0xc5, 0xfb, 0x26, 0x3a, 0xc0, 0x62,
	0x2c, 0x12, 0x93, 0x84, 0xd8, 0x6d, 0x9e, 0x16, 0x46, 0x8f, 0xd5, 0x4c, 0xaf, 0xff, 0x49, 0xec,
	0x14, 0x60, 0xd9, 0x72, 0x27, 0xe1, 0x25, 0x70, 0x3a, 0xf4, 0xdb, 0x4e, 0x60, 0xb8, 0xe2, 0x9c,
	0xe9, 0x95, 0xd1, 0xc0, 0xf0, 0xfd, 0xe7, 0x2c, 0x9d, 0x5f, 0x79, 0xf7, 0xa9, 0xa0, 0xde, 0x64,
	0x4e, 0x2b, 0x16, 0xc2, 0x1b, 0xd0, 0xf3, 0xb6, 0xc5, 0x0b, 0xf7, 0x47, 0xbf, 0xaa, 0x5d, 0x81,
	0xd8, 0x01, 0x5e, 0x83, 0x7e, 0x60, 0xd9, 0xcf, 0xcf, 0xdc, 0xf7, 0x00, 0x76, 0x80, 0x13, 0xa8,
	0xe7, 0x5e, 0xc2, 0x96, 0x84, 0x0f, 0x13, 0x6e, 0xa1, 0x6a, 0xf8, 0x8d, 0x8f, 0x54, 0x9e, 0xab,
	0xdd, 0xb2, 0xf7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x5a, 0x99, 0xd2, 0x6a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserMetaClient is the client API for UserMeta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserMetaClient interface {
	Add(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Message, error)
	Get(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Data, error)
	Activate(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Message, error)
	Delete(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Message, error)
}

type userMetaClient struct {
	cc *grpc.ClientConn
}

func NewUserMetaClient(cc *grpc.ClientConn) UserMetaClient {
	return &userMetaClient{cc}
}

func (c *userMetaClient) Add(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/UserMeta/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMetaClient) Get(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/UserMeta/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMetaClient) Activate(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/UserMeta/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMetaClient) Delete(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/UserMeta/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMetaServer is the server API for UserMeta service.
type UserMetaServer interface {
	Add(context.Context, *Identifier) (*Message, error)
	Get(context.Context, *Identifier) (*Data, error)
	Activate(context.Context, *Identifier) (*Message, error)
	Delete(context.Context, *Identifier) (*Message, error)
}

// UnimplementedUserMetaServer can be embedded to have forward compatible implementations.
type UnimplementedUserMetaServer struct {
}

func (*UnimplementedUserMetaServer) Add(ctx context.Context, req *Identifier) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedUserMetaServer) Get(ctx context.Context, req *Identifier) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserMetaServer) Activate(ctx context.Context, req *Identifier) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (*UnimplementedUserMetaServer) Delete(ctx context.Context, req *Identifier) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserMetaServer(s *grpc.Server, srv UserMetaServer) {
	s.RegisterService(&_UserMeta_serviceDesc, srv)
}

func _UserMeta_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMetaServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserMeta/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMetaServer).Add(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMeta_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMetaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserMeta/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMetaServer).Get(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMeta_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMetaServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserMeta/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMetaServer).Activate(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMeta_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMetaServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserMeta/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMetaServer).Delete(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserMeta_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserMeta",
	HandlerType: (*UserMetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserMeta_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserMeta_Get_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _UserMeta_Activate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserMeta_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}