// Code generated by protoc-gen-go.
// source: back.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	back.proto
	core.proto
	test.proto

It has these top-level messages:
	BackendLogin
	UpGetBackendLoginInfo
	DownGetBackendLoginInfo
	UpPing
	DownPing
	WebUpRegister
	WebDownRegister
	WebUpGetServerInfo
	WebDownGetServerInfo
	ServerInfo
	UpGetConfigToCenter
	SyncConfigDataToAll
	LoginKickPlayerToAgent
	PlayerLoginToUnique
	LoginCheckPlayerToAgent
	PlayerLoginToGame
	DownPlayerLogin
	WebUpRegisterPassword
	WebDownRegisterPassword
	TestStruct
	TestToAgent
	TestToGame
	UpJoinRoomToUnique
	DownJoinRoomToClient
	NoticeJoinRoomToClient
	UpSendMessageToGame
	DownSendMessageToClient
	NoticeSendMessageToClient
	JoinRoomServerIDToAgent
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// BackendLogin BackendLogin
type BackendLogin struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName" json:"UserName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *BackendLogin) Reset()                    { *m = BackendLogin{} }
func (m *BackendLogin) String() string            { return proto.CompactTextString(m) }
func (*BackendLogin) ProtoMessage()               {}
func (*BackendLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// UpGetBackendLoginInfo UpGetBackendLoginInfo
type UpGetBackendLoginInfo struct {
}

func (m *UpGetBackendLoginInfo) Reset()                    { *m = UpGetBackendLoginInfo{} }
func (m *UpGetBackendLoginInfo) String() string            { return proto.CompactTextString(m) }
func (*UpGetBackendLoginInfo) ProtoMessage()               {}
func (*UpGetBackendLoginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// DownGetBackendLoginInfo DownGetBackendLoginInfo
type DownGetBackendLoginInfo struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName" json:"UserName,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *DownGetBackendLoginInfo) Reset()                    { *m = DownGetBackendLoginInfo{} }
func (m *DownGetBackendLoginInfo) String() string            { return proto.CompactTextString(m) }
func (*DownGetBackendLoginInfo) ProtoMessage()               {}
func (*DownGetBackendLoginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*BackendLogin)(nil), "pb.BackendLogin")
	proto.RegisterType((*UpGetBackendLoginInfo)(nil), "pb.UpGetBackendLoginInfo")
	proto.RegisterType((*DownGetBackendLoginInfo)(nil), "pb.DownGetBackendLoginInfo")
}

func init() { proto.RegisterFile("back.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x4c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe3, 0xe2, 0x71, 0x02,
	0x8a, 0xa4, 0xe6, 0xa5, 0xf8, 0xe4, 0xa7, 0x67, 0xe6, 0x09, 0x49, 0x71, 0x71, 0x84, 0x16, 0xa7,
	0x16, 0xf9, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x20, 0xb9,
	0x80, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x26, 0x88, 0x1c, 0x8c, 0xaf, 0x24, 0xce,
	0x25, 0x1a, 0x5a, 0xe0, 0x9e, 0x5a, 0x82, 0x6c, 0x98, 0x67, 0x5e, 0x5a, 0xbe, 0x92, 0x27, 0x97,
	0xb8, 0x4b, 0x7e, 0x79, 0x1e, 0x16, 0x29, 0xbc, 0x76, 0x09, 0x71, 0xb1, 0x80, 0xc5, 0x21, 0xf6,
	0x80, 0xd9, 0x49, 0x6c, 0x60, 0x67, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x86, 0x45,
	0xca, 0xc4, 0x00, 0x00, 0x00,
}
