// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanfury/fantoken/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryFanTokenRequest is request type for the Query/FanToken RPC method
type QueryFanTokenRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryFanTokenRequest) Reset()         { *m = QueryFanTokenRequest{} }
func (m *QueryFanTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokenRequest) ProtoMessage()    {}
func (*QueryFanTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa40654ae0b6566, []int{0}
}
func (m *QueryFanTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokenRequest.Merge(m, src)
}
func (m *QueryFanTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokenRequest proto.InternalMessageInfo

func (m *QueryFanTokenRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryFanTokenResponse is response type for the Query/FanToken RPC method
type QueryFanTokenResponse struct {
	Fantoken *FanToken `protobuf:"bytes,1,opt,name=fantoken,proto3" json:"fantoken,omitempty"`
}

func (m *QueryFanTokenResponse) Reset()         { *m = QueryFanTokenResponse{} }
func (m *QueryFanTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokenResponse) ProtoMessage()    {}
func (*QueryFanTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa40654ae0b6566, []int{1}
}
func (m *QueryFanTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokenResponse.Merge(m, src)
}
func (m *QueryFanTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokenResponse proto.InternalMessageInfo

func (m *QueryFanTokenResponse) GetFantoken() *FanToken {
	if m != nil {
		return m.Fantoken
	}
	return nil
}

// QueryFanTokensRequest is request type for the Query/FanTokens RPC method
type QueryFanTokensRequest struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFanTokensRequest) Reset()         { *m = QueryFanTokensRequest{} }
func (m *QueryFanTokensRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokensRequest) ProtoMessage()    {}
func (*QueryFanTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa40654ae0b6566, []int{2}
}
func (m *QueryFanTokensRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokensRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokensRequest.Merge(m, src)
}
func (m *QueryFanTokensRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokensRequest proto.InternalMessageInfo

func (m *QueryFanTokensRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *QueryFanTokensRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryFanTokensResponse is response type for the Query/FanTokens RPC method
type QueryFanTokensResponse struct {
	Fantokens  []*FanToken         `protobuf:"bytes,1,rep,name=fantokens,proto3" json:"fantokens,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFanTokensResponse) Reset()         { *m = QueryFanTokensResponse{} }
func (m *QueryFanTokensResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokensResponse) ProtoMessage()    {}
func (*QueryFanTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa40654ae0b6566, []int{3}
}
func (m *QueryFanTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokensResponse.Merge(m, src)
}
func (m *QueryFanTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokensResponse proto.InternalMessageInfo

func (m *QueryFanTokensResponse) GetFantokens() []*FanToken {
	if m != nil {
		return m.Fantokens
	}
	return nil
}

func (m *QueryFanTokensResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryParametersRequest is request type for the Query/Parameters RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa40654ae0b6566, []int{4}
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

// QueryParametersResponse is response type for the Query/Parameters RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa40654ae0b6566, []int{5}
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
	proto.RegisterType((*QueryFanTokenRequest)(nil), "fanfury.fantoken.v1beta1.QueryFanTokenRequest")
	proto.RegisterType((*QueryFanTokenResponse)(nil), "fanfury.fantoken.v1beta1.QueryFanTokenResponse")
	proto.RegisterType((*QueryFanTokensRequest)(nil), "fanfury.fantoken.v1beta1.QueryFanTokensRequest")
	proto.RegisterType((*QueryFanTokensResponse)(nil), "fanfury.fantoken.v1beta1.QueryFanTokensResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "fanfury.fantoken.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "fanfury.fantoken.v1beta1.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("fanfury/fantoken/v1beta1/query.proto", fileDescriptor_2fa40654ae0b6566)
}

var fileDescriptor_2fa40654ae0b6566 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x8d, 0x55, 0xab, 0xb9, 0x99, 0x82, 0xaa, 0x68, 0x0a, 0x55, 0x60, 0x74, 0x40,
	0xb1, 0xd9, 0xb8, 0x23, 0xb4, 0x43, 0xb9, 0x8e, 0x0a, 0x84, 0xc4, 0xcd, 0x2d, 0x6e, 0x1a, 0x41,
	0xfc, 0xb2, 0xd8, 0x41, 0x54, 0x68, 0x17, 0xbe, 0x00, 0x48, 0x1c, 0xe1, 0xc6, 0x97, 0xd9, 0x71,
	0x12, 0x12, 0xe2, 0x84, 0x50, 0xcb, 0x07, 0x99, 0x6a, 0x3b, 0xe9, 0x5a, 0x35, 0x6a, 0x4e, 0x4d,
	0x9d, 0xff, 0xff, 0xbd, 0x9f, 0xdf, 0xdf, 0x0e, 0xbe, 0x3b, 0xe2, 0x72, 0x94, 0xa5, 0x13, 0x36,
	0xe2, 0x52, 0xc3, 0x3b, 0x21, 0xd9, 0x87, 0xc3, 0x81, 0xd0, 0xfc, 0x90, 0x9d, 0x66, 0x22, 0x9d,
	0xd0, 0x24, 0x05, 0x0d, 0xa4, 0xe5, 0x54, 0x34, 0x57, 0x51, 0xa7, 0xf2, 0xfc, 0x21, 0xa8, 0x18,
	0x14, 0x1b, 0x70, 0x25, 0x0a, 0xeb, 0x10, 0x22, 0x69, 0x9d, 0xde, 0x83, 0xab, 0xef, 0x4d, 0xc9,
	0x42, 0x95, 0xf0, 0x30, 0x92, 0x5c, 0x47, 0x90, 0x6b, 0x9b, 0x21, 0x84, 0x60, 0x1e, 0xd9, 0xfc,
	0xc9, 0xad, 0xee, 0x85, 0x00, 0xe1, 0x7b, 0xc1, 0x78, 0x12, 0x31, 0x2e, 0x25, 0x68, 0x63, 0x51,
	0xee, 0x6d, 0xa7, 0x94, 0xbf, 0x40, 0xb5, 0xc2, 0xfd, 0x52, 0x61, 0xc2, 0x53, 0x1e, 0xbb, 0x7a,
	0x41, 0x17, 0x37, 0x5f, 0xcc, 0x29, 0x7b, 0x5c, 0xbe, 0x9c, 0xab, 0xfa, 0xe2, 0x34, 0x13, 0x4a,
	0x93, 0x26, 0xde, 0x79, 0x2b, 0x24, 0xc4, 0x2d, 0xd4, 0x46, 0x07, 0x8d, 0xbe, 0xfd, 0x13, 0xbc,
	0xc6, 0x37, 0x57, 0xd4, 0x2a, 0x01, 0xa9, 0x04, 0x79, 0x8a, 0x77, 0xf3, 0x3e, 0xc6, 0x71, 0xfd,
	0x28, 0xa0, 0x65, 0x33, 0xa4, 0x85, 0xbb, 0xf0, 0x04, 0x67, 0x2b, 0x85, 0x55, 0xce, 0xb1, 0x87,
	0x1b, 0x3c, 0xd3, 0x63, 0x48, 0x23, 0x3d, 0x71, 0x2c, 0x8b, 0x05, 0xd2, 0xc3, 0x78, 0x31, 0xd5,
	0xd6, 0x96, 0x69, 0x7c, 0x8f, 0xda, 0x08, 0xe8, 0x3c, 0x02, 0x6a, 0x53, 0xcd, 0x3b, 0x9f, 0xf0,
	0x50, 0xb8, 0xca, 0xfd, 0x2b, 0xce, 0xe0, 0x27, 0xc2, 0xb7, 0x56, 0xfb, 0xbb, 0x9d, 0x3d, 0xc3,
	0x8d, 0x9c, 0x52, 0xb5, 0x50, 0x7b, 0xbb, 0xe2, 0xd6, 0x16, 0x26, 0xf2, 0x7c, 0x0d, 0x64, 0x67,
	0x23, 0xa4, 0x6d, 0xbf, 0x44, 0xd9, 0xc4, 0xc4, 0x40, 0x9e, 0x98, 0x00, 0xdd, 0x3e, 0x82, 0x57,
	0xf8, 0xc6, 0xd2, 0x6a, 0x91, 0x48, 0xdd, 0x06, 0xed, 0xf2, 0x68, 0x97, 0x43, 0x5b, 0xe7, 0xf1,
	0xb5, 0xf3, 0xbf, 0xb7, 0x6b, 0x7d, 0xe7, 0x3a, 0xfa, 0xbd, 0x8d, 0x77, 0x4c, 0x5d, 0xf2, 0x03,
	0xe1, 0xdd, 0x7c, 0x5f, 0x84, 0x96, 0x97, 0x59, 0x77, 0x8e, 0x3c, 0x56, 0x59, 0x6f, 0xb9, 0x03,
	0xf6, 0xf9, 0xd7, 0xff, 0x6f, 0x5b, 0xf7, 0x49, 0x87, 0x95, 0x1e, 0x60, 0x73, 0x16, 0xd9, 0x27,
	0xf3, 0x73, 0x46, 0xbe, 0x23, 0xdc, 0x28, 0x62, 0x23, 0x55, 0xfb, 0xe5, 0xe3, 0xf3, 0x1e, 0x57,
	0x37, 0x38, 0xc2, 0x87, 0x86, 0x70, 0x9f, 0xdc, 0x61, 0x1b, 0xef, 0xa2, 0x22, 0x5f, 0x10, 0xae,
	0xdb, 0xf9, 0x92, 0xee, 0x86, 0x4e, 0x4b, 0xb1, 0x7a, 0x8f, 0x2a, 0xaa, 0x1d, 0xd4, 0x81, 0x81,
	0x0a, 0x48, 0x9b, 0x6d, 0xb8, 0xf7, 0xc7, 0xbd, 0xf3, 0xa9, 0x8f, 0x2e, 0xa6, 0x3e, 0xfa, 0x37,
	0xf5, 0xd1, 0xd7, 0x99, 0x5f, 0xbb, 0x98, 0xf9, 0xb5, 0x3f, 0x33, 0xbf, 0xf6, 0xa6, 0x1b, 0x46,
	0x7a, 0x9c, 0x0d, 0xe8, 0x10, 0x62, 0x26, 0x45, 0x32, 0x8e, 0xd2, 0x28, 0x2e, 0xca, 0x7d, 0x5c,
	0x14, 0xd4, 0x93, 0x44, 0xa8, 0x41, 0xdd, 0x7c, 0x40, 0x9e, 0x5c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xe3, 0x44, 0x13, 0x5c, 0x52, 0x05, 0x00, 0x00,
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
	// FanToken returns fantoken with fantoken name
	FanToken(ctx context.Context, in *QueryFanTokenRequest, opts ...grpc.CallOption) (*QueryFanTokenResponse, error)
	// FanTokens returns the fantoken list
	FanTokens(ctx context.Context, in *QueryFanTokensRequest, opts ...grpc.CallOption) (*QueryFanTokensResponse, error)
	// Params queries the fantoken parameters
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) FanToken(ctx context.Context, in *QueryFanTokenRequest, opts ...grpc.CallOption) (*QueryFanTokenResponse, error) {
	out := new(QueryFanTokenResponse)
	err := c.cc.Invoke(ctx, "/fanfury.fantoken.v1beta1.Query/FanToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FanTokens(ctx context.Context, in *QueryFanTokensRequest, opts ...grpc.CallOption) (*QueryFanTokensResponse, error) {
	out := new(QueryFanTokensResponse)
	err := c.cc.Invoke(ctx, "/fanfury.fantoken.v1beta1.Query/FanTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/fanfury.fantoken.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// FanToken returns fantoken with fantoken name
	FanToken(context.Context, *QueryFanTokenRequest) (*QueryFanTokenResponse, error)
	// FanTokens returns the fantoken list
	FanTokens(context.Context, *QueryFanTokensRequest) (*QueryFanTokensResponse, error)
	// Params queries the fantoken parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) FanToken(ctx context.Context, req *QueryFanTokenRequest) (*QueryFanTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanToken not implemented")
}
func (*UnimplementedQueryServer) FanTokens(ctx context.Context, req *QueryFanTokensRequest) (*QueryFanTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanTokens not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_FanToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFanTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FanToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fanfury.fantoken.v1beta1.Query/FanToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FanToken(ctx, req.(*QueryFanTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FanTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFanTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FanTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fanfury.fantoken.v1beta1.Query/FanTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FanTokens(ctx, req.(*QueryFanTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/fanfury.fantoken.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fanfury.fantoken.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FanToken",
			Handler:    _Query_FanToken_Handler,
		},
		{
			MethodName: "FanTokens",
			Handler:    _Query_FanTokens_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fanfury/fantoken/v1beta1/query.proto",
}

func (m *QueryFanTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFanTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fantoken != nil {
		{
			size, err := m.Fantoken.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFanTokensRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokensRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokensRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFanTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Fantokens) > 0 {
		for iNdEx := len(m.Fantokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fantokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryFanTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFanTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fantoken != nil {
		l = m.Fantoken.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFanTokensRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFanTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fantokens) > 0 {
		for _, e := range m.Fantokens {
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
func (m *QueryFanTokenRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryFanTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fantoken", wireType)
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
			if m.Fantoken == nil {
				m.Fantoken = &FanToken{}
			}
			if err := m.Fantoken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFanTokensRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokensRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokensRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
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
func (m *QueryFanTokensResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fantokens", wireType)
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
			m.Fantokens = append(m.Fantokens, &FanToken{})
			if err := m.Fantokens[len(m.Fantokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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