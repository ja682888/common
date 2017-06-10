// Code generated by protoc-gen-go.
// source: common_client.proto
// DO NOT EDIT!

/*
Package yjprotogo is a generated protocol buffer package.

It is generated from these files:
	common_client.proto
	common_enum.proto
	common_server.proto
	hall_base.proto
	hall_common_server_award.proto
	hall_common_server_model.proto
	mj_changchun_base.proto
	mj_changchun_play.proto
	sdy_base.proto
	sdy_desk.proto
	sdy_hall.proto
	sdy_play.proto
	tdk.proto
	tdk_base.proto
	tdk_data.proto

It has these top-level messages:
	ProtoHeader
	ServerInfo
	QuickConn
	AckQuickConn
	WeixinInfo
	CommonReqReg
	CommonAckReg
	CmOffline
	CmHearbeat
	GameSession
	AwardReqGetNewUser
	AwardAckGetNewUser
	HallReqServer
	HallAckServer
	Hallserver
	HallReqFeedBack
	HallAckFeedBack
	HallReqNotice
	HallAckNotice
	HallNotice
	HallItemEvent
	HallMailItem
	HallItemTask
	HallReqUserData
	HallAckUserData
	HallUserData
	HallRankItem
	CoinZone
	DiamondZone
	ExchangeZone
	BuyZone
	GoodsItem
	HallStrongboxInfo
	HallAckStrongboxInfo
	HallReqStrongboxInfo
	HallReqStrongboxAccess
	HallAckStrongboxAccess
	CommonReqGameLogin
	CommonReqGameLoginViaInput
	CommonAckGameLogin
	HallReqDrawLottery
	HallAckDrawLottery
	HallReqDsLotteryInfo
	HallAckDsLotteryInfo
	HallReqBagItems
	HallAckBagItems
	HallReqGoodsList
	HallAckGoodsList
	HallReqGoodsBuy
	HallAckGoodsBuy
	HallGoodsItemMsg
	HallBagItem
	HallSignLotteryInfo
	HallDrawLotteryInfo
	HallLotteryItem
	HallReqNewRoomList
	HallAckNewRoomList
	HallNewRoom_List
	Taward
	AwardMOnline
	User
	TNotice
	TGameCounts
	TUserTask
	Push
	UserCoinBean
	RoomTypeInfo
	EndLotteryInfo
	PlayOptions
	DeskGameInfo
	PlayerInfo
	CardInfo
	BuCardInfo
	ComposeCard
	PlayerCard
	WinCoinInfo
	Game_SendGameInfo
	Game_AckActHu
	P16AckGameOpening
	Game_SendCurrentResult
	P16AckRoomInit
	P16BeanDeskInfo
	P16ReqCreateDesk
	P16AckCreateDesk
	P16ReqEnterDesk
	P16AckEnterDesk
	P16ReqDissolveBeginGame
	P16AckDissolveBeginGame
	P16DissolveBeginGame
	P16AckPlayerEnter
	P16AckPlayerExit
	P16ReqReady
	P16AckReady
	P16AckOpening
	P16BcMoPai
	P16AckGameOverturn
	P16AckGameDealCards
	GangOverTurn
	BuGangOverTurn
	ChiOverTurn
	JiaoInfo
	JiaoPaiInfo
	P16AckGameSendOutCard
	Game_SendEndLottery
	P16ReqGameSendOutCard
	P16ReqChi
	P16AckChi
	P16ReqGameActPeng
	P16AckGameActPeng
	P16ReqGameActGang
	P16AckGameActGang
	P16ReqGameActBugang
	P16AckGameActBugang
	P16ReqGameActGuo
	P16AckGameActGuo
	P16ReqGameActHu
	P16ReqGameTing
	P16AckGameTing
	P16AckGameDabao
	P16AckGameChangbao
	P16ReqRemainMajiang
	P16AckRemainMajiang
	P16ReqChangeMajiang
	P16AckChangeMajiang
	SdyBaseUserPaiIds
	SdyBaseRoomTypeInfo
	SdyBaseTimerInfo
	SdyBasePlayerInfo
	SdyBaseCommonRateInfo
	SdyBaseDeskInfo
	SdyReqReady
	SdyAckReady
	SdyBaseWinCoinInfo
	SdyBcCurrentResult
	SdyBaseBill
	SdyBcEndLotteryInfo
	SdyReUserOutPai
	SdyReReady
	SdyReHuanDi
	SdyRePlay
	SdyReJiaoFen
	SdyReLenHandPokers
	SdyBcReconnectInfo
	SdyBcIsOnLine
	SdyReqCreateDesk
	SdyAckCreateDesk
	SdyReqEnterDesk
	SdyAckEnterDesk
	SdyBcUserBreak
	SdyReqDissolveDeskOwner
	SdyBcDissolveDeskOwner
	SdyReqApplyDissolveDeskApllyer
	SdyBcApplyDissolveDeskApplyer
	SdyReqApplyDissolveDeskUser
	SdyAckApplyDissolveDeskUser
	SdyReqLeaveDesk
	SdyAckLeaveDesk
	SdyReqSendMessage
	SdyAckSendMessage
	SdyBcOpening
	SdyDealCards
	SdyBcPlayerPokers
	SdyBcJiaoFen
	SdyReqJiaoFen
	SdyAckJiaoFen
	SdyBcJiaoFenResult
	SdyBcStartPlay
	SdyReqOutCards
	SdyAckOutCards
	SdyBcOverTurn
	SdyBcWhatPai
	SdyBcScorePai
	SdyBcGameInfo
	SdyBcDingZhu
	SdyReqDingZhu
	SdyBcDingZhuResult
	SdyAckDingZhu
	SdyBcHuanDi
	SdyReqHuanDi
	SdyAckHuanDi
	TdkHeartBeat
	TdkJoinRoom
	TdkJoinRoomRsp
	TdkLeaveRoom
	TdkLeaveRoomRsp
	TdkJoinDesk
	TdkJoinDeskRsp
	TdkCreateDesk
	TdkCreateDeskRsp
	TdkLeaveDesk
	TdkLeaveDeskRsp
	TdkUserReady
	TdkUserReadyRsp
	TdkReturnDesk
	TdkDissolveDeskReq
	TdkDissolveDeskUser
	TdkDisDeskDesicionReq
	TdkDisDeskDesicionRsp
	TdkDisDeskReqResult
	TdkDisDesk
	TdkSendPoker
	TdkStartBet
	TdkBet
	TdkBetRsp
	TdkFold
	TdkFoldRsp
	TdkQiJiao
	TdkQiJiaoRsp
	TdkFanTi
	TdkFanTiRsp
	TdkPass
	TdkPassRsp
	TdkKaiPai
	TdkKaiPaiRsp
	TdkRoundEnd
	TdkZhanJi
	TdkZhanJiRsp
	TdkMatchReq
	TdkMatchRsp
	TdkMatchAddUserRsp
	TdkJoinPlayingDeskRsp
	TdkEnterGame
	TdkEnterGameRsp
	TdkDeskUserData
	TdkUserPokerData
	TdkDisUserData
	TdkDeskConfig
	TdkZhanJiData
*/
package yjprotogo

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

