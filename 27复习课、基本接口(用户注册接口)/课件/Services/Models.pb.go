// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package Services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserModel struct {
	UserId               int32                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string               `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPwd              string               `protobuf:"bytes,3,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	UserDate             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=user_date,json=userDate,proto3" json:"user_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserModel) Reset()         { *m = UserModel{} }
func (m *UserModel) String() string { return proto.CompactTextString(m) }
func (*UserModel) ProtoMessage()    {}
func (*UserModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *UserModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserModel.Unmarshal(m, b)
}
func (m *UserModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserModel.Marshal(b, m, deterministic)
}
func (m *UserModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModel.Merge(m, src)
}
func (m *UserModel) XXX_Size() int {
	return xxx_messageInfo_UserModel.Size(m)
}
func (m *UserModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserModel proto.InternalMessageInfo

func (m *UserModel) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserModel) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserModel) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

func (m *UserModel) GetUserDate() *timestamp.Timestamp {
	if m != nil {
		return m.UserDate
	}
	return nil
}

func init() {
	proto.RegisterType((*UserModel)(nil), "Services.UserModel")
}

func init() { proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a) }

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf1, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x08, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x2d, 0x96, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x27, 0x95, 0xa6,
	0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x94, 0x2a, 0x4d, 0x65, 0xe4,
	0xe2, 0x0c, 0x2d, 0x4e, 0x2d, 0x02, 0xeb, 0x17, 0x12, 0xe7, 0x62, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a,
	0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0xa4,
	0xb9, 0x38, 0xc1, 0x12, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c,
	0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x21, 0x49, 0x2e, 0x30, 0x3b, 0xbe, 0xa0, 0x3c, 0x45, 0x82,
	0x19, 0x2c, 0x07, 0x36, 0x25, 0xa0, 0x3c, 0x45, 0xc8, 0x1c, 0xaa, 0x2f, 0x25, 0xb1, 0x24, 0x55,
	0x82, 0x45, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4a, 0x0f, 0xe2, 0x26, 0x3d, 0x98, 0x9b, 0xf4, 0x42,
	0x60, 0x6e, 0x82, 0x98, 0xe9, 0x92, 0x58, 0x92, 0x9a, 0xc4, 0x06, 0x96, 0x35, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0x74, 0xbc, 0x94, 0xd9, 0x00, 0x00, 0x00,
}
