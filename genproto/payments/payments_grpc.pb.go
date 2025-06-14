// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: payments.proto

package payments

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_CreatePayment_FullMethodName           = "/payment.PaymentService/CreatePayment"
	PaymentService_GetPaymentInfo_FullMethodName          = "/payment.PaymentService/GetPaymentInfo"
	PaymentService_CreatePaymentURL_FullMethodName        = "/payment.PaymentService/CreatePaymentURL"
	PaymentService_CancelPaymentLink_FullMethodName       = "/payment.PaymentService/CancelPaymentLink"
	PaymentService_UpdatePaymentStatus_FullMethodName     = "/payment.PaymentService/UpdatePaymentStatus"
	PaymentService_UpdatePaymentMethod_FullMethodName     = "/payment.PaymentService/UpdatePaymentMethod"
	PaymentService_UpdatePaymentAmount_FullMethodName     = "/payment.PaymentService/UpdatePaymentAmount"
	PaymentService_UpdateBankPaymentStatus_FullMethodName = "/payment.PaymentService/UpdateBankPaymentStatus"
	PaymentService_GetAllPayment_FullMethodName           = "/payment.PaymentService/GetAllPayment"
	PaymentService_GetTotalRevenue_FullMethodName         = "/payment.PaymentService/GetTotalRevenue"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service thanh toán
