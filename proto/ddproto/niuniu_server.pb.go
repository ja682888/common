// Code generated by protoc-gen-go.
// source: niuniu_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of niuniu_client_poker from niuniu_base.proto

// Ignoring public import of niuniu_user_bill from niuniu_base.proto

// Ignoring public import of niuniu_desk_option from niuniu_base.proto

// Ignoring public import of niuniu_enum_protoid from niuniu_base.proto

// Ignoring public import of niuniu_enum_PokerType from niuniu_base.proto

// Ignoring public import of niuniu_enum_desk_state from niuniu_base.proto

// Ignoring public import of niuniu_enum_banker_rule from niuniu_base.proto

// 打出去的牌
type NiuniuSrvPoker struct {
	Pais             []*CommonSrvPokerPai  `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuSrvPoker) Reset()                    { *m = NiuniuSrvPoker{} }
func (m *NiuniuSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvPoker) ProtoMessage()               {}
func (*NiuniuSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{0} }

func (m *NiuniuSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuSrvPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

// desk 的信息
type NiuniuSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	DeskNumber       *string              `protobuf:"bytes,2,opt,name=deskNumber" json:"deskNumber,omitempty"`
	GameNumber       *int32               `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status           *NiuniuEnumDeskState `protobuf:"varint,5,opt,name=status,enum=ddproto.NiuniuEnumDeskState" json:"status,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	CircleNo         *int32               `protobuf:"varint,8,opt,name=circleNo" json:"circleNo,omitempty"`
	Owner            *uint32              `protobuf:"varint,10,opt,name=owner" json:"owner,omitempty"`
	CurrBanker       *uint32              `protobuf:"varint,11,opt,name=currBanker" json:"currBanker,omitempty"`
	DeskOption       *NiuniuDeskOption    `protobuf:"bytes,12,opt,name=deskOption" json:"deskOption,omitempty"`
	IsStart          *bool                `protobuf:"varint,13,opt,name=isStart" json:"isStart,omitempty"`
	IsOnDissolve     *bool                `protobuf:"varint,14,opt,name=isOnDissolve" json:"isOnDissolve,omitempty"`
	DissolveTime     *int64               `protobuf:"varint,15,opt,name=dissolveTime" json:"dissolveTime,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuSrvDesk) Reset()                    { *m = NiuniuSrvDesk{} }
func (m *NiuniuSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDesk) ProtoMessage()               {}
func (*NiuniuSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{1} }

func (m *NiuniuSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

func (m *NiuniuSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *NiuniuSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetStatus() NiuniuEnumDeskState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER
}

func (m *NiuniuSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *NiuniuSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCurrBanker() uint32 {
	if m != nil && m.CurrBanker != nil {
		return *m.CurrBanker
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskOption() *NiuniuDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *NiuniuSrvDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *NiuniuSrvDesk) GetIsOnDissolve() bool {
	if m != nil && m.IsOnDissolve != nil {
		return *m.IsOnDissolve
	}
	return false
}

func (m *NiuniuSrvDesk) GetDissolveTime() int64 {
	if m != nil && m.DissolveTime != nil {
		return *m.DissolveTime
	}
	return 0
}

// 用户属性
type NiuniuSrvUser struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *NiuniuUserBill `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool           `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32          `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers           *NiuniuSrvPoker `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	BankerScore      *int32          `protobuf:"varint,14,opt,name=bankerScore" json:"bankerScore,omitempty"`
	DoubleScore      *int32          `protobuf:"varint,15,opt,name=doubleScore" json:"doubleScore,omitempty"`
	IsReady          *bool           `protobuf:"varint,16,opt,name=isReady" json:"isReady,omitempty"`
	LastScore        *int32          `protobuf:"varint,17,opt,name=lastScore" json:"lastScore,omitempty"`
	DissolveState    *int32          `protobuf:"varint,18,opt,name=dissolveState" json:"dissolveState,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NiuniuSrvUser) Reset()                    { *m = NiuniuSrvUser{} }
func (m *NiuniuSrvUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvUser) ProtoMessage()               {}
func (*NiuniuSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{2} }

func (m *NiuniuSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *NiuniuSrvUser) GetBill() *NiuniuUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *NiuniuSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *NiuniuSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *NiuniuSrvUser) GetPokers() *NiuniuSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *NiuniuSrvUser) GetBankerScore() int32 {
	if m != nil && m.BankerScore != nil {
		return *m.BankerScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetDoubleScore() int32 {
	if m != nil && m.DoubleScore != nil {
		return *m.DoubleScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *NiuniuSrvUser) GetLastScore() int32 {
	if m != nil && m.LastScore != nil {
		return *m.LastScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetDissolveState() int32 {
	if m != nil && m.DissolveState != nil {
		return *m.DissolveState
	}
	return 0
}

// room 的信息
type NiuniuSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuSrvRoom) Reset()                    { *m = NiuniuSrvRoom{} }
func (m *NiuniuSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvRoom) ProtoMessage()               {}
func (*NiuniuSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{3} }

func (m *NiuniuSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func init() {
	proto.RegisterType((*NiuniuSrvPoker)(nil), "ddproto.niuniu_srv_poker")
	proto.RegisterType((*NiuniuSrvDesk)(nil), "ddproto.niuniu_srv_desk")
	proto.RegisterType((*NiuniuSrvUser)(nil), "ddproto.niuniu_srv_user")
	proto.RegisterType((*NiuniuSrvRoom)(nil), "ddproto.niuniu_srv_room")
}

var fileDescriptor32 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x2d, 0x63, 0xcb, 0xb1, 0xe9, 0x38, 0x4e, 0xd8, 0xa2, 0x60, 0xdc, 0x20, 0x15, 0x8c, 0x0e,
	0x5a, 0x6a, 0xb4, 0x5e, 0x3a, 0x74, 0x2b, 0xba, 0x04, 0x28, 0x12, 0x83, 0x36, 0xd0, 0xd1, 0x90,
	0x2d, 0xa2, 0x20, 0x2c, 0x89, 0x02, 0x29, 0xb9, 0xf1, 0xd8, 0xcf, 0xeb, 0x57, 0xf4, 0x57, 0x8a,
	0x3b, 0x52, 0xb6, 0x14, 0x64, 0xd2, 0xbd, 0x77, 0xef, 0x8e, 0xc7, 0xa7, 0x23, 0x7d, 0x9d, 0xab,
	0x2a, 0x57, 0xd5, 0xda, 0x4a, 0xb3, 0x97, 0x66, 0x56, 0x18, 0x5d, 0x6a, 0x76, 0x9e, 0x24, 0x18,
	0x4c, 0x6e, 0xb6, 0x3a, 0xcb, 0x74, 0xee, 0xb3, 0xeb, 0x42, 0xef, 0x6a, 0xcd, 0xe4, 0xda, 0x17,
	0x6e, 0x62, 0x2b, 0x1d, 0x35, 0x7d, 0xa2, 0x57, 0x75, 0x37, 0xb3, 0x77, 0x62, 0xf6, 0x89, 0x76,
	0x8b, 0x58, 0x59, 0x7e, 0x16, 0x76, 0xa2, 0xe1, 0xfc, 0x76, 0xe6, 0x3b, 0xcf, 0xea, 0xc6, 0xb5,
	0x70, 0x11, 0x2b, 0x81, 0x4a, 0x36, 0xa7, 0xdd, 0xf2, 0x50, 0x48, 0xde, 0x09, 0x49, 0x74, 0x39,
	0xbf, 0x3b, 0x56, 0xf8, 0xd6, 0x32, 0xaf, 0xb2, 0xf5, 0x02, 0x4a, 0x56, 0x87, 0x42, 0x0a, 0xd4,
	0x4e, 0xff, 0x76, 0xe8, 0xb8, 0x71, 0x74, 0x22, 0xed, 0x8e, 0xbd, 0xa5, 0x3d, 0xf8, 0xde, 0x27,
	0x9c, 0x84, 0x24, 0x0a, 0x84, 0x47, 0xec, 0x8e, 0x52, 0x88, 0x1e, 0xaa, 0x6c, 0x23, 0x0d, 0x3f,
	0x0b, 0x49, 0x34, 0x10, 0x0d, 0x06, 0xf2, 0xbf, 0xe2, 0x4c, 0xfa, 0x7c, 0x07, 0x6b, 0x1b, 0x0c,
	0xf4, 0x35, 0x5a, 0x67, 0xf7, 0x09, 0xef, 0xba, 0xbe, 0x0e, 0xb1, 0x2f, 0xb4, 0x67, 0xcb, 0xb8,
	0xac, 0x2c, 0x0f, 0x70, 0xf2, 0xf7, 0x2f, 0x4e, 0x0e, 0x07, 0xad, 0x41, 0x27, 0x85, 0x97, 0xb3,
	0x5b, 0x3a, 0x48, 0x63, 0x5b, 0xfe, 0x54, 0xb9, 0x34, 0xbc, 0x17, 0x92, 0x68, 0x24, 0x4e, 0x04,
	0x9b, 0xd0, 0xfe, 0x56, 0x99, 0x6d, 0x2a, 0x1f, 0x34, 0xef, 0xe3, 0x81, 0x47, 0xcc, 0xde, 0xd0,
	0x40, 0xff, 0x86, 0x2a, 0x8a, 0x55, 0x0e, 0xc0, 0x05, 0xb6, 0x95, 0x31, 0xdf, 0xe2, 0x7c, 0x27,
	0x0d, 0x1f, 0x62, 0xaa, 0xc1, 0xb0, 0xaf, 0xce, 0x80, 0xc7, 0xa2, 0x54, 0x3a, 0xe7, 0x17, 0x21,
	0x89, 0x86, 0xf3, 0x77, 0xcf, 0x87, 0xc5, 0x39, 0x35, 0x4a, 0x44, 0x43, 0xce, 0x38, 0x3d, 0x57,
	0x76, 0x59, 0xc6, 0xa6, 0xe4, 0xa3, 0x90, 0x44, 0x7d, 0x51, 0x43, 0x36, 0xa5, 0x17, 0xca, 0x3e,
	0xe6, 0xdf, 0x95, 0xb5, 0x3a, 0xdd, 0x4b, 0x7e, 0x89, 0xe9, 0x16, 0x07, 0x9a, 0xc4, 0xc7, 0x2b,
	0x95, 0x49, 0x3e, 0x0e, 0x49, 0xd4, 0x11, 0x2d, 0x6e, 0xfa, 0xef, 0xac, 0xf5, 0x2f, 0x2b, 0xeb,
	0x3c, 0x87, 0xaf, 0xff, 0x97, 0x23, 0xe1, 0x11, 0xfb, 0x48, 0xbb, 0x1b, 0x95, 0xa6, 0xe8, 0xda,
	0x70, 0x7e, 0xf3, 0xfc, 0x12, 0xa0, 0x5a, 0x83, 0x40, 0xa0, 0x0c, 0xbc, 0x84, 0x71, 0x52, 0x95,
	0x4b, 0xb4, 0xac, 0x2f, 0x8e, 0x18, 0xbc, 0x54, 0x79, 0x22, 0x9f, 0xd0, 0x90, 0x40, 0x38, 0xc0,
	0x3e, 0xd3, 0x1e, 0xae, 0xa7, 0xc5, 0xdb, 0xbe, 0x70, 0xc4, 0x71, 0x81, 0x85, 0x17, 0xb2, 0x90,
	0x0e, 0x37, 0x68, 0xf4, 0x72, 0xab, 0x8d, 0xb3, 0x21, 0x10, 0x4d, 0x0a, 0x14, 0x89, 0xae, 0x36,
	0xa9, 0x74, 0x8a, 0xb1, 0x53, 0x34, 0x28, 0xe7, 0xb2, 0x90, 0x71, 0x72, 0xe0, 0x57, 0xb5, 0xcb,
	0x08, 0xeb, 0x65, 0x71, 0x95, 0xd7, 0x58, 0x79, 0x22, 0xd8, 0x07, 0x3a, 0xaa, 0xbd, 0x5c, 0xc2,
	0x8e, 0x71, 0x86, 0x8a, 0x36, 0x39, 0xfd, 0x43, 0x5a, 0x0e, 0xc3, 0xfe, 0x36, 0xb6, 0x9a, 0xb4,
	0xb6, 0x7a, 0x42, 0xfb, 0x10, 0xc1, 0x5b, 0xc3, 0xb7, 0x12, 0x88, 0x23, 0x86, 0x59, 0x20, 0xfe,
	0x21, 0xf7, 0x32, 0xf5, 0x0f, 0xe5, 0x44, 0xd4, 0xd9, 0x95, 0x2a, 0x53, 0x89, 0x4f, 0x65, 0x20,
	0x4e, 0xc4, 0xe2, 0xd5, 0x82, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x60, 0xdf, 0xb7, 0x7c,
	0x04, 0x00, 0x00,
}
