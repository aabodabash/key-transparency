// Code generated by protoc-gen-go.
// source: github.com/gdbelvin/e2ekeys/proto/e2ekeys.proto
// DO NOT EDIT!

/*
Package google_security_e2ekeys_v1 is a generated protocol buffer package.

It is generated from these files:
	github.com/gdbelvin/e2ekeys/proto/e2ekeys.proto

It has these top-level messages:
	User
	SignedKey
	GetUserRequest
*/
package google_security_e2ekeys_v1

import proto "github.com/golang/protobuf/proto"
import google_protobuf "google/protobuf"

// discarding unused import google_api1 "google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// The User is the leaf object in the binary Merkle Tree. Its unique location
// in the tree is identified by a hash of an unpredictable verifiable function
// The value of the user leaf node is
// HMAC(nonce, (H(Meta) || H(k1) || H(k2) ... || H(kn))
type User struct {
	// keys is an ordered array of SignedKeys for this user.
	// When filtering key results, individual SignedKeys may be empty, but the
	// associated hash will remain in key_ides.
	SignedKeys []*SignedKey `protobuf:"bytes,2,rep,name=signed_keys" json:"signed_keys,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetSignedKeys() []*SignedKey {
	if m != nil {
		return m.SignedKeys
	}
	return nil
}

// SignedKey contains a Key and associated signatures, showing continuity of
// key ownership.
type SignedKey struct {
	// A key with metadata.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *SignedKey) Reset()         { *m = SignedKey{} }
func (m *SignedKey) String() string { return proto.CompactTextString(m) }
func (*SignedKey) ProtoMessage()    {}

// Get request for a user object.
type GetUserRequest struct {
	// Absence of the time field indicates a request for the current value.
	Time *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	// User identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,2,opt,name=user_id" json:"user_id,omitempty"`
	// Only return the keys belonging to this app.
	AppId string `protobuf:"bytes,3,opt,name=app_id" json:"app_id,omitempty"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}

func (m *GetUserRequest) GetTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
}

// Client API for E2EKeyProxy service

type E2EKeyProxyClient interface {
	// GetUser returns a user's keys and a proof that there is only one entry for
	// this user and that it is the same one being provided to everyone else.
	// GetUser also supports querying past values by setting the epoch field.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type e2EKeyProxyClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyProxyClient(cc *grpc.ClientConn) E2EKeyProxyClient {
	return &e2EKeyProxyClient{cc}
}

func (c *e2EKeyProxyClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyProxy service

type E2EKeyProxyServer interface {
	// GetUser returns a user's keys and a proof that there is only one entry for
	// this user and that it is the same one being provided to everyone else.
	// GetUser also supports querying past values by setting the epoch field.
	GetUser(context.Context, *GetUserRequest) (*User, error)
}

func RegisterE2EKeyProxyServer(s *grpc.Server, srv E2EKeyProxyServer) {
	s.RegisterService(&_E2EKeyProxy_serviceDesc, srv)
}

func _E2EKeyProxy_GetUser_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetUserRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _E2EKeyProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.security.e2ekeys.v1.E2EKeyProxy",
	HandlerType: (*E2EKeyProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _E2EKeyProxy_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}