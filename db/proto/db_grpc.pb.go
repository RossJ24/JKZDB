// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: db/proto/db.proto

package proto

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

// JKZDBClient is the client API for JKZDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JKZDBClient interface {
	SetEntryPrepare(ctx context.Context, in *SetEntryPrepareRequest, opts ...grpc.CallOption) (*SetEntryPrepareResponse, error)
	SetEntryCommit(ctx context.Context, in *SetEntryCommitRequest, opts ...grpc.CallOption) (*SetEntryCommitResponse, error)
	SetEntryAbort(ctx context.Context, in *SetEntryAbortRequest, opts ...grpc.CallOption) (*SetEntryAbortResponse, error)
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	GetEntryByIndexedField(ctx context.Context, in *GetEntryByIndexedFieldRequest, opts ...grpc.CallOption) (*GetEntryByIndexedFieldResponse, error)
}

type jKZDBClient struct {
	cc grpc.ClientConnInterface
}

func NewJKZDBClient(cc grpc.ClientConnInterface) JKZDBClient {
	return &jKZDBClient{cc}
}

func (c *jKZDBClient) SetEntryPrepare(ctx context.Context, in *SetEntryPrepareRequest, opts ...grpc.CallOption) (*SetEntryPrepareResponse, error) {
	out := new(SetEntryPrepareResponse)
	err := c.cc.Invoke(ctx, "/db.JKZDB/SetEntryPrepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jKZDBClient) SetEntryCommit(ctx context.Context, in *SetEntryCommitRequest, opts ...grpc.CallOption) (*SetEntryCommitResponse, error) {
	out := new(SetEntryCommitResponse)
	err := c.cc.Invoke(ctx, "/db.JKZDB/SetEntryCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jKZDBClient) SetEntryAbort(ctx context.Context, in *SetEntryAbortRequest, opts ...grpc.CallOption) (*SetEntryAbortResponse, error) {
	out := new(SetEntryAbortResponse)
	err := c.cc.Invoke(ctx, "/db.JKZDB/SetEntryAbort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jKZDBClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := c.cc.Invoke(ctx, "/db.JKZDB/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jKZDBClient) GetEntryByIndexedField(ctx context.Context, in *GetEntryByIndexedFieldRequest, opts ...grpc.CallOption) (*GetEntryByIndexedFieldResponse, error) {
	out := new(GetEntryByIndexedFieldResponse)
	err := c.cc.Invoke(ctx, "/db.JKZDB/GetEntryByIndexedField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JKZDBServer is the server API for JKZDB service.
// All implementations must embed UnimplementedJKZDBServer
// for forward compatibility
type JKZDBServer interface {
	SetEntryPrepare(context.Context, *SetEntryPrepareRequest) (*SetEntryPrepareResponse, error)
	SetEntryCommit(context.Context, *SetEntryCommitRequest) (*SetEntryCommitResponse, error)
	SetEntryAbort(context.Context, *SetEntryAbortRequest) (*SetEntryAbortResponse, error)
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	GetEntryByIndexedField(context.Context, *GetEntryByIndexedFieldRequest) (*GetEntryByIndexedFieldResponse, error)
	mustEmbedUnimplementedJKZDBServer()
}

// UnimplementedJKZDBServer must be embedded to have forward compatible implementations.
type UnimplementedJKZDBServer struct {
}

func (UnimplementedJKZDBServer) SetEntryPrepare(context.Context, *SetEntryPrepareRequest) (*SetEntryPrepareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEntryPrepare not implemented")
}
func (UnimplementedJKZDBServer) SetEntryCommit(context.Context, *SetEntryCommitRequest) (*SetEntryCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEntryCommit not implemented")
}
func (UnimplementedJKZDBServer) SetEntryAbort(context.Context, *SetEntryAbortRequest) (*SetEntryAbortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEntryAbort not implemented")
}
func (UnimplementedJKZDBServer) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedJKZDBServer) GetEntryByIndexedField(context.Context, *GetEntryByIndexedFieldRequest) (*GetEntryByIndexedFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntryByIndexedField not implemented")
}
func (UnimplementedJKZDBServer) mustEmbedUnimplementedJKZDBServer() {}

// UnsafeJKZDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JKZDBServer will
// result in compilation errors.
type UnsafeJKZDBServer interface {
	mustEmbedUnimplementedJKZDBServer()
}

func RegisterJKZDBServer(s grpc.ServiceRegistrar, srv JKZDBServer) {
	s.RegisterService(&JKZDB_ServiceDesc, srv)
}

func _JKZDB_SetEntryPrepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEntryPrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JKZDBServer).SetEntryPrepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.JKZDB/SetEntryPrepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JKZDBServer).SetEntryPrepare(ctx, req.(*SetEntryPrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JKZDB_SetEntryCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEntryCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JKZDBServer).SetEntryCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.JKZDB/SetEntryCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JKZDBServer).SetEntryCommit(ctx, req.(*SetEntryCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JKZDB_SetEntryAbort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEntryAbortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JKZDBServer).SetEntryAbort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.JKZDB/SetEntryAbort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JKZDBServer).SetEntryAbort(ctx, req.(*SetEntryAbortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JKZDB_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JKZDBServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.JKZDB/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JKZDBServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JKZDB_GetEntryByIndexedField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryByIndexedFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JKZDBServer).GetEntryByIndexedField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.JKZDB/GetEntryByIndexedField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JKZDBServer).GetEntryByIndexedField(ctx, req.(*GetEntryByIndexedFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JKZDB_ServiceDesc is the grpc.ServiceDesc for JKZDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JKZDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.JKZDB",
	HandlerType: (*JKZDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetEntryPrepare",
			Handler:    _JKZDB_SetEntryPrepare_Handler,
		},
		{
			MethodName: "SetEntryCommit",
			Handler:    _JKZDB_SetEntryCommit_Handler,
		},
		{
			MethodName: "SetEntryAbort",
			Handler:    _JKZDB_SetEntryAbort_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _JKZDB_GetEntry_Handler,
		},
		{
			MethodName: "GetEntryByIndexedField",
			Handler:    _JKZDB_GetEntryByIndexedField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db/proto/db.proto",
}
