// Code generated by protoc-gen-go.
// source: pez_desk.proto
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

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_timerInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_lastScore from pez_base.proto

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

// 房主解散房间(未开局)
type Pez_DissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_DissolveDesk) Reset()                    { *m = Pez_DissolveDesk{} }
func (m *Pez_DissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_DissolveDesk) ProtoMessage()               {}
func (*Pez_DissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{0} }

func (m *Pez_DissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_DissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 解散房间回复
type Pez_AckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckDissolveDesk) Reset()                    { *m = Pez_AckDissolveDesk{} }
func (m *Pez_AckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckDissolveDesk) ProtoMessage()               {}
func (*Pez_AckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{1} }

func (m *Pez_AckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *Pez_AckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 准备游戏
type Pez_Ready struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_Ready) Reset()                    { *m = Pez_Ready{} }
func (m *Pez_Ready) String() string            { return proto.CompactTextString(m) }
func (*Pez_Ready) ProtoMessage()               {}
func (*Pez_Ready) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{2} }

func (m *Pez_Ready) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Ready) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type Pez_AckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckReady) Reset()                    { *m = Pez_AckReady{} }
func (m *Pez_AckReady) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckReady) ProtoMessage()               {}
func (*Pez_AckReady) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{3} }

func (m *Pez_AckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *Pez_AckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type Pez_EndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	IsbigWin         *bool   `protobuf:"varint,3,opt,name=isbigWin" json:"isbigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	BaseValue        *int32  `protobuf:"varint,5,opt,name=baseValue" json:"baseValue,omitempty"`
	WinScore         *int64  `protobuf:"varint,6,opt,name=winScore" json:"winScore,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_EndLotteryInfo) Reset()                    { *m = Pez_EndLotteryInfo{} }
func (m *Pez_EndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_EndLotteryInfo) ProtoMessage()               {}
func (*Pez_EndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{4} }

func (m *Pez_EndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_EndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *Pez_EndLotteryInfo) GetIsbigWin() bool {
	if m != nil && m.IsbigWin != nil {
		return *m.IsbigWin
	}
	return false
}

func (m *Pez_EndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *Pez_EndLotteryInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *Pez_EndLotteryInfo) GetWinScore() int64 {
	if m != nil && m.WinScore != nil {
		return *m.WinScore
	}
	return 0
}

// 本局结果(广播)
type Pez_SendCurrentResult struct {
	CurrentRound     *int32          `protobuf:"varint,1,opt,name=currentRound" json:"currentRound,omitempty"`
	CurrBill         []*Pez_UserBean `protobuf:"bytes,2,rep,name=CurrBill" json:"CurrBill,omitempty"`
	BoardsCout       *int32          `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	CardsID          []int32         `protobuf:"varint,4,rep,name=cardsID" json:"cardsID,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_SendCurrentResult) Reset()                    { *m = Pez_SendCurrentResult{} }
func (m *Pez_SendCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendCurrentResult) ProtoMessage()               {}
func (*Pez_SendCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{5} }

func (m *Pez_SendCurrentResult) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *Pez_SendCurrentResult) GetCurrBill() []*Pez_UserBean {
	if m != nil {
		return m.CurrBill
	}
	return nil
}

func (m *Pez_SendCurrentResult) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *Pez_SendCurrentResult) GetCardsID() []int32 {
	if m != nil {
		return m.CardsID
	}
	return nil
}

type EndLottery struct {
	UserId           *uint32     `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	IsbigWin         *bool       `protobuf:"varint,2,opt,name=isbigWin" json:"isbigWin,omitempty"`
	IsOwner          *bool       `protobuf:"varint,3,opt,name=isOwner" json:"isOwner,omitempty"`
	WinScore         *int64      `protobuf:"varint,4,opt,name=winScore" json:"winScore,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,5,opt,name=wxInfo" json:"wxInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *EndLottery) Reset()                    { *m = EndLottery{} }
func (m *EndLottery) String() string            { return proto.CompactTextString(m) }
func (*EndLottery) ProtoMessage()               {}
func (*EndLottery) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{6} }

func (m *EndLottery) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *EndLottery) GetIsbigWin() bool {
	if m != nil && m.IsbigWin != nil {
		return *m.IsbigWin
	}
	return false
}

func (m *EndLottery) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *EndLottery) GetWinScore() int64 {
	if m != nil && m.WinScore != nil {
		return *m.WinScore
	}
	return 0
}

func (m *EndLottery) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

