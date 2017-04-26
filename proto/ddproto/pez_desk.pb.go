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

// 房主解散房间(未开局)
type Pez_DissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_DissolveDesk) Reset()                    { *m = Pez_DissolveDesk{} }
func (m *Pez_DissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_DissolveDesk) ProtoMessage()               {}
func (*Pez_DissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{0} }

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
func (*Pez_AckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{1} }

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

// 申请解散房间(游戏中)
type Pez_ApplyDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_ApplyDissolveDesk) Reset()                    { *m = Pez_ApplyDissolveDesk{} }
func (m *Pez_ApplyDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_ApplyDissolveDesk) ProtoMessage()               {}
func (*Pez_ApplyDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{2} }

func (m *Pez_ApplyDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_ApplyDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type Pez_AckApplyDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserIdAgree      []uint32     `protobuf:"varint,2,rep,name=userIdAgree" json:"userIdAgree,omitempty"`
	UserIdWait       []uint32     `protobuf:"varint,3,rep,name=userIdWait" json:"userIdWait,omitempty"`
	UserIdDisagree   []uint32     `protobuf:"varint,4,rep,name=userIdDisagree" json:"userIdDisagree,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckApplyDissolveDesk) Reset()                    { *m = Pez_AckApplyDissolveDesk{} }
func (m *Pez_AckApplyDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckApplyDissolveDesk) ProtoMessage()               {}
func (*Pez_AckApplyDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{3} }

func (m *Pez_AckApplyDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckApplyDissolveDesk) GetUserIdAgree() []uint32 {
	if m != nil {
		return m.UserIdAgree
	}
	return nil
}

func (m *Pez_AckApplyDissolveDesk) GetUserIdWait() []uint32 {
	if m != nil {
		return m.UserIdWait
	}
	return nil
}

func (m *Pez_AckApplyDissolveDesk) GetUserIdDisagree() []uint32 {
	if m != nil {
		return m.UserIdDisagree
	}
	return nil
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
func (*Pez_Ready) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{4} }

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
func (*Pez_AckReady) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{5} }

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

// 押注信息
type BetInfo struct {
	BetCount         *int64 `protobuf:"varint,1,opt,name=betCount" json:"betCount,omitempty"`
	ContinueBet      *bool  `protobuf:"varint,2,opt,name=continueBet" json:"continueBet,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BetInfo) Reset()                    { *m = BetInfo{} }
func (m *BetInfo) String() string            { return proto.CompactTextString(m) }
func (*BetInfo) ProtoMessage()               {}
func (*BetInfo) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{6} }

func (m *BetInfo) GetBetCount() int64 {
	if m != nil && m.BetCount != nil {
		return *m.BetCount
	}
	return 0
}

func (m *BetInfo) GetContinueBet() bool {
	if m != nil && m.ContinueBet != nil {
		return *m.ContinueBet
	}
	return false
}

