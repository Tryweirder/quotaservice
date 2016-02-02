// Code generated by protoc-gen-go.
// source: protos/quota_service.proto
// DO NOT EDIT!

/*
Package quotaservice is a generated protocol buffer package.

It is generated from these files:
	protos/quota_service.proto

It has these top-level messages:
	AllowRequest
	AllowResponse
*/
package quotaservice

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

type AllowRequest_EmptyBucketPolicyOverride int32

const (
	AllowRequest_SERVER_DEFAULTS AllowRequest_EmptyBucketPolicyOverride = 0
	AllowRequest_WAIT            AllowRequest_EmptyBucketPolicyOverride = 1
	AllowRequest_REJECT          AllowRequest_EmptyBucketPolicyOverride = 2
)

var AllowRequest_EmptyBucketPolicyOverride_name = map[int32]string{
	0: "SERVER_DEFAULTS",
	1: "WAIT",
	2: "REJECT",
}
var AllowRequest_EmptyBucketPolicyOverride_value = map[string]int32{
	"SERVER_DEFAULTS": 0,
	"WAIT":            1,
	"REJECT":          2,
}

func (x AllowRequest_EmptyBucketPolicyOverride) String() string {
	return proto.EnumName(AllowRequest_EmptyBucketPolicyOverride_name, int32(x))
}
func (AllowRequest_EmptyBucketPolicyOverride) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type AllowResponse_Status int32

const (
	AllowResponse_OK        AllowResponse_Status = 0
	AllowResponse_TIMED_OUT AllowResponse_Status = 1
	AllowResponse_REJECTED  AllowResponse_Status = 2
)

var AllowResponse_Status_name = map[int32]string{
	0: "OK",
	1: "TIMED_OUT",
	2: "REJECTED",
}
var AllowResponse_Status_value = map[string]int32{
	"OK":        0,
	"TIMED_OUT": 1,
	"REJECTED":  2,
}

