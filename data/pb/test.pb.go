// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TestToAgent TestToAgent
type TestToAgent struct {
	FieldA string `protobuf:"bytes,1,opt,name=FieldA" json:"FieldA,omitempty"`
}

func (m *TestToAgent) Reset()                    { *m = TestToAgent{} }
func (m *TestToAgent) String() string            { return proto.CompactTextString(m) }
func (*TestToAgent) ProtoMessage()               {}
func (*TestToAgent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// TestToGame TestToGame
type TestToGame struct {
	FieldA string `protobuf:"bytes,1,opt,name=FieldA" json:"FieldA,omitempty"`
	FieldB string `protobuf:"bytes,2,opt,name=FieldB" json:"FieldB,omitempty"`
}

func (m *TestToGame) Reset()                    { *m = TestToGame{} }
func (m *TestToGame) String() string            { return proto.CompactTextString(m) }
func (*TestToGame) ProtoMessage()               {}
func (*TestToGame) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// UpJoinRoomToUnique 申请加入房间
type UpJoinRoomToUnique struct {
}

func (m *UpJoinRoomToUnique) Reset()                    { *m = UpJoinRoomToUnique{} }
func (m *UpJoinRoomToUnique) String() string            { return proto.CompactTextString(m) }
func (*UpJoinRoomToUnique) ProtoMessage()               {}
func (*UpJoinRoomToUnique) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// DownJoinRoomToClient 加入客户端结果
type DownJoinRoomToClient struct {
	IsSucceed bool `protobuf:"varint,1,opt,name=IsSucceed" json:"IsSucceed,omitempty"`
}

func (m *DownJoinRoomToClient) Reset()                    { *m = DownJoinRoomToClient{} }
func (m *DownJoinRoomToClient) String() string            { return proto.CompactTextString(m) }
func (*DownJoinRoomToClient) ProtoMessage()               {}
func (*DownJoinRoomToClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// NoticeJoinRoomToClient 通知加入房间
type NoticeJoinRoomToClient struct {
	OtherUID string `protobuf:"bytes,2,opt,name=OtherUID" json:"OtherUID,omitempty"`
}

func (m *NoticeJoinRoomToClient) Reset()                    { *m = NoticeJoinRoomToClient{} }
func (m *NoticeJoinRoomToClient) String() string            { return proto.CompactTextString(m) }
func (*NoticeJoinRoomToClient) ProtoMessage()               {}
func (*NoticeJoinRoomToClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

// UpSendMessageToGame 上发房间说话
type UpSendMessageToGame struct {
}

func (m *UpSendMessageToGame) Reset()                    { *m = UpSendMessageToGame{} }
func (m *UpSendMessageToGame) String() string            { return proto.CompactTextString(m) }
func (*UpSendMessageToGame) ProtoMessage()               {}
func (*UpSendMessageToGame) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

// DownSendMessageToClient 房间说话结果
type DownSendMessageToClient struct {
	IsSucceed bool `protobuf:"varint,1,opt,name=IsSucceed" json:"IsSucceed,omitempty"`
}

func (m *DownSendMessageToClient) Reset()                    { *m = DownSendMessageToClient{} }
func (m *DownSendMessageToClient) String() string            { return proto.CompactTextString(m) }
func (*DownSendMessageToClient) ProtoMessage()               {}
func (*DownSendMessageToClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

// NoticeSendMessage 通知房间说话信息
type NoticeSendMessageToClient struct {
}

func (m *NoticeSendMessageToClient) Reset()                    { *m = NoticeSendMessageToClient{} }
func (m *NoticeSendMessageToClient) String() string            { return proto.CompactTextString(m) }
func (*NoticeSendMessageToClient) ProtoMessage()               {}
func (*NoticeSendMessageToClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

// JoinRoomServerIDToAgent 通知agent需要连接的服务器id
type JoinRoomServerIDToAgent struct {
	ServerID string `protobuf:"bytes,1,opt,name=ServerID" json:"ServerID,omitempty"`
}

func (m *JoinRoomServerIDToAgent) Reset()                    { *m = JoinRoomServerIDToAgent{} }
func (m *JoinRoomServerIDToAgent) String() string            { return proto.CompactTextString(m) }
func (*JoinRoomServerIDToAgent) ProtoMessage()               {}
func (*JoinRoomServerIDToAgent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func init() {
	proto.RegisterType((*TestToAgent)(nil), "pb.TestToAgent")
	proto.RegisterType((*TestToGame)(nil), "pb.TestToGame")
	proto.RegisterType((*UpJoinRoomToUnique)(nil), "pb.UpJoinRoomToUnique")
	proto.RegisterType((*DownJoinRoomToClient)(nil), "pb.DownJoinRoomToClient")
	proto.RegisterType((*NoticeJoinRoomToClient)(nil), "pb.NoticeJoinRoomToClient")
	proto.RegisterType((*UpSendMessageToGame)(nil), "pb.UpSendMessageToGame")
	proto.RegisterType((*DownSendMessageToClient)(nil), "pb.DownSendMessageToClient")
	proto.RegisterType((*NoticeSendMessageToClient)(nil), "pb.NoticeSendMessageToClient")
	proto.RegisterType((*JoinRoomServerIDToAgent)(nil), "pb.JoinRoomServerIDToAgent")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe5, 0xe2, 0x0e, 0x01,
	0x8a, 0x84, 0xe4, 0x3b, 0xa6, 0xa7, 0xe6, 0x95, 0x08, 0x89, 0x71, 0xb1, 0xb9, 0x65, 0xa6, 0xe6,
	0xa4, 0x38, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x4a, 0x36, 0x5c, 0x5c, 0x10,
	0x65, 0xee, 0x89, 0xb9, 0xa9, 0xb8, 0x54, 0xc1, 0xc5, 0x9d, 0x24, 0x98, 0x90, 0xc4, 0x9d, 0x94,
	0x44, 0xb8, 0x84, 0x42, 0x0b, 0xbc, 0xf2, 0x33, 0xf3, 0x82, 0xf2, 0xf3, 0x73, 0x43, 0xf2, 0x43,
	0xf3, 0x32, 0x0b, 0x4b, 0x53, 0x95, 0x4c, 0xb8, 0x44, 0x5c, 0xf2, 0xcb, 0xf3, 0x10, 0xe2, 0xce,
	0x39, 0x99, 0x20, 0x37, 0xc8, 0x70, 0x71, 0x7a, 0x16, 0x07, 0x97, 0x26, 0x27, 0xa7, 0xa6, 0xa6,
	0x80, 0x2d, 0xe0, 0x08, 0x42, 0x08, 0x00, 0x75, 0x89, 0xf9, 0xe5, 0x97, 0x64, 0x26, 0xa7, 0x62,
	0xe8, 0x93, 0xe2, 0xe2, 0xf0, 0x2f, 0xc9, 0x48, 0x2d, 0x0a, 0xf5, 0x74, 0x81, 0xda, 0x0f, 0xe7,
	0x2b, 0x89, 0x72, 0x09, 0x87, 0x16, 0x04, 0xa7, 0xe6, 0xa5, 0xf8, 0xa6, 0x16, 0x17, 0x27, 0xa6,
	0xa7, 0x42, 0x3c, 0xa2, 0x64, 0xce, 0x25, 0x0e, 0x72, 0x02, 0x8a, 0x04, 0x51, 0xae, 0x90, 0xe6,
	0x92, 0x84, 0xb8, 0x02, 0x8b, 0x56, 0x25, 0x53, 0x2e, 0x71, 0x98, 0xe3, 0x82, 0x53, 0x8b, 0xca,
	0x52, 0x8b, 0x3c, 0x5d, 0x60, 0xe1, 0x0b, 0x74, 0x23, 0x4c, 0x08, 0x1a, 0x76, 0x70, 0x7e, 0x12,
	0x1b, 0x38, 0x56, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x66, 0xd0, 0x26, 0xa3, 0x01,
	0x00, 0x00,
}
