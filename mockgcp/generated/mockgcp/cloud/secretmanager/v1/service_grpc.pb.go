// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/secretmanager/v1/service.proto

package secretmanagerpb

import (
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecretManagerServiceClient is the client API for SecretManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretManagerServiceClient interface {
	// Lists [Secrets][mockgcp.cloud.secretmanager.v1.Secret].
	ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error)
	// Creates a new [Secret][mockgcp.cloud.secretmanager.v1.Secret] containing no [SecretVersions][mockgcp.cloud.secretmanager.v1.SecretVersion].
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	// Creates a new [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] containing secret data and attaches
	// it to an existing [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	AddSecretVersion(ctx context.Context, in *AddSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error)
	// Gets metadata for a given [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	// Updates metadata of an existing [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	UpdateSecret(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	// Deletes a [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists [SecretVersions][mockgcp.cloud.secretmanager.v1.SecretVersion]. This call does not return secret
	// data.
	ListSecretVersions(ctx context.Context, in *ListSecretVersionsRequest, opts ...grpc.CallOption) (*ListSecretVersionsResponse, error)
	// Gets metadata for a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// `projects/*/secrets/*/versions/latest` is an alias to the most recently
	// created [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	GetSecretVersion(ctx context.Context, in *GetSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error)
	// Accesses a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion]. This call returns the secret data.
	//
	// `projects/*/secrets/*/versions/latest` is an alias to the most recently
	// created [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	AccessSecretVersion(ctx context.Context, in *AccessSecretVersionRequest, opts ...grpc.CallOption) (*AccessSecretVersionResponse, error)
	// Disables a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// Sets the [state][mockgcp.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] to
	// [DISABLED][mockgcp.cloud.secretmanager.v1.SecretVersion.State.DISABLED].
	DisableSecretVersion(ctx context.Context, in *DisableSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error)
	// Enables a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// Sets the [state][mockgcp.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] to
	// [ENABLED][mockgcp.cloud.secretmanager.v1.SecretVersion.State.ENABLED].
	EnableSecretVersion(ctx context.Context, in *EnableSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error)
	// Destroys a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// Sets the [state][mockgcp.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] to
	// [DESTROYED][mockgcp.cloud.secretmanager.v1.SecretVersion.State.DESTROYED] and irrevocably destroys the
	// secret data.
	DestroySecretVersion(ctx context.Context, in *DestroySecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error)
	// Sets the access control policy on the specified secret. Replaces any
	// existing policy.
	//
	// Permissions on [SecretVersions][mockgcp.cloud.secretmanager.v1.SecretVersion] are enforced according
	// to the policy set on the associated [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	SetIamPolicy(ctx context.Context, in *iampb.SetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error)
	// Gets the access control policy for a secret.
	// Returns empty policy if the secret exists and does not have a policy set.
	GetIamPolicy(ctx context.Context, in *iampb.GetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error)
	// Returns permissions that a caller has for the specified secret.
	// If the secret does not exist, this call returns an empty set of
	// permissions, not a NOT_FOUND error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(ctx context.Context, in *iampb.TestIamPermissionsRequest, opts ...grpc.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

type secretManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretManagerServiceClient(cc grpc.ClientConnInterface) SecretManagerServiceClient {
	return &secretManagerServiceClient{cc}
}

func (c *secretManagerServiceClient) ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error) {
	out := new(ListSecretsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) AddSecretVersion(ctx context.Context, in *AddSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error) {
	out := new(SecretVersion)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/AddSecretVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) UpdateSecret(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/UpdateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) ListSecretVersions(ctx context.Context, in *ListSecretVersionsRequest, opts ...grpc.CallOption) (*ListSecretVersionsResponse, error) {
	out := new(ListSecretVersionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/ListSecretVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) GetSecretVersion(ctx context.Context, in *GetSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error) {
	out := new(SecretVersion)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/GetSecretVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) AccessSecretVersion(ctx context.Context, in *AccessSecretVersionRequest, opts ...grpc.CallOption) (*AccessSecretVersionResponse, error) {
	out := new(AccessSecretVersionResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/AccessSecretVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) DisableSecretVersion(ctx context.Context, in *DisableSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error) {
	out := new(SecretVersion)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/DisableSecretVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) EnableSecretVersion(ctx context.Context, in *EnableSecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error) {
	out := new(SecretVersion)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/EnableSecretVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) DestroySecretVersion(ctx context.Context, in *DestroySecretVersionRequest, opts ...grpc.CallOption) (*SecretVersion, error) {
	out := new(SecretVersion)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/DestroySecretVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) SetIamPolicy(ctx context.Context, in *iampb.SetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error) {
	out := new(iampb.Policy)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) GetIamPolicy(ctx context.Context, in *iampb.GetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error) {
	out := new(iampb.Policy)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretManagerServiceClient) TestIamPermissions(ctx context.Context, in *iampb.TestIamPermissionsRequest, opts ...grpc.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	out := new(iampb.TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.secretmanager.v1.SecretManagerService/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretManagerServiceServer is the server API for SecretManagerService service.
// All implementations must embed UnimplementedSecretManagerServiceServer
// for forward compatibility
type SecretManagerServiceServer interface {
	// Lists [Secrets][mockgcp.cloud.secretmanager.v1.Secret].
	ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error)
	// Creates a new [Secret][mockgcp.cloud.secretmanager.v1.Secret] containing no [SecretVersions][mockgcp.cloud.secretmanager.v1.SecretVersion].
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	// Creates a new [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] containing secret data and attaches
	// it to an existing [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	AddSecretVersion(context.Context, *AddSecretVersionRequest) (*SecretVersion, error)
	// Gets metadata for a given [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	GetSecret(context.Context, *GetSecretRequest) (*Secret, error)
	// Updates metadata of an existing [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	UpdateSecret(context.Context, *UpdateSecretRequest) (*Secret, error)
	// Deletes a [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	DeleteSecret(context.Context, *DeleteSecretRequest) (*empty.Empty, error)
	// Lists [SecretVersions][mockgcp.cloud.secretmanager.v1.SecretVersion]. This call does not return secret
	// data.
	ListSecretVersions(context.Context, *ListSecretVersionsRequest) (*ListSecretVersionsResponse, error)
	// Gets metadata for a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// `projects/*/secrets/*/versions/latest` is an alias to the most recently
	// created [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	GetSecretVersion(context.Context, *GetSecretVersionRequest) (*SecretVersion, error)
	// Accesses a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion]. This call returns the secret data.
	//
	// `projects/*/secrets/*/versions/latest` is an alias to the most recently
	// created [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	AccessSecretVersion(context.Context, *AccessSecretVersionRequest) (*AccessSecretVersionResponse, error)
	// Disables a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// Sets the [state][mockgcp.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] to
	// [DISABLED][mockgcp.cloud.secretmanager.v1.SecretVersion.State.DISABLED].
	DisableSecretVersion(context.Context, *DisableSecretVersionRequest) (*SecretVersion, error)
	// Enables a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// Sets the [state][mockgcp.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] to
	// [ENABLED][mockgcp.cloud.secretmanager.v1.SecretVersion.State.ENABLED].
	EnableSecretVersion(context.Context, *EnableSecretVersionRequest) (*SecretVersion, error)
	// Destroys a [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion].
	//
	// Sets the [state][mockgcp.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][mockgcp.cloud.secretmanager.v1.SecretVersion] to
	// [DESTROYED][mockgcp.cloud.secretmanager.v1.SecretVersion.State.DESTROYED] and irrevocably destroys the
	// secret data.
	DestroySecretVersion(context.Context, *DestroySecretVersionRequest) (*SecretVersion, error)
	// Sets the access control policy on the specified secret. Replaces any
	// existing policy.
	//
	// Permissions on [SecretVersions][mockgcp.cloud.secretmanager.v1.SecretVersion] are enforced according
	// to the policy set on the associated [Secret][mockgcp.cloud.secretmanager.v1.Secret].
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest) (*iampb.Policy, error)
	// Gets the access control policy for a secret.
	// Returns empty policy if the secret exists and does not have a policy set.
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest) (*iampb.Policy, error)
	// Returns permissions that a caller has for the specified secret.
	// If the secret does not exist, this call returns an empty set of
	// permissions, not a NOT_FOUND error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error)
	mustEmbedUnimplementedSecretManagerServiceServer()
}

// UnimplementedSecretManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretManagerServiceServer struct {
}

func (UnimplementedSecretManagerServiceServer) ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}
func (UnimplementedSecretManagerServiceServer) CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedSecretManagerServiceServer) AddSecretVersion(context.Context, *AddSecretVersionRequest) (*SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecretVersion not implemented")
}
func (UnimplementedSecretManagerServiceServer) GetSecret(context.Context, *GetSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretManagerServiceServer) UpdateSecret(context.Context, *UpdateSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSecret not implemented")
}
func (UnimplementedSecretManagerServiceServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedSecretManagerServiceServer) ListSecretVersions(context.Context, *ListSecretVersionsRequest) (*ListSecretVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecretVersions not implemented")
}
func (UnimplementedSecretManagerServiceServer) GetSecretVersion(context.Context, *GetSecretVersionRequest) (*SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretVersion not implemented")
}
func (UnimplementedSecretManagerServiceServer) AccessSecretVersion(context.Context, *AccessSecretVersionRequest) (*AccessSecretVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessSecretVersion not implemented")
}
func (UnimplementedSecretManagerServiceServer) DisableSecretVersion(context.Context, *DisableSecretVersionRequest) (*SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableSecretVersion not implemented")
}
func (UnimplementedSecretManagerServiceServer) EnableSecretVersion(context.Context, *EnableSecretVersionRequest) (*SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableSecretVersion not implemented")
}
func (UnimplementedSecretManagerServiceServer) DestroySecretVersion(context.Context, *DestroySecretVersionRequest) (*SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroySecretVersion not implemented")
}
func (UnimplementedSecretManagerServiceServer) SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (UnimplementedSecretManagerServiceServer) GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (UnimplementedSecretManagerServiceServer) TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}
func (UnimplementedSecretManagerServiceServer) mustEmbedUnimplementedSecretManagerServiceServer() {}

// UnsafeSecretManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretManagerServiceServer will
// result in compilation errors.
type UnsafeSecretManagerServiceServer interface {
	mustEmbedUnimplementedSecretManagerServiceServer()
}

func RegisterSecretManagerServiceServer(s grpc.ServiceRegistrar, srv SecretManagerServiceServer) {
	s.RegisterService(&SecretManagerService_ServiceDesc, srv)
}

func _SecretManagerService_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).ListSecrets(ctx, req.(*ListSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_AddSecretVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSecretVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).AddSecretVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/AddSecretVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).AddSecretVersion(ctx, req.(*AddSecretVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_UpdateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).UpdateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/UpdateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).UpdateSecret(ctx, req.(*UpdateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_ListSecretVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).ListSecretVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/ListSecretVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).ListSecretVersions(ctx, req.(*ListSecretVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_GetSecretVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).GetSecretVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/GetSecretVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).GetSecretVersion(ctx, req.(*GetSecretVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_AccessSecretVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessSecretVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).AccessSecretVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/AccessSecretVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).AccessSecretVersion(ctx, req.(*AccessSecretVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_DisableSecretVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableSecretVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).DisableSecretVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/DisableSecretVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).DisableSecretVersion(ctx, req.(*DisableSecretVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_EnableSecretVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableSecretVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).EnableSecretVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/EnableSecretVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).EnableSecretVersion(ctx, req.(*EnableSecretVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_DestroySecretVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroySecretVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).DestroySecretVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/DestroySecretVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).DestroySecretVersion(ctx, req.(*DestroySecretVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).SetIamPolicy(ctx, req.(*iampb.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).GetIamPolicy(ctx, req.(*iampb.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretManagerService_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretManagerServiceServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.secretmanager.v1.SecretManagerService/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretManagerServiceServer).TestIamPermissions(ctx, req.(*iampb.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretManagerService_ServiceDesc is the grpc.ServiceDesc for SecretManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.secretmanager.v1.SecretManagerService",
	HandlerType: (*SecretManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSecrets",
			Handler:    _SecretManagerService_ListSecrets_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _SecretManagerService_CreateSecret_Handler,
		},
		{
			MethodName: "AddSecretVersion",
			Handler:    _SecretManagerService_AddSecretVersion_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _SecretManagerService_GetSecret_Handler,
		},
		{
			MethodName: "UpdateSecret",
			Handler:    _SecretManagerService_UpdateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _SecretManagerService_DeleteSecret_Handler,
		},
		{
			MethodName: "ListSecretVersions",
			Handler:    _SecretManagerService_ListSecretVersions_Handler,
		},
		{
			MethodName: "GetSecretVersion",
			Handler:    _SecretManagerService_GetSecretVersion_Handler,
		},
		{
			MethodName: "AccessSecretVersion",
			Handler:    _SecretManagerService_AccessSecretVersion_Handler,
		},
		{
			MethodName: "DisableSecretVersion",
			Handler:    _SecretManagerService_DisableSecretVersion_Handler,
		},
		{
			MethodName: "EnableSecretVersion",
			Handler:    _SecretManagerService_EnableSecretVersion_Handler,
		},
		{
			MethodName: "DestroySecretVersion",
			Handler:    _SecretManagerService_DestroySecretVersion_Handler,
		},
		{
			MethodName: "SetIamPolicy",
			Handler:    _SecretManagerService_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _SecretManagerService_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _SecretManagerService_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/secretmanager/v1/service.proto",
}