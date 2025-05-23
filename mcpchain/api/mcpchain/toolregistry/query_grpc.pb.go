// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: mcpchain/toolregistry/query.proto

package toolregistry

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

const (
	Query_Params_FullMethodName         = "/mcpchain.toolregistry.Query/Params"
	Query_ServerById_FullMethodName     = "/mcpchain.toolregistry.Query/ServerById"
	Query_ServersByOwner_FullMethodName = "/mcpchain.toolregistry.Query/ServersByOwner"
	Query_AllServers_FullMethodName     = "/mcpchain.toolregistry.Query/AllServers"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ServerById items.
	ServerById(ctx context.Context, in *QueryServerByIdRequest, opts ...grpc.CallOption) (*QueryServerByIdResponse, error)
	// Queries a list of ServersByOwner items.
	ServersByOwner(ctx context.Context, in *QueryServersByOwnerRequest, opts ...grpc.CallOption) (*QueryServersByOwnerResponse, error)
	// Queries a list of AllServers items.
	AllServers(ctx context.Context, in *QueryAllServersRequest, opts ...grpc.CallOption) (*QueryAllServersResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ServerById(ctx context.Context, in *QueryServerByIdRequest, opts ...grpc.CallOption) (*QueryServerByIdResponse, error) {
	out := new(QueryServerByIdResponse)
	err := c.cc.Invoke(ctx, Query_ServerById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ServersByOwner(ctx context.Context, in *QueryServersByOwnerRequest, opts ...grpc.CallOption) (*QueryServersByOwnerResponse, error) {
	out := new(QueryServersByOwnerResponse)
	err := c.cc.Invoke(ctx, Query_ServersByOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllServers(ctx context.Context, in *QueryAllServersRequest, opts ...grpc.CallOption) (*QueryAllServersResponse, error) {
	out := new(QueryAllServersResponse)
	err := c.cc.Invoke(ctx, Query_AllServers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ServerById items.
	ServerById(context.Context, *QueryServerByIdRequest) (*QueryServerByIdResponse, error)
	// Queries a list of ServersByOwner items.
	ServersByOwner(context.Context, *QueryServersByOwnerRequest) (*QueryServersByOwnerResponse, error)
	// Queries a list of AllServers items.
	AllServers(context.Context, *QueryAllServersRequest) (*QueryAllServersResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ServerById(context.Context, *QueryServerByIdRequest) (*QueryServerByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerById not implemented")
}
func (UnimplementedQueryServer) ServersByOwner(context.Context, *QueryServersByOwnerRequest) (*QueryServersByOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServersByOwner not implemented")
}
func (UnimplementedQueryServer) AllServers(context.Context, *QueryAllServersRequest) (*QueryAllServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllServers not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
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
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ServerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ServerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ServerById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ServerById(ctx, req.(*QueryServerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ServersByOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServersByOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ServersByOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ServersByOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ServersByOwner(ctx, req.(*QueryServersByOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllServers(ctx, req.(*QueryAllServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcpchain.toolregistry.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ServerById",
			Handler:    _Query_ServerById_Handler,
		},
		{
			MethodName: "ServersByOwner",
			Handler:    _Query_ServersByOwner_Handler,
		},
		{
			MethodName: "AllServers",
			Handler:    _Query_AllServers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcpchain/toolregistry/query.proto",
}
