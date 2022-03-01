// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/rewards/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("comdex/rewards/v1beta1/msg.proto", fileDescriptor_b7586e0f456d0ddd) }

var fileDescriptor_b7586e0f456d0ddd = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0xcf, 0x4d,
	0x49, 0xad, 0xd0, 0x2f, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x29, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x85, 0xa8,
	0xc8, 0x4f, 0x4b, 0xcb, 0x4c, 0xce, 0x4c, 0xcc, 0xd1, 0x83, 0x70, 0xf5, 0xa0, 0x1a, 0x8c, 0x78,
	0xb8, 0xb8, 0x7c, 0x8b, 0xd3, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x9d, 0xbc, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x30, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x09,
	0xa4, 0x5d, 0x1f, 0x62, 0x84, 0x2e, 0xcc, 0x48, 0x28, 0x5f, 0x1f, 0xe1, 0x8a, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x03, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0xbb, 0xec,
	0xfe, 0xa4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdexofficial.comdex.rewards.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "comdex/rewards/v1beta1/msg.proto",
}
