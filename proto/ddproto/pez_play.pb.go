// Code generated by protoc-gen-go.
// source: pez_play.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_tuozi from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_PEZTYPE from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_Base from pez_base.proto

// Ignoring public import of pez_enum_Bet from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of pez_enum_TuoziType from pez_base.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of pez_DissolveDesk from pez_desk.proto

// Ignoring public import of pez_AckDissolveDesk from pez_desk.proto

// Ignoring public import of pez_Ready from pez_desk.proto

// Ignoring public import of pez_AckReady from pez_desk.proto

// Ignoring public import of pez_EndLotteryInfo from pez_desk.proto

// Ignoring public import of pez_SendCurrentResult from pez_desk.proto

// Ignoring public import of EndLottery from pez_desk.proto

// Ignoring public import of pez_SendEndLottery from pez_desk.proto

// Ignoring public import of pez_UserBean from pez_desk.proto

// Ignoring public import of pez_Bill from pez_desk.proto

// 链接类型
type PEZ_RECONNECT_TYPE int32

const (
	PEZ_RECONNECT_TYPE_PEZ_NORMAL    PEZ_RECONNECT_TYPE = 1
	PEZ_RECONNECT_TYPE_PEZ_RECONNECT PEZ_RECONNECT_TYPE = 2
)

var PEZ_RECONNECT_TYPE_name = map[int32]string{
	1: "PEZ_NORMAL",
	2: "PEZ_RECONNECT",
}
var PEZ_RECONNECT_TYPE_value = map[string]int32{
	"PEZ_NORMAL":    1,
	"PEZ_RECONNECT": 2,
}

func (x PEZ_RECONNECT_TYPE) Enum() *PEZ_RECONNECT_TYPE {
	p := new(PEZ_RECONNECT_TYPE)
	*p = x
	return p
}
func (x PEZ_RECONNECT_TYPE) String() string {
	return proto.EnumName(PEZ_RECONNECT_TYPE_name, int32(x))
}
func (x *PEZ_RECONNECT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PEZ_RECONNECT_TYPE_value, data, "PEZ_RECONNECT_TYPE")
	if err != nil {
		return err
	}
	*x = PEZ_RECONNECT_TYPE(value)
	return nil
}
func (PEZ_RECONNECT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

// 积分
type PezUserCoinBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64  `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezUserCoinBean) Reset()                    { *m = PezUserCoinBean{} }
func (m *PezUserCoinBean) String() string            { return proto.CompactTextString(m) }
func (*PezUserCoinBean) ProtoMessage()               {}
func (*PezUserCoinBean) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

func (m *PezUserCoinBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezUserCoinBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 开局（接收服务端消息）
type Pez_Opening struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrPlayCount    *int32             `protobuf:"varint,2,opt,name=CurrPlayCount" json:"CurrPlayCount,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,3,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_Opening) Reset()                    { *m = Pez_Opening{} }
func (m *Pez_Opening) String() string            { return proto.CompactTextString(m) }
func (*Pez_Opening) ProtoMessage()               {}
func (*Pez_Opening) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{1} }

func (m *Pez_Opening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Opening) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *Pez_Opening) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

