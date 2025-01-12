// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: emissions/v1/tx.proto

package emissionsv1

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
	Msg_UpdateParams_FullMethodName               = "/emissions.v1.Msg/UpdateParams"
	Msg_InsertBulkWorkerPayload_FullMethodName    = "/emissions.v1.Msg/InsertBulkWorkerPayload"
	Msg_CreateNewTopic_FullMethodName             = "/emissions.v1.Msg/CreateNewTopic"
	Msg_Register_FullMethodName                   = "/emissions.v1.Msg/Register"
	Msg_RemoveRegistration_FullMethodName         = "/emissions.v1.Msg/RemoveRegistration"
	Msg_InsertBulkReputerPayload_FullMethodName   = "/emissions.v1.Msg/InsertBulkReputerPayload"
	Msg_AddStake_FullMethodName                   = "/emissions.v1.Msg/AddStake"
	Msg_StartRemoveStake_FullMethodName           = "/emissions.v1.Msg/StartRemoveStake"
	Msg_ConfirmRemoveStake_FullMethodName         = "/emissions.v1.Msg/ConfirmRemoveStake"
	Msg_DelegateStake_FullMethodName              = "/emissions.v1.Msg/DelegateStake"
	Msg_RewardDelegateStake_FullMethodName        = "/emissions.v1.Msg/RewardDelegateStake"
	Msg_StartRemoveDelegateStake_FullMethodName   = "/emissions.v1.Msg/StartRemoveDelegateStake"
	Msg_ConfirmRemoveDelegateStake_FullMethodName = "/emissions.v1.Msg/ConfirmRemoveDelegateStake"
	Msg_FundTopic_FullMethodName                  = "/emissions.v1.Msg/FundTopic"
	Msg_AddToWhitelistAdmin_FullMethodName        = "/emissions.v1.Msg/AddToWhitelistAdmin"
	Msg_RemoveFromWhitelistAdmin_FullMethodName   = "/emissions.v1.Msg/RemoveFromWhitelistAdmin"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	InsertBulkWorkerPayload(ctx context.Context, in *MsgInsertBulkWorkerPayload, opts ...grpc.CallOption) (*MsgInsertBulkWorkerPayloadResponse, error)
	CreateNewTopic(ctx context.Context, in *MsgCreateNewTopic, opts ...grpc.CallOption) (*MsgCreateNewTopicResponse, error)
	Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgRegisterResponse, error)
	RemoveRegistration(ctx context.Context, in *MsgRemoveRegistration, opts ...grpc.CallOption) (*MsgRemoveRegistrationResponse, error)
	InsertBulkReputerPayload(ctx context.Context, in *MsgInsertBulkReputerPayload, opts ...grpc.CallOption) (*MsgInsertBulkReputerPayloadResponse, error)
	AddStake(ctx context.Context, in *MsgAddStake, opts ...grpc.CallOption) (*MsgAddStakeResponse, error)
	StartRemoveStake(ctx context.Context, in *MsgStartRemoveStake, opts ...grpc.CallOption) (*MsgStartRemoveStakeResponse, error)
	ConfirmRemoveStake(ctx context.Context, in *MsgConfirmRemoveStake, opts ...grpc.CallOption) (*MsgConfirmRemoveStakeResponse, error)
	DelegateStake(ctx context.Context, in *MsgDelegateStake, opts ...grpc.CallOption) (*MsgDelegateStakeResponse, error)
	RewardDelegateStake(ctx context.Context, in *MsgRewardDelegateStake, opts ...grpc.CallOption) (*MsgRewardDelegateStakeResponse, error)
	StartRemoveDelegateStake(ctx context.Context, in *MsgStartRemoveDelegateStake, opts ...grpc.CallOption) (*MsgStartRemoveDelegateStakeResponse, error)
	ConfirmRemoveDelegateStake(ctx context.Context, in *MsgConfirmDelegateRemoveStake, opts ...grpc.CallOption) (*MsgConfirmRemoveDelegateStakeResponse, error)
	FundTopic(ctx context.Context, in *MsgFundTopic, opts ...grpc.CallOption) (*MsgFundTopicResponse, error)
	AddToWhitelistAdmin(ctx context.Context, in *MsgAddToWhitelistAdmin, opts ...grpc.CallOption) (*MsgAddToWhitelistAdminResponse, error)
	RemoveFromWhitelistAdmin(ctx context.Context, in *MsgRemoveFromWhitelistAdmin, opts ...grpc.CallOption) (*MsgRemoveFromWhitelistAdminResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InsertBulkWorkerPayload(ctx context.Context, in *MsgInsertBulkWorkerPayload, opts ...grpc.CallOption) (*MsgInsertBulkWorkerPayloadResponse, error) {
	out := new(MsgInsertBulkWorkerPayloadResponse)
	err := c.cc.Invoke(ctx, Msg_InsertBulkWorkerPayload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateNewTopic(ctx context.Context, in *MsgCreateNewTopic, opts ...grpc.CallOption) (*MsgCreateNewTopicResponse, error) {
	out := new(MsgCreateNewTopicResponse)
	err := c.cc.Invoke(ctx, Msg_CreateNewTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgRegisterResponse, error) {
	out := new(MsgRegisterResponse)
	err := c.cc.Invoke(ctx, Msg_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveRegistration(ctx context.Context, in *MsgRemoveRegistration, opts ...grpc.CallOption) (*MsgRemoveRegistrationResponse, error) {
	out := new(MsgRemoveRegistrationResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InsertBulkReputerPayload(ctx context.Context, in *MsgInsertBulkReputerPayload, opts ...grpc.CallOption) (*MsgInsertBulkReputerPayloadResponse, error) {
	out := new(MsgInsertBulkReputerPayloadResponse)
	err := c.cc.Invoke(ctx, Msg_InsertBulkReputerPayload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddStake(ctx context.Context, in *MsgAddStake, opts ...grpc.CallOption) (*MsgAddStakeResponse, error) {
	out := new(MsgAddStakeResponse)
	err := c.cc.Invoke(ctx, Msg_AddStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartRemoveStake(ctx context.Context, in *MsgStartRemoveStake, opts ...grpc.CallOption) (*MsgStartRemoveStakeResponse, error) {
	out := new(MsgStartRemoveStakeResponse)
	err := c.cc.Invoke(ctx, Msg_StartRemoveStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConfirmRemoveStake(ctx context.Context, in *MsgConfirmRemoveStake, opts ...grpc.CallOption) (*MsgConfirmRemoveStakeResponse, error) {
	out := new(MsgConfirmRemoveStakeResponse)
	err := c.cc.Invoke(ctx, Msg_ConfirmRemoveStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelegateStake(ctx context.Context, in *MsgDelegateStake, opts ...grpc.CallOption) (*MsgDelegateStakeResponse, error) {
	out := new(MsgDelegateStakeResponse)
	err := c.cc.Invoke(ctx, Msg_DelegateStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RewardDelegateStake(ctx context.Context, in *MsgRewardDelegateStake, opts ...grpc.CallOption) (*MsgRewardDelegateStakeResponse, error) {
	out := new(MsgRewardDelegateStakeResponse)
	err := c.cc.Invoke(ctx, Msg_RewardDelegateStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartRemoveDelegateStake(ctx context.Context, in *MsgStartRemoveDelegateStake, opts ...grpc.CallOption) (*MsgStartRemoveDelegateStakeResponse, error) {
	out := new(MsgStartRemoveDelegateStakeResponse)
	err := c.cc.Invoke(ctx, Msg_StartRemoveDelegateStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConfirmRemoveDelegateStake(ctx context.Context, in *MsgConfirmDelegateRemoveStake, opts ...grpc.CallOption) (*MsgConfirmRemoveDelegateStakeResponse, error) {
	out := new(MsgConfirmRemoveDelegateStakeResponse)
	err := c.cc.Invoke(ctx, Msg_ConfirmRemoveDelegateStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FundTopic(ctx context.Context, in *MsgFundTopic, opts ...grpc.CallOption) (*MsgFundTopicResponse, error) {
	out := new(MsgFundTopicResponse)
	err := c.cc.Invoke(ctx, Msg_FundTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddToWhitelistAdmin(ctx context.Context, in *MsgAddToWhitelistAdmin, opts ...grpc.CallOption) (*MsgAddToWhitelistAdminResponse, error) {
	out := new(MsgAddToWhitelistAdminResponse)
	err := c.cc.Invoke(ctx, Msg_AddToWhitelistAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveFromWhitelistAdmin(ctx context.Context, in *MsgRemoveFromWhitelistAdmin, opts ...grpc.CallOption) (*MsgRemoveFromWhitelistAdminResponse, error) {
	out := new(MsgRemoveFromWhitelistAdminResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveFromWhitelistAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	InsertBulkWorkerPayload(context.Context, *MsgInsertBulkWorkerPayload) (*MsgInsertBulkWorkerPayloadResponse, error)
	CreateNewTopic(context.Context, *MsgCreateNewTopic) (*MsgCreateNewTopicResponse, error)
	Register(context.Context, *MsgRegister) (*MsgRegisterResponse, error)
	RemoveRegistration(context.Context, *MsgRemoveRegistration) (*MsgRemoveRegistrationResponse, error)
	InsertBulkReputerPayload(context.Context, *MsgInsertBulkReputerPayload) (*MsgInsertBulkReputerPayloadResponse, error)
	AddStake(context.Context, *MsgAddStake) (*MsgAddStakeResponse, error)
	StartRemoveStake(context.Context, *MsgStartRemoveStake) (*MsgStartRemoveStakeResponse, error)
	ConfirmRemoveStake(context.Context, *MsgConfirmRemoveStake) (*MsgConfirmRemoveStakeResponse, error)
	DelegateStake(context.Context, *MsgDelegateStake) (*MsgDelegateStakeResponse, error)
	RewardDelegateStake(context.Context, *MsgRewardDelegateStake) (*MsgRewardDelegateStakeResponse, error)
	StartRemoveDelegateStake(context.Context, *MsgStartRemoveDelegateStake) (*MsgStartRemoveDelegateStakeResponse, error)
	ConfirmRemoveDelegateStake(context.Context, *MsgConfirmDelegateRemoveStake) (*MsgConfirmRemoveDelegateStakeResponse, error)
	FundTopic(context.Context, *MsgFundTopic) (*MsgFundTopicResponse, error)
	AddToWhitelistAdmin(context.Context, *MsgAddToWhitelistAdmin) (*MsgAddToWhitelistAdminResponse, error)
	RemoveFromWhitelistAdmin(context.Context, *MsgRemoveFromWhitelistAdmin) (*MsgRemoveFromWhitelistAdminResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) InsertBulkWorkerPayload(context.Context, *MsgInsertBulkWorkerPayload) (*MsgInsertBulkWorkerPayloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBulkWorkerPayload not implemented")
}
func (UnimplementedMsgServer) CreateNewTopic(context.Context, *MsgCreateNewTopic) (*MsgCreateNewTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTopic not implemented")
}
func (UnimplementedMsgServer) Register(context.Context, *MsgRegister) (*MsgRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMsgServer) RemoveRegistration(context.Context, *MsgRemoveRegistration) (*MsgRemoveRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRegistration not implemented")
}
func (UnimplementedMsgServer) InsertBulkReputerPayload(context.Context, *MsgInsertBulkReputerPayload) (*MsgInsertBulkReputerPayloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBulkReputerPayload not implemented")
}
func (UnimplementedMsgServer) AddStake(context.Context, *MsgAddStake) (*MsgAddStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStake not implemented")
}
func (UnimplementedMsgServer) StartRemoveStake(context.Context, *MsgStartRemoveStake) (*MsgStartRemoveStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRemoveStake not implemented")
}
func (UnimplementedMsgServer) ConfirmRemoveStake(context.Context, *MsgConfirmRemoveStake) (*MsgConfirmRemoveStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmRemoveStake not implemented")
}
func (UnimplementedMsgServer) DelegateStake(context.Context, *MsgDelegateStake) (*MsgDelegateStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateStake not implemented")
}
func (UnimplementedMsgServer) RewardDelegateStake(context.Context, *MsgRewardDelegateStake) (*MsgRewardDelegateStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RewardDelegateStake not implemented")
}
func (UnimplementedMsgServer) StartRemoveDelegateStake(context.Context, *MsgStartRemoveDelegateStake) (*MsgStartRemoveDelegateStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRemoveDelegateStake not implemented")
}
func (UnimplementedMsgServer) ConfirmRemoveDelegateStake(context.Context, *MsgConfirmDelegateRemoveStake) (*MsgConfirmRemoveDelegateStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmRemoveDelegateStake not implemented")
}
func (UnimplementedMsgServer) FundTopic(context.Context, *MsgFundTopic) (*MsgFundTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundTopic not implemented")
}
func (UnimplementedMsgServer) AddToWhitelistAdmin(context.Context, *MsgAddToWhitelistAdmin) (*MsgAddToWhitelistAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWhitelistAdmin not implemented")
}
func (UnimplementedMsgServer) RemoveFromWhitelistAdmin(context.Context, *MsgRemoveFromWhitelistAdmin) (*MsgRemoveFromWhitelistAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWhitelistAdmin not implemented")
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

func _Msg_InsertBulkWorkerPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertBulkWorkerPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertBulkWorkerPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InsertBulkWorkerPayload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertBulkWorkerPayload(ctx, req.(*MsgInsertBulkWorkerPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateNewTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateNewTopic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateNewTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateNewTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateNewTopic(ctx, req.(*MsgCreateNewTopic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Register(ctx, req.(*MsgRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveRegistration(ctx, req.(*MsgRemoveRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InsertBulkReputerPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertBulkReputerPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertBulkReputerPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InsertBulkReputerPayload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertBulkReputerPayload(ctx, req.(*MsgInsertBulkReputerPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddStake(ctx, req.(*MsgAddStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartRemoveStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartRemoveStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartRemoveStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StartRemoveStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartRemoveStake(ctx, req.(*MsgStartRemoveStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConfirmRemoveStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfirmRemoveStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfirmRemoveStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConfirmRemoveStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfirmRemoveStake(ctx, req.(*MsgConfirmRemoveStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelegateStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelegateStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DelegateStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateStake(ctx, req.(*MsgDelegateStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RewardDelegateStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRewardDelegateStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RewardDelegateStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RewardDelegateStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RewardDelegateStake(ctx, req.(*MsgRewardDelegateStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartRemoveDelegateStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartRemoveDelegateStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartRemoveDelegateStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StartRemoveDelegateStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartRemoveDelegateStake(ctx, req.(*MsgStartRemoveDelegateStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConfirmRemoveDelegateStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfirmDelegateRemoveStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfirmRemoveDelegateStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConfirmRemoveDelegateStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfirmRemoveDelegateStake(ctx, req.(*MsgConfirmDelegateRemoveStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FundTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundTopic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FundTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundTopic(ctx, req.(*MsgFundTopic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddToWhitelistAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddToWhitelistAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddToWhitelistAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddToWhitelistAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddToWhitelistAdmin(ctx, req.(*MsgAddToWhitelistAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveFromWhitelistAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveFromWhitelistAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveFromWhitelistAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveFromWhitelistAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveFromWhitelistAdmin(ctx, req.(*MsgRemoveFromWhitelistAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emissions.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "InsertBulkWorkerPayload",
			Handler:    _Msg_InsertBulkWorkerPayload_Handler,
		},
		{
			MethodName: "CreateNewTopic",
			Handler:    _Msg_CreateNewTopic_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Msg_Register_Handler,
		},
		{
			MethodName: "RemoveRegistration",
			Handler:    _Msg_RemoveRegistration_Handler,
		},
		{
			MethodName: "InsertBulkReputerPayload",
			Handler:    _Msg_InsertBulkReputerPayload_Handler,
		},
		{
			MethodName: "AddStake",
			Handler:    _Msg_AddStake_Handler,
		},
		{
			MethodName: "StartRemoveStake",
			Handler:    _Msg_StartRemoveStake_Handler,
		},
		{
			MethodName: "ConfirmRemoveStake",
			Handler:    _Msg_ConfirmRemoveStake_Handler,
		},
		{
			MethodName: "DelegateStake",
			Handler:    _Msg_DelegateStake_Handler,
		},
		{
			MethodName: "RewardDelegateStake",
			Handler:    _Msg_RewardDelegateStake_Handler,
		},
		{
			MethodName: "StartRemoveDelegateStake",
			Handler:    _Msg_StartRemoveDelegateStake_Handler,
		},
		{
			MethodName: "ConfirmRemoveDelegateStake",
			Handler:    _Msg_ConfirmRemoveDelegateStake_Handler,
		},
		{
			MethodName: "FundTopic",
			Handler:    _Msg_FundTopic_Handler,
		},
		{
			MethodName: "AddToWhitelistAdmin",
			Handler:    _Msg_AddToWhitelistAdmin_Handler,
		},
		{
			MethodName: "RemoveFromWhitelistAdmin",
			Handler:    _Msg_RemoveFromWhitelistAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emissions/v1/tx.proto",
}