// 赢牌信息：谁赢了多少
type Pez_WinScoreInfo struct {
	NickName         *string `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	WinScore         *int64  `protobuf:"varint,3,opt,name=winScore" json:"winScore,omitempty"`
	Coin             *int64  `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	CardTitle        *string `protobuf:"bytes,5,opt,name=cardTitle" json:"cardTitle,omitempty"`
	IsBanker         *bool   `protobuf:"varint,7,opt,name=isBanker" json:"isBanker,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_WinScoreInfo) Reset()                    { *m = Pez_WinScoreInfo{} }
func (m *Pez_WinScoreInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_WinScoreInfo) ProtoMessage()               {}
func (*Pez_WinScoreInfo) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{7} }

func (m *Pez_WinScoreInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *Pez_WinScoreInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_WinScoreInfo) GetWinScore() int64 {
	if m != nil && m.WinScore != nil {
		return *m.WinScore
	}
	return 0
}

func (m *Pez_WinScoreInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *Pez_WinScoreInfo) GetCardTitle() string {
	if m != nil && m.CardTitle != nil {
		return *m.CardTitle
	}
	return ""
}

func (m *Pez_WinScoreInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
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
func (*Pez_EndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{8} }

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
	CurrBill         []*Pez_UserBean `protobuf:"bytes,1,rep,name=CurrBill" json:"CurrBill,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_SendCurrentResult) Reset()                    { *m = Pez_SendCurrentResult{} }
func (m *Pez_SendCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendCurrentResult) ProtoMessage()               {}
func (*Pez_SendCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{9} }

func (m *Pez_SendCurrentResult) GetCurrBill() []*Pez_UserBean {
	if m != nil {
		return m.CurrBill
	}
	return nil
}

// 牌局结束(广播)
type Pez_SendEndLottery struct {
	AllBill          []*Pez_Bill `protobuf:"bytes,1,rep,name=allBill" json:"allBill,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Pez_SendEndLottery) Reset()                    { *m = Pez_SendEndLottery{} }
func (m *Pez_SendEndLottery) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendEndLottery) ProtoMessage()               {}
func (*Pez_SendEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{10} }

func (m *Pez_SendEndLottery) GetAllBill() []*Pez_Bill {
	if m != nil {
		return m.AllBill
	}
	return nil
}

type Pez_UserBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int64  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_UserBean) Reset()                    { *m = Pez_UserBean{} }
func (m *Pez_UserBean) String() string            { return proto.CompactTextString(m) }
func (*Pez_UserBean) ProtoMessage()               {}
func (*Pez_UserBean) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{11} }

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

type Pez_Bill struct {
	BeanList         []*Pez_UserBean `protobuf:"bytes,1,rep,name=BeanList" json:"BeanList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_Bill) Reset()                    { *m = Pez_Bill{} }
func (m *Pez_Bill) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bill) ProtoMessage()               {}
func (*Pez_Bill) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{12} }

func (m *Pez_Bill) GetBeanList() []*Pez_UserBean {
	if m != nil {
		return m.BeanList
	}
	return nil
}

func init() {
	proto.RegisterType((*Pez_DissolveDesk)(nil), "ddproto.pez_DissolveDesk")
	proto.RegisterType((*Pez_AckDissolveDesk)(nil), "ddproto.pez_AckDissolveDesk")
	proto.RegisterType((*Pez_ApplyDissolveDesk)(nil), "ddproto.pez_ApplyDissolveDesk")
	proto.RegisterType((*Pez_AckApplyDissolveDesk)(nil), "ddproto.pez_AckApplyDissolveDesk")
	proto.RegisterType((*Pez_Ready)(nil), "ddproto.pez_Ready")
	proto.RegisterType((*Pez_AckReady)(nil), "ddproto.pez_AckReady")
	proto.RegisterType((*BetInfo)(nil), "ddproto.BetInfo")
	proto.RegisterType((*Pez_WinScoreInfo)(nil), "ddproto.pez_WinScoreInfo")
	proto.RegisterType((*Pez_EndLotteryInfo)(nil), "ddproto.pez_EndLotteryInfo")
	proto.RegisterType((*Pez_SendCurrentResult)(nil), "ddproto.pez_SendCurrentResult")
	proto.RegisterType((*Pez_SendEndLottery)(nil), "ddproto.pez_SendEndLottery")
	proto.RegisterType((*Pez_UserBean)(nil), "ddproto.pez_UserBean")
	proto.RegisterType((*Pez_Bill)(nil), "ddproto.pez_Bill")
}