// 发牌
type Pez_DealCards struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Banker           *uint32         `protobuf:"varint,2,opt,name=banker" json:"banker,omitempty"`
	Dice1            *int32          `protobuf:"varint,3,opt,name=dice1" json:"dice1,omitempty"`
	Dice2            *int32          `protobuf:"varint,4,opt,name=dice2" json:"dice2,omitempty"`
	LastScore        []*PezLastScore `protobuf:"bytes,5,rep,name=lastScore" json:"lastScore,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_DealCards) Reset()                    { *m = Pez_DealCards{} }
func (m *Pez_DealCards) String() string            { return proto.CompactTextString(m) }
func (*Pez_DealCards) ProtoMessage()               {}
func (*Pez_DealCards) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{2} }

func (m *Pez_DealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_DealCards) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func (m *Pez_DealCards) GetDice1() int32 {
	if m != nil && m.Dice1 != nil {
		return *m.Dice1
	}
	return 0
}

func (m *Pez_DealCards) GetDice2() int32 {
	if m != nil && m.Dice2 != nil {
		return *m.Dice2
	}
	return 0
}

func (m *Pez_DealCards) GetLastScore() []*PezLastScore {
	if m != nil {
		return m.LastScore
	}
	return nil
}

type PezLastScore struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int64  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezLastScore) Reset()                    { *m = PezLastScore{} }
func (m *PezLastScore) String() string            { return proto.CompactTextString(m) }
func (*PezLastScore) ProtoMessage()               {}
func (*PezLastScore) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{3} }

func (m *PezLastScore) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezLastScore) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// 押注
type Pez_Bet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetNum           *int64  `protobuf:"varint,2,opt,name=betNum" json:"betNum,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_Bet) Reset()                    { *m = Pez_Bet{} }
func (m *Pez_Bet) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bet) ProtoMessage()               {}
func (*Pez_Bet) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{4} }

func (m *Pez_Bet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_Bet) GetBetNum() int64 {
	if m != nil && m.BetNum != nil {
		return *m.BetNum
	}
	return 0
}

func (m *Pez_Bet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_AckBet struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BetCount         *int64       `protobuf:"varint,3,opt,name=betCount" json:"betCount,omitempty"`
	Time             *int32       `protobuf:"varint,4,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckBet) Reset()                    { *m = Pez_AckBet{} }
func (m *Pez_AckBet) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckBet) ProtoMessage()               {}
func (*Pez_AckBet) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{5} }

func (m *Pez_AckBet) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckBet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckBet) GetBetCount() int64 {
	if m != nil && m.BetCount != nil {
		return *m.BetCount
	}
	return 0
}

func (m *Pez_AckBet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_BCOpenPai struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	HandPai          []*PezBase_PaiInfo `protobuf:"bytes,2,rep,name=handPai" json:"handPai,omitempty"`
	KeyValue         *int32             `protobuf:"varint,3,opt,name=keyValue" json:"keyValue,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_BCOpenPai) Reset()                    { *m = Pez_BCOpenPai{} }
func (m *Pez_BCOpenPai) String() string            { return proto.CompactTextString(m) }
func (*Pez_BCOpenPai) ProtoMessage()               {}
func (*Pez_BCOpenPai) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{6} }

func (m *Pez_BCOpenPai) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_BCOpenPai) GetHandPai() []*PezBase_PaiInfo {
	if m != nil {
		return m.HandPai
	}
	return nil
}

func (m *Pez_BCOpenPai) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

// 发送游戏信息(广播)
type Pez_SendGameInfo struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 1. 首先是牌桌的玩家数据（玩家数据包括其id昵称筹码头像等基本信息，其手牌数据，以及自己打开牌的数据，还有状态是否已经押注了，玩家在整局的总输赢）
	PlayerInfo []*PezBase_PlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	// 2. 桌面信息（包括：游戏是否结束，当前哪个玩家未押注，倒计时剩余时间）
	DeskGameInfo *PezBase_DeskGameInfo `protobuf:"bytes,3,opt,name=deskGameInfo" json:"deskGameInfo,omitempty"`
	//
	SenderUserId     *uint32 `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *bool   `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_SendGameInfo) Reset()                    { *m = Pez_SendGameInfo{} }
func (m *Pez_SendGameInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendGameInfo) ProtoMessage()               {}
func (*Pez_SendGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{7} }

