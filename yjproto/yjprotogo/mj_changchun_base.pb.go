// Code generated by protoc-gen-go.
// source: mj_changchun_base.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of GAME_ID from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

type CcmjProtoId int32

const (
	CcmjProtoId_id_p16_heartbeat              CcmjProtoId = 0
	CcmjProtoId_id_p16_req_createDesk         CcmjProtoId = 1
	CcmjProtoId_id_p16_ack_createDesk         CcmjProtoId = 2
	CcmjProtoId_id_p16_req_enterDesk          CcmjProtoId = 3
	CcmjProtoId_id_p16_ack_enterDesk          CcmjProtoId = 4
	CcmjProtoId_id_p16_req_dissolve_beginGame CcmjProtoId = 5
	CcmjProtoId_id_p16_ack_dissolve_beginGame CcmjProtoId = 6
	CcmjProtoId_id_p16_dissolve_beginGame     CcmjProtoId = 7
	CcmjProtoId_id_p16_ack_roomInit           CcmjProtoId = 8
	CcmjProtoId_id_p16_req_ready              CcmjProtoId = 9
	CcmjProtoId_id_p16_ack_ready              CcmjProtoId = 10
	CcmjProtoId_id_p16_ack_playerEnter        CcmjProtoId = 11
	CcmjProtoId_id_p16_ack_playerExit         CcmjProtoId = 12
	CcmjProtoId_id_p16_ack_game_deal_cards    CcmjProtoId = 13
	CcmjProtoId_id_p16_ack_game_opening       CcmjProtoId = 14
	CcmjProtoId_id_p16_ack_game_overturn      CcmjProtoId = 15
)

var CcmjProtoId_name = map[int32]string{
	0:  "id_p16_heartbeat",
	1:  "id_p16_req_createDesk",
	2:  "id_p16_ack_createDesk",
	3:  "id_p16_req_enterDesk",
	4:  "id_p16_ack_enterDesk",
	5:  "id_p16_req_dissolve_beginGame",
	6:  "id_p16_ack_dissolve_beginGame",
	7:  "id_p16_dissolve_beginGame",
	8:  "id_p16_ack_roomInit",
	9:  "id_p16_req_ready",
	10: "id_p16_ack_ready",
	11: "id_p16_ack_playerEnter",
	12: "id_p16_ack_playerExit",
	13: "id_p16_ack_game_deal_cards",
	14: "id_p16_ack_game_opening",
	15: "id_p16_ack_game_overturn",
}
var CcmjProtoId_value = map[string]int32{
	"id_p16_heartbeat":              0,
	"id_p16_req_createDesk":         1,
	"id_p16_ack_createDesk":         2,
	"id_p16_req_enterDesk":          3,
	"id_p16_ack_enterDesk":          4,
	"id_p16_req_dissolve_beginGame": 5,
	"id_p16_ack_dissolve_beginGame": 6,
	"id_p16_dissolve_beginGame":     7,
	"id_p16_ack_roomInit":           8,
	"id_p16_req_ready":              9,
	"id_p16_ack_ready":              10,
	"id_p16_ack_playerEnter":        11,
	"id_p16_ack_playerExit":         12,
	"id_p16_ack_game_deal_cards":    13,
	"id_p16_ack_game_opening":       14,
	"id_p16_ack_game_overturn":      15,
}