// 注册类型
type CommonEnumReg int32

const (
	CommonEnumReg_RET_TYPE_TOURIST CommonEnumReg = 1
	CommonEnumReg_RET_TYPE_WEIXIN  CommonEnumReg = 2
)

var CommonEnumReg_name = map[int32]string{
	1: "RET_TYPE_TOURIST",
	2: "RET_TYPE_WEIXIN",
}
var CommonEnumReg_value = map[string]int32{
	"RET_TYPE_TOURIST": 1,
	"RET_TYPE_WEIXIN":  2,
}

func (x CommonEnumReg) Enum() *CommonEnumReg {
	p := new(CommonEnumReg)
	*p = x
	return p
}
func (x CommonEnumReg) String() string {
	return proto.EnumName(CommonEnumReg_name, int32(x))
}
func (x *CommonEnumReg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumReg_value, data, "CommonEnumReg")
	if err != nil {
		return err
	}
	*x = CommonEnumReg(value)
	return nil
}
func (CommonEnumReg) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 客户端系统类型
type CommonEnumOsType int32

const (
	CommonEnumOsType_OS_IOS     CommonEnumOsType = 1
	CommonEnumOsType_OS_ANDROID CommonEnumOsType = 2
	CommonEnumOsType_OS_WEB     CommonEnumOsType = 3
)

var CommonEnumOsType_name = map[int32]string{
	1: "OS_IOS",
	2: "OS_ANDROID",
	3: "OS_WEB",
}
var CommonEnumOsType_value = map[string]int32{
	"OS_IOS":     1,
	"OS_ANDROID": 2,
	"OS_WEB":     3,
}

