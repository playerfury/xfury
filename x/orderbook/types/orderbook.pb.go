// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xfury/orderbook/orderbook.proto

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

// OrderBookStatus is the enum type for the status of the orderbook.
type OrderBookStatus int32

const (
	// invalid
	OrderBookStatus_ORDER_BOOK_STATUS_UNSPECIFIED OrderBookStatus = 0
	// active
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_ACTIVE OrderBookStatus = 1
	// resolved not settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_RESOLVED OrderBookStatus = 2
	// resolved and settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_SETTLED OrderBookStatus = 3
)

var OrderBookStatus_name = map[int32]string{
	0: "ORDER_BOOK_STATUS_UNSPECIFIED",
	1: "ORDER_BOOK_STATUS_STATUS_ACTIVE",
	2: "ORDER_BOOK_STATUS_STATUS_RESOLVED",
	3: "ORDER_BOOK_STATUS_STATUS_SETTLED",
}

var OrderBookStatus_value = map[string]int32{
	"ORDER_BOOK_STATUS_UNSPECIFIED":     0,
	"ORDER_BOOK_STATUS_STATUS_ACTIVE":   1,
	"ORDER_BOOK_STATUS_STATUS_RESOLVED": 2,
	"ORDER_BOOK_STATUS_STATUS_SETTLED":  3,
}

func (x OrderBookStatus) String() string {
	return proto.EnumName(OrderBookStatus_name, int32(x))
}

func (OrderBookStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3aebc657d8cdea59, []int{0}
}

// OrderBook represents the order book maintained against a sport event.
type OrderBook struct {
	// id corresponding to the book
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// participation_count is the count of participations in the order book
	ParticipationCount uint64 `protobuf:"varint,2,opt,name=participation_count,json=participationCount,proto3" json:"participation_count,omitempty" yaml:"participation_count"`
	// odds_count is the count of the odds in the order book
	OddsCount uint64 `protobuf:"varint,3,opt,name=odds_count,json=oddsCount,proto3" json:"odds_count,omitempty" yaml:"odds_count"`
	// order book status
	Status OrderBookStatus `protobuf:"varint,4,opt,name=status,proto3,enum=furysports.xfury.orderbook.OrderBookStatus" json:"status,omitempty"`
}

func (m *OrderBook) Reset()      { *m = OrderBook{} }
func (*OrderBook) ProtoMessage() {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aebc657d8cdea59, []int{0}
}
func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}
func (m *OrderBook) XXX_Size() int {
	return m.Size()
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("furysports.xfury.orderbook.OrderBookStatus", OrderBookStatus_name, OrderBookStatus_value)
	proto.RegisterType((*OrderBook)(nil), "furysports.xfury.orderbook.OrderBook")
}

func init() { proto.RegisterFile("xfury/orderbook/orderbook.proto", fileDescriptor_3aebc657d8cdea59) }