func (x CcmjProtoId) Enum() *CcmjProtoId {
	p := new(CcmjProtoId)
	*p = x
	return p
}
func (x CcmjProtoId) String() string {
	return proto.EnumName(CcmjProtoId_name, int32(x))
}
func (x *CcmjProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CcmjProtoId_value, data, "CcmjProtoId")
	if err != nil {
		return err
	}
	*x = CcmjProtoId(value)
	return nil
}
func (CcmjProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type RECONNECT_TYPE int32

const (
	RECONNECT_TYPE_NORMAL    RECONNECT_TYPE = 1
	RECONNECT_TYPE_RECONNECT RECONNECT_TYPE = 2
)

var RECONNECT_TYPE_name = map[int32]string{
	1: "NORMAL",
	2: "RECONNECT",
}
var RECONNECT_TYPE_value = map[string]int32{
	"NORMAL":    1,
	"RECONNECT": 2,
}

func (x RECONNECT_TYPE) Enum() *RECONNECT_TYPE {
	p := new(RECONNECT_TYPE)
	*p = x
	return p
}
func (x RECONNECT_TYPE) String() string {
	return proto.EnumName(RECONNECT_TYPE_name, int32(x))
}
func (x *RECONNECT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RECONNECT_TYPE_value, data, "RECONNECT_TYPE")
	if err != nil {
		return err
	}
	*x = RECONNECT_TYPE(value)
	return nil
}
func (RECONNECT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// 碰杠类型(客户端显示用)
type ComposeCardType int32

const (
	ComposeCardType_C_MINGGANG ComposeCardType = 1
	ComposeCardType_C_BAGANG   ComposeCardType = 2
	ComposeCardType_C_ANGANG   ComposeCardType = 3
	ComposeCardType_C_PENG     ComposeCardType = 4
	ComposeCardType_C_CHI      ComposeCardType = 5
)

var ComposeCardType_name = map[int32]string{
	1: "C_MINGGANG",
	2: "C_BAGANG",
	3: "C_ANGANG",
	4: "C_PENG",
	5: "C_CHI",
}
var ComposeCardType_value = map[string]int32{
	"C_MINGGANG": 1,
	"C_BAGANG":   2,
	"C_ANGANG":   3,
	"C_PENG":     4,
	"C_CHI":      5,
}

func (x ComposeCardType) Enum() *ComposeCardType {
	p := new(ComposeCardType)
	*p = x
	return p
}
func (x ComposeCardType) String() string {
	return proto.EnumName(ComposeCardType_name, int32(x))
}
func (x *ComposeCardType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ComposeCardType_value, data, "ComposeCardType")
	if err != nil {
		return err
	}
	*x = ComposeCardType(value)
	return nil
}
func (ComposeCardType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type DeskGameStatus int32

const (
	DeskGameStatus_INIT    DeskGameStatus = 0
	DeskGameStatus_FAPAI   DeskGameStatus = 1
	DeskGameStatus_PLAYING DeskGameStatus = 2
	DeskGameStatus_DINGQUE DeskGameStatus = 3
	DeskGameStatus_FINISH  DeskGameStatus = 4
	DeskGameStatus_PIAO    DeskGameStatus = 5
)

var DeskGameStatus_name = map[int32]string{
	0: "INIT",
	1: "FAPAI",
	2: "PLAYING",
	3: "DINGQUE",
	4: "FINISH",
	5: "PIAO",
}
var DeskGameStatus_value = map[string]int32{
	"INIT":    0,
	"FAPAI":   1,
	"PLAYING": 2,
	"DINGQUE": 3,
	"FINISH":  4,
	"PIAO":    5,
}

func (x DeskGameStatus) Enum() *DeskGameStatus {
	p := new(DeskGameStatus)
	*p = x
	return p
}
func (x DeskGameStatus) String() string {
	return proto.EnumName(DeskGameStatus_name, int32(x))
}
func (x *DeskGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeskGameStatus_value, data, "DeskGameStatus")
	if err != nil {
		return err
	}
	*x = DeskGameStatus(value)
	return nil
}
func (DeskGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type PaiType int32

const (
	PaiType_H_DuiDuiHu      PaiType = 1
	PaiType_H_QingYiSe      PaiType = 2
	PaiType_H_QiDui         PaiType = 3
	PaiType_H_DaiYaoJiu     PaiType = 4
	PaiType_H_LongQiDui     PaiType = 5
	PaiType_H_JiangDui      PaiType = 6
	PaiType_H_MenQing       PaiType = 7
	PaiType_H_ZhongZhang    PaiType = 8
	PaiType_H_QingLongQiDui PaiType = 9
	PaiType_H_QingQiDui     PaiType = 10
	PaiType_H_PingHu        PaiType = 11
)

var PaiType_name = map[int32]string{
	1:  "H_DuiDuiHu",
	2:  "H_QingYiSe",
	3:  "H_QiDui",
	4:  "H_DaiYaoJiu",
	5:  "H_LongQiDui",
	6:  "H_JiangDui",
	7:  "H_MenQing",
	8:  "H_ZhongZhang",
	9:  "H_QingLongQiDui",
	10: "H_QingQiDui",
	11: "H_PingHu",
}
var PaiType_value = map[string]int32{
	"H_DuiDuiHu":      1,
	"H_QingYiSe":      2,
	"H_QiDui":         3,
	"H_DaiYaoJiu":     4,
	"H_LongQiDui":     5,
	"H_JiangDui":      6,
	"H_MenQing":       7,
	"H_ZhongZhang":    8,
	"H_QingLongQiDui": 9,
	"H_QingQiDui":     10,
	"H_PingHu":        11,
}

func (x PaiType) Enum() *PaiType {
	p := new(PaiType)
	*p = x
	return p
}
func (x PaiType) String() string {
	return proto.EnumName(PaiType_name, int32(x))
}
func (x *PaiType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PaiType_value, data, "PaiType")
	if err != nil {
		return err
	}
	*x = PaiType(value)
	return nil
}
func (PaiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

// 房间类型枚举
type MJRoomType int32

const (
	MJRoomType_roomType_mj_changchun MJRoomType = 16
)

var MJRoomType_name = map[int32]string{
	16: "roomType_mj_changchun",
}
var MJRoomType_value = map[string]int32{
	"roomType_mj_changchun": 16,
}

func (x MJRoomType) Enum() *MJRoomType {
	p := new(MJRoomType)
	*p = x
	return p
}
func (x MJRoomType) String() string {
	return proto.EnumName(MJRoomType_name, int32(x))
}
func (x *MJRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MJRoomType_value, data, "MJRoomType")
	if err != nil {
		return err
	}
	*x = MJRoomType(value)
	return nil
}
func (MJRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type UserCoinBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64  `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserCoinBean) Reset()                    { *m = UserCoinBean{} }
func (m *UserCoinBean) String() string            { return proto.CompactTextString(m) }
func (*UserCoinBean) ProtoMessage()               {}
func (*UserCoinBean) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *UserCoinBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *UserCoinBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 房间信息
type RoomTypeInfo struct {
	MjRoomType       *MJRoomType  `protobuf:"varint,1,opt,name=mjRoomType,enum=yjprotogo.MJRoomType" json:"mjRoomType,omitempty"`
	BoardsCout       *int32       `protobuf:"varint,2,opt,name=boardsCout" json:"boardsCout,omitempty"`
	BaseValue        *int64       `protobuf:"varint,3,opt,name=baseValue" json:"baseValue,omitempty"`
	PlayOptions      *PlayOptions `protobuf:"bytes,4,opt,name=playOptions" json:"playOptions,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RoomTypeInfo) Reset()                    { *m = RoomTypeInfo{} }
func (m *RoomTypeInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomTypeInfo) ProtoMessage()               {}
func (*RoomTypeInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RoomTypeInfo) GetMjRoomType() MJRoomType {
	if m != nil && m.MjRoomType != nil {
		return *m.MjRoomType
	}
	return MJRoomType_roomType_mj_changchun
}

func (m *RoomTypeInfo) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *RoomTypeInfo) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *RoomTypeInfo) GetPlayOptions() *PlayOptions {
	if m != nil {
		return m.PlayOptions
	}
	return nil
}

type EndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	CountHu          *int32  `protobuf:"varint,6,opt,name=countHu" json:"countHu,omitempty"`
	CountZiMo        *int32  `protobuf:"varint,7,opt,name=countZiMo" json:"countZiMo,omitempty"`
	CountDianPao     *int32  `protobuf:"varint,8,opt,name=countDianPao" json:"countDianPao,omitempty"`
	CountAnGang      *int32  `protobuf:"varint,9,opt,name=countAnGang" json:"countAnGang,omitempty"`
	CountMingGang    *int32  `protobuf:"varint,10,opt,name=countMingGang" json:"countMingGang,omitempty"`
	CountDianGang    *int32  `protobuf:"varint,11,opt,name=countDianGang" json:"countDianGang,omitempty"`
	CountChaJiao     *int32  `protobuf:"varint,12,opt,name=countChaJiao" json:"countChaJiao,omitempty"`
	BestGunner       *bool   `protobuf:"varint,13,opt,name=bestGunner" json:"bestGunner,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EndLotteryInfo) Reset()                    { *m = EndLotteryInfo{} }
func (m *EndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*EndLotteryInfo) ProtoMessage()               {}
func (*EndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *EndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *EndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *EndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *EndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *EndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *EndLotteryInfo) GetCountHu() int32 {
	if m != nil && m.CountHu != nil {
		return *m.CountHu
	}
	return 0
}

func (m *EndLotteryInfo) GetCountZiMo() int32 {
	if m != nil && m.CountZiMo != nil {
		return *m.CountZiMo
	}
	return 0
}

func (m *EndLotteryInfo) GetCountDianPao() int32 {
	if m != nil && m.CountDianPao != nil {
		return *m.CountDianPao
	}
	return 0
}

func (m *EndLotteryInfo) GetCountAnGang() int32 {
	if m != nil && m.CountAnGang != nil {
		return *m.CountAnGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountMingGang() int32 {
	if m != nil && m.CountMingGang != nil {
		return *m.CountMingGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountDianGang() int32 {
	if m != nil && m.CountDianGang != nil {
		return *m.CountDianGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountChaJiao() int32 {
	if m != nil && m.CountChaJiao != nil {
		return *m.CountChaJiao
	}
	return 0
}

func (m *EndLotteryInfo) GetBestGunner() bool {
	if m != nil && m.BestGunner != nil {
		return *m.BestGunner
	}
	return false
}

// 玩法：包括自摸、点炮、以及可多选的玩法
type PlayOptions struct {
	DianPaoBaoSanJia *bool  `protobuf:"varint,1,opt,name=dianPaoBaoSanJia" json:"dianPaoBaoSanJia,omitempty"`
	XiaoJiFeiDan     *bool  `protobuf:"varint,2,opt,name=xiaoJiFeiDan" json:"xiaoJiFeiDan,omitempty"`
	SanFengDan       *bool  `protobuf:"varint,3,opt,name=sanFengDan" json:"sanFengDan,omitempty"`
	XiaDanSuanZhanLi *bool  `protobuf:"varint,4,opt,name=xiaDanSuanZhanLi" json:"xiaDanSuanZhanLi,omitempty"`
	DaiQueMen        *bool  `protobuf:"varint,5,opt,name=daiQueMen" json:"daiQueMen,omitempty"`
	UserCountLimit   *int32 `protobuf:"varint,6,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayOptions) Reset()                    { *m = PlayOptions{} }
func (m *PlayOptions) String() string            { return proto.CompactTextString(m) }
func (*PlayOptions) ProtoMessage()               {}
func (*PlayOptions) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *PlayOptions) GetDianPaoBaoSanJia() bool {
	if m != nil && m.DianPaoBaoSanJia != nil {
		return *m.DianPaoBaoSanJia
	}
	return false
}

func (m *PlayOptions) GetXiaoJiFeiDan() bool {
	if m != nil && m.XiaoJiFeiDan != nil {
		return *m.XiaoJiFeiDan
	}
	return false
}

func (m *PlayOptions) GetSanFengDan() bool {
	if m != nil && m.SanFengDan != nil {
		return *m.SanFengDan
	}
	return false
}

func (m *PlayOptions) GetXiaDanSuanZhanLi() bool {
	if m != nil && m.XiaDanSuanZhanLi != nil {
		return *m.XiaDanSuanZhanLi
	}
	return false
}

func (m *PlayOptions) GetDaiQueMen() bool {
	if m != nil && m.DaiQueMen != nil {
		return *m.DaiQueMen
	}
	return false
}

func (m *PlayOptions) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

