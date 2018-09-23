package driver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/nats-rpc/nrpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The tentative calculation request message
type TentativeRequest struct {
	Raw                  []byte   `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Wht                  []byte   `protobuf:"bytes,2,opt,name=wht,proto3" json:"wht,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TentativeRequest) Reset()         { *m = TentativeRequest{} }
func (m *TentativeRequest) String() string { return proto.CompactTextString(m) }
func (*TentativeRequest) ProtoMessage()    {}
func (*TentativeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{0}
}

func (m *TentativeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TentativeRequest.Unmarshal(m, b)
}
func (m *TentativeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TentativeRequest.Marshal(b, m, deterministic)
}
func (m *TentativeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TentativeRequest.Merge(m, src)
}
func (m *TentativeRequest) XXX_Size() int {
	return xxx_messageInfo_TentativeRequest.Size(m)
}
func (m *TentativeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TentativeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TentativeRequest proto.InternalMessageInfo

func (m *TentativeRequest) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *TentativeRequest) GetWht() []byte {
	if m != nil {
		return m.Wht
	}
	return nil
}

// The tentative calculation response message
type TentativeResponse struct {
	Msg                  []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TentativeResponse) Reset()         { *m = TentativeResponse{} }
func (m *TentativeResponse) String() string { return proto.CompactTextString(m) }
func (*TentativeResponse) ProtoMessage()    {}
func (*TentativeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{1}
}

func (m *TentativeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TentativeResponse.Unmarshal(m, b)
}
func (m *TentativeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TentativeResponse.Marshal(b, m, deterministic)
}
func (m *TentativeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TentativeResponse.Merge(m, src)
}
func (m *TentativeResponse) XXX_Size() int {
	return xxx_messageInfo_TentativeResponse.Size(m)
}
func (m *TentativeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TentativeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TentativeResponse proto.InternalMessageInfo

func (m *TentativeResponse) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// The incentive delivery request message
type IncentiveRequest struct {
	Raw                  []byte   `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Wht                  []byte   `protobuf:"bytes,2,opt,name=wht,proto3" json:"wht,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncentiveRequest) Reset()         { *m = IncentiveRequest{} }
func (m *IncentiveRequest) String() string { return proto.CompactTextString(m) }
func (*IncentiveRequest) ProtoMessage()    {}
func (*IncentiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{2}
}

func (m *IncentiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncentiveRequest.Unmarshal(m, b)
}
func (m *IncentiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncentiveRequest.Marshal(b, m, deterministic)
}
func (m *IncentiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveRequest.Merge(m, src)
}
func (m *IncentiveRequest) XXX_Size() int {
	return xxx_messageInfo_IncentiveRequest.Size(m)
}
func (m *IncentiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveRequest proto.InternalMessageInfo

func (m *IncentiveRequest) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *IncentiveRequest) GetWht() []byte {
	if m != nil {
		return m.Wht
	}
	return nil
}

// The incentive delivery response message
type IncentiveResponse struct {
	Msg                  []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncentiveResponse) Reset()         { *m = IncentiveResponse{} }
func (m *IncentiveResponse) String() string { return proto.CompactTextString(m) }
func (*IncentiveResponse) ProtoMessage()    {}
func (*IncentiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{3}
}

func (m *IncentiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncentiveResponse.Unmarshal(m, b)
}
func (m *IncentiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncentiveResponse.Marshal(b, m, deterministic)
}
func (m *IncentiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveResponse.Merge(m, src)
}
func (m *IncentiveResponse) XXX_Size() int {
	return xxx_messageInfo_IncentiveResponse.Size(m)
}
func (m *IncentiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveResponse proto.InternalMessageInfo

func (m *IncentiveResponse) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*TentativeRequest)(nil), "driver.TentativeRequest")
	proto.RegisterType((*TentativeResponse)(nil), "driver.TentativeResponse")
	proto.RegisterType((*IncentiveRequest)(nil), "driver.IncentiveRequest")
	proto.RegisterType((*IncentiveResponse)(nil), "driver.IncentiveResponse")
}

func init() { proto.RegisterFile("driver.proto", fileDescriptor_521003751d596b5e) }

var fileDescriptor_521003751d596b5e = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x29, 0xca, 0x2c,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0xb8, 0xf2, 0x8a,
	0x0a, 0x92, 0x21, 0x62, 0x4a, 0x66, 0x5c, 0x02, 0x21, 0xa9, 0x79, 0x25, 0x89, 0x25, 0x99, 0x65,
	0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x45, 0x89, 0xe5, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x20, 0x26, 0x48, 0xa4, 0x3c, 0xa3, 0x44, 0x82, 0x09, 0x22,
	0x52, 0x9e, 0x51, 0xa2, 0xa4, 0xca, 0x25, 0x88, 0xa4, 0xaf, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15,
	0xa4, 0x2c, 0xb7, 0x38, 0x1d, 0xa6, 0x31, 0xb7, 0x38, 0x1d, 0x64, 0xbc, 0x67, 0x5e, 0x72, 0x6a,
	0x1e, 0x19, 0xc6, 0x23, 0xe9, 0xc3, 0x65, 0xbc, 0xd1, 0x04, 0x46, 0x2e, 0x36, 0x17, 0xb0, 0xa7,
	0x84, 0x9c, 0xb8, 0x38, 0xe1, 0x0e, 0x12, 0x92, 0xd0, 0x83, 0x7a, 0x1c, 0xdd, 0x6f, 0x52, 0x92,
	0x58, 0x64, 0x20, 0xc6, 0x2b, 0x31, 0x80, 0xcc, 0x80, 0xdb, 0x8a, 0x30, 0x03, 0xdd, 0x03, 0x08,
	0x33, 0x30, 0x9c, 0xa8, 0xc4, 0xe0, 0x24, 0xcd, 0xc5, 0x9f, 0x9a, 0xa8, 0x57, 0x92, 0x9f, 0x9d,
	0x9a, 0xa7, 0x97, 0x52, 0x54, 0x96, 0x93, 0x5f, 0xe2, 0x04, 0x75, 0x62, 0x00, 0x63, 0x12, 0x1b,
	0x38, 0xd0, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x25, 0x29, 0x1e, 0x98, 0x01, 0x00,
	0x00,
}
