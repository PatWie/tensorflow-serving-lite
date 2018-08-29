// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/attr_value.proto

package framework // import "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"

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

// Protocol buffer representing the value for an attr used to configure an Op.
// Comment indicates the corresponding attr type.  Only the field matching the
// attr type may be filled.
type AttrValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttrValue_S
	//	*AttrValue_I
	//	*AttrValue_F
	//	*AttrValue_B
	//	*AttrValue_Type
	//	*AttrValue_Shape
	//	*AttrValue_Tensor
	//	*AttrValue_List
	//	*AttrValue_Func
	//	*AttrValue_Placeholder
	Value                isAttrValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttrValue) Reset()         { *m = AttrValue{} }
func (m *AttrValue) String() string { return proto.CompactTextString(m) }
func (*AttrValue) ProtoMessage()    {}
func (*AttrValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_attr_value_bfba2df0a0411d02, []int{0}
}
func (m *AttrValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrValue.Unmarshal(m, b)
}
func (m *AttrValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrValue.Marshal(b, m, deterministic)
}
func (dst *AttrValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrValue.Merge(dst, src)
}
func (m *AttrValue) XXX_Size() int {
	return xxx_messageInfo_AttrValue.Size(m)
}
func (m *AttrValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttrValue proto.InternalMessageInfo

type isAttrValue_Value interface {
	isAttrValue_Value()
}

type AttrValue_S struct {
	S []byte `protobuf:"bytes,2,opt,name=s,proto3,oneof"`
}

type AttrValue_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,proto3,oneof"`
}

type AttrValue_F struct {
	F float32 `protobuf:"fixed32,4,opt,name=f,proto3,oneof"`
}

type AttrValue_B struct {
	B bool `protobuf:"varint,5,opt,name=b,proto3,oneof"`
}

type AttrValue_Type struct {
	Type DataType `protobuf:"varint,6,opt,name=type,proto3,enum=tensorflow.DataType,oneof"`
}

type AttrValue_Shape struct {
	Shape *TensorShapeProto `protobuf:"bytes,7,opt,name=shape,proto3,oneof"`
}

type AttrValue_Tensor struct {
	Tensor *TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"`
}

type AttrValue_List struct {
	List *AttrValue_ListValue `protobuf:"bytes,1,opt,name=list,proto3,oneof"`
}

type AttrValue_Func struct {
	Func *NameAttrList `protobuf:"bytes,10,opt,name=func,proto3,oneof"`
}

type AttrValue_Placeholder struct {
	Placeholder string `protobuf:"bytes,9,opt,name=placeholder,proto3,oneof"`
}

func (*AttrValue_S) isAttrValue_Value() {}

func (*AttrValue_I) isAttrValue_Value() {}

func (*AttrValue_F) isAttrValue_Value() {}

func (*AttrValue_B) isAttrValue_Value() {}

func (*AttrValue_Type) isAttrValue_Value() {}

func (*AttrValue_Shape) isAttrValue_Value() {}

func (*AttrValue_Tensor) isAttrValue_Value() {}

func (*AttrValue_List) isAttrValue_Value() {}

func (*AttrValue_Func) isAttrValue_Value() {}

func (*AttrValue_Placeholder) isAttrValue_Value() {}

func (m *AttrValue) GetValue() isAttrValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttrValue) GetS() []byte {
	if x, ok := m.GetValue().(*AttrValue_S); ok {
		return x.S
	}
	return nil
}

func (m *AttrValue) GetI() int64 {
	if x, ok := m.GetValue().(*AttrValue_I); ok {
		return x.I
	}
	return 0
}

func (m *AttrValue) GetF() float32 {
	if x, ok := m.GetValue().(*AttrValue_F); ok {
		return x.F
	}
	return 0
}

func (m *AttrValue) GetB() bool {
	if x, ok := m.GetValue().(*AttrValue_B); ok {
		return x.B
	}
	return false
}

func (m *AttrValue) GetType() DataType {
	if x, ok := m.GetValue().(*AttrValue_Type); ok {
		return x.Type
	}
	return DataType_DT_INVALID
}

