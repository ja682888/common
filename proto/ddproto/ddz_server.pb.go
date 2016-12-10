// Code generated by protoc-gen-go.
// source: ddz_server.proto
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

type EDDZTYPE int32

const (
	EDDZTYPE_DDZ_TYPE_HUANLEDIZHU EDDZTYPE = 1
	EDDZTYPE_DDZ_TYPE_SICHUAN     EDDZTYPE = 2
)

var EDDZTYPE_name = map[int32]string{
	1: "DDZ_TYPE_HUANLEDIZHU",
	2: "DDZ_TYPE_SICHUAN",
}
var EDDZTYPE_value = map[string]int32{
	"DDZ_TYPE_HUANLEDIZHU": 1,
	"DDZ_TYPE_SICHUAN":     2,
}

func (x EDDZTYPE) Enum() *EDDZTYPE {
	p := new(EDDZTYPE)
	*p = x
	return p
}
func (x EDDZTYPE) String() string {
	return proto.EnumName(EDDZTYPE_name, int32(x))
}
func (x *EDDZTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDDZTYPE_value, data, "EDDZTYPE")
	if err != nil {
		return err
	}
	*x = EDDZTYPE(value)
	return nil
}
func (EDDZTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// 打出去的牌
type DdzSrvOutPokerPais struct {
	KeyValue         *int32               `protobuf:"varint,1,opt,name=keyValue" json:"keyValue,omitempty"`
	PokerPais        []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pokerPais" json:"pokerPais,omitempty"`
	Type             *int32               `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	IsBomb           *bool                `protobuf:"varint,4,opt,name=isBomb" json:"isBomb,omitempty"`
	CountDuizi       *int32               `protobuf:"varint,5,opt,name=countDuizi" json:"countDuizi,omitempty"`
	CountSanzhang    *int32               `protobuf:"varint,6,opt,name=countSanzhang" json:"countSanzhang,omitempty"`
	CountSizhang     *int32               `protobuf:"varint,7,opt,name=countSizhang" json:"countSizhang,omitempty"`
	CountYizhang     *int32               `protobuf:"varint,8,opt,name=countYizhang" json:"countYizhang,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzSrvOutPokerPais) Reset()                    { *m = DdzSrvOutPokerPais{} }
func (m *DdzSrvOutPokerPais) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvOutPokerPais) ProtoMessage()               {}
func (*DdzSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DdzSrvOutPokerPais) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetPokerPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.PokerPais
	}
	return nil
}

func (m *DdzSrvOutPokerPais) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetIsBomb() bool {
	if m != nil && m.IsBomb != nil {
		return *m.IsBomb
	}
	return false
}

