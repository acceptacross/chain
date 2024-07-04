// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tunnel/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryTunnelRequest is the request type for the Query/Tunnel RPC method.
type QueryTunnelsRequest struct {
	// is_active is a flag to filter active tunnels.
	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryTunnelsRequest) Reset()         { *m = QueryTunnelsRequest{} }
func (m *QueryTunnelsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTunnelsRequest) ProtoMessage()    {}
func (*QueryTunnelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_220124e1821d297e, []int{0}
}
func (m *QueryTunnelsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTunnelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTunnelsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTunnelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTunnelsRequest.Merge(m, src)
}
func (m *QueryTunnelsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTunnelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTunnelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTunnelsRequest proto.InternalMessageInfo

func (m *QueryTunnelsRequest) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *QueryTunnelsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryTunnelResponse is the response type for the Query/Tunnel RPC method.
type QueryTunnelsResponse struct {
	// Tunnels is a list of tunnels.
	Tunnels []*Tunnel `protobuf:"bytes,1,rep,name=tunnels,proto3" json:"tunnels,omitempty"`
	// pagination defines an optional pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryTunnelsResponse) Reset()         { *m = QueryTunnelsResponse{} }
func (m *QueryTunnelsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTunnelsResponse) ProtoMessage()    {}
func (*QueryTunnelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_220124e1821d297e, []int{1}
}
func (m *QueryTunnelsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTunnelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTunnelsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTunnelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTunnelsResponse.Merge(m, src)
}
func (m *QueryTunnelsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTunnelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTunnelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTunnelsResponse proto.InternalMessageInfo

func (m *QueryTunnelsResponse) GetTunnels() []*Tunnel {
	if m != nil {
		return m.Tunnels
	}
	return nil
}

func (m *QueryTunnelsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryTunnelRequest is the request type for the Query/Tunnel RPC method.
type QueryTunnelRequest struct {
	// TunnelID is the ID of the tunnel to query.
	TunnelId uint64 `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
}

func (m *QueryTunnelRequest) Reset()         { *m = QueryTunnelRequest{} }
func (m *QueryTunnelRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTunnelRequest) ProtoMessage()    {}
func (*QueryTunnelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_220124e1821d297e, []int{2}
}
func (m *QueryTunnelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTunnelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTunnelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTunnelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTunnelRequest.Merge(m, src)
}
func (m *QueryTunnelRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTunnelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTunnelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTunnelRequest proto.InternalMessageInfo

func (m *QueryTunnelRequest) GetTunnelId() uint64 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

// QueryTunnelResponse is the response type for the Query/Tunnel RPC method.
type QueryTunnelResponse struct {
	// Tunnel is the tunnel with the given ID.
	Tunnel Tunnel `protobuf:"bytes,1,opt,name=tunnel,proto3" json:"tunnel"`
}

func (m *QueryTunnelResponse) Reset()         { *m = QueryTunnelResponse{} }
func (m *QueryTunnelResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTunnelResponse) ProtoMessage()    {}
func (*QueryTunnelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_220124e1821d297e, []int{3}
}
func (m *QueryTunnelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTunnelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTunnelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTunnelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTunnelResponse.Merge(m, src)
}
func (m *QueryTunnelResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTunnelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTunnelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTunnelResponse proto.InternalMessageInfo

func (m *QueryTunnelResponse) GetTunnel() Tunnel {
	if m != nil {
		return m.Tunnel
	}
	return Tunnel{}
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_220124e1821d297e, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// Params is the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_220124e1821d297e, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryTunnelsRequest)(nil), "tunnel.v1beta1.QueryTunnelsRequest")
	proto.RegisterType((*QueryTunnelsResponse)(nil), "tunnel.v1beta1.QueryTunnelsResponse")
	proto.RegisterType((*QueryTunnelRequest)(nil), "tunnel.v1beta1.QueryTunnelRequest")
	proto.RegisterType((*QueryTunnelResponse)(nil), "tunnel.v1beta1.QueryTunnelResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "tunnel.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "tunnel.v1beta1.QueryParamsResponse")
}

func init() { proto.RegisterFile("tunnel/v1beta1/query.proto", fileDescriptor_220124e1821d297e) }

var fileDescriptor_220124e1821d297e = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0xe3, 0x50, 0xd2, 0x30, 0x95, 0x58, 0x0c, 0x51, 0x09, 0x2e, 0x72, 0x23, 0xb7, 0x82,
	0xa8, 0x8b, 0x19, 0x12, 0xb8, 0x00, 0x5d, 0x80, 0x2a, 0x36, 0xc5, 0x62, 0xc5, 0xa6, 0x1a, 0x3b,
	0x23, 0x77, 0xa4, 0x64, 0xc6, 0xc9, 0x8c, 0x23, 0x0a, 0xea, 0x86, 0x13, 0x80, 0xb8, 0x54, 0x97,
	0x95, 0xd8, 0xb0, 0x42, 0x28, 0xe1, 0x00, 0x1c, 0x01, 0x65, 0xde, 0x24, 0x8d, 0x8b, 0x1d, 0x76,
	0xc9, 0xbc, 0xef, 0xbd, 0xdf, 0xf7, 0xfe, 0x18, 0xf9, 0x26, 0x97, 0x92, 0x0f, 0xe9, 0xb4, 0x17,
	0x73, 0xc3, 0x7a, 0x74, 0x9c, 0xf3, 0xc9, 0x05, 0xc9, 0x26, 0xca, 0x28, 0x7c, 0x1f, 0x62, 0xc4,
	0xc5, 0xfc, 0x56, 0xaa, 0x52, 0x65, 0x43, 0x74, 0xf1, 0x0b, 0x54, 0xfe, 0x51, 0xa2, 0xf4, 0x48,
	0x69, 0x1a, 0x33, 0xcd, 0x21, 0x7d, 0x55, 0x2c, 0x63, 0xa9, 0x90, 0xcc, 0x08, 0x25, 0x9d, 0xf6,
	0x71, 0xaa, 0x54, 0x3a, 0xe4, 0x94, 0x65, 0x82, 0x32, 0x29, 0x95, 0xb1, 0x41, 0xed, 0xa2, 0x7b,
	0xb7, 0xbc, 0x64, 0x6c, 0xc2, 0x46, 0x55, 0x41, 0xe7, 0xcd, 0x06, 0xc3, 0x8f, 0xe8, 0xc1, 0xdb,
	0x05, 0xf9, 0x9d, 0x7d, 0xd4, 0x11, 0x1f, 0xe7, 0x5c, 0x1b, 0xbc, 0x87, 0xee, 0x09, 0x7d, 0xc6,
	0x12, 0x23, 0xa6, 0xbc, 0xed, 0x75, 0xbc, 0x6e, 0x33, 0x6a, 0x0a, 0xfd, 0xd2, 0xfe, 0xc7, 0xaf,
	0x10, 0xba, 0xf1, 0xd7, 0xae, 0x77, 0xbc, 0xee, 0x4e, 0xff, 0x09, 0x81, 0x66, 0xc8, 0xa2, 0x19,
	0x02, 0xb3, 0x70, 0x40, 0x72, 0xca, 0x52, 0xee, 0x0a, 0x47, 0x6b, 0x99, 0xe1, 0x57, 0x0f, 0xb5,
	0x8a, 0x70, 0x9d, 0x29, 0xa9, 0x39, 0x7e, 0x86, 0xb6, 0xc1, 0xa4, 0x6e, 0x7b, 0x9d, 0x3b, 0xdd,
	0x9d, 0xfe, 0x2e, 0x29, 0x0e, 0x94, 0x40, 0x46, 0xb4, 0x94, 0xe1, 0xd7, 0x25, 0x96, 0x9e, 0xfe,
	0xd7, 0x12, 0xe0, 0x0a, 0x9e, 0x7a, 0x08, 0xaf, 0x59, 0x5a, 0x1b, 0x07, 0x90, 0xce, 0xc4, 0xc0,
	0x8e, 0x63, 0x2b, 0x6a, 0xc2, 0xc3, 0xc9, 0x20, 0x7c, 0x53, 0x18, 0xe1, 0xaa, 0x89, 0x17, 0xa8,
	0x01, 0x12, 0x9b, 0x50, 0xd9, 0xc3, 0xf1, 0xd6, 0xd5, 0xcf, 0xfd, 0x5a, 0xe4, 0xb4, 0x61, 0xcb,
	0xf1, 0x4f, 0xed, 0x06, 0x1d, 0x7f, 0x85, 0x58, 0xbe, 0xde, 0x20, 0x60, 0xd3, 0x55, 0x08, 0xd0,
	0x2f, 0x11, 0xa0, 0xed, 0xff, 0xa9, 0xa3, 0xbb, 0xb6, 0x1a, 0x1e, 0xa3, 0x06, 0x28, 0x70, 0x78,
	0x3b, 0xf3, 0x5f, 0x13, 0xfe, 0xc1, 0x46, 0x0d, 0x58, 0x0a, 0x83, 0xcf, 0xdf, 0x7f, 0x7f, 0xab,
	0xb7, 0xf1, 0x2e, 0x2d, 0x3d, 0x49, 0x9c, 0xa3, 0x6d, 0xb7, 0x6d, 0x5c, 0x5e, 0xaf, 0x78, 0x88,
	0xfe, 0xe1, 0x66, 0x91, 0xa3, 0xee, 0x5b, 0xea, 0x23, 0xfc, 0x90, 0x96, 0xde, 0xba, 0xc6, 0x97,
	0xa8, 0x01, 0x39, 0x15, 0x9d, 0x16, 0xd6, 0xed, 0x1f, 0x6c, 0xd4, 0x38, 0xe6, 0x91, 0x65, 0x1e,
	0xe2, 0xb0, 0x9c, 0x49, 0x3f, 0xad, 0x2e, 0xe6, 0xf2, 0xf8, 0xe4, 0x6a, 0x16, 0x78, 0xd7, 0xb3,
	0xc0, 0xfb, 0x35, 0x0b, 0xbc, 0x2f, 0xf3, 0xa0, 0x76, 0x3d, 0x0f, 0x6a, 0x3f, 0xe6, 0x41, 0xed,
	0x3d, 0x4d, 0x85, 0x39, 0xcf, 0x63, 0x92, 0xa8, 0x11, 0x8d, 0x99, 0x1c, 0xd8, 0xaf, 0x32, 0x51,
	0x43, 0x9a, 0x9c, 0x33, 0x21, 0xe9, 0xb4, 0x4f, 0x3f, 0x2c, 0x0b, 0x9a, 0x8b, 0x8c, 0xeb, 0xb8,
	0x61, 0x15, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x99, 0x1f, 0x0e, 0x7f, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params is a RPC method that returns all parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Tunnels is a RPC method that returns all tunnels.
	Tunnels(ctx context.Context, in *QueryTunnelsRequest, opts ...grpc.CallOption) (*QueryTunnelsResponse, error)
	// Tunnel is a RPC method that returns a tunnel by its ID.
	Tunnel(ctx context.Context, in *QueryTunnelRequest, opts ...grpc.CallOption) (*QueryTunnelResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/tunnel.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Tunnels(ctx context.Context, in *QueryTunnelsRequest, opts ...grpc.CallOption) (*QueryTunnelsResponse, error) {
	out := new(QueryTunnelsResponse)
	err := c.cc.Invoke(ctx, "/tunnel.v1beta1.Query/Tunnels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Tunnel(ctx context.Context, in *QueryTunnelRequest, opts ...grpc.CallOption) (*QueryTunnelResponse, error) {
	out := new(QueryTunnelResponse)
	err := c.cc.Invoke(ctx, "/tunnel.v1beta1.Query/Tunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params is a RPC method that returns all parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Tunnels is a RPC method that returns all tunnels.
	Tunnels(context.Context, *QueryTunnelsRequest) (*QueryTunnelsResponse, error)
	// Tunnel is a RPC method that returns a tunnel by its ID.
	Tunnel(context.Context, *QueryTunnelRequest) (*QueryTunnelResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Tunnels(ctx context.Context, req *QueryTunnelsRequest) (*QueryTunnelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tunnels not implemented")
}
func (*UnimplementedQueryServer) Tunnel(ctx context.Context, req *QueryTunnelRequest) (*QueryTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tunnel not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tunnel.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Tunnels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTunnelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Tunnels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tunnel.v1beta1.Query/Tunnels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Tunnels(ctx, req.(*QueryTunnelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Tunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Tunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tunnel.v1beta1.Query/Tunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Tunnel(ctx, req.(*QueryTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tunnel.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Tunnels",
			Handler:    _Query_Tunnels_Handler,
		},
		{
			MethodName: "Tunnel",
			Handler:    _Query_Tunnel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tunnel/v1beta1/query.proto",
}

func (m *QueryTunnelsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTunnelsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTunnelsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryTunnelsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTunnelsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTunnelsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Tunnels) > 0 {
		for iNdEx := len(m.Tunnels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tunnels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryTunnelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTunnelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTunnelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TunnelId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.TunnelId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryTunnelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTunnelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTunnelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Tunnel.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryTunnelsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsActive {
		n += 2
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryTunnelsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tunnels) > 0 {
		for _, e := range m.Tunnels {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryTunnelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TunnelId != 0 {
		n += 1 + sovQuery(uint64(m.TunnelId))
	}
	return n
}

func (m *QueryTunnelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tunnel.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTunnelsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTunnelsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTunnelsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTunnelsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTunnelsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTunnelsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tunnels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tunnels = append(m.Tunnels, &Tunnel{})
			if err := m.Tunnels[len(m.Tunnels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTunnelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTunnelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTunnelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TunnelId", wireType)
			}
			m.TunnelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TunnelId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTunnelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTunnelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTunnelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tunnel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tunnel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