func (m *AttrValue) GetShape() *TensorShapeProto {
	if x, ok := m.GetValue().(*AttrValue_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *AttrValue) GetTensor() *TensorProto {
	if x, ok := m.GetValue().(*AttrValue_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (m *AttrValue) GetList() *AttrValue_ListValue {
	if x, ok := m.GetValue().(*AttrValue_List); ok {
		return x.List
	}
	return nil
}

func (m *AttrValue) GetFunc() *NameAttrList {
	if x, ok := m.GetValue().(*AttrValue_Func); ok {
		return x.Func
	}
	return nil
}

func (m *AttrValue) GetPlaceholder() string {
	if x, ok := m.GetValue().(*AttrValue_Placeholder); ok {
		return x.Placeholder
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AttrValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AttrValue_OneofMarshaler, _AttrValue_OneofUnmarshaler, _AttrValue_OneofSizer, []interface{}{
		(*AttrValue_S)(nil),
		(*AttrValue_I)(nil),
		(*AttrValue_F)(nil),
		(*AttrValue_B)(nil),
		(*AttrValue_Type)(nil),
		(*AttrValue_Shape)(nil),
		(*AttrValue_Tensor)(nil),
		(*AttrValue_List)(nil),
		(*AttrValue_Func)(nil),
		(*AttrValue_Placeholder)(nil),
	}
}

func _AttrValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AttrValue)
	// value
	switch x := m.Value.(type) {
	case *AttrValue_S:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.S)
	case *AttrValue_I:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.I))
	case *AttrValue_F:
		b.EncodeVarint(4<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.F)))
	case *AttrValue_B:
		t := uint64(0)
		if x.B {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *AttrValue_Type:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Type))
	case *AttrValue_Shape:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Shape); err != nil {
			return err
		}
	case *AttrValue_Tensor:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tensor); err != nil {
			return err
		}
	case *AttrValue_List:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case *AttrValue_Func:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Func); err != nil {
			return err
		}
	case *AttrValue_Placeholder:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Placeholder)
	case nil:
	default:
		return fmt.Errorf("AttrValue.Value has unexpected type %T", x)
	}
	return nil
}

func _AttrValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AttrValue)
	switch tag {
	case 2: // value.s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &AttrValue_S{x}
		return true, err
	case 3: // value.i
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_I{int64(x)}
		return true, err
	case 4: // value.f
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &AttrValue_F{math.Float32frombits(uint32(x))}
		return true, err
	case 5: // value.b
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_B{x != 0}
		return true, err
	case 6: // value.type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_Type{DataType(x)}
		return true, err
	case 7: // value.shape
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TensorShapeProto)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Shape{msg}
		return true, err
	case 8: // value.tensor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TensorProto)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Tensor{msg}
		return true, err
	case 1: // value.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AttrValue_ListValue)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_List{msg}
		return true, err
	case 10: // value.func
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NameAttrList)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Func{msg}
		return true, err
	case 9: // value.placeholder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &AttrValue_Placeholder{x}
		return true, err
	default:
		return false, nil
	}
}

func _AttrValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AttrValue)
	// value
	switch x := m.Value.(type) {
	case *AttrValue_S:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.S)))
		n += len(x.S)
	case *AttrValue_I:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.I))
	case *AttrValue_F:
		n += 1 // tag and wire
		n += 4
	case *AttrValue_B:
		n += 1 // tag and wire
		n += 1
	case *AttrValue_Type:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Type))
	case *AttrValue_Shape:
		s := proto.Size(x.Shape)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Tensor:
		s := proto.Size(x.Tensor)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_List:
		s := proto.Size(x.List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Func:
		s := proto.Size(x.Func)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Placeholder:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Placeholder)))
		n += len(x.Placeholder)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// LINT.IfChange
type AttrValue_ListValue struct {
	S                    [][]byte            `protobuf:"bytes,2,rep,name=s,proto3" json:"s,omitempty"`
	I                    []int64             `protobuf:"varint,3,rep,packed,name=i,proto3" json:"i,omitempty"`
	F                    []float32           `protobuf:"fixed32,4,rep,packed,name=f,proto3" json:"f,omitempty"`
	B                    []bool              `protobuf:"varint,5,rep,packed,name=b,proto3" json:"b,omitempty"`
	Type                 []DataType          `protobuf:"varint,6,rep,packed,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	Shape                []*TensorShapeProto `protobuf:"bytes,7,rep,name=shape,proto3" json:"shape,omitempty"`
	Tensor               []*TensorProto      `protobuf:"bytes,8,rep,name=tensor,proto3" json:"tensor,omitempty"`
	Func                 []*NameAttrList     `protobuf:"bytes,9,rep,name=func,proto3" json:"func,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AttrValue_ListValue) Reset()         { *m = AttrValue_ListValue{} }
func (m *AttrValue_ListValue) String() string { return proto.CompactTextString(m) }
func (*AttrValue_ListValue) ProtoMessage()    {}
func (*AttrValue_ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_attr_value_bfba2df0a0411d02, []int{0, 0}
}
func (m *AttrValue_ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrValue_ListValue.Unmarshal(m, b)
}
func (m *AttrValue_ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrValue_ListValue.Marshal(b, m, deterministic)
}
func (dst *AttrValue_ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrValue_ListValue.Merge(dst, src)
}
func (m *AttrValue_ListValue) XXX_Size() int {
	return xxx_messageInfo_AttrValue_ListValue.Size(m)
}
func (m *AttrValue_ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrValue_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttrValue_ListValue proto.InternalMessageInfo

func (m *AttrValue_ListValue) GetS() [][]byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *AttrValue_ListValue) GetI() []int64 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *AttrValue_ListValue) GetF() []float32 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *AttrValue_ListValue) GetB() []bool {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *AttrValue_ListValue) GetType() []DataType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AttrValue_ListValue) GetShape() []*TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *AttrValue_ListValue) GetTensor() []*TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *AttrValue_ListValue) GetFunc() []*NameAttrList {
	if m != nil {
		return m.Func
	}
	return nil
}

