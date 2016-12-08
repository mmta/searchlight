// Code generated by protoc-gen-go.
// source: ssh.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	ssh.proto

It has these top-level messages:
	SSHGetRequest
	SSHGetResponse
	SSHKey
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Use specific requests for protos
type SSHGetRequest struct {
	Namespace    string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	ClusterName  string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	InstanceName string `protobuf:"bytes,3,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *SSHGetRequest) Reset()                    { *m = SSHGetRequest{} }
func (m *SSHGetRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHGetRequest) ProtoMessage()               {}
func (*SSHGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SSHGetRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SSHGetRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *SSHGetRequest) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

// return phid ?
type SSHGetResponse struct {
	Status       *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SshKey       *SSHKey                 `protobuf:"bytes,2,opt,name=ssh_key,json=sshKey" json:"ssh_key,omitempty"`
	InstanceAddr string                  `protobuf:"bytes,3,opt,name=instance_addr,json=instanceAddr" json:"instance_addr,omitempty"`
	InstancePort int32                   `protobuf:"varint,4,opt,name=instance_port,json=instancePort" json:"instance_port,omitempty"`
	User         string                  `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Command      string                  `protobuf:"bytes,6,opt,name=command" json:"command,omitempty"`
}

func (m *SSHGetResponse) Reset()                    { *m = SSHGetResponse{} }
func (m *SSHGetResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHGetResponse) ProtoMessage()               {}
func (*SSHGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SSHGetResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SSHGetResponse) GetSshKey() *SSHKey {
	if m != nil {
		return m.SshKey
	}
	return nil
}

func (m *SSHGetResponse) GetInstanceAddr() string {
	if m != nil {
		return m.InstanceAddr
	}
	return ""
}

func (m *SSHGetResponse) GetInstancePort() int32 {
	if m != nil {
		return m.InstancePort
	}
	return 0
}

