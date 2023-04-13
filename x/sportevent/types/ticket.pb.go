// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xfury/sportevent/ticket.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// SportEventAddTicketPayload indicates data of add sport-event ticket
type SportEventAddTicketPayload struct {
	// uid is the universal unique identifier of the sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event.
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event.
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// odds is the list of odds of the sport-event.
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// status is the current status of the sport-event.
	Status SportEventStatus `protobuf:"varint,5,opt,name=status,proto3,enum=furysports.xfury.sportevent.SportEventStatus" json:"status,omitempty"`
	// creator is the address of the creator of sport-event.
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// min_bet_amount is the minimum allowed bet amount for a sport-event.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// meta contains human-readable metadata of the sport-event.
	Meta string `protobuf:"bytes,9,opt,name=meta,proto3" json:"meta,omitempty"`
	// sr_contribution_for_house is the portion of payout that should be paid by
	// sr
	SrContributionForHouse github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=sr_contribution_for_house,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sr_contribution_for_house"`
}

func (m *SportEventAddTicketPayload) Reset()         { *m = SportEventAddTicketPayload{} }
func (m *SportEventAddTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventAddTicketPayload) ProtoMessage()    {}
func (*SportEventAddTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_da3107fff151fd67, []int{0}
}
func (m *SportEventAddTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventAddTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventAddTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventAddTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventAddTicketPayload.Merge(m, src)
}
func (m *SportEventAddTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventAddTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventAddTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventAddTicketPayload proto.InternalMessageInfo

func (m *SportEventAddTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventAddTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEventAddTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEventAddTicketPayload) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *SportEventAddTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func (m *SportEventAddTicketPayload) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SportEventAddTicketPayload) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

// SportEventUpdateTicketPayload indicates data of update sport-event ticket
type SportEventUpdateTicketPayload struct {
	// uid is the uuid of the sport-event
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// min_bet_amount is the minimum allowed bet amount for a sport-event.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// status is the status of the resolution.
	Status SportEventStatus `protobuf:"varint,6,opt,name=status,proto3,enum=furysports.xfury.sportevent.SportEventStatus" json:"status,omitempty"`
}

func (m *SportEventUpdateTicketPayload) Reset()         { *m = SportEventUpdateTicketPayload{} }
func (m *SportEventUpdateTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventUpdateTicketPayload) ProtoMessage()    {}
func (*SportEventUpdateTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_da3107fff151fd67, []int{1}
}
func (m *SportEventUpdateTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventUpdateTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventUpdateTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventUpdateTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventUpdateTicketPayload.Merge(m, src)
}
func (m *SportEventUpdateTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventUpdateTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventUpdateTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventUpdateTicketPayload proto.InternalMessageInfo

func (m *SportEventUpdateTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventUpdateTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEventUpdateTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEventUpdateTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

// SportEventResolutionTicketPayload indicates data of the
// resolution of the sport-event ticket.
type SportEventResolutionTicketPayload struct {
	// uid is the universal unique identifier of sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// resolution_ts is the resolution timestamp of the sport-event.
	ResolutionTS uint64 `protobuf:"varint,2,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// winner_odds_uids is the universal unique identifier list of the winner
	// odds.
	WinnerOddsUIDs []string `protobuf:"bytes,3,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// status is the status of the resolution.
	Status SportEventStatus `protobuf:"varint,4,opt,name=status,proto3,enum=furysports.xfury.sportevent.SportEventStatus" json:"status,omitempty"`
}

func (m *SportEventResolutionTicketPayload) Reset()         { *m = SportEventResolutionTicketPayload{} }
func (m *SportEventResolutionTicketPayload) String() string { return proto.CompactTextString(m) }
func (*SportEventResolutionTicketPayload) ProtoMessage()    {}
func (*SportEventResolutionTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_da3107fff151fd67, []int{2}
}
func (m *SportEventResolutionTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEventResolutionTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEventResolutionTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEventResolutionTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEventResolutionTicketPayload.Merge(m, src)
}
func (m *SportEventResolutionTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *SportEventResolutionTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEventResolutionTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SportEventResolutionTicketPayload proto.InternalMessageInfo

func (m *SportEventResolutionTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEventResolutionTicketPayload) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *SportEventResolutionTicketPayload) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *SportEventResolutionTicketPayload) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SportEventAddTicketPayload)(nil), "furysports.xfury.sportevent.SportEventAddTicketPayload")
	proto.RegisterType((*SportEventUpdateTicketPayload)(nil), "furysports.xfury.sportevent.SportEventUpdateTicketPayload")
	proto.RegisterType((*SportEventResolutionTicketPayload)(nil), "furysports.xfury.sportevent.SportEventResolutionTicketPayload")
}

