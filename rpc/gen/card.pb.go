// Code generated by protoc-gen-go.
// source: card.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	card.proto
	gacha_request.proto
	gacha_response.proto
	gacha_service.proto

It has these top-level messages:
	Card
	Parameter
	GachaRequest
	GachaResponse
*/
package rpc

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

type Rarity int32

const (
	Rarity_COMMON    Rarity = 0
	Rarity_RARE      Rarity = 1
	Rarity_EPIC      Rarity = 2
	Rarity_LEGENDARY Rarity = 3
)

var Rarity_name = map[int32]string{
	0: "COMMON",
	1: "RARE",
	2: "EPIC",
	3: "LEGENDARY",
}
var Rarity_value = map[string]int32{
	"COMMON":    0,
	"RARE":      1,
	"EPIC":      2,
	"LEGENDARY": 3,
}

func (x Rarity) Enum() *Rarity {
	p := new(Rarity)
	*p = x
	return p
}
func (x Rarity) String() string {
	return proto.EnumName(Rarity_name, int32(x))
}
func (x *Rarity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Rarity_value, data, "Rarity")
	if err != nil {
		return err
	}
	*x = Rarity(value)
	return nil
}
func (Rarity) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Type int32

const (
	Type_TROOP    Type = 1
	Type_SPELL    Type = 2
	Type_BUILDING Type = 3
)

var Type_name = map[int32]string{
	1: "TROOP",
	2: "SPELL",
	3: "BUILDING",
}
var Type_value = map[string]int32{
	"TROOP":    1,
	"SPELL":    2,
	"BUILDING": 3,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Card struct {
	CardId           *int32     `protobuf:"varint,1,req,name=card_id,json=cardId" json:"card_id,omitempty"`
	Name             *string    `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Type             *Type      `protobuf:"varint,3,req,name=type,enum=rpc.Type" json:"type,omitempty"`
	Rarity           *Rarity    `protobuf:"varint,4,req,name=rarity,enum=rpc.Rarity" json:"rarity,omitempty"`
	Cost             *int32     `protobuf:"varint,5,req,name=cost" json:"cost,omitempty"`
	Parameter        *Parameter `protobuf:"bytes,6,opt,name=parameter" json:"parameter,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Card) Reset()                    { *m = Card{} }
func (m *Card) String() string            { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()               {}
func (*Card) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Card) GetCardId() int32 {
	if m != nil && m.CardId != nil {
		return *m.CardId
	}
	return 0
}

func (m *Card) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Card) GetType() Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Type_TROOP
}

func (m *Card) GetRarity() Rarity {
	if m != nil && m.Rarity != nil {
		return *m.Rarity
	}
	return Rarity_COMMON
}

func (m *Card) GetCost() int32 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return 0
}

func (m *Card) GetParameter() *Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