var fileDescriptor35 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xd1, 0x4e, 0x13, 0x41,
	0x14, 0x86, 0x9d, 0x2e, 0xa5, 0xed, 0x41, 0x08, 0x0c, 0x60, 0x36, 0x8d, 0x31, 0x9b, 0xbd, 0x30,
	0x4d, 0x34, 0x24, 0x72, 0xad, 0x17, 0x14, 0x8c, 0x60, 0x88, 0xe2, 0xa0, 0xd6, 0x1b, 0x43, 0x86,
	0xdd, 0x03, 0x4e, 0xba, 0x9d, 0x69, 0x66, 0x66, 0x25, 0xf8, 0x12, 0xbe, 0x8a, 0xf1, 0xca, 0xc7,
	0x33, 0x33, 0xb3, 0xbb, 0x6d, 0x49, 0x88, 0x89, 0xe2, 0x0d, 0x99, 0xff, 0x9f, 0xc3, 0x77, 0xfe,
	0x73, 0x3a, 0x0b, 0x6b, 0x53, 0xfc, 0x76, 0x96, 0xa3, 0x19, 0xef, 0x4c, 0xb5, 0xb2, 0x8a, 0x76,
	0xf2, 0xdc, 0x1f, 0xfa, 0x9b, 0x99, 0x9a, 0x4c, 0x94, 0x3c, 0xcb, 0x0a, 0x81, 0xd2, 0x86, 0xdb,
	0xfe, 0x46, 0x65, 0xa2, 0x2c, 0x27, 0xc1, 0x4a, 0x3f, 0xc1, 0xba, 0x43, 0x1c, 0x08, 0x63, 0x54,
	0xf1, 0x15, 0x0f, 0xd0, 0x8c, 0xe9, 0x53, 0x58, 0xfe, 0x82, 0x3c, 0x47, 0x1d, 0x93, 0x84, 0x0c,
	0x56, 0x76, 0xb7, 0x76, 0x2a, 0xea, 0xce, 0x89, 0xfb, 0x7b, 0xe8, 0xef, 0x58, 0x55, 0x43, 0x1f,
	0xc0, 0x72, 0x69, 0x50, 0x1f, 0xe5, 0x71, 0x2b, 0x21, 0x83, 0x55, 0x56, 0xa9, 0xf4, 0x3b, 0x81,
	0x4d, 0x87, 0xde, 0xcb, 0xc6, 0x77, 0x4f, 0x77, 0xbe, 0x1b, 0xfb, 0x28, 0x8f, 0xa3, 0x84, 0x0c,
	0xda, 0xac, 0x52, 0xb4, 0x0f, 0xdd, 0x29, 0x37, 0x66, 0xa4, 0x74, 0x1e, 0x2f, 0x25, 0x64, 0xd0,
	0x63, 0x8d, 0x4e, 0x3f, 0xc3, 0xb6, 0x0f, 0x34, 0x9d, 0x16, 0xd7, 0xff, 0x61, 0xe0, 0x9f, 0x04,
	0xe2, 0x6a, 0xe0, 0x7f, 0x6d, 0x91, 0xc0, 0x4a, 0x80, 0xee, 0x5d, 0x6a, 0xc4, 0xb8, 0x95, 0x44,
	0x83, 0x55, 0x36, 0x6f, 0xd1, 0x47, 0x00, 0x41, 0x8e, 0xb8, 0xb0, 0x71, 0xe4, 0x0b, 0xe6, 0x1c,
	0xfa, 0x18, 0xd6, 0x82, 0x3a, 0x10, 0x86, 0x7b, 0xc8, 0x92, 0xaf, 0xb9, 0xe1, 0xa6, 0xef, 0xa0,
	0xe7, 0x32, 0x33, 0xe4, 0xf9, 0xf5, 0x1d, 0xed, 0xe1, 0x02, 0xee, 0x57, 0x6b, 0xf8, 0x1b, 0xea,
	0x3a, 0x44, 0x13, 0x73, 0xe9, 0x91, 0x3d, 0xe6, 0x8e, 0x73, 0x7d, 0xa2, 0x85, 0x3e, 0xaf, 0xa0,
	0x33, 0x44, 0x7b, 0x24, 0x2f, 0x94, 0xfb, 0xd5, 0xcf, 0xd1, 0xee, 0xab, 0x52, 0x5a, 0xdf, 0x24,
	0x62, 0x8d, 0x76, 0xbb, 0xcc, 0x94, 0xb4, 0x42, 0x96, 0x38, 0x44, 0xeb, 0xc1, 0x5d, 0x36, 0x6f,
	0xa5, 0x3f, 0x48, 0xf8, 0x08, 0x46, 0x42, 0x9e, 0x66, 0x4a, 0x63, 0x8d, 0x94, 0x22, 0x1b, 0xbf,
	0xe1, 0x13, 0xf4, 0xc8, 0x1e, 0x6b, 0xf4, 0xad, 0x8f, 0xb2, 0x0f, 0xdd, 0xab, 0x8a, 0xe1, 0xb3,
	0x46, 0xac, 0xd1, 0x94, 0xc2, 0x52, 0xa6, 0x84, 0xf4, 0x8f, 0x32, 0x62, 0xfe, 0x4c, 0x1f, 0x42,
	0x2f, 0xe3, 0x3a, 0x7f, 0x2f, 0x6c, 0x81, 0x71, 0xdb, 0x37, 0x99, 0x19, 0x8e, 0x26, 0xcc, 0x90,
	0xcb, 0x31, 0xea, 0xb8, 0xe3, 0x53, 0x37, 0x3a, 0xfd, 0x45, 0x80, 0xba, 0xc8, 0x2f, 0x65, 0x7e,
	0xac, 0xac, 0x45, 0x7d, 0xed, 0x43, 0xcf, 0x82, 0x91, 0x9b, 0xc1, 0x9a, 0x61, 0x5a, 0x37, 0x86,
	0xf1, 0x6d, 0xce, 0xc5, 0xe5, 0x48, 0x48, 0x1f, 0xda, 0xb7, 0x09, 0x9a, 0xc6, 0xd0, 0x11, 0xe6,
	0xed, 0x95, 0x44, 0xed, 0x73, 0x77, 0x59, 0x2d, 0x5d, 0xf4, 0x73, 0x6e, 0xf0, 0x23, 0x2f, 0xca,
	0x10, 0xbd, 0xcd, 0x66, 0xc6, 0xc2, 0x22, 0x96, 0x17, 0x17, 0x91, 0xbe, 0x0e, 0x5f, 0xe1, 0x29,
	0xca, 0x7c, 0xbf, 0xd4, 0x1a, 0xa5, 0x65, 0x68, 0xca, 0xc2, 0xd2, 0x67, 0xd0, 0x75, 0xc6, 0x50,
	0x14, 0x45, 0x4c, 0x92, 0x68, 0xb0, 0xb2, 0xbb, 0xdd, 0xbc, 0x14, 0xf7, 0x1f, 0x1f, 0x0c, 0xea,
	0x21, 0x72, 0xc9, 0x9a, 0xb2, 0x74, 0x2f, 0x6c, 0xc1, 0xb1, 0x66, 0x9b, 0xa0, 0x4f, 0xa0, 0xc3,
	0x8b, 0x62, 0x8e, 0xb3, 0xb1, 0xc0, 0x71, 0x17, 0xac, 0xae, 0x48, 0x9f, 0x87, 0xd7, 0x5a, 0xc3,
	0x6f, 0x5d, 0xe1, 0x16, 0xb4, 0x8d, 0x9f, 0xa7, 0xe5, 0xe7, 0x09, 0x22, 0x7d, 0x01, 0xdd, 0x1a,
	0xe9, 0xf2, 0x3b, 0xc2, 0xb1, 0x30, 0xf6, 0x0f, 0xf9, 0xeb, 0xb2, 0x61, 0xeb, 0x30, 0x3a, 0xb9,
	0x77, 0x42, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x59, 0x03, 0x50, 0x99, 0xc7, 0x05, 0x00, 0x00,
}