func init() { proto.RegisterFile("xfury/sportevent/ticket.proto", fileDescriptor_da3107fff151fd67) }

var fileDescriptor_da3107fff151fd67 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x6f, 0xd6, 0x36, 0xdd, 0xcc, 0xa8, 0x90, 0x41, 0x10, 0x2a, 0x88, 0xb3, 0x1d, 0xa6, 0x4a,
	0xb0, 0x44, 0x6c, 0xbc, 0xc0, 0xca, 0xfe, 0x30, 0x71, 0x00, 0xb9, 0x9b, 0x90, 0xb8, 0x44, 0x69,
	0xed, 0x6e, 0xd1, 0x16, 0xbb, 0xb2, 0x1d, 0x60, 0x6f, 0xc0, 0x0d, 0xde, 0x81, 0x0b, 0x8f, 0xb2,
	0xe3, 0x8e, 0x88, 0x83, 0x85, 0xb2, 0x5b, 0x9f, 0x02, 0xd9, 0x2b, 0x4b, 0xb7, 0x69, 0x48, 0x2b,
	0x1c, 0xb8, 0xc4, 0x9f, 0xbf, 0xdf, 0xf7, 0xfb, 0xd9, 0xf9, 0x7e, 0x8e, 0x03, 0x1e, 0x0d, 0x72,
	0x71, 0x14, 0xc9, 0x21, 0x17, 0x8a, 0xbe, 0xa7, 0x4c, 0x45, 0x2a, 0xed, 0x1f, 0x50, 0x15, 0x0e,
	0x05, 0x57, 0x1c, 0xb6, 0x0c, 0x6a, 0x41, 0x19, 0x9a, 0x30, 0x2c, 0x0b, 0x5b, 0x0b, 0x97, 0x99,
	0x36, 0x8c, 0x6d, 0x7c, 0x46, 0x6f, 0xb5, 0x2e, 0x97, 0x70, 0x42, 0xe4, 0x18, 0xbb, 0xb7, 0xc7,
	0xf7, 0xb8, 0x0d, 0x23, 0x13, 0x9d, 0x65, 0x17, 0x3f, 0xd7, 0x41, 0xab, 0x6b, 0xea, 0x37, 0x4c,
	0xfd, 0x1a, 0x21, 0x3b, 0x76, 0x3b, 0x6f, 0x92, 0xa3, 0x43, 0x9e, 0x10, 0x18, 0x80, 0x6a, 0x9e,
	0x12, 0xcf, 0x09, 0x9c, 0xf6, 0x5c, 0xa7, 0x59, 0x68, 0x54, 0xdd, 0xdd, 0x5e, 0x1f, 0x69, 0x64,
	0xb2, 0xd8, 0x3c, 0xe0, 0x2a, 0x98, 0x95, 0x2a, 0x11, 0x2a, 0x56, 0xd2, 0x9b, 0x09, 0x9c, 0x76,
	0xad, 0xf3, 0xa0, 0xd0, 0xa8, 0xd1, 0x35, 0xb9, 0x9d, 0xee, 0x48, 0xa3, 0x73, 0x18, 0x9f, 0x47,
	0xf0, 0x09, 0x70, 0x29, 0x23, 0x86, 0x52, 0xb5, 0x94, 0xbb, 0x85, 0x46, 0xf5, 0x0d, 0x46, 0x2c,
	0x61, 0x0c, 0xe1, 0xf1, 0x08, 0x9f, 0x83, 0x9a, 0x79, 0x0d, 0xaf, 0x16, 0x54, 0xdb, 0xb7, 0x56,
	0x82, 0xf0, 0xfa, 0x16, 0x85, 0xaf, 0x09, 0x91, 0xd8, 0x56, 0xc3, 0x75, 0xe0, 0x4a, 0x95, 0xa8,
	0x5c, 0x7a, 0xf5, 0xc0, 0x69, 0x37, 0x57, 0x9e, 0xfe, 0x89, 0x57, 0x76, 0xa0, 0x6b, 0x39, 0x78,
	0xcc, 0x85, 0x1e, 0x68, 0xf4, 0x05, 0x4d, 0x14, 0x17, 0x9e, 0x6b, 0x7a, 0x80, 0x7f, 0x4f, 0xe1,
	0x0e, 0x68, 0x66, 0x29, 0x8b, 0x7b, 0x54, 0xc5, 0x49, 0xc6, 0x73, 0xa6, 0xbc, 0x86, 0x6d, 0x52,
	0x78, 0xac, 0x51, 0xe5, 0x87, 0x46, 0x4b, 0x7b, 0xa9, 0xda, 0xcf, 0x7b, 0x61, 0x9f, 0x67, 0x51,
	0x9f, 0xcb, 0x8c, 0xcb, 0xf1, 0xb0, 0x2c, 0xc9, 0x41, 0xa4, 0x8e, 0x86, 0x54, 0x86, 0xdb, 0x4c,
	0xe1, 0xf9, 0x2c, 0x65, 0x1d, 0xaa, 0xd6, 0xac, 0x06, 0xdc, 0x02, 0x0d, 0xa3, 0x38, 0xa0, 0xd4,
	0x9b, 0x9d, 0x4a, 0xce, 0xed, 0x51, 0xb5, 0x49, 0x29, 0x84, 0xa0, 0x96, 0x51, 0x95, 0x78, 0x73,
	0x76, 0xd7, 0x36, 0x86, 0x5f, 0x1d, 0xf0, 0x50, 0x8a, 0xb8, 0xcf, 0x99, 0x12, 0x69, 0x2f, 0x57,
	0x29, 0x67, 0xf1, 0x80, 0x8b, 0x78, 0x9f, 0xe7, 0x92, 0x7a, 0xc0, 0xae, 0x47, 0x6f, 0xb6, 0x5e,
	0xa1, 0xd1, 0xfd, 0xae, 0x78, 0x31, 0xa1, 0xb8, 0xc9, 0xc5, 0x4b, 0xa3, 0x37, 0xd2, 0xe8, 0xfa,
	0xc5, 0xf0, 0xf5, 0xd0, 0xe2, 0xa7, 0x2a, 0x78, 0x5c, 0xfa, 0xb1, 0x3b, 0x24, 0x89, 0xa2, 0xff,
	0xdf, 0xa1, 0xbc, 0x6a, 0x7f, 0xed, 0xdf, 0xda, 0x5f, 0xff, 0x2b, 0xfb, 0xcb, 0xd3, 0xef, 0x4e,
	0x7f, 0xfa, 0x17, 0xbf, 0xcd, 0x80, 0x85, 0x12, 0xc4, 0x54, 0xf2, 0x43, 0xeb, 0xd6, 0x4d, 0xed,
	0xd8, 0x02, 0xb7, 0xc5, 0x39, 0xb9, 0xf4, 0x64, 0xa1, 0xd0, 0x68, 0x7e, 0x42, 0xd5, 0xf4, 0xf9,
	0x62, 0x21, 0xbe, 0x38, 0x85, 0x18, 0xdc, 0xf9, 0x90, 0x32, 0x46, 0x45, 0x6c, 0xbe, 0xf1, 0x38,
	0x4f, 0x89, 0x31, 0xab, 0xda, 0x9e, 0xeb, 0x2c, 0x15, 0x1a, 0x35, 0xdf, 0x5a, 0xcc, 0x5c, 0x02,
	0xbb, 0xdb, 0xeb, 0x72, 0xa4, 0xd1, 0x95, 0x6a, 0x7c, 0x25, 0x33, 0xd1, 0xaa, 0xda, 0xf4, 0xad,
	0xea, 0xbc, 0x3a, 0x2e, 0x7c, 0xe7, 0xa4, 0xf0, 0x9d, 0x9f, 0x85, 0xef, 0x7c, 0x39, 0xf5, 0x2b,
	0x27, 0xa7, 0x7e, 0xe5, 0xfb, 0xa9, 0x5f, 0x79, 0xf7, 0x6c, 0xc2, 0xba, 0x41, 0xc2, 0x8c, 0xe2,
	0xf2, 0x99, 0x7a, 0x64, 0x6f, 0xeb, 0x8f, 0x17, 0x7e, 0x06, 0xc6, 0xc9, 0x9e, 0x6b, 0xef, 0xe6,
	0xd5, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x49, 0x0e, 0x47, 0x2c, 0x06, 0x00, 0x00,
}

func (m *SportEventAddTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventAddTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventAddTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SrContributionForHouse.Size()
		i -= size
		if _, err := m.SrContributionForHouse.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x4a
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SportEventUpdateTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventUpdateTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventUpdateTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SportEventResolutionTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEventResolutionTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEventResolutionTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SportEventAddTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.SrContributionForHouse.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func (m *SportEventUpdateTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func (m *SportEventResolutionTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovTicket(uint64(m.ResolutionTS))
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SportEventAddTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: SportEventAddTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventAddTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Odds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Odds = append(m.Odds, &Odds{})
			if err := m.Odds[len(m.Odds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrContributionForHouse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SrContributionForHouse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *SportEventUpdateTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: SportEventUpdateTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventUpdateTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *SportEventResolutionTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: SportEventResolutionTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEventResolutionTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolutionTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerOddsUIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerOddsUIDs = append(m.WinnerOddsUIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
