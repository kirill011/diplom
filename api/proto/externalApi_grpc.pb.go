// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: api/proto/externalApi.proto

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

const (
	Api_GetHardwareValue_FullMethodName     = "/api.api/GetHardwareValue"
	Api_UpdateParamValue_FullMethodName     = "/api.api/UpdateParamValue"
	Api_Registration_FullMethodName         = "/api.api/Registration"
	Api_RegistrationHardware_FullMethodName = "/api.api/RegistrationHardware"
	Api_GetHardwareId_FullMethodName        = "/api.api/GetHardwareId"
	Api_GetParamId_FullMethodName           = "/api.api/GetParamId"
	Api_RegistrationParams_FullMethodName   = "/api.api/RegistrationParams"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	// Метод позволяет получить список с параметрами
	// и их значением для необходимого оборудования.
	GetHardwareValue(ctx context.Context, in *HardwareRequest, opts ...grpc.CallOption) (*HardwareResponse, error)
	// Метод позволяет менять параметры оборудования.
	// Неуказанные параметры остаются неизменными.
	UpdateParamValue(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Метод позволяет зарегистрировать нового пользователя.
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	// Метод позволяет зарегистрировать оборудование пользователя
	RegistrationHardware(ctx context.Context, in *RegistrationHardwareRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	// Метод позволяет получить все id оборудования пользователя
	GetHardwareId(ctx context.Context, in *HardwareIdRequest, opts ...grpc.CallOption) (*HardwereIdResponce, error)
	// Метод позволяет получить все id параметров выбранного оборудования
	GetParamId(ctx context.Context, in *ParamIdRequest, opts ...grpc.CallOption) (*ParamIdResponce, error)
	// Метод позволяет зарегистрировать оборудование пользователя
	RegistrationParams(ctx context.Context, in *RegParamsReq, opts ...grpc.CallOption) (*RegParamsResponce, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetHardwareValue(ctx context.Context, in *HardwareRequest, opts ...grpc.CallOption) (*HardwareResponse, error) {
	out := new(HardwareResponse)
	err := c.cc.Invoke(ctx, Api_GetHardwareValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateParamValue(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, Api_UpdateParamValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, Api_Registration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RegistrationHardware(ctx context.Context, in *RegistrationHardwareRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, Api_RegistrationHardware_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetHardwareId(ctx context.Context, in *HardwareIdRequest, opts ...grpc.CallOption) (*HardwereIdResponce, error) {
	out := new(HardwereIdResponce)
	err := c.cc.Invoke(ctx, Api_GetHardwareId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetParamId(ctx context.Context, in *ParamIdRequest, opts ...grpc.CallOption) (*ParamIdResponce, error) {
	out := new(ParamIdResponce)
	err := c.cc.Invoke(ctx, Api_GetParamId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RegistrationParams(ctx context.Context, in *RegParamsReq, opts ...grpc.CallOption) (*RegParamsResponce, error) {
	out := new(RegParamsResponce)
	err := c.cc.Invoke(ctx, Api_RegistrationParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	// Метод позволяет получить список с параметрами
	// и их значением для необходимого оборудования.
	GetHardwareValue(context.Context, *HardwareRequest) (*HardwareResponse, error)
	// Метод позволяет менять параметры оборудования.
	// Неуказанные параметры остаются неизменными.
	UpdateParamValue(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Метод позволяет зарегистрировать нового пользователя.
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	// Метод позволяет зарегистрировать оборудование пользователя
	RegistrationHardware(context.Context, *RegistrationHardwareRequest) (*RegistrationResponse, error)
	// Метод позволяет получить все id оборудования пользователя
	GetHardwareId(context.Context, *HardwareIdRequest) (*HardwereIdResponce, error)
	// Метод позволяет получить все id параметров выбранного оборудования
	GetParamId(context.Context, *ParamIdRequest) (*ParamIdResponce, error)
	// Метод позволяет зарегистрировать оборудование пользователя
	RegistrationParams(context.Context, *RegParamsReq) (*RegParamsResponce, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) GetHardwareValue(context.Context, *HardwareRequest) (*HardwareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHardwareValue not implemented")
}
func (UnimplementedApiServer) UpdateParamValue(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParamValue not implemented")
}
func (UnimplementedApiServer) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedApiServer) RegistrationHardware(context.Context, *RegistrationHardwareRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationHardware not implemented")
}
func (UnimplementedApiServer) GetHardwareId(context.Context, *HardwareIdRequest) (*HardwereIdResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHardwareId not implemented")
}
func (UnimplementedApiServer) GetParamId(context.Context, *ParamIdRequest) (*ParamIdResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParamId not implemented")
}
func (UnimplementedApiServer) RegistrationParams(context.Context, *RegParamsReq) (*RegParamsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrationParams not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_GetHardwareValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HardwareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetHardwareValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetHardwareValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetHardwareValue(ctx, req.(*HardwareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateParamValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateParamValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpdateParamValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateParamValue(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Registration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RegistrationHardware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationHardwareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RegistrationHardware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_RegistrationHardware_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RegistrationHardware(ctx, req.(*RegistrationHardwareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetHardwareId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HardwareIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetHardwareId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetHardwareId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetHardwareId(ctx, req.(*HardwareIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetParamId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetParamId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetParamId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetParamId(ctx, req.(*ParamIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RegistrationParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegParamsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RegistrationParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_RegistrationParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RegistrationParams(ctx, req.(*RegParamsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHardwareValue",
			Handler:    _Api_GetHardwareValue_Handler,
		},
		{
			MethodName: "UpdateParamValue",
			Handler:    _Api_UpdateParamValue_Handler,
		},
		{
			MethodName: "Registration",
			Handler:    _Api_Registration_Handler,
		},
		{
			MethodName: "RegistrationHardware",
			Handler:    _Api_RegistrationHardware_Handler,
		},
		{
			MethodName: "GetHardwareId",
			Handler:    _Api_GetHardwareId_Handler,
		},
		{
			MethodName: "GetParamId",
			Handler:    _Api_GetParamId_Handler,
		},
		{
			MethodName: "RegistrationParams",
			Handler:    _Api_RegistrationParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/externalApi.proto",
}