// A list of attr names and their values. The whole list is attached
// with a string name.  E.g., MatMul[T=float].
type NameAttrList struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Attr                 map[string]*AttrValue `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NameAttrList) Reset()         { *m = NameAttrList{} }
func (m *NameAttrList) String() string { return proto.CompactTextString(m) }
func (*NameAttrList) ProtoMessage()    {}
func (*NameAttrList) Descriptor() ([]byte, []int) {
	return fileDescriptor_attr_value_bfba2df0a0411d02, []int{1}
}
func (m *NameAttrList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameAttrList.Unmarshal(m, b)
}
func (m *NameAttrList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameAttrList.Marshal(b, m, deterministic)
}
func (dst *NameAttrList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameAttrList.Merge(dst, src)
}
func (m *NameAttrList) XXX_Size() int {
	return xxx_messageInfo_NameAttrList.Size(m)
}
func (m *NameAttrList) XXX_DiscardUnknown() {
	xxx_messageInfo_NameAttrList.DiscardUnknown(m)
}

var xxx_messageInfo_NameAttrList proto.InternalMessageInfo

func (m *NameAttrList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameAttrList) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*AttrValue)(nil), "tensorflow.AttrValue")
	proto.RegisterType((*AttrValue_ListValue)(nil), "tensorflow.AttrValue.ListValue")
	proto.RegisterType((*NameAttrList)(nil), "tensorflow.NameAttrList")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.NameAttrList.AttrEntry")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/attr_value.proto", fileDescriptor_attr_value_bfba2df0a0411d02)
}

var fileDescriptor_attr_value_bfba2df0a0411d02 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0x99, 0xb4, 0xdb, 0x4c, 0xcb, 0x5a, 0x06, 0xc5, 0xa1, 0x08, 0x86, 0x82, 0x32,
	0xac, 0x25, 0xc1, 0xf8, 0x07, 0x11, 0xbc, 0xb0, 0x28, 0xf4, 0x42, 0x96, 0x25, 0x2e, 0x5e, 0x78,
	0xb3, 0x24, 0x75, 0xd2, 0x86, 0x4d, 0x3b, 0x61, 0x32, 0x75, 0xe9, 0x13, 0x78, 0xeb, 0x73, 0xf8,
	0x84, 0x5e, 0xca, 0x99, 0xc9, 0xa6, 0xc1, 0xdd, 0xee, 0xde, 0x9d, 0x73, 0xe6, 0xfb, 0x26, 0x27,
	0xbf, 0x39, 0x87, 0x9c, 0x68, 0xb1, 0xa9, 0xa4, 0xca, 0x0a, 0x79, 0x15, 0x2e, 0xa4, 0x12, 0x61,
	0xa6, 0x92, 0xb5, 0xb8, 0x92, 0xea, 0x32, 0x4c, 0xb4, 0x56, 0x17, 0x3f, 0x93, 0x62, 0x2b, 0x82,
	0x52, 0x49, 0x2d, 0x29, 0xd9, 0x6b, 0xc7, 0xcf, 0x0f, 0xfb, 0xec, 0x89, 0xf5, 0x8c, 0xa7, 0xf7,
	0xe9, 0x2e, 0xaa, 0x55, 0x52, 0xd6, 0x5f, 0x18, 0x3f, 0xbb, 0x43, 0xbd, 0x2b, 0x45, 0x65, 0x65,
	0x93, 0x5f, 0x5d, 0xe2, 0x7d, 0xd4, 0x5a, 0x7d, 0x83, 0xe6, 0xe8, 0x31, 0x41, 0x15, 0x73, 0x7c,
	0xc4, 0x87, 0xf3, 0x4e, 0x8c, 0x2a, 0xc8, 0x73, 0x86, 0x7d, 0xc4, 0x31, 0xe4, 0x39, 0xe4, 0x19,
	0x73, 0x7d, 0xc4, 0x1d, 0xc8, 0x33, 0xc8, 0x53, 0xd6, 0xf5, 0x11, 0xef, 0x43, 0x9e, 0xd2, 0x13,
	0xe2, 0xc2, 0xe5, 0xac, 0xe7, 0x23, 0x7e, 0x1c, 0x3d, 0x0c, 0xf6, 0x3d, 0x04, 0x9f, 0x12, 0x9d,
	0x9c, 0xef, 0x4a, 0x31, 0xef, 0xc4, 0x46, 0x43, 0x5f, 0x93, 0xae, 0xe9, 0x97, 0x1d, 0xf9, 0x88,
	0x0f, 0xa2, 0x27, 0x6d, 0xf1, 0xb9, 0x09, 0xbf, 0xc2, 0xf1, 0x19, 0xb4, 0x39, 0xef, 0xc4, 0x56,
	0x4c, 0x5f, 0x92, 0x9e, 0xd5, 0xb1, 0xbe, 0xb1, 0x3d, 0xbe, 0x69, 0xbb, 0x76, 0xd4, 0x42, 0xfa,
	0x86, 0xb8, 0x45, 0x5e, 0x69, 0x86, 0x8c, 0xe1, 0x69, 0xdb, 0xd0, 0xfc, 0x79, 0xf0, 0x25, 0xaf,
	0xb4, 0x89, 0xa0, 0x3f, 0x90, 0xd3, 0x80, 0xb8, 0xd9, 0x76, 0xb3, 0x60, 0xc4, 0xd8, 0x58, 0xdb,
	0x76, 0x9a, 0xac, 0x05, 0x58, 0xc1, 0x04, 0x7a, 0xd0, 0xd1, 0x09, 0x19, 0x94, 0x45, 0xb2, 0x10,
	0x2b, 0x59, 0xfc, 0x10, 0x8a, 0x79, 0x3e, 0xe2, 0xde, 0xbc, 0x13, 0xb7, 0x8b, 0xe3, 0xdf, 0x0e,
	0xf1, 0x9a, 0x2f, 0xd1, 0xa1, 0xa5, 0x8d, 0xf9, 0x10, 0x58, 0x8f, 0x2c, 0x6b, 0xcc, 0xf1, 0xcc,
	0x19, 0x21, 0xa0, 0x3d, 0xb2, 0xb4, 0x31, 0x77, 0x6c, 0x25, 0x83, 0x0a, 0xf0, 0xc6, 0xbc, 0x6f,
	0x2b, 0x29, 0x9d, 0x36, 0xc4, 0xf1, 0x21, 0xe2, 0x46, 0x6a, 0x99, 0x47, 0x7b, 0xe6, 0xf8, 0x3e,
	0xe6, 0xd7, 0xc4, 0xc3, 0x16, 0x71, 0x7c, 0x07, 0xf1, 0x86, 0xf7, 0xb4, 0x06, 0xe7, 0x19, 0xf9,
	0x41, 0x70, 0x16, 0xdb, 0xec, 0x88, 0x74, 0xcd, 0x62, 0x4c, 0xfe, 0x20, 0x32, 0x6c, 0x9f, 0x53,
	0x4a, 0xdc, 0x4d, 0xb2, 0x16, 0xe6, 0xdd, 0xbc, 0xd8, 0xc4, 0xf4, 0x2d, 0x71, 0x61, 0x97, 0x0c,
	0xb5, 0x41, 0x34, 0x39, 0x74, 0xb7, 0x79, 0xd8, 0xcf, 0x1b, 0xad, 0x76, 0xb1, 0xd1, 0x8f, 0x4f,
	0xed, 0x94, 0x9b, 0x12, 0x1d, 0x11, 0x7c, 0x29, 0x76, 0xf5, 0xbd, 0x10, 0xd2, 0x17, 0x75, 0x13,
	0x66, 0xf6, 0x07, 0xd1, 0xa3, 0x5b, 0x67, 0x24, 0xb6, 0x9a, 0xf7, 0xce, 0x3b, 0x34, 0x93, 0x84,
	0x49, 0xb5, 0x6c, 0xcb, 0x9a, 0xf5, 0x9a, 0x3d, 0x68, 0x1c, 0x86, 0x4b, 0x75, 0x86, 0xbe, 0x7f,
	0x58, 0xe6, 0x7a, 0xb5, 0x4d, 0x83, 0x85, 0x5c, 0x87, 0xad, 0xbd, 0xbc, 0x3d, 0x5c, 0xca, 0xff,
	0x16, 0xf6, 0x2f, 0x42, 0x69, 0xcf, 0xac, 0xeb, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab,
	0xbc, 0xb8, 0x90, 0x65, 0x04, 0x00, 0x00,
}
