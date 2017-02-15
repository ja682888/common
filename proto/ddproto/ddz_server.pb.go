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

// Ignoring public import of ddz_base_roomTypeInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_commonRateInfo from ddz_base.proto

// Ignoring public import of ddz_base_timerInfo from ddz_base.proto

// Ignoring public import of ddz_base_deskInfo from ddz_base.proto

// Ignoring public import of ddz_enum_protoId from ddz_base.proto

// Ignoring public import of ddz_enum_errorCode from ddz_base.proto

// Ignoring public import of ddz_enum_paiType from ddz_base.proto

// Ignoring public import of ddz_enum_actType from ddz_base.proto

// Ignoring public import of ddz_enum_gameStatus from ddz_base.proto

// Ignoring public import of ddz_enum_playerStatus from ddz_base.proto

// Ignoring public import of ddz_enum_roomType from ddz_base.proto

// Ignoring public import of ddz_enum_enterType from ddz_base.proto

// Ignoring public import of ddz_enum_coinRoomLevel from ddz_base.proto

// Ignoring public import of ddz_enum_deskGameStatus from ddz_base.proto

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
func (*DdzSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

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
func (*DdzSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

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
	AllPokerPai      []*CommonSrvPokerPai   `protobuf:"bytes,1,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*CommonSrvPokerPai   `protobuf:"bytes,2,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *DdzSrvOutPokerPais    `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	DizhuPaiUser     *uint32                `protobuf:"varint,4,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	DiZhuUserId      *uint32                `protobuf:"varint,5,opt,name=diZhuUserId" json:"diZhuUserId,omitempty"`
	ActiveUserId     *uint32                `protobuf:"varint,6,opt,name=activeUserId" json:"activeUserId,omitempty"`
	Tongji           *DdzSrvDeskTongJi      `protobuf:"bytes,7,opt,name=tongji" json:"tongji,omitempty"`
	DdzType          *int32                 `protobuf:"varint,8,opt,name=ddzType" json:"ddzType,omitempty"`
	RoomType         *int32                 `protobuf:"varint,9,opt,name=RoomType" json:"RoomType,omitempty"`
	BoardsCount      *int32                 `protobuf:"varint,10,opt,name=BoardsCount" json:"BoardsCount,omitempty"`
	CapMax           *int64                 `protobuf:"varint,11,opt,name=CapMax" json:"CapMax,omitempty"`
	IsJiaoFen        *bool                  `protobuf:"varint,12,opt,name=IsJiaoFen" json:"IsJiaoFen,omitempty"`
	CommonRateInfo   *DdzBaseCommonRateInfo `protobuf:"bytes,13,opt,name=commonRateInfo" json:"commonRateInfo,omitempty"`
	PlayRate         *int32                 `protobuf:"varint,14,opt,name=playRate" json:"playRate,omitempty"`
	GameStatus       *int32                 `protobuf:"varint,15,opt,name=GameStatus" json:"GameStatus,omitempty"`
	InitRoomCoin     *int64                 `protobuf:"varint,16,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32                 `protobuf:"varint,17,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32                 `protobuf:"varint,18,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	PlayerNum        *int32                 `protobuf:"varint,19,opt,name=playerNum" json:"playerNum,omitempty"`
	TimeOutSecond    *int32                 `protobuf:"varint,20,opt,name=timeOutSecond" json:"timeOutSecond,omitempty"`
	UserMinCoin      *int64                 `protobuf:"varint,21,opt,name=userMinCoin" json:"userMinCoin,omitempty"`
	UserMaxCoin      *int64                 `protobuf:"varint,22,opt,name=userMaxCoin" json:"userMaxCoin,omitempty"`
	CoinTicket       *int64                 `protobuf:"varint,23,opt,name=coinTicket" json:"coinTicket,omitempty"`
	CoinRoomLv       *DdzEnumCoinRoomLevel  `protobuf:"varint,24,opt,name=coinRoomLv,enum=ddproto.DdzEnumCoinRoomLevel" json:"coinRoomLv,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *DdzSrvDesk) Reset()                    { *m = DdzSrvDesk{} }
func (m *DdzSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDesk) ProtoMessage()               {}
func (*DdzSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

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

func (m *DdzSrvDesk) GetCommonRateInfo() *DdzBaseCommonRateInfo {
	if m != nil {
		return m.CommonRateInfo
	}
	return nil
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

func (m *DdzSrvDesk) GetTimeOutSecond() int32 {
	if m != nil && m.TimeOutSecond != nil {
		return *m.TimeOutSecond
	}
	return 0
}

func (m *DdzSrvDesk) GetUserMinCoin() int64 {
	if m != nil && m.UserMinCoin != nil {
		return *m.UserMinCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetUserMaxCoin() int64 {
	if m != nil && m.UserMaxCoin != nil {
		return *m.UserMaxCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *DdzSrvDesk) GetCoinRoomLv() DdzEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLv != nil {
		return *m.CoinRoomLv
	}
	return DdzEnumCoinRoomLevel_LV1
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
func (*DdzSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

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
func (*DdzSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

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
func (*DdzSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

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

type DdzSrvUserStatisticsRound struct {
	// 玩家每一局的统计信息
	Round            *int32                `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	WinCoin          *int64                `protobuf:"varint,2,opt,name=winCoin" json:"winCoin,omitempty"`
	CountBomb        *int32                `protobuf:"varint,3,opt,name=countBomb" json:"countBomb,omitempty"`
	BomBean          []*DdzSrvOutPokerPais `protobuf:"bytes,4,rep,name=bomBean" json:"bomBean,omitempty"`
	PlayRate         *int32                `protobuf:"varint,5,opt,name=playRate" json:"playRate,omitempty"`
	IsSpring         *bool                 `protobuf:"varint,6,opt,name=isSpring" json:"isSpring,omitempty"`
	IsDizhu          *bool                 `protobuf:"varint,7,opt,name=isDizhu" json:"isDizhu,omitempty"`
	IsWin            *bool                 `protobuf:"varint,8,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUserStatisticsRound) Reset()                    { *m = DdzSrvUserStatisticsRound{} }
func (m *DdzSrvUserStatisticsRound) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatisticsRound) ProtoMessage()               {}
func (*DdzSrvUserStatisticsRound) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *DdzSrvUserStatisticsRound) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetBomBean() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.BomBean
	}
	return nil
}

func (m *DdzSrvUserStatisticsRound) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsDizhu() bool {
	if m != nil && m.IsDizhu != nil {
		return *m.IsDizhu
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

type DdzSrvUserStatistics struct {
	// 玩家统计信息
	RoundBean        []*DdzSrvUserStatisticsRound `protobuf:"bytes,1,rep,name=roundBean" json:"roundBean,omitempty"`
	CountWin         *int32                       `protobuf:"varint,2,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32                       `protobuf:"varint,3,opt,name=countLose" json:"countLose,omitempty"`
	CountSpring      *int32                       `protobuf:"varint,4,opt,name=countSpring" json:"countSpring,omitempty"`
	CountDizhu       *int32                       `protobuf:"varint,5,opt,name=countDizhu" json:"countDizhu,omitempty"`
	CountBomb        *int32                       `protobuf:"varint,6,opt,name=countBomb" json:"countBomb,omitempty"`
	MaxWinCoin       *int64                       `protobuf:"varint,7,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *DdzSrvUserStatistics) Reset()                    { *m = DdzSrvUserStatistics{} }
func (m *DdzSrvUserStatistics) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatistics) ProtoMessage()               {}
func (*DdzSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *DdzSrvUserStatistics) GetRoundBean() []*DdzSrvUserStatisticsRound {
	if m != nil {
		return m.RoundBean
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

func (m *DdzSrvUserStatistics) GetCountSpring() int32 {
	if m != nil && m.CountSpring != nil {
		return *m.CountSpring
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountDizhu() int32 {
	if m != nil && m.CountDizhu != nil {
		return *m.CountDizhu
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetMaxWinCoin() int64 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

// user
type DdzSrvUser struct {
	UserId              *uint32               `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData            *DdzSrvGameData       `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	StatusHLQiang       *int32                `protobuf:"varint,3,opt,name=statusHLQiang" json:"statusHLQiang,omitempty"`
	StatusHLJiao        *int32                `protobuf:"varint,4,opt,name=statusHLJiao" json:"statusHLJiao,omitempty"`
	StatusHLJiaBei      *int32                `protobuf:"varint,5,opt,name=statusHLJiaBei" json:"statusHLJiaBei,omitempty"`
	StatusSCMenZhua     *int32                `protobuf:"varint,6,opt,name=statusSCMenZhua" json:"statusSCMenZhua,omitempty"`
	StatusSCZhuaPai     *int32                `protobuf:"varint,7,opt,name=statusSCZhuaPai" json:"statusSCZhuaPai,omitempty"`
	StatusSCLaDao       *int32                `protobuf:"varint,8,opt,name=statusSCLaDao" json:"statusSCLaDao,omitempty"`
	StatusJDJiao        *int32                `protobuf:"varint,9,opt,name=statusJDJiao" json:"statusJDJiao,omitempty"`
	StatusShowPokers    *int32                `protobuf:"varint,10,opt,name=statusShowPokers" json:"statusShowPokers,omitempty"`
	IsShowPokers        *bool                 `protobuf:"varint,11,opt,name=isShowPokers" json:"isShowPokers,omitempty"`
	Bill                *DdzSrvBill           `protobuf:"bytes,12,opt,name=bill" json:"bill,omitempty"`
	Coin                *int64                `protobuf:"varint,13,opt,name=coin" json:"coin,omitempty"`
	Score               *int64                `protobuf:"varint,14,opt,name=score" json:"score,omitempty"`
	Statistics          *DdzSrvUserStatistics `protobuf:"bytes,15,opt,name=statistics" json:"statistics,omitempty"`
	PlayRate            *int32                `protobuf:"varint,16,opt,name=playRate" json:"playRate,omitempty"`
	JiaoScore           *int32                `protobuf:"varint,17,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	TimeOutCount        *int32                `protobuf:"varint,18,opt,name=timeOutCount" json:"timeOutCount,omitempty"`
	IsAgent             *bool                 `protobuf:"varint,19,opt,name=isAgent" json:"isAgent,omitempty"`
	Sex                 *int32                `protobuf:"varint,20,opt,name=sex" json:"sex,omitempty"`
	RoomCard            *int32                `protobuf:"varint,21,opt,name=roomCard" json:"roomCard,omitempty"`
	StatusApplyDissolve *int32                `protobuf:"varint,22,opt,name=statusApplyDissolve" json:"statusApplyDissolve,omitempty"`
	XXX_unrecognized    []byte                `json:"-"`
}

func (m *DdzSrvUser) Reset()                    { *m = DdzSrvUser{} }
func (m *DdzSrvUser) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUser) ProtoMessage()               {}
func (*DdzSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

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

func (m *DdzSrvUser) GetStatusHLQiang() int32 {
	if m != nil && m.StatusHLQiang != nil {
		return *m.StatusHLQiang
	}
	return 0
}

func (m *DdzSrvUser) GetStatusHLJiao() int32 {
	if m != nil && m.StatusHLJiao != nil {
		return *m.StatusHLJiao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusHLJiaBei() int32 {
	if m != nil && m.StatusHLJiaBei != nil {
		return *m.StatusHLJiaBei
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCMenZhua() int32 {
	if m != nil && m.StatusSCMenZhua != nil {
		return *m.StatusSCMenZhua
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCZhuaPai() int32 {
	if m != nil && m.StatusSCZhuaPai != nil {
		return *m.StatusSCZhuaPai
	}
	return 0
}

func (m *DdzSrvUser) GetStatusSCLaDao() int32 {
	if m != nil && m.StatusSCLaDao != nil {
		return *m.StatusSCLaDao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusJDJiao() int32 {
	if m != nil && m.StatusJDJiao != nil {
		return *m.StatusJDJiao
	}
	return 0
}

func (m *DdzSrvUser) GetStatusShowPokers() int32 {
	if m != nil && m.StatusShowPokers != nil {
		return *m.StatusShowPokers
	}
	return 0
}

func (m *DdzSrvUser) GetIsShowPokers() bool {
	if m != nil && m.IsShowPokers != nil {
		return *m.IsShowPokers
	}
	return false
}

func (m *DdzSrvUser) GetBill() *DdzSrvBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *DdzSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvUser) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *DdzSrvUser) GetStatistics() *DdzSrvUserStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *DdzSrvUser) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUser) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

func (m *DdzSrvUser) GetTimeOutCount() int32 {
	if m != nil && m.TimeOutCount != nil {
		return *m.TimeOutCount
	}
	return 0
}

func (m *DdzSrvUser) GetIsAgent() bool {
	if m != nil && m.IsAgent != nil {
		return *m.IsAgent
	}
	return false
}

func (m *DdzSrvUser) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *DdzSrvUser) GetRoomCard() int32 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *DdzSrvUser) GetStatusApplyDissolve() int32 {
	if m != nil && m.StatusApplyDissolve != nil {
		return *m.StatusApplyDissolve
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
func (*DdzSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

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
func (*DdzSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

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
	proto.RegisterType((*DdzSrvUserStatisticsRound)(nil), "ddproto.ddz_srv_userStatisticsRound")
	proto.RegisterType((*DdzSrvUserStatistics)(nil), "ddproto.ddz_srv_userStatistics")
	proto.RegisterType((*DdzSrvUser)(nil), "ddproto.ddz_srv_user")
	proto.RegisterType((*DdzSrvRoom)(nil), "ddproto.ddz_srv_room")
	proto.RegisterType((*DdzSrvBak)(nil), "ddproto.ddz_srv_bak")
}

var fileDescriptor14 = []byte{
	// 1081 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x72, 0x1b, 0x45,
	0x10, 0x46, 0x91, 0x56, 0x96, 0x46, 0x96, 0x22, 0xaf, 0x7f, 0x32, 0x89, 0xf9, 0x71, 0x2d, 0x39,
	0xa4, 0x2a, 0x60, 0x88, 0x39, 0xc0, 0x35, 0x96, 0x0a, 0xb0, 0x4b, 0x06, 0xc5, 0x72, 0x30, 0x3f,
	0x07, 0xd5, 0x58, 0x3b, 0xd8, 0x13, 0xad, 0x76, 0x54, 0x3b, 0xbb, 0x8a, 0xed, 0x87, 0xe1, 0xc8,
	0x8d, 0xb7, 0xa0, 0x78, 0x0c, 0x9e, 0x85, 0xee, 0xde, 0x19, 0x69, 0x85, 0x55, 0xc4, 0x9c, 0xac,
	0xf9, 0xb6, 0x67, 0xfa, 0xeb, 0x9e, 0xaf, 0xbf, 0x31, 0x6b, 0x87, 0xe1, 0xed, 0xd0, 0xc8, 0x64,
	0x26, 0x93, 0xfd, 0x69, 0xa2, 0x53, 0xed, 0xaf, 0x85, 0x21, 0xfd, 0x78, 0xf2, 0x78, 0xa4, 0x27,
	0x13, 0x1d, 0xdb, 0xaf, 0xc3, 0xa9, 0x1e, 0xbb, 0x98, 0x27, 0x2d, 0xdc, 0x75, 0x21, 0x8c, 0xcc,
	0xd7, 0xc1, 0xdf, 0x25, 0xb6, 0x45, 0x07, 0x25, 0xb3, 0xa1, 0xce, 0xd2, 0x3e, 0x86, 0xf6, 0x85,
	0x32, 0x7e, 0x9b, 0xd5, 0xc6, 0xf2, 0xe6, 0x07, 0x11, 0x65, 0x92, 0x97, 0xf6, 0x4a, 0xcf, 0x3c,
	0xff, 0x33, 0x56, 0x9f, 0xba, 0xcf, 0xfc, 0xc1, 0x5e, 0xf9, 0x59, 0xe3, 0xe0, 0xfd, 0x7d, 0x9b,
	0x72, 0xdf, 0x65, 0x84, 0x63, 0x5c, 0x90, 0xbf, 0xce, 0x2a, 0xe9, 0xcd, 0x54, 0xf2, 0x32, 0x6d,
	0x6f, 0xb1, 0xaa, 0x32, 0x87, 0x7a, 0x72, 0xc1, 0x2b, 0xb0, 0xae, 0xf9, 0x3e, 0x63, 0x23, 0x9d,
	0xc5, 0x69, 0x37, 0x53, 0xb7, 0x8a, 0x7b, 0x14, 0xb3, 0xcd, 0x9a, 0x84, 0x0d, 0x44, 0x7c, 0x7b,
	0x25, 0xe2, 0x4b, 0x5e, 0x25, 0x78, 0x8b, 0xad, 0xe7, 0xb0, 0xca, 0xd1, 0xb5, 0x25, 0xf4, 0x27,
	0x8b, 0xd6, 0x5c, 0x9a, 0x0c, 0xea, 0x3e, 0x0a, 0x79, 0x1d, 0xd6, 0xcd, 0xe0, 0x17, 0xe6, 0xbb,
	0xfa, 0x42, 0x69, 0xc6, 0x67, 0x3a, 0xbe, 0x3c, 0x56, 0xfe, 0x27, 0xcc, 0xbb, 0x00, 0x2a, 0x06,
	0x4a, 0xc3, 0x3a, 0x3e, 0x98, 0xd7, 0xb1, 0xb2, 0x17, 0x8f, 0xd8, 0x43, 0xca, 0xf4, 0x4a, 0x41,
	0x9e, 0xae, 0xfa, 0xf9, 0x2a, 0x83, 0xfa, 0x21, 0x59, 0xf0, 0x87, 0xc7, 0xd6, 0x8b, 0xa7, 0xfb,
	0x2f, 0x58, 0x43, 0x44, 0x91, 0xdb, 0x69, 0x4f, 0xff, 0xef, 0x2e, 0x7d, 0xce, 0x58, 0xa8, 0xe6,
	0x3b, 0xee, 0xd3, 0xd7, 0x4f, 0x59, 0x15, 0xe9, 0x41, 0x34, 0x76, 0xf6, 0x9d, 0xec, 0xa1, 0x4f,
	0x21, 0xb4, 0x28, 0x83, 0xc5, 0x6b, 0xe8, 0x0c, 0xb5, 0xbf, 0xe9, 0x6f, 0xb2, 0x46, 0x88, 0x95,
	0xbc, 0xce, 0x9b, 0xe5, 0x11, 0x08, 0xa1, 0x62, 0x94, 0xaa, 0x99, 0xb4, 0x68, 0x95, 0xd0, 0xe7,
	0xac, 0x9a, 0x42, 0xdb, 0xde, 0x28, 0x6a, 0x7c, 0xe3, 0x60, 0xf7, 0x4e, 0xbe, 0x42, 0x67, 0x1f,
	0x32, 0x90, 0xe1, 0xed, 0x19, 0xde, 0x7b, 0x7e, 0x21, 0x20, 0xa4, 0x53, 0xad, 0x27, 0x84, 0xd4,
	0x09, 0x81, 0xd4, 0x87, 0x5a, 0x24, 0xa1, 0xe9, 0x60, 0x53, 0x39, 0x73, 0xf7, 0xd6, 0x11, 0xd3,
	0x13, 0x71, 0xcd, 0x1b, 0xb0, 0x2e, 0xfb, 0x1b, 0xac, 0x7e, 0x64, 0x8e, 0x95, 0xd0, 0x5f, 0xcb,
	0x98, 0xaf, 0x93, 0x62, 0xbe, 0x62, 0xad, 0xbc, 0x1d, 0xa7, 0x22, 0x95, 0x47, 0xf1, 0xaf, 0x9a,
	0x37, 0x89, 0xcf, 0xde, 0x12, 0x1f, 0x14, 0xf7, 0x70, 0x39, 0x0e, 0x39, 0x4c, 0x23, 0x71, 0x83,
	0x6b, 0xde, 0xa2, 0x74, 0xa0, 0xbe, 0x6f, 0xc4, 0x44, 0x0e, 0x52, 0x91, 0x66, 0x86, 0x3f, 0x74,
	0x82, 0x52, 0xb1, 0x4a, 0x91, 0x6d, 0x47, 0xab, 0x98, 0xb7, 0x89, 0x08, 0x6a, 0x32, 0x4b, 0x92,
	0x3e, 0xec, 0xcf, 0xf9, 0x6e, 0x50, 0xf0, 0x0e, 0x6b, 0xa5, 0x3a, 0x15, 0xd1, 0x02, 0xf7, 0x09,
	0x07, 0xde, 0x98, 0x4a, 0x26, 0xdf, 0x65, 0x13, 0xbe, 0xe9, 0x54, 0x9d, 0xaa, 0x89, 0xfc, 0x3e,
	0x4b, 0x07, 0x72, 0xa4, 0xe3, 0x90, 0x6f, 0xb9, 0x36, 0xa0, 0x52, 0x4f, 0x54, 0x4c, 0xd9, 0xb6,
	0x29, 0x9b, 0x03, 0xc5, 0x35, 0x81, 0x3b, 0x04, 0xd2, 0xa8, 0xa8, 0xf8, 0x4c, 0x8d, 0xc6, 0x32,
	0xe5, 0x8f, 0x08, 0xfb, 0x22, 0xc7, 0x90, 0x6c, 0x6f, 0xc6, 0x39, 0x60, 0xad, 0x83, 0x8f, 0x96,
	0x1a, 0x21, 0xe3, 0x6c, 0x32, 0x9c, 0xc7, 0xc8, 0x99, 0x8c, 0x82, 0xdf, 0x4b, 0xd6, 0x36, 0xe0,
	0xce, 0x2e, 0xa1, 0xfc, 0xae, 0x48, 0x05, 0x0a, 0x10, 0xe6, 0x27, 0x24, 0xc1, 0x98, 0x7b, 0x49,
	0xf6, 0x05, 0x63, 0xb9, 0x00, 0x7b, 0xca, 0xa4, 0x56, 0xb2, 0xef, 0x10, 0xe1, 0xff, 0xd3, 0x6c,
	0xf0, 0x6a, 0xc1, 0xf3, 0x42, 0x45, 0xd1, 0xa1, 0x14, 0x31, 0xda, 0x09, 0x56, 0x43, 0x6e, 0x54,
	0xc6, 0x2b, 0x8d, 0xb4, 0x21, 0xa1, 0xd2, 0x30, 0x36, 0x51, 0x79, 0x6f, 0x55, 0x4c, 0x40, 0x99,
	0x00, 0xd8, 0x00, 0xc2, 0x1c, 0x91, 0xe0, 0xeb, 0x41, 0x6f, 0x31, 0xaa, 0x78, 0xa4, 0x0d, 0xef,
	0x2c, 0x4e, 0x7c, 0xce, 0x6a, 0x2e, 0x97, 0xad, 0xe9, 0xf1, 0x1d, 0x92, 0x2e, 0x20, 0xf8, 0xb3,
	0xc4, 0x76, 0x1d, 0x88, 0x17, 0x86, 0x42, 0x82, 0x66, 0xa8, 0x91, 0x39, 0x05, 0x31, 0x84, 0x7e,
	0x93, 0x79, 0x09, 0xfe, 0xb0, 0xde, 0x59, 0x48, 0xf6, 0xc0, 0xc9, 0x9b, 0x2c, 0x85, 0x0c, 0x31,
	0x37, 0xc8, 0x7d, 0xb6, 0x06, 0x9e, 0x44, 0xe9, 0x2b, 0xf7, 0x69, 0x69, 0x51, 0xd4, 0x9e, 0x1b,
	0x35, 0x65, 0x06, 0xd3, 0x44, 0x59, 0xe7, 0xac, 0x61, 0x5e, 0x65, 0xba, 0x38, 0xfd, 0x34, 0xbb,
	0x35, 0xe4, 0xa5, 0xcc, 0x39, 0xd0, 0xc0, 0xe1, 0xac, 0x05, 0x7f, 0x95, 0xd8, 0xce, 0xea, 0x32,
	0xfc, 0x2f, 0x59, 0x9d, 0x2a, 0x20, 0x42, 0xb9, 0x2a, 0x9e, 0xde, 0x21, 0xb4, 0xaa, 0x74, 0x60,
	0x41, 0xa5, 0x9d, 0xdb, 0x62, 0xbd, 0x79, 0xb1, 0x3d, 0xb8, 0x30, 0x5b, 0x2c, 0xe8, 0x3c, 0xb7,
	0xf4, 0x9c, 0x6d, 0xc5, 0x0d, 0x65, 0xfe, 0x24, 0x10, 0x61, 0x6f, 0x69, 0x2f, 0x35, 0xaa, 0xea,
	0xc2, 0x26, 0xe2, 0xfa, 0xdc, 0xf6, 0x13, 0xeb, 0x2a, 0x07, 0xbf, 0x55, 0x16, 0xd7, 0x8b, 0xa4,
	0x0a, 0xef, 0x40, 0xc9, 0x9a, 0x58, 0xcd, 0x29, 0x9e, 0x58, 0xad, 0xba, 0xdd, 0xf9, 0x48, 0xc0,
	0xc4, 0x1a, 0x72, 0x86, 0x6f, 0x7b, 0xe4, 0xf9, 0x96, 0x34, 0x18, 0x84, 0x83, 0xd1, 0x99, 0x2c,
	0x6b, 0x70, 0x82, 0x02, 0x7a, 0x28, 0xdd, 0x63, 0x06, 0xaf, 0x46, 0x8e, 0x0f, 0x3a, 0x27, 0x32,
	0x06, 0xab, 0x15, 0x96, 0x7f, 0xe1, 0x03, 0xa2, 0x38, 0x14, 0x6b, 0xce, 0x28, 0xdc, 0x87, 0x9e,
	0xe8, 0x42, 0x82, 0xda, 0x72, 0xda, 0xe3, 0x2e, 0xa5, 0xcd, 0x5d, 0x94, 0xb3, 0xb6, 0x0d, 0xbe,
	0xd2, 0x6f, 0xed, 0xf0, 0xb2, 0xb9, 0x8f, 0x15, 0xd1, 0x06, 0xdd, 0xfc, 0xc7, 0xac, 0x82, 0xea,
	0x25, 0x2f, 0x6d, 0x1c, 0x6c, 0xaf, 0x94, 0xf6, 0x7c, 0xc6, 0x9a, 0x24, 0x52, 0x10, 0x8b, 0x19,
	0xe9, 0x24, 0xf7, 0x4c, 0xb2, 0x1c, 0x33, 0xbf, 0x6b, 0xf2, 0xcc, 0xc6, 0xbf, 0x2c, 0x67, 0x85,
	0x8c, 0x8a, 0x2a, 0x6d, 0xbb, 0x1b, 0x7d, 0x03, 0x65, 0x0c, 0xe8, 0xe4, 0x0d, 0xc7, 0xd8, 0x3a,
	0x64, 0xd1, 0x4a, 0x49, 0xbc, 0x2f, 0x2f, 0x25, 0x00, 0x9b, 0x54, 0x42, 0x83, 0x95, 0x8d, 0xbc,
	0xb6, 0xf6, 0x09, 0x07, 0x27, 0xe8, 0xd4, 0xf0, 0x90, 0x90, 0x77, 0x7a, 0xfe, 0x2e, 0xdb, 0xcc,
	0x3b, 0xf2, 0x72, 0x3a, 0x8d, 0x6e, 0xba, 0xca, 0x18, 0x1d, 0xcd, 0x24, 0x79, 0xa8, 0x17, 0x7c,
	0xb8, 0xd0, 0x07, 0x6e, 0x43, 0x7d, 0xe0, 0x5f, 0xab, 0x0f, 0x2f, 0xf8, 0x11, 0xde, 0x43, 0xd7,
	0x09, 0x31, 0xc6, 0x6e, 0xe1, 0xa3, 0x46, 0x1f, 0x57, 0x75, 0x8b, 0x5e, 0xfb, 0xa7, 0xcc, 0xc3,
	0x6a, 0xdd, 0x7f, 0x43, 0xdb, 0x2b, 0x7b, 0xd1, 0x7f, 0xaf, 0x5f, 0xfa, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xea, 0xe5, 0x36, 0x24, 0xad, 0x09, 0x00, 0x00,
}
