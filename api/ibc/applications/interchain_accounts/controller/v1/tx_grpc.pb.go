// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ibc/applications/interchain_accounts/controller/v1/tx.proto

package controllerv1

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
	Msg_RegisterInterchainAccount_FullMethodName = "/ibc.applications.interchain_accounts.controller.v1.Msg/RegisterInterchainAccount"
	Msg_SendTx_FullMethodName                    = "/ibc.applications.interchain_accounts.controller.v1.Msg/SendTx"
	Msg_UpdateParams_FullMethodName              = "/ibc.applications.interchain_accounts.controller.v1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterInterchainAccount defines a rpc handler for MsgRegisterInterchainAccount.
	RegisterInterchainAccount(ctx context.Context, in *MsgRegisterInterchainAccount, opts ...grpc.CallOption) (*MsgRegisterInterchainAccountResponse, error)
	// SendTx defines a rpc handler for MsgSendTx.
	SendTx(ctx context.Context, in *MsgSendTx, opts ...grpc.CallOption) (*MsgSendTxResponse, error)
	// UpdateParams defines a rpc handler for MsgUpdateParams.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterInterchainAccount(ctx context.Context, in *MsgRegisterInterchainAccount, opts ...grpc.CallOption) (*MsgRegisterInterchainAccountResponse, error) {
	out := new(MsgRegisterInterchainAccountResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterInterchainAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendTx(ctx context.Context, in *MsgSendTx, opts ...grpc.CallOption) (*MsgSendTxResponse, error) {
	out := new(MsgSendTxResponse)
	err := c.cc.Invoke(ctx, Msg_SendTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// RegisterInterchainAccount defines a rpc handler for MsgRegisterInterchainAccount.
	RegisterInterchainAccount(context.Context, *MsgRegisterInterchainAccount) (*MsgRegisterInterchainAccountResponse, error)
	// SendTx defines a rpc handler for MsgSendTx.
	SendTx(context.Context, *MsgSendTx) (*MsgSendTxResponse, error)
	// UpdateParams defines a rpc handler for MsgUpdateParams.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterInterchainAccount(context.Context, *MsgRegisterInterchainAccount) (*MsgRegisterInterchainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInterchainAccount not implemented")
}
func (UnimplementedMsgServer) SendTx(context.Context, *MsgSendTx) (*MsgSendTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RegisterInterchainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterInterchainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterInterchainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterInterchainAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterInterchainAccount(ctx, req.(*MsgRegisterInterchainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SendTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendTx(ctx, req.(*MsgSendTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.interchain_accounts.controller.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInterchainAccount",
			Handler:    _Msg_RegisterInterchainAccount_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _Msg_SendTx_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/interchain_accounts/controller/v1/tx.proto",
}