func (m *Pez_SendGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_SendGameInfo) GetPlayerInfo() []*PezBase_PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetDeskGameInfo() *PezBase_DeskGameInfo {
	if m != nil {
		return m.DeskGameInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *Pez_SendGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

func init() {
	proto.RegisterType((*PezUserCoinBean)(nil), "ddproto.pez_user_coin_bean")
	proto.RegisterType((*Pez_Opening)(nil), "ddproto.pez_Opening")
	proto.RegisterType((*Pez_DealCards)(nil), "ddproto.pez_DealCards")
	proto.RegisterType((*PezLastScore)(nil), "ddproto.pez_lastScore")
	proto.RegisterType((*Pez_Bet)(nil), "ddproto.pez_Bet")
	proto.RegisterType((*Pez_AckBet)(nil), "ddproto.pez_AckBet")
	proto.RegisterType((*Pez_BCOpenPai)(nil), "ddproto.pez_BCOpenPai")
	proto.RegisterType((*Pez_SendGameInfo)(nil), "ddproto.pez_SendGameInfo")
	proto.RegisterEnum("ddproto.PEZ_RECONNECT_TYPE", PEZ_RECONNECT_TYPE_name, PEZ_RECONNECT_TYPE_value)
}

var fileDescriptor39 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0xd2, 0x94, 0x49, 0x5c, 0xa5, 0xdb, 0x82, 0x42, 0x41, 0xa8, 0xb2, 0x38, 0x94,
	0x1e, 0x22, 0x88, 0x90, 0x7a, 0x6e, 0xdc, 0x88, 0x56, 0x82, 0xc4, 0x4a, 0x0b, 0x12, 0x48, 0xc8,
	0xda, 0xd8, 0x03, 0xb5, 0x12, 0xaf, 0x23, 0x7f, 0x1c, 0xca, 0x81, 0x9f, 0xc0, 0x1f, 0xe3, 0x4f,
	0x31, 0xb3, 0xeb, 0x46, 0x8d, 0xdc, 0x43, 0x2e, 0x96, 0x67, 0xe6, 0xcd, 0xbc, 0xf7, 0x66, 0x07,
	0xf6, 0x56, 0xf8, 0x3b, 0x58, 0x2d, 0xe5, 0xdd, 0x60, 0x95, 0xa5, 0x45, 0x2a, 0xda, 0x51, 0xa4,
	0x7f, 0x8e, 0x0e, 0xc2, 0x34, 0x49, 0x52, 0x15, 0x84, 0xcb, 0x18, 0x55, 0x61, 0xaa, 0x47, 0x1a,
	0x3d, 0x97, 0x39, 0x56, 0xf1, 0x7e, 0x05, 0x42, 0x55, 0x26, 0x0f, 0x21, 0x11, 0xe6, 0x0b, 0x13,
	0xbb, 0x43, 0x10, 0x9c, 0x29, 0x73, 0xcc, 0x82, 0x30, 0x8d, 0x55, 0x30, 0x47, 0xa9, 0xc4, 0x1e,
	0xec, 0x70, 0xe6, 0x2a, 0xea, 0x5b, 0xc7, 0xd6, 0x89, 0x23, 0xba, 0xd0, 0xe4, 0x62, 0xbf, 0x41,
	0x91, 0xed, 0xfe, 0x81, 0x0e, 0xf7, 0x4c, 0x57, 0xa8, 0x62, 0xf5, 0x4b, 0xbc, 0x81, 0x9d, 0x5b,
	0x94, 0x11, 0x66, 0x1a, 0xdc, 0x19, 0x1e, 0x0e, 0x2a, 0x91, 0x03, 0x9f, 0xbf, 0x97, 0xba, 0x26,
	0x9e, 0x81, 0xe3, 0x95, 0x59, 0xe6, 0x93, 0x17, 0x2f, 0x2d, 0x55, 0xa1, 0x67, 0xb5, 0xc4, 0x10,
	0x1c, 0x66, 0xf2, 0x68, 0xfa, 0x88, 0x98, 0xf3, 0xbe, 0x7d, 0x6c, 0xd3, 0x8c, 0x97, 0xeb, 0x19,
	0x75, 0x75, 0xee, 0x5f, 0x0b, 0x1c, 0x4e, 0x5f, 0xa0, 0x5c, 0x7a, 0x32, 0x8b, 0xf2, 0x2d, 0x25,
	0x90, 0xab, 0xb9, 0x54, 0x0b, 0x42, 0x35, 0xb4, 0x2b, 0x07, 0x5a, 0x51, 0x1c, 0xe2, 0x7b, 0xe2,
	0x64, 0x29, 0x55, 0x38, 0xec, 0x37, 0x75, 0xf8, 0x16, 0x9e, 0x2e, 0x65, 0x5e, 0x5c, 0x87, 0x69,
	0x86, 0xfd, 0x96, 0x56, 0xf5, 0x7c, 0x43, 0xd5, 0xba, 0xea, 0x0e, 0x8c, 0x9e, 0x75, 0xa2, 0xb6,
	0x3f, 0x1a, 0x9d, 0xeb, 0x39, 0x66, 0x81, 0x67, 0xd0, 0x66, 0xfc, 0x08, 0x8b, 0x1a, 0x92, 0x35,
	0x62, 0x31, 0x29, 0x13, 0x03, 0xe5, 0xcd, 0x17, 0x71, 0x82, 0x46, 0xa2, 0x3b, 0x07, 0xe0, 0xc6,
	0xf3, 0x70, 0xc1, 0xbd, 0x5b, 0xbb, 0xae, 0x18, 0x8c, 0xeb, 0x1e, 0xec, 0x12, 0x83, 0x79, 0x03,
	0x7b, 0x83, 0x43, 0xfb, 0x76, 0x7f, 0x18, 0x33, 0x23, 0x8f, 0xdf, 0xd7, 0x97, 0x71, 0x4d, 0xe2,
	0x29, 0xb4, 0x6f, 0xa5, 0x8a, 0xa8, 0x44, 0x13, 0x79, 0x2d, 0x2f, 0x36, 0xd6, 0xc2, 0xf7, 0x17,
	0x50, 0xf1, 0x4a, 0xfd, 0x4c, 0x99, 0x6c, 0x81, 0x77, 0x5f, 0xe5, 0xb2, 0xbc, 0xb7, 0xf0, 0xcf,
	0x82, 0x1e, 0xc3, 0xae, 0x51, 0x45, 0x1f, 0x65, 0x82, 0x1a, 0xb6, 0x9d, 0x93, 0x77, 0xe4, 0x9e,
	0xce, 0x87, 0xa4, 0x50, 0x4f, 0xc5, 0xfd, 0xea, 0x11, 0xee, 0x35, 0x46, 0x7c, 0x80, 0x2e, 0xdf,
	0xfa, 0x3d, 0x8f, 0x96, 0xd0, 0x19, 0xbe, 0xae, 0xf7, 0x5c, 0x3c, 0x40, 0x89, 0x43, 0xe8, 0xe6,
	0xa4, 0x0e, 0xb3, 0x2f, 0xc6, 0x76, 0x53, 0xdb, 0x3e, 0x80, 0x4e, 0x9c, 0xcf, 0x30, 0x4c, 0x95,
	0xc2, 0xb0, 0xa0, 0x8b, 0xb0, 0x4e, 0x76, 0x4f, 0xcf, 0x40, 0xf8, 0xe3, 0xef, 0xc1, 0x6c, 0xec,
	0x4d, 0x27, 0x93, 0xb1, 0x77, 0x13, 0xdc, 0x7c, 0xf3, 0xc7, 0xb4, 0x31, 0xe0, 0xec, 0x64, 0x3a,
	0xfb, 0x7c, 0xfe, 0xa9, 0x67, 0x89, 0x7d, 0x70, 0x36, 0x50, 0xbd, 0xc6, 0xa8, 0x71, 0x69, 0xfb,
	0x4f, 0x7c, 0xcb, 0x6f, 0xf8, 0xf6, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xc3, 0x0f, 0x0f,
	0xe5, 0x03, 0x00, 0x00,
}
