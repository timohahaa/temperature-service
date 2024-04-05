// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/record/record.proto

package record

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RecordClient is the client API for Record service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordClient interface {
	GetAvgTemp(ctx context.Context, in *GetAvgTempReq, opts ...grpc.CallOption) (*GetAvgTempResp, error)
}

type recordClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordClient(cc grpc.ClientConnInterface) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) GetAvgTemp(ctx context.Context, in *GetAvgTempReq, opts ...grpc.CallOption) (*GetAvgTempResp, error) {
	out := new(GetAvgTempResp)
	err := c.cc.Invoke(ctx, "/Record/GetAvgTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServer is the server API for Record service.
// All implementations must embed UnimplementedRecordServer
// for forward compatibility
type RecordServer interface {
	GetAvgTemp(context.Context, *GetAvgTempReq) (*GetAvgTempResp, error)
	mustEmbedUnimplementedRecordServer()
}

// UnimplementedRecordServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServer struct {
}

func (UnimplementedRecordServer) GetAvgTemp(context.Context, *GetAvgTempReq) (*GetAvgTempResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvgTemp not implemented")
}
func (UnimplementedRecordServer) mustEmbedUnimplementedRecordServer() {}

// UnsafeRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServer will
// result in compilation errors.
type UnsafeRecordServer interface {
	mustEmbedUnimplementedRecordServer()
}

func RegisterRecordServer(s grpc.ServiceRegistrar, srv RecordServer) {
	s.RegisterService(&Record_ServiceDesc, srv)
}

func _Record_GetAvgTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvgTempReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).GetAvgTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Record/GetAvgTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).GetAvgTemp(ctx, req.(*GetAvgTempReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Record_ServiceDesc is the grpc.ServiceDesc for Record service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Record_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvgTemp",
			Handler:    _Record_GetAvgTemp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/record/record.proto",
}
