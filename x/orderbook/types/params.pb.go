// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xfury/orderbook/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the parameters for the orderbook module.
type Params struct {
	// max_book_participations is the maximum number of participations per book.
	MaxBookParticipations uint64 `protobuf:"varint,1,opt,name=max_book_participations,json=maxBookParticipations,proto3" json:"max_book_participations,omitempty" yaml:"max_book_participations"`
	// batch_settlement_count is the batch settlement deposit counts.
	BatchSettlementCount uint64 `protobuf:"varint,2,opt,name=batch_settlement_count,json=batchSettlementCount,proto3" json:"batch_settlement_count,omitempty" yaml:"batch_settlement_count"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e65bc3b240323e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxBookParticipations() uint64 {
	if m != nil {
		return m.MaxBookParticipations
	}
	return 0
}

func (m *Params) GetBatchSettlementCount() uint64 {
	if m != nil {
		return m.BatchSettlementCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "furysports.xfury.orderbook.Params")
}

func init() { proto.RegisterFile("xfury/orderbook/params.proto", fileDescriptor_21e65bc3b240323e) }

var fileDescriptor_21e65bc3b240323e = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2b, 0x2d, 0xaa,
	0xd4, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0xd6, 0x2f, 0x48, 0x2c, 0x4a, 0xcc,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x04, 0x49, 0x16, 0x17, 0xe4, 0x17, 0x95,
	0x14, 0xeb, 0x81, 0x98, 0x7a, 0x70, 0x75, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x55, 0xfa,
	0x20, 0x16, 0x44, 0x83, 0xd2, 0x71, 0x46, 0x2e, 0xb6, 0x00, 0xb0, 0x09, 0x42, 0x51, 0x5c, 0xe2,
	0xb9, 0x89, 0x15, 0xf1, 0x20, 0xc5, 0xf1, 0x05, 0x89, 0x45, 0x25, 0x99, 0xc9, 0x99, 0x05, 0x89,
	0x25, 0x99, 0xf9, 0x79, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x4e, 0x4a, 0x9f, 0xee, 0xc9,
	0xcb, 0x55, 0x26, 0xe6, 0xe6, 0x58, 0x29, 0xe1, 0x50, 0xa8, 0x14, 0x24, 0x9a, 0x9b, 0x58, 0xe1,
	0x94, 0x9f, 0x9f, 0x1d, 0x80, 0x22, 0x2e, 0x14, 0xce, 0x25, 0x96, 0x94, 0x58, 0x92, 0x9c, 0x11,
	0x5f, 0x9c, 0x5a, 0x52, 0x92, 0x93, 0x9a, 0x9b, 0x9a, 0x57, 0x12, 0x9f, 0x9c, 0x5f, 0x9a, 0x57,
	0x22, 0xc1, 0x04, 0x36, 0x5a, 0xf1, 0xd3, 0x3d, 0x79, 0x59, 0x88, 0xd1, 0xd8, 0xd5, 0x29, 0x05,
	0x89, 0x80, 0x25, 0x82, 0xe1, 0xe2, 0xce, 0x20, 0x61, 0x2b, 0x8e, 0x19, 0x0b, 0xe4, 0x19, 0x5e,
	0x2c, 0x90, 0x67, 0x74, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xb4, 0xc4, 0x3c, 0x50,
	0xb8, 0xe8, 0x42, 0xc2, 0x48, 0x1f, 0x1c, 0x96, 0x15, 0x48, 0xa1, 0x59, 0x52, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0x0e, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x47, 0x97, 0xe1,
	0x6c, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxBookParticipations != that1.MaxBookParticipations {
		return false
	}
	if this.BatchSettlementCount != that1.BatchSettlementCount {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchSettlementCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BatchSettlementCount))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxBookParticipations != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxBookParticipations))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxBookParticipations != 0 {
		n += 1 + sovParams(uint64(m.MaxBookParticipations))
	}
	if m.BatchSettlementCount != 0 {
		n += 1 + sovParams(uint64(m.BatchSettlementCount))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBookParticipations", wireType)
			}
			m.MaxBookParticipations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBookParticipations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSettlementCount", wireType)
			}
			m.BatchSettlementCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSettlementCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