func (x CommonEnumOsType) Enum() *CommonEnumOsType {
	p := new(CommonEnumOsType)
	*p = x
	return p
}
func (x CommonEnumOsType) String() string {
	return proto.EnumName(CommonEnumOsType_name, int32(x))
}
func (x *CommonEnumOsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumOsType_value, data, "CommonEnumOsType")
	if err != nil {
		return err
	}
	*x = CommonEnumOsType(value)
	return nil
}
func (CommonEnumOsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// ProtoHeader 需要在每个 Message 中作为第一个字段
type ProtoHeader struct {
	Version          *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Code             *int32  `protobuf:"varint,3,opt,name=code" json:"code,omitempty"`
	Error            *string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoHeader) Reset()                    { *m = ProtoHeader{} }
func (m *ProtoHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtoHeader) ProtoMessage()               {}
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoHeader) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ProtoHeader) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ProtoHeader) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ProtoHeader) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// 服务器信息
type ServerInfo struct {
	Ip               *string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port             *int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	ReleaseTag       *int32  `protobuf:"varint,3,opt,name=releaseTag" json:"releaseTag,omitempty"`
	CurrVersion      *int32  `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	IsUpdate         *int32  `protobuf:"varint,5,opt,name=isUpdate" json:"isUpdate,omitempty"`
	DownloadUrl      *string `protobuf:"bytes,6,opt,name=downloadUrl" json:"downloadUrl,omitempty"`
	VersionInfo      *string `protobuf:"bytes,7,opt,name=versionInfo" json:"versionInfo,omitempty"`
	IsMaintain       *int32  `protobuf:"varint,8,opt,name=isMaintain" json:"isMaintain,omitempty"`
	MaintainMsg      *string `protobuf:"bytes,9,opt,name=maintainMsg" json:"maintainMsg,omitempty"`
	Status           *int32  `protobuf:"varint,10,opt,name=status" json:"status,omitempty"`
	GameId           *int32  `protobuf:"varint,11,opt,name=gameId" json:"gameId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ServerInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ServerInfo) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

func (m *ServerInfo) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *ServerInfo) GetIsUpdate() int32 {
	if m != nil && m.IsUpdate != nil {
		return *m.IsUpdate
	}
	return 0
}

func (m *ServerInfo) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *ServerInfo) GetVersionInfo() string {
	if m != nil && m.VersionInfo != nil {
		return *m.VersionInfo
	}
	return ""
}

func (m *ServerInfo) GetIsMaintain() int32 {
	if m != nil && m.IsMaintain != nil {
		return *m.IsMaintain
	}
	return 0
}

func (m *ServerInfo) GetMaintainMsg() string {
	if m != nil && m.MaintainMsg != nil {
		return *m.MaintainMsg
	}
	return ""
}

func (m *ServerInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *ServerInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

// 接入服务器
type QuickConn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ChannelId        *string      `protobuf:"bytes,2,opt,name=channelId" json:"channelId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	CurrVersion      *int32       `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	LanguageId       *int32       `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	UserId           *uint32      `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *QuickConn) Reset()                    { *m = QuickConn{} }
func (m *QuickConn) String() string            { return proto.CompactTextString(m) }
func (*QuickConn) ProtoMessage()               {}
func (*QuickConn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *QuickConn) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

func (m *QuickConn) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *QuickConn) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *QuickConn) GetLanguageId() int32 {
	if m != nil && m.LanguageId != nil {
		return *m.LanguageId
	}
	return 0
}

