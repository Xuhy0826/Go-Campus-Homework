// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// PatientClient is the client API for Patient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientClient interface {
	GetPatientInfo(ctx context.Context, in *GetPatientInfoRequest, opts ...grpc.CallOption) (*GetPatientInfoReply, error)
	ListPatientInfo(ctx context.Context, in *ListPatientInfoRequest, opts ...grpc.CallOption) (*ListPatientInfoReply, error)
}

type patientClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientClient(cc grpc.ClientConnInterface) PatientClient {
	return &patientClient{cc}
}

func (c *patientClient) GetPatientInfo(ctx context.Context, in *GetPatientInfoRequest, opts ...grpc.CallOption) (*GetPatientInfoReply, error) {
	out := new(GetPatientInfoReply)
	err := c.cc.Invoke(ctx, "/api.patient.v1.Patient/GetPatientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientClient) ListPatientInfo(ctx context.Context, in *ListPatientInfoRequest, opts ...grpc.CallOption) (*ListPatientInfoReply, error) {
	out := new(ListPatientInfoReply)
	err := c.cc.Invoke(ctx, "/api.patient.v1.Patient/ListPatientInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServer is the server API for Patient service.
// All implementations must embed UnimplementedPatientServer
// for forward compatibility
type PatientServer interface {
	GetPatientInfo(context.Context, *GetPatientInfoRequest) (*GetPatientInfoReply, error)
	ListPatientInfo(context.Context, *ListPatientInfoRequest) (*ListPatientInfoReply, error)
	mustEmbedUnimplementedPatientServer()
}

// UnimplementedPatientServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServer struct {
}

func (UnimplementedPatientServer) GetPatientInfo(context.Context, *GetPatientInfoRequest) (*GetPatientInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientInfo not implemented")
}
func (UnimplementedPatientServer) ListPatientInfo(context.Context, *ListPatientInfoRequest) (*ListPatientInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientInfo not implemented")
}
func (UnimplementedPatientServer) mustEmbedUnimplementedPatientServer() {}

// UnsafePatientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientServer will
// result in compilation errors.
type UnsafePatientServer interface {
	mustEmbedUnimplementedPatientServer()
}

func RegisterPatientServer(s grpc.ServiceRegistrar, srv PatientServer) {
	s.RegisterService(&Patient_ServiceDesc, srv)
}

func _Patient_GetPatientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServer).GetPatientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.patient.v1.Patient/GetPatientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServer).GetPatientInfo(ctx, req.(*GetPatientInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Patient_ListPatientInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServer).ListPatientInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.patient.v1.Patient/ListPatientInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServer).ListPatientInfo(ctx, req.(*ListPatientInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Patient_ServiceDesc is the grpc.ServiceDesc for Patient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Patient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.patient.v1.Patient",
	HandlerType: (*PatientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPatientInfo",
			Handler:    _Patient_GetPatientInfo_Handler,
		},
		{
			MethodName: "ListPatientInfo",
			Handler:    _Patient_ListPatientInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patient.proto",
}