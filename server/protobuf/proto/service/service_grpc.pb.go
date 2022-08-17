// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: proto/service/service.proto

package service

import (
	context "context"
	message "github.com/dtc03012/me/protobuf/proto/service/message"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeClient is the client API for Me service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeClient interface {
	FindAdminUUID(ctx context.Context, in *message.FindAdminUUIDRequest, opts ...grpc.CallOption) (*message.FindAdminUUIDResponse, error)
	InsertAdminUUID(ctx context.Context, in *message.InsertAdminUUIDRequest, opts ...grpc.CallOption) (*message.InsertAdminUUIDResponse, error)
	LoginAdmin(ctx context.Context, in *message.LoginAdminRequest, opts ...grpc.CallOption) (*message.LoginAdminResponse, error)
	FetchDistrictWeather(ctx context.Context, in *message.FetchDistrictWeatherRequest, opts ...grpc.CallOption) (*message.FetchDistrictWeatherResponse, error)
	UploadPost(ctx context.Context, in *message.UploadPostRequest, opts ...grpc.CallOption) (*message.UploadPostResponse, error)
	FetchPostList(ctx context.Context, in *message.FetchPostListRequest, opts ...grpc.CallOption) (*message.FetchPostListResponse, error)
	FetchPost(ctx context.Context, in *message.FetchPostRequest, opts ...grpc.CallOption) (*message.FetchPostResponse, error)
	DeletePost(ctx context.Context, in *message.DeletePostRequest, opts ...grpc.CallOption) (*message.DeletePostResponse, error)
	UpdatePost(ctx context.Context, in *message.UpdatePostRequest, opts ...grpc.CallOption) (*message.UpdatePostResponse, error)
	CheckPostPassword(ctx context.Context, in *message.CheckPostPasswordRequest, opts ...grpc.CallOption) (*message.CheckPostPasswordResponse, error)
	IncrementView(ctx context.Context, in *message.IncrementViewRequest, opts ...grpc.CallOption) (*message.IncrementViewResponse, error)
	LeaveComment(ctx context.Context, in *message.LeaveCommentRequest, opts ...grpc.CallOption) (*message.LeaveCommentResponse, error)
	FetchCommentList(ctx context.Context, in *message.FetchCommentListRequest, opts ...grpc.CallOption) (*message.FetchCommentListResponse, error)
	DeleteComment(ctx context.Context, in *message.DeleteCommentRequest, opts ...grpc.CallOption) (*message.DeleteCommentResponse, error)
	IncrementLike(ctx context.Context, in *message.IncrementLikeRequest, opts ...grpc.CallOption) (*message.IncrementLikeResponse, error)
	DecrementLike(ctx context.Context, in *message.DecrementLikeRequest, opts ...grpc.CallOption) (*message.DecrementLikeResponse, error)
	CheckValidPostId(ctx context.Context, in *message.CheckValidPostIdRequest, opts ...grpc.CallOption) (*message.CheckValidPostIdResponse, error)
	CheckValidCommentId(ctx context.Context, in *message.CheckValidCommentIdRequest, opts ...grpc.CallOption) (*message.CheckValidCommentIdResponse, error)
}

type meClient struct {
	cc grpc.ClientConnInterface
}

func NewMeClient(cc grpc.ClientConnInterface) MeClient {
	return &meClient{cc}
}

func (c *meClient) FindAdminUUID(ctx context.Context, in *message.FindAdminUUIDRequest, opts ...grpc.CallOption) (*message.FindAdminUUIDResponse, error) {
	out := new(message.FindAdminUUIDResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/FindAdminUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) InsertAdminUUID(ctx context.Context, in *message.InsertAdminUUIDRequest, opts ...grpc.CallOption) (*message.InsertAdminUUIDResponse, error) {
	out := new(message.InsertAdminUUIDResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/InsertAdminUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) LoginAdmin(ctx context.Context, in *message.LoginAdminRequest, opts ...grpc.CallOption) (*message.LoginAdminResponse, error) {
	out := new(message.LoginAdminResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/LoginAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) FetchDistrictWeather(ctx context.Context, in *message.FetchDistrictWeatherRequest, opts ...grpc.CallOption) (*message.FetchDistrictWeatherResponse, error) {
	out := new(message.FetchDistrictWeatherResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/FetchDistrictWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) UploadPost(ctx context.Context, in *message.UploadPostRequest, opts ...grpc.CallOption) (*message.UploadPostResponse, error) {
	out := new(message.UploadPostResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/UploadPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) FetchPostList(ctx context.Context, in *message.FetchPostListRequest, opts ...grpc.CallOption) (*message.FetchPostListResponse, error) {
	out := new(message.FetchPostListResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/FetchPostList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) FetchPost(ctx context.Context, in *message.FetchPostRequest, opts ...grpc.CallOption) (*message.FetchPostResponse, error) {
	out := new(message.FetchPostResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/FetchPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) DeletePost(ctx context.Context, in *message.DeletePostRequest, opts ...grpc.CallOption) (*message.DeletePostResponse, error) {
	out := new(message.DeletePostResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) UpdatePost(ctx context.Context, in *message.UpdatePostRequest, opts ...grpc.CallOption) (*message.UpdatePostResponse, error) {
	out := new(message.UpdatePostResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) CheckPostPassword(ctx context.Context, in *message.CheckPostPasswordRequest, opts ...grpc.CallOption) (*message.CheckPostPasswordResponse, error) {
	out := new(message.CheckPostPasswordResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/CheckPostPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) IncrementView(ctx context.Context, in *message.IncrementViewRequest, opts ...grpc.CallOption) (*message.IncrementViewResponse, error) {
	out := new(message.IncrementViewResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/IncrementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) LeaveComment(ctx context.Context, in *message.LeaveCommentRequest, opts ...grpc.CallOption) (*message.LeaveCommentResponse, error) {
	out := new(message.LeaveCommentResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/LeaveComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) FetchCommentList(ctx context.Context, in *message.FetchCommentListRequest, opts ...grpc.CallOption) (*message.FetchCommentListResponse, error) {
	out := new(message.FetchCommentListResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/FetchCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) DeleteComment(ctx context.Context, in *message.DeleteCommentRequest, opts ...grpc.CallOption) (*message.DeleteCommentResponse, error) {
	out := new(message.DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) IncrementLike(ctx context.Context, in *message.IncrementLikeRequest, opts ...grpc.CallOption) (*message.IncrementLikeResponse, error) {
	out := new(message.IncrementLikeResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/IncrementLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) DecrementLike(ctx context.Context, in *message.DecrementLikeRequest, opts ...grpc.CallOption) (*message.DecrementLikeResponse, error) {
	out := new(message.DecrementLikeResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/DecrementLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) CheckValidPostId(ctx context.Context, in *message.CheckValidPostIdRequest, opts ...grpc.CallOption) (*message.CheckValidPostIdResponse, error) {
	out := new(message.CheckValidPostIdResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/CheckValidPostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meClient) CheckValidCommentId(ctx context.Context, in *message.CheckValidCommentIdRequest, opts ...grpc.CallOption) (*message.CheckValidCommentIdResponse, error) {
	out := new(message.CheckValidCommentIdResponse)
	err := c.cc.Invoke(ctx, "/v2.service.me/CheckValidCommentId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeServer is the server API for Me service.
// All implementations must embed UnimplementedMeServer
// for forward compatibility
type MeServer interface {
	FindAdminUUID(context.Context, *message.FindAdminUUIDRequest) (*message.FindAdminUUIDResponse, error)
	InsertAdminUUID(context.Context, *message.InsertAdminUUIDRequest) (*message.InsertAdminUUIDResponse, error)
	LoginAdmin(context.Context, *message.LoginAdminRequest) (*message.LoginAdminResponse, error)
	FetchDistrictWeather(context.Context, *message.FetchDistrictWeatherRequest) (*message.FetchDistrictWeatherResponse, error)
	UploadPost(context.Context, *message.UploadPostRequest) (*message.UploadPostResponse, error)
	FetchPostList(context.Context, *message.FetchPostListRequest) (*message.FetchPostListResponse, error)
	FetchPost(context.Context, *message.FetchPostRequest) (*message.FetchPostResponse, error)
	DeletePost(context.Context, *message.DeletePostRequest) (*message.DeletePostResponse, error)
	UpdatePost(context.Context, *message.UpdatePostRequest) (*message.UpdatePostResponse, error)
	CheckPostPassword(context.Context, *message.CheckPostPasswordRequest) (*message.CheckPostPasswordResponse, error)
	IncrementView(context.Context, *message.IncrementViewRequest) (*message.IncrementViewResponse, error)
	LeaveComment(context.Context, *message.LeaveCommentRequest) (*message.LeaveCommentResponse, error)
	FetchCommentList(context.Context, *message.FetchCommentListRequest) (*message.FetchCommentListResponse, error)
	DeleteComment(context.Context, *message.DeleteCommentRequest) (*message.DeleteCommentResponse, error)
	IncrementLike(context.Context, *message.IncrementLikeRequest) (*message.IncrementLikeResponse, error)
	DecrementLike(context.Context, *message.DecrementLikeRequest) (*message.DecrementLikeResponse, error)
	CheckValidPostId(context.Context, *message.CheckValidPostIdRequest) (*message.CheckValidPostIdResponse, error)
	CheckValidCommentId(context.Context, *message.CheckValidCommentIdRequest) (*message.CheckValidCommentIdResponse, error)
	mustEmbedUnimplementedMeServer()
}

// UnimplementedMeServer must be embedded to have forward compatible implementations.
type UnimplementedMeServer struct {
}

func (UnimplementedMeServer) FindAdminUUID(context.Context, *message.FindAdminUUIDRequest) (*message.FindAdminUUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAdminUUID not implemented")
}
func (UnimplementedMeServer) InsertAdminUUID(context.Context, *message.InsertAdminUUIDRequest) (*message.InsertAdminUUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertAdminUUID not implemented")
}
func (UnimplementedMeServer) LoginAdmin(context.Context, *message.LoginAdminRequest) (*message.LoginAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAdmin not implemented")
}
func (UnimplementedMeServer) FetchDistrictWeather(context.Context, *message.FetchDistrictWeatherRequest) (*message.FetchDistrictWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchDistrictWeather not implemented")
}
func (UnimplementedMeServer) UploadPost(context.Context, *message.UploadPostRequest) (*message.UploadPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPost not implemented")
}
func (UnimplementedMeServer) FetchPostList(context.Context, *message.FetchPostListRequest) (*message.FetchPostListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPostList not implemented")
}
func (UnimplementedMeServer) FetchPost(context.Context, *message.FetchPostRequest) (*message.FetchPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPost not implemented")
}
func (UnimplementedMeServer) DeletePost(context.Context, *message.DeletePostRequest) (*message.DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedMeServer) UpdatePost(context.Context, *message.UpdatePostRequest) (*message.UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedMeServer) CheckPostPassword(context.Context, *message.CheckPostPasswordRequest) (*message.CheckPostPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPostPassword not implemented")
}
func (UnimplementedMeServer) IncrementView(context.Context, *message.IncrementViewRequest) (*message.IncrementViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementView not implemented")
}
func (UnimplementedMeServer) LeaveComment(context.Context, *message.LeaveCommentRequest) (*message.LeaveCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveComment not implemented")
}
func (UnimplementedMeServer) FetchCommentList(context.Context, *message.FetchCommentListRequest) (*message.FetchCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCommentList not implemented")
}
func (UnimplementedMeServer) DeleteComment(context.Context, *message.DeleteCommentRequest) (*message.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedMeServer) IncrementLike(context.Context, *message.IncrementLikeRequest) (*message.IncrementLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementLike not implemented")
}
func (UnimplementedMeServer) DecrementLike(context.Context, *message.DecrementLikeRequest) (*message.DecrementLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementLike not implemented")
}
func (UnimplementedMeServer) CheckValidPostId(context.Context, *message.CheckValidPostIdRequest) (*message.CheckValidPostIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckValidPostId not implemented")
}
func (UnimplementedMeServer) CheckValidCommentId(context.Context, *message.CheckValidCommentIdRequest) (*message.CheckValidCommentIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckValidCommentId not implemented")
}
func (UnimplementedMeServer) mustEmbedUnimplementedMeServer() {}

// UnsafeMeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeServer will
// result in compilation errors.
type UnsafeMeServer interface {
	mustEmbedUnimplementedMeServer()
}

func RegisterMeServer(s grpc.ServiceRegistrar, srv MeServer) {
	s.RegisterService(&Me_ServiceDesc, srv)
}

func _Me_FindAdminUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FindAdminUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).FindAdminUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/FindAdminUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).FindAdminUUID(ctx, req.(*message.FindAdminUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_InsertAdminUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.InsertAdminUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).InsertAdminUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/InsertAdminUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).InsertAdminUUID(ctx, req.(*message.InsertAdminUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_LoginAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.LoginAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).LoginAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/LoginAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).LoginAdmin(ctx, req.(*message.LoginAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_FetchDistrictWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FetchDistrictWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).FetchDistrictWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/FetchDistrictWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).FetchDistrictWeather(ctx, req.(*message.FetchDistrictWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_UploadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.UploadPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).UploadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/UploadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).UploadPost(ctx, req.(*message.UploadPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_FetchPostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FetchPostListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).FetchPostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/FetchPostList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).FetchPostList(ctx, req.(*message.FetchPostListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_FetchPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FetchPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).FetchPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/FetchPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).FetchPost(ctx, req.(*message.FetchPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).DeletePost(ctx, req.(*message.DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).UpdatePost(ctx, req.(*message.UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_CheckPostPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CheckPostPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).CheckPostPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/CheckPostPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).CheckPostPassword(ctx, req.(*message.CheckPostPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_IncrementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.IncrementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).IncrementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/IncrementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).IncrementView(ctx, req.(*message.IncrementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_LeaveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.LeaveCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).LeaveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/LeaveComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).LeaveComment(ctx, req.(*message.LeaveCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_FetchCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.FetchCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).FetchCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/FetchCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).FetchCommentList(ctx, req.(*message.FetchCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).DeleteComment(ctx, req.(*message.DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_IncrementLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.IncrementLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).IncrementLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/IncrementLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).IncrementLike(ctx, req.(*message.IncrementLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_DecrementLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.DecrementLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).DecrementLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/DecrementLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).DecrementLike(ctx, req.(*message.DecrementLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_CheckValidPostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CheckValidPostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).CheckValidPostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/CheckValidPostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).CheckValidPostId(ctx, req.(*message.CheckValidPostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Me_CheckValidCommentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.CheckValidCommentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeServer).CheckValidCommentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2.service.me/CheckValidCommentId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeServer).CheckValidCommentId(ctx, req.(*message.CheckValidCommentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Me_ServiceDesc is the grpc.ServiceDesc for Me service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Me_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2.service.me",
	HandlerType: (*MeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAdminUUID",
			Handler:    _Me_FindAdminUUID_Handler,
		},
		{
			MethodName: "InsertAdminUUID",
			Handler:    _Me_InsertAdminUUID_Handler,
		},
		{
			MethodName: "LoginAdmin",
			Handler:    _Me_LoginAdmin_Handler,
		},
		{
			MethodName: "FetchDistrictWeather",
			Handler:    _Me_FetchDistrictWeather_Handler,
		},
		{
			MethodName: "UploadPost",
			Handler:    _Me_UploadPost_Handler,
		},
		{
			MethodName: "FetchPostList",
			Handler:    _Me_FetchPostList_Handler,
		},
		{
			MethodName: "FetchPost",
			Handler:    _Me_FetchPost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Me_DeletePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Me_UpdatePost_Handler,
		},
		{
			MethodName: "CheckPostPassword",
			Handler:    _Me_CheckPostPassword_Handler,
		},
		{
			MethodName: "IncrementView",
			Handler:    _Me_IncrementView_Handler,
		},
		{
			MethodName: "LeaveComment",
			Handler:    _Me_LeaveComment_Handler,
		},
		{
			MethodName: "FetchCommentList",
			Handler:    _Me_FetchCommentList_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Me_DeleteComment_Handler,
		},
		{
			MethodName: "IncrementLike",
			Handler:    _Me_IncrementLike_Handler,
		},
		{
			MethodName: "DecrementLike",
			Handler:    _Me_DecrementLike_Handler,
		},
		{
			MethodName: "CheckValidPostId",
			Handler:    _Me_CheckValidPostId_Handler,
		},
		{
			MethodName: "CheckValidCommentId",
			Handler:    _Me_CheckValidCommentId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/service.proto",
}