func (m *DdzSrvOutPokerPais) GetCountDuizi() int32 {
	if m != nil && m.CountDuizi != nil {
		return *m.CountDuizi
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSanzhang() int32 {
	if m != nil && m.CountSanzhang != nil {
		return *m.CountSanzhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSizhang() int32 {
	if m != nil && m.CountSizhang != nil {
		return *m.CountSizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountYizhang() int32 {
	if m != nil && m.CountYizhang != nil {
		return *m.CountYizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 对 desk的统计
type DdzSrvDeskTongJi struct {
	Bombs            []*DdzSrvOutPokerPais `protobuf:"bytes,1,rep,name=bombs" json:"bombs,omitempty"`
	CountQiangDiZhu  *int32                `protobuf:"varint,2,opt,name=countQiangDiZhu" json:"countQiangDiZhu,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvDeskTongJi) Reset()                    { *m = DdzSrvDeskTongJi{} }
func (m *DdzSrvDeskTongJi) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDeskTongJi) ProtoMessage()               {}
func (*DdzSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *DdzSrvDeskTongJi) GetBombs() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.Bombs
	}
	return nil
}

func (m *DdzSrvDeskTongJi) GetCountQiangDiZhu() int32 {
	if m != nil && m.CountQiangDiZhu != nil {
		return *m.CountQiangDiZhu
	}
	return 0
}

// desk
type DdzSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Key              *string              `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	UserCountLimit   *int32               `protobuf:"varint,3,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	AllPokerPai      []*CommonSrvPokerPai `protobuf:"bytes,4,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*CommonSrvPokerPai `protobuf:"bytes,5,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *DdzSrvOutPokerPais  `protobuf:"bytes,6,opt,name=outPai" json:"outPai,omitempty"`
	Owner            *uint32              `protobuf:"varint,7,opt,name=owner" json:"owner,omitempty"`
	DizhuPaiUser     *uint32              `protobuf:"varint,8,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	DiZhuUserId      *uint32              `protobuf:"varint,9,opt,name=diZhuUserId" json:"diZhuUserId,omitempty"`
	ActiveUserId     *uint32              `protobuf:"varint,10,opt,name=activeUserId" json:"activeUserId,omitempty"`
	Tongji           *DdzSrvDeskTongJi    `protobuf:"bytes,11,opt,name=tongji" json:"tongji,omitempty"`
	BaseValue        *int64               `protobuf:"varint,12,opt,name=baseValue" json:"baseValue,omitempty"`
	QingDizhuValue   *int64               `protobuf:"varint,13,opt,name=qingDizhuValue" json:"qingDizhuValue,omitempty"`
	WinValue         *int64               `protobuf:"varint,14,opt,name=winValue" json:"winValue,omitempty"`
	DdzType          *int32               `protobuf:"varint,15,opt,name=ddzType" json:"ddzType,omitempty"`
	RoomType         *int32               `protobuf:"varint,16,opt,name=RoomType" json:"RoomType,omitempty"`
	BoardsCount      *int32               `protobuf:"varint,17,opt,name=BoardsCount" json:"BoardsCount,omitempty"`
	CapMax           *int64               `protobuf:"varint,18,opt,name=CapMax" json:"CapMax,omitempty"`
	IsJiaoFen        *bool                `protobuf:"varint,19,opt,name=IsJiaoFen" json:"IsJiaoFen,omitempty"`
	RoomId           *int32               `protobuf:"varint,20,opt,name=roomId" json:"roomId,omitempty"`
	FootRate         *int32               `protobuf:"varint,21,opt,name=footRate" json:"footRate,omitempty"`
	PlayRate         *int32               `protobuf:"varint,22,opt,name=playRate" json:"playRate,omitempty"`
	GameStatus       *int32               `protobuf:"varint,23,opt,name=GameStatus" json:"GameStatus,omitempty"`
	InitRoomCoin     *int64               `protobuf:"varint,24,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32               `protobuf:"varint,25,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32               `protobuf:"varint,26,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	PlayerNum        *int32               `protobuf:"varint,27,opt,name=playerNum" json:"playerNum,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzSrvDesk) Reset()                    { *m = DdzSrvDesk{} }
func (m *DdzSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDesk) ProtoMessage()               {}
func (*DdzSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *DdzSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzSrvDesk) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *DdzSrvDesk) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func (m *DdzSrvDesk) GetAllPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetDiPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.DiPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *DdzSrvDesk) GetDizhuPaiUser() uint32 {
	if m != nil && m.DizhuPaiUser != nil {
		return *m.DizhuPaiUser
	}
	return 0
}

func (m *DdzSrvDesk) GetDiZhuUserId() uint32 {
	if m != nil && m.DiZhuUserId != nil {
		return *m.DiZhuUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetTongji() *DdzSrvDeskTongJi {
	if m != nil {
		return m.Tongji
	}
	return nil
}

func (m *DdzSrvDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *DdzSrvDesk) GetQingDizhuValue() int64 {
	if m != nil && m.QingDizhuValue != nil {
		return *m.QingDizhuValue
	}
	return 0
}

func (m *DdzSrvDesk) GetWinValue() int64 {
	if m != nil && m.WinValue != nil {
		return *m.WinValue
	}
	return 0
}

func (m *DdzSrvDesk) GetDdzType() int32 {
	if m != nil && m.DdzType != nil {
		return *m.DdzType
	}
	return 0
}

func (m *DdzSrvDesk) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *DdzSrvDesk) GetBoardsCount() int32 {
	if m != nil && m.BoardsCount != nil {
		return *m.BoardsCount
	}
	return 0
}

func (m *DdzSrvDesk) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *DdzSrvDesk) GetIsJiaoFen() bool {
	if m != nil && m.IsJiaoFen != nil {
		return *m.IsJiaoFen
	}
	return false
}

func (m *DdzSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *DdzSrvDesk) GetFootRate() int32 {
	if m != nil && m.FootRate != nil {
		return *m.FootRate
	}
	return 0
}

func (m *DdzSrvDesk) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *DdzSrvDesk) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

// 游戏数据
type DdzSrvGameData struct {
	HandPokers       []*CommonSrvPokerPai  `protobuf:"bytes,1,rep,name=handPokers" json:"handPokers,omitempty"`
	OutPaiList       []*DdzSrvOutPokerPais `protobuf:"bytes,2,rep,name=outPaiList" json:"outPaiList,omitempty"`
	OutPai           *DdzSrvOutPokerPais   `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvGameData) Reset()                    { *m = DdzSrvGameData{} }
func (m *DdzSrvGameData) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvGameData) ProtoMessage()               {}
func (*DdzSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DdzSrvGameData) GetHandPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPaiList() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

type DdzSrvBillBean struct {
	// 斗地主的账单
	Coin             *int64  `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	LoseUser         *uint32 `protobuf:"varint,2,opt,name=loseUser" json:"loseUser,omitempty"`
	WinUser          *uint32 `protobuf:"varint,3,opt,name=winUser" json:"winUser,omitempty"`
	Desc             *string `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DdzSrvBillBean) Reset()                    { *m = DdzSrvBillBean{} }
func (m *DdzSrvBillBean) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBillBean) ProtoMessage()               {}
func (*DdzSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *DdzSrvBillBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvBillBean) GetLoseUser() uint32 {
	if m != nil && m.LoseUser != nil {
		return *m.LoseUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetWinUser() uint32 {
	if m != nil && m.WinUser != nil {
		return *m.WinUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

type DdzSrvBill struct {
	// 斗地主的账单
	WinCoin          *int64            `protobuf:"varint,1,opt,name=winCoin" json:"winCoin,omitempty"`
	BillBean         []*DdzSrvBillBean `protobuf:"bytes,2,rep,name=billBean" json:"billBean,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DdzSrvBill) Reset()                    { *m = DdzSrvBill{} }
func (m *DdzSrvBill) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBill) ProtoMessage()               {}
func (*DdzSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *DdzSrvBill) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvBill) GetBillBean() []*DdzSrvBillBean {
	if m != nil {
		return m.BillBean
	}
	return nil
}

type DdzSrvUserStatistics struct {
	BomBean          []*DdzSrvOutPokerPais `protobuf:"bytes,1,rep,name=bomBean" json:"bomBean,omitempty"`
	CountWin         *int32                `protobuf:"varint,2,opt,name=CountWin" json:"CountWin,omitempty"`
	CountLose        *int32                `protobuf:"varint,3,opt,name=CountLose" json:"CountLose,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUserStatistics) Reset()                    { *m = DdzSrvUserStatistics{} }
func (m *DdzSrvUserStatistics) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatistics) ProtoMessage()               {}
func (*DdzSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *DdzSrvUserStatistics) GetBomBean() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.BomBean
	}
	return nil
}

func (m *DdzSrvUserStatistics) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

// user
type DdzSrvUser struct {
	UserId           *uint32               `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData         *DdzSrvGameData       `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	Status           *int32                `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	IsBreak          *bool                 `protobuf:"varint,4,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool                 `protobuf:"varint,5,opt,name=isLeave" json:"isLeave,omitempty"`
	QiangDiZhuStatus *int32                `protobuf:"varint,7,opt,name=qiangDiZhuStatus" json:"qiangDiZhuStatus,omitempty"`
	JiabeiStatus     *int32                `protobuf:"varint,8,opt,name=jiabeiStatus" json:"jiabeiStatus,omitempty"`
	Bill             *DdzSrvBill           `protobuf:"bytes,9,opt,name=bill" json:"bill,omitempty"`
	DeskId           *int32                `protobuf:"varint,10,opt,name=deskId" json:"deskId,omitempty"`
	RoomId           *int32                `protobuf:"varint,11,opt,name=roomId" json:"roomId,omitempty"`
	Coin             *int64                `protobuf:"varint,12,opt,name=coin" json:"coin,omitempty"`
	Statistics       *DdzSrvUserStatistics `protobuf:"bytes,13,opt,name=statistics" json:"statistics,omitempty"`
	LaDaoStatus      *int32                `protobuf:"varint,14,opt,name=laDaoStatus" json:"laDaoStatus,omitempty"`
	PlayRate         *int32                `protobuf:"varint,15,opt,name=playRate" json:"playRate,omitempty"`
	Bomb             *int32                `protobuf:"varint,16,opt,name=bomb" json:"bomb,omitempty"`
	JiaoScore        *int32                `protobuf:"varint,17,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUser) Reset()                    { *m = DdzSrvUser{} }
func (m *DdzSrvUser) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUser) ProtoMessage()               {}
func (*DdzSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *DdzSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzSrvUser) GetGameData() *DdzSrvGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *DdzSrvUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *DdzSrvUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *DdzSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *DdzSrvUser) GetQiangDiZhuStatus() int32 {
	if m != nil && m.QiangDiZhuStatus != nil {
		return *m.QiangDiZhuStatus
	}
	return 0
}

func (m *DdzSrvUser) GetJiabeiStatus() int32 {
	if m != nil && m.JiabeiStatus != nil {
		return *m.JiabeiStatus
	}
	return 0
}

func (m *DdzSrvUser) GetBill() *DdzSrvBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *DdzSrvUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzSrvUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *DdzSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvUser) GetStatistics() *DdzSrvUserStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *DdzSrvUser) GetLaDaoStatus() int32 {
	if m != nil && m.LaDaoStatus != nil {
		return *m.LaDaoStatus
	}
	return 0
}

func (m *DdzSrvUser) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUser) GetBomb() int32 {
	if m != nil && m.Bomb != nil {
		return *m.Bomb
	}
	return 0
}

func (m *DdzSrvUser) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

// room
type DdzSrvRoom struct {
	RoomId           *int32 `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DdzSrvRoom) Reset()                    { *m = DdzSrvRoom{} }
func (m *DdzSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvRoom) ProtoMessage()               {}
func (*DdzSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *DdzSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

// 备份专用...
type DdzSrvBak struct {
	Desk             *DdzSrvDesk   `protobuf:"bytes,1,opt,name=desk" json:"desk,omitempty"`
	Users            []*DdzSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DdzSrvBak) Reset()                    { *m = DdzSrvBak{} }
func (m *DdzSrvBak) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBak) ProtoMessage()               {}
func (*DdzSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *DdzSrvBak) GetDesk() *DdzSrvDesk {
	if m != nil {
		return m.Desk
	}
	return nil
}

func (m *DdzSrvBak) GetUsers() []*DdzSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*DdzSrvOutPokerPais)(nil), "ddproto.ddz_srv_outPokerPais")
	proto.RegisterType((*DdzSrvDeskTongJi)(nil), "ddproto.ddz_srv_deskTongJi")
	proto.RegisterType((*DdzSrvDesk)(nil), "ddproto.ddz_srv_desk")
	proto.RegisterType((*DdzSrvGameData)(nil), "ddproto.ddz_srv_gameData")
	proto.RegisterType((*DdzSrvBillBean)(nil), "ddproto.ddz_srv_billBean")
	proto.RegisterType((*DdzSrvBill)(nil), "ddproto.ddz_srv_bill")
	proto.RegisterType((*DdzSrvUserStatistics)(nil), "ddproto.ddz_srv_userStatistics")
	proto.RegisterType((*DdzSrvUser)(nil), "ddproto.ddz_srv_user")
	proto.RegisterType((*DdzSrvRoom)(nil), "ddproto.ddz_srv_room")
	proto.RegisterType((*DdzSrvBak)(nil), "ddproto.ddz_srv_bak")
	proto.RegisterEnum("ddproto.EDDZTYPE", EDDZTYPE_name, EDDZTYPE_value)
}

var fileDescriptor7 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdb, 0x72, 0x1b, 0x45,
	0x10, 0x45, 0x96, 0x65, 0x4b, 0x2d, 0xcb, 0x56, 0x36, 0xb6, 0x33, 0x4e, 0xb8, 0xa4, 0x16, 0x1e,
	0x52, 0x04, 0x0c, 0x31, 0x6f, 0xbc, 0x61, 0xc9, 0x10, 0xa7, 0x44, 0x4a, 0x89, 0x6d, 0x20, 0xe6,
	0x41, 0x35, 0xd2, 0x0e, 0xce, 0x44, 0xbb, 0x3b, 0xca, 0x5e, 0x6c, 0xec, 0x6f, 0xe0, 0x3b, 0xf8,
	0x2c, 0xfe, 0x84, 0x2a, 0xba, 0x7b, 0x66, 0x24, 0x81, 0x54, 0xd8, 0x4f, 0xd6, 0x9e, 0xe9, 0x99,
	0xee, 0x3e, 0x7d, 0xfa, 0x18, 0xda, 0x51, 0x74, 0x33, 0xc8, 0x55, 0x76, 0xa9, 0xb2, 0xfd, 0x49,
	0x66, 0x0a, 0x13, 0xac, 0x47, 0x11, 0xff, 0x78, 0xb8, 0x37, 0x32, 0x49, 0x62, 0x52, 0x77, 0x3a,
	0x98, 0x98, 0xb1, 0x8f, 0x09, 0xff, 0xaa, 0xc0, 0x36, 0x5f, 0xcc, 0x2e, 0x07, 0xa6, 0x2c, 0xfa,
	0x74, 0xd4, 0x97, 0x3a, 0x0f, 0xda, 0x50, 0x1f, 0xab, 0xeb, 0x9f, 0x64, 0x5c, 0x2a, 0x51, 0x79,
	0x5c, 0x79, 0x52, 0x0b, 0xbe, 0x82, 0xc6, 0xc4, 0x1f, 0x8b, 0x95, 0xc7, 0xd5, 0x27, 0xcd, 0x83,
	0x0f, 0xf7, 0x5d, 0x8a, 0x7d, 0x9f, 0x01, 0x9f, 0xf1, 0x41, 0xc1, 0x06, 0xac, 0x16, 0xd7, 0x13,
	0x25, 0xaa, 0x7c, 0x7d, 0x13, 0xd6, 0x74, 0x7e, 0x68, 0x92, 0xa1, 0x58, 0xc5, 0xef, 0x7a, 0x10,
	0x00, 0x8c, 0x4c, 0x99, 0x16, 0xdd, 0x52, 0xdf, 0x68, 0x51, 0xe3, 0x98, 0x1d, 0x68, 0x31, 0x76,
	0x22, 0xd3, 0x9b, 0xb7, 0x32, 0xbd, 0x10, 0x6b, 0x0c, 0x6f, 0xc3, 0x86, 0x85, 0xb5, 0x45, 0xd7,
	0xff, 0x85, 0xbe, 0x71, 0x68, 0xdd, 0xa7, 0x29, 0xb1, 0xcf, 0xe3, 0x48, 0x34, 0xf0, 0xbb, 0x15,
	0xfe, 0x0a, 0x81, 0xef, 0x2f, 0x52, 0xf9, 0xf8, 0xd4, 0xa4, 0x17, 0x2f, 0x74, 0xf0, 0x05, 0xd4,
	0x86, 0x58, 0x4a, 0x8e, 0xad, 0x51, 0x1f, 0x1f, 0x4d, 0xfb, 0x58, 0xca, 0xc5, 0x03, 0xd8, 0xe2,
	0x4c, 0xaf, 0x34, 0xe6, 0xe9, 0xea, 0xf3, 0xb7, 0x25, 0xf6, 0x8f, 0xc9, 0xc2, 0x3f, 0x6a, 0xb0,
	0x31, 0xff, 0x3a, 0x65, 0xa7, 0xbf, 0x98, 0xdd, 0x72, 0xd6, 0x84, 0x2a, 0xb2, 0xc8, 0xd1, 0x8d,
	0x60, 0x17, 0x36, 0xa9, 0xb4, 0x0e, 0x3d, 0xd5, 0xd3, 0x89, 0x2e, 0x1c, 0x33, 0xcf, 0xa0, 0x29,
	0xe3, 0xd8, 0xa7, 0x43, 0x7a, 0x6e, 0xa7, 0xf6, 0x6b, 0x80, 0x48, 0x4f, 0x6f, 0xd4, 0xee, 0x70,
	0xe3, 0x4b, 0x58, 0xa3, 0x9e, 0x30, 0x9a, 0x38, 0xbd, 0xb5, 0xe5, 0x16, 0xd4, 0xcc, 0x55, 0xaa,
	0x32, 0xe6, 0xba, 0x45, 0x5c, 0x47, 0x48, 0x73, 0x89, 0x67, 0x67, 0xd8, 0x02, 0x73, 0xdd, 0x0a,
	0xee, 0x43, 0x33, 0x22, 0x36, 0xce, 0xe6, 0x08, 0xa7, 0x50, 0x39, 0x2a, 0xf4, 0xa5, 0x72, 0x28,
	0x30, 0xfa, 0x14, 0xd6, 0x0a, 0xa4, 0xfe, 0x9d, 0x16, 0x4d, 0x4e, 0xff, 0x68, 0x21, 0xfd, 0xdc,
	0x74, 0xee, 0x41, 0x63, 0x28, 0x73, 0x65, 0xc5, 0xb7, 0x81, 0xf1, 0x55, 0xe2, 0xee, 0xbd, 0x26,
	0xf2, 0xb1, 0x08, 0x8b, 0xb7, 0x18, 0x47, 0x99, 0x5e, 0xe9, 0xd4, 0x22, 0x9b, 0x8c, 0x6c, 0x01,
	0xea, 0xfe, 0xe6, 0x94, 0x84, 0xb7, 0xc5, 0xf4, 0x62, 0xc8, 0x6b, 0x63, 0x12, 0x46, 0xda, 0x8c,
	0x60, 0xdd, 0x87, 0x46, 0x66, 0x51, 0xce, 0xa3, 0x10, 0xf7, 0xbc, 0x70, 0x3a, 0x72, 0xf2, 0xa3,
	0xfc, 0x5d, 0x04, 0xfc, 0x0e, 0x16, 0x71, 0x9c, 0xbf, 0xd0, 0xd2, 0x7c, 0xaf, 0x52, 0x71, 0x9f,
	0x25, 0x8b, 0x21, 0x19, 0xbe, 0x84, 0x4d, 0x6d, 0xfb, 0x97, 0x7f, 0x33, 0xa6, 0x78, 0x2d, 0x0b,
	0x25, 0x76, 0x3c, 0x32, 0x89, 0xe5, 0x35, 0x23, 0xbb, 0x8c, 0xa0, 0xcc, 0x7f, 0x90, 0x89, 0x3a,
	0x29, 0x64, 0x51, 0xe6, 0xe2, 0x81, 0x57, 0xae, 0x4e, 0x75, 0x41, 0x55, 0x75, 0x8c, 0x4e, 0x85,
	0xe0, 0x84, 0x24, 0xfe, 0x32, 0xcb, 0xfa, 0x78, 0xdf, 0xd6, 0xb5, 0xc7, 0xc1, 0xd8, 0x79, 0x61,
	0x0a, 0x19, 0xcf, 0xf0, 0x87, 0x8c, 0x63, 0x7d, 0x94, 0x4a, 0x65, 0x2f, 0xcb, 0x44, 0x3c, 0x62,
	0x39, 0xfe, 0x59, 0x71, 0x2e, 0x80, 0x74, 0x5e, 0x60, 0xd2, 0xae, 0x2c, 0x24, 0x49, 0x05, 0xd7,
	0x23, 0xe2, 0xd1, 0x7a, 0xbd, 0xff, 0xbf, 0x54, 0x9e, 0x01, 0x58, 0xa9, 0xf4, 0x74, 0x5e, 0xb8,
	0x4d, 0xbf, 0x45, 0x2e, 0x33, 0x75, 0x55, 0xef, 0xa0, 0xae, 0xf0, 0xd5, 0xac, 0xce, 0xa1, 0x8e,
	0xe3, 0x43, 0x25, 0x53, 0x72, 0x8b, 0x11, 0x91, 0x51, 0xf1, 0x73, 0x8d, 0x4d, 0xce, 0x1a, 0xe2,
	0xed, 0x69, 0xd1, 0x5c, 0x71, 0xd2, 0x0c, 0x54, 0x19, 0xc0, 0x0b, 0xa8, 0x99, 0x11, 0xdb, 0x49,
	0x23, 0xec, 0xcd, 0x36, 0x91, 0x9e, 0x74, 0xe1, 0x9d, 0xd9, 0x8b, 0x4f, 0xa1, 0xee, 0x73, 0xb9,
	0x9e, 0xf6, 0x16, 0x8a, 0xf4, 0x01, 0x61, 0x02, 0xbb, 0x1e, 0xa3, 0x95, 0xa5, 0xe9, 0x21, 0x17,
	0x7a, 0x94, 0x07, 0xfb, 0xb0, 0x8e, 0xce, 0xc1, 0xaf, 0xdc, 0xc9, 0x3b, 0xb0, 0x11, 0x9e, 0xda,
	0xcf, 0x3a, 0xb5, 0xa6, 0x41, 0x83, 0xb3, 0x16, 0x80, 0xfd, 0x59, 0x07, 0x08, 0xff, 0x5e, 0x99,
	0x55, 0x4f, 0xf9, 0xe6, 0x5c, 0xac, 0xe2, 0xd6, 0xa7, 0xee, 0x07, 0xca, 0xaf, 0x2c, 0x2b, 0x7e,
	0x3a, 0x71, 0xbc, 0x9c, 0x5b, 0xb9, 0x59, 0x7f, 0x41, 0x2a, 0xd0, 0x79, 0x33, 0x25, 0xc7, 0xce,
	0x7a, 0x19, 0xe8, 0x29, 0x79, 0xa9, 0xd8, 0x77, 0xeb, 0x81, 0x80, 0xf6, 0xfb, 0xa9, 0xb7, 0x39,
	0xa9, 0x4e, 0x4d, 0xf6, 0x9d, 0x96, 0x43, 0xa5, 0x1d, 0x6a, 0x4d, 0xf6, 0x53, 0x58, 0x25, 0xaa,
	0x78, 0xe3, 0x9b, 0x07, 0x3b, 0x4b, 0x79, 0x9c, 0xf3, 0x42, 0xf0, 0x0b, 0xe6, 0xb6, 0xa7, 0xc9,
	0xdf, 0x7e, 0xe0, 0x76, 0xc1, 0xbf, 0x01, 0xc8, 0xa7, 0x2c, 0xf3, 0x72, 0x37, 0x0f, 0x3e, 0x59,
	0x78, 0xf8, 0x3f, 0xc3, 0xc0, 0x45, 0x8e, 0x65, 0x57, 0x1a, 0x57, 0xdc, 0xe6, 0xc2, 0x0e, 0x6e,
	0xf9, 0x4c, 0xe4, 0xf6, 0x6e, 0xfb, 0x91, 0x7f, 0x6c, 0xc9, 0x9c, 0x8c, 0x4c, 0xa6, 0xec, 0xee,
	0x87, 0x1f, 0xcf, 0xe8, 0xa7, 0x12, 0xe7, 0x4a, 0x65, 0x1b, 0x0f, 0x7f, 0x41, 0xa3, 0xf3, 0xad,
	0xc9, 0x31, 0xb5, 0x4f, 0x9d, 0xf1, 0xe1, 0xb2, 0xf6, 0xf9, 0x5f, 0xc1, 0x67, 0x50, 0xa3, 0x6a,
	0xfd, 0xbf, 0xca, 0x9d, 0xa5, 0xbd, 0x7c, 0xfe, 0x2d, 0xd4, 0x8f, 0xba, 0xdd, 0xf3, 0xd3, 0x37,
	0xfd, 0x23, 0x9c, 0xc2, 0x36, 0xfe, 0x1c, 0xd0, 0xef, 0xc1, 0xf3, 0xb3, 0xef, 0x5e, 0xf6, 0x8e,
	0xba, 0xc7, 0xe7, 0xcf, 0xcf, 0xda, 0x15, 0x9c, 0x42, 0x7b, 0x7a, 0x72, 0x72, 0xdc, 0xa1, 0xc3,
	0xf6, 0x4a, 0xff, 0x83, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x40, 0x98, 0x5a, 0x27, 0xf4, 0x07,
	0x00, 0x00,
}