type PaymentServiceClient interface {
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	GetPaymentInfo(ctx context.Context, in *GetPaymentInfoRequest, opts ...grpc.CallOption) (*GetPaymentInfoResponse, error)
	CreatePaymentURL(ctx context.Context, in *CreatePaymentURLRequest, opts ...grpc.CallOption) (*CreatePaymentURLResponse, error)
	CancelPaymentLink(ctx context.Context, in *CancelPaymentLinkRequest, opts ...grpc.CallOption) (*CancelPaymentLinkResponse, error)
	UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error)
	UpdatePaymentMethod(ctx context.Context, in *UpdatePaymentMethodRequest, opts ...grpc.CallOption) (*UpdatePaymentMethodResponse, error)
	UpdatePaymentAmount(ctx context.Context, in *UpdatePaymentAmountRequest, opts ...grpc.CallOption) (*UpdatePaymentAmountResponse, error)
	UpdateBankPaymentStatus(ctx context.Context, in *UpdateBankPaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error)
	GetAllPayment(ctx context.Context, in *GetAllPaymentsRequest, opts ...grpc.CallOption) (*GetAllPaymentsResponse, error)
	GetTotalRevenue(ctx context.Context, in *GetTotalRevenueRequest, opts ...grpc.CallOption) (*GetTotalRevenueResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_CreatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPaymentInfo(ctx context.Context, in *GetPaymentInfoRequest, opts ...grpc.CallOption) (*GetPaymentInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentInfoResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetPaymentInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreatePaymentURL(ctx context.Context, in *CreatePaymentURLRequest, opts ...grpc.CallOption) (*CreatePaymentURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentURLResponse)
	err := c.cc.Invoke(ctx, PaymentService_CreatePaymentURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CancelPaymentLink(ctx context.Context, in *CancelPaymentLinkRequest, opts ...grpc.CallOption) (*CancelPaymentLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelPaymentLinkResponse)
	err := c.cc.Invoke(ctx, PaymentService_CancelPaymentLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentStatusResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePaymentMethod(ctx context.Context, in *UpdatePaymentMethodRequest, opts ...grpc.CallOption) (*UpdatePaymentMethodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentMethodResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePaymentMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePaymentAmount(ctx context.Context, in *UpdatePaymentAmountRequest, opts ...grpc.CallOption) (*UpdatePaymentAmountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentAmountResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePaymentAmount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdateBankPaymentStatus(ctx context.Context, in *UpdateBankPaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentStatusResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdateBankPaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetAllPayment(ctx context.Context, in *GetAllPaymentsRequest, opts ...grpc.CallOption) (*GetAllPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllPaymentsResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetAllPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetTotalRevenue(ctx context.Context, in *GetTotalRevenueRequest, opts ...grpc.CallOption) (*GetTotalRevenueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTotalRevenueResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetTotalRevenue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
//
// Service thanh toán
type PaymentServiceServer interface {
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	GetPaymentInfo(context.Context, *GetPaymentInfoRequest) (*GetPaymentInfoResponse, error)
	CreatePaymentURL(context.Context, *CreatePaymentURLRequest) (*CreatePaymentURLResponse, error)
	CancelPaymentLink(context.Context, *CancelPaymentLinkRequest) (*CancelPaymentLinkResponse, error)
	UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error)
	UpdatePaymentMethod(context.Context, *UpdatePaymentMethodRequest) (*UpdatePaymentMethodResponse, error)
	UpdatePaymentAmount(context.Context, *UpdatePaymentAmountRequest) (*UpdatePaymentAmountResponse, error)
	UpdateBankPaymentStatus(context.Context, *UpdateBankPaymentStatusRequest) (*UpdatePaymentStatusResponse, error)
	GetAllPayment(context.Context, *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error)
	GetTotalRevenue(context.Context, *GetTotalRevenueRequest) (*GetTotalRevenueResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPaymentInfo(context.Context, *GetPaymentInfoRequest) (*GetPaymentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentInfo not implemented")
}
func (UnimplementedPaymentServiceServer) CreatePaymentURL(context.Context, *CreatePaymentURLRequest) (*CreatePaymentURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePaymentURL not implemented")
}
func (UnimplementedPaymentServiceServer) CancelPaymentLink(context.Context, *CancelPaymentLinkRequest) (*CancelPaymentLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPaymentLink not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentStatus not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePaymentMethod(context.Context, *UpdatePaymentMethodRequest) (*UpdatePaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentMethod not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePaymentAmount(context.Context, *UpdatePaymentAmountRequest) (*UpdatePaymentAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentAmount not implemented")
}
func (UnimplementedPaymentServiceServer) UpdateBankPaymentStatus(context.Context, *UpdateBankPaymentStatusRequest) (*UpdatePaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBankPaymentStatus not implemented")
}
func (UnimplementedPaymentServiceServer) GetAllPayment(context.Context, *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetTotalRevenue(context.Context, *GetTotalRevenueRequest) (*GetTotalRevenueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalRevenue not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPaymentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPaymentInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentInfo(ctx, req.(*GetPaymentInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CreatePaymentURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePaymentURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePaymentURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePaymentURL(ctx, req.(*CreatePaymentURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CancelPaymentLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPaymentLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CancelPaymentLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CancelPaymentLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CancelPaymentLink(ctx, req.(*CancelPaymentLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, req.(*UpdatePaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePaymentMethod(ctx, req.(*UpdatePaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePaymentAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePaymentAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePaymentAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePaymentAmount(ctx, req.(*UpdatePaymentAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdateBankPaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankPaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdateBankPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdateBankPaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdateBankPaymentStatus(ctx, req.(*UpdateBankPaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetAllPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetAllPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetAllPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetAllPayment(ctx, req.(*GetAllPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetTotalRevenue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalRevenueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetTotalRevenue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetTotalRevenue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetTotalRevenue(ctx, req.(*GetTotalRevenueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
		{
			MethodName: "GetPaymentInfo",
			Handler:    _PaymentService_GetPaymentInfo_Handler,
		},
		{
			MethodName: "CreatePaymentURL",
			Handler:    _PaymentService_CreatePaymentURL_Handler,
		},
		{
			MethodName: "CancelPaymentLink",
			Handler:    _PaymentService_CancelPaymentLink_Handler,
		},
		{
			MethodName: "UpdatePaymentStatus",
			Handler:    _PaymentService_UpdatePaymentStatus_Handler,
		},
		{
			MethodName: "UpdatePaymentMethod",
			Handler:    _PaymentService_UpdatePaymentMethod_Handler,
		},
		{
			MethodName: "UpdatePaymentAmount",
			Handler:    _PaymentService_UpdatePaymentAmount_Handler,
		},
		{
			MethodName: "UpdateBankPaymentStatus",
			Handler:    _PaymentService_UpdateBankPaymentStatus_Handler,
		},
		{
			MethodName: "GetAllPayment",
			Handler:    _PaymentService_GetAllPayment_Handler,
		},
		{
			MethodName: "GetTotalRevenue",
			Handler:    _PaymentService_GetTotalRevenue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payments.proto",
}
