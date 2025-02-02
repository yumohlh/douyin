// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: UserOptService.proto

package userOptPb

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

// UserOptServiceClient is the client API for UserOptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserOptServiceClient interface {
	// -----------------------userFavoriteList-----------------------
	GetUserFavorite(ctx context.Context, in *GetUserFavoriteReq, opts ...grpc.CallOption) (*GetUserFavoriteResp, error)
	UpdateFavoriteStatus(ctx context.Context, in *UpdateFavoriteStatusReq, opts ...grpc.CallOption) (*UpdateFavoriteStatusResp, error)
}

type userOptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserOptServiceClient(cc grpc.ClientConnInterface) UserOptServiceClient {
	return &userOptServiceClient{cc}
}

func (c *userOptServiceClient) GetUserFavorite(ctx context.Context, in *GetUserFavoriteReq, opts ...grpc.CallOption) (*GetUserFavoriteResp, error) {
	out := new(GetUserFavoriteResp)
	err := c.cc.Invoke(ctx, "/pb.UserOptService/GetUserFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOptServiceClient) UpdateFavoriteStatus(ctx context.Context, in *UpdateFavoriteStatusReq, opts ...grpc.CallOption) (*UpdateFavoriteStatusResp, error) {
	out := new(UpdateFavoriteStatusResp)
	err := c.cc.Invoke(ctx, "/pb.UserOptService/UpdateFavoriteStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserOptServiceServer is the server API for UserOptService service.
// All implementations must embed UnimplementedUserOptServiceServer
// for forward compatibility
type UserOptServiceServer interface {
	// -----------------------userFavoriteList-----------------------
	GetUserFavorite(context.Context, *GetUserFavoriteReq) (*GetUserFavoriteResp, error)
	UpdateFavoriteStatus(context.Context, *UpdateFavoriteStatusReq) (*UpdateFavoriteStatusResp, error)
	mustEmbedUnimplementedUserOptServiceServer()
}

// UnimplementedUserOptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserOptServiceServer struct {
}

func (UnimplementedUserOptServiceServer) GetUserFavorite(context.Context, *GetUserFavoriteReq) (*GetUserFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavorite not implemented")
}
func (UnimplementedUserOptServiceServer) UpdateFavoriteStatus(context.Context, *UpdateFavoriteStatusReq) (*UpdateFavoriteStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavoriteStatus not implemented")
}
func (UnimplementedUserOptServiceServer) mustEmbedUnimplementedUserOptServiceServer() {}

// UnsafeUserOptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserOptServiceServer will
// result in compilation errors.
type UnsafeUserOptServiceServer interface {
	mustEmbedUnimplementedUserOptServiceServer()
}

func RegisterUserOptServiceServer(s grpc.ServiceRegistrar, srv UserOptServiceServer) {
	s.RegisterService(&UserOptService_ServiceDesc, srv)
}

func _UserOptService_GetUserFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOptServiceServer).GetUserFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOptService/GetUserFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOptServiceServer).GetUserFavorite(ctx, req.(*GetUserFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOptService_UpdateFavoriteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavoriteStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOptServiceServer).UpdateFavoriteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOptService/UpdateFavoriteStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOptServiceServer).UpdateFavoriteStatus(ctx, req.(*UpdateFavoriteStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserOptService_ServiceDesc is the grpc.ServiceDesc for UserOptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserOptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserOptService",
	HandlerType: (*UserOptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserFavorite",
			Handler:    _UserOptService_GetUserFavorite_Handler,
		},
		{
			MethodName: "UpdateFavoriteStatus",
			Handler:    _UserOptService_UpdateFavoriteStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserOptService.proto",
}