func (m *QuickConn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type AckQuickConn struct {
	Header            *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GameServer        []*ServerInfo `protobuf:"bytes,2,rep,name=gameServer" json:"gameServer,omitempty"`
	ServerListVersion *int32        `protobuf:"varint,3,opt,name=serverListVersion" json:"serverListVersion,omitempty"`
	XXX_unrecognized  []byte        `json:"-"`
}

func (m *AckQuickConn) Reset()                    { *m = AckQuickConn{} }
func (m *AckQuickConn) String() string            { return proto.CompactTextString(m) }
func (*AckQuickConn) ProtoMessage()               {}
func (*AckQuickConn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AckQuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AckQuickConn) GetGameServer() []*ServerInfo {
	if m != nil {
		return m.GameServer
	}
	return nil
}

func (m *AckQuickConn) GetServerListVersion() int32 {
	if m != nil && m.ServerListVersion != nil {
		return *m.ServerListVersion
	}
	return 0
}

// 微信信息
type WeixinInfo struct {
	OpenId           *string `protobuf:"bytes,1,opt,name=openId" json:"openId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	HeadUrl          *string `protobuf:"bytes,3,opt,name=headUrl" json:"headUrl,omitempty"`
	Sex              *int32  `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	City             *string `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	UnionId          *string `protobuf:"bytes,6,opt,name=unionId" json:"unionId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WeixinInfo) Reset()                    { *m = WeixinInfo{} }
func (m *WeixinInfo) String() string            { return proto.CompactTextString(m) }
func (*WeixinInfo) ProtoMessage()               {}
func (*WeixinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WeixinInfo) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *WeixinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WeixinInfo) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *WeixinInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *WeixinInfo) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *WeixinInfo) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

