// Code generated by protoc-gen-go. DO NOT EDIT.
// source: paoyao_server.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of paoyao_client_poker from paoyao_base.proto

// Ignoring public import of paoyao_user_bill from paoyao_base.proto

// Ignoring public import of paoyao_desk_option from paoyao_base.proto

// Ignoring public import of paoyao_srv_room from paoyao_base.proto

// Ignoring public import of paoyao_enum_protoid from paoyao_base.proto

// Ignoring public import of paoyao_enum_pokerType from paoyao_base.proto

// Ignoring public import of paoyao_snow_type from paoyao_base.proto

// Ignoring public import of paoyao_enum_desk_status from paoyao_base.proto

// 打出去的牌
type PaoyaoSrvPoker struct {
	Pais             []*CommonSrvPokerPai `protobuf:"bytes,1,rep,name=pais" json:"pais,omitempty"`
	Type             *PaoyaoEnumPokerType `protobuf:"varint,2,opt,name=Type,enum=ddproto.PaoyaoEnumPokerType" json:"Type,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PaoyaoSrvPoker) Reset()                    { *m = PaoyaoSrvPoker{} }
func (m *PaoyaoSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoSrvPoker) ProtoMessage()               {}
func (*PaoyaoSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

func (m *PaoyaoSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *PaoyaoSrvPoker) GetType() PaoyaoEnumPokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PaoyaoEnumPokerType_PAOYAO_POKER_TYPE_OTHER
}

// desk 的信息
type PaoyaoSrvDesk struct {
	DeskId         *int32                `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Pwd            *string               `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	GameNumber     *int32                `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId         *int32                `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status         *PaoyaoEnumDeskStatus `protobuf:"varint,5,opt,name=status,enum=ddproto.PaoyaoEnumDeskStatus" json:"status,omitempty"`
	CircleNo       *int32                `protobuf:"varint,7,opt,name=circleNo" json:"circleNo,omitempty"`
	CurrDeskScore  *int32                `protobuf:"varint,8,opt,name=currDeskScore" json:"currDeskScore,omitempty"`
	Owner          *uint32               `protobuf:"varint,11,opt,name=owner" json:"owner,omitempty"`
	IsDaikai       *bool                 `protobuf:"varint,12,opt,name=isDaikai" json:"isDaikai,omitempty"`
	DaikaiUser     *uint32               `protobuf:"varint,13,opt,name=daikaiUser" json:"daikaiUser,omitempty"`
	DeskOption     *PaoyaoDeskOption     `protobuf:"bytes,14,opt,name=deskOption" json:"deskOption,omitempty"`
	LastActUser    *uint32               `protobuf:"varint,15,opt,name=lastActUser" json:"lastActUser,omitempty"`
	LastChupaiUser *uint32               `protobuf:"varint,16,opt,name=lastChupaiUser" json:"lastChupaiUser,omitempty"`
	CurrActUser    *uint32               `protobuf:"varint,26,opt,name=currActUser" json:"currActUser,omitempty"`
	CanGuo         *bool                 `protobuf:"varint,27,opt,name=canGuo" json:"canGuo,omitempty"`
	IsOnDissolve   *bool                 `protobuf:"varint,18,opt,name=isOnDissolve" json:"isOnDissolve,omitempty"`
	DissolveTime   *int64                `protobuf:"varint,19,opt,name=dissolveTime" json:"dissolveTime,omitempty"`
	IsStart        *bool                 `protobuf:"varint,22,opt,name=isStart" json:"isStart,omitempty"`
	OneStartTime   *int64                `protobuf:"varint,23,opt,name=oneStartTime" json:"oneStartTime,omitempty"`
	AllStartTime   *int64                `protobuf:"varint,24,opt,name=allStartTime" json:"allStartTime,omitempty"`
	DissolveUser   *uint32               `protobuf:"varint,25,opt,name=dissolveUser" json:"dissolveUser,omitempty"`
	LastKangqi     *uint32               `protobuf:"varint,28,opt,name=lastKangqi" json:"lastKangqi,omitempty"`
	// 金币场
	IsCoinRoom       *bool  `protobuf:"varint,20,opt,name=isCoinRoom" json:"isCoinRoom,omitempty"`
	SurplusTime      *int32 `protobuf:"varint,21,opt,name=surplusTime" json:"surplusTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PaoyaoSrvDesk) Reset()                    { *m = PaoyaoSrvDesk{} }
func (m *PaoyaoSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoSrvDesk) ProtoMessage()               {}
func (*PaoyaoSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{1} }

func (m *PaoyaoSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *PaoyaoSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetStatus() PaoyaoEnumDeskStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PaoyaoEnumDeskStatus_PAOYAO_DESK_STATUS_WAIT_READY
}

func (m *PaoyaoSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetCurrDeskScore() int32 {
	if m != nil && m.CurrDeskScore != nil {
		return *m.CurrDeskScore
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

func (m *PaoyaoSrvDesk) GetDaikaiUser() uint32 {
	if m != nil && m.DaikaiUser != nil {
		return *m.DaikaiUser
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetDeskOption() *PaoyaoDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *PaoyaoSrvDesk) GetLastActUser() uint32 {
	if m != nil && m.LastActUser != nil {
		return *m.LastActUser
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetLastChupaiUser() uint32 {
	if m != nil && m.LastChupaiUser != nil {
		return *m.LastChupaiUser
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetCurrActUser() uint32 {
	if m != nil && m.CurrActUser != nil {
		return *m.CurrActUser
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetCanGuo() bool {
	if m != nil && m.CanGuo != nil {
		return *m.CanGuo
	}
	return false
}

func (m *PaoyaoSrvDesk) GetIsOnDissolve() bool {
	if m != nil && m.IsOnDissolve != nil {
		return *m.IsOnDissolve
	}
	return false
}

func (m *PaoyaoSrvDesk) GetDissolveTime() int64 {
	if m != nil && m.DissolveTime != nil {
		return *m.DissolveTime
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *PaoyaoSrvDesk) GetOneStartTime() int64 {
	if m != nil && m.OneStartTime != nil {
		return *m.OneStartTime
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetAllStartTime() int64 {
	if m != nil && m.AllStartTime != nil {
		return *m.AllStartTime
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetDissolveUser() uint32 {
	if m != nil && m.DissolveUser != nil {
		return *m.DissolveUser
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetLastKangqi() uint32 {
	if m != nil && m.LastKangqi != nil {
		return *m.LastKangqi
	}
	return 0
}

func (m *PaoyaoSrvDesk) GetIsCoinRoom() bool {
	if m != nil && m.IsCoinRoom != nil {
		return *m.IsCoinRoom
	}
	return false
}

func (m *PaoyaoSrvDesk) GetSurplusTime() int32 {
	if m != nil && m.SurplusTime != nil {
		return *m.SurplusTime
	}
	return 0
}

// 用户属性
type PaoyaoSrvUser struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *PaoyaoUserBill `protobuf:"bytes,2,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool           `protobuf:"varint,3,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32          `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Pokers           *PaoyaoSrvPoker `protobuf:"bytes,5,opt,name=pokers" json:"pokers,omitempty"`
	OutPai           *PaoyaoSrvPoker `protobuf:"bytes,6,opt,name=out_pai,json=outPai" json:"out_pai,omitempty"`
	IsPass           *bool           `protobuf:"varint,7,opt,name=isPass" json:"isPass,omitempty"`
	DeskScore        *int32          `protobuf:"varint,16,opt,name=deskScore" json:"deskScore,omitempty"`
	TeamMate         *uint32         `protobuf:"varint,17,opt,name=teamMate" json:"teamMate,omitempty"`
	IsKangQi         *bool           `protobuf:"varint,18,opt,name=isKangQi" json:"isKangQi,omitempty"`
	IsReady          *bool           `protobuf:"varint,8,opt,name=isReady" json:"isReady,omitempty"`
	LastScore        *int32          `protobuf:"varint,9,opt,name=lastScore" json:"lastScore,omitempty"`
	WxInfo           *WeixinInfo     `protobuf:"bytes,10,opt,name=wxInfo" json:"wxInfo,omitempty"`
	DissolveState    *int32          `protobuf:"varint,11,opt,name=dissolveState" json:"dissolveState,omitempty"`
	IsRobot          *bool           `protobuf:"varint,12,opt,name=isRobot" json:"isRobot,omitempty"`
	IsOnWhiteList    *bool           `protobuf:"varint,13,opt,name=isOnWhiteList" json:"isOnWhiteList,omitempty"`
	WhiteWinRate     *int32          `protobuf:"varint,14,opt,name=whiteWinRate" json:"whiteWinRate,omitempty"`
	IsLeave          *bool           `protobuf:"varint,15,opt,name=isLeave" json:"isLeave,omitempty"`
	PaoNo            *int32          `protobuf:"varint,19,opt,name=paoNo" json:"paoNo,omitempty"`
	IsOnGaming       *bool           `protobuf:"varint,20,opt,name=isOnGaming" json:"isOnGaming,omitempty"`
	CanPreEnd        *bool           `protobuf:"varint,21,opt,name=canPreEnd" json:"canPreEnd,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PaoyaoSrvUser) Reset()                    { *m = PaoyaoSrvUser{} }
func (m *PaoyaoSrvUser) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoSrvUser) ProtoMessage()               {}
func (*PaoyaoSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{2} }

func (m *PaoyaoSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PaoyaoSrvUser) GetBill() *PaoyaoUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *PaoyaoSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *PaoyaoSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *PaoyaoSrvUser) GetPokers() *PaoyaoSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *PaoyaoSrvUser) GetOutPai() *PaoyaoSrvPoker {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *PaoyaoSrvUser) GetIsPass() bool {
	if m != nil && m.IsPass != nil {
		return *m.IsPass
	}
	return false
}

func (m *PaoyaoSrvUser) GetDeskScore() int32 {
	if m != nil && m.DeskScore != nil {
		return *m.DeskScore
	}
	return 0
}

func (m *PaoyaoSrvUser) GetTeamMate() uint32 {
	if m != nil && m.TeamMate != nil {
		return *m.TeamMate
	}
	return 0
}

func (m *PaoyaoSrvUser) GetIsKangQi() bool {
	if m != nil && m.IsKangQi != nil {
		return *m.IsKangQi
	}
	return false
}

func (m *PaoyaoSrvUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *PaoyaoSrvUser) GetLastScore() int32 {
	if m != nil && m.LastScore != nil {
		return *m.LastScore
	}
	return 0
}

func (m *PaoyaoSrvUser) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PaoyaoSrvUser) GetDissolveState() int32 {
	if m != nil && m.DissolveState != nil {
		return *m.DissolveState
	}
	return 0
}

func (m *PaoyaoSrvUser) GetIsRobot() bool {
	if m != nil && m.IsRobot != nil {
		return *m.IsRobot
	}
	return false
}

func (m *PaoyaoSrvUser) GetIsOnWhiteList() bool {
	if m != nil && m.IsOnWhiteList != nil {
		return *m.IsOnWhiteList
	}
	return false
}

func (m *PaoyaoSrvUser) GetWhiteWinRate() int32 {
	if m != nil && m.WhiteWinRate != nil {
		return *m.WhiteWinRate
	}
	return 0
}

func (m *PaoyaoSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *PaoyaoSrvUser) GetPaoNo() int32 {
	if m != nil && m.PaoNo != nil {
		return *m.PaoNo
	}
	return 0
}

func (m *PaoyaoSrvUser) GetIsOnGaming() bool {
	if m != nil && m.IsOnGaming != nil {
		return *m.IsOnGaming
	}
	return false
}

func (m *PaoyaoSrvUser) GetCanPreEnd() bool {
	if m != nil && m.CanPreEnd != nil {
		return *m.CanPreEnd
	}
	return false
}

// desk快照索引列表
type PaoyaoSrvDeskSnapshotIdIndex struct {
	DeskId           []int32 `protobuf:"varint,1,rep,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PaoyaoSrvDeskSnapshotIdIndex) Reset()                    { *m = PaoyaoSrvDeskSnapshotIdIndex{} }
func (m *PaoyaoSrvDeskSnapshotIdIndex) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoSrvDeskSnapshotIdIndex) ProtoMessage()               {}
func (*PaoyaoSrvDeskSnapshotIdIndex) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{3} }

func (m *PaoyaoSrvDeskSnapshotIdIndex) GetDeskId() []int32 {
	if m != nil {
		return m.DeskId
	}
	return nil
}

// 牌桌快照
type PaoyaoSrvDeskSnapshot struct {
	DeskState        *PaoyaoSrvDesk   `protobuf:"bytes,1,opt,name=deskState" json:"deskState,omitempty"`
	Users            []*PaoyaoSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PaoyaoSrvDeskSnapshot) Reset()                    { *m = PaoyaoSrvDeskSnapshot{} }
func (m *PaoyaoSrvDeskSnapshot) String() string            { return proto.CompactTextString(m) }
func (*PaoyaoSrvDeskSnapshot) ProtoMessage()               {}
func (*PaoyaoSrvDeskSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{4} }

func (m *PaoyaoSrvDeskSnapshot) GetDeskState() *PaoyaoSrvDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

func (m *PaoyaoSrvDeskSnapshot) GetUsers() []*PaoyaoSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*PaoyaoSrvPoker)(nil), "ddproto.paoyao_srv_poker")
	proto.RegisterType((*PaoyaoSrvDesk)(nil), "ddproto.paoyao_srv_desk")
	proto.RegisterType((*PaoyaoSrvUser)(nil), "ddproto.paoyao_srv_user")
	proto.RegisterType((*PaoyaoSrvDeskSnapshotIdIndex)(nil), "ddproto.paoyao_srv_desk_snapshot_id_index")
	proto.RegisterType((*PaoyaoSrvDeskSnapshot)(nil), "ddproto.paoyao_srv_desk_snapshot")
}

func init() { proto.RegisterFile("paoyao_server.proto", fileDescriptor38) }

var fileDescriptor38 = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xcf, 0x6f, 0xdb, 0x36,
	0x14, 0x9e, 0xea, 0xd8, 0x71, 0xe8, 0x26, 0x4d, 0x99, 0xae, 0x63, 0xd2, 0xa0, 0xd0, 0x8c, 0x61,
	0x30, 0x30, 0x2c, 0xd8, 0x7c, 0x18, 0x06, 0xf4, 0x34, 0x34, 0x43, 0x11, 0xac, 0x4b, 0x3c, 0xa6,
	0x43, 0x8e, 0x02, 0x23, 0x71, 0x09, 0x11, 0x89, 0xd4, 0x44, 0x2a, 0x76, 0xae, 0xfb, 0x93, 0x76,
	0xde, 0x1f, 0x37, 0xbc, 0x47, 0x4a, 0x96, 0x0c, 0x0f, 0x3b, 0x89, 0xdf, 0xc7, 0xef, 0xfd, 0x22,
	0xdf, 0xa3, 0xc8, 0x51, 0x29, 0xcc, 0x93, 0x30, 0x89, 0x95, 0xd5, 0xa3, 0xac, 0xce, 0xca, 0xca,
	0x38, 0x43, 0x77, 0xb3, 0x0c, 0x17, 0x27, 0xc7, 0xa9, 0x29, 0x0a, 0xa3, 0xc3, 0x6e, 0x52, 0x9a,
	0x87, 0x46, 0x73, 0xf2, 0x32, 0x18, 0xde, 0x0a, 0x2b, 0x3d, 0x35, 0x5d, 0x91, 0xc3, 0xc6, 0x5b,
	0xf5, 0xe8, 0xc5, 0xf4, 0x3b, 0xb2, 0x53, 0x0a, 0x65, 0x59, 0x14, 0x0f, 0x66, 0x93, 0xf9, 0xe9,
	0x59, 0xf0, 0x7c, 0xd6, 0x38, 0x6e, 0x84, 0x0b, 0xa1, 0x38, 0x2a, 0xe9, 0x9c, 0xec, 0x7c, 0x7a,
	0x2a, 0x25, 0x7b, 0x16, 0x47, 0xb3, 0x83, 0xf9, 0xdb, 0xd6, 0x22, 0xb8, 0x96, 0xba, 0x2e, 0xbc,
	0x09, 0xa8, 0x38, 0x6a, 0xa7, 0x7f, 0x8f, 0xc8, 0x8b, 0x4e, 0xe8, 0x4c, 0xda, 0x07, 0xfa, 0x9a,
	0x8c, 0xe0, 0x7b, 0x91, 0xb1, 0x28, 0x8e, 0x66, 0x43, 0x1e, 0x10, 0x3d, 0x24, 0x83, 0x72, 0x99,
	0xa1, 0xfb, 0x3d, 0x0e, 0x4b, 0xfa, 0x96, 0x90, 0x3b, 0x51, 0xc8, 0xcb, 0xba, 0xb8, 0x95, 0x15,
	0x1b, 0xa0, 0xba, 0xc3, 0x80, 0xa7, 0xca, 0x98, 0xe2, 0x22, 0x63, 0x3b, 0xde, 0x93, 0x47, 0xf4,
	0x47, 0x32, 0xb2, 0x4e, 0xb8, 0xda, 0xb2, 0x21, 0xe6, 0x1a, 0x6f, 0xcd, 0x15, 0xc2, 0x26, 0x5e,
	0xc7, 0x83, 0x9e, 0x9e, 0x90, 0x71, 0xaa, 0xaa, 0x34, 0x97, 0x97, 0x86, 0xed, 0xa2, 0xcf, 0x16,
	0xd3, 0xaf, 0xc8, 0x7e, 0x5a, 0x57, 0xd5, 0xb9, 0xb4, 0x0f, 0xd7, 0xa9, 0xa9, 0x24, 0x1b, 0xa3,
	0xa0, 0x4f, 0xd2, 0x57, 0x64, 0x68, 0x96, 0x5a, 0x56, 0x6c, 0x12, 0x47, 0xb3, 0x7d, 0xee, 0x01,
	0xf8, 0x55, 0xf6, 0x5c, 0xa8, 0x07, 0xa1, 0xd8, 0xf3, 0x38, 0x9a, 0x8d, 0x79, 0x8b, 0xa1, 0xca,
	0x0c, 0x57, 0xbf, 0x5b, 0x59, 0xb1, 0x7d, 0x34, 0xeb, 0x30, 0xf4, 0x1d, 0x21, 0x90, 0xea, 0x55,
	0xe9, 0x94, 0xd1, 0xec, 0x20, 0x8e, 0x66, 0x93, 0xf9, 0x9b, 0xcd, 0x8a, 0xb0, 0x18, 0x83, 0x12,
	0xde, 0x91, 0xd3, 0x98, 0x4c, 0x72, 0x61, 0xdd, 0x4f, 0xa9, 0x43, 0xef, 0x2f, 0xd0, 0x7b, 0x97,
	0xa2, 0x5f, 0x93, 0x03, 0x80, 0xef, 0xef, 0xeb, 0x32, 0xa4, 0x70, 0x88, 0xa2, 0x0d, 0x16, 0x3c,
	0x41, 0xa5, 0x8d, 0xa7, 0x13, 0xef, 0xa9, 0x43, 0xc1, 0x75, 0xa4, 0x42, 0x7f, 0xa8, 0x0d, 0x7b,
	0x83, 0x25, 0x06, 0x44, 0xa7, 0xe4, 0xb9, 0xb2, 0x57, 0xfa, 0x5c, 0x59, 0x6b, 0xf2, 0x47, 0xc9,
	0x28, 0xee, 0xf6, 0x38, 0xd0, 0x64, 0x61, 0xfd, 0x49, 0x15, 0x92, 0x1d, 0xc5, 0xd1, 0x6c, 0xc0,
	0x7b, 0x1c, 0x65, 0x64, 0x57, 0xd9, 0x6b, 0x27, 0x2a, 0xc7, 0x5e, 0xa3, 0x8b, 0x06, 0x82, 0xb5,
	0xd1, 0x12, 0xd7, 0x68, 0xfd, 0x85, 0xb7, 0xee, 0x72, 0xa0, 0x11, 0x79, 0xbe, 0xd6, 0x30, 0xaf,
	0xe9, 0x72, 0xdd, 0x2c, 0xb0, 0xc8, 0x63, 0x2c, 0xb2, 0xc7, 0xc1, 0x75, 0xc1, 0xc9, 0xfc, 0x22,
	0xf4, 0xdd, 0x9f, 0x8a, 0x9d, 0xfa, 0xeb, 0x5a, 0x33, 0xb0, 0xaf, 0xec, 0x7b, 0xa3, 0x34, 0x37,
	0xa6, 0x60, 0xaf, 0x30, 0xd1, 0x0e, 0x03, 0xe7, 0x68, 0xeb, 0xaa, 0xcc, 0x6b, 0x8b, 0x69, 0x7c,
	0x8e, 0x4d, 0xd4, 0xa5, 0xa6, 0xff, 0x0c, 0x7b, 0x43, 0x53, 0x87, 0xb3, 0x85, 0x6f, 0x18, 0x9a,
	0x7d, 0x1e, 0x10, 0xfd, 0x96, 0xec, 0xdc, 0xaa, 0x3c, 0xc7, 0xa9, 0x99, 0xcc, 0x8f, 0x37, 0xdb,
	0x02, 0x54, 0x09, 0x08, 0x38, 0xca, 0x7c, 0x1f, 0x5e, 0xe9, 0x5c, 0x69, 0x89, 0xf3, 0x84, 0x7d,
	0xe8, 0x31, 0x74, 0xae, 0xd2, 0x99, 0x5c, 0x85, 0x61, 0xf2, 0x80, 0x7e, 0x4f, 0x46, 0x38, 0xd4,
	0x7e, 0x96, 0xb6, 0x84, 0x68, 0x5f, 0x0a, 0x1e, 0x84, 0x74, 0x4e, 0x76, 0x4d, 0xed, 0x92, 0x52,
	0x28, 0x36, 0xfa, 0x5f, 0x1b, 0x53, 0xbb, 0x85, 0x50, 0x50, 0x9f, 0xb2, 0x0b, 0x61, 0x2d, 0x8e,
	0xdd, 0x98, 0x07, 0x44, 0x4f, 0xc9, 0x5e, 0xd6, 0x0e, 0xdc, 0x21, 0x26, 0xb6, 0x26, 0xa0, 0x1c,
	0x27, 0x45, 0xf1, 0xab, 0x70, 0x92, 0xbd, 0xc4, 0x73, 0x69, 0xb1, 0x2f, 0x15, 0xee, 0xe4, 0x37,
	0x15, 0x3a, 0xae, 0xc5, 0xbe, 0x93, 0xb8, 0x14, 0xd9, 0x13, 0x0e, 0x31, 0x76, 0x12, 0x42, 0x88,
	0x07, 0x77, 0xe9, 0xe3, 0xed, 0xf9, 0x78, 0x2d, 0x41, 0xbf, 0x21, 0xa3, 0xe5, 0xea, 0x42, 0xff,
	0x61, 0x18, 0xc1, 0xc2, 0x8e, 0xda, 0xc2, 0x6e, 0xa4, 0x5a, 0x29, 0x0d, 0x5b, 0x3c, 0x48, 0xe0,
	0xbd, 0x68, 0x1a, 0xe7, 0xda, 0x41, 0x86, 0x13, 0xff, 0x5e, 0xf4, 0xc8, 0x90, 0x8a, 0xb9, 0x35,
	0x2e, 0x3c, 0x0c, 0x0d, 0x04, 0x7b, 0xb8, 0x9b, 0x9b, 0x7b, 0xe5, 0xe4, 0x47, 0x65, 0x1d, 0x3e,
	0x0d, 0x63, 0xde, 0x27, 0xa1, 0x65, 0x97, 0x00, 0x6e, 0x94, 0xe6, 0x10, 0xe4, 0x00, 0x83, 0xf4,
	0x38, 0x1f, 0xe3, 0xa3, 0x14, 0x8f, 0x12, 0x1f, 0x00, 0x8c, 0x81, 0x10, 0xee, 0xbc, 0x14, 0xe6,
	0xd2, 0xe0, 0xbc, 0x0d, 0xb9, 0x07, 0xbe, 0x85, 0xaf, 0xf4, 0x07, 0x51, 0x28, 0x7d, 0xb7, 0x6e,
	0xe1, 0x86, 0x81, 0x43, 0x4a, 0x85, 0x5e, 0x54, 0xf2, 0x67, 0x9d, 0x61, 0x03, 0x8f, 0xf9, 0x9a,
	0x98, 0xbe, 0x23, 0x5f, 0x6e, 0x3c, 0xf9, 0x89, 0xd5, 0xa2, 0xb4, 0xf7, 0xc6, 0x25, 0x2a, 0x4b,
	0x7c, 0x5b, 0x75, 0x7f, 0x02, 0x83, 0xf5, 0x4f, 0x60, 0xfa, 0x57, 0x44, 0xd8, 0x7f, 0x59, 0xd3,
	0x1f, 0x42, 0x33, 0xe0, 0x69, 0x46, 0x78, 0x03, 0x6c, 0x5b, 0x6b, 0x81, 0x88, 0xaf, 0xa5, 0xf4,
	0x8c, 0x0c, 0x61, 0x10, 0x2c, 0x7b, 0x86, 0x3f, 0xbb, 0xad, 0x36, 0x20, 0xe0, 0x5e, 0xb6, 0xf8,
	0x6c, 0x11, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe9, 0x77, 0xfb, 0x80, 0x07, 0x00, 0x00,
}
