// protoc --go_out=proto/artist/. --go-grpc_out=proto/artist/.
// proto/artist/artist.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/artist/artist.proto

package artistService

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
	ArtistService_FindByID_FullMethodName = "/artistService.ArtistService/FindByID"
)

// ArtistServiceClient is the client API for ArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistServiceClient interface {
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error)
}

type artistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistServiceClient(cc grpc.ClientConnInterface) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindByIDResponse)
	err := c.cc.Invoke(ctx, ArtistService_FindByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistServiceServer is the server API for ArtistService service.
// All implementations must embed UnimplementedArtistServiceServer
// for forward compatibility.
type ArtistServiceServer interface {
	FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error)
	mustEmbedUnimplementedArtistServiceServer()
}

// UnimplementedArtistServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedArtistServiceServer struct{}

func (UnimplementedArtistServiceServer) FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedArtistServiceServer) mustEmbedUnimplementedArtistServiceServer() {}
func (UnimplementedArtistServiceServer) testEmbeddedByValue()                       {}

// UnsafeArtistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtistServiceServer will
// result in compilation errors.
type UnsafeArtistServiceServer interface {
	mustEmbedUnimplementedArtistServiceServer()
}

func RegisterArtistServiceServer(s grpc.ServiceRegistrar, srv ArtistServiceServer) {
	// If the following call pancis, it indicates UnimplementedArtistServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ArtistService_ServiceDesc, srv)
}

func _ArtistService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_FindByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtistService_ServiceDesc is the grpc.ServiceDesc for ArtistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "artistService.ArtistService",
	HandlerType: (*ArtistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByID",
			Handler:    _ArtistService_FindByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/artist/artist.proto",
}
