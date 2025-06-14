// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: appointments.proto

package appointments

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
	AppointmentService_CreateAppointment_FullMethodName            = "/appointments.AppointmentService/CreateAppointment"
	AppointmentService_GetAppointmentsByCustomer_FullMethodName    = "/appointments.AppointmentService/GetAppointmentsByCustomer"
	AppointmentService_GetAppointmentsByEmployee_FullMethodName    = "/appointments.AppointmentService/GetAppointmentsByEmployee"
	AppointmentService_GetAppointmentsByBranch_FullMethodName      = "/appointments.AppointmentService/GetAppointmentsByBranch"
	AppointmentService_GetAllAppointments_FullMethodName           = "/appointments.AppointmentService/GetAllAppointments"
	AppointmentService_UpdateAppointmentStatus_FullMethodName      = "/appointments.AppointmentService/UpdateAppointmentStatus"
	AppointmentService_UpdateEmployeeForAppointment_FullMethodName = "/appointments.AppointmentService/UpdateEmployeeForAppointment"
	AppointmentService_GetAppointmentDetails_FullMethodName        = "/appointments.AppointmentService/GetAppointmentDetails"
	AppointmentService_CreateService_FullMethodName                = "/appointments.AppointmentService/CreateService"
	AppointmentService_GetServices_FullMethodName                  = "/appointments.AppointmentService/GetServices"
	AppointmentService_UpdateService_FullMethodName                = "/appointments.AppointmentService/UpdateService"
	AppointmentService_DeleteService_FullMethodName                = "/appointments.AppointmentService/DeleteService"
)

// AppointmentServiceClient is the client API for AppointmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppointmentServiceClient interface {
	// Đặt lịch hẹn tại nhà khách hàng
	CreateAppointment(ctx context.Context, in *CreateAppointmentRequest, opts ...grpc.CallOption) (*CreateAppointmentResponse, error)
	// Lấy danh sách lịch hẹn theo khách hàng hoặc nhân viên
	GetAppointmentsByCustomer(ctx context.Context, in *GetAppointmentsByCustomerRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error)
	GetAppointmentsByEmployee(ctx context.Context, in *GetAppointmentsByEmployeeRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error)
	GetAppointmentsByBranch(ctx context.Context, in *GetAppointmentsByBranchRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error)
	GetAllAppointments(ctx context.Context, in *GetAllAppointmentsRequest, opts ...grpc.CallOption) (*GetAllAppointmentsResponse, error)
	// Cập nhật trạng thái lịch hẹn
	UpdateAppointmentStatus(ctx context.Context, in *UpdateAppointmentStatusRequest, opts ...grpc.CallOption) (*UpdateAppointmentStatusResponse, error)
	UpdateEmployeeForAppointment(ctx context.Context, in *UpdateEmployeeForAppointmentRequest, opts ...grpc.CallOption) (*UpdateEmployeeForAppointmentResponse, error)
	// Lấy chi tiết lịch hẹn
	GetAppointmentDetails(ctx context.Context, in *GetAppointmentDetailsRequest, opts ...grpc.CallOption) (*GetAppointmentDetailsResponse, error)
	// Quản lý danh sách dịch vụ
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
}

type appointmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentServiceClient(cc grpc.ClientConnInterface) AppointmentServiceClient {
	return &appointmentServiceClient{cc}
}