type DeskGameInfo struct {
	GameStatus       *int32        `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	PlayerNum        *int32        `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32       `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	ActionTime       *int32        `protobuf:"varint,5,opt,name=actionTime" json:"actionTime,omitempty"`
	DelayTime        *int32        `protobuf:"varint,6,opt,name=delayTime" json:"delayTime,omitempty"`
	NInitActionTime  *int32        `protobuf:"varint,7,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	NInitDelayTime   *int32        `protobuf:"varint,8,opt,name=nInitDelayTime" json:"nInitDelayTime,omitempty"`
	InitRoomCoin     *int64        `protobuf:"varint,9,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32        `protobuf:"varint,10,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32        `protobuf:"varint,11,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string       `protobuf:"bytes,12,opt,name=roomNumber" json:"roomNumber,omitempty"`
	RemainCards      *int32        `protobuf:"varint,13,opt,name=remainCards" json:"remainCards,omitempty"`
	Banker           *uint32       `protobuf:"varint,14,opt,name=Banker" json:"Banker,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DeskGameInfo) Reset()                    { *m = DeskGameInfo{} }
func (m *DeskGameInfo) String() string            { return proto.CompactTextString(m) }
func (*DeskGameInfo) ProtoMessage()               {}
func (*DeskGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *DeskGameInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *DeskGameInfo) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *DeskGameInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *DeskGameInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *DeskGameInfo) GetActionTime() int32 {
	if m != nil && m.ActionTime != nil {
		return *m.ActionTime
	}
	return 0
}

func (m *DeskGameInfo) GetDelayTime() int32 {
	if m != nil && m.DelayTime != nil {
		return *m.DelayTime
	}
	return 0
}

func (m *DeskGameInfo) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *DeskGameInfo) GetNInitDelayTime() int32 {
	if m != nil && m.NInitDelayTime != nil {
		return *m.NInitDelayTime
	}
	return 0
}

func (m *DeskGameInfo) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *DeskGameInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *DeskGameInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *DeskGameInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *DeskGameInfo) GetRemainCards() int32 {
	if m != nil && m.RemainCards != nil {
		return *m.RemainCards
	}
	return 0
}

func (m *DeskGameInfo) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

type PlayerInfo struct {
	IsBanker         *bool       `protobuf:"varint,1,opt,name=isBanker" json:"isBanker,omitempty"`
	PlayerCard       *PlayerCard `protobuf:"bytes,2,opt,name=playerCard" json:"playerCard,omitempty"`
	Coin             *int64      `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	NickName         *string     `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32      `protobuf:"varint,5,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32     `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool       `protobuf:"varint,7,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32      `protobuf:"varint,8,opt,name=bReady" json:"bReady,omitempty"`
	NHuPai           *int32      `protobuf:"varint,9,opt,name=nHuPai" json:"nHuPai,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,10,opt,name=wxInfo" json:"wxInfo,omitempty"`
	GameStatus       *int32      `protobuf:"varint,11,opt,name=GameStatus" json:"GameStatus,omitempty"`
	AgentMode        *bool       `protobuf:"varint,12,opt,name=agentMode" json:"agentMode,omitempty"`
	Ip               *string     `protobuf:"bytes,13,opt,name=ip" json:"ip,omitempty"`
	IsBaoTing        *bool       `protobuf:"varint,14,opt,name=isBaoTing" json:"isBaoTing,omitempty"`
	Address          *string     `protobuf:"bytes,15,opt,name=address" json:"address,omitempty"`
	Site             *int32      `protobuf:"varint,16,opt,name=site" json:"site,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *PlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *PlayerInfo) GetPlayerCard() *PlayerCard {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *PlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *PlayerInfo) GetNHuPai() int32 {
	if m != nil && m.NHuPai != nil {
		return *m.NHuPai
	}
	return 0
}

func (m *PlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PlayerInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PlayerInfo) GetAgentMode() bool {
	if m != nil && m.AgentMode != nil {
		return *m.AgentMode
	}
	return false
}

func (m *PlayerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *PlayerInfo) GetIsBaoTing() bool {
	if m != nil && m.IsBaoTing != nil {
		return *m.IsBaoTing
	}
	return false
}

func (m *PlayerInfo) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *PlayerInfo) GetSite() int32 {
	if m != nil && m.Site != nil {
		return *m.Site
	}
	return 0
}

// 麻将牌
type CardInfo struct {
	Type             *int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Value            *int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Id               *int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CardInfo) Reset()                    { *m = CardInfo{} }
func (m *CardInfo) String() string            { return proto.CompactTextString(m) }
func (*CardInfo) ProtoMessage()               {}
func (*CardInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *CardInfo) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *CardInfo) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *CardInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type ComposeCard struct {
	Value            *int32  `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	ChiValue         []int32 `protobuf:"varint,3,rep,name=chiValue" json:"chiValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ComposeCard) Reset()                    { *m = ComposeCard{} }
func (m *ComposeCard) String() string            { return proto.CompactTextString(m) }
func (*ComposeCard) ProtoMessage()               {}
func (*ComposeCard) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ComposeCard) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ComposeCard) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ComposeCard) GetChiValue() []int32 {
	if m != nil {
		return m.ChiValue
	}
	return nil
}

type PlayerCard struct {
	HandCard         []*CardInfo    `protobuf:"bytes,1,rep,name=handCard" json:"handCard,omitempty"`
	ComposeCard      []*ComposeCard `protobuf:"bytes,2,rep,name=composeCard" json:"composeCard,omitempty"`
	OutCard          []*CardInfo    `protobuf:"bytes,3,rep,name=outCard" json:"outCard,omitempty"`
	HuCard           []*CardInfo    `protobuf:"bytes,4,rep,name=huCard" json:"huCard,omitempty"`
	UserId           *uint32        `protobuf:"varint,5,opt,name=UserId" json:"UserId,omitempty"`
	HandCardCount    *int32         `protobuf:"varint,6,opt,name=handCardCount" json:"handCardCount,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PlayerCard) Reset()                    { *m = PlayerCard{} }
func (m *PlayerCard) String() string            { return proto.CompactTextString(m) }
func (*PlayerCard) ProtoMessage()               {}
func (*PlayerCard) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *PlayerCard) GetHandCard() []*CardInfo {
	if m != nil {
		return m.HandCard
	}
	return nil
}

func (m *PlayerCard) GetComposeCard() []*ComposeCard {
	if m != nil {
		return m.ComposeCard
	}
	return nil
}

func (m *PlayerCard) GetOutCard() []*CardInfo {
	if m != nil {
		return m.OutCard
	}
	return nil
}

func (m *PlayerCard) GetHuCard() []*CardInfo {
	if m != nil {
		return m.HuCard
	}
	return nil
}

func (m *PlayerCard) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PlayerCard) GetHandCardCount() int32 {
	if m != nil && m.HandCardCount != nil {
		return *m.HandCardCount
	}
	return 0
}

