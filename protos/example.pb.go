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

type Data struct {
	Counter              int64    `protobuf:"varint,1,opt,name=Counter,proto3" json:"Counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
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

func (m *Data) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*Data)(nil), "example.Data")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x14, 0xb8,
	0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x24, 0xb8, 0xd8, 0x9d, 0xf3, 0x4b, 0xf3, 0x4a, 0x52, 0x8b,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x60, 0x5c, 0x23, 0x6b, 0x2e, 0x3e, 0xcf, 0x92, 0xd4,
	0xa2, 0xc4, 0x92, 0x54, 0xa8, 0x88, 0x90, 0x26, 0x17, 0x3b, 0x54, 0x44, 0x88, 0x57, 0x0f, 0x66,
	0x2e, 0xc8, 0x14, 0x29, 0x54, 0xae, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x3a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0x6e, 0x67, 0x91, 0x7f, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IterateCounterClient is the client API for IterateCounter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IterateCounterClient interface {
	Iterate(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
}

type iterateCounterClient struct {
	cc *grpc.ClientConn
}

func NewIterateCounterClient(cc *grpc.ClientConn) IterateCounterClient {
	return &iterateCounterClient{cc}
}

func (c *iterateCounterClient) Iterate(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/example.IterateCounter/Iterate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IterateCounterServer is the server API for IterateCounter service.
type IterateCounterServer interface {
	Iterate(context.Context, *Data) (*Data, error)
}

func RegisterIterateCounterServer(s *grpc.Server, srv IterateCounterServer) {
	s.RegisterService(&_IterateCounter_serviceDesc, srv)
}

func _IterateCounter_Iterate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IterateCounterServer).Iterate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.IterateCounter/Iterate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IterateCounterServer).Iterate(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _IterateCounter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.IterateCounter",
	HandlerType: (*IterateCounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Iterate",
			Handler:    _IterateCounter_Iterate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}