func (x AllowResponse_Status) String() string {
	return proto.EnumName(AllowResponse_Status_name, int32(x))
}
func (AllowResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type AllowRequest struct {
	BucketName        string                                 `protobuf:"bytes,1,opt,name=bucket_name" json:"bucket_name,omitempty"`
	TokensRequested   int32                                  `protobuf:"varint,2,opt,name=tokens_requested" json:"tokens_requested,omitempty"`
	EmptyBucketPolicy AllowRequest_EmptyBucketPolicyOverride `protobuf:"varint,3,opt,name=empty_bucket_policy,enum=quotaservice.AllowRequest_EmptyBucketPolicyOverride" json:"empty_bucket_policy,omitempty"`
}

func (m *AllowRequest) Reset()                    { *m = AllowRequest{} }
func (m *AllowRequest) String() string            { return proto.CompactTextString(m) }
func (*AllowRequest) ProtoMessage()               {}
func (*AllowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllowResponse struct {
	TokensGranted int32                `protobuf:"varint,1,opt,name=tokens_granted" json:"tokens_granted,omitempty"`
	Status        AllowResponse_Status `protobuf:"varint,2,opt,name=status,enum=quotaservice.AllowResponse_Status" json:"status,omitempty"`
}

func (m *AllowResponse) Reset()                    { *m = AllowResponse{} }
func (m *AllowResponse) String() string            { return proto.CompactTextString(m) }
func (*AllowResponse) ProtoMessage()               {}
func (*AllowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*AllowRequest)(nil), "quotaservice.AllowRequest")
	proto.RegisterType((*AllowResponse)(nil), "quotaservice.AllowResponse")
	proto.RegisterEnum("quotaservice.AllowRequest_EmptyBucketPolicyOverride", AllowRequest_EmptyBucketPolicyOverride_name, AllowRequest_EmptyBucketPolicyOverride_value)
	proto.RegisterEnum("quotaservice.AllowResponse_Status", AllowResponse_Status_name, AllowResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for QuotaService service

type QuotaServiceClient interface {
	Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error)
}

type quotaServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuotaServiceClient(cc *grpc.ClientConn) QuotaServiceClient {
	return &quotaServiceClient{cc}
}

func (c *quotaServiceClient) Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error) {
	out := new(AllowResponse)
	err := grpc.Invoke(ctx, "/quotaservice.QuotaService/Allow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QuotaService service

type QuotaServiceServer interface {
	Allow(context.Context, *AllowRequest) (*AllowResponse, error)
}

func RegisterQuotaServiceServer(s *grpc.Server, srv QuotaServiceServer) {
	s.RegisterService(&_QuotaService_serviceDesc, srv)
}

func _QuotaService_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(QuotaServiceServer).Allow(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _QuotaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quotaservice.QuotaService",
	HandlerType: (*QuotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allow",
			Handler:    _QuotaService_Allow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x69, 0x95, 0x06, 0xae, 0x05, 0x9b, 0x21, 0x31, 0x15, 0x37, 0xa6, 0x2b, 0x37, 0xd6,
	0xa4, 0xfa, 0x02, 0x20, 0x43, 0x82, 0x3f, 0x41, 0xda, 0xa2, 0xcb, 0xa6, 0xc0, 0x8d, 0x21, 0x40,
	0xa7, 0xcc, 0x4c, 0x31, 0xbc, 0x82, 0x4f, 0xe8, 0xe3, 0xd8, 0x4e, 0xbb, 0x60, 0x21, 0xee, 0x9a,
	0x7b, 0xef, 0x39, 0xe7, 0x3b, 0x1d, 0xe8, 0xa6, 0x9c, 0x49, 0x26, 0xee, 0xb6, 0x19, 0x93, 0x71,
	0x24, 0x90, 0xef, 0x96, 0x73, 0x74, 0xd5, 0x90, 0x98, 0x6a, 0x58, 0xcd, 0x9c, 0x1f, 0x0d, 0xcc,
	0xde, 0x7a, 0xcd, 0xbe, 0x7c, 0xdc, 0x66, 0x28, 0x24, 0xe9, 0xc0, 0xd9, 0x2c, 0x9b, 0xaf, 0x50,
	0x46, 0x49, 0xbc, 0x41, 0x5b, 0xbb, 0xd6, 0x6e, 0x9a, 0xc4, 0x06, 0x4b, 0xb2, 0x15, 0x26, 0x22,
	0xe2, 0xe5, 0x19, 0x2e, 0x6c, 0x3d, 0xdf, 0xd4, 0xc9, 0x04, 0x3a, 0xb8, 0x49, 0xe5, 0x3e, 0xaa,
	0x44, 0x29, 0x5b, 0x2f, 0xe7, 0x7b, 0xfb, 0x24, 0x5f, 0xb6, 0xbd, 0x07, 0xf7, 0x30, 0xcb, 0x3d,
	0xcc, 0x71, 0x69, 0xa1, 0xea, 0x2b, 0xd1, 0x9b, 0xd2, 0x8c, 0x77, 0xc8, 0xf9, 0x72, 0x81, 0xce,
	0x10, 0x2e, 0x8f, 0x2e, 0x73, 0xbc, 0xf3, 0x80, 0xfa, 0xef, 0xd4, 0x8f, 0x06, 0x74, 0xd8, 0x9b,
	0xbe, 0x84, 0x81, 0x55, 0x23, 0x0d, 0x38, 0xfd, 0xe8, 0x8d, 0x42, 0x4b, 0x23, 0x00, 0x86, 0x4f,
	0x9f, 0xe8, 0x63, 0x68, 0xe9, 0xce, 0xb7, 0x06, 0xad, 0x2a, 0x52, 0xa4, 0x2c, 0x11, 0x48, 0x2e,
	0xa0, 0x5d, 0xd5, 0xf8, 0xe4, 0x71, 0x52, 0x94, 0xd0, 0x54, 0x09, 0x0f, 0x0c, 0x21, 0x63, 0x99,
	0x09, 0x55, 0xaa, 0xed, 0x39, 0x7f, 0x72, 0x97, 0x26, 0x6e, 0xa0, 0x2e, 0x9d, 0x5b, 0x30, 0xca,
	0x2f, 0x62, 0x80, 0x3e, 0x7e, 0xce, 0x29, 0x5a, 0xd0, 0x0c, 0x47, 0xaf, 0x74, 0x10, 0x8d, 0xa7,
	0x05, 0x8a, 0x09, 0x8d, 0x12, 0x85, 0x0e, 0x2c, 0xdd, 0xf3, 0xc1, 0x9c, 0x14, 0x9e, 0x41, 0xe9,
	0x49, 0xfa, 0x50, 0x57, 0xb6, 0xa4, 0x7b, 0xfc, 0x1f, 0x75, 0xaf, 0xfe, 0xe1, 0x70, 0x6a, 0x33,
	0x43, 0x3d, 0xe8, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xc4, 0x9b, 0x34, 0xee, 0x01,
	0x00, 0x00,
}