// 牌局结束(广播)
type Pez_SendEndLottery struct {
	EndLotteryInfo   []*EndLottery `protobuf:"bytes,1,rep,name=endLotteryInfo" json:"endLotteryInfo,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Pez_SendEndLottery) Reset()                    { *m = Pez_SendEndLottery{} }
func (m *Pez_SendEndLottery) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendEndLottery) ProtoMessage()               {}
func (*Pez_SendEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{7} }

func (m *Pez_SendEndLottery) GetEndLotteryInfo() []*EndLottery {
	if m != nil {
		return m.EndLotteryInfo
	}
	return nil
}

type Pez_UserBean struct {
	UserId           *uint32   `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int64    `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	Coin             *int64    `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Tuozi            *PezTuozi `protobuf:"bytes,4,opt,name=tuozi" json:"tuozi,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Pez_UserBean) Reset()                    { *m = Pez_UserBean{} }
func (m *Pez_UserBean) String() string            { return proto.CompactTextString(m) }
func (*Pez_UserBean) ProtoMessage()               {}
func (*Pez_UserBean) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{8} }

func (m *Pez_UserBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_UserBean) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Pez_UserBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *Pez_UserBean) GetTuozi() *PezTuozi {
	if m != nil {
		return m.Tuozi
	}
	return nil
}

type Pez_Bill struct {
	BeanList         []*Pez_UserBean `protobuf:"bytes,1,rep,name=BeanList" json:"BeanList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_Bill) Reset()                    { *m = Pez_Bill{} }
func (m *Pez_Bill) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bill) ProtoMessage()               {}
func (*Pez_Bill) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{9} }

func (m *Pez_Bill) GetBeanList() []*Pez_UserBean {
	if m != nil {
		return m.BeanList
	}
	return nil
}

func init() {
	proto.RegisterType((*Pez_DissolveDesk)(nil), "ddproto.pez_DissolveDesk")
	proto.RegisterType((*Pez_AckDissolveDesk)(nil), "ddproto.pez_AckDissolveDesk")
	proto.RegisterType((*Pez_Ready)(nil), "ddproto.pez_Ready")
	proto.RegisterType((*Pez_AckReady)(nil), "ddproto.pez_AckReady")
	proto.RegisterType((*Pez_EndLotteryInfo)(nil), "ddproto.pez_EndLotteryInfo")
	proto.RegisterType((*Pez_SendCurrentResult)(nil), "ddproto.pez_SendCurrentResult")
	proto.RegisterType((*EndLottery)(nil), "ddproto.EndLottery")
	proto.RegisterType((*Pez_SendEndLottery)(nil), "ddproto.pez_SendEndLottery")
	proto.RegisterType((*Pez_UserBean)(nil), "ddproto.pez_UserBean")
	proto.RegisterType((*Pez_Bill)(nil), "ddproto.pez_Bill")
}

