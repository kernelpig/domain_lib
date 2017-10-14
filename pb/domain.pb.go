// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	domain.proto

It has these top-level messages:
	SubOrderParam
	CreateOrderRequest
	CreateOrderRespose
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

// 子订单参数，具体请参考链接
// https://help.aliyun.com/document_detail/42902.html?spm=5176.doc42892.6.615.jDsihJ
type SubOrderParam struct {
	Action           string `protobuf:"bytes,1,opt,name=Action" json:"Action,omitempty"`
	RelatedName      string `protobuf:"bytes,2,opt,name=RelatedName" json:"RelatedName,omitempty"`
	Period           int32  `protobuf:"varint,3,opt,name=Period" json:"Period,omitempty"`
	DomainTemplateID string `protobuf:"bytes,4,opt,name=DomainTemplateID" json:"DomainTemplateID,omitempty"`
}

func (m *SubOrderParam) Reset()                    { *m = SubOrderParam{} }
func (m *SubOrderParam) String() string            { return proto.CompactTextString(m) }
func (*SubOrderParam) ProtoMessage()               {}
func (*SubOrderParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubOrderParam) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SubOrderParam) GetRelatedName() string {
	if m != nil {
		return m.RelatedName
	}
	return ""
}

func (m *SubOrderParam) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *SubOrderParam) GetDomainTemplateID() string {
	if m != nil {
		return m.DomainTemplateID
	}
	return ""
}

// 创建订单接口请求参数, 具体请参考链接
// https://help.aliyun.com/document_detail/42892.html?spm=5176.doc42893.6.596.f2vn3v
type CreateOrderRequest struct {
	Action        string           `protobuf:"bytes,1,opt,name=Action" json:"Action,omitempty"`
	SubOrderParam []*SubOrderParam `protobuf:"bytes,2,rep,name=SubOrderParam" json:"SubOrderParam,omitempty"`
}

func (m *CreateOrderRequest) Reset()                    { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()               {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateOrderRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateOrderRequest) GetSubOrderParam() []*SubOrderParam {
	if m != nil {
		return m.SubOrderParam
	}
	return nil
}

// 创建订单响应参数, 具体请参考链接
// https://help.aliyun.com/document_detail/42892.html?spm=5176.doc42893.6.596.f2vn3v
type CreateOrderRespose struct {
	RequestId string `protobuf:"bytes,1,opt,name=RequestId" json:"RequestId,omitempty"`
	OrderID   string `protobuf:"bytes,2,opt,name=OrderID" json:"OrderID,omitempty"`
}

func (m *CreateOrderRespose) Reset()                    { *m = CreateOrderRespose{} }
func (m *CreateOrderRespose) String() string            { return proto.CompactTextString(m) }
func (*CreateOrderRespose) ProtoMessage()               {}
func (*CreateOrderRespose) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateOrderRespose) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *CreateOrderRespose) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func init() {
	proto.RegisterType((*SubOrderParam)(nil), "pb.SubOrderParam")
	proto.RegisterType((*CreateOrderRequest)(nil), "pb.CreateOrderRequest")
	proto.RegisterType((*CreateOrderRespose)(nil), "pb.CreateOrderRespose")
}

func init() { proto.RegisterFile("domain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x85, 0x49, 0xab, 0x95, 0xde, 0x2a, 0x68, 0x06, 0xc9, 0xe0, 0x10, 0x3a, 0x15, 0x87, 0x0e,
	0x3a, 0x38, 0x8b, 0x5d, 0x0a, 0xa2, 0x25, 0xfa, 0x02, 0x89, 0xb9, 0x43, 0xc1, 0x34, 0x31, 0x4d,
	0x1f, 0xc3, 0x77, 0x96, 0x86, 0x88, 0xb6, 0x3f, 0xff, 0x78, 0xce, 0x3d, 0x87, 0xef, 0x24, 0x70,
	0xae, 0xad, 0x91, 0xe3, 0xd4, 0x3a, 0x6f, 0x83, 0xa5, 0x99, 0x53, 0xf5, 0x37, 0x81, 0x8b, 0xb7,
	0x45, 0xbd, 0x7a, 0x8d, 0x7e, 0x90, 0x5e, 0x1a, 0x7a, 0x0d, 0xc5, 0xe3, 0x47, 0x18, 0xed, 0xc4,
	0x08, 0x27, 0x4d, 0x29, 0x92, 0xa2, 0x1c, 0x2a, 0x81, 0x9f, 0x32, 0xa0, 0x7e, 0x91, 0x06, 0x59,
	0x16, 0x8f, 0xff, 0xad, 0xb5, 0x39, 0xa0, 0x1f, 0xad, 0x66, 0x39, 0x27, 0xcd, 0xa9, 0x48, 0x8a,
	0xde, 0xc2, 0x65, 0x17, 0xb9, 0xef, 0x68, 0xdc, 0x1a, 0xef, 0x3b, 0x76, 0x12, 0xeb, 0x07, 0x7e,
	0x8d, 0x40, 0x9f, 0x3c, 0xca, 0x80, 0x71, 0x91, 0xc0, 0xaf, 0x05, 0xe7, 0x70, 0x74, 0xd3, 0xc3,
	0x6e, 0x3c, 0xcb, 0x78, 0xde, 0x54, 0x77, 0x57, 0xad, 0x53, 0xed, 0xe6, 0x20, 0xb6, 0xb9, 0xfa,
	0x79, 0x87, 0x99, 0x9d, 0x9d, 0x91, 0xde, 0x40, 0x99, 0x88, 0xbd, 0x4e, 0xa4, 0x3f, 0x83, 0x32,
	0x38, 0x8b, 0xe9, 0xbe, 0x4b, 0x8f, 0xff, 0x95, 0xaa, 0x88, 0xff, 0x79, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x68, 0xf5, 0x3c, 0x8a, 0x5f, 0x01, 0x00, 0x00,
}
