// Code generated by protoc-gen-go. DO NOT EDIT.
// source: birthday.proto

package proto3

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

type Date struct {
	Day                  int64    `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Month                int64    `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Year                 int64    `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57480d75fb8881b, []int{0}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetDay() int64 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Date) GetMonth() int64 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetYear() int64 {
	if m != nil {
		return m.Year
	}
	return 0
}

type BirthdayStatus struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BirthdayStatus) Reset()         { *m = BirthdayStatus{} }
func (m *BirthdayStatus) String() string { return proto.CompactTextString(m) }
func (*BirthdayStatus) ProtoMessage()    {}
func (*BirthdayStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57480d75fb8881b, []int{1}
}

func (m *BirthdayStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BirthdayStatus.Unmarshal(m, b)
}
func (m *BirthdayStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BirthdayStatus.Marshal(b, m, deterministic)
}
func (m *BirthdayStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BirthdayStatus.Merge(m, src)
}
func (m *BirthdayStatus) XXX_Size() int {
	return xxx_messageInfo_BirthdayStatus.Size(m)
}
func (m *BirthdayStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BirthdayStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BirthdayStatus proto.InternalMessageInfo

func (m *BirthdayStatus) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *BirthdayStatus) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "proto3.Date")
	proto.RegisterType((*BirthdayStatus)(nil), "proto3.BirthdayStatus")
}

func init() { proto.RegisterFile("birthday.proto", fileDescriptor_b57480d75fb8881b) }

var fileDescriptor_b57480d75fb8881b = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xca, 0x2c, 0x2a,
	0xc9, 0x48, 0x49, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc6, 0x4a,
	0x4e, 0x5c, 0x2c, 0x2e, 0x89, 0x25, 0xa9, 0x42, 0x02, 0x5c, 0xcc, 0x29, 0x89, 0x95, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x6e, 0x7e, 0x5e, 0x49, 0x86,
	0x04, 0x13, 0x58, 0x0c, 0xc2, 0x11, 0x12, 0xe2, 0x62, 0xa9, 0x4c, 0x4d, 0x2c, 0x92, 0x60, 0x06,
	0x0b, 0x82, 0xd9, 0x4a, 0x56, 0x5c, 0x7c, 0x4e, 0x50, 0xd3, 0x83, 0x4b, 0x12, 0x4b, 0x4a, 0x8b,
	0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0xc1, 0x2c, 0xb0, 0x81, 0x1c, 0x41, 0x50, 0x1e, 0xc8, 0x96, 0xc4,
	0xf4, 0x54, 0xa8, 0x89, 0x20, 0xa6, 0x91, 0x07, 0x17, 0x3f, 0x4c, 0xaf, 0x73, 0x46, 0x6a, 0x72,
	0x76, 0x6a, 0x91, 0x90, 0x29, 0x17, 0x6f, 0x32, 0x88, 0x09, 0x13, 0x17, 0xe2, 0x81, 0xb8, 0xd9,
	0x58, 0x0f, 0xe4, 0x52, 0x29, 0x31, 0x18, 0x0f, 0xd5, 0xce, 0x24, 0xa8, 0x8f, 0x00, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0xa6, 0x4e, 0x49, 0xea, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BirthdayCheckerClient is the client API for BirthdayChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BirthdayCheckerClient interface {
	CheckBirthday(ctx context.Context, in *Date, opts ...grpc.CallOption) (*BirthdayStatus, error)
}

type birthdayCheckerClient struct {
	cc *grpc.ClientConn
}

func NewBirthdayCheckerClient(cc *grpc.ClientConn) BirthdayCheckerClient {
	return &birthdayCheckerClient{cc}
}

func (c *birthdayCheckerClient) CheckBirthday(ctx context.Context, in *Date, opts ...grpc.CallOption) (*BirthdayStatus, error) {
	out := new(BirthdayStatus)
	err := c.cc.Invoke(ctx, "/proto3.BirthdayChecker/checkBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirthdayCheckerServer is the server API for BirthdayChecker service.
type BirthdayCheckerServer interface {
	CheckBirthday(context.Context, *Date) (*BirthdayStatus, error)
}

// UnimplementedBirthdayCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedBirthdayCheckerServer struct {
}

func (*UnimplementedBirthdayCheckerServer) CheckBirthday(ctx context.Context, req *Date) (*BirthdayStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBirthday not implemented")
}

func RegisterBirthdayCheckerServer(s *grpc.Server, srv BirthdayCheckerServer) {
	s.RegisterService(&_BirthdayChecker_serviceDesc, srv)
}

func _BirthdayChecker_CheckBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayCheckerServer).CheckBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto3.BirthdayChecker/CheckBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayCheckerServer).CheckBirthday(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

var _BirthdayChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto3.BirthdayChecker",
	HandlerType: (*BirthdayCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkBirthday",
			Handler:    _BirthdayChecker_CheckBirthday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "birthday.proto",
}