// 请求
type CommonReqReg struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RegType          *int32       `protobuf:"varint,2,opt,name=regType" json:"regType,omitempty"`
	WxInfo           *WeixinInfo  `protobuf:"bytes,3,opt,name=wxInfo" json:"wxInfo,omitempty"`
	ClientOSType     *int32       `protobuf:"varint,4,opt,name=clientOSType" json:"clientOSType,omitempty"`
	ChannelId        *string      `protobuf:"bytes,5,opt,name=channelId" json:"channelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonReqReg) Reset()                    { *m = CommonReqReg{} }
func (m *CommonReqReg) String() string            { return proto.CompactTextString(m) }
func (*CommonReqReg) ProtoMessage()               {}
func (*CommonReqReg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CommonReqReg) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonReqReg) GetRegType() int32 {
	if m != nil && m.RegType != nil {
		return *m.RegType
	}
	return 0
}

func (m *CommonReqReg) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *CommonReqReg) GetClientOSType() int32 {
	if m != nil && m.ClientOSType != nil {
		return *m.ClientOSType
	}
	return 0
}

func (m *CommonReqReg) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

// 回复
type CommonAckReg struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonAckReg) Reset()                    { *m = CommonAckReg{} }
func (m *CommonAckReg) String() string            { return proto.CompactTextString(m) }
func (*CommonAckReg) ProtoMessage()               {}
func (*CommonAckReg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CommonAckReg) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonAckReg) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 表示玩家掉线
type CmOffline struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmOffline) Reset()                    { *m = CmOffline{} }
func (m *CmOffline) String() string            { return proto.CompactTextString(m) }
func (*CmOffline) ProtoMessage()               {}
func (*CmOffline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 表示心跳
type CmHearbeat struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CmHearbeat) Reset()                    { *m = CmHearbeat{} }
func (m *CmHearbeat) String() string            { return proto.CompactTextString(m) }
func (*CmHearbeat) ProtoMessage()               {}
func (*CmHearbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CmHearbeat) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "yjprotogo.ProtoHeader")
	proto.RegisterType((*ServerInfo)(nil), "yjprotogo.ServerInfo")
	proto.RegisterType((*QuickConn)(nil), "yjprotogo.QuickConn")
	proto.RegisterType((*AckQuickConn)(nil), "yjprotogo.AckQuickConn")
	proto.RegisterType((*WeixinInfo)(nil), "yjprotogo.WeixinInfo")
	proto.RegisterType((*CommonReqReg)(nil), "yjprotogo.common_req_reg")
	proto.RegisterType((*CommonAckReg)(nil), "yjprotogo.common_ack_reg")
	proto.RegisterType((*CmOffline)(nil), "yjprotogo.cm_offline")
	proto.RegisterType((*CmHearbeat)(nil), "yjprotogo.cm_hearbeat")
	proto.RegisterEnum("yjprotogo.CommonEnumReg", CommonEnumReg_name, CommonEnumReg_value)
	proto.RegisterEnum("yjprotogo.CommonEnumOsType", CommonEnumOsType_name, CommonEnumOsType_value)
}

var fileDescriptor0 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xda, 0xb5, 0x5b, 0x5e, 0xc7, 0x16, 0xbc, 0x31, 0x59, 0x08, 0x4d, 0x55, 0x4e, 0xd5,
	0x04, 0x3d, 0x4c, 0xe2, 0xc6, 0x0e, 0x83, 0x55, 0x22, 0x12, 0x5b, 0x47, 0xda, 0xb1, 0x71, 0x8a,
	0x4c, 0xe2, 0x65, 0xa6, 0x89, 0x1d, 0x9c, 0x64, 0x3f, 0xfe, 0x15, 0xae, 0xfc, 0x2b, 0xdc, 0xf8,
	0xa7, 0x90, 0x1d, 0xb7, 0xf5, 0x40, 0x48, 0x6c, 0xb7, 0xf7, 0x7d, 0x79, 0xef, 0xf9, 0xbd, 0xcf,
	0x9f, 0x03, 0x5b, 0xb1, 0xc8, 0x73, 0xc1, 0xa3, 0x38, 0x63, 0x94, 0x57, 0xc3, 0x42, 0x8a, 0x4a,
	0x20, 0xf7, 0xee, 0xab, 0x0e, 0x52, 0xe1, 0x33, 0xe8, 0x9d, 0xaa, 0xf0, 0x3d, 0x25, 0x09, 0x95,
	0x08, 0xc3, 0xea, 0x35, 0x95, 0x25, 0x13, 0x1c, 0x3b, 0x7d, 0x67, 0xe0, 0x86, 0x73, 0x88, 0x76,
	0xa0, 0x5b, 0x97, 0x54, 0x06, 0x09, 0x6e, 0xf5, 0x9d, 0xc1, 0x93, 0xd0, 0x20, 0x84, 0x60, 0x25,
	0x16, 0x09, 0xc5, 0xed, 0xbe, 0x33, 0xe8, 0x84, 0x3a, 0x46, 0xdb, 0xd0, 0xa1, 0x52, 0x0a, 0x89,
	0x57, 0x74, 0x8f, 0x06, 0xf8, 0x3f, 0x5b, 0x00, 0x13, 0x2a, 0xaf, 0xa9, 0x0c, 0xf8, 0xa5, 0x40,
	0x1b, 0xd0, 0x62, 0x85, 0x39, 0xa5, 0xc5, 0x0a, 0xd5, 0xa8, 0x10, 0xb2, 0xd2, 0xed, 0x3b, 0xa1,
	0x8e, 0xd1, 0x2e, 0x80, 0xa4, 0x19, 0x25, 0x25, 0x9d, 0x92, 0xd4, 0x1c, 0x61, 0x31, 0xa8, 0x0f,
	0xbd, 0xb8, 0x96, 0xf2, 0x93, 0x19, 0x79, 0x45, 0x27, 0xd8, 0x14, 0x7a, 0x0e, 0x6b, 0xac, 0x3c,
	0x2b, 0x12, 0x52, 0x51, 0xdc, 0xd1, 0x9f, 0x17, 0x58, 0x55, 0x27, 0xe2, 0x86, 0x67, 0x82, 0x24,
	0x67, 0x32, 0xc3, 0x5d, 0x3d, 0x8a, 0x4d, 0xa9, 0x0c, 0xb3, 0xbf, 0x1a, 0x19, 0xaf, 0x36, 0x19,
	0x16, 0xa5, 0x26, 0x64, 0xe5, 0x31, 0x61, 0xbc, 0x22, 0x8c, 0xe3, 0xb5, 0x66, 0xc2, 0x25, 0xa3,
	0x3a, 0xe4, 0x26, 0x3e, 0x2e, 0x53, 0xec, 0x36, 0x1d, 0x2c, 0x4a, 0x09, 0x5b, 0x56, 0xa4, 0xaa,
	0x4b, 0x0c, 0xba, 0xda, 0x20, 0xc5, 0xa7, 0x24, 0xa7, 0x41, 0x82, 0x7b, 0x0d, 0xdf, 0x20, 0xff,
	0x97, 0x03, 0xee, 0xc7, 0x9a, 0xc5, 0xb3, 0x77, 0x82, 0x73, 0x34, 0x84, 0xee, 0x95, 0xbe, 0x3a,
	0xad, 0x64, 0x6f, 0x7f, 0x67, 0xb8, 0xb8, 0xdb, 0xa1, 0x75, 0xb1, 0xa1, 0xc9, 0x42, 0x2f, 0xc0,
	0x8d, 0xaf, 0x08, 0xe7, 0x34, 0x33, 0x37, 0xe9, 0x86, 0x4b, 0xc2, 0x3a, 0xb3, 0x6d, 0x9f, 0xf9,
	0x1f, 0x3a, 0xef, 0x02, 0x64, 0x84, 0xa7, 0x35, 0x49, 0x55, 0x75, 0xa3, 0xb4, 0xc5, 0x58, 0xf6,
	0xe9, 0xda, 0xf6, 0xf1, 0x7f, 0x38, 0xb0, 0x7e, 0x18, 0xcf, 0x1e, 0xbf, 0xd0, 0x6b, 0x00, 0x35,
	0x64, 0x63, 0x2c, 0xdc, 0xea, 0xb7, 0x07, 0xbd, 0xfd, 0x67, 0x56, 0xcd, 0xd2, 0x71, 0xa1, 0x95,
	0x88, 0x5e, 0xc2, 0xd3, 0x52, 0x47, 0x1f, 0x58, 0x59, 0xcd, 0xf7, 0x6a, 0x96, 0xfe, 0xfb, 0x83,
	0xff, 0xdd, 0x01, 0x38, 0xa7, 0xec, 0x96, 0x35, 0x97, 0xbe, 0x03, 0x5d, 0x51, 0x50, 0x1e, 0x24,
	0xc6, 0xbe, 0x06, 0x29, 0xb3, 0x71, 0x16, 0xcf, 0x4e, 0x48, 0x4e, 0x8d, 0xb6, 0x0b, 0xac, 0x5e,
	0x96, 0x9a, 0x58, 0x19, 0xad, 0xdd, 0xbc, 0x2c, 0x03, 0x91, 0x07, 0xed, 0x92, 0xde, 0x1a, 0x51,
	0x55, 0xa8, 0xdf, 0x14, 0xab, 0xee, 0xb4, 0x8c, 0x6e, 0xa8, 0x63, 0x55, 0x5f, 0x73, 0xe5, 0xba,
	0xc4, 0x18, 0x75, 0x0e, 0x95, 0x21, 0x36, 0xcc, 0x2b, 0x97, 0xf4, 0x5b, 0x24, 0x69, 0xfa, 0x60,
	0x11, 0x31, 0xac, 0x4a, 0x9a, 0x4e, 0xef, 0x0a, 0x6a, 0x9e, 0xdf, 0x1c, 0xa2, 0x57, 0xd0, 0xbd,
	0xb9, 0xd5, 0xe6, 0x6f, 0xeb, 0x4e, 0xb6, 0xb4, 0x4b, 0x45, 0x42, 0x93, 0x84, 0x7c, 0x58, 0x6f,
	0xfe, 0x34, 0xe3, 0x89, 0xee, 0xd6, 0x2c, 0x75, 0x8f, 0xbb, 0x6f, 0xc1, 0xce, 0x1f, 0x16, 0xf4,
	0x2f, 0x16, 0xcb, 0x90, 0x78, 0xf6, 0xa8, 0x65, 0xfe, 0xf1, 0xa7, 0xf2, 0xd7, 0x01, 0xe2, 0x3c,
	0x12, 0x97, 0x97, 0x19, 0xe3, 0xd4, 0x3f, 0x80, 0x5e, 0x9c, 0x47, 0x57, 0x94, 0xc8, 0x2f, 0x94,
	0x54, 0x0f, 0x3d, 0x64, 0xef, 0x0d, 0x6c, 0x9a, 0x31, 0x29, 0xaf, 0x73, 0x3d, 0xe7, 0x36, 0x78,
	0xe1, 0x68, 0x1a, 0x4d, 0x3f, 0x9f, 0x8e, 0xa2, 0xe9, 0xf8, 0x2c, 0x0c, 0x26, 0x53, 0xcf, 0x41,
	0x5b, 0xb0, 0xb9, 0x60, 0xcf, 0x47, 0xc1, 0x45, 0x70, 0xe2, 0xb5, 0xf6, 0x0e, 0x16, 0xff, 0x65,
	0x5d, 0x2d, 0xca, 0xa8, 0x52, 0xca, 0x00, 0x74, 0xc7, 0x93, 0x28, 0x18, 0x4f, 0x3c, 0x07, 0x6d,
	0x00, 0x8c, 0x27, 0xd1, 0xe1, 0xc9, 0x51, 0x38, 0x0e, 0x8e, 0xbc, 0x96, 0xf9, 0x76, 0x3e, 0x7a,
	0xeb, 0xb5, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xce, 0x41, 0xb6, 0x2f, 0xd5, 0x05, 0x00, 0x00,
}