// 赢牌信息：谁赢了多少
type WinCoinInfo struct {
	NickName         *string     `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32     `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	WinCoin          *int64      `protobuf:"varint,3,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64      `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	CardTitle        *string     `protobuf:"bytes,5,opt,name=cardTitle" json:"cardTitle,omitempty"`
	Cards            *PlayerCard `protobuf:"bytes,6,opt,name=cards" json:"cards,omitempty"`
	IsDealer         *bool       `protobuf:"varint,7,opt,name=isDealer" json:"isDealer,omitempty"`
	HuCount          *int32      `protobuf:"varint,8,opt,name=huCount" json:"huCount,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *WinCoinInfo) Reset()                    { *m = WinCoinInfo{} }
func (m *WinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*WinCoinInfo) ProtoMessage()               {}
func (*WinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *WinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *WinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *WinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WinCoinInfo) GetCardTitle() string {
	if m != nil && m.CardTitle != nil {
		return *m.CardTitle
	}
	return ""
}

func (m *WinCoinInfo) GetCards() *PlayerCard {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *WinCoinInfo) GetIsDealer() bool {
	if m != nil && m.IsDealer != nil {
		return *m.IsDealer
	}
	return false
}

func (m *WinCoinInfo) GetHuCount() int32 {
	if m != nil && m.HuCount != nil {
		return *m.HuCount
	}
	return 0
}

func init() {
	proto.RegisterType((*UserCoinBean)(nil), "yjprotogo.user_coin_bean")
	proto.RegisterType((*RoomTypeInfo)(nil), "yjprotogo.RoomTypeInfo")
	proto.RegisterType((*EndLotteryInfo)(nil), "yjprotogo.EndLotteryInfo")
	proto.RegisterType((*PlayOptions)(nil), "yjprotogo.PlayOptions")
	proto.RegisterType((*DeskGameInfo)(nil), "yjprotogo.DeskGameInfo")
	proto.RegisterType((*PlayerInfo)(nil), "yjprotogo.PlayerInfo")
	proto.RegisterType((*CardInfo)(nil), "yjprotogo.CardInfo")
	proto.RegisterType((*ComposeCard)(nil), "yjprotogo.ComposeCard")
	proto.RegisterType((*PlayerCard)(nil), "yjprotogo.PlayerCard")
	proto.RegisterType((*WinCoinInfo)(nil), "yjprotogo.WinCoinInfo")
	proto.RegisterEnum("yjprotogo.CcmjProtoId", CcmjProtoId_name, CcmjProtoId_value)
	proto.RegisterEnum("yjprotogo.RECONNECT_TYPE", RECONNECT_TYPE_name, RECONNECT_TYPE_value)
	proto.RegisterEnum("yjprotogo.ComposeCardType", ComposeCardType_name, ComposeCardType_value)
	proto.RegisterEnum("yjprotogo.DeskGameStatus", DeskGameStatus_name, DeskGameStatus_value)
	proto.RegisterEnum("yjprotogo.PaiType", PaiType_name, PaiType_value)
	proto.RegisterEnum("yjprotogo.MJRoomType", MJRoomType_name, MJRoomType_value)
}

var fileDescriptor2 = []byte{
	// 1571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x57, 0x5d, 0x6e, 0x1b, 0xc9,
	0x11, 0x36, 0xff, 0x87, 0x45, 0x8a, 0x6a, 0xb4, 0xd7, 0xf6, 0xac, 0xe2, 0x35, 0x14, 0x62, 0x91,
	0x08, 0x32, 0xd6, 0x41, 0x0c, 0x24, 0x08, 0x90, 0xbc, 0xd0, 0xa4, 0x2c, 0x52, 0x90, 0x28, 0x7a,
	0x24, 0xc7, 0xb0, 0x5f, 0x06, 0x2d, 0x4e, 0x67, 0xd8, 0x36, 0xd9, 0xc3, 0xcc, 0x8f, 0x2d, 0x1d,
	0x21, 0x8f, 0xb9, 0x40, 0x6e, 0xb2, 0x17, 0x08, 0x72, 0x97, 0xe4, 0x08, 0x41, 0x55, 0xcf, 0x70,
	0x9a, 0x94, 0xbd, 0x6f, 0x53, 0x5f, 0x7d, 0x5d, 0x5d, 0xff, 0x4d, 0xc2, 0x93, 0xd5, 0x47, 0x7f,
	0xbe, 0x10, 0x3a, 0x9c, 0x2f, 0x32, 0xed, 0xdf, 0x88, 0x44, 0xbe, 0x58, 0xc7, 0x51, 0x1a, 0xf1,
	0xf6, 0xdd, 0x47, 0xfa, 0x08, 0xa3, 0x83, 0x87, 0xf3, 0x68, 0xb5, 0x8a, 0xb4, 0x3f, 0x5f, 0x2a,
	0xa9, 0x53, 0xa3, 0xef, 0xff, 0x05, 0x7a, 0x59, 0x22, 0x63, 0x7f, 0x1e, 0x29, 0xed, 0xdf, 0x48,
	0xa1, 0xf9, 0x63, 0x68, 0x22, 0x32, 0x09, 0xdc, 0xca, 0x61, 0xe5, 0x68, 0xcf, 0xcb, 0x25, 0xce,
	0xa1, 0x8e, 0x24, 0xb7, 0x7a, 0x58, 0x39, 0xaa, 0x79, 0xf4, 0xdd, 0xff, 0xb9, 0x02, 0x5d, 0x2f,
	0x8a, 0x56, 0xd7, 0x77, 0x6b, 0x39, 0xd1, 0x7f, 0x8b, 0xf8, 0x1f, 0x00, 0x56, 0x1f, 0x0b, 0x84,
	0x0c, 0xf4, 0x5e, 0x3e, 0x7a, 0xb1, 0xf1, 0xe1, 0xc5, 0xc5, 0x59, 0xa1, 0xf4, 0x2c, 0x22, 0x7f,
	0x06, 0x70, 0x13, 0x89, 0x38, 0x48, 0x86, 0x51, 0x96, 0xd2, 0x0d, 0x0d, 0xcf, 0x42, 0xf8, 0x53,
	0x68, 0x63, 0x4c, 0x7f, 0x15, 0xcb, 0x4c, 0xba, 0x35, 0x72, 0xa0, 0x04, 0xf8, 0x9f, 0xa0, 0xb3,
	0x5e, 0x8a, 0xbb, 0xcb, 0x75, 0xaa, 0x22, 0x9d, 0xb8, 0xf5, 0xc3, 0xca, 0x51, 0xe7, 0xe5, 0x63,
	0xeb, 0xd6, 0x59, 0xa9, 0xf5, 0x6c, 0x6a, 0xff, 0x5f, 0x35, 0xe8, 0x9d, 0xe8, 0xe0, 0x3c, 0x4a,
	0x53, 0x19, 0xdf, 0x51, 0x04, 0xdf, 0x0a, 0xff, 0x00, 0x1c, 0xad, 0xe6, 0x9f, 0xa6, 0x62, 0x25,
	0xc9, 0xc1, 0xb6, 0xb7, 0x91, 0xf1, 0xcc, 0x8d, 0x0a, 0xdf, 0x29, 0x4d, 0xbe, 0x39, 0x5e, 0x2e,
	0x71, 0x17, 0x5a, 0x2a, 0xb9, 0xfc, 0xa2, 0x65, 0x4c, 0x4e, 0x39, 0x5e, 0x21, 0xa2, 0xe6, 0x8b,
	0xd2, 0x43, 0xcc, 0x67, 0x83, 0xc2, 0x29, 0x44, 0xd4, 0xcc, 0xa3, 0x4c, 0xa7, 0xe3, 0xcc, 0x6d,
	0x52, 0x1e, 0x0a, 0x11, 0x93, 0x40, 0x9f, 0x1f, 0xd4, 0x45, 0xe4, 0xb6, 0x48, 0x57, 0x02, 0xbc,
	0x0f, 0x5d, 0x12, 0x46, 0x4a, 0xe8, 0x99, 0x88, 0x5c, 0x87, 0x08, 0x5b, 0x18, 0x3f, 0x84, 0x0e,
	0xc9, 0x03, 0x7d, 0x2a, 0x74, 0xe8, 0xb6, 0x89, 0x62, 0x43, 0xfc, 0x47, 0xd8, 0x23, 0xf1, 0x42,
	0xe9, 0x90, 0x38, 0x40, 0x9c, 0x6d, 0x70, 0xc3, 0x42, 0xbb, 0xc4, 0xea, 0x58, 0xac, 0x02, 0xdc,
	0x78, 0x34, 0x5c, 0x88, 0x33, 0x25, 0x22, 0xb7, 0x6b, 0x79, 0x94, 0x63, 0x54, 0x78, 0x99, 0xa4,
	0xa7, 0x99, 0xc6, 0x24, 0xed, 0x51, 0x92, 0x2c, 0xa4, 0xff, 0xdf, 0x0a, 0x74, 0xac, 0xea, 0xf1,
	0x63, 0x60, 0x81, 0x09, 0xe6, 0x95, 0x88, 0xae, 0x84, 0x3e, 0x53, 0x82, 0xea, 0xe4, 0x78, 0xf7,
	0x70, 0xbc, 0xff, 0x56, 0x89, 0xe8, 0x4c, 0xbd, 0x96, 0x6a, 0x24, 0x4c, 0xe3, 0x3a, 0xde, 0x16,
	0x86, 0xf7, 0x27, 0x42, 0xbf, 0x96, 0x3a, 0x44, 0x86, 0xa9, 0x9e, 0x85, 0xe0, 0x7d, 0xb7, 0x4a,
	0x8c, 0x84, 0xbe, 0xca, 0x84, 0xfe, 0xb0, 0x10, 0xfa, 0x5c, 0xe5, 0xa5, 0xbc, 0x87, 0x63, 0x7d,
	0x02, 0xa1, 0xde, 0x64, 0xf2, 0x42, 0x9a, 0xaa, 0x3a, 0x5e, 0x09, 0xf0, 0xdf, 0x98, 0x41, 0x1b,
	0x62, 0xf4, 0xe7, 0x6a, 0xa5, 0xd2, 0xbc, 0xbc, 0x3b, 0x68, 0xff, 0x1f, 0x75, 0xe8, 0x8e, 0x64,
	0xf2, 0xe9, 0x54, 0xac, 0xcc, 0x48, 0x3d, 0x03, 0xc0, 0xef, 0xab, 0x54, 0xa4, 0x59, 0x42, 0xc1,
	0x36, 0x3c, 0x0b, 0xe1, 0x7f, 0x86, 0x6e, 0x6c, 0x8d, 0x20, 0x85, 0xd9, 0x79, 0xf9, 0xc4, 0x6a,
	0x7f, 0x7b, 0x42, 0xbd, 0x2d, 0x32, 0xfa, 0x8c, 0xf3, 0x20, 0xe3, 0x69, 0xb6, 0xa2, 0xf0, 0x1b,
	0x5e, 0x09, 0x60, 0x06, 0xc5, 0x3c, 0x55, 0x9f, 0xe5, 0x5b, 0x33, 0x11, 0x75, 0x9a, 0x88, 0x2d,
	0x0c, 0xdd, 0x43, 0x39, 0xd2, 0xd7, 0x6a, 0x25, 0x29, 0xec, 0x86, 0x67, 0x21, 0x94, 0x15, 0xb9,
	0x14, 0x77, 0xa4, 0x36, 0x21, 0x97, 0x00, 0x3f, 0x82, 0x7d, 0x3d, 0xd1, 0x2a, 0x1d, 0x94, 0x26,
	0x4c, 0x67, 0xef, 0xc2, 0x98, 0x3f, 0x82, 0x46, 0x1b, 0x63, 0xa6, 0xc3, 0x77, 0x50, 0xf4, 0x59,
	0x69, 0x95, 0x62, 0xcc, 0x34, 0x5e, 0x6d, 0x1a, 0xaf, 0x2d, 0x8c, 0xfa, 0x37, 0x8b, 0x63, 0x6c,
	0x2c, 0xca, 0xfc, 0xa6, 0xcb, 0x6d, 0x10, 0x6f, 0x4c, 0xa3, 0x54, 0x2c, 0x4b, 0x9a, 0x69, 0xf3,
	0x1d, 0x14, 0x33, 0x80, 0x39, 0x9d, 0x66, 0xab, 0x1b, 0x19, 0x53, 0x97, 0xb7, 0x3d, 0x0b, 0xc1,
	0xa9, 0x8b, 0xe5, 0x4a, 0x28, 0x3d, 0xc4, 0x7d, 0x46, 0x4d, 0xde, 0xf0, 0x6c, 0x08, 0xf7, 0xc7,
	0x2b, 0xa1, 0x3f, 0xc9, 0xd8, 0xed, 0x99, 0x9d, 0x63, 0xa4, 0xfe, 0x7f, 0x6a, 0x00, 0x33, 0xaa,
	0x06, 0x15, 0xeb, 0x00, 0x1c, 0x95, 0xe4, 0x44, 0xd3, 0xf4, 0x1b, 0x19, 0x17, 0xaf, 0xa9, 0x1b,
	0x5a, 0xcc, 0x7b, 0xe0, 0xd1, 0xce, 0x0a, 0x34, 0x4a, 0xcf, 0x22, 0x6e, 0x96, 0x7a, 0xad, 0x5c,
	0xea, 0x5b, 0x9b, 0xae, 0xbe, 0xb3, 0xe9, 0x18, 0xd4, 0x12, 0x79, 0x9b, 0x97, 0x19, 0x3f, 0xad,
	0x7d, 0xd9, 0xdc, 0xda, 0x97, 0xd6, 0xee, 0x6b, 0x6d, 0xef, 0x3e, 0xdc, 0x96, 0x9e, 0x14, 0xc1,
	0x5d, 0x5e, 0xc1, 0x5c, 0x42, 0x5c, 0x8f, 0xb3, 0x99, 0x50, 0xf9, 0x62, 0xca, 0x25, 0xfe, 0x13,
	0x34, 0xbf, 0xdc, 0x52, 0x6b, 0xc3, 0xbd, 0xb0, 0xde, 0x49, 0x75, 0xab, 0x34, 0x35, 0x76, 0x4e,
	0xda, 0x99, 0x97, 0xce, 0xbd, 0x79, 0x79, 0x0a, 0x6d, 0x11, 0x4a, 0x9d, 0x5e, 0x44, 0x81, 0xa4,
	0x6a, 0x39, 0x5e, 0x09, 0xf0, 0x1e, 0x54, 0xd5, 0x9a, 0x6a, 0xd4, 0xf6, 0xaa, 0x6a, 0x8d, 0x6c,
	0xcc, 0x71, 0x74, 0xad, 0x74, 0x48, 0xd5, 0x71, 0xbc, 0x12, 0xc0, 0x20, 0x45, 0x10, 0xc4, 0x32,
	0x49, 0xdc, 0x7d, 0x3a, 0x52, 0x88, 0x98, 0xd8, 0x44, 0xa5, 0xd2, 0x65, 0x74, 0x3f, 0x7d, 0xf7,
	0x47, 0xe0, 0x60, 0xd2, 0xc9, 0x4b, 0x0e, 0xf5, 0xb4, 0x78, 0x22, 0x1b, 0x1e, 0x7d, 0xf3, 0xef,
	0xa0, 0xf1, 0x99, 0x5e, 0x38, 0xf3, 0x00, 0x1a, 0x81, 0x3c, 0x0a, 0xf2, 0xd9, 0xac, 0xaa, 0xa0,
	0x7f, 0x05, 0x9d, 0x61, 0xb4, 0x5a, 0x47, 0x89, 0xa4, 0x0a, 0x6e, 0x0e, 0x55, 0xec, 0x43, 0x85,
	0xf9, 0xaa, 0x65, 0xfe, 0x00, 0x9c, 0xf9, 0x42, 0x15, 0x6f, 0x68, 0xed, 0xa8, 0xe1, 0x6d, 0xe4,
	0xfe, 0x3f, 0xab, 0x45, 0xa7, 0x91, 0xd1, 0xdf, 0x81, 0xb3, 0x10, 0x3a, 0xa0, 0x5e, 0xaa, 0x1c,
	0xd6, 0x8e, 0x3a, 0x2f, 0x1f, 0x5a, 0x49, 0x2f, 0x82, 0xf0, 0x36, 0x24, 0x7c, 0x82, 0xe7, 0xa5,
	0x53, 0x6e, 0x95, 0xce, 0xd8, 0x4f, 0xb0, 0xe5, 0xb2, 0x67, 0x53, 0xf9, 0x4f, 0xd0, 0x8a, 0xb2,
	0x94, 0x4e, 0xd5, 0xbe, 0x7d, 0x53, 0xc1, 0xe1, 0xcf, 0xa1, 0xb9, 0xc8, 0x88, 0x5d, 0xff, 0x36,
	0x3b, 0xa7, 0x60, 0x47, 0xe5, 0x9b, 0xab, 0x61, 0x7a, 0x33, 0xdf, 0x59, 0x3f, 0xc2, 0x5e, 0xe1,
	0xb9, 0x19, 0x6c, 0xb3, 0x97, 0xb6, 0xc1, 0xfe, 0xff, 0x2a, 0xd0, 0x79, 0x67, 0x5e, 0xe5, 0x62,
	0xfc, 0x36, 0x73, 0x51, 0xb9, 0xff, 0x0b, 0x20, 0x9f, 0x82, 0xea, 0xee, 0x14, 0x14, 0xef, 0x7c,
	0x6d, 0xfb, 0x9d, 0x2f, 0x26, 0xaf, 0x6e, 0x4d, 0x1e, 0xbe, 0xf0, 0x22, 0x0e, 0xae, 0x55, 0xba,
	0x34, 0xab, 0xb4, 0xed, 0x95, 0x00, 0x7f, 0x0e, 0x8d, 0x39, 0x6d, 0x90, 0xe6, 0x2f, 0x4d, 0xb7,
	0xe1, 0x98, 0x5d, 0x31, 0x92, 0x62, 0xb9, 0x99, 0xbf, 0x8d, 0x8c, 0x4e, 0x2d, 0x32, 0x13, 0xb8,
	0x99, 0xc0, 0x42, 0x3c, 0xfe, 0x77, 0x0d, 0xba, 0xf3, 0xf9, 0xea, 0xa3, 0x4f, 0x86, 0x27, 0xd8,
	0x5d, 0x4c, 0x05, 0xfe, 0xfa, 0xf7, 0x7f, 0xf4, 0x17, 0x52, 0xc4, 0xe9, 0x8d, 0x14, 0x29, 0x7b,
	0xc0, 0xbf, 0x87, 0x47, 0x39, 0x1a, 0xcb, 0xbf, 0xfb, 0xf3, 0x58, 0x8a, 0x54, 0xe2, 0x9b, 0xc5,
	0x2a, 0x96, 0x4a, 0xcc, 0x3f, 0xd9, 0xaa, 0x2a, 0x77, 0xe1, 0x3b, 0xeb, 0x94, 0xd4, 0xa9, 0x8c,
	0x49, 0x53, 0xb3, 0x34, 0x78, 0xa8, 0xd4, 0xd4, 0xf9, 0xaf, 0xe1, 0x07, 0xeb, 0x4c, 0xa0, 0x92,
	0x24, 0x5a, 0x7e, 0x96, 0xfe, 0x8d, 0x0c, 0x95, 0xc6, 0x99, 0x66, 0x0d, 0x8b, 0x82, 0x87, 0xbf,
	0x42, 0x69, 0xf2, 0x1f, 0xe0, 0xfb, 0x9c, 0xf2, 0x15, 0x75, 0x8b, 0x3f, 0x81, 0x87, 0x96, 0x05,
	0xdc, 0xdc, 0xf8, 0xa4, 0x30, 0xc7, 0x8a, 0x1e, 0x6f, 0x8f, 0x71, 0x4b, 0xb1, 0xb6, 0x85, 0x12,
	0x9d, 0x50, 0xe0, 0x07, 0xf0, 0xd8, 0x42, 0xcd, 0x8a, 0x3d, 0xc1, 0x48, 0x58, 0x67, 0x27, 0x29,
	0xb9, 0xee, 0x56, 0xa5, 0xac, 0xcb, 0x9f, 0xc1, 0x81, 0xa5, 0x0a, 0xc5, 0x4a, 0xfa, 0x81, 0x14,
	0x4b, 0x9f, 0xaa, 0xc8, 0xf6, 0xf8, 0xaf, 0xe0, 0xc9, 0xae, 0x3e, 0x5a, 0x4b, 0xad, 0x74, 0xc8,
	0x7a, 0xfc, 0x29, 0xb8, 0xf7, 0x94, 0x9f, 0x65, 0x9c, 0x66, 0xb1, 0x66, 0xfb, 0xc7, 0xcf, 0xa1,
	0xe7, 0x9d, 0x0c, 0x2f, 0xa7, 0xd3, 0x93, 0xe1, 0xb5, 0x7f, 0xfd, 0x7e, 0x76, 0xc2, 0x01, 0x9a,
	0xd3, 0x4b, 0xef, 0x62, 0x70, 0xce, 0x2a, 0x7c, 0x0f, 0xda, 0x1b, 0x2d, 0xab, 0x1e, 0x5f, 0xc3,
	0xbe, 0x35, 0xa2, 0xf4, 0xa3, 0xbc, 0x07, 0x30, 0xf4, 0x2f, 0x26, 0xd3, 0xd3, 0xd3, 0xc1, 0xf4,
	0x94, 0x55, 0x78, 0x17, 0x9c, 0xa1, 0xff, 0x6a, 0x40, 0x52, 0xd5, 0x48, 0x83, 0x29, 0x49, 0x35,
	0xb4, 0x3c, 0xf4, 0x67, 0x27, 0xd3, 0x53, 0x56, 0xe7, 0x6d, 0x68, 0x0c, 0xfd, 0xe1, 0x78, 0xc2,
	0x1a, 0xc7, 0x6f, 0xa1, 0x57, 0xfc, 0x96, 0xc9, 0xb7, 0xaf, 0x03, 0xf5, 0xc9, 0x74, 0x72, 0xcd,
	0x1e, 0x20, 0xed, 0xf5, 0x60, 0x36, 0x98, 0xb0, 0x0a, 0xef, 0x40, 0x6b, 0x76, 0x3e, 0x78, 0x3f,
	0x21, 0xc3, 0x1d, 0x68, 0x8d, 0x26, 0xd3, 0xd3, 0x37, 0x6f, 0x4f, 0x8c, 0xdd, 0xd7, 0x93, 0xe9,
	0xe4, 0x6a, 0xcc, 0xea, 0x78, 0x74, 0x36, 0x19, 0x5c, 0xb2, 0xc6, 0xf1, 0xcf, 0x15, 0x68, 0xcd,
	0x84, 0x2a, 0xbc, 0x1c, 0xfb, 0xa3, 0x4c, 0x8d, 0x32, 0x35, 0xce, 0x58, 0xc5, 0xc8, 0x6f, 0x94,
	0x0e, 0xdf, 0xab, 0x2b, 0x69, 0xcc, 0xa1, 0x3c, 0xca, 0x14, 0xab, 0xf1, 0x7d, 0xe8, 0x8c, 0xfd,
	0x91, 0x50, 0xef, 0xf1, 0x27, 0x60, 0xc6, 0xea, 0x06, 0x38, 0x8f, 0x74, 0x68, 0x18, 0x0d, 0x73,
	0xfc, 0x4c, 0x09, 0x1d, 0xa2, 0xdc, 0xc4, 0x34, 0x8d, 0xfd, 0x0b, 0xa9, 0xd1, 0x22, 0x6b, 0x71,
	0x06, 0xdd, 0xb1, 0xff, 0x61, 0x11, 0xe9, 0x10, 0x7f, 0xf4, 0x85, 0xcc, 0xe1, 0x0f, 0x61, 0xdf,
	0xdc, 0x57, 0x5a, 0x69, 0x1b, 0xb3, 0x08, 0x1a, 0x00, 0x30, 0x5b, 0x63, 0x7f, 0xa6, 0x74, 0x38,
	0xce, 0x58, 0xe7, 0xf8, 0xb7, 0x00, 0xe5, 0x1f, 0x21, 0xec, 0x8e, 0xe2, 0x37, 0x99, 0x6f, 0xff,
	0x8d, 0x63, 0x6c, 0xf6, 0xe0, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x0f, 0xbf, 0xca, 0xd9,
	0x0d, 0x00, 0x00,
}
