// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xfury/bet/odds_type.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OddsType is the representation of the type of the odds.
type OddsType int32

const (
	// invalid odds type
	OddsType_ODDS_TYPE_UNSPECIFIED OddsType = 0
	// decimal odds type (european)
	OddsType_ODDS_TYPE_DECIMAL OddsType = 1
	// fractional odds type (british)
	OddsType_ODDS_TYPE_FRACTIONAL OddsType = 2
	// moneyline odds type (american)
	OddsType_ODDS_TYPE_MONEYLINE OddsType = 3
)

var OddsType_name = map[int32]string{
	0: "ODDS_TYPE_UNSPECIFIED",
	1: "ODDS_TYPE_DECIMAL",
	2: "ODDS_TYPE_FRACTIONAL",
	3: "ODDS_TYPE_MONEYLINE",
}

var OddsType_value = map[string]int32{
	"ODDS_TYPE_UNSPECIFIED": 0,
	"ODDS_TYPE_DECIMAL":     1,
	"ODDS_TYPE_FRACTIONAL":  2,
	"ODDS_TYPE_MONEYLINE":   3,
}

func (x OddsType) String() string {
	return proto.EnumName(OddsType_name, int32(x))
}

func (OddsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1281e1b8ca89efb8, []int{0}
}

func init() {
	proto.RegisterEnum("furysports.xfury.bet.OddsType", OddsType_name, OddsType_value)
}

func init() { proto.RegisterFile("xfury/bet/odds_type.proto", fileDescriptor_1281e1b8ca89efb8) }

var fileDescriptor_1281e1b8ca89efb8 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2b, 0x2d, 0xaa,
	0xd4, 0x4f, 0x4a, 0x2d, 0xd1, 0xcf, 0x4f, 0x49, 0x29, 0x8e, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06, 0xc9, 0x14, 0x17, 0xe4, 0x17, 0x95, 0x14, 0xeb, 0x81,
	0x98, 0x7a, 0x49, 0xa9, 0x25, 0x5a, 0xf9, 0x5c, 0x1c, 0xfe, 0x29, 0x29, 0xc5, 0x21, 0x95, 0x05,
	0xa9, 0x42, 0x92, 0x5c, 0xa2, 0xfe, 0x2e, 0x2e, 0xc1, 0xf1, 0x21, 0x91, 0x01, 0xae, 0xf1, 0xa1,
	0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c, 0x42, 0xa2, 0x5c, 0x82,
	0x08, 0x29, 0x17, 0x57, 0x67, 0x4f, 0x5f, 0x47, 0x1f, 0x01, 0x46, 0x21, 0x09, 0x2e, 0x11, 0x84,
	0xb0, 0x5b, 0x90, 0xa3, 0x73, 0x88, 0xa7, 0xbf, 0x9f, 0xa3, 0x8f, 0x00, 0x93, 0x90, 0x38, 0x97,
	0x30, 0x42, 0xc6, 0xd7, 0xdf, 0xcf, 0x35, 0xd2, 0xc7, 0xd3, 0xcf, 0x55, 0x80, 0xd9, 0xc9, 0xe5,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd2, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0x12, 0xf3, 0x40, 0x4e, 0xd4, 0x85, 0x38, 0x57, 0x1f,
	0xec, 0xa7, 0x0a, 0xb0, 0xaf, 0x40, 0x1e, 0x2a, 0x4e, 0x62, 0x03, 0x7b, 0xc9, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x2c, 0xdb, 0x9a, 0xee, 0x00, 0x00, 0x00,
}
