// Code generated by protoc-gen-go. DO NOT EDIT.
// source: venus.proto

package venus

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bbe65601b4a05ea, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bbe65601b4a05ea, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NotifyPhoneCallStatusRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Callee               string   `protobuf:"bytes,2,opt,name=callee,proto3" json:"callee,omitempty"`
	Caller               string   `protobuf:"bytes,3,opt,name=caller,proto3" json:"caller,omitempty"`
	RecordUrl            string   `protobuf:"bytes,4,opt,name=record_url,json=recordUrl,proto3" json:"record_url,omitempty"`
	Duration             int64    `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	StartAtSec           int64    `protobuf:"varint,6,opt,name=start_at_sec,json=startAtSec,proto3" json:"start_at_sec,omitempty"`
	EndAtSec             int64    `protobuf:"varint,7,opt,name=end_at_sec,json=endAtSec,proto3" json:"end_at_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyPhoneCallStatusRequest) Reset()         { *m = NotifyPhoneCallStatusRequest{} }
func (m *NotifyPhoneCallStatusRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyPhoneCallStatusRequest) ProtoMessage()    {}
func (*NotifyPhoneCallStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bbe65601b4a05ea, []int{2}
}

func (m *NotifyPhoneCallStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyPhoneCallStatusRequest.Unmarshal(m, b)
}
func (m *NotifyPhoneCallStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyPhoneCallStatusRequest.Marshal(b, m, deterministic)
}
func (m *NotifyPhoneCallStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyPhoneCallStatusRequest.Merge(m, src)
}
func (m *NotifyPhoneCallStatusRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyPhoneCallStatusRequest.Size(m)
}
func (m *NotifyPhoneCallStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyPhoneCallStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyPhoneCallStatusRequest proto.InternalMessageInfo

func (m *NotifyPhoneCallStatusRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *NotifyPhoneCallStatusRequest) GetCallee() string {
	if m != nil {
		return m.Callee
	}
	return ""
}

func (m *NotifyPhoneCallStatusRequest) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *NotifyPhoneCallStatusRequest) GetRecordUrl() string {
	if m != nil {
		return m.RecordUrl
	}
	return ""
}

func (m *NotifyPhoneCallStatusRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *NotifyPhoneCallStatusRequest) GetStartAtSec() int64 {
	if m != nil {
		return m.StartAtSec
	}
	return 0
}

func (m *NotifyPhoneCallStatusRequest) GetEndAtSec() int64 {
	if m != nil {
		return m.EndAtSec
	}
	return 0
}

type NotifyPhoneCallStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyPhoneCallStatusResponse) Reset()         { *m = NotifyPhoneCallStatusResponse{} }
func (m *NotifyPhoneCallStatusResponse) String() string { return proto.CompactTextString(m) }
func (*NotifyPhoneCallStatusResponse) ProtoMessage()    {}
func (*NotifyPhoneCallStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bbe65601b4a05ea, []int{3}
}