func (c *appointmentServiceClient) CreateAppointment(ctx context.Context, in *CreateAppointmentRequest, opts ...grpc.CallOption) (*CreateAppointmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAppointmentResponse)
	err := c.cc.Invoke(ctx, AppointmentService_CreateAppointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetAppointmentsByCustomer(ctx context.Context, in *GetAppointmentsByCustomerRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppointmentsResponse)
	err := c.cc.Invoke(ctx, AppointmentService_GetAppointmentsByCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetAppointmentsByEmployee(ctx context.Context, in *GetAppointmentsByEmployeeRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppointmentsResponse)
	err := c.cc.Invoke(ctx, AppointmentService_GetAppointmentsByEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetAppointmentsByBranch(ctx context.Context, in *GetAppointmentsByBranchRequest, opts ...grpc.CallOption) (*GetAppointmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppointmentsResponse)
	err := c.cc.Invoke(ctx, AppointmentService_GetAppointmentsByBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetAllAppointments(ctx context.Context, in *GetAllAppointmentsRequest, opts ...grpc.CallOption) (*GetAllAppointmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllAppointmentsResponse)
	err := c.cc.Invoke(ctx, AppointmentService_GetAllAppointments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) UpdateAppointmentStatus(ctx context.Context, in *UpdateAppointmentStatusRequest, opts ...grpc.CallOption) (*UpdateAppointmentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAppointmentStatusResponse)
	err := c.cc.Invoke(ctx, AppointmentService_UpdateAppointmentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) UpdateEmployeeForAppointment(ctx context.Context, in *UpdateEmployeeForAppointmentRequest, opts ...grpc.CallOption) (*UpdateEmployeeForAppointmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEmployeeForAppointmentResponse)
	err := c.cc.Invoke(ctx, AppointmentService_UpdateEmployeeForAppointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetAppointmentDetails(ctx context.Context, in *GetAppointmentDetailsRequest, opts ...grpc.CallOption) (*GetAppointmentDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppointmentDetailsResponse)
	err := c.cc.Invoke(ctx, AppointmentService_GetAppointmentDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, AppointmentService_CreateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, AppointmentService_GetServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateServiceResponse)
	err := c.cc.Invoke(ctx, AppointmentService_UpdateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, AppointmentService_DeleteService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServiceServer is the server API for AppointmentService service.
// All implementations must embed UnimplementedAppointmentServiceServer
// for forward compatibility.
type AppointmentServiceServer interface {
	// Đặt lịch hẹn tại nhà khách hàng
	CreateAppointment(context.Context, *CreateAppointmentRequest) (*CreateAppointmentResponse, error)
	// Lấy danh sách lịch hẹn theo khách hàng hoặc nhân viên
	GetAppointmentsByCustomer(context.Context, *GetAppointmentsByCustomerRequest) (*GetAppointmentsResponse, error)
	GetAppointmentsByEmployee(context.Context, *GetAppointmentsByEmployeeRequest) (*GetAppointmentsResponse, error)
	GetAppointmentsByBranch(context.Context, *GetAppointmentsByBranchRequest) (*GetAppointmentsResponse, error)
	GetAllAppointments(context.Context, *GetAllAppointmentsRequest) (*GetAllAppointmentsResponse, error)
	// Cập nhật trạng thái lịch hẹn
	UpdateAppointmentStatus(context.Context, *UpdateAppointmentStatusRequest) (*UpdateAppointmentStatusResponse, error)
	UpdateEmployeeForAppointment(context.Context, *UpdateEmployeeForAppointmentRequest) (*UpdateEmployeeForAppointmentResponse, error)
	// Lấy chi tiết lịch hẹn
	GetAppointmentDetails(context.Context, *GetAppointmentDetailsRequest) (*GetAppointmentDetailsResponse, error)
	// Quản lý danh sách dịch vụ
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	mustEmbedUnimplementedAppointmentServiceServer()
}

// UnimplementedAppointmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppointmentServiceServer struct{}

func (UnimplementedAppointmentServiceServer) CreateAppointment(context.Context, *CreateAppointmentRequest) (*CreateAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppointment not implemented")
}
func (UnimplementedAppointmentServiceServer) GetAppointmentsByCustomer(context.Context, *GetAppointmentsByCustomerRequest) (*GetAppointmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppointmentsByCustomer not implemented")
}
func (UnimplementedAppointmentServiceServer) GetAppointmentsByEmployee(context.Context, *GetAppointmentsByEmployeeRequest) (*GetAppointmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppointmentsByEmployee not implemented")
}
func (UnimplementedAppointmentServiceServer) GetAppointmentsByBranch(context.Context, *GetAppointmentsByBranchRequest) (*GetAppointmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppointmentsByBranch not implemented")
}
func (UnimplementedAppointmentServiceServer) GetAllAppointments(context.Context, *GetAllAppointmentsRequest) (*GetAllAppointmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAppointments not implemented")
}
func (UnimplementedAppointmentServiceServer) UpdateAppointmentStatus(context.Context, *UpdateAppointmentStatusRequest) (*UpdateAppointmentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppointmentStatus not implemented")
}
func (UnimplementedAppointmentServiceServer) UpdateEmployeeForAppointment(context.Context, *UpdateEmployeeForAppointmentRequest) (*UpdateEmployeeForAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployeeForAppointment not implemented")
}
func (UnimplementedAppointmentServiceServer) GetAppointmentDetails(context.Context, *GetAppointmentDetailsRequest) (*GetAppointmentDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppointmentDetails not implemented")
}
func (UnimplementedAppointmentServiceServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedAppointmentServiceServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedAppointmentServiceServer) UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedAppointmentServiceServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedAppointmentServiceServer) mustEmbedUnimplementedAppointmentServiceServer() {}
func (UnimplementedAppointmentServiceServer) testEmbeddedByValue()                            {}

// UnsafeAppointmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServiceServer will
// result in compilation errors.
type UnsafeAppointmentServiceServer interface {
	mustEmbedUnimplementedAppointmentServiceServer()
}

func RegisterAppointmentServiceServer(s grpc.ServiceRegistrar, srv AppointmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedAppointmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppointmentService_ServiceDesc, srv)
}

func _AppointmentService_CreateAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).CreateAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_CreateAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).CreateAppointment(ctx, req.(*CreateAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetAppointmentsByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppointmentsByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetAppointmentsByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetAppointmentsByCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetAppointmentsByCustomer(ctx, req.(*GetAppointmentsByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetAppointmentsByEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppointmentsByEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetAppointmentsByEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetAppointmentsByEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetAppointmentsByEmployee(ctx, req.(*GetAppointmentsByEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetAppointmentsByBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppointmentsByBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetAppointmentsByBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetAppointmentsByBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetAppointmentsByBranch(ctx, req.(*GetAppointmentsByBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetAllAppointments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAppointmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetAllAppointments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetAllAppointments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetAllAppointments(ctx, req.(*GetAllAppointmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_UpdateAppointmentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppointmentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).UpdateAppointmentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_UpdateAppointmentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).UpdateAppointmentStatus(ctx, req.(*UpdateAppointmentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_UpdateEmployeeForAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeForAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).UpdateEmployeeForAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_UpdateEmployeeForAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).UpdateEmployeeForAppointment(ctx, req.(*UpdateEmployeeForAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetAppointmentDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppointmentDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetAppointmentDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetAppointmentDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetAppointmentDetails(ctx, req.(*GetAppointmentDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_GetServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_UpdateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppointmentService_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppointmentService_ServiceDesc is the grpc.ServiceDesc for AppointmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppointmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appointments.AppointmentService",
	HandlerType: (*AppointmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppointment",
			Handler:    _AppointmentService_CreateAppointment_Handler,
		},
		{
			MethodName: "GetAppointmentsByCustomer",
			Handler:    _AppointmentService_GetAppointmentsByCustomer_Handler,
		},
		{
			MethodName: "GetAppointmentsByEmployee",
			Handler:    _AppointmentService_GetAppointmentsByEmployee_Handler,
		},
		{
			MethodName: "GetAppointmentsByBranch",
			Handler:    _AppointmentService_GetAppointmentsByBranch_Handler,
		},
		{
			MethodName: "GetAllAppointments",
			Handler:    _AppointmentService_GetAllAppointments_Handler,
		},
		{
			MethodName: "UpdateAppointmentStatus",
			Handler:    _AppointmentService_UpdateAppointmentStatus_Handler,
		},
		{
			MethodName: "UpdateEmployeeForAppointment",
			Handler:    _AppointmentService_UpdateEmployeeForAppointment_Handler,
		},
		{
			MethodName: "GetAppointmentDetails",
			Handler:    _AppointmentService_GetAppointmentDetails_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _AppointmentService_CreateService_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _AppointmentService_GetServices_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _AppointmentService_UpdateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _AppointmentService_DeleteService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appointments.proto",
}
