// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IntMsg struct {
	IntValue             int64    `protobuf:"varint,1,opt,name=IntValue,proto3" json:"IntValue,omitempty"`
	IntMultiple          int64    `protobuf:"varint,2,opt,name=IntMultiple,proto3" json:"IntMultiple,omitempty"`
	IntCalc              int64    `protobuf:"varint,3,opt,name=IntCalc,proto3" json:"IntCalc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntMsg) Reset()         { *m = IntMsg{} }
func (m *IntMsg) String() string { return proto.CompactTextString(m) }
func (*IntMsg) ProtoMessage()    {}
func (*IntMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *IntMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMsg.Unmarshal(m, b)
}
func (m *IntMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMsg.Marshal(b, m, deterministic)
}
func (m *IntMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMsg.Merge(m, src)
}
func (m *IntMsg) XXX_Size() int {
	return xxx_messageInfo_IntMsg.Size(m)
}
func (m *IntMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMsg.DiscardUnknown(m)
}

var xxx_messageInfo_IntMsg proto.InternalMessageInfo

func (m *IntMsg) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *IntMsg) GetIntMultiple() int64 {
	if m != nil {
		return m.IntMultiple
	}
	return 0
}

func (m *IntMsg) GetIntCalc() int64 {
	if m != nil {
		return m.IntCalc
	}
	return 0
}

func init() {
	proto.RegisterType((*IntMsg)(nil), "example.IntMsg")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x12, 0xb8,
	0xd8, 0x3c, 0xf3, 0x4a, 0x7c, 0x8b, 0xd3, 0x85, 0xa4, 0xb8, 0x38, 0x3c, 0xf3, 0x4a, 0xc2, 0x12,
	0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xe0, 0x7c, 0x21, 0x05, 0x2e, 0x6e,
	0x90, 0xaa, 0xd2, 0x9c, 0x92, 0xcc, 0x82, 0x9c, 0x54, 0x09, 0x26, 0xb0, 0x34, 0xb2, 0x90, 0x90,
	0x04, 0x17, 0xbb, 0x67, 0x5e, 0x89, 0x73, 0x62, 0x4e, 0xb2, 0x04, 0x33, 0x58, 0x16, 0xc6, 0x35,
	0x72, 0xe6, 0xe2, 0x75, 0xca, 0x74, 0xc9, 0x2c, 0x4a, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x4b, 0xcc,
	0x11, 0x32, 0xe2, 0xe2, 0x80, 0x6b, 0xe3, 0xd7, 0x83, 0xb9, 0x0b, 0xe2, 0x0a, 0x29, 0x74, 0x01,
	0x25, 0x06, 0x0d, 0x46, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x72, 0xdb, 0x7a, 0xc7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BiDirectionalClient is the client API for BiDirectional service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BiDirectionalClient interface {
	Multiple(ctx context.Context, opts ...grpc.CallOption) (BiDirectional_MultipleClient, error)
}

type biDirectionalClient struct {
	cc *grpc.ClientConn
}

func NewBiDirectionalClient(cc *grpc.ClientConn) BiDirectionalClient {
	return &biDirectionalClient{cc}
}

func (c *biDirectionalClient) Multiple(ctx context.Context, opts ...grpc.CallOption) (BiDirectional_MultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BiDirectional_serviceDesc.Streams[0], "/example.BiDirectional/Multiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &biDirectionalMultipleClient{stream}
	return x, nil
}

type BiDirectional_MultipleClient interface {
	Send(*IntMsg) error
	Recv() (*IntMsg, error)
	grpc.ClientStream
}

type biDirectionalMultipleClient struct {
	grpc.ClientStream
}

func (x *biDirectionalMultipleClient) Send(m *IntMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *biDirectionalMultipleClient) Recv() (*IntMsg, error) {
	m := new(IntMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BiDirectionalServer is the server API for BiDirectional service.
type BiDirectionalServer interface {
	Multiple(BiDirectional_MultipleServer) error
}

func RegisterBiDirectionalServer(s *grpc.Server, srv BiDirectionalServer) {
	s.RegisterService(&_BiDirectional_serviceDesc, srv)
}

func _BiDirectional_Multiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BiDirectionalServer).Multiple(&biDirectionalMultipleServer{stream})
}

type BiDirectional_MultipleServer interface {
	Send(*IntMsg) error
	Recv() (*IntMsg, error)
	grpc.ServerStream
}

type biDirectionalMultipleServer struct {
	grpc.ServerStream
}

func (x *biDirectionalMultipleServer) Send(m *IntMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *biDirectionalMultipleServer) Recv() (*IntMsg, error) {
	m := new(IntMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BiDirectional_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.BiDirectional",
	HandlerType: (*BiDirectionalServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Multiple",
			Handler:       _BiDirectional_Multiple_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}