func (m *SSHGetResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SSHGetResponse) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type SSHKey struct {
	PublicKey          []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey         []byte `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	AwsFingerprint     string `protobuf:"bytes,3,opt,name=aws_fingerprint,json=awsFingerprint" json:"aws_fingerprint,omitempty"`
	OpensshFingerprint string `protobuf:"bytes,4,opt,name=openssh_fingerprint,json=opensshFingerprint" json:"openssh_fingerprint,omitempty"`
}

func (m *SSHKey) Reset()                    { *m = SSHKey{} }
func (m *SSHKey) String() string            { return proto.CompactTextString(m) }
func (*SSHKey) ProtoMessage()               {}
func (*SSHKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SSHKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SSHKey) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *SSHKey) GetAwsFingerprint() string {
	if m != nil {
		return m.AwsFingerprint
	}
	return ""
}

func (m *SSHKey) GetOpensshFingerprint() string {
	if m != nil {
		return m.OpensshFingerprint
	}
	return ""
}

func init() {
	proto.RegisterType((*SSHGetRequest)(nil), "appscode.ssh.v1beta1.SSHGetRequest")
	proto.RegisterType((*SSHGetResponse)(nil), "appscode.ssh.v1beta1.SSHGetResponse")
	proto.RegisterType((*SSHKey)(nil), "appscode.ssh.v1beta1.SSHKey")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SSH service

type SSHClient interface {
	Get(ctx context.Context, in *SSHGetRequest, opts ...grpc.CallOption) (*SSHGetResponse, error)
}

type sSHClient struct {
	cc *grpc.ClientConn
}

func NewSSHClient(cc *grpc.ClientConn) SSHClient {
	return &sSHClient{cc}
}

func (c *sSHClient) Get(ctx context.Context, in *SSHGetRequest, opts ...grpc.CallOption) (*SSHGetResponse, error) {
	out := new(SSHGetResponse)
	err := grpc.Invoke(ctx, "/appscode.ssh.v1beta1.SSH/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SSH service

type SSHServer interface {
	Get(context.Context, *SSHGetRequest) (*SSHGetResponse, error)
}

func RegisterSSHServer(s *grpc.Server, srv SSHServer) {
	s.RegisterService(&_SSH_serviceDesc, srv)
}

func _SSH_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSHServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ssh.v1beta1.SSH/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSHServer).Get(ctx, req.(*SSHGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSH_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ssh.v1beta1.SSH",
	HandlerType: (*SSHServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SSH_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ssh.proto",
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x26, 0x75, 0xc8, 0x24, 0x2d, 0xd2, 0x82, 0x84, 0x15, 0xa5, 0x4a, 0x49, 0x41,
	0xf4, 0xc9, 0x56, 0x53, 0x71, 0x00, 0xfa, 0x40, 0x23, 0x55, 0x42, 0x91, 0xfd, 0xc6, 0x4b, 0xb4,
	0xb1, 0x87, 0xc4, 0x22, 0xde, 0x5d, 0x76, 0xd6, 0x8d, 0xf2, 0xc2, 0x03, 0x57, 0xe0, 0x06, 0x5c,
	0x89, 0x2b, 0x70, 0x00, 0x8e, 0x80, 0x76, 0xed, 0xd4, 0x09, 0x82, 0xbe, 0x58, 0xeb, 0x7f, 0xbe,
	0x99, 0xfd, 0x67, 0x76, 0xa0, 0x4b, 0xb4, 0x0a, 0x95, 0x96, 0x46, 0xb2, 0xe7, 0x5c, 0x29, 0x4a,
	0x65, 0x86, 0xa1, 0xd5, 0xee, 0xaf, 0x16, 0x68, 0xf8, 0xd5, 0x60, 0xb8, 0x94, 0x72, 0xb9, 0xc6,
	0x88, 0xab, 0x3c, 0xe2, 0x42, 0x48, 0xc3, 0x4d, 0x2e, 0x05, 0x55, 0x39, 0x83, 0xd1, 0x2e, 0xc7,
	0xc5, 0x33, 0xb3, 0x55, 0x48, 0x91, 0xfb, 0x56, 0xc0, 0xb8, 0x84, 0x93, 0x24, 0x99, 0xde, 0xa2,
	0x89, 0xf1, 0x4b, 0x89, 0x64, 0xd8, 0x10, 0xba, 0x82, 0x17, 0x48, 0x8a, 0xa7, 0x18, 0x78, 0xe7,
	0xde, 0x65, 0x37, 0x6e, 0x04, 0xf6, 0x12, 0xfa, 0xe9, 0xba, 0x24, 0x83, 0x7a, 0x6e, 0xc5, 0xe0,
	0xc8, 0x01, 0xbd, 0x5a, 0xfb, 0xc0, 0x0b, 0x64, 0x17, 0x70, 0x92, 0x0b, 0x32, 0x5c, 0xa4, 0x58,
	0x31, 0x2d, 0xc7, 0xf4, 0x77, 0xa2, 0x85, 0xc6, 0xbf, 0x3d, 0x38, 0xdd, 0xdd, 0x4b, 0x4a, 0x0a,
	0x42, 0x16, 0x81, 0x4f, 0x86, 0x9b, 0x92, 0xdc, 0xad, 0xbd, 0xc9, 0x8b, 0xf0, 0xa1, 0xdf, 0xca,
	0x77, 0x98, 0xb8, 0x70, 0x5c, 0x63, 0xec, 0x2d, 0x74, 0x88, 0x56, 0xf3, 0xcf, 0xb8, 0x75, 0x36,
	0x7a, 0x93, 0x61, 0xf8, 0xaf, 0x09, 0x85, 0x49, 0x32, 0xbd, 0xc3, 0x6d, 0xec, 0x13, 0xad, 0xee,
	0x70, 0x7b, 0xe0, 0x8f, 0x67, 0x99, 0xfe, 0xdb, 0xdf, 0xbb, 0x2c, 0xd3, 0x07, 0x90, 0x92, 0xda,
	0x04, 0xed, 0x73, 0xef, 0xf2, 0xb8, 0x81, 0x66, 0x52, 0x1b, 0xc6, 0xa0, 0x5d, 0x12, 0xea, 0xe0,
	0xd8, 0x15, 0x70, 0x67, 0x16, 0x40, 0x27, 0x95, 0x45, 0xc1, 0x45, 0x16, 0xf8, 0x4e, 0xde, 0xfd,
	0x8e, 0x7f, 0x78, 0xe0, 0x57, 0x56, 0xd8, 0x19, 0x80, 0x2a, 0x17, 0xeb, 0x3c, 0x75, 0xe6, 0x6d,
	0xbb, 0xfd, 0xb8, 0x5b, 0x29, 0x36, 0x3c, 0x82, 0x9e, 0xd2, 0xf9, 0x3d, 0x37, 0xf8, 0xd0, 0x5c,
	0x3f, 0x86, 0x5a, 0xb2, 0xc0, 0x1b, 0x78, 0xca, 0x37, 0x34, 0xff, 0x94, 0x8b, 0x25, 0x6a, 0xa5,
	0x73, 0x61, 0xea, 0x26, 0x4e, 0xf9, 0x86, 0xde, 0x37, 0x2a, 0x8b, 0xe0, 0x99, 0x54, 0x28, 0xec,
	0x98, 0xf6, 0xe1, 0xb6, 0x83, 0x59, 0x1d, 0xda, 0x4b, 0x98, 0x7c, 0x85, 0x56, 0x92, 0x4c, 0xd9,
	0x06, 0x5a, 0xb7, 0x68, 0xd8, 0xc5, 0x7f, 0x07, 0xda, 0x2c, 0xcc, 0xe0, 0xd5, 0xe3, 0x50, 0xf5,
	0xba, 0xe3, 0xd7, 0xdf, 0x7e, 0xfe, 0xfa, 0x7e, 0x34, 0x62, 0x67, 0xd1, 0xc1, 0x46, 0x12, 0xad,
	0xa2, 0x3a, 0xc3, 0x9e, 0x6f, 0xae, 0x61, 0x98, 0xca, 0xa2, 0xa9, 0xc8, 0x55, 0xbe, 0x5f, 0xf5,
	0xe6, 0x49, 0x92, 0x4c, 0x67, 0x76, 0x71, 0x67, 0xde, 0xc7, 0x4e, 0x2d, 0x2e, 0x7c, 0xb7, 0xca,
	0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x72, 0x02, 0x25, 0x2c, 0x03, 0x00, 0x00,
}