var fileDescriptor37 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x51, 0x6b, 0xdb, 0x3c,
	0x14, 0xfd, 0x64, 0xc7, 0xa9, 0x73, 0xd3, 0xaf, 0x74, 0x6a, 0x3b, 0x4c, 0x18, 0xc3, 0xf8, 0xc9,
	0xb0, 0x51, 0x68, 0x5e, 0xc7, 0x1e, 0xd6, 0x76, 0xd0, 0x40, 0xd9, 0x52, 0x95, 0xad, 0x7b, 0x2b,
	0x8e, 0x7d, 0xdb, 0x89, 0x38, 0x52, 0xb0, 0xec, 0xa5, 0xcd, 0x9f, 0xd8, 0x8f, 0xd8, 0xcb, 0x7e,
	0xc2, 0x7e, 0xde, 0x90, 0xac, 0x38, 0xf1, 0x20, 0x7b, 0x18, 0x7d, 0x31, 0x3a, 0xe7, 0x5e, 0x9f,
	0xab, 0x73, 0x74, 0x61, 0x6f, 0x8e, 0xcb, 0xdb, 0x0c, 0xd5, 0xf4, 0x78, 0x5e, 0xc8, 0x52, 0xd2,
	0x9d, 0x2c, 0x33, 0x87, 0xc1, 0x41, 0x2a, 0x67, 0x33, 0x29, 0x6e, 0xd3, 0x9c, 0xa3, 0x28, 0xeb,
	0xea, 0xe0, 0x99, 0x25, 0x51, 0x54, 0x33, 0x4b, 0x19, 0x81, 0x49, 0xa2, 0xb0, 0xc6, 0xd1, 0x17,
	0xd8, 0xd7, 0xcc, 0x39, 0x57, 0x4a, 0xe6, 0xdf, 0xf0, 0x1c, 0xd5, 0x94, 0xbe, 0x86, 0xee, 0x57,
	0x4c, 0x32, 0x2c, 0x02, 0x12, 0x92, 0xb8, 0x3f, 0x3c, 0x3c, 0xb6, 0x53, 0x8e, 0xc7, 0xfa, 0x7b,
	0x61, 0x6a, 0xcc, 0xf6, 0xd0, 0xe7, 0xd0, 0xad, 0x14, 0x16, 0xa3, 0x2c, 0x70, 0x42, 0x12, 0xff,
	0xcf, 0x2c, 0x8a, 0xbe, 0x13, 0x38, 0xd0, 0xd2, 0xef, 0xd2, 0xe9, 0xd3, 0xab, 0x6b, 0x5e, 0xc7,
	0x30, 0xca, 0x02, 0x37, 0x24, 0xb1, 0xc7, 0x2c, 0xa2, 0x03, 0xf0, 0xe7, 0x89, 0x52, 0x37, 0xb2,
	0xc8, 0x82, 0x4e, 0x48, 0xe2, 0x1e, 0x6b, 0x70, 0x74, 0x05, 0x3d, 0x7d, 0x21, 0x86, 0x49, 0xf6,
	0xf8, 0x44, 0x26, 0xef, 0x60, 0xd7, 0x7a, 0xfc, 0x17, 0xd5, 0x7d, 0x70, 0x67, 0xea, 0xde, 0x48,
	0xf6, 0x98, 0x3e, 0x6e, 0xcc, 0x71, 0x5b, 0x73, 0x7e, 0x11, 0xa0, 0x7a, 0xd0, 0x7b, 0x91, 0x5d,
	0xca, 0xb2, 0xc4, 0xe2, 0x71, 0x24, 0xee, 0xe4, 0x46, 0x3b, 0x69, 0xa5, 0x33, 0x00, 0x5f, 0xf0,
	0x74, 0xfa, 0x21, 0x99, 0xa1, 0x55, 0x6f, 0xb0, 0xae, 0x71, 0x35, 0xe1, 0xf7, 0x37, 0x5c, 0x98,
	0x21, 0x3e, 0x6b, 0x30, 0x0d, 0x60, 0x87, 0xab, 0x8f, 0x0b, 0x81, 0x85, 0x09, 0xcf, 0x67, 0x2b,
	0x48, 0x5f, 0x40, 0x4f, 0x6f, 0xcd, 0xe7, 0x24, 0xaf, 0x30, 0xf0, 0x4c, 0xe4, 0x6b, 0x42, 0x6b,
	0x2e, 0xb8, 0xb8, 0x4e, 0x65, 0x81, 0x41, 0x37, 0x24, 0xb1, 0xcb, 0x1a, 0x1c, 0xfd, 0x24, 0x70,
	0xa4, 0xaf, 0x7e, 0x8d, 0x22, 0x3b, 0xab, 0x8a, 0x02, 0x45, 0xc9, 0x50, 0x55, 0x79, 0x49, 0x23,
	0xd8, 0x4d, 0x2d, 0x21, 0x2b, 0x51, 0x7b, 0xf0, 0x58, 0x8b, 0xa3, 0x27, 0xe0, 0xeb, 0x9f, 0x4e,
	0x79, 0x9e, 0x07, 0x4e, 0xe8, 0xc6, 0xfd, 0xe1, 0x51, 0x13, 0xa9, 0x56, 0xfd, 0xa4, 0xb0, 0x38,
	0xc5, 0x44, 0xb0, 0xa6, 0x8d, 0xbe, 0x04, 0x98, 0xc8, 0xa4, 0xc8, 0xd4, 0x99, 0xac, 0x4a, 0xbb,
	0x1e, 0x1b, 0x8c, 0x36, 0x99, 0x6a, 0x30, 0x3a, 0x0f, 0x3a, 0xa1, 0x1b, 0x7b, 0x6c, 0x05, 0xa3,
	0x1f, 0x04, 0x60, 0x9d, 0xf0, 0xdf, 0xd2, 0x6d, 0x12, 0x74, 0xb6, 0x27, 0xe8, 0xb6, 0x13, 0xdc,
	0xcc, 0xa8, 0xd3, 0xce, 0x88, 0xbe, 0x82, 0xee, 0xe2, 0x41, 0xbf, 0xa8, 0x89, 0xb6, 0x3f, 0x3c,
	0x68, 0x3c, 0xde, 0x20, 0x7f, 0xe0, 0x42, 0x97, 0x98, 0x6d, 0x89, 0xae, 0xea, 0x55, 0xd0, 0x79,
	0x6e, 0x5c, 0xf6, 0x0d, 0xec, 0x61, 0x6b, 0x39, 0x02, 0x62, 0xe2, 0x5a, 0x4b, 0xad, 0x9b, 0xd9,
	0x1f, 0xad, 0xd1, 0xb2, 0x5e, 0xe3, 0x55, 0x98, 0x5b, 0x9d, 0x1f, 0x82, 0xa7, 0x8c, 0x01, 0xc7,
	0x18, 0xa8, 0x01, 0xa5, 0xd0, 0x49, 0xa5, 0xdd, 0x26, 0x97, 0x99, 0x33, 0x8d, 0xc1, 0x2b, 0x2b,
	0xb9, 0xe4, 0xc6, 0x6a, 0x7f, 0x48, 0x5b, 0x8f, 0x66, 0x2a, 0xac, 0x6e, 0x88, 0xde, 0x82, 0xaf,
	0x39, 0xf3, 0x74, 0x27, 0xe0, 0xeb, 0xf9, 0x97, 0x5c, 0x95, 0xf6, 0xfa, 0xdb, 0x5e, 0x7b, 0xd5,
	0x76, 0xea, 0x5c, 0xb8, 0xe3, 0xff, 0xc6, 0x64, 0xec, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x6a, 0xce, 0x91, 0x1c, 0x05, 0x00, 0x00,
}