type Parameter struct {
	Hp               *int32   `protobuf:"varint,1,opt,name=hp" json:"hp,omitempty"`
	Damage           *int32   `protobuf:"varint,2,opt,name=damage" json:"damage,omitempty"`
	HitSpeed         *float32 `protobuf:"fixed32,3,opt,name=hit_speed,json=hitSpeed" json:"hit_speed,omitempty"`
	Range            *int32   `protobuf:"varint,4,opt,name=range" json:"range,omitempty"`
	Radius           *float32 `protobuf:"fixed32,5,opt,name=radius" json:"radius,omitempty"`
	Lifetime         *int32   `protobuf:"varint,6,opt,name=lifetime" json:"lifetime,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Parameter) GetHp() int32 {
	if m != nil && m.Hp != nil {
		return *m.Hp
	}
	return 0
}

func (m *Parameter) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *Parameter) GetHitSpeed() float32 {
	if m != nil && m.HitSpeed != nil {
		return *m.HitSpeed
	}
	return 0
}

func (m *Parameter) GetRange() int32 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

func (m *Parameter) GetRadius() float32 {
	if m != nil && m.Radius != nil {
		return *m.Radius
	}
	return 0
}

func (m *Parameter) GetLifetime() int32 {
	if m != nil && m.Lifetime != nil {
		return *m.Lifetime
	}
	return 0
}

func init() {
	proto.RegisterType((*Card)(nil), "rpc.Card")
	proto.RegisterType((*Parameter)(nil), "rpc.Parameter")
	proto.RegisterEnum("rpc.Rarity", Rarity_name, Rarity_value)
	proto.RegisterEnum("rpc.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("card.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0x4f, 0xab, 0xda, 0x40,
	0x14, 0xc5, 0x3b, 0x93, 0x3f, 0x4d, 0xae, 0xad, 0x84, 0x4b, 0x69, 0x87, 0x96, 0x42, 0xb0, 0x9b,
	0x20, 0xc5, 0x85, 0x9b, 0xae, 0xad, 0x06, 0x09, 0x44, 0x13, 0x46, 0xbb, 0xe8, 0x4a, 0x86, 0x64,
	0x6a, 0x02, 0x8d, 0x0e, 0xe3, 0x74, 0xe1, 0x57, 0xe9, 0x37, 0x79, 0xdf, 0xee, 0x31, 0xa3, 0xcf,
	0xb7, 0x3b, 0xbf, 0x73, 0xce, 0xbd, 0xb9, 0x19, 0x80, 0x46, 0xe8, 0x76, 0xa6, 0xf4, 0xd9, 0x9c,
	0xd1, 0xd3, 0xaa, 0x99, 0x3c, 0x11, 0xf0, 0x97, 0x42, 0xb7, 0xf8, 0x09, 0xde, 0xda, 0xec, 0xd0,
	0xb7, 0x8c, 0xa4, 0x34, 0x0b, 0x78, 0x68, 0xb1, 0x68, 0x11, 0xc1, 0x3f, 0x89, 0x41, 0x32, 0x9a,
	0xd2, 0x2c, 0xe6, 0x4e, 0xe3, 0x57, 0xf0, 0xcd, 0x55, 0x49, 0xe6, 0xa5, 0x34, 0x1b, 0xcf, 0xe3,
	0x99, 0x56, 0xcd, 0x6c, 0x7f, 0x55, 0x92, 0x3b, 0x1b, 0xbf, 0x41, 0xa8, 0x85, 0xee, 0xcd, 0x95,
	0xf9, 0xae, 0x30, 0x72, 0x05, 0xee, 0x2c, 0x7e, 0x8f, 0xec, 0xde, 0xe6, 0x7c, 0x31, 0x2c, 0x70,
	0x5f, 0x73, 0x1a, 0xbf, 0x43, 0xac, 0x84, 0x16, 0x83, 0x34, 0x52, 0xb3, 0x30, 0x25, 0xd9, 0x68,
	0x3e, 0x76, 0xb3, 0xf5, 0x8b, 0xcb, 0x5f, 0x0b, 0x93, 0xff, 0x04, 0xe2, 0x47, 0x80, 0x63, 0xa0,
	0x9d, 0x62, 0x24, 0x25, 0x59, 0xc0, 0x69, 0xa7, 0xf0, 0x23, 0x84, 0xad, 0x18, 0xc4, 0xd1, 0x5e,
	0x6e, 0xbd, 0x3b, 0xe1, 0x17, 0x88, 0xbb, 0xde, 0x1c, 0x2e, 0x4a, 0xca, 0x96, 0x79, 0x29, 0xc9,
	0x28, 0x8f, 0xba, 0xde, 0xec, 0x2c, 0xe3, 0x07, 0x08, 0xb4, 0x38, 0x1d, 0x25, 0xf3, 0xdd, 0xcc,
	0x0d, 0xec, 0x2a, 0x2d, 0xda, 0xfe, 0xdf, 0x85, 0x05, 0xae, 0x7f, 0x27, 0xfc, 0x0c, 0xd1, 0xdf,
	0xfe, 0x8f, 0x34, 0xfd, 0x20, 0xdd, 0xb5, 0x01, 0x7f, 0xf0, 0xf4, 0x07, 0x84, 0xb7, 0x1f, 0x46,
	0x80, 0x70, 0x59, 0x6d, 0x36, 0xd5, 0x36, 0x79, 0x83, 0x11, 0xf8, 0x7c, 0xc1, 0xf3, 0x84, 0x58,
	0x95, 0xd7, 0xc5, 0x32, 0xa1, 0xf8, 0x1e, 0xe2, 0x32, 0x5f, 0xe7, 0xdb, 0xd5, 0x82, 0xff, 0x4e,
	0xbc, 0xe9, 0x14, 0x7c, 0xfb, 0x94, 0x18, 0x43, 0xb0, 0xe7, 0x55, 0x55, 0x27, 0xc4, 0xca, 0x5d,
	0x9d, 0x97, 0x65, 0x42, 0xf1, 0x1d, 0x44, 0x3f, 0x7f, 0x15, 0xe5, 0xaa, 0xd8, 0xae, 0x13, 0xef,
	0x39, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xd5, 0x6d, 0xf3, 0xcf, 0x01, 0x00, 0x00,
}