func (m *NotifyPhoneCallStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyPhoneCallStatusResponse.Unmarshal(m, b)
}
func (m *NotifyPhoneCallStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyPhoneCallStatusResponse.Marshal(b, m, deterministic)
}
func (m *NotifyPhoneCallStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyPhoneCallStatusResponse.Merge(m, src)
}
func (m *NotifyPhoneCallStatusResponse) XXX_Size() int {
	return xxx_messageInfo_NotifyPhoneCallStatusResponse.Size(m)
}
func (m *NotifyPhoneCallStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyPhoneCallStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyPhoneCallStatusResponse proto.InternalMessageInfo

type MakePhoneCallRequest struct {
	Callee               string   `protobuf:"bytes,1,opt,name=callee,proto3" json:"callee,omitempty"`
	Caller               string   `protobuf:"bytes,2,opt,name=caller,proto3" json:"caller,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakePhoneCallRequest) Reset()         { *m = MakePhoneCallRequest{} }
func (m *MakePhoneCallRequest) String() string { return proto.CompactTextString(m) }
func (*MakePhoneCallRequest) ProtoMessage()    {}
func (*MakePhoneCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bbe65601b4a05ea, []int{4}
}

func (m *MakePhoneCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakePhoneCallRequest.Unmarshal(m, b)
}
func (m *MakePhoneCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakePhoneCallRequest.Marshal(b, m, deterministic)
}
func (m *MakePhoneCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakePhoneCallRequest.Merge(m, src)
}
func (m *MakePhoneCallRequest) XXX_Size() int {
	return xxx_messageInfo_MakePhoneCallRequest.Size(m)
}
func (m *MakePhoneCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakePhoneCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakePhoneCallRequest proto.InternalMessageInfo

func (m *MakePhoneCallRequest) GetCallee() string {
	if m != nil {
		return m.Callee
	}
	return ""
}

func (m *MakePhoneCallRequest) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

type MakePhoneCallResponse struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakePhoneCallResponse) Reset()         { *m = MakePhoneCallResponse{} }
func (m *MakePhoneCallResponse) String() string { return proto.CompactTextString(m) }
func (*MakePhoneCallResponse) ProtoMessage()    {}
func (*MakePhoneCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bbe65601b4a05ea, []int{5}
}

func (m *MakePhoneCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakePhoneCallResponse.Unmarshal(m, b)
}
func (m *MakePhoneCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakePhoneCallResponse.Marshal(b, m, deterministic)
}
func (m *MakePhoneCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakePhoneCallResponse.Merge(m, src)
}
func (m *MakePhoneCallResponse) XXX_Size() int {
	return xxx_messageInfo_MakePhoneCallResponse.Size(m)
}
func (m *MakePhoneCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MakePhoneCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MakePhoneCallResponse proto.InternalMessageInfo

func (m *MakePhoneCallResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *MakePhoneCallResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "venus.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "venus.HelloResponse")
	proto.RegisterType((*NotifyPhoneCallStatusRequest)(nil), "venus.NotifyPhoneCallStatusRequest")
	proto.RegisterType((*NotifyPhoneCallStatusResponse)(nil), "venus.NotifyPhoneCallStatusResponse")
	proto.RegisterType((*MakePhoneCallRequest)(nil), "venus.MakePhoneCallRequest")
	proto.RegisterType((*MakePhoneCallResponse)(nil), "venus.MakePhoneCallResponse")
}

func init() { proto.RegisterFile("venus.proto", fileDescriptor_4bbe65601b4a05ea) }

var fileDescriptor_4bbe65601b4a05ea = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x4d, 0x1f, 0x1f, 0x0f, 0x2e, 0xb0, 0x99, 0x07, 0x2f, 0x4d, 0x85, 0x48, 0x8a, 0x0b, 0x56,
	0x98, 0xe0, 0x2f, 0x50, 0x13, 0xa3, 0x26, 0x12, 0x53, 0xa2, 0x0b, 0x37, 0xcd, 0xd0, 0xb9, 0xc6,
	0x86, 0x71, 0x06, 0x67, 0xa6, 0x24, 0xfe, 0x63, 0xd7, 0xfe, 0x02, 0xc3, 0x74, 0x40, 0x30, 0x05,
	0x77, 0xbd, 0xe7, 0x9c, 0x7b, 0xa6, 0xe7, 0xde, 0x0b, 0x8d, 0x25, 0x8a, 0x4c, 0x8f, 0x16, 0x4a,
	0x1a, 0x49, 0x2a, 0xb6, 0x08, 0x43, 0x68, 0x5e, 0x23, 0xe7, 0x32, 0xc2, 0xb7, 0x0c, 0xb5, 0x21,
	0x04, 0xca, 0x82, 0xbe, 0xa2, 0xef, 0xf5, 0xbd, 0x61, 0x3d, 0xb2, 0xdf, 0xe1, 0x00, 0x5a, 0x4e,
	0xa3, 0x17, 0x52, 0x68, 0x2c, 0x14, 0x7d, 0x78, 0xd0, 0x9d, 0x48, 0x93, 0x3e, 0xbf, 0xdf, 0xbf,
	0x48, 0x81, 0x97, 0x94, 0xf3, 0xa9, 0xa1, 0x26, 0xd3, 0x6b, 0xe7, 0x1e, 0x80, 0x46, 0xad, 0x53,
	0x29, 0xe2, 0x94, 0xb9, 0xd6, 0xba, 0x43, 0x6e, 0x18, 0xf9, 0x0f, 0xd5, 0x84, 0x72, 0x8e, 0xe8,
	0xff, 0xb1, 0x94, 0xab, 0x36, 0xb8, 0xf2, 0x4b, 0x5b, 0xb8, 0x5a, 0xd9, 0x29, 0x4c, 0xa4, 0x62,
	0x71, 0xa6, 0xb8, 0x5f, 0xce, 0xed, 0x72, 0xe4, 0x41, 0x71, 0x12, 0x40, 0x8d, 0x65, 0x8a, 0x9a,
	0x54, 0x0a, 0xbf, 0xd2, 0xf7, 0x86, 0xa5, 0x68, 0x53, 0x93, 0x3e, 0x34, 0xb5, 0xa1, 0xca, 0xc4,
	0xd4, 0xc4, 0x1a, 0x13, 0xbf, 0x6a, 0x79, 0xb0, 0xd8, 0xb9, 0x99, 0x62, 0x42, 0xba, 0x00, 0x28,
	0xd8, 0x9a, 0xff, 0x9b, 0xf7, 0xa3, 0x60, 0x96, 0x0d, 0x8f, 0xa1, 0xb7, 0x27, 0x69, 0x3e, 0x9f,
	0xf0, 0x0a, 0xda, 0x77, 0x74, 0x8e, 0x1b, 0x7a, 0x3d, 0x82, 0xef, 0x8c, 0xde, 0x9e, 0x8c, 0xdb,
	0xd9, 0x55, 0x38, 0x81, 0xce, 0x0f, 0x1f, 0xb7, 0x80, 0xdf, 0x67, 0xa9, 0xed, 0x1f, 0x59, 0xbf,
	0x5a, 0xe4, 0xaa, 0xf1, 0xa7, 0x07, 0xcd, 0xc7, 0xd5, 0xda, 0xa7, 0xa8, 0x96, 0x69, 0x82, 0x64,
	0x0c, 0x15, 0xbb, 0x59, 0xf2, 0x6f, 0x94, 0xdf, 0xc6, 0xf6, 0x2d, 0x04, 0xed, 0x5d, 0xd0, 0xbd,
	0x3d, 0x83, 0x4e, 0x61, 0x7a, 0x32, 0x70, 0xf2, 0x43, 0x57, 0x10, 0x9c, 0x1c, 0x16, 0xb9, 0x37,
	0x6e, 0xa1, 0xb5, 0x13, 0x9c, 0x1c, 0xb9, 0xb6, 0xa2, 0xb1, 0x06, 0xdd, 0x62, 0x32, 0xf7, 0xba,
	0x68, 0x3c, 0xd5, 0xe7, 0x6c, 0x76, 0x6a, 0x25, 0xb3, 0xaa, 0x3d, 0xfe, 0xb3, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x07, 0x27, 0x96, 0x2e, 0x0b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VenusServiceClient is the client API for VenusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VenusServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	NotifyPhoneCallStatus(ctx context.Context, in *NotifyPhoneCallStatusRequest, opts ...grpc.CallOption) (*NotifyPhoneCallStatusResponse, error)
	MakePhoneCall(ctx context.Context, in *MakePhoneCallRequest, opts ...grpc.CallOption) (*MakePhoneCallResponse, error)
}

type venusServiceClient struct {
	cc *grpc.ClientConn
}

func NewVenusServiceClient(cc *grpc.ClientConn) VenusServiceClient {
	return &venusServiceClient{cc}
}

func (c *venusServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/venus.VenusService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *venusServiceClient) NotifyPhoneCallStatus(ctx context.Context, in *NotifyPhoneCallStatusRequest, opts ...grpc.CallOption) (*NotifyPhoneCallStatusResponse, error) {
	out := new(NotifyPhoneCallStatusResponse)
	err := c.cc.Invoke(ctx, "/venus.VenusService/NotifyPhoneCallStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *venusServiceClient) MakePhoneCall(ctx context.Context, in *MakePhoneCallRequest, opts ...grpc.CallOption) (*MakePhoneCallResponse, error) {
	out := new(MakePhoneCallResponse)
	err := c.cc.Invoke(ctx, "/venus.VenusService/MakePhoneCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VenusServiceServer is the server API for VenusService service.
type VenusServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	NotifyPhoneCallStatus(context.Context, *NotifyPhoneCallStatusRequest) (*NotifyPhoneCallStatusResponse, error)
	MakePhoneCall(context.Context, *MakePhoneCallRequest) (*MakePhoneCallResponse, error)
}

// UnimplementedVenusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVenusServiceServer struct {
}

func (*UnimplementedVenusServiceServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedVenusServiceServer) NotifyPhoneCallStatus(ctx context.Context, req *NotifyPhoneCallStatusRequest) (*NotifyPhoneCallStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyPhoneCallStatus not implemented")
}
func (*UnimplementedVenusServiceServer) MakePhoneCall(ctx context.Context, req *MakePhoneCallRequest) (*MakePhoneCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePhoneCall not implemented")
}

func RegisterVenusServiceServer(s *grpc.Server, srv VenusServiceServer) {
	s.RegisterService(&_VenusService_serviceDesc, srv)
}

func _VenusService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VenusServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/venus.VenusService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VenusServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VenusService_NotifyPhoneCallStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyPhoneCallStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VenusServiceServer).NotifyPhoneCallStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/venus.VenusService/NotifyPhoneCallStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VenusServiceServer).NotifyPhoneCallStatus(ctx, req.(*NotifyPhoneCallStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VenusService_MakePhoneCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakePhoneCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VenusServiceServer).MakePhoneCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/venus.VenusService/MakePhoneCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VenusServiceServer).MakePhoneCall(ctx, req.(*MakePhoneCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VenusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "venus.VenusService",
	HandlerType: (*VenusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _VenusService_Hello_Handler,
		},
		{
			MethodName: "NotifyPhoneCallStatus",
			Handler:    _VenusService_NotifyPhoneCallStatus_Handler,
		},
		{
			MethodName: "MakePhoneCall",
			Handler:    _VenusService_MakePhoneCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "venus.proto",
}