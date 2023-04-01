// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pres.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Pres struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pres) Reset()         { *m = Pres{} }
func (m *Pres) String() string { return proto.CompactTextString(m) }
func (*Pres) ProtoMessage()    {}
func (*Pres) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b290bd8880a60f, []int{0}
}

func (m *Pres) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pres.Unmarshal(m, b)
}
func (m *Pres) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pres.Marshal(b, m, deterministic)
}
func (m *Pres) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pres.Merge(m, src)
}
func (m *Pres) XXX_Size() int {
	return xxx_messageInfo_Pres.Size(m)
}
func (m *Pres) XXX_DiscardUnknown() {
	xxx_messageInfo_Pres.DiscardUnknown(m)
}

var xxx_messageInfo_Pres proto.InternalMessageInfo

func (m *Pres) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pres) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Slide struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Idx                  uint32   `protobuf:"varint,2,opt,name=idx,proto3" json:"idx,omitempty"`
	ImageWidth           uint32   `protobuf:"varint,3,opt,name=image_width,json=imageWidth,proto3" json:"image_width,omitempty"`
	ImageHeight          uint32   `protobuf:"varint,4,opt,name=image_height,json=imageHeight,proto3" json:"image_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slide) Reset()         { *m = Slide{} }
func (m *Slide) String() string { return proto.CompactTextString(m) }
func (*Slide) ProtoMessage()    {}
func (*Slide) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b290bd8880a60f, []int{1}
}

func (m *Slide) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slide.Unmarshal(m, b)
}
func (m *Slide) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slide.Marshal(b, m, deterministic)
}
func (m *Slide) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slide.Merge(m, src)
}
func (m *Slide) XXX_Size() int {
	return xxx_messageInfo_Slide.Size(m)
}
func (m *Slide) XXX_DiscardUnknown() {
	xxx_messageInfo_Slide.DiscardUnknown(m)
}

var xxx_messageInfo_Slide proto.InternalMessageInfo

func (m *Slide) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Slide) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *Slide) GetImageWidth() uint32 {
	if m != nil {
		return m.ImageWidth
	}
	return 0
}

func (m *Slide) GetImageHeight() uint32 {
	if m != nil {
		return m.ImageHeight
	}
	return 0
}

type Slides struct {
	Num                  uint32   `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Slide                []*Slide `protobuf:"bytes,2,rep,name=slide,proto3" json:"slide,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slides) Reset()         { *m = Slides{} }
func (m *Slides) String() string { return proto.CompactTextString(m) }
func (*Slides) ProtoMessage()    {}
func (*Slides) Descriptor() ([]byte, []int) {
	return fileDescriptor_58b290bd8880a60f, []int{2}
}

func (m *Slides) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slides.Unmarshal(m, b)
}
func (m *Slides) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slides.Marshal(b, m, deterministic)
}
func (m *Slides) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slides.Merge(m, src)
}
func (m *Slides) XXX_Size() int {
	return xxx_messageInfo_Slides.Size(m)
}
func (m *Slides) XXX_DiscardUnknown() {
	xxx_messageInfo_Slides.DiscardUnknown(m)
}

var xxx_messageInfo_Slides proto.InternalMessageInfo

func (m *Slides) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Slides) GetSlide() []*Slide {
	if m != nil {
		return m.Slide
	}
	return nil
}

func init() {
	proto.RegisterType((*Pres)(nil), "pres.Pres")
	proto.RegisterType((*Slide)(nil), "pres.Slide")
	proto.RegisterType((*Slides)(nil), "pres.Slides")
}

func init() {
	proto.RegisterFile("pres.proto", fileDescriptor_58b290bd8880a60f)
}

var fileDescriptor_58b290bd8880a60f = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xd8, 0x2d, 0xe2, 0xd2, 0x22, 0x74, 0x93, 0xc5, 0x42, 0x9a, 0x29, 0x42, 0x28,
	0x43, 0x99, 0x59, 0x98, 0x18, 0x2b, 0x77, 0x40, 0x62, 0x41, 0x41, 0xb6, 0x92, 0x93, 0x9a, 0x34,
	0xd8, 0x46, 0xf0, 0xf3, 0xd1, 0x9d, 0x17, 0x86, 0x6e, 0x9f, 0xde, 0xfb, 0xce, 0xa7, 0x33, 0xc0,
	0x12, 0x7c, 0xec, 0x96, 0x70, 0x4e, 0x67, 0xd4, 0xcc, 0xcd, 0x03, 0xe8, 0x43, 0xf0, 0x11, 0x11,
	0xf4, 0xdc, 0x4f, 0xde, 0x14, 0x75, 0xd1, 0x5e, 0x5b, 0x61, 0xbc, 0x81, 0x92, 0x9c, 0x29, 0xeb,
	0xa2, 0xd5, 0xb6, 0x24, 0xd7, 0x7c, 0xc1, 0xea, 0x78, 0x22, 0xe7, 0x2f, 0xca, 0xb7, 0xa0, 0xc8,
	0xfd, 0x8a, 0xbd, 0xb5, 0x8c, 0x78, 0x0f, 0x15, 0x4d, 0xfd, 0xe0, 0x3f, 0x7e, 0xc8, 0xa5, 0xd1,
	0x28, 0x69, 0x40, 0xa2, 0x37, 0x4e, 0x70, 0x07, 0x9b, 0x2c, 0x8c, 0x9e, 0x86, 0x31, 0x19, 0x2d,
	0x46, 0x1e, 0x7a, 0x95, 0xa8, 0x79, 0x86, 0xb5, 0xac, 0x8c, 0xfc, 0xfe, 0xfc, 0x3d, 0xc9, 0xca,
	0xad, 0x65, 0xc4, 0x1d, 0xac, 0x22, 0x77, 0xa6, 0xac, 0x55, 0x5b, 0xed, 0xab, 0x4e, 0x8e, 0x13,
	0xdd, 0xe6, 0x66, 0xff, 0x08, 0x57, 0x87, 0x3e, 0x44, 0x9a, 0x07, 0xb6, 0x8f, 0xcb, 0x89, 0x12,
	0x42, 0xf6, 0xf8, 0xea, 0xbb, 0xcd, 0xbf, 0x99, 0xf8, 0xa2, 0xde, 0x8b, 0xee, 0x73, 0x2d, 0xbf,
	0xf3, 0xf4, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x89, 0x1f, 0xf4, 0x2b, 0x01, 0x00, 0x00,
}