var fileDescriptor_3aebc657d8cdea59 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2b, 0x2d, 0xaa,
	0xd4, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0x46, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x24, 0x41, 0xf2, 0xc5, 0x05, 0xf9, 0x45, 0x25, 0xc5, 0x7a, 0x20, 0xa6, 0x1e,
	0x5c, 0x81, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x95, 0x3e, 0x88, 0x05, 0xd1, 0xa0, 0xd4,
	0xc4, 0xc4, 0xc5, 0xe9, 0x0f, 0x52, 0xe3, 0x94, 0x9f, 0x9f, 0x2d, 0x24, 0xc3, 0xc5, 0x94, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xf3, 0xe8, 0x9e, 0x3c, 0x93, 0xa7, 0xcb, 0xab,
	0x7b, 0xf2, 0x4c, 0x99, 0x29, 0x41, 0x4c, 0x99, 0x29, 0x42, 0xfe, 0x5c, 0xc2, 0x05, 0x89, 0x45,
	0x25, 0x99, 0xc9, 0x99, 0x05, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xf1, 0xc9, 0xf9, 0xa5, 0x79, 0x25,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x2c, 0x4e, 0x72, 0x9f, 0xee, 0xc9, 0x4b, 0x55, 0x26, 0xe6, 0xe6,
	0x58, 0x29, 0x61, 0x51, 0xa4, 0x14, 0x24, 0x84, 0x22, 0xea, 0x0c, 0x12, 0x14, 0x32, 0xe1, 0xe2,
	0xca, 0x4f, 0x49, 0x29, 0x86, 0x9a, 0xc3, 0x0c, 0x36, 0x47, 0xf4, 0xd3, 0x3d, 0x79, 0x41, 0x88,
	0x39, 0x08, 0x39, 0xa5, 0x20, 0x4e, 0x10, 0x07, 0xa2, 0xcb, 0x89, 0x8b, 0xad, 0xb8, 0x24, 0xb1,
	0xa4, 0xb4, 0x58, 0x82, 0x45, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x4b, 0x0f, 0xa7, 0xa7, 0xf5, 0xe0,
	0x5e, 0x0b, 0x06, 0xeb, 0x08, 0x82, 0xea, 0xb4, 0xe2, 0xe9, 0x58, 0x20, 0xcf, 0x30, 0x63, 0x81,
	0x3c, 0xc3, 0x8b, 0x05, 0xf2, 0x0c, 0x5a, 0xcb, 0x18, 0xb9, 0xf8, 0xd1, 0x54, 0x0a, 0x29, 0x72,
	0xc9, 0xfa, 0x07, 0xb9, 0xb8, 0x06, 0xc5, 0x3b, 0xf9, 0xfb, 0x7b, 0xc7, 0x07, 0x87, 0x38, 0x86,
	0x84, 0x06, 0xc7, 0x87, 0xfa, 0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba, 0xba, 0x08, 0x30,
	0x08, 0x29, 0x73, 0xc9, 0x63, 0x2a, 0x81, 0x52, 0x8e, 0xce, 0x21, 0x9e, 0x61, 0xae, 0x02, 0x8c,
	0x42, 0xaa, 0x5c, 0x8a, 0x38, 0x15, 0x05, 0xb9, 0x06, 0xfb, 0xfb, 0x84, 0xb9, 0xba, 0x08, 0x30,
	0x09, 0xa9, 0x70, 0x29, 0xe0, 0x54, 0x16, 0xec, 0x1a, 0x12, 0xe2, 0xe3, 0xea, 0x22, 0xc0, 0xec,
	0xe4, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x06, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x69, 0x89, 0x79, 0xa0, 0x60, 0xd0, 0x85, 0x04,
	0x89, 0x3e, 0x38, 0xc9, 0x54, 0x20, 0x25, 0x9a, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70,
	0x02, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xd5, 0xb0, 0x34, 0x53, 0x02, 0x00, 0x00,
}

func (m *OrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.OddsCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.OddsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.ParticipationCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.ParticipationCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintOrderbook(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderbook(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderbook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovOrderbook(uint64(l))
	}
	if m.ParticipationCount != 0 {
		n += 1 + sovOrderbook(uint64(m.ParticipationCount))
	}
	if m.OddsCount != 0 {
		n += 1 + sovOrderbook(uint64(m.OddsCount))
	}
	if m.Status != 0 {
		n += 1 + sovOrderbook(uint64(m.Status))
	}
	return n
}

func sovOrderbook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderbook(x uint64) (n int) {
	return sovOrderbook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderbook
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
			return fmt.Errorf("proto: OrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrderbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationCount", wireType)
			}
			m.ParticipationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsCount", wireType)
			}
			m.OddsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderBookStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrderbook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderbook
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
func skipOrderbook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
				return 0, ErrInvalidLengthOrderbook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderbook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderbook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderbook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderbook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderbook = fmt.Errorf("proto: unexpected end of group")
)
